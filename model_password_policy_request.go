/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PasswordPolicyRequest Password Policy Serializer
type PasswordPolicyRequest struct {
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Field key to check, field keys defined in Prompt stages are available.
	PasswordField *string `json:"password_field,omitempty"`
	AmountUppercase *int32 `json:"amount_uppercase,omitempty"`
	AmountLowercase *int32 `json:"amount_lowercase,omitempty"`
	AmountSymbols *int32 `json:"amount_symbols,omitempty"`
	LengthMin *int32 `json:"length_min,omitempty"`
	SymbolCharset *string `json:"symbol_charset,omitempty"`
	ErrorMessage string `json:"error_message"`
}

// NewPasswordPolicyRequest instantiates a new PasswordPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyRequest(errorMessage string) *PasswordPolicyRequest {
	this := PasswordPolicyRequest{}
	this.ErrorMessage = errorMessage
	return &this
}

// NewPasswordPolicyRequestWithDefaults instantiates a new PasswordPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyRequestWithDefaults() *PasswordPolicyRequest {
	this := PasswordPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PasswordPolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PasswordPolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PasswordPolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PasswordPolicyRequest) UnsetName() {
	o.Name.Unset()
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

// GetErrorMessage returns the ErrorMessage field value
func (o *PasswordPolicyRequest) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicyRequest) GetErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *PasswordPolicyRequest) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

func (o PasswordPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.PasswordField != nil {
		toSerialize["password_field"] = o.PasswordField
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
	if true {
		toSerialize["error_message"] = o.ErrorMessage
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


