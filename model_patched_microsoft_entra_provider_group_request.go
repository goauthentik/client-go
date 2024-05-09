/*
authentik

Making authentication simple.

API version: 2024.4.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedMicrosoftEntraProviderGroupRequest MicrosoftEntraProviderGroup Serializer
type PatchedMicrosoftEntraProviderGroupRequest struct {
	Group *string `json:"group,omitempty"`
}

// NewPatchedMicrosoftEntraProviderGroupRequest instantiates a new PatchedMicrosoftEntraProviderGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedMicrosoftEntraProviderGroupRequest() *PatchedMicrosoftEntraProviderGroupRequest {
	this := PatchedMicrosoftEntraProviderGroupRequest{}
	return &this
}

// NewPatchedMicrosoftEntraProviderGroupRequestWithDefaults instantiates a new PatchedMicrosoftEntraProviderGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedMicrosoftEntraProviderGroupRequestWithDefaults() *PatchedMicrosoftEntraProviderGroupRequest {
	this := PatchedMicrosoftEntraProviderGroupRequest{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PatchedMicrosoftEntraProviderGroupRequest) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMicrosoftEntraProviderGroupRequest) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PatchedMicrosoftEntraProviderGroupRequest) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *PatchedMicrosoftEntraProviderGroupRequest) SetGroup(v string) {
	o.Group = &v
}

func (o PatchedMicrosoftEntraProviderGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedMicrosoftEntraProviderGroupRequest struct {
	value *PatchedMicrosoftEntraProviderGroupRequest
	isSet bool
}

func (v NullablePatchedMicrosoftEntraProviderGroupRequest) Get() *PatchedMicrosoftEntraProviderGroupRequest {
	return v.value
}

func (v *NullablePatchedMicrosoftEntraProviderGroupRequest) Set(val *PatchedMicrosoftEntraProviderGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedMicrosoftEntraProviderGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedMicrosoftEntraProviderGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedMicrosoftEntraProviderGroupRequest(val *PatchedMicrosoftEntraProviderGroupRequest) *NullablePatchedMicrosoftEntraProviderGroupRequest {
	return &NullablePatchedMicrosoftEntraProviderGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedMicrosoftEntraProviderGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedMicrosoftEntraProviderGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}