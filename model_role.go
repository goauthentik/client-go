/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Role Role serializer
type Role struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
}

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole(pk string, name string) *Role {
	this := Role{}
	this.Pk = pk
	this.Name = name
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetPk returns the Pk field value
func (o *Role) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Role) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Role) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *Role) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Role) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Role) SetName(v string) {
	o.Name = v
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
