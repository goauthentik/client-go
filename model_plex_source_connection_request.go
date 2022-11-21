/*
authentik

Making authentication simple.

API version: 2022.11.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PlexSourceConnectionRequest Plex Source connection Serializer
type PlexSourceConnectionRequest struct {
	Identifier string `json:"identifier"`
	PlexToken  string `json:"plex_token"`
}

// NewPlexSourceConnectionRequest instantiates a new PlexSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlexSourceConnectionRequest(identifier string, plexToken string) *PlexSourceConnectionRequest {
	this := PlexSourceConnectionRequest{}
	this.Identifier = identifier
	this.PlexToken = plexToken
	return &this
}

// NewPlexSourceConnectionRequestWithDefaults instantiates a new PlexSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlexSourceConnectionRequestWithDefaults() *PlexSourceConnectionRequest {
	this := PlexSourceConnectionRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *PlexSourceConnectionRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *PlexSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = v
}

// GetPlexToken returns the PlexToken field value
func (o *PlexSourceConnectionRequest) GetPlexToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlexToken
}

// GetPlexTokenOk returns a tuple with the PlexToken field value
// and a boolean to check if the value has been set.
func (o *PlexSourceConnectionRequest) GetPlexTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlexToken, true
}

// SetPlexToken sets field value
func (o *PlexSourceConnectionRequest) SetPlexToken(v string) {
	o.PlexToken = v
}

func (o PlexSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["plex_token"] = o.PlexToken
	}
	return json.Marshal(toSerialize)
}

type NullablePlexSourceConnectionRequest struct {
	value *PlexSourceConnectionRequest
	isSet bool
}

func (v NullablePlexSourceConnectionRequest) Get() *PlexSourceConnectionRequest {
	return v.value
}

func (v *NullablePlexSourceConnectionRequest) Set(val *PlexSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePlexSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePlexSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlexSourceConnectionRequest(val *PlexSourceConnectionRequest) *NullablePlexSourceConnectionRequest {
	return &NullablePlexSourceConnectionRequest{value: val, isSet: true}
}

func (v NullablePlexSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlexSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
