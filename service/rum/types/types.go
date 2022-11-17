// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A RUM app monitor collects telemetry data from your application and sends that
// data to RUM. The data includes performance and reliability information such as
// page load time, client-side errors, and user behavior.
type AppMonitor struct {

	// A structure that contains much of the configuration data for the app monitor.
	AppMonitorConfiguration *AppMonitorConfiguration

	// The date and time that this app monitor was created.
	Created *string

	// Specifies whether this app monitor allows the web client to define and send
	// custom events. For more information about custom events, see Send custom events
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-custom-events.html).
	CustomEvents *CustomEvents

	// A structure that contains information about whether this app monitor stores a
	// copy of the telemetry data that RUM collects using CloudWatch Logs.
	DataStorage *DataStorage

	// The top-level internet domain name for which your application has administrative
	// authority.
	Domain *string

	// The unique ID of this app monitor.
	Id *string

	// The date and time of the most recent changes to this app monitor's
	// configuration.
	LastModified *string

	// The name of the app monitor.
	Name *string

	// The current state of the app monitor.
	State StateEnum

	// The list of tag keys and values associated with this app monitor.
	Tags map[string]string

	noSmithyDocumentSerde
}

// This structure contains much of the configuration data for the app monitor.
type AppMonitorConfiguration struct {

	// If you set this to true, the RUM web client sets two cookies, a session cookie
	// and a user cookie. The cookies allow the RUM web client to collect data relating
	// to the number of users an application has and the behavior of the application
	// across a sequence of events. Cookies are stored in the top-level domain of the
	// current page.
	AllowCookies *bool

	// If you set this to true, RUM enables X-Ray tracing for the user sessions that
	// RUM samples. RUM adds an X-Ray trace header to allowed HTTP requests. It also
	// records an X-Ray segment for allowed HTTP requests. You can see traces and
	// segments from these user sessions in the X-Ray console and the CloudWatch
	// ServiceLens console. For more information, see What is X-Ray?
	// (https://docs.aws.amazon.com/xray/latest/devguide/aws-xray.html)
	EnableXRay *bool

	// A list of URLs in your website or application to exclude from RUM data
	// collection. You can't include both ExcludedPages and IncludedPages in the same
	// operation.
	ExcludedPages []string

	// A list of pages in your application that are to be displayed with a "favorite"
	// icon in the CloudWatch RUM console.
	FavoritePages []string

	// The ARN of the guest IAM role that is attached to the Amazon Cognito identity
	// pool that is used to authorize the sending of data to RUM.
	GuestRoleArn *string

	// The ID of the Amazon Cognito identity pool that is used to authorize the sending
	// of data to RUM.
	IdentityPoolId *string

	// If this app monitor is to collect data from only certain pages in your
	// application, this structure lists those pages. You can't include both
	// ExcludedPages and IncludedPages in the same operation.
	IncludedPages []string

	// Specifies the portion of user sessions to use for RUM data collection. Choosing
	// a higher portion gives you more data but also incurs more costs. The range for
	// this value is 0 to 1 inclusive. Setting this to 1 means that 100% of user
	// sessions are sampled, and setting it to 0.1 means that 10% of user sessions are
	// sampled. If you omit this parameter, the default of 0.1 is used, and 10% of
	// sessions will be sampled.
	SessionSampleRate float64

	// An array that lists the types of telemetry data that this app monitor is to
	// collect.
	//
	// * errors indicates that RUM collects data about unhandled JavaScript
	// errors raised by your application.
	//
	// * performance indicates that RUM collects
	// performance data about how your application and its resources are loaded and
	// rendered. This includes Core Web Vitals.
	//
	// * http indicates that RUM collects
	// data about HTTP errors thrown by your application.
	Telemetries []Telemetry

	noSmithyDocumentSerde
}

// A structure that contains information about the RUM app monitor.
type AppMonitorDetails struct {

	// The unique ID of the app monitor.
	Id *string

	// The name of the app monitor.
	Name *string

	// The version of the app monitor.
	Version *string

	noSmithyDocumentSerde
}

// A structure that includes some data about app monitors and their settings.
type AppMonitorSummary struct {

	// The date and time that the app monitor was created.
	Created *string

	// The unique ID of this app monitor.
	Id *string

	// The date and time of the most recent changes to this app monitor's
	// configuration.
	LastModified *string

	// The name of this app monitor.
	Name *string

	// The current state of this app monitor.
	State StateEnum

	noSmithyDocumentSerde
}

// A structure that defines one error caused by a BatchCreateRumMetricsDefinitions
// (https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_BatchCreateRumMetricsDefinitions.html)
// operation.
type BatchCreateRumMetricDefinitionsError struct {

	// The error code.
	//
	// This member is required.
	ErrorCode *string

	// The error message for this metric definition.
	//
	// This member is required.
	ErrorMessage *string

	// The metric definition that caused this error.
	//
	// This member is required.
	MetricDefinition *MetricDefinitionRequest

	noSmithyDocumentSerde
}

// A structure that defines one error caused by a BatchCreateRumMetricsDefinitions
// (https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_BatchDeleteRumMetricsDefinitions.html)
// operation.
type BatchDeleteRumMetricDefinitionsError struct {

	// The error code.
	//
	// This member is required.
	ErrorCode *string

	// The error message for this metric definition.
	//
	// This member is required.
	ErrorMessage *string

	// The ID of the metric definition that caused this error.
	//
	// This member is required.
	MetricDefinitionId *string

	noSmithyDocumentSerde
}

// A structure that contains information about custom events for this app monitor.
type CustomEvents struct {

	// Specifies whether this app monitor allows the web client to define and send
	// custom events. The default is for custom events to be DISABLED.
	Status CustomEventsStatus

	noSmithyDocumentSerde
}

// A structure that contains the information about whether the app monitor stores
// copies of the data that RUM collects in CloudWatch Logs. If it does, this
// structure also contains the name of the log group.
type CwLog struct {

	// Indicated whether the app monitor stores copies of the data that RUM collects in
	// CloudWatch Logs.
	CwLogEnabled *bool

	// The name of the log group where the copies are stored.
	CwLogGroup *string

	noSmithyDocumentSerde
}

// A structure that contains information about whether this app monitor stores a
// copy of the telemetry data that RUM collects using CloudWatch Logs.
type DataStorage struct {

	// A structure that contains the information about whether the app monitor stores
	// copies of the data that RUM collects in CloudWatch Logs. If it does, this
	// structure also contains the name of the log group.
	CwLog *CwLog

	noSmithyDocumentSerde
}

// A structure that displays the definition of one extended metric that RUM sends
// to CloudWatch or CloudWatch Evidently. For more information, see  Additional
// metrics that you can send to CloudWatch and CloudWatch Evidently
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-vended-metrics.html).
type MetricDefinition struct {

	// The ID of this metric definition.
	//
	// This member is required.
	MetricDefinitionId *string

	// The name of the metric that is defined in this structure.
	//
	// This member is required.
	Name *string

	// This field is a map of field paths to dimension names. It defines the dimensions
	// to associate with this metric in CloudWatch The value of this field is used only
	// if the metric destination is CloudWatch. If the metric destination is Evidently,
	// the value of DimensionKeys is ignored.
	DimensionKeys map[string]string

	// The pattern that defines the metric. RUM checks events that happen in a user's
	// session against the pattern, and events that match the pattern are sent to the
	// metric destination. If the metrics destination is CloudWatch and the event also
	// matches a value in DimensionKeys, then the metric is published with the
	// specified dimensions.
	EventPattern *string

	// Use this field only if you are sending this metric to CloudWatch. It defines the
	// CloudWatch metric unit that this metric is measured in.
	UnitLabel *string

	// The field within the event object that the metric value is sourced from.
	ValueKey *string

	noSmithyDocumentSerde
}

// Use this structure to define one extended metric that RUM will send to
// CloudWatch or CloudWatch Evidently. For more information, see  Additional
// metrics that you can send to CloudWatch and CloudWatch Evidently
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-vended-metrics.html).
// Only certain combinations of values for Name, ValueKey, and EventPattern are
// valid. In addition to what is displayed in the list below, the EventPattern can
// also include information used by the DimensionKeys field.
//
// * If Name is
// PerformanceNavigationDuration, then ValueKeymust be event_details.duration and
// the EventPattern must include
// {"event_type":["com.amazon.rum.performance_navigation_event"]}
//
// * If Name is
// PerformanceResourceDuration, then ValueKeymust be event_details.duration and the
// EventPattern must include
// {"event_type":["com.amazon.rum.performance_resource_event"]}
//
// * If Name is
// NavigationSatisfiedTransaction, then ValueKeymust be null and the EventPattern
// must include { "event_type": ["com.amazon.rum.performance_navigation_event"],
// "event_details": { "duration": [{ "numeric": [">",2000] }] } }
//
// * If Name is
// NavigationToleratedTransaction, then ValueKeymust be null and the EventPattern
// must include { "event_type": ["com.amazon.rum.performance_navigation_event"],
// "event_details": { "duration": [{ "numeric": [">=",2000,"<"8000] }] } }
//
// * If
// Name is NavigationFrustratedTransaction, then ValueKeymust be null and the
// EventPattern must include { "event_type":
// ["com.amazon.rum.performance_navigation_event"], "event_details": { "duration":
// [{ "numeric": [">=",8000] }] } }
//
// * If Name is WebVitalsCumulativeLayoutShift,
// then ValueKeymust be event_details.value and the EventPattern must include
// {"event_type":["com.amazon.rum.cumulative_layout_shift_event"]}
//
// * If Name is
// WebVitalsFirstInputDelay, then ValueKeymust be event_details.value and the
// EventPattern must include
// {"event_type":["com.amazon.rum.first_input_delay_event"]}
//
// * If Name is
// WebVitalsLargestContentfulPaint, then ValueKeymust be event_details.value and
// the EventPattern must include
// {"event_type":["com.amazon.rum.largest_contentful_paint_event"]}
//
// * If Name is
// JsErrorCount, then ValueKeymust be null and the EventPattern must include
// {"event_type":["com.amazon.rum.js_error_event"]}
//
// * If Name is HttpErrorCount,
// then ValueKeymust be null and the EventPattern must include
// {"event_type":["com.amazon.rum.http_event"]}
//
// * If Name is SessionCount, then
// ValueKeymust be null and the EventPattern must include
// {"event_type":["com.amazon.rum.session_start_event"]}
type MetricDefinitionRequest struct {

	// The name for the metric that is defined in this structure. Valid values are the
	// following:
	//
	// * PerformanceNavigationDuration
	//
	// * PerformanceResourceDuration
	//
	// *
	// NavigationSatisfiedTransaction
	//
	// * NavigationToleratedTransaction
	//
	// *
	// NavigationFrustratedTransaction
	//
	// * WebVitalsCumulativeLayoutShift
	//
	// *
	// WebVitalsFirstInputDelay
	//
	// * WebVitalsLargestContentfulPaint
	//
	// * JsErrorCount
	//
	// *
	// HttpErrorCount
	//
	// * SessionCount
	//
	// This member is required.
	Name *string

	// Use this field only if you are sending the metric to CloudWatch. This field is a
	// map of field paths to dimension names. It defines the dimensions to associate
	// with this metric in CloudWatch. Valid values for the entries in this field are
	// the following:
	//
	// * "metadata.pageId": "PageId"
	//
	// * "metadata.browserName":
	// "BrowserName"
	//
	// * "metadata.deviceType": "DeviceType"
	//
	// * "metadata.osName":
	// "OSName"
	//
	// * "metadata.countryCode": "CountryCode"
	//
	// * "event_details.fileType":
	// "FileType"
	//
	// All dimensions listed in this field must also be included in
	// EventPattern.
	DimensionKeys map[string]string

	// The pattern that defines the metric, specified as a JSON object. RUM checks
	// events that happen in a user's session against the pattern, and events that
	// match the pattern are sent to the metric destination. When you define extended
	// metrics, the metric definition is not valid if EventPattern is omitted. Example
	// event patterns:
	//
	// * '{ "event_type": ["com.amazon.rum.js_error_event"],
	// "metadata": { "browserName": [ "Chrome", "Safari" ], } }'
	//
	// * '{ "event_type":
	// ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [
	// "Chrome", "Firefox" ] }, "event_details": { "duration": [{ "numeric": [ "<",
	// 2000 ] }] } }'
	//
	// * '{ "event_type":
	// ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [
	// "Chrome", "Safari" ], "countryCode": [ "US" ] }, "event_details": { "duration":
	// [{ "numeric": [ ">=", 2000, "<", 8000 ] }] } }'
	//
	// If the metrics destination' is
	// CloudWatch and the event also matches a value in DimensionKeys, then the metric
	// is published with the specified dimensions.
	EventPattern *string

	// The CloudWatch metric unit to use for this metric. If you omit this field, the
	// metric is recorded with no unit.
	UnitLabel *string

	// The field within the event object that the metric value is sourced from. If you
	// omit this field, a hardcoded value of 1 is pushed as the metric value. This is
	// useful if you just want to count the number of events that the filter catches.
	// If this metric is sent to CloudWatch Evidently, this field will be passed to
	// Evidently raw and Evidently will handle data extraction from the event.
	ValueKey *string

	noSmithyDocumentSerde
}

// A structure that displays information about one destination that CloudWatch RUM
// sends extended metrics to.
type MetricDestinationSummary struct {

	// Specifies whether the destination is CloudWatch or Evidently.
	Destination MetricDestination

	// If the destination is Evidently, this specifies the ARN of the Evidently
	// experiment that receives the metrics.
	DestinationArn *string

	// This field appears only when the destination is Evidently. It specifies the ARN
	// of the IAM role that is used to write to the Evidently experiment that receives
	// the metrics.
	IamRoleArn *string

	noSmithyDocumentSerde
}

// A structure that defines a key and values that you can use to filter the
// results. The only performance events that are returned are those that have
// values matching the ones that you specify in one of your QueryFilter structures.
// For example, you could specify Browser as the Name and specify Chrome,Firefox as
// the Values to return events generated only from those browsers. Specifying
// Invert as the Name works as a "not equal to" filter. For example, specify Invert
// as the Name and specify Chrome as the value to return all events except events
// from user sessions with the Chrome browser.
type QueryFilter struct {

	// The name of a key to search for. The filter returns only the events that match
	// the Name and Values that you specify. Valid values for Name are Browser | Device
	// | Country | Page | OS | EventType | Invert
	Name *string

	// The values of the Name that are to be be included in the returned results.
	Values []string

	noSmithyDocumentSerde
}

// A structure that contains the information for one performance event that RUM
// collects from a user session with your application.
type RumEvent struct {

	// A string containing details about the event.
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	Details *string

	// A unique ID for this event.
	//
	// This member is required.
	Id *string

	// The exact time that this event occurred.
	//
	// This member is required.
	Timestamp *time.Time

	// The JSON schema that denotes the type of event this is, such as a page load or a
	// new session.
	//
	// This member is required.
	Type *string

	// Metadata about this event, which contains a JSON serialization of the identity
	// of the user for this session. The user information comes from information such
	// as the HTTP user-agent request header and document interface.
	//
	// This value conforms to the media type: application/json
	Metadata *string

	noSmithyDocumentSerde
}

// A structure that defines the time range that you want to retrieve results from.
type TimeRange struct {

	// The beginning of the time range to retrieve performance events from.
	//
	// This member is required.
	After int64

	// The end of the time range to retrieve performance events from. If you omit this,
	// the time range extends to the time that this operation is performed.
	Before int64

	noSmithyDocumentSerde
}

// A structure that contains information about the user session that this batch of
// events was collected from.
type UserDetails struct {

	// The session ID that the performance events are from.
	SessionId *string

	// The ID of the user for this user session. This ID is generated by RUM and does
	// not include any personally identifiable information about the user.
	UserId *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
