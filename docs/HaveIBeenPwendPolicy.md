# HaveIBeenPwendPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Component** | **string** |  | [readonly] 
**VerboseName** | **string** |  | [readonly] 
**VerboseNamePlural** | **string** |  | [readonly] 
**BoundTo** | **int32** |  | [readonly] 
**PasswordField** | Pointer to **string** | Field key to check, field keys defined in Prompt stages are available. | [optional] 
**AllowedCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewHaveIBeenPwendPolicy

`func NewHaveIBeenPwendPolicy(pk string, component string, verboseName string, verboseNamePlural string, boundTo int32, ) *HaveIBeenPwendPolicy`

NewHaveIBeenPwendPolicy instantiates a new HaveIBeenPwendPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHaveIBeenPwendPolicyWithDefaults

`func NewHaveIBeenPwendPolicyWithDefaults() *HaveIBeenPwendPolicy`

NewHaveIBeenPwendPolicyWithDefaults instantiates a new HaveIBeenPwendPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *HaveIBeenPwendPolicy) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *HaveIBeenPwendPolicy) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *HaveIBeenPwendPolicy) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *HaveIBeenPwendPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HaveIBeenPwendPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HaveIBeenPwendPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HaveIBeenPwendPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HaveIBeenPwendPolicy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HaveIBeenPwendPolicy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetExecutionLogging

`func (o *HaveIBeenPwendPolicy) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *HaveIBeenPwendPolicy) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *HaveIBeenPwendPolicy) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *HaveIBeenPwendPolicy) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetComponent

`func (o *HaveIBeenPwendPolicy) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *HaveIBeenPwendPolicy) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *HaveIBeenPwendPolicy) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *HaveIBeenPwendPolicy) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *HaveIBeenPwendPolicy) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *HaveIBeenPwendPolicy) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *HaveIBeenPwendPolicy) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *HaveIBeenPwendPolicy) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *HaveIBeenPwendPolicy) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetBoundTo

`func (o *HaveIBeenPwendPolicy) GetBoundTo() int32`

GetBoundTo returns the BoundTo field if non-nil, zero value otherwise.

### GetBoundToOk

`func (o *HaveIBeenPwendPolicy) GetBoundToOk() (*int32, bool)`

GetBoundToOk returns a tuple with the BoundTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTo

`func (o *HaveIBeenPwendPolicy) SetBoundTo(v int32)`

SetBoundTo sets BoundTo field to given value.


### GetPasswordField

`func (o *HaveIBeenPwendPolicy) GetPasswordField() string`

GetPasswordField returns the PasswordField field if non-nil, zero value otherwise.

### GetPasswordFieldOk

`func (o *HaveIBeenPwendPolicy) GetPasswordFieldOk() (*string, bool)`

GetPasswordFieldOk returns a tuple with the PasswordField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordField

`func (o *HaveIBeenPwendPolicy) SetPasswordField(v string)`

SetPasswordField sets PasswordField field to given value.

### HasPasswordField

`func (o *HaveIBeenPwendPolicy) HasPasswordField() bool`

HasPasswordField returns a boolean if a field has been set.

### GetAllowedCount

`func (o *HaveIBeenPwendPolicy) GetAllowedCount() int32`

GetAllowedCount returns the AllowedCount field if non-nil, zero value otherwise.

### GetAllowedCountOk

`func (o *HaveIBeenPwendPolicy) GetAllowedCountOk() (*int32, bool)`

GetAllowedCountOk returns a tuple with the AllowedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCount

`func (o *HaveIBeenPwendPolicy) SetAllowedCount(v int32)`

SetAllowedCount sets AllowedCount field to given value.

### HasAllowedCount

`func (o *HaveIBeenPwendPolicy) HasAllowedCount() bool`

HasAllowedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


