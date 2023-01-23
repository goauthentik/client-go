/*
authentik

Making authentication simple.

API version: 2023.1.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedStaticDeviceRequest Serializer for static authenticator devices
type PatchedStaticDeviceRequest struct {
	// The human-readable name of this device.
	Name *string `json:"name,omitempty"`
}

// NewPatchedStaticDeviceRequest instantiates a new PatchedStaticDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedStaticDeviceRequest() *PatchedStaticDeviceRequest {
	this := PatchedStaticDeviceRequest{}
	return &this
}

// NewPatchedStaticDeviceRequestWithDefaults instantiates a new PatchedStaticDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedStaticDeviceRequestWithDefaults() *PatchedStaticDeviceRequest {
	this := PatchedStaticDeviceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedStaticDeviceRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedStaticDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedStaticDeviceRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedStaticDeviceRequest) SetName(v string) {
	o.Name = &v
}

func (o PatchedStaticDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedStaticDeviceRequest struct {
	value *PatchedStaticDeviceRequest
	isSet bool
}

func (v NullablePatchedStaticDeviceRequest) Get() *PatchedStaticDeviceRequest {
	return v.value
}

func (v *NullablePatchedStaticDeviceRequest) Set(val *PatchedStaticDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedStaticDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedStaticDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedStaticDeviceRequest(val *PatchedStaticDeviceRequest) *NullablePatchedStaticDeviceRequest {
	return &NullablePatchedStaticDeviceRequest{value: val, isSet: true}
}

func (v NullablePatchedStaticDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedStaticDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
