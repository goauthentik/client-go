# PatchedHaveIBeenPwendPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**PasswordField** | Pointer to **string** | Field key to check, field keys defined in Prompt stages are available. | [optional] 
**AllowedCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedHaveIBeenPwendPolicyRequest

`func NewPatchedHaveIBeenPwendPolicyRequest() *PatchedHaveIBeenPwendPolicyRequest`

NewPatchedHaveIBeenPwendPolicyRequest instantiates a new PatchedHaveIBeenPwendPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedHaveIBeenPwendPolicyRequestWithDefaults

`func NewPatchedHaveIBeenPwendPolicyRequestWithDefaults() *PatchedHaveIBeenPwendPolicyRequest`

NewPatchedHaveIBeenPwendPolicyRequestWithDefaults instantiates a new PatchedHaveIBeenPwendPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedHaveIBeenPwendPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedHaveIBeenPwendPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedHaveIBeenPwendPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedHaveIBeenPwendPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedHaveIBeenPwendPolicyRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedHaveIBeenPwendPolicyRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetExecutionLogging

`func (o *PatchedHaveIBeenPwendPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *PatchedHaveIBeenPwendPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *PatchedHaveIBeenPwendPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *PatchedHaveIBeenPwendPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetPasswordField

`func (o *PatchedHaveIBeenPwendPolicyRequest) GetPasswordField() string`

GetPasswordField returns the PasswordField field if non-nil, zero value otherwise.

### GetPasswordFieldOk

`func (o *PatchedHaveIBeenPwendPolicyRequest) GetPasswordFieldOk() (*string, bool)`

GetPasswordFieldOk returns a tuple with the PasswordField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordField

`func (o *PatchedHaveIBeenPwendPolicyRequest) SetPasswordField(v string)`

SetPasswordField sets PasswordField field to given value.

### HasPasswordField

`func (o *PatchedHaveIBeenPwendPolicyRequest) HasPasswordField() bool`

HasPasswordField returns a boolean if a field has been set.

### GetAllowedCount

`func (o *PatchedHaveIBeenPwendPolicyRequest) GetAllowedCount() int32`

GetAllowedCount returns the AllowedCount field if non-nil, zero value otherwise.

### GetAllowedCountOk

`func (o *PatchedHaveIBeenPwendPolicyRequest) GetAllowedCountOk() (*int32, bool)`

GetAllowedCountOk returns a tuple with the AllowedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCount

`func (o *PatchedHaveIBeenPwendPolicyRequest) SetAllowedCount(v int32)`

SetAllowedCount sets AllowedCount field to given value.

### HasAllowedCount

`func (o *PatchedHaveIBeenPwendPolicyRequest) HasAllowedCount() bool`

HasAllowedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


