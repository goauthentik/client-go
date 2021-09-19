/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DeviceChallenge Single device challenge
type DeviceChallenge struct {
	DeviceClass string `json:"device_class"`
	DeviceUid string `json:"device_uid"`
	Challenge map[string]interface{} `json:"challenge"`
}

// NewDeviceChallenge instantiates a new DeviceChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceChallenge(deviceClass string, deviceUid string, challenge map[string]interface{}) *DeviceChallenge {
	this := DeviceChallenge{}
	this.DeviceClass = deviceClass
	this.DeviceUid = deviceUid
	this.Challenge = challenge
	return &this
}

// NewDeviceChallengeWithDefaults instantiates a new DeviceChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceChallengeWithDefaults() *DeviceChallenge {
	this := DeviceChallenge{}
	return &this
}

// GetDeviceClass returns the DeviceClass field value
func (o *DeviceChallenge) GetDeviceClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceClass
}

// GetDeviceClassOk returns a tuple with the DeviceClass field value
// and a boolean to check if the value has been set.
func (o *DeviceChallenge) GetDeviceClassOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DeviceClass, true
}

// SetDeviceClass sets field value
func (o *DeviceChallenge) SetDeviceClass(v string) {
	o.DeviceClass = v
}

// GetDeviceUid returns the DeviceUid field value
func (o *DeviceChallenge) GetDeviceUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceUid
}

// GetDeviceUidOk returns a tuple with the DeviceUid field value
// and a boolean to check if the value has been set.
func (o *DeviceChallenge) GetDeviceUidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DeviceUid, true
}

// SetDeviceUid sets field value
func (o *DeviceChallenge) SetDeviceUid(v string) {
	o.DeviceUid = v
}

// GetChallenge returns the Challenge field value
func (o *DeviceChallenge) GetChallenge() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value
// and a boolean to check if the value has been set.
func (o *DeviceChallenge) GetChallengeOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Challenge, true
}

// SetChallenge sets field value
func (o *DeviceChallenge) SetChallenge(v map[string]interface{}) {
	o.Challenge = v
}

func (o DeviceChallenge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["device_class"] = o.DeviceClass
	}
	if true {
		toSerialize["device_uid"] = o.DeviceUid
	}
	if true {
		toSerialize["challenge"] = o.Challenge
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceChallenge struct {
	value *DeviceChallenge
	isSet bool
}

func (v NullableDeviceChallenge) Get() *DeviceChallenge {
	return v.value
}

func (v *NullableDeviceChallenge) Set(val *DeviceChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceChallenge(val *DeviceChallenge) *NullableDeviceChallenge {
	return &NullableDeviceChallenge{value: val, isSet: true}
}

func (v NullableDeviceChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


