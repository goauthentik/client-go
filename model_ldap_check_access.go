/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LDAPCheckAccess Base serializer class which doesn't implement create/update methods
type LDAPCheckAccess struct {
	HasSearchPermission *bool            `json:"has_search_permission,omitempty"`
	Access              PolicyTestResult `json:"access"`
}

// NewLDAPCheckAccess instantiates a new LDAPCheckAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPCheckAccess(access PolicyTestResult) *LDAPCheckAccess {
	this := LDAPCheckAccess{}
	this.Access = access
	return &this
}

// NewLDAPCheckAccessWithDefaults instantiates a new LDAPCheckAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPCheckAccessWithDefaults() *LDAPCheckAccess {
	this := LDAPCheckAccess{}
	return &this
}

// GetHasSearchPermission returns the HasSearchPermission field value if set, zero value otherwise.
func (o *LDAPCheckAccess) GetHasSearchPermission() bool {
	if o == nil || o.HasSearchPermission == nil {
		var ret bool
		return ret
	}
	return *o.HasSearchPermission
}

// GetHasSearchPermissionOk returns a tuple with the HasSearchPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPCheckAccess) GetHasSearchPermissionOk() (*bool, bool) {
	if o == nil || o.HasSearchPermission == nil {
		return nil, false
	}
	return o.HasSearchPermission, true
}

// HasHasSearchPermission returns a boolean if a field has been set.
func (o *LDAPCheckAccess) HasHasSearchPermission() bool {
	if o != nil && o.HasSearchPermission != nil {
		return true
	}

	return false
}

// SetHasSearchPermission gets a reference to the given bool and assigns it to the HasSearchPermission field.
func (o *LDAPCheckAccess) SetHasSearchPermission(v bool) {
	o.HasSearchPermission = &v
}

// GetAccess returns the Access field value
func (o *LDAPCheckAccess) GetAccess() PolicyTestResult {
	if o == nil {
		var ret PolicyTestResult
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *LDAPCheckAccess) GetAccessOk() (*PolicyTestResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *LDAPCheckAccess) SetAccess(v PolicyTestResult) {
	o.Access = v
}

func (o LDAPCheckAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HasSearchPermission != nil {
		toSerialize["has_search_permission"] = o.HasSearchPermission
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPCheckAccess struct {
	value *LDAPCheckAccess
	isSet bool
}

func (v NullableLDAPCheckAccess) Get() *LDAPCheckAccess {
	return v.value
}

func (v *NullableLDAPCheckAccess) Set(val *LDAPCheckAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPCheckAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPCheckAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPCheckAccess(val *LDAPCheckAccess) *NullableLDAPCheckAccess {
	return &NullableLDAPCheckAccess{value: val, isSet: true}
}

func (v NullableLDAPCheckAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPCheckAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
