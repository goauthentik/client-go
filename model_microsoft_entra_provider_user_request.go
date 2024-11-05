/*
authentik

Making authentication simple.

API version: 2024.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MicrosoftEntraProviderUserRequest MicrosoftEntraProviderUser Serializer
type MicrosoftEntraProviderUserRequest struct {
	MicrosoftId string `json:"microsoft_id"`
	User        int32  `json:"user"`
	Provider    int32  `json:"provider"`
}

// NewMicrosoftEntraProviderUserRequest instantiates a new MicrosoftEntraProviderUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftEntraProviderUserRequest(microsoftId string, user int32, provider int32) *MicrosoftEntraProviderUserRequest {
	this := MicrosoftEntraProviderUserRequest{}
	this.MicrosoftId = microsoftId
	this.User = user
	this.Provider = provider
	return &this
}

// NewMicrosoftEntraProviderUserRequestWithDefaults instantiates a new MicrosoftEntraProviderUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftEntraProviderUserRequestWithDefaults() *MicrosoftEntraProviderUserRequest {
	this := MicrosoftEntraProviderUserRequest{}
	return &this
}

// GetMicrosoftId returns the MicrosoftId field value
func (o *MicrosoftEntraProviderUserRequest) GetMicrosoftId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicrosoftId
}

// GetMicrosoftIdOk returns a tuple with the MicrosoftId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderUserRequest) GetMicrosoftIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicrosoftId, true
}

// SetMicrosoftId sets field value
func (o *MicrosoftEntraProviderUserRequest) SetMicrosoftId(v string) {
	o.MicrosoftId = v
}

// GetUser returns the User field value
func (o *MicrosoftEntraProviderUserRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderUserRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *MicrosoftEntraProviderUserRequest) SetUser(v int32) {
	o.User = v
}

// GetProvider returns the Provider field value
func (o *MicrosoftEntraProviderUserRequest) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderUserRequest) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *MicrosoftEntraProviderUserRequest) SetProvider(v int32) {
	o.Provider = v
}

func (o MicrosoftEntraProviderUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["microsoft_id"] = o.MicrosoftId
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftEntraProviderUserRequest struct {
	value *MicrosoftEntraProviderUserRequest
	isSet bool
}

func (v NullableMicrosoftEntraProviderUserRequest) Get() *MicrosoftEntraProviderUserRequest {
	return v.value
}

func (v *NullableMicrosoftEntraProviderUserRequest) Set(val *MicrosoftEntraProviderUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftEntraProviderUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftEntraProviderUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftEntraProviderUserRequest(val *MicrosoftEntraProviderUserRequest) *NullableMicrosoftEntraProviderUserRequest {
	return &NullableMicrosoftEntraProviderUserRequest{value: val, isSet: true}
}

func (v NullableMicrosoftEntraProviderUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftEntraProviderUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
