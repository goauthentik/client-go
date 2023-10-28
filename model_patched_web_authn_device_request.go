/*
authentik

Making authentication simple.

API version: 2023.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedWebAuthnDeviceRequest Serializer for WebAuthn authenticator devices
type PatchedWebAuthnDeviceRequest struct {
	Name *string `json:"name,omitempty"`
}

// NewPatchedWebAuthnDeviceRequest instantiates a new PatchedWebAuthnDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWebAuthnDeviceRequest() *PatchedWebAuthnDeviceRequest {
	this := PatchedWebAuthnDeviceRequest{}
	return &this
}

// NewPatchedWebAuthnDeviceRequestWithDefaults instantiates a new PatchedWebAuthnDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWebAuthnDeviceRequestWithDefaults() *PatchedWebAuthnDeviceRequest {
	this := PatchedWebAuthnDeviceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWebAuthnDeviceRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWebAuthnDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWebAuthnDeviceRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWebAuthnDeviceRequest) SetName(v string) {
	o.Name = &v
}

func (o PatchedWebAuthnDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedWebAuthnDeviceRequest struct {
	value *PatchedWebAuthnDeviceRequest
	isSet bool
}

func (v NullablePatchedWebAuthnDeviceRequest) Get() *PatchedWebAuthnDeviceRequest {
	return v.value
}

func (v *NullablePatchedWebAuthnDeviceRequest) Set(val *PatchedWebAuthnDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWebAuthnDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWebAuthnDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWebAuthnDeviceRequest(val *PatchedWebAuthnDeviceRequest) *NullablePatchedWebAuthnDeviceRequest {
	return &NullablePatchedWebAuthnDeviceRequest{value: val, isSet: true}
}

func (v NullablePatchedWebAuthnDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWebAuthnDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
