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

// PermissionRequest Global permission
type PermissionRequest struct {
	Name     string `json:"name"`
	Codename string `json:"codename"`
}

// NewPermissionRequest instantiates a new PermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionRequest(name string, codename string) *PermissionRequest {
	this := PermissionRequest{}
	this.Name = name
	this.Codename = codename
	return &this
}

// NewPermissionRequestWithDefaults instantiates a new PermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionRequestWithDefaults() *PermissionRequest {
	this := PermissionRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PermissionRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PermissionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PermissionRequest) SetName(v string) {
	o.Name = v
}

// GetCodename returns the Codename field value
func (o *PermissionRequest) GetCodename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Codename
}

// GetCodenameOk returns a tuple with the Codename field value
// and a boolean to check if the value has been set.
func (o *PermissionRequest) GetCodenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Codename, true
}

// SetCodename sets field value
func (o *PermissionRequest) SetCodename(v string) {
	o.Codename = v
}

func (o PermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["codename"] = o.Codename
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionRequest struct {
	value *PermissionRequest
	isSet bool
}

func (v NullablePermissionRequest) Get() *PermissionRequest {
	return v.value
}

func (v *NullablePermissionRequest) Set(val *PermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionRequest(val *PermissionRequest) *NullablePermissionRequest {
	return &NullablePermissionRequest{value: val, isSet: true}
}

func (v NullablePermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
