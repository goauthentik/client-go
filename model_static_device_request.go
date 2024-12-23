/*
authentik

Making authentication simple.

API version: 2024.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// StaticDeviceRequest Serializer for static authenticator devices
type StaticDeviceRequest struct {
	// The human-readable name of this device.
	Name string `json:"name"`
}

// NewStaticDeviceRequest instantiates a new StaticDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticDeviceRequest(name string) *StaticDeviceRequest {
	this := StaticDeviceRequest{}
	this.Name = name
	return &this
}

// NewStaticDeviceRequestWithDefaults instantiates a new StaticDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticDeviceRequestWithDefaults() *StaticDeviceRequest {
	this := StaticDeviceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *StaticDeviceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StaticDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StaticDeviceRequest) SetName(v string) {
	o.Name = v
}

func (o StaticDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableStaticDeviceRequest struct {
	value *StaticDeviceRequest
	isSet bool
}

func (v NullableStaticDeviceRequest) Get() *StaticDeviceRequest {
	return v.value
}

func (v *NullableStaticDeviceRequest) Set(val *StaticDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticDeviceRequest(val *StaticDeviceRequest) *NullableStaticDeviceRequest {
	return &NullableStaticDeviceRequest{value: val, isSet: true}
}

func (v NullableStaticDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
