// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of control panels in an account or in a cluster.
func (c *Client) ListControlPanels(ctx context.Context, params *ListControlPanelsInput, optFns ...func(*Options)) (*ListControlPanelsOutput, error) {
	if params == nil {
		params = &ListControlPanelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListControlPanels", params, optFns, c.addOperationListControlPanelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListControlPanelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListControlPanelsInput struct {

	// The Amazon Resource Name (ARN) of a cluster.
	ClusterArn *string

	// The number of objects that you want to return with this call.
	MaxResults int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type ListControlPanelsOutput struct {

	// The result of a successful ListControlPanel request.
	ControlPanels []types.ControlPanel

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListControlPanelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListControlPanels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListControlPanels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListControlPanels(options.Region), middleware.Before); err != nil {
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

// ListControlPanelsAPIClient is a client that implements the ListControlPanels
// operation.
type ListControlPanelsAPIClient interface {
	ListControlPanels(context.Context, *ListControlPanelsInput, ...func(*Options)) (*ListControlPanelsOutput, error)
}

var _ ListControlPanelsAPIClient = (*Client)(nil)

// ListControlPanelsPaginatorOptions is the paginator options for ListControlPanels
type ListControlPanelsPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListControlPanelsPaginator is a paginator for ListControlPanels
type ListControlPanelsPaginator struct {
	options   ListControlPanelsPaginatorOptions
	client    ListControlPanelsAPIClient
	params    *ListControlPanelsInput
	nextToken *string
	firstPage bool
}

// NewListControlPanelsPaginator returns a new ListControlPanelsPaginator
func NewListControlPanelsPaginator(client ListControlPanelsAPIClient, params *ListControlPanelsInput, optFns ...func(*ListControlPanelsPaginatorOptions)) *ListControlPanelsPaginator {
	if params == nil {
		params = &ListControlPanelsInput{}
	}

	options := ListControlPanelsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListControlPanelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListControlPanelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListControlPanels page.
func (p *ListControlPanelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListControlPanelsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListControlPanels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListControlPanels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-control-config",
		OperationName: "ListControlPanels",
	}
}
