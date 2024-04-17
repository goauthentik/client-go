/*
authentik

Making authentication simple.

API version: 2024.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// StaticDevice Serializer for static authenticator devices
type StaticDevice struct {
	// The human-readable name of this device.
	Name     string              `json:"name"`
	TokenSet []StaticDeviceToken `json:"token_set"`
	Pk       int32               `json:"pk"`
}

// NewStaticDevice instantiates a new StaticDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticDevice(name string, tokenSet []StaticDeviceToken, pk int32) *StaticDevice {
	this := StaticDevice{}
	this.Name = name
	this.TokenSet = tokenSet
	this.Pk = pk
	return &this
}

// NewStaticDeviceWithDefaults instantiates a new StaticDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticDeviceWithDefaults() *StaticDevice {
	this := StaticDevice{}
	return &this
}

// GetName returns the Name field value
func (o *StaticDevice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StaticDevice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StaticDevice) SetName(v string) {
	o.Name = v
}

// GetTokenSet returns the TokenSet field value
func (o *StaticDevice) GetTokenSet() []StaticDeviceToken {
	if o == nil {
		var ret []StaticDeviceToken
		return ret
	}

	return o.TokenSet
}

// GetTokenSetOk returns a tuple with the TokenSet field value
// and a boolean to check if the value has been set.
func (o *StaticDevice) GetTokenSetOk() ([]StaticDeviceToken, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenSet, true
}

// SetTokenSet sets field value
func (o *StaticDevice) SetTokenSet(v []StaticDeviceToken) {
	o.TokenSet = v
}

// GetPk returns the Pk field value
func (o *StaticDevice) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *StaticDevice) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *StaticDevice) SetPk(v int32) {
	o.Pk = v
}

func (o StaticDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["token_set"] = o.TokenSet
	}
	if true {
		toSerialize["pk"] = o.Pk
	}
	return json.Marshal(toSerialize)
}

type NullableStaticDevice struct {
	value *StaticDevice
	isSet bool
}

func (v NullableStaticDevice) Get() *StaticDevice {
	return v.value
}

func (v *NullableStaticDevice) Set(val *StaticDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticDevice(val *StaticDevice) *NullableStaticDevice {
	return &NullableStaticDevice{value: val, isSet: true}
}

func (v NullableStaticDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
