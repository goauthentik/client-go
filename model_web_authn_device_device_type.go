/*
authentik

Making authentication simple.

API version: 2024.4.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// WebAuthnDeviceDeviceType struct for WebAuthnDeviceDeviceType
type WebAuthnDeviceDeviceType struct {
	Aaguid      string `json:"aaguid"`
	Description string `json:"description"`
}

// NewWebAuthnDeviceDeviceType instantiates a new WebAuthnDeviceDeviceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebAuthnDeviceDeviceType(aaguid string, description string) *WebAuthnDeviceDeviceType {
	this := WebAuthnDeviceDeviceType{}
	this.Aaguid = aaguid
	this.Description = description
	return &this
}

// NewWebAuthnDeviceDeviceTypeWithDefaults instantiates a new WebAuthnDeviceDeviceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebAuthnDeviceDeviceTypeWithDefaults() *WebAuthnDeviceDeviceType {
	this := WebAuthnDeviceDeviceType{}
	return &this
}

// GetAaguid returns the Aaguid field value
func (o *WebAuthnDeviceDeviceType) GetAaguid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aaguid
}

// GetAaguidOk returns a tuple with the Aaguid field value
// and a boolean to check if the value has been set.
func (o *WebAuthnDeviceDeviceType) GetAaguidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aaguid, true
}

// SetAaguid sets field value
func (o *WebAuthnDeviceDeviceType) SetAaguid(v string) {
	o.Aaguid = v
}

// GetDescription returns the Description field value
func (o *WebAuthnDeviceDeviceType) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *WebAuthnDeviceDeviceType) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *WebAuthnDeviceDeviceType) SetDescription(v string) {
	o.Description = v
}

func (o WebAuthnDeviceDeviceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["aaguid"] = o.Aaguid
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableWebAuthnDeviceDeviceType struct {
	value *WebAuthnDeviceDeviceType
	isSet bool
}

func (v NullableWebAuthnDeviceDeviceType) Get() *WebAuthnDeviceDeviceType {
	return v.value
}

func (v *NullableWebAuthnDeviceDeviceType) Set(val *WebAuthnDeviceDeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebAuthnDeviceDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebAuthnDeviceDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebAuthnDeviceDeviceType(val *WebAuthnDeviceDeviceType) *NullableWebAuthnDeviceDeviceType {
	return &NullableWebAuthnDeviceDeviceType{value: val, isSet: true}
}

func (v NullableWebAuthnDeviceDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebAuthnDeviceDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
