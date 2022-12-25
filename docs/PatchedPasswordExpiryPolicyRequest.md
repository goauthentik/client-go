# PatchedPasswordExpiryPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Days** | Pointer to **int32** |  | [optional] 
**DenyOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedPasswordExpiryPolicyRequest

`func NewPatchedPasswordExpiryPolicyRequest() *PatchedPasswordExpiryPolicyRequest`

NewPatchedPasswordExpiryPolicyRequest instantiates a new PatchedPasswordExpiryPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPasswordExpiryPolicyRequestWithDefaults

`func NewPatchedPasswordExpiryPolicyRequestWithDefaults() *PatchedPasswordExpiryPolicyRequest`

NewPatchedPasswordExpiryPolicyRequestWithDefaults instantiates a new PatchedPasswordExpiryPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedPasswordExpiryPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedPasswordExpiryPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedPasswordExpiryPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedPasswordExpiryPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExecutionLogging

`func (o *PatchedPasswordExpiryPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *PatchedPasswordExpiryPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *PatchedPasswordExpiryPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *PatchedPasswordExpiryPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetDays

`func (o *PatchedPasswordExpiryPolicyRequest) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *PatchedPasswordExpiryPolicyRequest) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *PatchedPasswordExpiryPolicyRequest) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *PatchedPasswordExpiryPolicyRequest) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetDenyOnly

`func (o *PatchedPasswordExpiryPolicyRequest) GetDenyOnly() bool`

GetDenyOnly returns the DenyOnly field if non-nil, zero value otherwise.

### GetDenyOnlyOk

`func (o *PatchedPasswordExpiryPolicyRequest) GetDenyOnlyOk() (*bool, bool)`

GetDenyOnlyOk returns a tuple with the DenyOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyOnly

`func (o *PatchedPasswordExpiryPolicyRequest) SetDenyOnly(v bool)`

SetDenyOnly sets DenyOnly field to given value.

### HasDenyOnly

`func (o *PatchedPasswordExpiryPolicyRequest) HasDenyOnly() bool`

HasDenyOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


