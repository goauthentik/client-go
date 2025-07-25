/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserPath struct for UserPath
type UserPath struct {
	Paths []string `json:"paths"`
}

// NewUserPath instantiates a new UserPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPath(paths []string) *UserPath {
	this := UserPath{}
	this.Paths = paths
	return &this
}

// NewUserPathWithDefaults instantiates a new UserPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPathWithDefaults() *UserPath {
	this := UserPath{}
	return &this
}

// GetPaths returns the Paths field value
func (o *UserPath) GetPaths() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value
// and a boolean to check if the value has been set.
func (o *UserPath) GetPathsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Paths, true
}

// SetPaths sets field value
func (o *UserPath) SetPaths(v []string) {
	o.Paths = v
}

func (o UserPath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["paths"] = o.Paths
	}
	return json.Marshal(toSerialize)
}

type NullableUserPath struct {
	value *UserPath
	isSet bool
}

func (v NullableUserPath) Get() *UserPath {
	return v.value
}

func (v *NullableUserPath) Set(val *UserPath) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPath) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPath(val *UserPath) *NullableUserPath {
	return &NullableUserPath{value: val, isSet: true}
}

func (v NullableUserPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
