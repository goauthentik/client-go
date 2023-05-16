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

// PatchedSMSDeviceRequest Serializer for sms authenticator devices
type PatchedSMSDeviceRequest struct {
	// The human-readable name of this device.
	Name *string `json:"name,omitempty"`
}

// NewPatchedSMSDeviceRequest instantiates a new PatchedSMSDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSMSDeviceRequest() *PatchedSMSDeviceRequest {
	this := PatchedSMSDeviceRequest{}
	return &this
}

// NewPatchedSMSDeviceRequestWithDefaults instantiates a new PatchedSMSDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSMSDeviceRequestWithDefaults() *PatchedSMSDeviceRequest {
	this := PatchedSMSDeviceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedSMSDeviceRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSMSDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedSMSDeviceRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedSMSDeviceRequest) SetName(v string) {
	o.Name = &v
}

func (o PatchedSMSDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedSMSDeviceRequest struct {
	value *PatchedSMSDeviceRequest
	isSet bool
}

func (v NullablePatchedSMSDeviceRequest) Get() *PatchedSMSDeviceRequest {
	return v.value
}

func (v *NullablePatchedSMSDeviceRequest) Set(val *PatchedSMSDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSMSDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSMSDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSMSDeviceRequest(val *PatchedSMSDeviceRequest) *NullablePatchedSMSDeviceRequest {
	return &NullablePatchedSMSDeviceRequest{value: val, isSet: true}
}

func (v NullablePatchedSMSDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSMSDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
