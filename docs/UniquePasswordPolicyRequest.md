# UniquePasswordPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**PasswordField** | Pointer to **string** | Field key to check, field keys defined in Prompt stages are available. | [optional] 
**NumHistoricalPasswords** | Pointer to **int32** | Number of passwords to check against. | [optional] 

## Methods

### NewUniquePasswordPolicyRequest

`func NewUniquePasswordPolicyRequest(name string, ) *UniquePasswordPolicyRequest`

NewUniquePasswordPolicyRequest instantiates a new UniquePasswordPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniquePasswordPolicyRequestWithDefaults

`func NewUniquePasswordPolicyRequestWithDefaults() *UniquePasswordPolicyRequest`

NewUniquePasswordPolicyRequestWithDefaults instantiates a new UniquePasswordPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UniquePasswordPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniquePasswordPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniquePasswordPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *UniquePasswordPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *UniquePasswordPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *UniquePasswordPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *UniquePasswordPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetPasswordField

`func (o *UniquePasswordPolicyRequest) GetPasswordField() string`

GetPasswordField returns the PasswordField field if non-nil, zero value otherwise.

### GetPasswordFieldOk

`func (o *UniquePasswordPolicyRequest) GetPasswordFieldOk() (*string, bool)`

GetPasswordFieldOk returns a tuple with the PasswordField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordField

`func (o *UniquePasswordPolicyRequest) SetPasswordField(v string)`

SetPasswordField sets PasswordField field to given value.

### HasPasswordField

`func (o *UniquePasswordPolicyRequest) HasPasswordField() bool`

HasPasswordField returns a boolean if a field has been set.

### GetNumHistoricalPasswords

`func (o *UniquePasswordPolicyRequest) GetNumHistoricalPasswords() int32`

GetNumHistoricalPasswords returns the NumHistoricalPasswords field if non-nil, zero value otherwise.

### GetNumHistoricalPasswordsOk

`func (o *UniquePasswordPolicyRequest) GetNumHistoricalPasswordsOk() (*int32, bool)`

GetNumHistoricalPasswordsOk returns a tuple with the NumHistoricalPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHistoricalPasswords

`func (o *UniquePasswordPolicyRequest) SetNumHistoricalPasswords(v int32)`

SetNumHistoricalPasswords sets NumHistoricalPasswords field to given value.

### HasNumHistoricalPasswords

`func (o *UniquePasswordPolicyRequest) HasNumHistoricalPasswords() bool`

HasNumHistoricalPasswords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


