/*
authentik

Making authentication simple.

API version: 2022.11.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedHaveIBeenPwendPolicyRequest Have I Been Pwned Policy Serializer
type PatchedHaveIBeenPwendPolicyRequest struct {
	Name *string `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Field key to check, field keys defined in Prompt stages are available.
	PasswordField *string `json:"password_field,omitempty"`
	AllowedCount  *int32  `json:"allowed_count,omitempty"`
}

// NewPatchedHaveIBeenPwendPolicyRequest instantiates a new PatchedHaveIBeenPwendPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedHaveIBeenPwendPolicyRequest() *PatchedHaveIBeenPwendPolicyRequest {
	this := PatchedHaveIBeenPwendPolicyRequest{}
	return &this
}

// NewPatchedHaveIBeenPwendPolicyRequestWithDefaults instantiates a new PatchedHaveIBeenPwendPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedHaveIBeenPwendPolicyRequestWithDefaults() *PatchedHaveIBeenPwendPolicyRequest {
	this := PatchedHaveIBeenPwendPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedHaveIBeenPwendPolicyRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedHaveIBeenPwendPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedHaveIBeenPwendPolicyRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedHaveIBeenPwendPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PatchedHaveIBeenPwendPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedHaveIBeenPwendPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PatchedHaveIBeenPwendPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PatchedHaveIBeenPwendPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetPasswordField returns the PasswordField field value if set, zero value otherwise.
func (o *PatchedHaveIBeenPwendPolicyRequest) GetPasswordField() string {
	if o == nil || o.PasswordField == nil {
		var ret string
		return ret
	}
	return *o.PasswordField
}

// GetPasswordFieldOk returns a tuple with the PasswordField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedHaveIBeenPwendPolicyRequest) GetPasswordFieldOk() (*string, bool) {
	if o == nil || o.PasswordField == nil {
		return nil, false
	}
	return o.PasswordField, true
}

// HasPasswordField returns a boolean if a field has been set.
func (o *PatchedHaveIBeenPwendPolicyRequest) HasPasswordField() bool {
	if o != nil && o.PasswordField != nil {
		return true
	}

	return false
}

// SetPasswordField gets a reference to the given string and assigns it to the PasswordField field.
func (o *PatchedHaveIBeenPwendPolicyRequest) SetPasswordField(v string) {
	o.PasswordField = &v
}

// GetAllowedCount returns the AllowedCount field value if set, zero value otherwise.
func (o *PatchedHaveIBeenPwendPolicyRequest) GetAllowedCount() int32 {
	if o == nil || o.AllowedCount == nil {
		var ret int32
		return ret
	}
	return *o.AllowedCount
}

// GetAllowedCountOk returns a tuple with the AllowedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedHaveIBeenPwendPolicyRequest) GetAllowedCountOk() (*int32, bool) {
	if o == nil || o.AllowedCount == nil {
		return nil, false
	}
	return o.AllowedCount, true
}

// HasAllowedCount returns a boolean if a field has been set.
func (o *PatchedHaveIBeenPwendPolicyRequest) HasAllowedCount() bool {
	if o != nil && o.AllowedCount != nil {
		return true
	}

	return false
}

// SetAllowedCount gets a reference to the given int32 and assigns it to the AllowedCount field.
func (o *PatchedHaveIBeenPwendPolicyRequest) SetAllowedCount(v int32) {
	o.AllowedCount = &v
}

func (o PatchedHaveIBeenPwendPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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

type NullablePatchedHaveIBeenPwendPolicyRequest struct {
	value *PatchedHaveIBeenPwendPolicyRequest
	isSet bool
}

func (v NullablePatchedHaveIBeenPwendPolicyRequest) Get() *PatchedHaveIBeenPwendPolicyRequest {
	return v.value
}

func (v *NullablePatchedHaveIBeenPwendPolicyRequest) Set(val *PatchedHaveIBeenPwendPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedHaveIBeenPwendPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedHaveIBeenPwendPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedHaveIBeenPwendPolicyRequest(val *PatchedHaveIBeenPwendPolicyRequest) *NullablePatchedHaveIBeenPwendPolicyRequest {
	return &NullablePatchedHaveIBeenPwendPolicyRequest{value: val, isSet: true}
}

func (v NullablePatchedHaveIBeenPwendPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedHaveIBeenPwendPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
