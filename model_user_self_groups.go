/*
authentik

Making authentication simple.

API version: 2025.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserSelfGroups struct for UserSelfGroups
type UserSelfGroups struct {
	Name string `json:"name"`
	Pk   string `json:"pk"`
}

// NewUserSelfGroups instantiates a new UserSelfGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSelfGroups(name string, pk string) *UserSelfGroups {
	this := UserSelfGroups{}
	this.Name = name
	this.Pk = pk
	return &this
}

// NewUserSelfGroupsWithDefaults instantiates a new UserSelfGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSelfGroupsWithDefaults() *UserSelfGroups {
	this := UserSelfGroups{}
	return &this
}

// GetName returns the Name field value
func (o *UserSelfGroups) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserSelfGroups) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserSelfGroups) SetName(v string) {
	o.Name = v
}

// GetPk returns the Pk field value
func (o *UserSelfGroups) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserSelfGroups) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserSelfGroups) SetPk(v string) {
	o.Pk = v
}

func (o UserSelfGroups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["pk"] = o.Pk
	}
	return json.Marshal(toSerialize)
}

type NullableUserSelfGroups struct {
	value *UserSelfGroups
	isSet bool
}

func (v NullableUserSelfGroups) Get() *UserSelfGroups {
	return v.value
}

func (v *NullableUserSelfGroups) Set(val *UserSelfGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSelfGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSelfGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSelfGroups(val *UserSelfGroups) *NullableUserSelfGroups {
	return &NullableUserSelfGroups{value: val, isSet: true}
}

func (v NullableUserSelfGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSelfGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
