/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserServiceAccountResponse struct for UserServiceAccountResponse
type UserServiceAccountResponse struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

// NewUserServiceAccountResponse instantiates a new UserServiceAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserServiceAccountResponse(username string, token string) *UserServiceAccountResponse {
	this := UserServiceAccountResponse{}
	this.Username = username
	this.Token = token
	return &this
}

// NewUserServiceAccountResponseWithDefaults instantiates a new UserServiceAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserServiceAccountResponseWithDefaults() *UserServiceAccountResponse {
	this := UserServiceAccountResponse{}
	return &this
}

// GetUsername returns the Username field value
func (o *UserServiceAccountResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserServiceAccountResponse) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserServiceAccountResponse) SetUsername(v string) {
	o.Username = v
}

// GetToken returns the Token field value
func (o *UserServiceAccountResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UserServiceAccountResponse) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UserServiceAccountResponse) SetToken(v string) {
	o.Token = v
}

func (o UserServiceAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableUserServiceAccountResponse struct {
	value *UserServiceAccountResponse
	isSet bool
}

func (v NullableUserServiceAccountResponse) Get() *UserServiceAccountResponse {
	return v.value
}

func (v *NullableUserServiceAccountResponse) Set(val *UserServiceAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserServiceAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserServiceAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserServiceAccountResponse(val *UserServiceAccountResponse) *NullableUserServiceAccountResponse {
	return &NullableUserServiceAccountResponse{value: val, isSet: true}
}

func (v NullableUserServiceAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserServiceAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


