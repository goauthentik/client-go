/*
authentik

Making authentication simple.

API version: 2021.9.7
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedDuoDeviceRequest Serializer for Duo authenticator devices
type PatchedDuoDeviceRequest struct {
	// The human-readable name of this device.
	Name *string `json:"name,omitempty"`
}

// NewPatchedDuoDeviceRequest instantiates a new PatchedDuoDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDuoDeviceRequest() *PatchedDuoDeviceRequest {
	this := PatchedDuoDeviceRequest{}
	return &this
}

// NewPatchedDuoDeviceRequestWithDefaults instantiates a new PatchedDuoDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDuoDeviceRequestWithDefaults() *PatchedDuoDeviceRequest {
	this := PatchedDuoDeviceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedDuoDeviceRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDuoDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedDuoDeviceRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedDuoDeviceRequest) SetName(v string) {
	o.Name = &v
}

func (o PatchedDuoDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedDuoDeviceRequest struct {
	value *PatchedDuoDeviceRequest
	isSet bool
}

func (v NullablePatchedDuoDeviceRequest) Get() *PatchedDuoDeviceRequest {
	return v.value
}

func (v *NullablePatchedDuoDeviceRequest) Set(val *PatchedDuoDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDuoDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDuoDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDuoDeviceRequest(val *PatchedDuoDeviceRequest) *NullablePatchedDuoDeviceRequest {
	return &NullablePatchedDuoDeviceRequest{value: val, isSet: true}
}

func (v NullablePatchedDuoDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDuoDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
