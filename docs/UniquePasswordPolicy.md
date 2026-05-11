# UniquePasswordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**BoundTo** | **int32** | Return objects policy is bound to | [readonly] 
**PasswordField** | Pointer to **string** | Field key to check, field keys defined in Prompt stages are available. | [optional] 
**NumHistoricalPasswords** | Pointer to **int32** | Number of passwords to check against. | [optional] 

## Methods

### NewUniquePasswordPolicy

`func NewUniquePasswordPolicy(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32, ) *UniquePasswordPolicy`

NewUniquePasswordPolicy instantiates a new UniquePasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniquePasswordPolicyWithDefaults

`func NewUniquePasswordPolicyWithDefaults() *UniquePasswordPolicy`

NewUniquePasswordPolicyWithDefaults instantiates a new UniquePasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *UniquePasswordPolicy) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *UniquePasswordPolicy) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *UniquePasswordPolicy) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *UniquePasswordPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniquePasswordPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniquePasswordPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *UniquePasswordPolicy) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *UniquePasswordPolicy) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *UniquePasswordPolicy) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *UniquePasswordPolicy) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetComponent

`func (o *UniquePasswordPolicy) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *UniquePasswordPolicy) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *UniquePasswordPolicy) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *UniquePasswordPolicy) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *UniquePasswordPolicy) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *UniquePasswordPolicy) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *UniquePasswordPolicy) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *UniquePasswordPolicy) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *UniquePasswordPolicy) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *UniquePasswordPolicy) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *UniquePasswordPolicy) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *UniquePasswordPolicy) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetBoundTo

`func (o *UniquePasswordPolicy) GetBoundTo() int32`

GetBoundTo returns the BoundTo field if non-nil, zero value otherwise.

### GetBoundToOk

`func (o *UniquePasswordPolicy) GetBoundToOk() (*int32, bool)`

GetBoundToOk returns a tuple with the BoundTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTo

`func (o *UniquePasswordPolicy) SetBoundTo(v int32)`

SetBoundTo sets BoundTo field to given value.


### GetPasswordField

`func (o *UniquePasswordPolicy) GetPasswordField() string`

GetPasswordField returns the PasswordField field if non-nil, zero value otherwise.

### GetPasswordFieldOk

`func (o *UniquePasswordPolicy) GetPasswordFieldOk() (*string, bool)`

GetPasswordFieldOk returns a tuple with the PasswordField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordField

`func (o *UniquePasswordPolicy) SetPasswordField(v string)`

SetPasswordField sets PasswordField field to given value.

### HasPasswordField

`func (o *UniquePasswordPolicy) HasPasswordField() bool`

HasPasswordField returns a boolean if a field has been set.

### GetNumHistoricalPasswords

`func (o *UniquePasswordPolicy) GetNumHistoricalPasswords() int32`

GetNumHistoricalPasswords returns the NumHistoricalPasswords field if non-nil, zero value otherwise.

### GetNumHistoricalPasswordsOk

`func (o *UniquePasswordPolicy) GetNumHistoricalPasswordsOk() (*int32, bool)`

GetNumHistoricalPasswordsOk returns a tuple with the NumHistoricalPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHistoricalPasswords

`func (o *UniquePasswordPolicy) SetNumHistoricalPasswords(v int32)`

SetNumHistoricalPasswords sets NumHistoricalPasswords field to given value.

### HasNumHistoricalPasswords

`func (o *UniquePasswordPolicy) HasNumHistoricalPasswords() bool`

HasNumHistoricalPasswords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


