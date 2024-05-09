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

// MicrosoftEntraProviderGroupRequest MicrosoftEntraProviderGroup Serializer
type MicrosoftEntraProviderGroupRequest struct {
	Group string `json:"group"`
}

// NewMicrosoftEntraProviderGroupRequest instantiates a new MicrosoftEntraProviderGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftEntraProviderGroupRequest(group string) *MicrosoftEntraProviderGroupRequest {
	this := MicrosoftEntraProviderGroupRequest{}
	this.Group = group
	return &this
}

// NewMicrosoftEntraProviderGroupRequestWithDefaults instantiates a new MicrosoftEntraProviderGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftEntraProviderGroupRequestWithDefaults() *MicrosoftEntraProviderGroupRequest {
	this := MicrosoftEntraProviderGroupRequest{}
	return &this
}

// GetGroup returns the Group field value
func (o *MicrosoftEntraProviderGroupRequest) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderGroupRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *MicrosoftEntraProviderGroupRequest) SetGroup(v string) {
	o.Group = v
}

func (o MicrosoftEntraProviderGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftEntraProviderGroupRequest struct {
	value *MicrosoftEntraProviderGroupRequest
	isSet bool
}

func (v NullableMicrosoftEntraProviderGroupRequest) Get() *MicrosoftEntraProviderGroupRequest {
	return v.value
}

func (v *NullableMicrosoftEntraProviderGroupRequest) Set(val *MicrosoftEntraProviderGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftEntraProviderGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftEntraProviderGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftEntraProviderGroupRequest(val *MicrosoftEntraProviderGroupRequest) *NullableMicrosoftEntraProviderGroupRequest {
	return &NullableMicrosoftEntraProviderGroupRequest{value: val, isSet: true}
}

func (v NullableMicrosoftEntraProviderGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftEntraProviderGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
