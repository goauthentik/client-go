# PasswordExpiryPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Days** | **int32** |  | 
**DenyOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewPasswordExpiryPolicyRequest

`func NewPasswordExpiryPolicyRequest(name string, days int32, ) *PasswordExpiryPolicyRequest`

NewPasswordExpiryPolicyRequest instantiates a new PasswordExpiryPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordExpiryPolicyRequestWithDefaults

`func NewPasswordExpiryPolicyRequestWithDefaults() *PasswordExpiryPolicyRequest`

NewPasswordExpiryPolicyRequestWithDefaults instantiates a new PasswordExpiryPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PasswordExpiryPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordExpiryPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordExpiryPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *PasswordExpiryPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *PasswordExpiryPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *PasswordExpiryPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *PasswordExpiryPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetDays

`func (o *PasswordExpiryPolicyRequest) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *PasswordExpiryPolicyRequest) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *PasswordExpiryPolicyRequest) SetDays(v int32)`

SetDays sets Days field to given value.


### GetDenyOnly

`func (o *PasswordExpiryPolicyRequest) GetDenyOnly() bool`

GetDenyOnly returns the DenyOnly field if non-nil, zero value otherwise.

### GetDenyOnlyOk

`func (o *PasswordExpiryPolicyRequest) GetDenyOnlyOk() (*bool, bool)`

GetDenyOnlyOk returns a tuple with the DenyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyOnly

`func (o *PasswordExpiryPolicyRequest) SetDenyOnly(v bool)`

SetDenyOnly sets DenyOnly field to given value.

### HasDenyOnly

`func (o *PasswordExpiryPolicyRequest) HasDenyOnly() bool`

HasDenyOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


