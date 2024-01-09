/*
authentik

Making authentication simple.

API version: 2023.10.6
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserAccountRequest Account adding/removing operations
type UserAccountRequest struct {
	Pk int32 `json:"pk"`
}

// NewUserAccountRequest instantiates a new UserAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccountRequest(pk int32) *UserAccountRequest {
	this := UserAccountRequest{}
	this.Pk = pk
	return &this
}

// NewUserAccountRequestWithDefaults instantiates a new UserAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccountRequestWithDefaults() *UserAccountRequest {
	this := UserAccountRequest{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserAccountRequest) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserAccountRequest) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserAccountRequest) SetPk(v int32) {
	o.Pk = v
}

func (o UserAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	return json.Marshal(toSerialize)
}

type NullableUserAccountRequest struct {
	value *UserAccountRequest
	isSet bool
}

func (v NullableUserAccountRequest) Get() *UserAccountRequest {
	return v.value
}

func (v *NullableUserAccountRequest) Set(val *UserAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccountRequest(val *UserAccountRequest) *NullableUserAccountRequest {
	return &NullableUserAccountRequest{value: val, isSet: true}
}

func (v NullableUserAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
