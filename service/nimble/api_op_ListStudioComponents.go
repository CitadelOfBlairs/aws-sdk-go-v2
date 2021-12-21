// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the StudioComponents in a studio.
func (c *Client) ListStudioComponents(ctx context.Context, params *ListStudioComponentsInput, optFns ...func(*Options)) (*ListStudioComponentsOutput, error) {
	if params == nil {
		params = &ListStudioComponentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStudioComponents", params, optFns, c.addOperationListStudioComponentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStudioComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStudioComponentsInput struct {

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// The max number of results to return in the response.
	MaxResults int32

	// The token to request the next page of results.
	NextToken *string

	// Filters the request to studio components that are in one of the given states.
	States []string

	// Filters the request to studio components that are of one of the given types.
	Types []string

	noSmithyDocumentSerde
}

type ListStudioComponentsOutput struct {

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// A collection of studio components.
	StudioComponents []types.StudioComponent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStudioComponentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStudioComponents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStudioComponents{}, middleware.After)
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
	if err = addOpListStudioComponentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStudioComponents(options.Region), middleware.Before); err != nil {
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

// ListStudioComponentsAPIClient is a client that implements the
// ListStudioComponents operation.
type ListStudioComponentsAPIClient interface {
	ListStudioComponents(context.Context, *ListStudioComponentsInput, ...func(*Options)) (*ListStudioComponentsOutput, error)
}

var _ ListStudioComponentsAPIClient = (*Client)(nil)

// ListStudioComponentsPaginatorOptions is the paginator options for
// ListStudioComponents
type ListStudioComponentsPaginatorOptions struct {
	// The max number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStudioComponentsPaginator is a paginator for ListStudioComponents
type ListStudioComponentsPaginator struct {
	options   ListStudioComponentsPaginatorOptions
	client    ListStudioComponentsAPIClient
	params    *ListStudioComponentsInput
	nextToken *string
	firstPage bool
}

// NewListStudioComponentsPaginator returns a new ListStudioComponentsPaginator
func NewListStudioComponentsPaginator(client ListStudioComponentsAPIClient, params *ListStudioComponentsInput, optFns ...func(*ListStudioComponentsPaginatorOptions)) *ListStudioComponentsPaginator {
	if params == nil {
		params = &ListStudioComponentsInput{}
	}

	options := ListStudioComponentsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStudioComponentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStudioComponentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStudioComponents page.
func (p *ListStudioComponentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStudioComponentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListStudioComponents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStudioComponents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "ListStudioComponents",
	}
}
