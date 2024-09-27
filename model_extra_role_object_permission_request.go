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

// ExtraRoleObjectPermissionRequest User permission with additional object-related data
type ExtraRoleObjectPermissionRequest struct {
	ObjectPk string `json:"object_pk"`
}

// NewExtraRoleObjectPermissionRequest instantiates a new ExtraRoleObjectPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtraRoleObjectPermissionRequest(objectPk string) *ExtraRoleObjectPermissionRequest {
	this := ExtraRoleObjectPermissionRequest{}
	this.ObjectPk = objectPk
	return &this
}

// NewExtraRoleObjectPermissionRequestWithDefaults instantiates a new ExtraRoleObjectPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtraRoleObjectPermissionRequestWithDefaults() *ExtraRoleObjectPermissionRequest {
	this := ExtraRoleObjectPermissionRequest{}
	return &this
}

// GetObjectPk returns the ObjectPk field value
func (o *ExtraRoleObjectPermissionRequest) GetObjectPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectPk
}

// GetObjectPkOk returns a tuple with the ObjectPk field value
// and a boolean to check if the value has been set.
func (o *ExtraRoleObjectPermissionRequest) GetObjectPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectPk, true
}

// SetObjectPk sets field value
func (o *ExtraRoleObjectPermissionRequest) SetObjectPk(v string) {
	o.ObjectPk = v
}

func (o ExtraRoleObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["object_pk"] = o.ObjectPk
	}
	return json.Marshal(toSerialize)
}

type NullableExtraRoleObjectPermissionRequest struct {
	value *ExtraRoleObjectPermissionRequest
	isSet bool
}

func (v NullableExtraRoleObjectPermissionRequest) Get() *ExtraRoleObjectPermissionRequest {
	return v.value
}

func (v *NullableExtraRoleObjectPermissionRequest) Set(val *ExtraRoleObjectPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExtraRoleObjectPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExtraRoleObjectPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtraRoleObjectPermissionRequest(val *ExtraRoleObjectPermissionRequest) *NullableExtraRoleObjectPermissionRequest {
	return &NullableExtraRoleObjectPermissionRequest{value: val, isSet: true}
}

func (v NullableExtraRoleObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtraRoleObjectPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
