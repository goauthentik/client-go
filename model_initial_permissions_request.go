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

// InitialPermissionsRequest InitialPermissions serializer
type InitialPermissionsRequest struct {
	Name        string                     `json:"name"`
	Mode        InitialPermissionsModeEnum `json:"mode"`
	Role        string                     `json:"role"`
	Permissions []int32                    `json:"permissions,omitempty"`
}

// NewInitialPermissionsRequest instantiates a new InitialPermissionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitialPermissionsRequest(name string, mode InitialPermissionsModeEnum, role string) *InitialPermissionsRequest {
	this := InitialPermissionsRequest{}
	this.Name = name
	this.Mode = mode
	this.Role = role
	return &this
}

// NewInitialPermissionsRequestWithDefaults instantiates a new InitialPermissionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitialPermissionsRequestWithDefaults() *InitialPermissionsRequest {
	this := InitialPermissionsRequest{}
	return &this
}

// GetName returns the Name field value
func (o *InitialPermissionsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InitialPermissionsRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InitialPermissionsRequest) SetName(v string) {
	o.Name = v
}

// GetMode returns the Mode field value
func (o *InitialPermissionsRequest) GetMode() InitialPermissionsModeEnum {
	if o == nil {
		var ret InitialPermissionsModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *InitialPermissionsRequest) GetModeOk() (*InitialPermissionsModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *InitialPermissionsRequest) SetMode(v InitialPermissionsModeEnum) {
	o.Mode = v
}

// GetRole returns the Role field value
func (o *InitialPermissionsRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *InitialPermissionsRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *InitialPermissionsRequest) SetRole(v string) {
	o.Role = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *InitialPermissionsRequest) GetPermissions() []int32 {
	if o == nil || o.Permissions == nil {
		var ret []int32
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialPermissionsRequest) GetPermissionsOk() ([]int32, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *InitialPermissionsRequest) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []int32 and assigns it to the Permissions field.
func (o *InitialPermissionsRequest) SetPermissions(v []int32) {
	o.Permissions = v
}

func (o InitialPermissionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["role"] = o.Role
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableInitialPermissionsRequest struct {
	value *InitialPermissionsRequest
	isSet bool
}

func (v NullableInitialPermissionsRequest) Get() *InitialPermissionsRequest {
	return v.value
}

func (v *NullableInitialPermissionsRequest) Set(val *InitialPermissionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInitialPermissionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInitialPermissionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitialPermissionsRequest(val *InitialPermissionsRequest) *NullableInitialPermissionsRequest {
	return &NullableInitialPermissionsRequest{value: val, isSet: true}
}

func (v NullableInitialPermissionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitialPermissionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
