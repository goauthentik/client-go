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

// GoogleWorkspaceProviderUserRequest GoogleWorkspaceProviderUser Serializer
type GoogleWorkspaceProviderUserRequest struct {
	GoogleId string `json:"google_id"`
	User     int32  `json:"user"`
	Provider int32  `json:"provider"`
}

// NewGoogleWorkspaceProviderUserRequest instantiates a new GoogleWorkspaceProviderUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleWorkspaceProviderUserRequest(googleId string, user int32, provider int32) *GoogleWorkspaceProviderUserRequest {
	this := GoogleWorkspaceProviderUserRequest{}
	this.GoogleId = googleId
	this.User = user
	this.Provider = provider
	return &this
}

// NewGoogleWorkspaceProviderUserRequestWithDefaults instantiates a new GoogleWorkspaceProviderUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleWorkspaceProviderUserRequestWithDefaults() *GoogleWorkspaceProviderUserRequest {
	this := GoogleWorkspaceProviderUserRequest{}
	return &this
}

// GetGoogleId returns the GoogleId field value
func (o *GoogleWorkspaceProviderUserRequest) GetGoogleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoogleId
}

// GetGoogleIdOk returns a tuple with the GoogleId field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderUserRequest) GetGoogleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoogleId, true
}

// SetGoogleId sets field value
func (o *GoogleWorkspaceProviderUserRequest) SetGoogleId(v string) {
	o.GoogleId = v
}

// GetUser returns the User field value
func (o *GoogleWorkspaceProviderUserRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderUserRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GoogleWorkspaceProviderUserRequest) SetUser(v int32) {
	o.User = v
}

// GetProvider returns the Provider field value
func (o *GoogleWorkspaceProviderUserRequest) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderUserRequest) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *GoogleWorkspaceProviderUserRequest) SetProvider(v int32) {
	o.Provider = v
}

func (o GoogleWorkspaceProviderUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["google_id"] = o.GoogleId
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleWorkspaceProviderUserRequest struct {
	value *GoogleWorkspaceProviderUserRequest
	isSet bool
}

func (v NullableGoogleWorkspaceProviderUserRequest) Get() *GoogleWorkspaceProviderUserRequest {
	return v.value
}

func (v *NullableGoogleWorkspaceProviderUserRequest) Set(val *GoogleWorkspaceProviderUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleWorkspaceProviderUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleWorkspaceProviderUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleWorkspaceProviderUserRequest(val *GoogleWorkspaceProviderUserRequest) *NullableGoogleWorkspaceProviderUserRequest {
	return &NullableGoogleWorkspaceProviderUserRequest{value: val, isSet: true}
}

func (v NullableGoogleWorkspaceProviderUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleWorkspaceProviderUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
