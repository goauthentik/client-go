/*
authentik

Making authentication simple.

API version: 2023.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatedSessionUserAgent Get parsed user agent
type AuthenticatedSessionUserAgent struct {
	Device    AuthenticatedSessionUserAgentDevice    `json:"device"`
	Os        AuthenticatedSessionUserAgentOs        `json:"os"`
	UserAgent AuthenticatedSessionUserAgentUserAgent `json:"user_agent"`
	String    string                                 `json:"string"`
}

// NewAuthenticatedSessionUserAgent instantiates a new AuthenticatedSessionUserAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSessionUserAgent(device AuthenticatedSessionUserAgentDevice, os AuthenticatedSessionUserAgentOs, userAgent AuthenticatedSessionUserAgentUserAgent, string_ string) *AuthenticatedSessionUserAgent {
	this := AuthenticatedSessionUserAgent{}
	this.Device = device
	this.Os = os
	this.UserAgent = userAgent
	this.String = string_
	return &this
}

// NewAuthenticatedSessionUserAgentWithDefaults instantiates a new AuthenticatedSessionUserAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionUserAgentWithDefaults() *AuthenticatedSessionUserAgent {
	this := AuthenticatedSessionUserAgent{}
	return &this
}

// GetDevice returns the Device field value
func (o *AuthenticatedSessionUserAgent) GetDevice() AuthenticatedSessionUserAgentDevice {
	if o == nil {
		var ret AuthenticatedSessionUserAgentDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgent) GetDeviceOk() (*AuthenticatedSessionUserAgentDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *AuthenticatedSessionUserAgent) SetDevice(v AuthenticatedSessionUserAgentDevice) {
	o.Device = v
}

// GetOs returns the Os field value
func (o *AuthenticatedSessionUserAgent) GetOs() AuthenticatedSessionUserAgentOs {
	if o == nil {
		var ret AuthenticatedSessionUserAgentOs
		return ret
	}

	return o.Os
}

// GetOsOk returns a tuple with the Os field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgent) GetOsOk() (*AuthenticatedSessionUserAgentOs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Os, true
}

// SetOs sets field value
func (o *AuthenticatedSessionUserAgent) SetOs(v AuthenticatedSessionUserAgentOs) {
	o.Os = v
}

// GetUserAgent returns the UserAgent field value
func (o *AuthenticatedSessionUserAgent) GetUserAgent() AuthenticatedSessionUserAgentUserAgent {
	if o == nil {
		var ret AuthenticatedSessionUserAgentUserAgent
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgent) GetUserAgentOk() (*AuthenticatedSessionUserAgentUserAgent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *AuthenticatedSessionUserAgent) SetUserAgent(v AuthenticatedSessionUserAgentUserAgent) {
	o.UserAgent = v
}

// GetString returns the String field value
func (o *AuthenticatedSessionUserAgent) GetString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.String
}

// GetStringOk returns a tuple with the String field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgent) GetStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.String, true
}

// SetString sets field value
func (o *AuthenticatedSessionUserAgent) SetString(v string) {
	o.String = v
}

func (o AuthenticatedSessionUserAgent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["device"] = o.Device
	}
	if true {
		toSerialize["os"] = o.Os
	}
	if true {
		toSerialize["user_agent"] = o.UserAgent
	}
	if true {
		toSerialize["string"] = o.String
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatedSessionUserAgent struct {
	value *AuthenticatedSessionUserAgent
	isSet bool
}

func (v NullableAuthenticatedSessionUserAgent) Get() *AuthenticatedSessionUserAgent {
	return v.value
}

func (v *NullableAuthenticatedSessionUserAgent) Set(val *AuthenticatedSessionUserAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatedSessionUserAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatedSessionUserAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatedSessionUserAgent(val *AuthenticatedSessionUserAgent) *NullableAuthenticatedSessionUserAgent {
	return &NullableAuthenticatedSessionUserAgent{value: val, isSet: true}
}

func (v NullableAuthenticatedSessionUserAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatedSessionUserAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
