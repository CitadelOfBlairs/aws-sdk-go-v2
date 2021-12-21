// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Get streaming image.
func (c *Client) GetStreamingImage(ctx context.Context, params *GetStreamingImageInput, optFns ...func(*Options)) (*GetStreamingImageOutput, error) {
	if params == nil {
		params = &GetStreamingImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStreamingImage", params, optFns, c.addOperationGetStreamingImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStreamingImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetStreamingImageInput struct {

	// The streaming image ID.
	//
	// This member is required.
	StreamingImageId *string

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	noSmithyDocumentSerde
}

type GetStreamingImageOutput struct {

	// The streaming image.
	StreamingImage *types.StreamingImage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetStreamingImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetStreamingImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetStreamingImage{}, middleware.After)
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
	if err = addOpGetStreamingImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetStreamingImage(options.Region), middleware.Before); err != nil {
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

// GetStreamingImageAPIClient is a client that implements the GetStreamingImage
// operation.
type GetStreamingImageAPIClient interface {
	GetStreamingImage(context.Context, *GetStreamingImageInput, ...func(*Options)) (*GetStreamingImageOutput, error)
}

var _ GetStreamingImageAPIClient = (*Client)(nil)

// StreamingImageReadyWaiterOptions are waiter options for
// StreamingImageReadyWaiter
type StreamingImageReadyWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// StreamingImageReadyWaiter will use default minimum delay of 2 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, StreamingImageReadyWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *GetStreamingImageInput, *GetStreamingImageOutput, error) (bool, error)
}

// StreamingImageReadyWaiter defines the waiters for StreamingImageReady
type StreamingImageReadyWaiter struct {
	client GetStreamingImageAPIClient

	options StreamingImageReadyWaiterOptions
}

// NewStreamingImageReadyWaiter constructs a StreamingImageReadyWaiter.
func NewStreamingImageReadyWaiter(client GetStreamingImageAPIClient, optFns ...func(*StreamingImageReadyWaiterOptions)) *StreamingImageReadyWaiter {
	options := StreamingImageReadyWaiterOptions{}
	options.MinDelay = 2 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = streamingImageReadyStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &StreamingImageReadyWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for StreamingImageReady waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *StreamingImageReadyWaiter) Wait(ctx context.Context, params *GetStreamingImageInput, maxWaitDur time.Duration, optFns ...func(*StreamingImageReadyWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for StreamingImageReady waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *StreamingImageReadyWaiter) WaitForOutput(ctx context.Context, params *GetStreamingImageInput, maxWaitDur time.Duration, optFns ...func(*StreamingImageReadyWaiterOptions)) (*GetStreamingImageOutput, error) {
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

		out, err := w.client.GetStreamingImage(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for StreamingImageReady waiter")
}

func streamingImageReadyStateRetryable(ctx context.Context, input *GetStreamingImageInput, output *GetStreamingImageOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("streamingImage.state", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "READY"
		value, ok := pathValue.(types.StreamingImageState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StreamingImageState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("streamingImage.state", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_FAILED"
		value, ok := pathValue.(types.StreamingImageState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StreamingImageState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("streamingImage.state", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "UPDATE_FAILED"
		value, ok := pathValue.(types.StreamingImageState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StreamingImageState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// StreamingImageDeletedWaiterOptions are waiter options for
// StreamingImageDeletedWaiter
type StreamingImageDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// StreamingImageDeletedWaiter will use default minimum delay of 2 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, StreamingImageDeletedWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *GetStreamingImageInput, *GetStreamingImageOutput, error) (bool, error)
}

// StreamingImageDeletedWaiter defines the waiters for StreamingImageDeleted
type StreamingImageDeletedWaiter struct {
	client GetStreamingImageAPIClient

	options StreamingImageDeletedWaiterOptions
}

// NewStreamingImageDeletedWaiter constructs a StreamingImageDeletedWaiter.
func NewStreamingImageDeletedWaiter(client GetStreamingImageAPIClient, optFns ...func(*StreamingImageDeletedWaiterOptions)) *StreamingImageDeletedWaiter {
	options := StreamingImageDeletedWaiterOptions{}
	options.MinDelay = 2 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = streamingImageDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &StreamingImageDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for StreamingImageDeleted waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *StreamingImageDeletedWaiter) Wait(ctx context.Context, params *GetStreamingImageInput, maxWaitDur time.Duration, optFns ...func(*StreamingImageDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for StreamingImageDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *StreamingImageDeletedWaiter) WaitForOutput(ctx context.Context, params *GetStreamingImageInput, maxWaitDur time.Duration, optFns ...func(*StreamingImageDeletedWaiterOptions)) (*GetStreamingImageOutput, error) {
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

		out, err := w.client.GetStreamingImage(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for StreamingImageDeleted waiter")
}

func streamingImageDeletedStateRetryable(ctx context.Context, input *GetStreamingImageInput, output *GetStreamingImageOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("streamingImage.state", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETED"
		value, ok := pathValue.(types.StreamingImageState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StreamingImageState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("streamingImage.state", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETE_FAILED"
		value, ok := pathValue.(types.StreamingImageState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.StreamingImageState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetStreamingImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "GetStreamingImage",
	}
}
