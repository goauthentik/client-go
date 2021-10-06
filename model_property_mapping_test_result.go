/*
authentik

Making authentication simple.

API version: 2021.9.7
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PropertyMappingTestResult Result of a Property-mapping test
type PropertyMappingTestResult struct {
	Result     string `json:"result"`
	Successful bool   `json:"successful"`
}

// NewPropertyMappingTestResult instantiates a new PropertyMappingTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyMappingTestResult(result string, successful bool) *PropertyMappingTestResult {
	this := PropertyMappingTestResult{}
	this.Result = result
	this.Successful = successful
	return &this
}

// NewPropertyMappingTestResultWithDefaults instantiates a new PropertyMappingTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyMappingTestResultWithDefaults() *PropertyMappingTestResult {
	this := PropertyMappingTestResult{}
	return &this
}

// GetResult returns the Result field value
func (o *PropertyMappingTestResult) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *PropertyMappingTestResult) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *PropertyMappingTestResult) SetResult(v string) {
	o.Result = v
}

// GetSuccessful returns the Successful field value
func (o *PropertyMappingTestResult) GetSuccessful() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Successful
}

// GetSuccessfulOk returns a tuple with the Successful field value
// and a boolean to check if the value has been set.
func (o *PropertyMappingTestResult) GetSuccessfulOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Successful, true
}

// SetSuccessful sets field value
func (o *PropertyMappingTestResult) SetSuccessful(v bool) {
	o.Successful = v
}

func (o PropertyMappingTestResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["result"] = o.Result
	}
	if true {
		toSerialize["successful"] = o.Successful
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyMappingTestResult struct {
	value *PropertyMappingTestResult
	isSet bool
}

func (v NullablePropertyMappingTestResult) Get() *PropertyMappingTestResult {
	return v.value
}

func (v *NullablePropertyMappingTestResult) Set(val *PropertyMappingTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyMappingTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyMappingTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyMappingTestResult(val *PropertyMappingTestResult) *NullablePropertyMappingTestResult {
	return &NullablePropertyMappingTestResult{value: val, isSet: true}
}

func (v NullablePropertyMappingTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyMappingTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
