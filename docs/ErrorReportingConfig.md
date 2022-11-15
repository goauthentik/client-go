# ErrorReportingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | [readonly] 
**SentryDsn** | **string** |  | [readonly] 
**Environment** | **string** |  | [readonly] 
**SendPii** | **bool** |  | [readonly] 
**TracesSampleRate** | **float64** |  | [readonly] 

## Methods

### NewErrorReportingConfig

`func NewErrorReportingConfig(enabled bool, sentryDsn string, environment string, sendPii bool, tracesSampleRate float64, ) *ErrorReportingConfig`

NewErrorReportingConfig instantiates a new ErrorReportingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorReportingConfigWithDefaults

`func NewErrorReportingConfigWithDefaults() *ErrorReportingConfig`

NewErrorReportingConfigWithDefaults instantiates a new ErrorReportingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ErrorReportingConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ErrorReportingConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ErrorReportingConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSentryDsn

`func (o *ErrorReportingConfig) GetSentryDsn() string`

GetSentryDsn returns the SentryDsn field if non-nil, zero value otherwise.

### GetSentryDsnOk

`func (o *ErrorReportingConfig) GetSentryDsnOk() (*string, bool)`

GetSentryDsnOk returns a tuple with the SentryDsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentryDsn

`func (o *ErrorReportingConfig) SetSentryDsn(v string)`

SetSentryDsn sets SentryDsn field to given value.


### GetEnvironment

`func (o *ErrorReportingConfig) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ErrorReportingConfig) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ErrorReportingConfig) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetSendPii

`func (o *ErrorReportingConfig) GetSendPii() bool`

GetSendPii returns the SendPii field if non-nil, zero value otherwise.

### GetSendPiiOk

`func (o *ErrorReportingConfig) GetSendPiiOk() (*bool, bool)`

GetSendPiiOk returns a tuple with the SendPii field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPii

`func (o *ErrorReportingConfig) SetSendPii(v bool)`

SetSendPii sets SendPii field to given value.


### GetTracesSampleRate

`func (o *ErrorReportingConfig) GetTracesSampleRate() float64`

GetTracesSampleRate returns the TracesSampleRate field if non-nil, zero value otherwise.

### GetTracesSampleRateOk

`func (o *ErrorReportingConfig) GetTracesSampleRateOk() (*float64, bool)`

GetTracesSampleRateOk returns a tuple with the TracesSampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracesSampleRate

`func (o *ErrorReportingConfig) SetTracesSampleRate(v float64)`

SetTracesSampleRate sets TracesSampleRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


