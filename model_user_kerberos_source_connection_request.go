/*
authentik

Making authentication simple.

API version: 2024.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserKerberosSourceConnectionRequest Kerberos Source Serializer
type UserKerberosSourceConnectionRequest struct {
	User       int32  `json:"user"`
	Source     string `json:"source"`
	Identifier string `json:"identifier"`
}

// NewUserKerberosSourceConnectionRequest instantiates a new UserKerberosSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserKerberosSourceConnectionRequest(user int32, source string, identifier string) *UserKerberosSourceConnectionRequest {
	this := UserKerberosSourceConnectionRequest{}
	this.User = user
	this.Source = source
	this.Identifier = identifier
	return &this
}

// NewUserKerberosSourceConnectionRequestWithDefaults instantiates a new UserKerberosSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserKerberosSourceConnectionRequestWithDefaults() *UserKerberosSourceConnectionRequest {
	this := UserKerberosSourceConnectionRequest{}
	return &this
}

// GetUser returns the User field value
func (o *UserKerberosSourceConnectionRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserKerberosSourceConnectionRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserKerberosSourceConnectionRequest) SetUser(v int32) {
	o.User = v
}

// GetSource returns the Source field value
func (o *UserKerberosSourceConnectionRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UserKerberosSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UserKerberosSourceConnectionRequest) SetSource(v string) {
	o.Source = v
}

// GetIdentifier returns the Identifier field value
func (o *UserKerberosSourceConnectionRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *UserKerberosSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *UserKerberosSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = v
}

func (o UserKerberosSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullableUserKerberosSourceConnectionRequest struct {
	value *UserKerberosSourceConnectionRequest
	isSet bool
}

func (v NullableUserKerberosSourceConnectionRequest) Get() *UserKerberosSourceConnectionRequest {
	return v.value
}

func (v *NullableUserKerberosSourceConnectionRequest) Set(val *UserKerberosSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserKerberosSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserKerberosSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserKerberosSourceConnectionRequest(val *UserKerberosSourceConnectionRequest) *NullableUserKerberosSourceConnectionRequest {
	return &NullableUserKerberosSourceConnectionRequest{value: val, isSet: true}
}

func (v NullableUserKerberosSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserKerberosSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
