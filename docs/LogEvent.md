# LogEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | 
**LogLevel** | [**LogLevelEnum**](LogLevelEnum.md) |  | 
**Logger** | **string** |  | 
**Event** | **string** |  | 
**Attributes** | **map[string]interface{}** |  | 

## Methods

### NewLogEvent

`func NewLogEvent(timestamp time.Time, logLevel LogLevelEnum, logger string, event string, attributes map[string]interface{}, ) *LogEvent`

NewLogEvent instantiates a new LogEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogEventWithDefaults

`func NewLogEventWithDefaults() *LogEvent`

NewLogEventWithDefaults instantiates a new LogEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *LogEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LogEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetLogLevel

`func (o *LogEvent) GetLogLevel() LogLevelEnum`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *LogEvent) GetLogLevelOk() (*LogLevelEnum, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *LogEvent) SetLogLevel(v LogLevelEnum)`

SetLogLevel sets LogLevel field to given value.


### GetLogger

`func (o *LogEvent) GetLogger() string`

GetLogger returns the Logger field if non-nil, zero value otherwise.

### GetLoggerOk

`func (o *LogEvent) GetLoggerOk() (*string, bool)`

GetLoggerOk returns a tuple with the Logger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogger

`func (o *LogEvent) SetLogger(v string)`

SetLogger sets Logger field to given value.


### GetEvent

`func (o *LogEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *LogEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *LogEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetAttributes

`func (o *LogEvent) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LogEvent) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LogEvent) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


