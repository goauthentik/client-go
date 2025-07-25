/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PasswordPolicyRequest Password Policy Serializer
type PasswordPolicyRequest struct {
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Field key to check, field keys defined in Prompt stages are available.
	PasswordField       *string `json:"password_field,omitempty"`
	AmountDigits        *int32  `json:"amount_digits,omitempty"`
	AmountUppercase     *int32  `json:"amount_uppercase,omitempty"`
	AmountLowercase     *int32  `json:"amount_lowercase,omitempty"`
	AmountSymbols       *int32  `json:"amount_symbols,omitempty"`
	LengthMin           *int32  `json:"length_min,omitempty"`
	SymbolCharset       *string `json:"symbol_charset,omitempty"`
	ErrorMessage        *string `json:"error_message,omitempty"`
	CheckStaticRules    *bool   `json:"check_static_rules,omitempty"`
	CheckHaveIBeenPwned *bool   `json:"check_have_i_been_pwned,omitempty"`
	CheckZxcvbn         *bool   `json:"check_zxcvbn,omitempty"`
	// How many times the password hash is allowed to be on haveibeenpwned
	HibpAllowedCount *int32 `json:"hibp_allowed_count,omitempty"`
	// If the zxcvbn score is equal or less than this value, the policy will fail.
	ZxcvbnScoreThreshold *int32 `json:"zxcvbn_score_threshold,omitempty"`
}

// NewPasswordPolicyRequest instantiates a new PasswordPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRequest(name string) *PasswordPolicyRequest {
	this := PasswordPolicyRequest{}
	this.Name = name
	return &this
}

// NewPasswordPolicyRequestWithDefaults instantiates a new PasswordPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRequestWithDefaults() *PasswordPolicyRequest {
	this := PasswordPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PasswordPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PasswordPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PasswordPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetPasswordField returns the PasswordField field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetPasswordField() string {
	if o == nil || o.PasswordField == nil {
		var ret string
		return ret
	}
	return *o.PasswordField
}

// GetPasswordFieldOk returns a tuple with the PasswordField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetPasswordFieldOk() (*string, bool) {
	if o == nil || o.PasswordField == nil {
		return nil, false
	}
	return o.PasswordField, true
}

// HasPasswordField returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasPasswordField() bool {
	if o != nil && o.PasswordField != nil {
		return true
	}

	return false
}

// SetPasswordField gets a reference to the given string and assigns it to the PasswordField field.
func (o *PasswordPolicyRequest) SetPasswordField(v string) {
	o.PasswordField = &v
}

// GetAmountDigits returns the AmountDigits field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetAmountDigits() int32 {
	if o == nil || o.AmountDigits == nil {
		var ret int32
		return ret
	}
	return *o.AmountDigits
}

// GetAmountDigitsOk returns a tuple with the AmountDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetAmountDigitsOk() (*int32, bool) {
	if o == nil || o.AmountDigits == nil {
		return nil, false
	}
	return o.AmountDigits, true
}

// HasAmountDigits returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasAmountDigits() bool {
	if o != nil && o.AmountDigits != nil {
		return true
	}

	return false
}

// SetAmountDigits gets a reference to the given int32 and assigns it to the AmountDigits field.
func (o *PasswordPolicyRequest) SetAmountDigits(v int32) {
	o.AmountDigits = &v
}

// GetAmountUppercase returns the AmountUppercase field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetAmountUppercase() int32 {
	if o == nil || o.AmountUppercase == nil {
		var ret int32
		return ret
	}
	return *o.AmountUppercase
}

// GetAmountUppercaseOk returns a tuple with the AmountUppercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetAmountUppercaseOk() (*int32, bool) {
	if o == nil || o.AmountUppercase == nil {
		return nil, false
	}
	return o.AmountUppercase, true
}

// HasAmountUppercase returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasAmountUppercase() bool {
	if o != nil && o.AmountUppercase != nil {
		return true
	}

	return false
}

// SetAmountUppercase gets a reference to the given int32 and assigns it to the AmountUppercase field.
func (o *PasswordPolicyRequest) SetAmountUppercase(v int32) {
	o.AmountUppercase = &v
}

// GetAmountLowercase returns the AmountLowercase field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetAmountLowercase() int32 {
	if o == nil || o.AmountLowercase == nil {
		var ret int32
		return ret
	}
	return *o.AmountLowercase
}

// GetAmountLowercaseOk returns a tuple with the AmountLowercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetAmountLowercaseOk() (*int32, bool) {
	if o == nil || o.AmountLowercase == nil {
		return nil, false
	}
	return o.AmountLowercase, true
}

// HasAmountLowercase returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasAmountLowercase() bool {
	if o != nil && o.AmountLowercase != nil {
		return true
	}

	return false
}

// SetAmountLowercase gets a reference to the given int32 and assigns it to the AmountLowercase field.
func (o *PasswordPolicyRequest) SetAmountLowercase(v int32) {
	o.AmountLowercase = &v
}

// GetAmountSymbols returns the AmountSymbols field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetAmountSymbols() int32 {
	if o == nil || o.AmountSymbols == nil {
		var ret int32
		return ret
	}
	return *o.AmountSymbols
}

// GetAmountSymbolsOk returns a tuple with the AmountSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetAmountSymbolsOk() (*int32, bool) {
	if o == nil || o.AmountSymbols == nil {
		return nil, false
	}
	return o.AmountSymbols, true
}

// HasAmountSymbols returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasAmountSymbols() bool {
	if o != nil && o.AmountSymbols != nil {
		return true
	}

	return false
}

// SetAmountSymbols gets a reference to the given int32 and assigns it to the AmountSymbols field.
func (o *PasswordPolicyRequest) SetAmountSymbols(v int32) {
	o.AmountSymbols = &v
}

// GetLengthMin returns the LengthMin field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetLengthMin() int32 {
	if o == nil || o.LengthMin == nil {
		var ret int32
		return ret
	}
	return *o.LengthMin
}

// GetLengthMinOk returns a tuple with the LengthMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetLengthMinOk() (*int32, bool) {
	if o == nil || o.LengthMin == nil {
		return nil, false
	}
	return o.LengthMin, true
}

// HasLengthMin returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasLengthMin() bool {
	if o != nil && o.LengthMin != nil {
		return true
	}

	return false
}

// SetLengthMin gets a reference to the given int32 and assigns it to the LengthMin field.
func (o *PasswordPolicyRequest) SetLengthMin(v int32) {
	o.LengthMin = &v
}

// GetSymbolCharset returns the SymbolCharset field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetSymbolCharset() string {
	if o == nil || o.SymbolCharset == nil {
		var ret string
		return ret
	}
	return *o.SymbolCharset
}

// GetSymbolCharsetOk returns a tuple with the SymbolCharset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetSymbolCharsetOk() (*string, bool) {
	if o == nil || o.SymbolCharset == nil {
		return nil, false
	}
	return o.SymbolCharset, true
}

// HasSymbolCharset returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasSymbolCharset() bool {
	if o != nil && o.SymbolCharset != nil {
		return true
	}

	return false
}

// SetSymbolCharset gets a reference to the given string and assigns it to the SymbolCharset field.
func (o *PasswordPolicyRequest) SetSymbolCharset(v string) {
	o.SymbolCharset = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *PasswordPolicyRequest) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetCheckStaticRules returns the CheckStaticRules field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetCheckStaticRules() bool {
	if o == nil || o.CheckStaticRules == nil {
		var ret bool
		return ret
	}
	return *o.CheckStaticRules
}

// GetCheckStaticRulesOk returns a tuple with the CheckStaticRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetCheckStaticRulesOk() (*bool, bool) {
	if o == nil || o.CheckStaticRules == nil {
		return nil, false
	}
	return o.CheckStaticRules, true
}

// HasCheckStaticRules returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasCheckStaticRules() bool {
	if o != nil && o.CheckStaticRules != nil {
		return true
	}

	return false
}

// SetCheckStaticRules gets a reference to the given bool and assigns it to the CheckStaticRules field.
func (o *PasswordPolicyRequest) SetCheckStaticRules(v bool) {
	o.CheckStaticRules = &v
}

// GetCheckHaveIBeenPwned returns the CheckHaveIBeenPwned field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetCheckHaveIBeenPwned() bool {
	if o == nil || o.CheckHaveIBeenPwned == nil {
		var ret bool
		return ret
	}
	return *o.CheckHaveIBeenPwned
}

// GetCheckHaveIBeenPwnedOk returns a tuple with the CheckHaveIBeenPwned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetCheckHaveIBeenPwnedOk() (*bool, bool) {
	if o == nil || o.CheckHaveIBeenPwned == nil {
		return nil, false
	}
	return o.CheckHaveIBeenPwned, true
}

// HasCheckHaveIBeenPwned returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasCheckHaveIBeenPwned() bool {
	if o != nil && o.CheckHaveIBeenPwned != nil {
		return true
	}

	return false
}

// SetCheckHaveIBeenPwned gets a reference to the given bool and assigns it to the CheckHaveIBeenPwned field.
func (o *PasswordPolicyRequest) SetCheckHaveIBeenPwned(v bool) {
	o.CheckHaveIBeenPwned = &v
}

// GetCheckZxcvbn returns the CheckZxcvbn field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetCheckZxcvbn() bool {
	if o == nil || o.CheckZxcvbn == nil {
		var ret bool
		return ret
	}
	return *o.CheckZxcvbn
}

// GetCheckZxcvbnOk returns a tuple with the CheckZxcvbn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetCheckZxcvbnOk() (*bool, bool) {
	if o == nil || o.CheckZxcvbn == nil {
		return nil, false
	}
	return o.CheckZxcvbn, true
}

// HasCheckZxcvbn returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasCheckZxcvbn() bool {
	if o != nil && o.CheckZxcvbn != nil {
		return true
	}

	return false
}

// SetCheckZxcvbn gets a reference to the given bool and assigns it to the CheckZxcvbn field.
func (o *PasswordPolicyRequest) SetCheckZxcvbn(v bool) {
	o.CheckZxcvbn = &v
}

// GetHibpAllowedCount returns the HibpAllowedCount field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetHibpAllowedCount() int32 {
	if o == nil || o.HibpAllowedCount == nil {
		var ret int32
		return ret
	}
	return *o.HibpAllowedCount
}

// GetHibpAllowedCountOk returns a tuple with the HibpAllowedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetHibpAllowedCountOk() (*int32, bool) {
	if o == nil || o.HibpAllowedCount == nil {
		return nil, false
	}
	return o.HibpAllowedCount, true
}

// HasHibpAllowedCount returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasHibpAllowedCount() bool {
	if o != nil && o.HibpAllowedCount != nil {
		return true
	}

	return false
}

// SetHibpAllowedCount gets a reference to the given int32 and assigns it to the HibpAllowedCount field.
func (o *PasswordPolicyRequest) SetHibpAllowedCount(v int32) {
	o.HibpAllowedCount = &v
}

// GetZxcvbnScoreThreshold returns the ZxcvbnScoreThreshold field value if set, zero value otherwise.
func (o *PasswordPolicyRequest) GetZxcvbnScoreThreshold() int32 {
	if o == nil || o.ZxcvbnScoreThreshold == nil {
		var ret int32
		return ret
	}
	return *o.ZxcvbnScoreThreshold
}

// GetZxcvbnScoreThresholdOk returns a tuple with the ZxcvbnScoreThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetZxcvbnScoreThresholdOk() (*int32, bool) {
	if o == nil || o.ZxcvbnScoreThreshold == nil {
		return nil, false
	}
	return o.ZxcvbnScoreThreshold, true
}

// HasZxcvbnScoreThreshold returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasZxcvbnScoreThreshold() bool {
	if o != nil && o.ZxcvbnScoreThreshold != nil {
		return true
	}

	return false
}

// SetZxcvbnScoreThreshold gets a reference to the given int32 and assigns it to the ZxcvbnScoreThreshold field.
func (o *PasswordPolicyRequest) SetZxcvbnScoreThreshold(v int32) {
	o.ZxcvbnScoreThreshold = &v
}

func (o PasswordPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.PasswordField != nil {
		toSerialize["password_field"] = o.PasswordField
	}
	if o.AmountDigits != nil {
		toSerialize["amount_digits"] = o.AmountDigits
	}
	if o.AmountUppercase != nil {
		toSerialize["amount_uppercase"] = o.AmountUppercase
	}
	if o.AmountLowercase != nil {
		toSerialize["amount_lowercase"] = o.AmountLowercase
	}
	if o.AmountSymbols != nil {
		toSerialize["amount_symbols"] = o.AmountSymbols
	}
	if o.LengthMin != nil {
		toSerialize["length_min"] = o.LengthMin
	}
	if o.SymbolCharset != nil {
		toSerialize["symbol_charset"] = o.SymbolCharset
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.CheckStaticRules != nil {
		toSerialize["check_static_rules"] = o.CheckStaticRules
	}
	if o.CheckHaveIBeenPwned != nil {
		toSerialize["check_have_i_been_pwned"] = o.CheckHaveIBeenPwned
	}
	if o.CheckZxcvbn != nil {
		toSerialize["check_zxcvbn"] = o.CheckZxcvbn
	}
	if o.HibpAllowedCount != nil {
		toSerialize["hibp_allowed_count"] = o.HibpAllowedCount
	}
	if o.ZxcvbnScoreThreshold != nil {
		toSerialize["zxcvbn_score_threshold"] = o.ZxcvbnScoreThreshold
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordPolicyRequest struct {
	value *PasswordPolicyRequest
	isSet bool
}

func (v NullablePasswordPolicyRequest) Get() *PasswordPolicyRequest {
	return v.value
}

func (v *NullablePasswordPolicyRequest) Set(val *PasswordPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyRequest(val *PasswordPolicyRequest) *NullablePasswordPolicyRequest {
	return &NullablePasswordPolicyRequest{value: val, isSet: true}
}

func (v NullablePasswordPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
