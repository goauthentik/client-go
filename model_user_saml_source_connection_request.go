/*
authentik

Making authentication simple.

API version: 2022.11.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserSAMLSourceConnectionRequest SAML Source Serializer
type UserSAMLSourceConnectionRequest struct {
	User       int32  `json:"user"`
	Identifier string `json:"identifier"`
}

// NewUserSAMLSourceConnectionRequest instantiates a new UserSAMLSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSAMLSourceConnectionRequest(user int32, identifier string) *UserSAMLSourceConnectionRequest {
	this := UserSAMLSourceConnectionRequest{}
	this.User = user
	this.Identifier = identifier
	return &this
}

// NewUserSAMLSourceConnectionRequestWithDefaults instantiates a new UserSAMLSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSAMLSourceConnectionRequestWithDefaults() *UserSAMLSourceConnectionRequest {
	this := UserSAMLSourceConnectionRequest{}
	return &this
}

// GetUser returns the User field value
func (o *UserSAMLSourceConnectionRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserSAMLSourceConnectionRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserSAMLSourceConnectionRequest) SetUser(v int32) {
	o.User = v
}

// GetIdentifier returns the Identifier field value
func (o *UserSAMLSourceConnectionRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *UserSAMLSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *UserSAMLSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = v
}

func (o UserSAMLSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullableUserSAMLSourceConnectionRequest struct {
	value *UserSAMLSourceConnectionRequest
	isSet bool
}

func (v NullableUserSAMLSourceConnectionRequest) Get() *UserSAMLSourceConnectionRequest {
	return v.value
}

func (v *NullableUserSAMLSourceConnectionRequest) Set(val *UserSAMLSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSAMLSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSAMLSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSAMLSourceConnectionRequest(val *UserSAMLSourceConnectionRequest) *NullableUserSAMLSourceConnectionRequest {
	return &NullableUserSAMLSourceConnectionRequest{value: val, isSet: true}
}

func (v NullableUserSAMLSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSAMLSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
