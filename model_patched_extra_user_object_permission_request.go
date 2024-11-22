/*
authentik

Making authentication simple.

API version: 2024.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedExtraUserObjectPermissionRequest User permission with additional object-related data
type PatchedExtraUserObjectPermissionRequest struct {
	ObjectPk *string `json:"object_pk,omitempty"`
}

// NewPatchedExtraUserObjectPermissionRequest instantiates a new PatchedExtraUserObjectPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedExtraUserObjectPermissionRequest() *PatchedExtraUserObjectPermissionRequest {
	this := PatchedExtraUserObjectPermissionRequest{}
	return &this
}

// NewPatchedExtraUserObjectPermissionRequestWithDefaults instantiates a new PatchedExtraUserObjectPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedExtraUserObjectPermissionRequestWithDefaults() *PatchedExtraUserObjectPermissionRequest {
	this := PatchedExtraUserObjectPermissionRequest{}
	return &this
}

// GetObjectPk returns the ObjectPk field value if set, zero value otherwise.
func (o *PatchedExtraUserObjectPermissionRequest) GetObjectPk() string {
	if o == nil || o.ObjectPk == nil {
		var ret string
		return ret
	}
	return *o.ObjectPk
}

// GetObjectPkOk returns a tuple with the ObjectPk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExtraUserObjectPermissionRequest) GetObjectPkOk() (*string, bool) {
	if o == nil || o.ObjectPk == nil {
		return nil, false
	}
	return o.ObjectPk, true
}

// HasObjectPk returns a boolean if a field has been set.
func (o *PatchedExtraUserObjectPermissionRequest) HasObjectPk() bool {
	if o != nil && o.ObjectPk != nil {
		return true
	}

	return false
}

// SetObjectPk gets a reference to the given string and assigns it to the ObjectPk field.
func (o *PatchedExtraUserObjectPermissionRequest) SetObjectPk(v string) {
	o.ObjectPk = &v
}

func (o PatchedExtraUserObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectPk != nil {
		toSerialize["object_pk"] = o.ObjectPk
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedExtraUserObjectPermissionRequest struct {
	value *PatchedExtraUserObjectPermissionRequest
	isSet bool
}

func (v NullablePatchedExtraUserObjectPermissionRequest) Get() *PatchedExtraUserObjectPermissionRequest {
	return v.value
}

func (v *NullablePatchedExtraUserObjectPermissionRequest) Set(val *PatchedExtraUserObjectPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedExtraUserObjectPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedExtraUserObjectPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedExtraUserObjectPermissionRequest(val *PatchedExtraUserObjectPermissionRequest) *NullablePatchedExtraUserObjectPermissionRequest {
	return &NullablePatchedExtraUserObjectPermissionRequest{value: val, isSet: true}
}

func (v NullablePatchedExtraUserObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedExtraUserObjectPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
