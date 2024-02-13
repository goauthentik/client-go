/*
authentik

Making authentication simple.

API version: 2023.10.7
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OAuthDeviceCodeFinishChallengeResponseRequest Response that device has been authenticated and tab can be closed
type OAuthDeviceCodeFinishChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
}

// NewOAuthDeviceCodeFinishChallengeResponseRequest instantiates a new OAuthDeviceCodeFinishChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthDeviceCodeFinishChallengeResponseRequest() *OAuthDeviceCodeFinishChallengeResponseRequest {
	this := OAuthDeviceCodeFinishChallengeResponseRequest{}
	var component string = "ak-provider-oauth2-device-code-finish"
	this.Component = &component
	return &this
}

// NewOAuthDeviceCodeFinishChallengeResponseRequestWithDefaults instantiates a new OAuthDeviceCodeFinishChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthDeviceCodeFinishChallengeResponseRequestWithDefaults() *OAuthDeviceCodeFinishChallengeResponseRequest {
	this := OAuthDeviceCodeFinishChallengeResponseRequest{}
	var component string = "ak-provider-oauth2-device-code-finish"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *OAuthDeviceCodeFinishChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthDeviceCodeFinishChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *OAuthDeviceCodeFinishChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *OAuthDeviceCodeFinishChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

func (o OAuthDeviceCodeFinishChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullableOAuthDeviceCodeFinishChallengeResponseRequest struct {
	value *OAuthDeviceCodeFinishChallengeResponseRequest
	isSet bool
}

func (v NullableOAuthDeviceCodeFinishChallengeResponseRequest) Get() *OAuthDeviceCodeFinishChallengeResponseRequest {
	return v.value
}

func (v *NullableOAuthDeviceCodeFinishChallengeResponseRequest) Set(val *OAuthDeviceCodeFinishChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthDeviceCodeFinishChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthDeviceCodeFinishChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthDeviceCodeFinishChallengeResponseRequest(val *OAuthDeviceCodeFinishChallengeResponseRequest) *NullableOAuthDeviceCodeFinishChallengeResponseRequest {
	return &NullableOAuthDeviceCodeFinishChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableOAuthDeviceCodeFinishChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthDeviceCodeFinishChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
