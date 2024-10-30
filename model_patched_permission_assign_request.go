/*
authentik

Making authentication simple.

API version: 2024.10.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedPermissionAssignRequest Request to assign a new permission
type PatchedPermissionAssignRequest struct {
	Permissions []string   `json:"permissions,omitempty"`
	Model       *ModelEnum `json:"model,omitempty"`
	ObjectPk    *string    `json:"object_pk,omitempty"`
}

// NewPatchedPermissionAssignRequest instantiates a new PatchedPermissionAssignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPermissionAssignRequest() *PatchedPermissionAssignRequest {
	this := PatchedPermissionAssignRequest{}
	return &this
}

// NewPatchedPermissionAssignRequestWithDefaults instantiates a new PatchedPermissionAssignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPermissionAssignRequestWithDefaults() *PatchedPermissionAssignRequest {
	this := PatchedPermissionAssignRequest{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *PatchedPermissionAssignRequest) GetPermissions() []string {
	if o == nil || o.Permissions == nil {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPermissionAssignRequest) GetPermissionsOk() ([]string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PatchedPermissionAssignRequest) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *PatchedPermissionAssignRequest) SetPermissions(v []string) {
	o.Permissions = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *PatchedPermissionAssignRequest) GetModel() ModelEnum {
	if o == nil || o.Model == nil {
		var ret ModelEnum
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPermissionAssignRequest) GetModelOk() (*ModelEnum, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *PatchedPermissionAssignRequest) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given ModelEnum and assigns it to the Model field.
func (o *PatchedPermissionAssignRequest) SetModel(v ModelEnum) {
	o.Model = &v
}

// GetObjectPk returns the ObjectPk field value if set, zero value otherwise.
func (o *PatchedPermissionAssignRequest) GetObjectPk() string {
	if o == nil || o.ObjectPk == nil {
		var ret string
		return ret
	}
	return *o.ObjectPk
}

// GetObjectPkOk returns a tuple with the ObjectPk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPermissionAssignRequest) GetObjectPkOk() (*string, bool) {
	if o == nil || o.ObjectPk == nil {
		return nil, false
	}
	return o.ObjectPk, true
}

// HasObjectPk returns a boolean if a field has been set.
func (o *PatchedPermissionAssignRequest) HasObjectPk() bool {
	if o != nil && o.ObjectPk != nil {
		return true
	}

	return false
}

// SetObjectPk gets a reference to the given string and assigns it to the ObjectPk field.
func (o *PatchedPermissionAssignRequest) SetObjectPk(v string) {
	o.ObjectPk = &v
}

func (o PatchedPermissionAssignRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.ObjectPk != nil {
		toSerialize["object_pk"] = o.ObjectPk
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedPermissionAssignRequest struct {
	value *PatchedPermissionAssignRequest
	isSet bool
}

func (v NullablePatchedPermissionAssignRequest) Get() *PatchedPermissionAssignRequest {
	return v.value
}

func (v *NullablePatchedPermissionAssignRequest) Set(val *PatchedPermissionAssignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPermissionAssignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPermissionAssignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPermissionAssignRequest(val *PatchedPermissionAssignRequest) *NullablePatchedPermissionAssignRequest {
	return &NullablePatchedPermissionAssignRequest{value: val, isSet: true}
}

func (v NullablePatchedPermissionAssignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPermissionAssignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
