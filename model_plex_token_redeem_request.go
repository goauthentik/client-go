/*
authentik

Making authentication simple.

API version: 2022.11.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PlexTokenRedeemRequest Serializer to redeem a plex token
type PlexTokenRedeemRequest struct {
	PlexToken string `json:"plex_token"`
}

// NewPlexTokenRedeemRequest instantiates a new PlexTokenRedeemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlexTokenRedeemRequest(plexToken string) *PlexTokenRedeemRequest {
	this := PlexTokenRedeemRequest{}
	this.PlexToken = plexToken
	return &this
}

// NewPlexTokenRedeemRequestWithDefaults instantiates a new PlexTokenRedeemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlexTokenRedeemRequestWithDefaults() *PlexTokenRedeemRequest {
	this := PlexTokenRedeemRequest{}
	return &this
}

// GetPlexToken returns the PlexToken field value
func (o *PlexTokenRedeemRequest) GetPlexToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlexToken
}

// GetPlexTokenOk returns a tuple with the PlexToken field value
// and a boolean to check if the value has been set.
func (o *PlexTokenRedeemRequest) GetPlexTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlexToken, true
}

// SetPlexToken sets field value
func (o *PlexTokenRedeemRequest) SetPlexToken(v string) {
	o.PlexToken = v
}

func (o PlexTokenRedeemRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["plex_token"] = o.PlexToken
	}
	return json.Marshal(toSerialize)
}

type NullablePlexTokenRedeemRequest struct {
	value *PlexTokenRedeemRequest
	isSet bool
}

func (v NullablePlexTokenRedeemRequest) Get() *PlexTokenRedeemRequest {
	return v.value
}

func (v *NullablePlexTokenRedeemRequest) Set(val *PlexTokenRedeemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePlexTokenRedeemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePlexTokenRedeemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlexTokenRedeemRequest(val *PlexTokenRedeemRequest) *NullablePlexTokenRedeemRequest {
	return &NullablePlexTokenRedeemRequest{value: val, isSet: true}
}

func (v NullablePlexTokenRedeemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlexTokenRedeemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
