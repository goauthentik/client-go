/*
authentik

Making authentication simple.

API version: 2024.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PolicyRequest Policy Serializer
type PolicyRequest struct {
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
}

// NewPolicyRequest instantiates a new PolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRequest(name string) *PolicyRequest {
	this := PolicyRequest{}
	this.Name = name
	return &this
}

// NewPolicyRequestWithDefaults instantiates a new PolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRequestWithDefaults() *PolicyRequest {
	this := PolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PolicyRequest) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

func (o PolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyRequest struct {
	value *PolicyRequest
	isSet bool
}

func (v NullablePolicyRequest) Get() *PolicyRequest {
	return v.value
}

func (v *NullablePolicyRequest) Set(val *PolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRequest(val *PolicyRequest) *NullablePolicyRequest {
	return &NullablePolicyRequest{value: val, isSet: true}
}

func (v NullablePolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
