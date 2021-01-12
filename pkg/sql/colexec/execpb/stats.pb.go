// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/colexec/execpb/stats.proto

package execpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// VectorizedStats represents the stats collected from an operator.
type VectorizedStats struct {
	ID         int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NumBatches int64         `protobuf:"varint,2,opt,name=num_batches,json=numBatches,proto3" json:"num_batches,omitempty"`
	NumTuples  int64         `protobuf:"varint,3,opt,name=num_tuples,json=numTuples,proto3" json:"num_tuples,omitempty"`
	Time       time.Duration `protobuf:"bytes,4,opt,name=time,proto3,stdduration" json:"time"`
	// stall indicates whether stall time or execution time is being tracked.
	Stall            bool  `protobuf:"varint,5,opt,name=stall,proto3" json:"stall,omitempty"`
	MaxAllocatedMem  int64 `protobuf:"varint,6,opt,name=max_allocated_mem,json=maxAllocatedMem,proto3" json:"max_allocated_mem,omitempty"`
	MaxAllocatedDisk int64 `protobuf:"varint,7,opt,name=max_allocated_disk,json=maxAllocatedDisk,proto3" json:"max_allocated_disk,omitempty"`
}

func (m *VectorizedStats) Reset()         { *m = VectorizedStats{} }
func (m *VectorizedStats) String() string { return proto.CompactTextString(m) }
func (*VectorizedStats) ProtoMessage()    {}
func (*VectorizedStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_stats_fe732b30c56248d9, []int{0}
}
func (m *VectorizedStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VectorizedStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *VectorizedStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VectorizedStats.Merge(dst, src)
}
func (m *VectorizedStats) XXX_Size() int {
	return m.Size()
}
func (m *VectorizedStats) XXX_DiscardUnknown() {
	xxx_messageInfo_VectorizedStats.DiscardUnknown(m)
}

var xxx_messageInfo_VectorizedStats proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VectorizedStats)(nil), "cockroach.sql.execpb.VectorizedStats")
}
func (m *VectorizedStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VectorizedStats) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.ID))
	}
	if m.NumBatches != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.NumBatches))
	}
	if m.NumTuples != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.NumTuples))
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintStats(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.Time)))
	n1, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Time, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Stall {
		dAtA[i] = 0x28
		i++
		if m.Stall {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.MaxAllocatedMem != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.MaxAllocatedMem))
	}
	if m.MaxAllocatedDisk != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.MaxAllocatedDisk))
	}
	return i, nil
}

func encodeVarintStats(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VectorizedStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovStats(uint64(m.ID))
	}
	if m.NumBatches != 0 {
		n += 1 + sovStats(uint64(m.NumBatches))
	}
	if m.NumTuples != 0 {
		n += 1 + sovStats(uint64(m.NumTuples))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Time)
	n += 1 + l + sovStats(uint64(l))
	if m.Stall {
		n += 2
	}
	if m.MaxAllocatedMem != 0 {
		n += 1 + sovStats(uint64(m.MaxAllocatedMem))
	}
	if m.MaxAllocatedDisk != 0 {
		n += 1 + sovStats(uint64(m.MaxAllocatedDisk))
	}
	return n
}

func sovStats(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStats(x uint64) (n int) {
	return sovStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VectorizedStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VectorizedStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VectorizedStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBatches", wireType)
			}
			m.NumBatches = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBatches |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumTuples", wireType)
			}
			m.NumTuples = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumTuples |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stall", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Stall = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAllocatedMem", wireType)
			}
			m.MaxAllocatedMem = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAllocatedMem |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAllocatedDisk", wireType)
			}
			m.MaxAllocatedDisk = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAllocatedDisk |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStats
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStats
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStats
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStats(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStats = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStats   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("sql/colexec/execpb/stats.proto", fileDescriptor_stats_fe732b30c56248d9)
}

var fileDescriptor_stats_fe732b30c56248d9 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x6b, 0xdb, 0x30,
	0x18, 0x86, 0x2d, 0x27, 0xf1, 0x32, 0xe5, 0x90, 0x4d, 0x84, 0xe1, 0x05, 0x26, 0x9b, 0x9d, 0xcc,
	0x18, 0x36, 0x6c, 0x87, 0x9d, 0x67, 0x72, 0xd9, 0x61, 0x17, 0xb7, 0xf4, 0xd0, 0x8b, 0x91, 0x65,
	0xd5, 0x31, 0x91, 0xac, 0xc4, 0x92, 0x21, 0xf4, 0x57, 0x94, 0x9e, 0xfa, 0x93, 0x72, 0xcc, 0x31,
	0xa7, 0xb4, 0x75, 0xfe, 0x48, 0xb1, 0xdc, 0x40, 0x7b, 0x11, 0x7a, 0xdf, 0xef, 0x11, 0xdf, 0x83,
	0x20, 0x56, 0x1b, 0x1e, 0x51, 0xc9, 0xd9, 0x96, 0xd1, 0xa8, 0x3b, 0xd6, 0x59, 0xa4, 0x34, 0xd1,
	0x2a, 0x5c, 0xd7, 0x52, 0x4b, 0x34, 0xa3, 0x92, 0xae, 0x6a, 0x49, 0xe8, 0x32, 0x54, 0x1b, 0x1e,
	0xf6, 0xc4, 0x7c, 0x56, 0xc8, 0x42, 0x1a, 0x20, 0xea, 0x6e, 0x3d, 0x3b, 0xc7, 0x85, 0x94, 0x05,
	0x67, 0x91, 0x49, 0x59, 0x73, 0x13, 0xe5, 0x4d, 0x4d, 0x74, 0x29, 0xab, 0x7e, 0xfe, 0xfd, 0xde,
	0x86, 0xd3, 0x2b, 0x46, 0xb5, 0xac, 0xcb, 0x5b, 0x96, 0x5f, 0x74, 0x5b, 0xd0, 0x17, 0x68, 0x97,
	0xb9, 0x0b, 0x7c, 0x10, 0x8c, 0x62, 0xa7, 0x3d, 0x7a, 0xf6, 0xbf, 0x45, 0x62, 0x97, 0x39, 0xf2,
	0xe0, 0xa4, 0x6a, 0x44, 0x9a, 0x11, 0x4d, 0x97, 0x4c, 0xb9, 0xb6, 0x0f, 0x82, 0x41, 0x02, 0xab,
	0x46, 0xc4, 0x7d, 0x83, 0xbe, 0xc1, 0x2e, 0xa5, 0xba, 0x59, 0x73, 0xa6, 0xdc, 0x81, 0x99, 0x7f,
	0xac, 0x1a, 0x71, 0x69, 0x0a, 0xf4, 0x07, 0x0e, 0x75, 0x29, 0x98, 0x3b, 0xf4, 0x41, 0x30, 0xf9,
	0xf5, 0x35, 0xec, 0xd5, 0xc2, 0xb3, 0x5a, 0xb8, 0x78, 0x55, 0x8b, 0xc7, 0xbb, 0xa3, 0x67, 0x3d,
	0x3c, 0x7a, 0x20, 0x31, 0x0f, 0xd0, 0x0c, 0x8e, 0x94, 0x26, 0x9c, 0xbb, 0x23, 0x1f, 0x04, 0xe3,
	0xa4, 0x0f, 0xe8, 0x07, 0xfc, 0x2c, 0xc8, 0x36, 0x25, 0x9c, 0x4b, 0x4a, 0x34, 0xcb, 0x53, 0xc1,
	0x84, 0xeb, 0x98, 0xa5, 0x53, 0x41, 0xb6, 0x7f, 0xcf, 0xfd, 0x7f, 0x26, 0xd0, 0x4f, 0x88, 0xde,
	0xb3, 0x79, 0xa9, 0x56, 0xee, 0x07, 0x03, 0x7f, 0x7a, 0x0b, 0x2f, 0x4a, 0xb5, 0x8a, 0x83, 0xdd,
	0x33, 0xb6, 0x76, 0x2d, 0x06, 0xfb, 0x16, 0x83, 0x43, 0x8b, 0xc1, 0x53, 0x8b, 0xc1, 0xdd, 0x09,
	0x5b, 0xfb, 0x13, 0xb6, 0x0e, 0x27, 0x6c, 0x5d, 0x3b, 0xfd, 0xa7, 0x67, 0x8e, 0x91, 0xff, 0xfd,
	0x12, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xaf, 0xd1, 0x5f, 0xb3, 0x01, 0x00, 0x00,
}
