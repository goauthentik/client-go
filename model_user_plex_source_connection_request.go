/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserPlexSourceConnectionRequest Plex Source connection Serializer
type UserPlexSourceConnectionRequest struct {
	Identifier string `json:"identifier"`
	PlexToken  string `json:"plex_token"`
}

// NewUserPlexSourceConnectionRequest instantiates a new UserPlexSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPlexSourceConnectionRequest(identifier string, plexToken string) *UserPlexSourceConnectionRequest {
	this := UserPlexSourceConnectionRequest{}
	this.Identifier = identifier
	this.PlexToken = plexToken
	return &this
}

// NewUserPlexSourceConnectionRequestWithDefaults instantiates a new UserPlexSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPlexSourceConnectionRequestWithDefaults() *UserPlexSourceConnectionRequest {
	this := UserPlexSourceConnectionRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *UserPlexSourceConnectionRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *UserPlexSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *UserPlexSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = v
}

// GetPlexToken returns the PlexToken field value
func (o *UserPlexSourceConnectionRequest) GetPlexToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlexToken
}

// GetPlexTokenOk returns a tuple with the PlexToken field value
// and a boolean to check if the value has been set.
func (o *UserPlexSourceConnectionRequest) GetPlexTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlexToken, true
}

// SetPlexToken sets field value
func (o *UserPlexSourceConnectionRequest) SetPlexToken(v string) {
	o.PlexToken = v
}

func (o UserPlexSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["plex_token"] = o.PlexToken
	}
	return json.Marshal(toSerialize)
}

type NullableUserPlexSourceConnectionRequest struct {
	value *UserPlexSourceConnectionRequest
	isSet bool
}

func (v NullableUserPlexSourceConnectionRequest) Get() *UserPlexSourceConnectionRequest {
	return v.value
}

func (v *NullableUserPlexSourceConnectionRequest) Set(val *UserPlexSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPlexSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPlexSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPlexSourceConnectionRequest(val *UserPlexSourceConnectionRequest) *NullableUserPlexSourceConnectionRequest {
	return &NullableUserPlexSourceConnectionRequest{value: val, isSet: true}
}

func (v NullableUserPlexSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPlexSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
