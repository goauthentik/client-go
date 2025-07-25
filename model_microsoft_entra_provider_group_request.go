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

// MicrosoftEntraProviderGroupRequest MicrosoftEntraProviderGroup Serializer
type MicrosoftEntraProviderGroupRequest struct {
	MicrosoftId string `json:"microsoft_id"`
	Group       string `json:"group"`
	Provider    int32  `json:"provider"`
}

// NewMicrosoftEntraProviderGroupRequest instantiates a new MicrosoftEntraProviderGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftEntraProviderGroupRequest(microsoftId string, group string, provider int32) *MicrosoftEntraProviderGroupRequest {
	this := MicrosoftEntraProviderGroupRequest{}
	this.MicrosoftId = microsoftId
	this.Group = group
	this.Provider = provider
	return &this
}

// NewMicrosoftEntraProviderGroupRequestWithDefaults instantiates a new MicrosoftEntraProviderGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftEntraProviderGroupRequestWithDefaults() *MicrosoftEntraProviderGroupRequest {
	this := MicrosoftEntraProviderGroupRequest{}
	return &this
}

// GetMicrosoftId returns the MicrosoftId field value
func (o *MicrosoftEntraProviderGroupRequest) GetMicrosoftId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicrosoftId
}

// GetMicrosoftIdOk returns a tuple with the MicrosoftId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderGroupRequest) GetMicrosoftIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicrosoftId, true
}

// SetMicrosoftId sets field value
func (o *MicrosoftEntraProviderGroupRequest) SetMicrosoftId(v string) {
	o.MicrosoftId = v
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

// GetProvider returns the Provider field value
func (o *MicrosoftEntraProviderGroupRequest) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderGroupRequest) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *MicrosoftEntraProviderGroupRequest) SetProvider(v int32) {
	o.Provider = v
}

func (o MicrosoftEntraProviderGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["microsoft_id"] = o.MicrosoftId
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["provider"] = o.Provider
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
