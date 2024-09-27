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

// PatchedUserPlexSourceConnectionRequest Plex Source connection Serializer
type PatchedUserPlexSourceConnectionRequest struct {
	Identifier *string `json:"identifier,omitempty"`
	PlexToken  *string `json:"plex_token,omitempty"`
}

// NewPatchedUserPlexSourceConnectionRequest instantiates a new PatchedUserPlexSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserPlexSourceConnectionRequest() *PatchedUserPlexSourceConnectionRequest {
	this := PatchedUserPlexSourceConnectionRequest{}
	return &this
}

// NewPatchedUserPlexSourceConnectionRequestWithDefaults instantiates a new PatchedUserPlexSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserPlexSourceConnectionRequestWithDefaults() *PatchedUserPlexSourceConnectionRequest {
	this := PatchedUserPlexSourceConnectionRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PatchedUserPlexSourceConnectionRequest) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserPlexSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PatchedUserPlexSourceConnectionRequest) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PatchedUserPlexSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetPlexToken returns the PlexToken field value if set, zero value otherwise.
func (o *PatchedUserPlexSourceConnectionRequest) GetPlexToken() string {
	if o == nil || o.PlexToken == nil {
		var ret string
		return ret
	}
	return *o.PlexToken
}

// GetPlexTokenOk returns a tuple with the PlexToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserPlexSourceConnectionRequest) GetPlexTokenOk() (*string, bool) {
	if o == nil || o.PlexToken == nil {
		return nil, false
	}
	return o.PlexToken, true
}

// HasPlexToken returns a boolean if a field has been set.
func (o *PatchedUserPlexSourceConnectionRequest) HasPlexToken() bool {
	if o != nil && o.PlexToken != nil {
		return true
	}

	return false
}

// SetPlexToken gets a reference to the given string and assigns it to the PlexToken field.
func (o *PatchedUserPlexSourceConnectionRequest) SetPlexToken(v string) {
	o.PlexToken = &v
}

func (o PatchedUserPlexSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.PlexToken != nil {
		toSerialize["plex_token"] = o.PlexToken
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedUserPlexSourceConnectionRequest struct {
	value *PatchedUserPlexSourceConnectionRequest
	isSet bool
}

func (v NullablePatchedUserPlexSourceConnectionRequest) Get() *PatchedUserPlexSourceConnectionRequest {
	return v.value
}

func (v *NullablePatchedUserPlexSourceConnectionRequest) Set(val *PatchedUserPlexSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserPlexSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserPlexSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserPlexSourceConnectionRequest(val *PatchedUserPlexSourceConnectionRequest) *NullablePatchedUserPlexSourceConnectionRequest {
	return &NullablePatchedUserPlexSourceConnectionRequest{value: val, isSet: true}
}

func (v NullablePatchedUserPlexSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserPlexSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
