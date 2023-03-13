/*
authentik

Making authentication simple.

API version: 2023.3.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorDuoStageDeviceImportResponse struct for AuthenticatorDuoStageDeviceImportResponse
type AuthenticatorDuoStageDeviceImportResponse struct {
	Count int32  `json:"count"`
	Error string `json:"error"`
}

// NewAuthenticatorDuoStageDeviceImportResponse instantiates a new AuthenticatorDuoStageDeviceImportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorDuoStageDeviceImportResponse(count int32, error_ string) *AuthenticatorDuoStageDeviceImportResponse {
	this := AuthenticatorDuoStageDeviceImportResponse{}
	this.Count = count
	this.Error = error_
	return &this
}

// NewAuthenticatorDuoStageDeviceImportResponseWithDefaults instantiates a new AuthenticatorDuoStageDeviceImportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorDuoStageDeviceImportResponseWithDefaults() *AuthenticatorDuoStageDeviceImportResponse {
	this := AuthenticatorDuoStageDeviceImportResponse{}
	return &this
}

// GetCount returns the Count field value
func (o *AuthenticatorDuoStageDeviceImportResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoStageDeviceImportResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *AuthenticatorDuoStageDeviceImportResponse) SetCount(v int32) {
	o.Count = v
}

// GetError returns the Error field value
func (o *AuthenticatorDuoStageDeviceImportResponse) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoStageDeviceImportResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AuthenticatorDuoStageDeviceImportResponse) SetError(v string) {
	o.Error = v
}

func (o AuthenticatorDuoStageDeviceImportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorDuoStageDeviceImportResponse struct {
	value *AuthenticatorDuoStageDeviceImportResponse
	isSet bool
}

func (v NullableAuthenticatorDuoStageDeviceImportResponse) Get() *AuthenticatorDuoStageDeviceImportResponse {
	return v.value
}

func (v *NullableAuthenticatorDuoStageDeviceImportResponse) Set(val *AuthenticatorDuoStageDeviceImportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorDuoStageDeviceImportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorDuoStageDeviceImportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorDuoStageDeviceImportResponse(val *AuthenticatorDuoStageDeviceImportResponse) *NullableAuthenticatorDuoStageDeviceImportResponse {
	return &NullableAuthenticatorDuoStageDeviceImportResponse{value: val, isSet: true}
}

func (v NullableAuthenticatorDuoStageDeviceImportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorDuoStageDeviceImportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
