/*
authentik

Making authentication simple.

API version: 2022.7.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// HaveIBeenPwendPolicyRequest Have I Been Pwned Policy Serializer
type HaveIBeenPwendPolicyRequest struct {
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Field key to check, field keys defined in Prompt stages are available.
	PasswordField *string `json:"password_field,omitempty"`
	AllowedCount  *int32  `json:"allowed_count,omitempty"`
}

// NewHaveIBeenPwendPolicyRequest instantiates a new HaveIBeenPwendPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHaveIBeenPwendPolicyRequest() *HaveIBeenPwendPolicyRequest {
	this := HaveIBeenPwendPolicyRequest{}
	return &this
}

// NewHaveIBeenPwendPolicyRequestWithDefaults instantiates a new HaveIBeenPwendPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHaveIBeenPwendPolicyRequestWithDefaults() *HaveIBeenPwendPolicyRequest {
	this := HaveIBeenPwendPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HaveIBeenPwendPolicyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HaveIBeenPwendPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *HaveIBeenPwendPolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *HaveIBeenPwendPolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *HaveIBeenPwendPolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *HaveIBeenPwendPolicyRequest) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *HaveIBeenPwendPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaveIBeenPwendPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *HaveIBeenPwendPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *HaveIBeenPwendPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetPasswordField returns the PasswordField field value if set, zero value otherwise.
func (o *HaveIBeenPwendPolicyRequest) GetPasswordField() string {
	if o == nil || o.PasswordField == nil {
		var ret string
		return ret
	}
	return *o.PasswordField
}

// GetPasswordFieldOk returns a tuple with the PasswordField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaveIBeenPwendPolicyRequest) GetPasswordFieldOk() (*string, bool) {
	if o == nil || o.PasswordField == nil {
		return nil, false
	}
	return o.PasswordField, true
}

// HasPasswordField returns a boolean if a field has been set.
func (o *HaveIBeenPwendPolicyRequest) HasPasswordField() bool {
	if o != nil && o.PasswordField != nil {
		return true
	}

	return false
}

// SetPasswordField gets a reference to the given string and assigns it to the PasswordField field.
func (o *HaveIBeenPwendPolicyRequest) SetPasswordField(v string) {
	o.PasswordField = &v
}

// GetAllowedCount returns the AllowedCount field value if set, zero value otherwise.
func (o *HaveIBeenPwendPolicyRequest) GetAllowedCount() int32 {
	if o == nil || o.AllowedCount == nil {
		var ret int32
		return ret
	}
	return *o.AllowedCount
}

// GetAllowedCountOk returns a tuple with the AllowedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HaveIBeenPwendPolicyRequest) GetAllowedCountOk() (*int32, bool) {
	if o == nil || o.AllowedCount == nil {
		return nil, false
	}
	return o.AllowedCount, true
}

// HasAllowedCount returns a boolean if a field has been set.
func (o *HaveIBeenPwendPolicyRequest) HasAllowedCount() bool {
	if o != nil && o.AllowedCount != nil {
		return true
	}

	return false
}

// SetAllowedCount gets a reference to the given int32 and assigns it to the AllowedCount field.
func (o *HaveIBeenPwendPolicyRequest) SetAllowedCount(v int32) {
	o.AllowedCount = &v
}

func (o HaveIBeenPwendPolicyRequest) MarshalJSON() ([]byte, error) {
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
	if o.AllowedCount != nil {
		toSerialize["allowed_count"] = o.AllowedCount
	}
	return json.Marshal(toSerialize)
}

type NullableHaveIBeenPwendPolicyRequest struct {
	value *HaveIBeenPwendPolicyRequest
	isSet bool
}

func (v NullableHaveIBeenPwendPolicyRequest) Get() *HaveIBeenPwendPolicyRequest {
	return v.value
}

func (v *NullableHaveIBeenPwendPolicyRequest) Set(val *HaveIBeenPwendPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHaveIBeenPwendPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHaveIBeenPwendPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHaveIBeenPwendPolicyRequest(val *HaveIBeenPwendPolicyRequest) *NullableHaveIBeenPwendPolicyRequest {
	return &NullableHaveIBeenPwendPolicyRequest{value: val, isSet: true}
}

func (v NullableHaveIBeenPwendPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHaveIBeenPwendPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
