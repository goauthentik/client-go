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

// RoleRequest Role serializer
type RoleRequest struct {
	Name string `json:"name"`
}

// NewRoleRequest instantiates a new RoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleRequest(name string) *RoleRequest {
	this := RoleRequest{}
	this.Name = name
	return &this
}

// NewRoleRequestWithDefaults instantiates a new RoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleRequestWithDefaults() *RoleRequest {
	this := RoleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *RoleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleRequest) SetName(v string) {
	o.Name = v
}

func (o RoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableRoleRequest struct {
	value *RoleRequest
	isSet bool
}

func (v NullableRoleRequest) Get() *RoleRequest {
	return v.value
}

func (v *NullableRoleRequest) Set(val *RoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleRequest(val *RoleRequest) *NullableRoleRequest {
	return &NullableRoleRequest{value: val, isSet: true}
}

func (v NullableRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
