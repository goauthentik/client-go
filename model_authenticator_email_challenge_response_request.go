/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorEmailChallengeResponseRequest Authenticator Email Challenge response, device is set by get_response_instance
type AuthenticatorEmailChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
	Code      *int32  `json:"code,omitempty"`
	Email     *string `json:"email,omitempty"`
}

// NewAuthenticatorEmailChallengeResponseRequest instantiates a new AuthenticatorEmailChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEmailChallengeResponseRequest() *AuthenticatorEmailChallengeResponseRequest {
	this := AuthenticatorEmailChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-email"
	this.Component = &component
	return &this
}

// NewAuthenticatorEmailChallengeResponseRequestWithDefaults instantiates a new AuthenticatorEmailChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEmailChallengeResponseRequestWithDefaults() *AuthenticatorEmailChallengeResponseRequest {
	this := AuthenticatorEmailChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-email"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorEmailChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorEmailChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorEmailChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AuthenticatorEmailChallengeResponseRequest) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailChallengeResponseRequest) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AuthenticatorEmailChallengeResponseRequest) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *AuthenticatorEmailChallengeResponseRequest) SetCode(v int32) {
	o.Code = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AuthenticatorEmailChallengeResponseRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailChallengeResponseRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AuthenticatorEmailChallengeResponseRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AuthenticatorEmailChallengeResponseRequest) SetEmail(v string) {
	o.Email = &v
}

func (o AuthenticatorEmailChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorEmailChallengeResponseRequest struct {
	value *AuthenticatorEmailChallengeResponseRequest
	isSet bool
}

func (v NullableAuthenticatorEmailChallengeResponseRequest) Get() *AuthenticatorEmailChallengeResponseRequest {
	return v.value
}

func (v *NullableAuthenticatorEmailChallengeResponseRequest) Set(val *AuthenticatorEmailChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEmailChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEmailChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEmailChallengeResponseRequest(val *AuthenticatorEmailChallengeResponseRequest) *NullableAuthenticatorEmailChallengeResponseRequest {
	return &NullableAuthenticatorEmailChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorEmailChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEmailChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
