/*
authentik

Making authentication simple.

API version: 2023.10.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedLicenseRequest License Serializer
type PatchedLicenseRequest struct {
	Key *string `json:"key,omitempty"`
}

// NewPatchedLicenseRequest instantiates a new PatchedLicenseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedLicenseRequest() *PatchedLicenseRequest {
	this := PatchedLicenseRequest{}
	return &this
}

// NewPatchedLicenseRequestWithDefaults instantiates a new PatchedLicenseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedLicenseRequestWithDefaults() *PatchedLicenseRequest {
	this := PatchedLicenseRequest{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PatchedLicenseRequest) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLicenseRequest) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PatchedLicenseRequest) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PatchedLicenseRequest) SetKey(v string) {
	o.Key = &v
}

func (o PatchedLicenseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedLicenseRequest struct {
	value *PatchedLicenseRequest
	isSet bool
}

func (v NullablePatchedLicenseRequest) Get() *PatchedLicenseRequest {
	return v.value
}

func (v *NullablePatchedLicenseRequest) Set(val *PatchedLicenseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedLicenseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedLicenseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedLicenseRequest(val *PatchedLicenseRequest) *NullablePatchedLicenseRequest {
	return &NullablePatchedLicenseRequest{value: val, isSet: true}
}

func (v NullablePatchedLicenseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedLicenseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
