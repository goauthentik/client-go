/*
authentik

Making authentication simple.

API version: 2024.6.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorDuoStageManualDeviceImportRequest struct for AuthenticatorDuoStageManualDeviceImportRequest
type AuthenticatorDuoStageManualDeviceImportRequest struct {
	DuoUserId string `json:"duo_user_id"`
	Username  string `json:"username"`
}

// NewAuthenticatorDuoStageManualDeviceImportRequest instantiates a new AuthenticatorDuoStageManualDeviceImportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorDuoStageManualDeviceImportRequest(duoUserId string, username string) *AuthenticatorDuoStageManualDeviceImportRequest {
	this := AuthenticatorDuoStageManualDeviceImportRequest{}
	this.DuoUserId = duoUserId
	this.Username = username
	return &this
}

// NewAuthenticatorDuoStageManualDeviceImportRequestWithDefaults instantiates a new AuthenticatorDuoStageManualDeviceImportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorDuoStageManualDeviceImportRequestWithDefaults() *AuthenticatorDuoStageManualDeviceImportRequest {
	this := AuthenticatorDuoStageManualDeviceImportRequest{}
	return &this
}

// GetDuoUserId returns the DuoUserId field value
func (o *AuthenticatorDuoStageManualDeviceImportRequest) GetDuoUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DuoUserId
}

// GetDuoUserIdOk returns a tuple with the DuoUserId field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoStageManualDeviceImportRequest) GetDuoUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DuoUserId, true
}

// SetDuoUserId sets field value
func (o *AuthenticatorDuoStageManualDeviceImportRequest) SetDuoUserId(v string) {
	o.DuoUserId = v
}

// GetUsername returns the Username field value
func (o *AuthenticatorDuoStageManualDeviceImportRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoStageManualDeviceImportRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthenticatorDuoStageManualDeviceImportRequest) SetUsername(v string) {
	o.Username = v
}

func (o AuthenticatorDuoStageManualDeviceImportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["duo_user_id"] = o.DuoUserId
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorDuoStageManualDeviceImportRequest struct {
	value *AuthenticatorDuoStageManualDeviceImportRequest
	isSet bool
}

func (v NullableAuthenticatorDuoStageManualDeviceImportRequest) Get() *AuthenticatorDuoStageManualDeviceImportRequest {
	return v.value
}

func (v *NullableAuthenticatorDuoStageManualDeviceImportRequest) Set(val *AuthenticatorDuoStageManualDeviceImportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorDuoStageManualDeviceImportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorDuoStageManualDeviceImportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorDuoStageManualDeviceImportRequest(val *AuthenticatorDuoStageManualDeviceImportRequest) *NullableAuthenticatorDuoStageManualDeviceImportRequest {
	return &NullableAuthenticatorDuoStageManualDeviceImportRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorDuoStageManualDeviceImportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorDuoStageManualDeviceImportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
