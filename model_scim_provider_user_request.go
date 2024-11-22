/*
authentik

Making authentication simple.

API version: 2024.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMProviderUserRequest SCIMProviderUser Serializer
type SCIMProviderUserRequest struct {
	ScimId   string `json:"scim_id"`
	User     int32  `json:"user"`
	Provider int32  `json:"provider"`
}

// NewSCIMProviderUserRequest instantiates a new SCIMProviderUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMProviderUserRequest(scimId string, user int32, provider int32) *SCIMProviderUserRequest {
	this := SCIMProviderUserRequest{}
	this.ScimId = scimId
	this.User = user
	this.Provider = provider
	return &this
}

// NewSCIMProviderUserRequestWithDefaults instantiates a new SCIMProviderUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMProviderUserRequestWithDefaults() *SCIMProviderUserRequest {
	this := SCIMProviderUserRequest{}
	return &this
}

// GetScimId returns the ScimId field value
func (o *SCIMProviderUserRequest) GetScimId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScimId
}

// GetScimIdOk returns a tuple with the ScimId field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderUserRequest) GetScimIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScimId, true
}

// SetScimId sets field value
func (o *SCIMProviderUserRequest) SetScimId(v string) {
	o.ScimId = v
}

// GetUser returns the User field value
func (o *SCIMProviderUserRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderUserRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *SCIMProviderUserRequest) SetUser(v int32) {
	o.User = v
}

// GetProvider returns the Provider field value
func (o *SCIMProviderUserRequest) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderUserRequest) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SCIMProviderUserRequest) SetProvider(v int32) {
	o.Provider = v
}

func (o SCIMProviderUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scim_id"] = o.ScimId
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMProviderUserRequest struct {
	value *SCIMProviderUserRequest
	isSet bool
}

func (v NullableSCIMProviderUserRequest) Get() *SCIMProviderUserRequest {
	return v.value
}

func (v *NullableSCIMProviderUserRequest) Set(val *SCIMProviderUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMProviderUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMProviderUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMProviderUserRequest(val *SCIMProviderUserRequest) *NullableSCIMProviderUserRequest {
	return &NullableSCIMProviderUserRequest{value: val, isSet: true}
}

func (v NullableSCIMProviderUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMProviderUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
