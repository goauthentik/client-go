/*
authentik

Making authentication simple.

API version: 2025.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EmailDeviceRequest Serializer for email authenticator devices
type EmailDeviceRequest struct {
	// The human-readable name of this device.
	Name string `json:"name"`
}

// NewEmailDeviceRequest instantiates a new EmailDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDeviceRequest(name string) *EmailDeviceRequest {
	this := EmailDeviceRequest{}
	this.Name = name
	return &this
}

// NewEmailDeviceRequestWithDefaults instantiates a new EmailDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDeviceRequestWithDefaults() *EmailDeviceRequest {
	this := EmailDeviceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *EmailDeviceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmailDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EmailDeviceRequest) SetName(v string) {
	o.Name = v
}

func (o EmailDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableEmailDeviceRequest struct {
	value *EmailDeviceRequest
	isSet bool
}

func (v NullableEmailDeviceRequest) Get() *EmailDeviceRequest {
	return v.value
}

func (v *NullableEmailDeviceRequest) Set(val *EmailDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDeviceRequest(val *EmailDeviceRequest) *NullableEmailDeviceRequest {
	return &NullableEmailDeviceRequest{value: val, isSet: true}
}

func (v NullableEmailDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
