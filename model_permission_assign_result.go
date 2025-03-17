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

// PermissionAssignResult Result from assigning permissions to a user/role
type PermissionAssignResult struct {
	Id string `json:"id"`
}

// NewPermissionAssignResult instantiates a new PermissionAssignResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionAssignResult(id string) *PermissionAssignResult {
	this := PermissionAssignResult{}
	this.Id = id
	return &this
}

// NewPermissionAssignResultWithDefaults instantiates a new PermissionAssignResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionAssignResultWithDefaults() *PermissionAssignResult {
	this := PermissionAssignResult{}
	return &this
}

// GetId returns the Id field value
func (o *PermissionAssignResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PermissionAssignResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PermissionAssignResult) SetId(v string) {
	o.Id = v
}

func (o PermissionAssignResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionAssignResult struct {
	value *PermissionAssignResult
	isSet bool
}

func (v NullablePermissionAssignResult) Get() *PermissionAssignResult {
	return v.value
}

func (v *NullablePermissionAssignResult) Set(val *PermissionAssignResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionAssignResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionAssignResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionAssignResult(val *PermissionAssignResult) *NullablePermissionAssignResult {
	return &NullablePermissionAssignResult{value: val, isSet: true}
}

func (v NullablePermissionAssignResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionAssignResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
