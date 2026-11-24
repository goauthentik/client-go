# PatchedPasswordPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
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

### NewPatchedPasswordPolicyRequest

`func NewPatchedPasswordPolicyRequest() *PatchedPasswordPolicyRequest`

NewPatchedPasswordPolicyRequest instantiates a new PatchedPasswordPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedPasswordPolicyRequestWithDefaults

`func NewPatchedPasswordPolicyRequestWithDefaults() *PatchedPasswordPolicyRequest`

NewPatchedPasswordPolicyRequestWithDefaults instantiates a new PatchedPasswordPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedPasswordPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedPasswordPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedPasswordPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedPasswordPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExecutionLogging

`func (o *PatchedPasswordPolicyRequest) GetExecutionLogging() bool`

GetExecutionLogging returns the ExecutionLogging field if non-nil, zero value otherwise.

### GetExecutionLoggingOk

`func (o *PatchedPasswordPolicyRequest) GetExecutionLoggingOk() (*bool, bool)`

GetExecutionLoggingOk returns a tuple with the ExecutionLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLogging

`func (o *PatchedPasswordPolicyRequest) SetExecutionLogging(v bool)`

SetExecutionLogging sets ExecutionLogging field to given value.

### HasExecutionLogging

`func (o *PatchedPasswordPolicyRequest) HasExecutionLogging() bool`

HasExecutionLogging returns a boolean if a field has been set.

### GetPasswordField

`func (o *PatchedPasswordPolicyRequest) GetPasswordField() string`

GetPasswordField returns the PasswordField field if non-nil, zero value otherwise.

### GetPasswordFieldOk

`func (o *PatchedPasswordPolicyRequest) GetPasswordFieldOk() (*string, bool)`

GetPasswordFieldOk returns a tuple with the PasswordField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordField

`func (o *PatchedPasswordPolicyRequest) SetPasswordField(v string)`

SetPasswordField sets PasswordField field to given value.

### HasPasswordField

`func (o *PatchedPasswordPolicyRequest) HasPasswordField() bool`

HasPasswordField returns a boolean if a field has been set.

### GetAmountDigits

`func (o *PatchedPasswordPolicyRequest) GetAmountDigits() int32`

GetAmountDigits returns the AmountDigits field if non-nil, zero value otherwise.

### GetAmountDigitsOk

`func (o *PatchedPasswordPolicyRequest) GetAmountDigitsOk() (*int32, bool)`

GetAmountDigitsOk returns a tuple with the AmountDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDigits

`func (o *PatchedPasswordPolicyRequest) SetAmountDigits(v int32)`

SetAmountDigits sets AmountDigits field to given value.

### HasAmountDigits

`func (o *PatchedPasswordPolicyRequest) HasAmountDigits() bool`

HasAmountDigits returns a boolean if a field has been set.

### GetAmountUppercase

`func (o *PatchedPasswordPolicyRequest) GetAmountUppercase() int32`

GetAmountUppercase returns the AmountUppercase field if non-nil, zero value otherwise.

### GetAmountUppercaseOk

`func (o *PatchedPasswordPolicyRequest) GetAmountUppercaseOk() (*int32, bool)`

GetAmountUppercaseOk returns a tuple with the AmountUppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountUppercase

`func (o *PatchedPasswordPolicyRequest) SetAmountUppercase(v int32)`

SetAmountUppercase sets AmountUppercase field to given value.

### HasAmountUppercase

`func (o *PatchedPasswordPolicyRequest) HasAmountUppercase() bool`

HasAmountUppercase returns a boolean if a field has been set.

### GetAmountLowercase

`func (o *PatchedPasswordPolicyRequest) GetAmountLowercase() int32`

GetAmountLowercase returns the AmountLowercase field if non-nil, zero value otherwise.

### GetAmountLowercaseOk

`func (o *PatchedPasswordPolicyRequest) GetAmountLowercaseOk() (*int32, bool)`

GetAmountLowercaseOk returns a tuple with the AmountLowercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLowercase

`func (o *PatchedPasswordPolicyRequest) SetAmountLowercase(v int32)`

SetAmountLowercase sets AmountLowercase field to given value.

### HasAmountLowercase

`func (o *PatchedPasswordPolicyRequest) HasAmountLowercase() bool`

HasAmountLowercase returns a boolean if a field has been set.

### GetAmountSymbols

`func (o *PatchedPasswordPolicyRequest) GetAmountSymbols() int32`

GetAmountSymbols returns the AmountSymbols field if non-nil, zero value otherwise.

### GetAmountSymbolsOk

`func (o *PatchedPasswordPolicyRequest) GetAmountSymbolsOk() (*int32, bool)`

GetAmountSymbolsOk returns a tuple with the AmountSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSymbols

`func (o *PatchedPasswordPolicyRequest) SetAmountSymbols(v int32)`

SetAmountSymbols sets AmountSymbols field to given value.

### HasAmountSymbols

`func (o *PatchedPasswordPolicyRequest) HasAmountSymbols() bool`

HasAmountSymbols returns a boolean if a field has been set.

### GetLengthMin

`func (o *PatchedPasswordPolicyRequest) GetLengthMin() int32`

GetLengthMin returns the LengthMin field if non-nil, zero value otherwise.

### GetLengthMinOk

`func (o *PatchedPasswordPolicyRequest) GetLengthMinOk() (*int32, bool)`

GetLengthMinOk returns a tuple with the LengthMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthMin

`func (o *PatchedPasswordPolicyRequest) SetLengthMin(v int32)`

SetLengthMin sets LengthMin field to given value.

### HasLengthMin

`func (o *PatchedPasswordPolicyRequest) HasLengthMin() bool`

HasLengthMin returns a boolean if a field has been set.

### GetSymbolCharset

`func (o *PatchedPasswordPolicyRequest) GetSymbolCharset() string`

GetSymbolCharset returns the SymbolCharset field if non-nil, zero value otherwise.

### GetSymbolCharsetOk

`func (o *PatchedPasswordPolicyRequest) GetSymbolCharsetOk() (*string, bool)`

GetSymbolCharsetOk returns a tuple with the SymbolCharset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolCharset

`func (o *PatchedPasswordPolicyRequest) SetSymbolCharset(v string)`

SetSymbolCharset sets SymbolCharset field to given value.

### HasSymbolCharset

`func (o *PatchedPasswordPolicyRequest) HasSymbolCharset() bool`

HasSymbolCharset returns a boolean if a field has been set.

### GetErrorMessage

`func (o *PatchedPasswordPolicyRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PatchedPasswordPolicyRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PatchedPasswordPolicyRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PatchedPasswordPolicyRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCheckStaticRules

`func (o *PatchedPasswordPolicyRequest) GetCheckStaticRules() bool`

GetCheckStaticRules returns the CheckStaticRules field if non-nil, zero value otherwise.

### GetCheckStaticRulesOk

`func (o *PatchedPasswordPolicyRequest) GetCheckStaticRulesOk() (*bool, bool)`

GetCheckStaticRulesOk returns a tuple with the CheckStaticRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStaticRules

`func (o *PatchedPasswordPolicyRequest) SetCheckStaticRules(v bool)`

SetCheckStaticRules sets CheckStaticRules field to given value.

### HasCheckStaticRules

`func (o *PatchedPasswordPolicyRequest) HasCheckStaticRules() bool`

HasCheckStaticRules returns a boolean if a field has been set.

### GetCheckHaveIBeenPwned

`func (o *PatchedPasswordPolicyRequest) GetCheckHaveIBeenPwned() bool`

GetCheckHaveIBeenPwned returns the CheckHaveIBeenPwned field if non-nil, zero value otherwise.

### GetCheckHaveIBeenPwnedOk

`func (o *PatchedPasswordPolicyRequest) GetCheckHaveIBeenPwnedOk() (*bool, bool)`

GetCheckHaveIBeenPwnedOk returns a tuple with the CheckHaveIBeenPwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckHaveIBeenPwned

`func (o *PatchedPasswordPolicyRequest) SetCheckHaveIBeenPwned(v bool)`

SetCheckHaveIBeenPwned sets CheckHaveIBeenPwned field to given value.

### HasCheckHaveIBeenPwned

`func (o *PatchedPasswordPolicyRequest) HasCheckHaveIBeenPwned() bool`

HasCheckHaveIBeenPwned returns a boolean if a field has been set.

### GetCheckZxcvbn

`func (o *PatchedPasswordPolicyRequest) GetCheckZxcvbn() bool`

GetCheckZxcvbn returns the CheckZxcvbn field if non-nil, zero value otherwise.

### GetCheckZxcvbnOk

`func (o *PatchedPasswordPolicyRequest) GetCheckZxcvbnOk() (*bool, bool)`

GetCheckZxcvbnOk returns a tuple with the CheckZxcvbn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckZxcvbn

`func (o *PatchedPasswordPolicyRequest) SetCheckZxcvbn(v bool)`

SetCheckZxcvbn sets CheckZxcvbn field to given value.

### HasCheckZxcvbn

`func (o *PatchedPasswordPolicyRequest) HasCheckZxcvbn() bool`

HasCheckZxcvbn returns a boolean if a field has been set.

### GetHibpAllowedCount

`func (o *PatchedPasswordPolicyRequest) GetHibpAllowedCount() int32`

GetHibpAllowedCount returns the HibpAllowedCount field if non-nil, zero value otherwise.

### GetHibpAllowedCountOk

`func (o *PatchedPasswordPolicyRequest) GetHibpAllowedCountOk() (*int32, bool)`

GetHibpAllowedCountOk returns a tuple with the HibpAllowedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHibpAllowedCount

`func (o *PatchedPasswordPolicyRequest) SetHibpAllowedCount(v int32)`

SetHibpAllowedCount sets HibpAllowedCount field to given value.

### HasHibpAllowedCount

`func (o *PatchedPasswordPolicyRequest) HasHibpAllowedCount() bool`

HasHibpAllowedCount returns a boolean if a field has been set.

### GetZxcvbnScoreThreshold

`func (o *PatchedPasswordPolicyRequest) GetZxcvbnScoreThreshold() int32`

GetZxcvbnScoreThreshold returns the ZxcvbnScoreThreshold field if non-nil, zero value otherwise.

### GetZxcvbnScoreThresholdOk

`func (o *PatchedPasswordPolicyRequest) GetZxcvbnScoreThresholdOk() (*int32, bool)`

GetZxcvbnScoreThresholdOk returns a tuple with the ZxcvbnScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZxcvbnScoreThreshold

`func (o *PatchedPasswordPolicyRequest) SetZxcvbnScoreThreshold(v int32)`

SetZxcvbnScoreThreshold sets ZxcvbnScoreThreshold field to given value.

### HasZxcvbnScoreThreshold

`func (o *PatchedPasswordPolicyRequest) HasZxcvbnScoreThreshold() bool`

HasZxcvbnScoreThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


