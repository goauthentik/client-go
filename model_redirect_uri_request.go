/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RedirectURIRequest A single allowed redirect URI entry
type RedirectURIRequest struct {
	MatchingMode MatchingModeEnum `json:"matching_mode"`
	Url          string           `json:"url"`
}

// NewRedirectURIRequest instantiates a new RedirectURIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectURIRequest(matchingMode MatchingModeEnum, url string) *RedirectURIRequest {
	this := RedirectURIRequest{}
	this.MatchingMode = matchingMode
	this.Url = url
	return &this
}

// NewRedirectURIRequestWithDefaults instantiates a new RedirectURIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectURIRequestWithDefaults() *RedirectURIRequest {
	this := RedirectURIRequest{}
	return &this
}

// GetMatchingMode returns the MatchingMode field value
func (o *RedirectURIRequest) GetMatchingMode() MatchingModeEnum {
	if o == nil {
		var ret MatchingModeEnum
		return ret
	}

	return o.MatchingMode
}

// GetMatchingModeOk returns a tuple with the MatchingMode field value
// and a boolean to check if the value has been set.
func (o *RedirectURIRequest) GetMatchingModeOk() (*MatchingModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchingMode, true
}

// SetMatchingMode sets field value
func (o *RedirectURIRequest) SetMatchingMode(v MatchingModeEnum) {
	o.MatchingMode = v
}

// GetUrl returns the Url field value
func (o *RedirectURIRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RedirectURIRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RedirectURIRequest) SetUrl(v string) {
	o.Url = v
}

func (o RedirectURIRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["matching_mode"] = o.MatchingMode
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableRedirectURIRequest struct {
	value *RedirectURIRequest
	isSet bool
}

func (v NullableRedirectURIRequest) Get() *RedirectURIRequest {
	return v.value
}

func (v *NullableRedirectURIRequest) Set(val *RedirectURIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectURIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectURIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectURIRequest(val *RedirectURIRequest) *NullableRedirectURIRequest {
	return &NullableRedirectURIRequest{value: val, isSet: true}
}

func (v NullableRedirectURIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectURIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
