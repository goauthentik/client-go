/*
authentik

Making authentication simple.

API version: 2024.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Workers struct for Workers
type Workers struct {
	Count int32 `json:"count"`
}

// NewWorkers instantiates a new Workers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkers(count int32) *Workers {
	this := Workers{}
	this.Count = count
	return &this
}

// NewWorkersWithDefaults instantiates a new Workers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkersWithDefaults() *Workers {
	this := Workers{}
	return &this
}

// GetCount returns the Count field value
func (o *Workers) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *Workers) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *Workers) SetCount(v int32) {
	o.Count = v
}

func (o Workers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableWorkers struct {
	value *Workers
	isSet bool
}

func (v NullableWorkers) Get() *Workers {
	return v.value
}

func (v *NullableWorkers) Set(val *Workers) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkers) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkers(val *Workers) *NullableWorkers {
	return &NullableWorkers{value: val, isSet: true}
}

func (v NullableWorkers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
