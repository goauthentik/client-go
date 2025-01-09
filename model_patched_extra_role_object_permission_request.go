/*
authentik

Making authentication simple.

API version: 2024.12.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedExtraRoleObjectPermissionRequest User permission with additional object-related data
type PatchedExtraRoleObjectPermissionRequest struct {
	ObjectPk *string `json:"object_pk,omitempty"`
}

// NewPatchedExtraRoleObjectPermissionRequest instantiates a new PatchedExtraRoleObjectPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedExtraRoleObjectPermissionRequest() *PatchedExtraRoleObjectPermissionRequest {
	this := PatchedExtraRoleObjectPermissionRequest{}
	return &this
}

// NewPatchedExtraRoleObjectPermissionRequestWithDefaults instantiates a new PatchedExtraRoleObjectPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedExtraRoleObjectPermissionRequestWithDefaults() *PatchedExtraRoleObjectPermissionRequest {
	this := PatchedExtraRoleObjectPermissionRequest{}
	return &this
}

// GetObjectPk returns the ObjectPk field value if set, zero value otherwise.
func (o *PatchedExtraRoleObjectPermissionRequest) GetObjectPk() string {
	if o == nil || o.ObjectPk == nil {
		var ret string
		return ret
	}
	return *o.ObjectPk
}

// GetObjectPkOk returns a tuple with the ObjectPk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExtraRoleObjectPermissionRequest) GetObjectPkOk() (*string, bool) {
	if o == nil || o.ObjectPk == nil {
		return nil, false
	}
	return o.ObjectPk, true
}

// HasObjectPk returns a boolean if a field has been set.
func (o *PatchedExtraRoleObjectPermissionRequest) HasObjectPk() bool {
	if o != nil && o.ObjectPk != nil {
		return true
	}

	return false
}

// SetObjectPk gets a reference to the given string and assigns it to the ObjectPk field.
func (o *PatchedExtraRoleObjectPermissionRequest) SetObjectPk(v string) {
	o.ObjectPk = &v
}

func (o PatchedExtraRoleObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectPk != nil {
		toSerialize["object_pk"] = o.ObjectPk
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedExtraRoleObjectPermissionRequest struct {
	value *PatchedExtraRoleObjectPermissionRequest
	isSet bool
}

func (v NullablePatchedExtraRoleObjectPermissionRequest) Get() *PatchedExtraRoleObjectPermissionRequest {
	return v.value
}

func (v *NullablePatchedExtraRoleObjectPermissionRequest) Set(val *PatchedExtraRoleObjectPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedExtraRoleObjectPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedExtraRoleObjectPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedExtraRoleObjectPermissionRequest(val *PatchedExtraRoleObjectPermissionRequest) *NullablePatchedExtraRoleObjectPermissionRequest {
	return &NullablePatchedExtraRoleObjectPermissionRequest{value: val, isSet: true}
}

func (v NullablePatchedExtraRoleObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedExtraRoleObjectPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
