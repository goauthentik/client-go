# PatchedUniquePasswordPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**PasswordField** | Pointer to **string** | Field key to check, field keys defined in Prompt stages are available. | [optional] 
**NumHistoricalPasswords** | Pointer to **int32** | Number of passwords to check against. | [optional] 

## Methods

### NewPatchedUniquePasswordPolicyRequest

`func NewPatchedUniquePasswordPolicyRequest() *PatchedUniquePasswordPolicyRequest`

NewPatchedUniquePasswordPolicyRequest instantiates a new PatchedUniquePasswordPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUniquePasswordPolicyRequestWithDefaults

`func NewPatchedUniquePasswordPolicyRequestWithDefaults() *PatchedUniquePasswordPolicyRequest`

NewPatchedUniquePasswordPolicyRequestWithDefaults instantiates a new PatchedUniquePasswordPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedUniquePasswordPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedUniquePasswordPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedUniquePasswordPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedUniquePasswordPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExecutionLogging

`func (o *PatchedUniquePasswordPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *PatchedUniquePasswordPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *PatchedUniquePasswordPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *PatchedUniquePasswordPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetPasswordField

`func (o *PatchedUniquePasswordPolicyRequest) GetPasswordField() string`

GetPasswordField returns the PasswordField field if non-nil, zero value otherwise.

### GetPasswordFieldOk

`func (o *PatchedUniquePasswordPolicyRequest) GetPasswordFieldOk() (*string, bool)`

GetPasswordFieldOk returns a tuple with the PasswordField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordField

`func (o *PatchedUniquePasswordPolicyRequest) SetPasswordField(v string)`

SetPasswordField sets PasswordField field to given value.

### HasPasswordField

`func (o *PatchedUniquePasswordPolicyRequest) HasPasswordField() bool`

HasPasswordField returns a boolean if a field has been set.

### GetNumHistoricalPasswords

`func (o *PatchedUniquePasswordPolicyRequest) GetNumHistoricalPasswords() int32`

GetNumHistoricalPasswords returns the NumHistoricalPasswords field if non-nil, zero value otherwise.

### GetNumHistoricalPasswordsOk

`func (o *PatchedUniquePasswordPolicyRequest) GetNumHistoricalPasswordsOk() (*int32, bool)`

GetNumHistoricalPasswordsOk returns a tuple with the NumHistoricalPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHistoricalPasswords

`func (o *PatchedUniquePasswordPolicyRequest) SetNumHistoricalPasswords(v int32)`

SetNumHistoricalPasswords sets NumHistoricalPasswords field to given value.

### HasNumHistoricalPasswords

`func (o *PatchedUniquePasswordPolicyRequest) HasNumHistoricalPasswords() bool`

HasNumHistoricalPasswords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


