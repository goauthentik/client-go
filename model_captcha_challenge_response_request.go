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

// CaptchaChallengeResponseRequest Validate captcha token
type CaptchaChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
	Token     string  `json:"token"`
}

// NewCaptchaChallengeResponseRequest instantiates a new CaptchaChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptchaChallengeResponseRequest(token string) *CaptchaChallengeResponseRequest {
	this := CaptchaChallengeResponseRequest{}
	var component string = "ak-stage-captcha"
	this.Component = &component
	this.Token = token
	return &this
}

// NewCaptchaChallengeResponseRequestWithDefaults instantiates a new CaptchaChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptchaChallengeResponseRequestWithDefaults() *CaptchaChallengeResponseRequest {
	this := CaptchaChallengeResponseRequest{}
	var component string = "ak-stage-captcha"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *CaptchaChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *CaptchaChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *CaptchaChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

// GetToken returns the Token field value
func (o *CaptchaChallengeResponseRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CaptchaChallengeResponseRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CaptchaChallengeResponseRequest) SetToken(v string) {
	o.Token = v
}

func (o CaptchaChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableCaptchaChallengeResponseRequest struct {
	value *CaptchaChallengeResponseRequest
	isSet bool
}

func (v NullableCaptchaChallengeResponseRequest) Get() *CaptchaChallengeResponseRequest {
	return v.value
}

func (v *NullableCaptchaChallengeResponseRequest) Set(val *CaptchaChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptchaChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptchaChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptchaChallengeResponseRequest(val *CaptchaChallengeResponseRequest) *NullableCaptchaChallengeResponseRequest {
	return &NullableCaptchaChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableCaptchaChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptchaChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
