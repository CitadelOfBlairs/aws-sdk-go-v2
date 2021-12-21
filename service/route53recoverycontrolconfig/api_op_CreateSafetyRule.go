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

// Creates a safety rule in a control panel. Safety rules let you add safeguards
// around changing routing control states, and for enabling and disabling routing
// controls, to help prevent unexpected outcomes. There are two types of safety
// rules: assertion rules and gating rules. Assertion rule: An assertion rule
// enforces that, when you change a routing control state, that a certain criteria
// is met. For example, the criteria might be that at least one routing control
// state is On after the transation so that traffic continues to flow to at least
// one cell for the application. This ensures that you avoid a fail-open scenario.
// Gating rule: A gating rule lets you configure a gating routing control as an
// overall "on/off" switch for a group of routing controls. Or, you can configure
// more complex gating scenarios, for example by configuring multiple gating
// routing controls. For more information, see Safety rules
// (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.safety-rules.html)
// in the Amazon Route 53 Application Recovery Controller Developer Guide.
func (c *Client) CreateSafetyRule(ctx context.Context, params *CreateSafetyRuleInput, optFns ...func(*Options)) (*CreateSafetyRuleOutput, error) {
	if params == nil {
		params = &CreateSafetyRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSafetyRule", params, optFns, c.addOperationCreateSafetyRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSafetyRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body that you include when you create a safety rule.
type CreateSafetyRuleInput struct {

	// The assertion rule requested.
	AssertionRule *types.NewAssertionRule

	// A unique, case-sensitive string of up to 64 ASCII characters. To make an
	// idempotent API request with an action, specify a client token in the request.
	ClientToken *string

	// The gating rule requested.
	GatingRule *types.NewGatingRule

	// The tags associated with the safety rule.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSafetyRuleOutput struct {

	// The assertion rule created.
	AssertionRule *types.AssertionRule

	// The gating rule created.
	GatingRule *types.GatingRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSafetyRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSafetyRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSafetyRule{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateSafetyRuleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSafetyRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSafetyRule(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateSafetyRule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateSafetyRule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateSafetyRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateSafetyRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateSafetyRuleInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateSafetyRuleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateSafetyRule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateSafetyRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-control-config",
		OperationName: "CreateSafetyRule",
	}
}
