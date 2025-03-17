/*
authentik

Making authentication simple.

API version: 2025.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OAuthDeviceCodeChallengeResponseRequest Response that includes the user-entered device code
type OAuthDeviceCodeChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
	Code      string  `json:"code"`
}

// NewOAuthDeviceCodeChallengeResponseRequest instantiates a new OAuthDeviceCodeChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthDeviceCodeChallengeResponseRequest(code string) *OAuthDeviceCodeChallengeResponseRequest {
	this := OAuthDeviceCodeChallengeResponseRequest{}
	var component string = "ak-provider-oauth2-device-code"
	this.Component = &component
	this.Code = code
	return &this
}

// NewOAuthDeviceCodeChallengeResponseRequestWithDefaults instantiates a new OAuthDeviceCodeChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthDeviceCodeChallengeResponseRequestWithDefaults() *OAuthDeviceCodeChallengeResponseRequest {
	this := OAuthDeviceCodeChallengeResponseRequest{}
	var component string = "ak-provider-oauth2-device-code"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *OAuthDeviceCodeChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthDeviceCodeChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *OAuthDeviceCodeChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *OAuthDeviceCodeChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

// GetCode returns the Code field value
func (o *OAuthDeviceCodeChallengeResponseRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OAuthDeviceCodeChallengeResponseRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OAuthDeviceCodeChallengeResponseRequest) SetCode(v string) {
	o.Code = v
}

func (o OAuthDeviceCodeChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthDeviceCodeChallengeResponseRequest struct {
	value *OAuthDeviceCodeChallengeResponseRequest
	isSet bool
}

func (v NullableOAuthDeviceCodeChallengeResponseRequest) Get() *OAuthDeviceCodeChallengeResponseRequest {
	return v.value
}

func (v *NullableOAuthDeviceCodeChallengeResponseRequest) Set(val *OAuthDeviceCodeChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthDeviceCodeChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthDeviceCodeChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthDeviceCodeChallengeResponseRequest(val *OAuthDeviceCodeChallengeResponseRequest) *NullableOAuthDeviceCodeChallengeResponseRequest {
	return &NullableOAuthDeviceCodeChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableOAuthDeviceCodeChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthDeviceCodeChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
