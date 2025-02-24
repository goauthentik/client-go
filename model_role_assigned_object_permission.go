/*
authentik

Making authentication simple.

API version: 2025.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RoleAssignedObjectPermission Roles assigned object permission serializer
type RoleAssignedObjectPermission struct {
	RolePk      string                 `json:"role_pk"`
	Name        string                 `json:"name"`
	Permissions []RoleObjectPermission `json:"permissions"`
}

// NewRoleAssignedObjectPermission instantiates a new RoleAssignedObjectPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignedObjectPermission(rolePk string, name string, permissions []RoleObjectPermission) *RoleAssignedObjectPermission {
	this := RoleAssignedObjectPermission{}
	this.RolePk = rolePk
	this.Name = name
	this.Permissions = permissions
	return &this
}

// NewRoleAssignedObjectPermissionWithDefaults instantiates a new RoleAssignedObjectPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignedObjectPermissionWithDefaults() *RoleAssignedObjectPermission {
	this := RoleAssignedObjectPermission{}
	return &this
}

// GetRolePk returns the RolePk field value
func (o *RoleAssignedObjectPermission) GetRolePk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RolePk
}

// GetRolePkOk returns a tuple with the RolePk field value
// and a boolean to check if the value has been set.
func (o *RoleAssignedObjectPermission) GetRolePkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RolePk, true
}

// SetRolePk sets field value
func (o *RoleAssignedObjectPermission) SetRolePk(v string) {
	o.RolePk = v
}

// GetName returns the Name field value
func (o *RoleAssignedObjectPermission) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleAssignedObjectPermission) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleAssignedObjectPermission) SetName(v string) {
	o.Name = v
}

// GetPermissions returns the Permissions field value
func (o *RoleAssignedObjectPermission) GetPermissions() []RoleObjectPermission {
	if o == nil {
		var ret []RoleObjectPermission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *RoleAssignedObjectPermission) GetPermissionsOk() ([]RoleObjectPermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *RoleAssignedObjectPermission) SetPermissions(v []RoleObjectPermission) {
	o.Permissions = v
}

func (o RoleAssignedObjectPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role_pk"] = o.RolePk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableRoleAssignedObjectPermission struct {
	value *RoleAssignedObjectPermission
	isSet bool
}

func (v NullableRoleAssignedObjectPermission) Get() *RoleAssignedObjectPermission {
	return v.value
}

func (v *NullableRoleAssignedObjectPermission) Set(val *RoleAssignedObjectPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignedObjectPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignedObjectPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignedObjectPermission(val *RoleAssignedObjectPermission) *NullableRoleAssignedObjectPermission {
	return &NullableRoleAssignedObjectPermission{value: val, isSet: true}
}

func (v NullableRoleAssignedObjectPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignedObjectPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
