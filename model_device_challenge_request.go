/*
authentik

Making authentication simple.

API version: 2023.8.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DeviceChallengeRequest Single device challenge
type DeviceChallengeRequest struct {
	DeviceClass string                 `json:"device_class"`
	DeviceUid   string                 `json:"device_uid"`
	Challenge   map[string]interface{} `json:"challenge"`
}

// NewDeviceChallengeRequest instantiates a new DeviceChallengeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceChallengeRequest(deviceClass string, deviceUid string, challenge map[string]interface{}) *DeviceChallengeRequest {
	this := DeviceChallengeRequest{}
	this.DeviceClass = deviceClass
	this.DeviceUid = deviceUid
	this.Challenge = challenge
	return &this
}

// NewDeviceChallengeRequestWithDefaults instantiates a new DeviceChallengeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceChallengeRequestWithDefaults() *DeviceChallengeRequest {
	this := DeviceChallengeRequest{}
	return &this
}

// GetDeviceClass returns the DeviceClass field value
func (o *DeviceChallengeRequest) GetDeviceClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceClass
}

// GetDeviceClassOk returns a tuple with the DeviceClass field value
// and a boolean to check if the value has been set.
func (o *DeviceChallengeRequest) GetDeviceClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceClass, true
}

// SetDeviceClass sets field value
func (o *DeviceChallengeRequest) SetDeviceClass(v string) {
	o.DeviceClass = v
}

// GetDeviceUid returns the DeviceUid field value
func (o *DeviceChallengeRequest) GetDeviceUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceUid
}

// GetDeviceUidOk returns a tuple with the DeviceUid field value
// and a boolean to check if the value has been set.
func (o *DeviceChallengeRequest) GetDeviceUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceUid, true
}

// SetDeviceUid sets field value
func (o *DeviceChallengeRequest) SetDeviceUid(v string) {
	o.DeviceUid = v
}

// GetChallenge returns the Challenge field value
func (o *DeviceChallengeRequest) GetChallenge() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value
// and a boolean to check if the value has been set.
func (o *DeviceChallengeRequest) GetChallengeOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Challenge, true
}

// SetChallenge sets field value
func (o *DeviceChallengeRequest) SetChallenge(v map[string]interface{}) {
	o.Challenge = v
}

func (o DeviceChallengeRequest) MarshalJSON() ([]byte, error) {
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

type NullableDeviceChallengeRequest struct {
	value *DeviceChallengeRequest
	isSet bool
}

func (v NullableDeviceChallengeRequest) Get() *DeviceChallengeRequest {
	return v.value
}

func (v *NullableDeviceChallengeRequest) Set(val *DeviceChallengeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceChallengeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceChallengeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceChallengeRequest(val *DeviceChallengeRequest) *NullableDeviceChallengeRequest {
	return &NullableDeviceChallengeRequest{value: val, isSet: true}
}

func (v NullableDeviceChallengeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceChallengeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
