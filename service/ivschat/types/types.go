// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Specifies a CloudWatch Logs location where chat logs will be stored.
type CloudWatchLogsDestinationConfiguration struct {

	// Name of the Amazon Cloudwatch Logs destination where chat activity will be
	// logged.
	//
	// This member is required.
	LogGroupName *string

	noSmithyDocumentSerde
}

// A complex type that describes a location where chat logs will be stored. Each
// member represents the configuration of one log destination. For logging, you
// define only one type of destination (for CloudWatch Logs, Kinesis Firehose, or
// S3).
//
// The following types satisfy this interface:
//
//	DestinationConfigurationMemberCloudWatchLogs
//	DestinationConfigurationMemberFirehose
//	DestinationConfigurationMemberS3
type DestinationConfiguration interface {
	isDestinationConfiguration()
}

// Name of the Amazon CloudWatch Logs destination where chat activity will be
// logged.
type DestinationConfigurationMemberCloudWatchLogs struct {
	Value CloudWatchLogsDestinationConfiguration

	noSmithyDocumentSerde
}

func (*DestinationConfigurationMemberCloudWatchLogs) isDestinationConfiguration() {}

// Name of the Amazon Kinesis Data Firehose destination where chat activity will be
// logged
type DestinationConfigurationMemberFirehose struct {
	Value FirehoseDestinationConfiguration

	noSmithyDocumentSerde
}

func (*DestinationConfigurationMemberFirehose) isDestinationConfiguration() {}

// Name of the Amazon S3 bucket where chat activity will be logged.
type DestinationConfigurationMemberS3 struct {
	Value S3DestinationConfiguration

	noSmithyDocumentSerde
}

func (*DestinationConfigurationMemberS3) isDestinationConfiguration() {}

// Specifies a Kinesis Firehose location where chat logs will be stored.
type FirehoseDestinationConfiguration struct {

	// Name of the Amazon Kinesis Firehose delivery stream where chat activity will be
	// logged.
	//
	// This member is required.
	DeliveryStreamName *string

	noSmithyDocumentSerde
}

// Summary information about a logging configuration.
type LoggingConfigurationSummary struct {

	// Logging-configuration ARN.
	Arn *string

	// Time when the logging configuration was created. This is an ISO 8601 timestamp;
	// note that this is returned as a string.
	CreateTime *time.Time

	// A complex type that contains a destination configuration for where chat content
	// will be logged.
	DestinationConfiguration DestinationConfiguration

	// Logging-configuration ID, generated by the system. This is a relative
	// identifier, the part of the ARN that uniquely identifies the room.
	Id *string

	// Logging-configuration name. The value does not need to be unique.
	Name *string

	// The state of the logging configuration. When this is ACTIVE, the configuration
	// is ready for logging chat content.
	State LoggingConfigurationState

	// Tags to attach to the resource. Array of maps, each of the form string:string
	// (key:value). See Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS Chat has no constraints on tags beyond what is
	// documented there.
	Tags map[string]string

	// Time of the logging configuration’s last update. This is an ISO 8601 timestamp;
	// note that this is returned as a string.
	UpdateTime *time.Time

	noSmithyDocumentSerde
}

// Configuration information for optional message review.
type MessageReviewHandler struct {

	// Specifies the fallback behavior (whether the message is allowed or denied) if
	// the handler does not return a valid response, encounters an error, or times out.
	// (For the timeout period, see  Service Quotas
	// (https://docs.aws.amazon.com/ivs/latest/userguide/service-quotas.html).) If
	// allowed, the message is delivered with returned content to all users connected
	// to the room. If denied, the message is not delivered to any user. Default:
	// ALLOW.
	FallbackResult FallbackResult

	// Identifier of the message review handler. Currently this must be an ARN of a
	// lambda function.
	Uri *string

	noSmithyDocumentSerde
}

// Summary information about a room.
type RoomSummary struct {

	// Room ARN.
	Arn *string

	// Time when the room was created. This is an ISO 8601 timestamp; note that this is
	// returned as a string.
	CreateTime *time.Time

	// Room ID, generated by the system. This is a relative identifier, the part of the
	// ARN that uniquely identifies the room.
	Id *string

	// List of logging-configuration identifiers attached to the room.
	LoggingConfigurationIdentifiers []string

	// Configuration information for optional review of messages.
	MessageReviewHandler *MessageReviewHandler

	// Room name. The value does not need to be unique.
	Name *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value). See Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) for details,
	// including restrictions that apply to tags and "Tag naming limits and
	// requirements"; Amazon IVS Chat has no constraints beyond what is documented
	// there.
	Tags map[string]string

	// Time of the room’s last update. This is an ISO 8601 timestamp; note that this is
	// returned as a string.
	UpdateTime *time.Time

	noSmithyDocumentSerde
}

// Specifies an S3 location where chat logs will be stored.
type S3DestinationConfiguration struct {

	// Name of the Amazon S3 bucket where chat activity will be logged.
	//
	// This member is required.
	BucketName *string

	noSmithyDocumentSerde
}

// This object is used in the ValidationException error.
type ValidationExceptionField struct {

	// Explanation of the reason for the validation error.
	//
	// This member is required.
	Message *string

	// Name of the field which failed validation.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isDestinationConfiguration() {}
