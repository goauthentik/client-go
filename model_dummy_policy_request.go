/*
authentik

Making authentication simple.

API version: 2021.8.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DummyPolicyRequest Dummy Policy Serializer
type DummyPolicyRequest struct {
	Name NullableString `json:"name,omitempty"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	Result *bool `json:"result,omitempty"`
	WaitMin *int32 `json:"wait_min,omitempty"`
	WaitMax *int32 `json:"wait_max,omitempty"`
}

// NewDummyPolicyRequest instantiates a new DummyPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDummyPolicyRequest() *DummyPolicyRequest {
	this := DummyPolicyRequest{}
	return &this
}

// NewDummyPolicyRequestWithDefaults instantiates a new DummyPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDummyPolicyRequestWithDefaults() *DummyPolicyRequest {
	this := DummyPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DummyPolicyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DummyPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DummyPolicyRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DummyPolicyRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DummyPolicyRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DummyPolicyRequest) UnsetName() {
	o.Name.Unset()
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *DummyPolicyRequest) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *DummyPolicyRequest) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *DummyPolicyRequest) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DummyPolicyRequest) GetResult() bool {
	if o == nil || o.Result == nil {
		var ret bool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicyRequest) GetResultOk() (*bool, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DummyPolicyRequest) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given bool and assigns it to the Result field.
func (o *DummyPolicyRequest) SetResult(v bool) {
	o.Result = &v
}

// GetWaitMin returns the WaitMin field value if set, zero value otherwise.
func (o *DummyPolicyRequest) GetWaitMin() int32 {
	if o == nil || o.WaitMin == nil {
		var ret int32
		return ret
	}
	return *o.WaitMin
}

// GetWaitMinOk returns a tuple with the WaitMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicyRequest) GetWaitMinOk() (*int32, bool) {
	if o == nil || o.WaitMin == nil {
		return nil, false
	}
	return o.WaitMin, true
}

// HasWaitMin returns a boolean if a field has been set.
func (o *DummyPolicyRequest) HasWaitMin() bool {
	if o != nil && o.WaitMin != nil {
		return true
	}

	return false
}

// SetWaitMin gets a reference to the given int32 and assigns it to the WaitMin field.
func (o *DummyPolicyRequest) SetWaitMin(v int32) {
	o.WaitMin = &v
}

// GetWaitMax returns the WaitMax field value if set, zero value otherwise.
func (o *DummyPolicyRequest) GetWaitMax() int32 {
	if o == nil || o.WaitMax == nil {
		var ret int32
		return ret
	}
	return *o.WaitMax
}

// GetWaitMaxOk returns a tuple with the WaitMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicyRequest) GetWaitMaxOk() (*int32, bool) {
	if o == nil || o.WaitMax == nil {
		return nil, false
	}
	return o.WaitMax, true
}

// HasWaitMax returns a boolean if a field has been set.
func (o *DummyPolicyRequest) HasWaitMax() bool {
	if o != nil && o.WaitMax != nil {
		return true
	}

	return false
}

// SetWaitMax gets a reference to the given int32 and assigns it to the WaitMax field.
func (o *DummyPolicyRequest) SetWaitMax(v int32) {
	o.WaitMax = &v
}

func (o DummyPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.WaitMin != nil {
		toSerialize["wait_min"] = o.WaitMin
	}
	if o.WaitMax != nil {
		toSerialize["wait_max"] = o.WaitMax
	}
	return json.Marshal(toSerialize)
}

type NullableDummyPolicyRequest struct {
	value *DummyPolicyRequest
	isSet bool
}

func (v NullableDummyPolicyRequest) Get() *DummyPolicyRequest {
	return v.value
}

func (v *NullableDummyPolicyRequest) Set(val *DummyPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDummyPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDummyPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDummyPolicyRequest(val *DummyPolicyRequest) *NullableDummyPolicyRequest {
	return &NullableDummyPolicyRequest{value: val, isSet: true}
}

func (v NullableDummyPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDummyPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


