/*
authentik

Making authentication simple.

API version: 2024.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// WebAuthnDeviceTypeRequest WebAuthnDeviceType Serializer
type WebAuthnDeviceTypeRequest struct {
	Aaguid      string `json:"aaguid"`
	Description string `json:"description"`
}

// NewWebAuthnDeviceTypeRequest instantiates a new WebAuthnDeviceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebAuthnDeviceTypeRequest(aaguid string, description string) *WebAuthnDeviceTypeRequest {
	this := WebAuthnDeviceTypeRequest{}
	this.Aaguid = aaguid
	this.Description = description
	return &this
}

// NewWebAuthnDeviceTypeRequestWithDefaults instantiates a new WebAuthnDeviceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebAuthnDeviceTypeRequestWithDefaults() *WebAuthnDeviceTypeRequest {
	this := WebAuthnDeviceTypeRequest{}
	return &this
}

// GetAaguid returns the Aaguid field value
func (o *WebAuthnDeviceTypeRequest) GetAaguid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aaguid
}

// GetAaguidOk returns a tuple with the Aaguid field value
// and a boolean to check if the value has been set.
func (o *WebAuthnDeviceTypeRequest) GetAaguidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aaguid, true
}

// SetAaguid sets field value
func (o *WebAuthnDeviceTypeRequest) SetAaguid(v string) {
	o.Aaguid = v
}

// GetDescription returns the Description field value
func (o *WebAuthnDeviceTypeRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *WebAuthnDeviceTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *WebAuthnDeviceTypeRequest) SetDescription(v string) {
	o.Description = v
}

func (o WebAuthnDeviceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["aaguid"] = o.Aaguid
	}
	if true {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableWebAuthnDeviceTypeRequest struct {
	value *WebAuthnDeviceTypeRequest
	isSet bool
}

func (v NullableWebAuthnDeviceTypeRequest) Get() *WebAuthnDeviceTypeRequest {
	return v.value
}

func (v *NullableWebAuthnDeviceTypeRequest) Set(val *WebAuthnDeviceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebAuthnDeviceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebAuthnDeviceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebAuthnDeviceTypeRequest(val *WebAuthnDeviceTypeRequest) *NullableWebAuthnDeviceTypeRequest {
	return &NullableWebAuthnDeviceTypeRequest{value: val, isSet: true}
}

func (v NullableWebAuthnDeviceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebAuthnDeviceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
