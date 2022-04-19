// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the SSH public keys associated with the specified IAM
// user. If none exists, the operation returns an empty list. The SSH public keys
// returned by this operation are used only for authenticating the IAM user to an
// CodeCommit repository. For more information about using SSH keys to authenticate
// to an CodeCommit repository, see Set up CodeCommit for SSH connections
// (https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html)
// in the CodeCommit User Guide. Although each user is limited to a small number of
// keys, you can still paginate the results using the MaxItems and Marker
// parameters.
func (c *Client) ListSSHPublicKeys(ctx context.Context, params *ListSSHPublicKeysInput, optFns ...func(*Options)) (*ListSSHPublicKeysOutput, error) {
	if params == nil {
		params = &ListSSHPublicKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSSHPublicKeys", params, optFns, c.addOperationListSSHPublicKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSSHPublicKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSSHPublicKeysInput struct {

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32

	// The name of the IAM user to list SSH public keys for. If none is specified, the
	// UserName field is determined implicitly based on the Amazon Web Services access
	// key used to sign the request. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	UserName *string

	noSmithyDocumentSerde
}

// Contains the response to a successful ListSSHPublicKeys request.
type ListSSHPublicKeysOutput struct {

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you receive
	// all your results.
	IsTruncated bool

	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// A list of the SSH public keys assigned to IAM user.
	SSHPublicKeys []types.SSHPublicKeyMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSSHPublicKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListSSHPublicKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListSSHPublicKeys{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSSHPublicKeys(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListSSHPublicKeysAPIClient is a client that implements the ListSSHPublicKeys
// operation.
type ListSSHPublicKeysAPIClient interface {
	ListSSHPublicKeys(context.Context, *ListSSHPublicKeysInput, ...func(*Options)) (*ListSSHPublicKeysOutput, error)
}

var _ ListSSHPublicKeysAPIClient = (*Client)(nil)

// ListSSHPublicKeysPaginatorOptions is the paginator options for ListSSHPublicKeys
type ListSSHPublicKeysPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSSHPublicKeysPaginator is a paginator for ListSSHPublicKeys
type ListSSHPublicKeysPaginator struct {
	options   ListSSHPublicKeysPaginatorOptions
	client    ListSSHPublicKeysAPIClient
	params    *ListSSHPublicKeysInput
	nextToken *string
	firstPage bool
}

// NewListSSHPublicKeysPaginator returns a new ListSSHPublicKeysPaginator
func NewListSSHPublicKeysPaginator(client ListSSHPublicKeysAPIClient, params *ListSSHPublicKeysInput, optFns ...func(*ListSSHPublicKeysPaginatorOptions)) *ListSSHPublicKeysPaginator {
	if params == nil {
		params = &ListSSHPublicKeysInput{}
	}

	options := ListSSHPublicKeysPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSSHPublicKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSSHPublicKeysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSSHPublicKeys page.
func (p *ListSSHPublicKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSSHPublicKeysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListSSHPublicKeys(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListSSHPublicKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListSSHPublicKeys",
	}
}
