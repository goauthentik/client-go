# PasswordPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ExecutionLogging** | Pointer to **bool** | When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged. | [optional] 
**PasswordField** | Pointer to **string** | Field key to check, field keys defined in Prompt stages are available. | [optional] 
**AmountDigits** | Pointer to **int32** |  | [optional] 
**AmountUppercase** | Pointer to **int32** |  | [optional] 
**AmountLowercase** | Pointer to **int32** |  | [optional] 
**AmountSymbols** | Pointer to **int32** |  | [optional] 
**LengthMin** | Pointer to **int32** |  | [optional] 
**SymbolCharset** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**CheckStaticRules** | Pointer to **bool** |  | [optional] 
**CheckHaveIBeenPwned** | Pointer to **bool** |  | [optional] 
**CheckZxcvbn** | Pointer to **bool** |  | [optional] 
**HibpAllowedCount** | Pointer to **int32** | How many times the password hash is allowed to be on haveibeenpwned | [optional] 
**ZxcvbnScoreThreshold** | Pointer to **int32** | If the zxcvbn score is equal or less than this value, the policy will fail. | [optional] 

## Methods

### NewPasswordPolicyRequest

`func NewPasswordPolicyRequest(name string, ) *PasswordPolicyRequest`

NewPasswordPolicyRequest instantiates a new PasswordPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyRequestWithDefaults

`func NewPasswordPolicyRequestWithDefaults() *PasswordPolicyRequest`

NewPasswordPolicyRequestWithDefaults instantiates a new PasswordPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PasswordPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExecutionLogging

`func (o *PasswordPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *PasswordPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *PasswordPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *PasswordPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetPasswordField

`func (o *PasswordPolicyRequest) GetPasswordField() string`

GetPasswordField returns the PasswordField field if non-nil, zero value otherwise.

### GetPasswordFieldOk

`func (o *PasswordPolicyRequest) GetPasswordFieldOk() (*string, bool)`

GetPasswordFieldOk returns a tuple with the PasswordField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordField

`func (o *PasswordPolicyRequest) SetPasswordField(v string)`

SetPasswordField sets PasswordField field to given value.

### HasPasswordField

`func (o *PasswordPolicyRequest) HasPasswordField() bool`

HasPasswordField returns a boolean if a field has been set.

### GetAmountDigits

`func (o *PasswordPolicyRequest) GetAmountDigits() int32`

GetAmountDigits returns the AmountDigits field if non-nil, zero value otherwise.

### GetAmountDigitsOk

`func (o *PasswordPolicyRequest) GetAmountDigitsOk() (*int32, bool)`

GetAmountDigitsOk returns a tuple with the AmountDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDigits

`func (o *PasswordPolicyRequest) SetAmountDigits(v int32)`

SetAmountDigits sets AmountDigits field to given value.

### HasAmountDigits

`func (o *PasswordPolicyRequest) HasAmountDigits() bool`

HasAmountDigits returns a boolean if a field has been set.

### GetAmountUppercase

`func (o *PasswordPolicyRequest) GetAmountUppercase() int32`

GetAmountUppercase returns the AmountUppercase field if non-nil, zero value otherwise.

### GetAmountUppercaseOk

`func (o *PasswordPolicyRequest) GetAmountUppercaseOk() (*int32, bool)`

GetAmountUppercaseOk returns a tuple with the AmountUppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountUppercase

`func (o *PasswordPolicyRequest) SetAmountUppercase(v int32)`

SetAmountUppercase sets AmountUppercase field to given value.

### HasAmountUppercase

`func (o *PasswordPolicyRequest) HasAmountUppercase() bool`

HasAmountUppercase returns a boolean if a field has been set.

### GetAmountLowercase

`func (o *PasswordPolicyRequest) GetAmountLowercase() int32`

GetAmountLowercase returns the AmountLowercase field if non-nil, zero value otherwise.

### GetAmountLowercaseOk

`func (o *PasswordPolicyRequest) GetAmountLowercaseOk() (*int32, bool)`

GetAmountLowercaseOk returns a tuple with the AmountLowercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLowercase

`func (o *PasswordPolicyRequest) SetAmountLowercase(v int32)`

SetAmountLowercase sets AmountLowercase field to given value.

### HasAmountLowercase

`func (o *PasswordPolicyRequest) HasAmountLowercase() bool`

HasAmountLowercase returns a boolean if a field has been set.

### GetAmountSymbols

`func (o *PasswordPolicyRequest) GetAmountSymbols() int32`

GetAmountSymbols returns the AmountSymbols field if non-nil, zero value otherwise.

### GetAmountSymbolsOk

`func (o *PasswordPolicyRequest) GetAmountSymbolsOk() (*int32, bool)`

GetAmountSymbolsOk returns a tuple with the AmountSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSymbols

`func (o *PasswordPolicyRequest) SetAmountSymbols(v int32)`

SetAmountSymbols sets AmountSymbols field to given value.

### HasAmountSymbols

`func (o *PasswordPolicyRequest) HasAmountSymbols() bool`

HasAmountSymbols returns a boolean if a field has been set.

### GetLengthMin

`func (o *PasswordPolicyRequest) GetLengthMin() int32`

GetLengthMin returns the LengthMin field if non-nil, zero value otherwise.

### GetLengthMinOk

`func (o *PasswordPolicyRequest) GetLengthMinOk() (*int32, bool)`

GetLengthMinOk returns a tuple with the LengthMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthMin

`func (o *PasswordPolicyRequest) SetLengthMin(v int32)`

SetLengthMin sets LengthMin field to given value.

### HasLengthMin

`func (o *PasswordPolicyRequest) HasLengthMin() bool`

HasLengthMin returns a boolean if a field has been set.

### GetSymbolCharset

`func (o *PasswordPolicyRequest) GetSymbolCharset() string`

GetSymbolCharset returns the SymbolCharset field if non-nil, zero value otherwise.

### GetSymbolCharsetOk

`func (o *PasswordPolicyRequest) GetSymbolCharsetOk() (*string, bool)`

GetSymbolCharsetOk returns a tuple with the SymbolCharset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolCharset

`func (o *PasswordPolicyRequest) SetSymbolCharset(v string)`

SetSymbolCharset sets SymbolCharset field to given value.

### HasSymbolCharset

`func (o *PasswordPolicyRequest) HasSymbolCharset() bool`

HasSymbolCharset returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PasswordPolicyRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PasswordPolicyRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PasswordPolicyRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PasswordPolicyRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCheckStaticRules

`func (o *PasswordPolicyRequest) GetCheckStaticRules() bool`

GetCheckStaticRules returns the CheckStaticRules field if non-nil, zero value otherwise.

### GetCheckStaticRulesOk

`func (o *PasswordPolicyRequest) GetCheckStaticRulesOk() (*bool, bool)`

GetCheckStaticRulesOk returns a tuple with the CheckStaticRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStaticRules

`func (o *PasswordPolicyRequest) SetCheckStaticRules(v bool)`

SetCheckStaticRules sets CheckStaticRules field to given value.

### HasCheckStaticRules

`func (o *PasswordPolicyRequest) HasCheckStaticRules() bool`

HasCheckStaticRules returns a boolean if a field has been set.

### GetCheckHaveIBeenPwned

`func (o *PasswordPolicyRequest) GetCheckHaveIBeenPwned() bool`

GetCheckHaveIBeenPwned returns the CheckHaveIBeenPwned field if non-nil, zero value otherwise.

### GetCheckHaveIBeenPwnedOk

`func (o *PasswordPolicyRequest) GetCheckHaveIBeenPwnedOk() (*bool, bool)`

GetCheckHaveIBeenPwnedOk returns a tuple with the CheckHaveIBeenPwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckHaveIBeenPwned

`func (o *PasswordPolicyRequest) SetCheckHaveIBeenPwned(v bool)`

SetCheckHaveIBeenPwned sets CheckHaveIBeenPwned field to given value.

### HasCheckHaveIBeenPwned

`func (o *PasswordPolicyRequest) HasCheckHaveIBeenPwned() bool`

HasCheckHaveIBeenPwned returns a boolean if a field has been set.

### GetCheckZxcvbn

`func (o *PasswordPolicyRequest) GetCheckZxcvbn() bool`

GetCheckZxcvbn returns the CheckZxcvbn field if non-nil, zero value otherwise.

### GetCheckZxcvbnOk

`func (o *PasswordPolicyRequest) GetCheckZxcvbnOk() (*bool, bool)`

GetCheckZxcvbnOk returns a tuple with the CheckZxcvbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckZxcvbn

`func (o *PasswordPolicyRequest) SetCheckZxcvbn(v bool)`

SetCheckZxcvbn sets CheckZxcvbn field to given value.

### HasCheckZxcvbn

`func (o *PasswordPolicyRequest) HasCheckZxcvbn() bool`

HasCheckZxcvbn returns a boolean if a field has been set.

### GetHibpAllowedCount

`func (o *PasswordPolicyRequest) GetHibpAllowedCount() int32`

GetHibpAllowedCount returns the HibpAllowedCount field if non-nil, zero value otherwise.

### GetHibpAllowedCountOk

`func (o *PasswordPolicyRequest) GetHibpAllowedCountOk() (*int32, bool)`

GetHibpAllowedCountOk returns a tuple with the HibpAllowedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHibpAllowedCount

`func (o *PasswordPolicyRequest) SetHibpAllowedCount(v int32)`

SetHibpAllowedCount sets HibpAllowedCount field to given value.

### HasHibpAllowedCount

`func (o *PasswordPolicyRequest) HasHibpAllowedCount() bool`

HasHibpAllowedCount returns a boolean if a field has been set.

### GetZxcvbnScoreThreshold

`func (o *PasswordPolicyRequest) GetZxcvbnScoreThreshold() int32`

GetZxcvbnScoreThreshold returns the ZxcvbnScoreThreshold field if non-nil, zero value otherwise.

### GetZxcvbnScoreThresholdOk

`func (o *PasswordPolicyRequest) GetZxcvbnScoreThresholdOk() (*int32, bool)`

GetZxcvbnScoreThresholdOk returns a tuple with the ZxcvbnScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZxcvbnScoreThreshold

`func (o *PasswordPolicyRequest) SetZxcvbnScoreThreshold(v int32)`

SetZxcvbnScoreThreshold sets ZxcvbnScoreThreshold field to given value.

### HasZxcvbnScoreThreshold

`func (o *PasswordPolicyRequest) HasZxcvbnScoreThreshold() bool`

HasZxcvbnScoreThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


