# PasswordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**Component** | **string** |  | [readonly] 
**VerboseName** | **string** |  | [readonly] 
**VerboseNamePlural** | **string** |  | [readonly] 
**MetaModelName** | **string** |  | [readonly] 
**BoundTo** | **int32** |  | [readonly] 
**PasswordField** | Pointer to **string** | Field key to check, field keys defined in Prompt stages are available. | [optional] 
**AmountDigits** | Pointer to **int32** |  | [optional] 
**AmountUppercase** | Pointer to **int32** |  | [optional] 
**AmountLowercase** | Pointer to **int32** |  | [optional] 
**AmountSymbols** | Pointer to **int32** |  | [optional] 
**LengthMin** | Pointer to **int32** |  | [optional] 
**SymbolCharset** | Pointer to **string** |  | [optional] 
**ErrorMessage** | **string** |  | 

## Methods

### NewPasswordPolicy

`func NewPasswordPolicy(pk string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32, errorMessage string, ) *PasswordPolicy`

NewPasswordPolicy instantiates a new PasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyWithDefaults

`func NewPasswordPolicyWithDefaults() *PasswordPolicy`

NewPasswordPolicyWithDefaults instantiates a new PasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *PasswordPolicy) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *PasswordPolicy) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *PasswordPolicy) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *PasswordPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PasswordPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PasswordPolicy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PasswordPolicy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetExecutionLogging

`func (o *PasswordPolicy) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *PasswordPolicy) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *PasswordPolicy) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *PasswordPolicy) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetComponent

`func (o *PasswordPolicy) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PasswordPolicy) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PasswordPolicy) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *PasswordPolicy) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *PasswordPolicy) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *PasswordPolicy) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *PasswordPolicy) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *PasswordPolicy) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *PasswordPolicy) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *PasswordPolicy) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *PasswordPolicy) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *PasswordPolicy) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetBoundTo

`func (o *PasswordPolicy) GetBoundTo() int32`

GetBoundTo returns the BoundTo field if non-nil, zero value otherwise.

### GetBoundToOk

`func (o *PasswordPolicy) GetBoundToOk() (*int32, bool)`

GetBoundToOk returns a tuple with the BoundTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundTo

`func (o *PasswordPolicy) SetBoundTo(v int32)`

SetBoundTo sets BoundTo field to given value.


### GetPasswordField

`func (o *PasswordPolicy) GetPasswordField() string`

GetPasswordField returns the PasswordField field if non-nil, zero value otherwise.

### GetPasswordFieldOk

`func (o *PasswordPolicy) GetPasswordFieldOk() (*string, bool)`

GetPasswordFieldOk returns a tuple with the PasswordField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordField

`func (o *PasswordPolicy) SetPasswordField(v string)`

SetPasswordField sets PasswordField field to given value.

### HasPasswordField

`func (o *PasswordPolicy) HasPasswordField() bool`

HasPasswordField returns a boolean if a field has been set.

### GetAmountDigits

`func (o *PasswordPolicy) GetAmountDigits() int32`

GetAmountDigits returns the AmountDigits field if non-nil, zero value otherwise.

### GetAmountDigitsOk

`func (o *PasswordPolicy) GetAmountDigitsOk() (*int32, bool)`

GetAmountDigitsOk returns a tuple with the AmountDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDigits

`func (o *PasswordPolicy) SetAmountDigits(v int32)`

SetAmountDigits sets AmountDigits field to given value.

### HasAmountDigits

`func (o *PasswordPolicy) HasAmountDigits() bool`

HasAmountDigits returns a boolean if a field has been set.

### GetAmountUppercase

`func (o *PasswordPolicy) GetAmountUppercase() int32`

GetAmountUppercase returns the AmountUppercase field if non-nil, zero value otherwise.

### GetAmountUppercaseOk

`func (o *PasswordPolicy) GetAmountUppercaseOk() (*int32, bool)`

GetAmountUppercaseOk returns a tuple with the AmountUppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountUppercase

`func (o *PasswordPolicy) SetAmountUppercase(v int32)`

SetAmountUppercase sets AmountUppercase field to given value.

### HasAmountUppercase

`func (o *PasswordPolicy) HasAmountUppercase() bool`

HasAmountUppercase returns a boolean if a field has been set.

### GetAmountLowercase

`func (o *PasswordPolicy) GetAmountLowercase() int32`

GetAmountLowercase returns the AmountLowercase field if non-nil, zero value otherwise.

### GetAmountLowercaseOk

`func (o *PasswordPolicy) GetAmountLowercaseOk() (*int32, bool)`

GetAmountLowercaseOk returns a tuple with the AmountLowercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLowercase

`func (o *PasswordPolicy) SetAmountLowercase(v int32)`

SetAmountLowercase sets AmountLowercase field to given value.

### HasAmountLowercase

`func (o *PasswordPolicy) HasAmountLowercase() bool`

HasAmountLowercase returns a boolean if a field has been set.

### GetAmountSymbols

`func (o *PasswordPolicy) GetAmountSymbols() int32`

GetAmountSymbols returns the AmountSymbols field if non-nil, zero value otherwise.

### GetAmountSymbolsOk

`func (o *PasswordPolicy) GetAmountSymbolsOk() (*int32, bool)`

GetAmountSymbolsOk returns a tuple with the AmountSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSymbols

`func (o *PasswordPolicy) SetAmountSymbols(v int32)`

SetAmountSymbols sets AmountSymbols field to given value.

### HasAmountSymbols

`func (o *PasswordPolicy) HasAmountSymbols() bool`

HasAmountSymbols returns a boolean if a field has been set.

### GetLengthMin

`func (o *PasswordPolicy) GetLengthMin() int32`

GetLengthMin returns the LengthMin field if non-nil, zero value otherwise.

### GetLengthMinOk

`func (o *PasswordPolicy) GetLengthMinOk() (*int32, bool)`

GetLengthMinOk returns a tuple with the LengthMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthMin

`func (o *PasswordPolicy) SetLengthMin(v int32)`

SetLengthMin sets LengthMin field to given value.

### HasLengthMin

`func (o *PasswordPolicy) HasLengthMin() bool`

HasLengthMin returns a boolean if a field has been set.

### GetSymbolCharset

`func (o *PasswordPolicy) GetSymbolCharset() string`

GetSymbolCharset returns the SymbolCharset field if non-nil, zero value otherwise.

### GetSymbolCharsetOk

`func (o *PasswordPolicy) GetSymbolCharsetOk() (*string, bool)`

GetSymbolCharsetOk returns a tuple with the SymbolCharset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolCharset

`func (o *PasswordPolicy) SetSymbolCharset(v string)`

SetSymbolCharset sets SymbolCharset field to given value.

### HasSymbolCharset

`func (o *PasswordPolicy) HasSymbolCharset() bool`

HasSymbolCharset returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PasswordPolicy) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PasswordPolicy) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PasswordPolicy) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


