/*
authentik

Making authentication simple.

API version: 2021.10.1-rc2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatedSessionUserAgent struct for AuthenticatedSessionUserAgent
type AuthenticatedSessionUserAgent struct {
	Device    *AuthenticatedSessionUserAgentDevice    `json:"device,omitempty"`
	Os        *AuthenticatedSessionUserAgentOs        `json:"os,omitempty"`
	UserAgent *AuthenticatedSessionUserAgentUserAgent `json:"user_agent,omitempty"`
	String    *string                                 `json:"string,omitempty"`
}

// NewAuthenticatedSessionUserAgent instantiates a new AuthenticatedSessionUserAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSessionUserAgent() *AuthenticatedSessionUserAgent {
	this := AuthenticatedSessionUserAgent{}
	return &this
}

// NewAuthenticatedSessionUserAgentWithDefaults instantiates a new AuthenticatedSessionUserAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionUserAgentWithDefaults() *AuthenticatedSessionUserAgent {
	this := AuthenticatedSessionUserAgent{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgent) GetDevice() AuthenticatedSessionUserAgentDevice {
	if o == nil || o.Device == nil {
		var ret AuthenticatedSessionUserAgentDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgent) GetDeviceOk() (*AuthenticatedSessionUserAgentDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgent) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AuthenticatedSessionUserAgentDevice and assigns it to the Device field.
func (o *AuthenticatedSessionUserAgent) SetDevice(v AuthenticatedSessionUserAgentDevice) {
	o.Device = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgent) GetOs() AuthenticatedSessionUserAgentOs {
	if o == nil || o.Os == nil {
		var ret AuthenticatedSessionUserAgentOs
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgent) GetOsOk() (*AuthenticatedSessionUserAgentOs, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgent) HasOs() bool {
	if o != nil && o.Os != nil {
		return true
	}

	return false
}

// SetOs gets a reference to the given AuthenticatedSessionUserAgentOs and assigns it to the Os field.
func (o *AuthenticatedSessionUserAgent) SetOs(v AuthenticatedSessionUserAgentOs) {
	o.Os = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgent) GetUserAgent() AuthenticatedSessionUserAgentUserAgent {
	if o == nil || o.UserAgent == nil {
		var ret AuthenticatedSessionUserAgentUserAgent
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgent) GetUserAgentOk() (*AuthenticatedSessionUserAgentUserAgent, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgent) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given AuthenticatedSessionUserAgentUserAgent and assigns it to the UserAgent field.
func (o *AuthenticatedSessionUserAgent) SetUserAgent(v AuthenticatedSessionUserAgentUserAgent) {
	o.UserAgent = &v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgent) GetString() string {
	if o == nil || o.String == nil {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgent) GetStringOk() (*string, bool) {
	if o == nil || o.String == nil {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgent) HasString() bool {
	if o != nil && o.String != nil {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *AuthenticatedSessionUserAgent) SetString(v string) {
	o.String = &v
}

func (o AuthenticatedSessionUserAgent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Os != nil {
		toSerialize["os"] = o.Os
	}
	if o.UserAgent != nil {
		toSerialize["user_agent"] = o.UserAgent
	}
	if o.String != nil {
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
