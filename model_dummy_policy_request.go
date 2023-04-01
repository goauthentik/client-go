/*
authentik

Making authentication simple.

API version: 2023.3.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DummyPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DummyPolicyRequest{}

// DummyPolicyRequest Dummy Policy Serializer
type DummyPolicyRequest struct {
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool  `json:"execution_logging,omitempty"`
	Result           *bool  `json:"result,omitempty"`
	WaitMin          *int32 `json:"wait_min,omitempty"`
	WaitMax          *int32 `json:"wait_max,omitempty"`
}

// NewDummyPolicyRequest instantiates a new DummyPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDummyPolicyRequest(name string) *DummyPolicyRequest {
	this := DummyPolicyRequest{}
	this.Name = name
	return &this
}

// NewDummyPolicyRequestWithDefaults instantiates a new DummyPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDummyPolicyRequestWithDefaults() *DummyPolicyRequest {
	this := DummyPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DummyPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DummyPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DummyPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *DummyPolicyRequest) GetExecutionLogging() bool {
	if o == nil || IsNil(o.ExecutionLogging) {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicyRequest) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || IsNil(o.ExecutionLogging) {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *DummyPolicyRequest) HasExecutionLogging() bool {
	if o != nil && !IsNil(o.ExecutionLogging) {
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
	if o == nil || IsNil(o.Result) {
		var ret bool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicyRequest) GetResultOk() (*bool, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DummyPolicyRequest) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
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
	if o == nil || IsNil(o.WaitMin) {
		var ret int32
		return ret
	}
	return *o.WaitMin
}

// GetWaitMinOk returns a tuple with the WaitMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicyRequest) GetWaitMinOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitMin) {
		return nil, false
	}
	return o.WaitMin, true
}

// HasWaitMin returns a boolean if a field has been set.
func (o *DummyPolicyRequest) HasWaitMin() bool {
	if o != nil && !IsNil(o.WaitMin) {
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
	if o == nil || IsNil(o.WaitMax) {
		var ret int32
		return ret
	}
	return *o.WaitMax
}

// GetWaitMaxOk returns a tuple with the WaitMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicyRequest) GetWaitMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitMax) {
		return nil, false
	}
	return o.WaitMax, true
}

// HasWaitMax returns a boolean if a field has been set.
func (o *DummyPolicyRequest) HasWaitMax() bool {
	if o != nil && !IsNil(o.WaitMax) {
		return true
	}

	return false
}

// SetWaitMax gets a reference to the given int32 and assigns it to the WaitMax field.
func (o *DummyPolicyRequest) SetWaitMax(v int32) {
	o.WaitMax = &v
}

func (o DummyPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DummyPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.ExecutionLogging) {
		toSerialize["execution_logging"] = o.ExecutionLogging
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.WaitMin) {
		toSerialize["wait_min"] = o.WaitMin
	}
	if !IsNil(o.WaitMax) {
		toSerialize["wait_max"] = o.WaitMax
	}
	return toSerialize, nil
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
