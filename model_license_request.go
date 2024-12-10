/*
authentik

Making authentication simple.

API version: 2024.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LicenseRequest License Serializer
type LicenseRequest struct {
	Key string `json:"key"`
}

// NewLicenseRequest instantiates a new LicenseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseRequest(key string) *LicenseRequest {
	this := LicenseRequest{}
	this.Key = key
	return &this
}

// NewLicenseRequestWithDefaults instantiates a new LicenseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseRequestWithDefaults() *LicenseRequest {
	this := LicenseRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *LicenseRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *LicenseRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *LicenseRequest) SetKey(v string) {
	o.Key = v
}

func (o LicenseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseRequest struct {
	value *LicenseRequest
	isSet bool
}

func (v NullableLicenseRequest) Get() *LicenseRequest {
	return v.value
}

func (v *NullableLicenseRequest) Set(val *LicenseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseRequest(val *LicenseRequest) *NullableLicenseRequest {
	return &NullableLicenseRequest{value: val, isSet: true}
}

func (v NullableLicenseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
