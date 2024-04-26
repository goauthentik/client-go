/*
authentik

Making authentication simple.

API version: 2024.4.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TOTPDevice Serializer for totp authenticator devices
type TOTPDevice struct {
	// The human-readable name of this device.
	Name string `json:"name"`
	Pk   int32  `json:"pk"`
}

// NewTOTPDevice instantiates a new TOTPDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTOTPDevice(name string, pk int32) *TOTPDevice {
	this := TOTPDevice{}
	this.Name = name
	this.Pk = pk
	return &this
}

// NewTOTPDeviceWithDefaults instantiates a new TOTPDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTOTPDeviceWithDefaults() *TOTPDevice {
	this := TOTPDevice{}
	return &this
}

// GetName returns the Name field value
func (o *TOTPDevice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TOTPDevice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TOTPDevice) SetName(v string) {
	o.Name = v
}

// GetPk returns the Pk field value
func (o *TOTPDevice) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *TOTPDevice) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *TOTPDevice) SetPk(v int32) {
	o.Pk = v
}

func (o TOTPDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["pk"] = o.Pk
	}
	return json.Marshal(toSerialize)
}

type NullableTOTPDevice struct {
	value *TOTPDevice
	isSet bool
}

func (v NullableTOTPDevice) Get() *TOTPDevice {
	return v.value
}

func (v *NullableTOTPDevice) Set(val *TOTPDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableTOTPDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableTOTPDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTOTPDevice(val *TOTPDevice) *NullableTOTPDevice {
	return &NullableTOTPDevice{value: val, isSet: true}
}

func (v NullableTOTPDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTOTPDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
