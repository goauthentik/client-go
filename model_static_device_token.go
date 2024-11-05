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

// StaticDeviceToken Serializer for static device's tokens
type StaticDeviceToken struct {
	Token string `json:"token"`
}

// NewStaticDeviceToken instantiates a new StaticDeviceToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticDeviceToken(token string) *StaticDeviceToken {
	this := StaticDeviceToken{}
	this.Token = token
	return &this
}

// NewStaticDeviceTokenWithDefaults instantiates a new StaticDeviceToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticDeviceTokenWithDefaults() *StaticDeviceToken {
	this := StaticDeviceToken{}
	return &this
}

// GetToken returns the Token field value
func (o *StaticDeviceToken) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *StaticDeviceToken) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *StaticDeviceToken) SetToken(v string) {
	o.Token = v
}

func (o StaticDeviceToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableStaticDeviceToken struct {
	value *StaticDeviceToken
	isSet bool
}

func (v NullableStaticDeviceToken) Get() *StaticDeviceToken {
	return v.value
}

func (v *NullableStaticDeviceToken) Set(val *StaticDeviceToken) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticDeviceToken) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticDeviceToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticDeviceToken(val *StaticDeviceToken) *NullableStaticDeviceToken {
	return &NullableStaticDeviceToken{value: val, isSet: true}
}

func (v NullableStaticDeviceToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticDeviceToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
