/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// InstallID struct for InstallID
type InstallID struct {
	InstallId string `json:"install_id"`
}

// NewInstallID instantiates a new InstallID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallID(installId string) *InstallID {
	this := InstallID{}
	this.InstallId = installId
	return &this
}

// NewInstallIDWithDefaults instantiates a new InstallID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallIDWithDefaults() *InstallID {
	this := InstallID{}
	return &this
}

// GetInstallId returns the InstallId field value
func (o *InstallID) GetInstallId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallId
}

// GetInstallIdOk returns a tuple with the InstallId field value
// and a boolean to check if the value has been set.
func (o *InstallID) GetInstallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallId, true
}

// SetInstallId sets field value
func (o *InstallID) SetInstallId(v string) {
	o.InstallId = v
}

func (o InstallID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["install_id"] = o.InstallId
	}
	return json.Marshal(toSerialize)
}

type NullableInstallID struct {
	value *InstallID
	isSet bool
}

func (v NullableInstallID) Get() *InstallID {
	return v.value
}

func (v *NullableInstallID) Set(val *InstallID) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallID) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallID(val *InstallID) *NullableInstallID {
	return &NullableInstallID{value: val, isSet: true}
}

func (v NullableInstallID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
