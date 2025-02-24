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

// GoogleWorkspaceProviderGroupRequest GoogleWorkspaceProviderGroup Serializer
type GoogleWorkspaceProviderGroupRequest struct {
	GoogleId string `json:"google_id"`
	Group    string `json:"group"`
	Provider int32  `json:"provider"`
}

// NewGoogleWorkspaceProviderGroupRequest instantiates a new GoogleWorkspaceProviderGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleWorkspaceProviderGroupRequest(googleId string, group string, provider int32) *GoogleWorkspaceProviderGroupRequest {
	this := GoogleWorkspaceProviderGroupRequest{}
	this.GoogleId = googleId
	this.Group = group
	this.Provider = provider
	return &this
}

// NewGoogleWorkspaceProviderGroupRequestWithDefaults instantiates a new GoogleWorkspaceProviderGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleWorkspaceProviderGroupRequestWithDefaults() *GoogleWorkspaceProviderGroupRequest {
	this := GoogleWorkspaceProviderGroupRequest{}
	return &this
}

// GetGoogleId returns the GoogleId field value
func (o *GoogleWorkspaceProviderGroupRequest) GetGoogleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoogleId
}

// GetGoogleIdOk returns a tuple with the GoogleId field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderGroupRequest) GetGoogleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoogleId, true
}

// SetGoogleId sets field value
func (o *GoogleWorkspaceProviderGroupRequest) SetGoogleId(v string) {
	o.GoogleId = v
}

// GetGroup returns the Group field value
func (o *GoogleWorkspaceProviderGroupRequest) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderGroupRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *GoogleWorkspaceProviderGroupRequest) SetGroup(v string) {
	o.Group = v
}

// GetProvider returns the Provider field value
func (o *GoogleWorkspaceProviderGroupRequest) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderGroupRequest) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *GoogleWorkspaceProviderGroupRequest) SetProvider(v int32) {
	o.Provider = v
}

func (o GoogleWorkspaceProviderGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["google_id"] = o.GoogleId
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleWorkspaceProviderGroupRequest struct {
	value *GoogleWorkspaceProviderGroupRequest
	isSet bool
}

func (v NullableGoogleWorkspaceProviderGroupRequest) Get() *GoogleWorkspaceProviderGroupRequest {
	return v.value
}

func (v *NullableGoogleWorkspaceProviderGroupRequest) Set(val *GoogleWorkspaceProviderGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleWorkspaceProviderGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleWorkspaceProviderGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleWorkspaceProviderGroupRequest(val *GoogleWorkspaceProviderGroupRequest) *NullableGoogleWorkspaceProviderGroupRequest {
	return &NullableGoogleWorkspaceProviderGroupRequest{value: val, isSet: true}
}

func (v NullableGoogleWorkspaceProviderGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleWorkspaceProviderGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
