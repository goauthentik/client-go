/*
authentik

Making authentication simple.

API version: 2022.5.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PasswordExpiryPolicyRequest Password Expiry Policy Serializer
type PasswordExpiryPolicyRequest struct {
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	Days             int32 `json:"days"`
	DenyOnly         *bool `json:"deny_only,omitempty"`
}

// NewPasswordExpiryPolicyRequest instantiates a new PasswordExpiryPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordExpiryPolicyRequest(days int32) *PasswordExpiryPolicyRequest {
	this := PasswordExpiryPolicyRequest{}
	this.Days = days
	return &this
}

// NewPasswordExpiryPolicyRequestWithDefaults instantiates a new PasswordExpiryPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordExpiryPolicyRequestWithDefaults() *PasswordExpiryPolicyRequest {
	this := PasswordExpiryPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordExpiryPolicyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordExpiryPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PasswordExpiryPolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PasswordExpiryPolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PasswordExpiryPolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PasswordExpiryPolicyRequest) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PasswordExpiryPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PasswordExpiryPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PasswordExpiryPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetDays returns the Days field value
func (o *PasswordExpiryPolicyRequest) GetDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicyRequest) GetDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Days, true
}

// SetDays sets field value
func (o *PasswordExpiryPolicyRequest) SetDays(v int32) {
	o.Days = v
}

// GetDenyOnly returns the DenyOnly field value if set, zero value otherwise.
func (o *PasswordExpiryPolicyRequest) GetDenyOnly() bool {
	if o == nil || o.DenyOnly == nil {
		var ret bool
		return ret
	}
	return *o.DenyOnly
}

// GetDenyOnlyOk returns a tuple with the DenyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordExpiryPolicyRequest) GetDenyOnlyOk() (*bool, bool) {
	if o == nil || o.DenyOnly == nil {
		return nil, false
	}
	return o.DenyOnly, true
}

// HasDenyOnly returns a boolean if a field has been set.
func (o *PasswordExpiryPolicyRequest) HasDenyOnly() bool {
	if o != nil && o.DenyOnly != nil {
		return true
	}

	return false
}

// SetDenyOnly gets a reference to the given bool and assigns it to the DenyOnly field.
func (o *PasswordExpiryPolicyRequest) SetDenyOnly(v bool) {
	o.DenyOnly = &v
}

func (o PasswordExpiryPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if true {
		toSerialize["days"] = o.Days
	}
	if o.DenyOnly != nil {
		toSerialize["deny_only"] = o.DenyOnly
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordExpiryPolicyRequest struct {
	value *PasswordExpiryPolicyRequest
	isSet bool
}

func (v NullablePasswordExpiryPolicyRequest) Get() *PasswordExpiryPolicyRequest {
	return v.value
}

func (v *NullablePasswordExpiryPolicyRequest) Set(val *PasswordExpiryPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordExpiryPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordExpiryPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordExpiryPolicyRequest(val *PasswordExpiryPolicyRequest) *NullablePasswordExpiryPolicyRequest {
	return &NullablePasswordExpiryPolicyRequest{value: val, isSet: true}
}

func (v NullablePasswordExpiryPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordExpiryPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
