/*
authentik

Making authentication simple.

API version: 2022.7.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedTOTPDeviceRequest Serializer for totp authenticator devices
type PatchedTOTPDeviceRequest struct {
	// The human-readable name of this device.
	Name *string `json:"name,omitempty"`
}

// NewPatchedTOTPDeviceRequest instantiates a new PatchedTOTPDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedTOTPDeviceRequest() *PatchedTOTPDeviceRequest {
	this := PatchedTOTPDeviceRequest{}
	return &this
}

// NewPatchedTOTPDeviceRequestWithDefaults instantiates a new PatchedTOTPDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedTOTPDeviceRequestWithDefaults() *PatchedTOTPDeviceRequest {
	this := PatchedTOTPDeviceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedTOTPDeviceRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedTOTPDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedTOTPDeviceRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedTOTPDeviceRequest) SetName(v string) {
	o.Name = &v
}

func (o PatchedTOTPDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedTOTPDeviceRequest struct {
	value *PatchedTOTPDeviceRequest
	isSet bool
}

func (v NullablePatchedTOTPDeviceRequest) Get() *PatchedTOTPDeviceRequest {
	return v.value
}

func (v *NullablePatchedTOTPDeviceRequest) Set(val *PatchedTOTPDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedTOTPDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedTOTPDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedTOTPDeviceRequest(val *PatchedTOTPDeviceRequest) *NullablePatchedTOTPDeviceRequest {
	return &NullablePatchedTOTPDeviceRequest{value: val, isSet: true}
}

func (v NullablePatchedTOTPDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedTOTPDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
