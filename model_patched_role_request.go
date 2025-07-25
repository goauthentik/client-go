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

// PatchedRoleRequest Role serializer
type PatchedRoleRequest struct {
	Name *string `json:"name,omitempty"`
}

// NewPatchedRoleRequest instantiates a new PatchedRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRoleRequest() *PatchedRoleRequest {
	this := PatchedRoleRequest{}
	return &this
}

// NewPatchedRoleRequestWithDefaults instantiates a new PatchedRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRoleRequestWithDefaults() *PatchedRoleRequest {
	this := PatchedRoleRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedRoleRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRoleRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedRoleRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedRoleRequest) SetName(v string) {
	o.Name = &v
}

func (o PatchedRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedRoleRequest struct {
	value *PatchedRoleRequest
	isSet bool
}

func (v NullablePatchedRoleRequest) Get() *PatchedRoleRequest {
	return v.value
}

func (v *NullablePatchedRoleRequest) Set(val *PatchedRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRoleRequest(val *PatchedRoleRequest) *NullablePatchedRoleRequest {
	return &NullablePatchedRoleRequest{value: val, isSet: true}
}

func (v NullablePatchedRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
