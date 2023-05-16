/*
authentik

Making authentication simple.

API version: 2023.5.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DuoDevice Serializer for Duo authenticator devices
type DuoDevice struct {
	Pk int32 `json:"pk"`
	// The human-readable name of this device.
	Name string `json:"name"`
}

// NewDuoDevice instantiates a new DuoDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDuoDevice(pk int32, name string) *DuoDevice {
	this := DuoDevice{}
	this.Pk = pk
	this.Name = name
	return &this
}

// NewDuoDeviceWithDefaults instantiates a new DuoDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDuoDeviceWithDefaults() *DuoDevice {
	this := DuoDevice{}
	return &this
}

// GetPk returns the Pk field value
func (o *DuoDevice) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *DuoDevice) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *DuoDevice) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *DuoDevice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DuoDevice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DuoDevice) SetName(v string) {
	o.Name = v
}

func (o DuoDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDuoDevice struct {
	value *DuoDevice
	isSet bool
}

func (v NullableDuoDevice) Get() *DuoDevice {
	return v.value
}

func (v *NullableDuoDevice) Set(val *DuoDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableDuoDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDuoDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuoDevice(val *DuoDevice) *NullableDuoDevice {
	return &NullableDuoDevice{value: val, isSet: true}
}

func (v NullableDuoDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuoDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
