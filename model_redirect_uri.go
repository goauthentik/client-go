/*
authentik

Making authentication simple.

API version: 2024.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RedirectURI A single allowed redirect URI entry
type RedirectURI struct {
	MatchingMode MatchingModeEnum `json:"matching_mode"`
	Url          string           `json:"url"`
}

// NewRedirectURI instantiates a new RedirectURI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectURI(matchingMode MatchingModeEnum, url string) *RedirectURI {
	this := RedirectURI{}
	this.MatchingMode = matchingMode
	this.Url = url
	return &this
}

// NewRedirectURIWithDefaults instantiates a new RedirectURI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectURIWithDefaults() *RedirectURI {
	this := RedirectURI{}
	return &this
}

// GetMatchingMode returns the MatchingMode field value
func (o *RedirectURI) GetMatchingMode() MatchingModeEnum {
	if o == nil {
		var ret MatchingModeEnum
		return ret
	}

	return o.MatchingMode
}

// GetMatchingModeOk returns a tuple with the MatchingMode field value
// and a boolean to check if the value has been set.
func (o *RedirectURI) GetMatchingModeOk() (*MatchingModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchingMode, true
}

// SetMatchingMode sets field value
func (o *RedirectURI) SetMatchingMode(v MatchingModeEnum) {
	o.MatchingMode = v
}

// GetUrl returns the Url field value
func (o *RedirectURI) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RedirectURI) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RedirectURI) SetUrl(v string) {
	o.Url = v
}

func (o RedirectURI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["matching_mode"] = o.MatchingMode
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableRedirectURI struct {
	value *RedirectURI
	isSet bool
}

func (v NullableRedirectURI) Get() *RedirectURI {
	return v.value
}

func (v *NullableRedirectURI) Set(val *RedirectURI) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectURI) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectURI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectURI(val *RedirectURI) *NullableRedirectURI {
	return &NullableRedirectURI{value: val, isSet: true}
}

func (v NullableRedirectURI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectURI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
