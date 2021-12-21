// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous job to detect medication entities and link them to the
// RxNorm ontology. Use the DescribeRxNormInferenceJob operation to track the
// status of a job.
func (c *Client) StartRxNormInferenceJob(ctx context.Context, params *StartRxNormInferenceJobInput, optFns ...func(*Options)) (*StartRxNormInferenceJobOutput, error) {
	if params == nil {
		params = &StartRxNormInferenceJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartRxNormInferenceJob", params, optFns, c.addOperationStartRxNormInferenceJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartRxNormInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartRxNormInferenceJobInput struct {

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that grants Comprehend Medical; read access to your input data. For more
	// information, see  Role-Based Permissions Required for Asynchronous Operations
	// (https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions-med.html#auth-role-permissions-med).
	//
	// This member is required.
	DataAccessRoleArn *string

	// Specifies the format and location of the input data for the job.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// The language of the input documents. All documents must be in the same language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// Specifies where to send the output files.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// A unique identifier for the request. If you don't set the client request token,
	// Comprehend Medical; generates one.
	ClientRequestToken *string

	// The identifier of the job.
	JobName *string

	// An AWS Key Management Service key to encrypt your output files. If you do not
	// specify a key, the files are written in plain text.
	KMSKey *string

	noSmithyDocumentSerde
}

type StartRxNormInferenceJobOutput struct {

	// The identifier of the job.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartRxNormInferenceJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartRxNormInferenceJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartRxNormInferenceJob{}, middleware.After)
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
	if err = addIdempotencyToken_opStartRxNormInferenceJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartRxNormInferenceJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartRxNormInferenceJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartRxNormInferenceJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartRxNormInferenceJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartRxNormInferenceJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartRxNormInferenceJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartRxNormInferenceJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartRxNormInferenceJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartRxNormInferenceJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartRxNormInferenceJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "StartRxNormInferenceJob",
	}
}
