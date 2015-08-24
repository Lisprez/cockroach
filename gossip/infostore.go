// Copyright 2014 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Spencer Kimball (spencer.kimball@gmail.com)

package gossip

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/cockroachdb/cockroach/proto"
	"github.com/cockroachdb/cockroach/util"
	"github.com/cockroachdb/cockroach/util/log"
)

// callback holds regexp pattern match and GossipCallback method.
type callback struct {
	pattern *regexp.Regexp
	method  Callback
}

// infoStore objects manage maps of Info objects. They maintain a
// sequence number generator which they use to allocate new info
// objects.
//
// infoStores can be queried for incremental updates occurring since a
// specified sequence number.
//
// infoStores can be combined using deltas from peer nodes.
//
// infoStores are not thread safe.
type infoStore struct {
	Infos     infoMap             `json:"infos,omitempty"` // Map from key to info
	NodeID    proto.NodeID        `json:"-"`               // Owning node's ID
	NodeAddr  util.UnresolvedAddr `json:"-"`               // Address of node owning this info store: "host:port"
	MaxSeq    int64               `json:"-"`               // Maximum sequence number inserted
	seqGen    int64               // Sequence generator incremented each time info is added
	callbacks []callback
}

// monotonicUnixNano returns a monotonically increasing value for
// nanoseconds in Unix time. Since equal times are ignored with
// updates to infos, we're careful to avoid incorrectly ignoring a
// newly created value in the event one is created within the same
// nanosecond. Really unlikely except for the case of unittests, but
// better safe than sorry.
func monotonicUnixNano() int64 {
	monoTimeMu.Lock()
	defer monoTimeMu.Unlock()

	now := time.Now().UnixNano()
	if now <= lastTime {
		now = lastTime + 1
	}
	lastTime = now
	return now
}

// String returns a string representation of an infostore.
func (is *infoStore) String() string {
	buf := bytes.Buffer{}
	if infoCount := len(is.Infos); infoCount > 0 {
		buf.WriteString(fmt.Sprintf("infostore with %d info(s): ", infoCount))
	} else {
		return "infostore (empty)"
	}

	prepend := ""

	if err := is.visitInfos(func(i *info) error {
		str := fmt.Sprintf("%sinfo %q: %+v", prepend, i.Key, i.Val)
		prepend = ", "
		_, err := buf.WriteString(str)
		return err
	}); err != nil {
		log.Errorf("failed to properly construct string representation of infoStore: %s", err)
	}
	return buf.String()
}

var (
	monoTimeMu sync.Mutex
	lastTime   int64
)

// newInfoStore allocates and returns a new infoStore.
// "NodeAddr" is the address of the node owning the infostore
// in "host:port" format.
func newInfoStore(nodeID proto.NodeID, nodeAddr util.UnresolvedAddr) *infoStore {
	return &infoStore{
		Infos:    infoMap{},
		NodeID:   nodeID,
		NodeAddr: nodeAddr,
	}
}

// newInfo allocates and returns a new info object using specified key,
// value, and time-to-live.
func (is *infoStore) newInfo(key string, val interface{}, ttl time.Duration) *info {
	is.seqGen++
	now := monotonicUnixNano()
	ttlStamp := now + int64(ttl)
	if ttl == 0 {
		ttlStamp = math.MaxInt64
	}
	return &info{
		Key:       key,
		Val:       val,
		Timestamp: now,
		TTLStamp:  ttlStamp,
		NodeID:    is.NodeID,
		peerID:    is.NodeID,
		seq:       is.seqGen,
	}
}

// getInfo returns an info object by key or nil if it doesn't exist.
func (is *infoStore) getInfo(key string) *info {
	if info, ok := is.Infos[key]; ok {
		// Check TTL and discard if too old.
		if info.expired(time.Now().UnixNano()) {
			delete(is.Infos, key)
			return nil
		}
		return info
	}
	return nil
}

// addInfo adds or updates an info in the infos map.
//
// Returns nil if info was added; error otherwise.
func (is *infoStore) addInfo(i *info) error {
	// Only replace an existing info if new timestamp is greater, or if
	// timestamps are equal, but new hops is smaller.
	var contentsChanged bool
	if existingInfo, ok := is.Infos[i.Key]; ok {
		if i.Timestamp < existingInfo.Timestamp ||
			(i.Timestamp == existingInfo.Timestamp && i.Hops >= existingInfo.Hops) {
			return util.Errorf("info %+v older than current info %+v", i, existingInfo)
		}
		contentsChanged = !reflect.DeepEqual(existingInfo.Val, i.Val)
	} else {
		// No preexisting info means contentsChanged is true.
		contentsChanged = true
	}
	// Update info map.
	is.Infos[i.Key] = i
	if i.seq > is.MaxSeq {
		is.MaxSeq = i.seq
	}
	is.processCallbacks(i.Key, contentsChanged, i.Val)
	return nil
}

// maxHops returns the maximum hops across all infos in the store.
// This is the maximum number of gossip exchanges between any
// originator and this node.
func (is *infoStore) maxHops() uint32 {
	var maxHops uint32
	// will never error because `return nil` below
	_ = is.visitInfos(func(i *info) error {
		if i.Hops > maxHops {
			maxHops = i.Hops
		}
		return nil
	})
	return maxHops
}

// registerCallback compiles a regexp for pattern and adds it to
// the callbacks slice.
func (is *infoStore) registerCallback(pattern string, method Callback) {
	re := regexp.MustCompile(pattern)
	is.callbacks = append(is.callbacks, callback{pattern: re, method: method})
	var infos []*info
	if err := is.visitInfos(func(i *info) error {
		if re.MatchString(i.Key) {
			infos = append(infos, i)
		}
		return nil
	}); err != nil {
		log.Errorf("failed to properly run registered callback while visiting pre-existing info: %s", err)
	}
	// Run callbacks in a goroutine to avoid mutex reentry.
	go func() {
		for _, i := range infos {
			method(i.Key, true /* contentsChanged */, i.Val)
		}
	}()
}

// processCallbacks processes callbacks for the specified key by
// matching callback regular expression against the key and invoking
// the corresponding callback method on a match.
func (is *infoStore) processCallbacks(key string, contentsChanged bool, content interface{}) {
	var matches []callback
	for _, cb := range is.callbacks {
		if cb.pattern.MatchString(key) {
			matches = append(matches, cb)
		}
	}
	// Run callbacks in a goroutine to avoid mutex reentry.
	go func() {
		for _, cb := range matches {
			cb.method(key, contentsChanged, content)
		}
	}()
}

// visitInfos implements a visitor pattern to run the visitInfo
// function against each info in turn. Be sure to skip over any expired
// infos.
func (is *infoStore) visitInfos(visitInfo func(*info) error) error {
	now := time.Now().UnixNano()

	if visitInfo != nil {
		for _, i := range is.Infos {
			if i.expired(now) {
				delete(is.Infos, i.Key)
				continue
			}
			if err := visitInfo(i); err != nil {
				return err
			}
		}
	}

	return nil
}

// combine combines an incremental delta with the current infoStore.
// The sequence numbers on all info objects are reset using the info
// store's sequence generator. All hop distances on infos are
// incremented to indicate they've arrived from an external source.
// Returns the count of "fresh" infos in the provided delta.
func (is *infoStore) combine(delta *infoStore) int {
	var freshCount int
	if err := delta.visitInfos(func(i *info) error {
		is.seqGen++
		i.seq = is.seqGen
		i.Hops++
		i.peerID = delta.NodeID
		// Errors from addInfo here are not a problem; they simply
		// indicate that the data in *is is newer than in *delta.
		if err := is.addInfo(i); err == nil {
			freshCount++
		}
		return nil
	}); err != nil {
		log.Errorf("failed to properly combine infoStores: %s", err)
	}
	return freshCount
}

// delta returns an incremental delta of infos added to the info store
// since (not including) the specified sequence number. These deltas
// are intended for efficiently updating peer nodes. Any infos passed
// from node requesting delta are ignored.
//
// Returns nil if there are no deltas.
func (is *infoStore) delta(nodeID proto.NodeID, seq int64) *infoStore {
	if seq >= is.MaxSeq {
		return nil
	}

	delta := newInfoStore(is.NodeID, is.NodeAddr)

	// Compute delta of infos.
	if err := is.visitInfos(func(i *info) error {
		if i.isFresh(nodeID, seq) {
			return delta.addInfo(i)
		}
		return nil
	}); err != nil {
		log.Errorf("failed to properly create delta infoStore: %s", err)
	}

	delta.MaxSeq = is.MaxSeq
	return delta
}

// distant returns a nodeSet for gossip peers which originated infos
// with info.Hops > maxHops.
func (is *infoStore) distant(maxHops uint32) *nodeSet {
	ns := newNodeSet(0)
	// will never error because `return nil` below
	_ = is.visitInfos(func(i *info) error {
		if i.Hops > maxHops {
			ns.addNode(i.NodeID)
		}
		return nil
	})
	return ns
}

// leastUseful determines which node ID from amongst the set is
// currently contributing the least. Returns 0 if nodes is empty.
func (is *infoStore) leastUseful(nodes *nodeSet) proto.NodeID {
	contrib := make(map[proto.NodeID]int, nodes.len())
	for node := range nodes.nodes {
		contrib[node] = 0
	}
	// will never error because `return nil` below
	_ = is.visitInfos(func(i *info) error {
		contrib[i.peerID]++
		return nil
	})

	least := math.MaxInt32
	var leastNode proto.NodeID
	for id, count := range contrib {
		if nodes.hasNode(id) {
			if count < least {
				least = count
				leastNode = id
			}
		}
	}
	return leastNode
}
