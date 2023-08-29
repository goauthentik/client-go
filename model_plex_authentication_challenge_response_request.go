/*
authentik

Making authentication simple.

API version: 2023.8.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PlexAuthenticationChallengeResponseRequest Pseudo class for plex response
type PlexAuthenticationChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
}

// NewPlexAuthenticationChallengeResponseRequest instantiates a new PlexAuthenticationChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlexAuthenticationChallengeResponseRequest() *PlexAuthenticationChallengeResponseRequest {
	this := PlexAuthenticationChallengeResponseRequest{}
	var component string = "ak-source-plex"
	this.Component = &component
	return &this
}

// NewPlexAuthenticationChallengeResponseRequestWithDefaults instantiates a new PlexAuthenticationChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlexAuthenticationChallengeResponseRequestWithDefaults() *PlexAuthenticationChallengeResponseRequest {
	this := PlexAuthenticationChallengeResponseRequest{}
	var component string = "ak-source-plex"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *PlexAuthenticationChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexAuthenticationChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *PlexAuthenticationChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *PlexAuthenticationChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

func (o PlexAuthenticationChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullablePlexAuthenticationChallengeResponseRequest struct {
	value *PlexAuthenticationChallengeResponseRequest
	isSet bool
}

func (v NullablePlexAuthenticationChallengeResponseRequest) Get() *PlexAuthenticationChallengeResponseRequest {
	return v.value
}

func (v *NullablePlexAuthenticationChallengeResponseRequest) Set(val *PlexAuthenticationChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePlexAuthenticationChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePlexAuthenticationChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlexAuthenticationChallengeResponseRequest(val *PlexAuthenticationChallengeResponseRequest) *NullablePlexAuthenticationChallengeResponseRequest {
	return &NullablePlexAuthenticationChallengeResponseRequest{value: val, isSet: true}
}

func (v NullablePlexAuthenticationChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlexAuthenticationChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
