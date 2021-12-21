// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the current list of Amazon Web Services services and a list of service
// categories for each service. You then use service names and categories in your
// CreateCase requests. Each Amazon Web Services service has its own set of
// categories. The service codes and category codes correspond to the values that
// appear in the Service and Category lists on the Amazon Web Services Support
// Center Create Case (https://console.aws.amazon.com/support/home#/case/create)
// page. The values in those fields don't necessarily match the service codes and
// categories returned by the DescribeServices operation. Always use the service
// codes and categories that the DescribeServices operation returns, so that you
// have the most recent set of service and category codes.
//
// * You must have a
// Business, Enterprise On-Ramp, or Enterprise Support plan to use the Amazon Web
// Services Support API.
//
// * If you call the Amazon Web Services Support API from an
// account that does not have a Business, Enterprise On-Ramp, or Enterprise Support
// plan, the SubscriptionRequiredException error message appears. For information
// about changing your support plan, see Amazon Web Services Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) DescribeServices(ctx context.Context, params *DescribeServicesInput, optFns ...func(*Options)) (*DescribeServicesOutput, error) {
	if params == nil {
		params = &DescribeServicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeServices", params, optFns, c.addOperationDescribeServicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeServicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeServicesInput struct {

	// The ISO 639-1 code for the language in which Amazon Web Services provides
	// support. Amazon Web Services Support currently supports English ("en") and
	// Japanese ("ja"). Language parameters must be passed explicitly for operations
	// that take them.
	Language *string

	// A JSON-formatted list of service codes available for Amazon Web Services
	// services.
	ServiceCodeList []string

	noSmithyDocumentSerde
}

// The list of Amazon Web Services services returned by the DescribeServices
// operation.
type DescribeServicesOutput struct {

	// A JSON-formatted list of Amazon Web Services services.
	Services []types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeServicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeServices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeServices{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServices(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeServices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeServices",
	}
}
