/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReputationPolicyRequest Reputation Policy Serializer
type ReputationPolicyRequest struct {
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool  `json:"execution_logging,omitempty"`
	CheckIp          *bool  `json:"check_ip,omitempty"`
	CheckUsername    *bool  `json:"check_username,omitempty"`
	Threshold        *int32 `json:"threshold,omitempty"`
}

// NewReputationPolicyRequest instantiates a new ReputationPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReputationPolicyRequest(name string) *ReputationPolicyRequest {
	this := ReputationPolicyRequest{}
	this.Name = name
	return &this
}

// NewReputationPolicyRequestWithDefaults instantiates a new ReputationPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReputationPolicyRequestWithDefaults() *ReputationPolicyRequest {
	this := ReputationPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ReputationPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReputationPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReputationPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *ReputationPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReputationPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *ReputationPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *ReputationPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetCheckIp returns the CheckIp field value if set, zero value otherwise.
func (o *ReputationPolicyRequest) GetCheckIp() bool {
	if o == nil || o.CheckIp == nil {
		var ret bool
		return ret
	}
	return *o.CheckIp
}

// GetCheckIpOk returns a tuple with the CheckIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReputationPolicyRequest) GetCheckIpOk() (*bool, bool) {
	if o == nil || o.CheckIp == nil {
		return nil, false
	}
	return o.CheckIp, true
}

// HasCheckIp returns a boolean if a field has been set.
func (o *ReputationPolicyRequest) HasCheckIp() bool {
	if o != nil && o.CheckIp != nil {
		return true
	}

	return false
}

// SetCheckIp gets a reference to the given bool and assigns it to the CheckIp field.
func (o *ReputationPolicyRequest) SetCheckIp(v bool) {
	o.CheckIp = &v
}

// GetCheckUsername returns the CheckUsername field value if set, zero value otherwise.
func (o *ReputationPolicyRequest) GetCheckUsername() bool {
	if o == nil || o.CheckUsername == nil {
		var ret bool
		return ret
	}
	return *o.CheckUsername
}

// GetCheckUsernameOk returns a tuple with the CheckUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReputationPolicyRequest) GetCheckUsernameOk() (*bool, bool) {
	if o == nil || o.CheckUsername == nil {
		return nil, false
	}
	return o.CheckUsername, true
}

// HasCheckUsername returns a boolean if a field has been set.
func (o *ReputationPolicyRequest) HasCheckUsername() bool {
	if o != nil && o.CheckUsername != nil {
		return true
	}

	return false
}

// SetCheckUsername gets a reference to the given bool and assigns it to the CheckUsername field.
func (o *ReputationPolicyRequest) SetCheckUsername(v bool) {
	o.CheckUsername = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *ReputationPolicyRequest) GetThreshold() int32 {
	if o == nil || o.Threshold == nil {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReputationPolicyRequest) GetThresholdOk() (*int32, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *ReputationPolicyRequest) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *ReputationPolicyRequest) SetThreshold(v int32) {
	o.Threshold = &v
}

func (o ReputationPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.CheckIp != nil {
		toSerialize["check_ip"] = o.CheckIp
	}
	if o.CheckUsername != nil {
		toSerialize["check_username"] = o.CheckUsername
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	return json.Marshal(toSerialize)
}

type NullableReputationPolicyRequest struct {
	value *ReputationPolicyRequest
	isSet bool
}

func (v NullableReputationPolicyRequest) Get() *ReputationPolicyRequest {
	return v.value
}

func (v *NullableReputationPolicyRequest) Set(val *ReputationPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReputationPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReputationPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReputationPolicyRequest(val *ReputationPolicyRequest) *NullableReputationPolicyRequest {
	return &NullableReputationPolicyRequest{value: val, isSet: true}
}

func (v NullableReputationPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReputationPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
