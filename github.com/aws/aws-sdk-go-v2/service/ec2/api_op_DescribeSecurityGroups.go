// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"strconv"
	"time"
)

// Describes the specified security groups or all of your security groups. A
// security group is for use with instances either in the EC2-Classic platform or
// in a specific VPC. For more information, see Amazon EC2 security groups
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html)
// in the Amazon Elastic Compute Cloud User Guide and Security groups for your VPC
// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) DescribeSecurityGroups(ctx context.Context, params *DescribeSecurityGroupsInput, optFns ...func(*Options)) (*DescribeSecurityGroupsOutput, error) {
	if params == nil {
		params = &DescribeSecurityGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSecurityGroups", params, optFns, c.addOperationDescribeSecurityGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSecurityGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSecurityGroupsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters. If using multiple filters for rules, the results include security
	// groups for which any combination of rules - not necessarily a single rule -
	// match all filters.
	//
	// * description - The description of the security group.
	//
	// *
	// egress.ip-permission.cidr - An IPv4 CIDR block for an outbound security group
	// rule.
	//
	// * egress.ip-permission.from-port - For an outbound rule, the start of
	// port range for the TCP and UDP protocols, or an ICMP type number.
	//
	// *
	// egress.ip-permission.group-id - The ID of a security group that has been
	// referenced in an outbound security group rule.
	//
	// *
	// egress.ip-permission.group-name - The name of a security group that is
	// referenced in an outbound security group rule.
	//
	// * egress.ip-permission.ipv6-cidr
	// - An IPv6 CIDR block for an outbound security group rule.
	//
	// *
	// egress.ip-permission.prefix-list-id - The ID of a prefix list to which a
	// security group rule allows outbound access.
	//
	// * egress.ip-permission.protocol -
	// The IP protocol for an outbound security group rule (tcp | udp | icmp, a
	// protocol number, or -1 for all protocols).
	//
	// * egress.ip-permission.to-port - For
	// an outbound rule, the end of port range for the TCP and UDP protocols, or an
	// ICMP code.
	//
	// * egress.ip-permission.user-id - The ID of an Amazon Web Services
	// account that has been referenced in an outbound security group rule.
	//
	// * group-id
	// - The ID of the security group.
	//
	// * group-name - The name of the security
	// group.
	//
	// * ip-permission.cidr - An IPv4 CIDR block for an inbound security group
	// rule.
	//
	// * ip-permission.from-port - For an inbound rule, the start of port range
	// for the TCP and UDP protocols, or an ICMP type number.
	//
	// * ip-permission.group-id
	// - The ID of a security group that has been referenced in an inbound security
	// group rule.
	//
	// * ip-permission.group-name - The name of a security group that is
	// referenced in an inbound security group rule.
	//
	// * ip-permission.ipv6-cidr - An
	// IPv6 CIDR block for an inbound security group rule.
	//
	// *
	// ip-permission.prefix-list-id - The ID of a prefix list from which a security
	// group rule allows inbound access.
	//
	// * ip-permission.protocol - The IP protocol
	// for an inbound security group rule (tcp | udp | icmp, a protocol number, or -1
	// for all protocols).
	//
	// * ip-permission.to-port - For an inbound rule, the end of
	// port range for the TCP and UDP protocols, or an ICMP code.
	//
	// *
	// ip-permission.user-id - The ID of an Amazon Web Services account that has been
	// referenced in an inbound security group rule.
	//
	// * owner-id - The Amazon Web
	// Services account ID of the owner of the security group.
	//
	// * tag: - The key/value
	// combination of a tag assigned to the resource. Use the tag key in the filter
	// name and the tag value as the filter value. For example, to find all resources
	// that have a tag with the key Owner and the value TeamA, specify tag:Owner for
	// the filter name and TeamA for the filter value.
	//
	// * tag-key - The key of a tag
	// assigned to the resource. Use this filter to find all resources assigned a tag
	// with a specific key, regardless of the tag value.
	//
	// * vpc-id - The ID of the VPC
	// specified when the security group was created.
	Filters []types.Filter

	// The IDs of the security groups. Required for security groups in a nondefault
	// VPC. Default: Describes all of your security groups.
	GroupIds []string

	// [EC2-Classic and default VPC only] The names of the security groups. You can
	// specify either the security group name or the security group ID. For security
	// groups in a nondefault VPC, use the group-name filter to describe security
	// groups by name. Default: Describes all of your security groups.
	GroupNames []string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another request with the returned NextToken value. This
	// value can be between 5 and 1000. If this parameter is not specified, then all
	// results are returned.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeSecurityGroupsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the security groups.
	SecurityGroups []types.SecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSecurityGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSecurityGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSecurityGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSecurityGroups(options.Region), middleware.Before); err != nil {
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

// DescribeSecurityGroupsAPIClient is a client that implements the
// DescribeSecurityGroups operation.
type DescribeSecurityGroupsAPIClient interface {
	DescribeSecurityGroups(context.Context, *DescribeSecurityGroupsInput, ...func(*Options)) (*DescribeSecurityGroupsOutput, error)
}

var _ DescribeSecurityGroupsAPIClient = (*Client)(nil)

// DescribeSecurityGroupsPaginatorOptions is the paginator options for
// DescribeSecurityGroups
type DescribeSecurityGroupsPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another request with the returned NextToken value. This
	// value can be between 5 and 1000. If this parameter is not specified, then all
	// results are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSecurityGroupsPaginator is a paginator for DescribeSecurityGroups
type DescribeSecurityGroupsPaginator struct {
	options   DescribeSecurityGroupsPaginatorOptions
	client    DescribeSecurityGroupsAPIClient
	params    *DescribeSecurityGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeSecurityGroupsPaginator returns a new DescribeSecurityGroupsPaginator
func NewDescribeSecurityGroupsPaginator(client DescribeSecurityGroupsAPIClient, params *DescribeSecurityGroupsInput, optFns ...func(*DescribeSecurityGroupsPaginatorOptions)) *DescribeSecurityGroupsPaginator {
	if params == nil {
		params = &DescribeSecurityGroupsInput{}
	}

	options := DescribeSecurityGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSecurityGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSecurityGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSecurityGroups page.
func (p *DescribeSecurityGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSecurityGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeSecurityGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// SecurityGroupExistsWaiterOptions are waiter options for
// SecurityGroupExistsWaiter
type SecurityGroupExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// SecurityGroupExistsWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, SecurityGroupExistsWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeSecurityGroupsInput, *DescribeSecurityGroupsOutput, error) (bool, error)
}

// SecurityGroupExistsWaiter defines the waiters for SecurityGroupExists
type SecurityGroupExistsWaiter struct {
	client DescribeSecurityGroupsAPIClient

	options SecurityGroupExistsWaiterOptions
}

// NewSecurityGroupExistsWaiter constructs a SecurityGroupExistsWaiter.
func NewSecurityGroupExistsWaiter(client DescribeSecurityGroupsAPIClient, optFns ...func(*SecurityGroupExistsWaiterOptions)) *SecurityGroupExistsWaiter {
	options := SecurityGroupExistsWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = securityGroupExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &SecurityGroupExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for SecurityGroupExists waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *SecurityGroupExistsWaiter) Wait(ctx context.Context, params *DescribeSecurityGroupsInput, maxWaitDur time.Duration, optFns ...func(*SecurityGroupExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for SecurityGroupExists waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *SecurityGroupExistsWaiter) WaitForOutput(ctx context.Context, params *DescribeSecurityGroupsInput, maxWaitDur time.Duration, optFns ...func(*SecurityGroupExistsWaiterOptions)) (*DescribeSecurityGroupsOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeSecurityGroups(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for SecurityGroupExists waiter")
}

func securityGroupExistsStateRetryable(ctx context.Context, input *DescribeSecurityGroupsInput, output *DescribeSecurityGroupsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("length(SecurityGroups[].GroupId) > `0`", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, nil
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "InvalidGroup.NotFound" == apiErr.ErrorCode() {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeSecurityGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSecurityGroups",
	}
}
