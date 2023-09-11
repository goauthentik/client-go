/*
authentik

Making authentication simple.

API version: 2023.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TOTPDeviceRequest Serializer for totp authenticator devices
type TOTPDeviceRequest struct {
	// The human-readable name of this device.
	Name string `json:"name"`
}

// NewTOTPDeviceRequest instantiates a new TOTPDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTOTPDeviceRequest(name string) *TOTPDeviceRequest {
	this := TOTPDeviceRequest{}
	this.Name = name
	return &this
}

// NewTOTPDeviceRequestWithDefaults instantiates a new TOTPDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTOTPDeviceRequestWithDefaults() *TOTPDeviceRequest {
	this := TOTPDeviceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *TOTPDeviceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TOTPDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TOTPDeviceRequest) SetName(v string) {
	o.Name = v
}

func (o TOTPDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableTOTPDeviceRequest struct {
	value *TOTPDeviceRequest
	isSet bool
}

func (v NullableTOTPDeviceRequest) Get() *TOTPDeviceRequest {
	return v.value
}

func (v *NullableTOTPDeviceRequest) Set(val *TOTPDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTOTPDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTOTPDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTOTPDeviceRequest(val *TOTPDeviceRequest) *NullableTOTPDeviceRequest {
	return &NullableTOTPDeviceRequest{value: val, isSet: true}
}

func (v NullableTOTPDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTOTPDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
