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

// PatchedEndpointDeviceRequest Serializer for Endpoint authenticator devices
type PatchedEndpointDeviceRequest struct {
	Pk *string `json:"pk,omitempty"`
	// The human-readable name of this device.
	Name *string `json:"name,omitempty"`
}

// NewPatchedEndpointDeviceRequest instantiates a new PatchedEndpointDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEndpointDeviceRequest() *PatchedEndpointDeviceRequest {
	this := PatchedEndpointDeviceRequest{}
	return &this
}

// NewPatchedEndpointDeviceRequestWithDefaults instantiates a new PatchedEndpointDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEndpointDeviceRequestWithDefaults() *PatchedEndpointDeviceRequest {
	this := PatchedEndpointDeviceRequest{}
	return &this
}

// GetPk returns the Pk field value if set, zero value otherwise.
func (o *PatchedEndpointDeviceRequest) GetPk() string {
	if o == nil || o.Pk == nil {
		var ret string
		return ret
	}
	return *o.Pk
}

// GetPkOk returns a tuple with the Pk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEndpointDeviceRequest) GetPkOk() (*string, bool) {
	if o == nil || o.Pk == nil {
		return nil, false
	}
	return o.Pk, true
}

// HasPk returns a boolean if a field has been set.
func (o *PatchedEndpointDeviceRequest) HasPk() bool {
	if o != nil && o.Pk != nil {
		return true
	}

	return false
}

// SetPk gets a reference to the given string and assigns it to the Pk field.
func (o *PatchedEndpointDeviceRequest) SetPk(v string) {
	o.Pk = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedEndpointDeviceRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEndpointDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedEndpointDeviceRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedEndpointDeviceRequest) SetName(v string) {
	o.Name = &v
}

func (o PatchedEndpointDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pk != nil {
		toSerialize["pk"] = o.Pk
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedEndpointDeviceRequest struct {
	value *PatchedEndpointDeviceRequest
	isSet bool
}

func (v NullablePatchedEndpointDeviceRequest) Get() *PatchedEndpointDeviceRequest {
	return v.value
}

func (v *NullablePatchedEndpointDeviceRequest) Set(val *PatchedEndpointDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEndpointDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEndpointDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEndpointDeviceRequest(val *PatchedEndpointDeviceRequest) *NullablePatchedEndpointDeviceRequest {
	return &NullablePatchedEndpointDeviceRequest{value: val, isSet: true}
}

func (v NullablePatchedEndpointDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEndpointDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
