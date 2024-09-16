/*
authentik

Making authentication simple.

API version: 2024.8.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMProviderGroupRequest SCIMProviderGroup Serializer
type SCIMProviderGroupRequest struct {
	ScimId   string `json:"scim_id"`
	Group    string `json:"group"`
	Provider int32  `json:"provider"`
}

// NewSCIMProviderGroupRequest instantiates a new SCIMProviderGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMProviderGroupRequest(scimId string, group string, provider int32) *SCIMProviderGroupRequest {
	this := SCIMProviderGroupRequest{}
	this.ScimId = scimId
	this.Group = group
	this.Provider = provider
	return &this
}

// NewSCIMProviderGroupRequestWithDefaults instantiates a new SCIMProviderGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMProviderGroupRequestWithDefaults() *SCIMProviderGroupRequest {
	this := SCIMProviderGroupRequest{}
	return &this
}

// GetScimId returns the ScimId field value
func (o *SCIMProviderGroupRequest) GetScimId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScimId
}

// GetScimIdOk returns a tuple with the ScimId field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderGroupRequest) GetScimIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScimId, true
}

// SetScimId sets field value
func (o *SCIMProviderGroupRequest) SetScimId(v string) {
	o.ScimId = v
}

// GetGroup returns the Group field value
func (o *SCIMProviderGroupRequest) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderGroupRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SCIMProviderGroupRequest) SetGroup(v string) {
	o.Group = v
}

// GetProvider returns the Provider field value
func (o *SCIMProviderGroupRequest) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderGroupRequest) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SCIMProviderGroupRequest) SetProvider(v int32) {
	o.Provider = v
}

func (o SCIMProviderGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scim_id"] = o.ScimId
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMProviderGroupRequest struct {
	value *SCIMProviderGroupRequest
	isSet bool
}

func (v NullableSCIMProviderGroupRequest) Get() *SCIMProviderGroupRequest {
	return v.value
}

func (v *NullableSCIMProviderGroupRequest) Set(val *SCIMProviderGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMProviderGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMProviderGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMProviderGroupRequest(val *SCIMProviderGroupRequest) *NullableSCIMProviderGroupRequest {
	return &NullableSCIMProviderGroupRequest{value: val, isSet: true}
}

func (v NullableSCIMProviderGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMProviderGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
