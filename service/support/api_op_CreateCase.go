// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a case in the Amazon Web Services Support Center. This operation is
// similar to how you create a case in the Amazon Web Services Support Center
// Create Case (https://console.aws.amazon.com/support/home#/case/create) page. The
// Amazon Web Services Support API doesn't support requesting service limit
// increases. You can submit a service limit increase in the following ways:
//
// *
// Submit a request from the Amazon Web Services Support Center Create Case
// (https://console.aws.amazon.com/support/home#/case/create) page.
//
// * Use the
// Service Quotas RequestServiceQuotaIncrease
// (https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_RequestServiceQuotaIncrease.html)
// operation.
//
// A successful CreateCase request returns an Amazon Web Services
// Support case number. You can use the DescribeCases operation and specify the
// case number to get existing Amazon Web Services Support cases. After you create
// a case, use the AddCommunicationToCase operation to add additional communication
// or attachments to an existing case. The caseId is separate from the displayId
// that appears in the Amazon Web Services Support Center
// (https://console.aws.amazon.com/support). Use the DescribeCases operation to get
// the displayId.
//
// * You must have a Business, Enterprise On-Ramp, or Enterprise
// Support plan to use the Amazon Web Services Support API.
//
// * If you call the
// Amazon Web Services Support API from an account that does not have a Business,
// Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see Amazon Web Services Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) CreateCase(ctx context.Context, params *CreateCaseInput, optFns ...func(*Options)) (*CreateCaseOutput, error) {
	if params == nil {
		params = &CreateCaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCase", params, optFns, c.addOperationCreateCaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCaseInput struct {

	// The communication body text that describes the issue. This text appears in the
	// Description field on the Amazon Web Services Support Center Create Case
	// (https://console.aws.amazon.com/support/home#/case/create) page.
	//
	// This member is required.
	CommunicationBody *string

	// The title of the support case. The title appears in the Subject field on the
	// Amazon Web Services Support Center Create Case
	// (https://console.aws.amazon.com/support/home#/case/create) page.
	//
	// This member is required.
	Subject *string

	// The ID of a set of one or more attachments for the case. Create the set by using
	// the AddAttachmentsToSet operation.
	AttachmentSetId *string

	// The category of problem for the support case. You also use the DescribeServices
	// operation to get the category code for a service. Each Amazon Web Services
	// service defines its own set of category codes.
	CategoryCode *string

	// A list of email addresses that Amazon Web Services Support copies on case
	// correspondence. Amazon Web Services Support identifies the account that creates
	// the case when you specify your Amazon Web Services credentials in an HTTP POST
	// method or use the Amazon Web Services SDKs (http://aws.amazon.com/tools/).
	CcEmailAddresses []string

	// The type of issue for the case. You can specify customer-service or technical.
	// If you don't specify a value, the default is technical.
	IssueType *string

	// The language in which Amazon Web Services Support handles the case. You must
	// specify the ISO 639-1 code for the language parameter if you want support in
	// that language. Currently, English ("en") and Japanese ("ja") are supported.
	Language *string

	// The code for the Amazon Web Services service. You can use the DescribeServices
	// operation to get the possible serviceCode values.
	ServiceCode *string

	// A value that indicates the urgency of the case. This value determines the
	// response time according to your service level agreement with Amazon Web Services
	// Support. You can use the DescribeSeverityLevels operation to get the possible
	// values for severityCode. For more information, see SeverityLevel and Choosing a
	// Severity
	// (https://docs.aws.amazon.com/awssupport/latest/user/getting-started.html#choosing-severity)
	// in the Amazon Web Services Support User Guide. The availability of severity
	// levels depends on the support plan for the Amazon Web Services account.
	SeverityCode *string

	noSmithyDocumentSerde
}

// The support case ID returned by a successful completion of the CreateCase
// operation.
type CreateCaseOutput struct {

	// The support case ID requested or returned in the call. The case ID is an
	// alphanumeric string in the following format:
	// case-12345678910-2013-c4c1d2bf33c5cf47
	CaseId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCase{}, middleware.After)
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
	if err = addOpCreateCaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "CreateCase",
	}
}
