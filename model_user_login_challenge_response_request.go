/*
authentik

Making authentication simple.

API version: 2024.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserLoginChallengeResponseRequest User login challenge
type UserLoginChallengeResponseRequest struct {
	Component  *string `json:"component,omitempty"`
	RememberMe bool    `json:"remember_me"`
}

// NewUserLoginChallengeResponseRequest instantiates a new UserLoginChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLoginChallengeResponseRequest(rememberMe bool) *UserLoginChallengeResponseRequest {
	this := UserLoginChallengeResponseRequest{}
	var component string = "ak-stage-user-login"
	this.Component = &component
	this.RememberMe = rememberMe
	return &this
}

// NewUserLoginChallengeResponseRequestWithDefaults instantiates a new UserLoginChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLoginChallengeResponseRequestWithDefaults() *UserLoginChallengeResponseRequest {
	this := UserLoginChallengeResponseRequest{}
	var component string = "ak-stage-user-login"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *UserLoginChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *UserLoginChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *UserLoginChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

// GetRememberMe returns the RememberMe field value
func (o *UserLoginChallengeResponseRequest) GetRememberMe() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RememberMe
}

// GetRememberMeOk returns a tuple with the RememberMe field value
// and a boolean to check if the value has been set.
func (o *UserLoginChallengeResponseRequest) GetRememberMeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RememberMe, true
}

// SetRememberMe sets field value
func (o *UserLoginChallengeResponseRequest) SetRememberMe(v bool) {
	o.RememberMe = v
}

func (o UserLoginChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["remember_me"] = o.RememberMe
	}
	return json.Marshal(toSerialize)
}

type NullableUserLoginChallengeResponseRequest struct {
	value *UserLoginChallengeResponseRequest
	isSet bool
}

func (v NullableUserLoginChallengeResponseRequest) Get() *UserLoginChallengeResponseRequest {
	return v.value
}

func (v *NullableUserLoginChallengeResponseRequest) Set(val *UserLoginChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLoginChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLoginChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLoginChallengeResponseRequest(val *UserLoginChallengeResponseRequest) *NullableUserLoginChallengeResponseRequest {
	return &NullableUserLoginChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableUserLoginChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLoginChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
