/*
authentik

Making authentication simple.

API version: 2021.10.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserOAuthSourceConnectionRequest OAuth Source Serializer
type UserOAuthSourceConnectionRequest struct {
	Source     string `json:"source"`
	Identifier string `json:"identifier"`
}

// NewUserOAuthSourceConnectionRequest instantiates a new UserOAuthSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserOAuthSourceConnectionRequest(source string, identifier string) *UserOAuthSourceConnectionRequest {
	this := UserOAuthSourceConnectionRequest{}
	this.Source = source
	this.Identifier = identifier
	return &this
}

// NewUserOAuthSourceConnectionRequestWithDefaults instantiates a new UserOAuthSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserOAuthSourceConnectionRequestWithDefaults() *UserOAuthSourceConnectionRequest {
	this := UserOAuthSourceConnectionRequest{}
	return &this
}

// GetSource returns the Source field value
func (o *UserOAuthSourceConnectionRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UserOAuthSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UserOAuthSourceConnectionRequest) SetSource(v string) {
	o.Source = v
}

// GetIdentifier returns the Identifier field value
func (o *UserOAuthSourceConnectionRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *UserOAuthSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *UserOAuthSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = v
}

func (o UserOAuthSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullableUserOAuthSourceConnectionRequest struct {
	value *UserOAuthSourceConnectionRequest
	isSet bool
}

func (v NullableUserOAuthSourceConnectionRequest) Get() *UserOAuthSourceConnectionRequest {
	return v.value
}

func (v *NullableUserOAuthSourceConnectionRequest) Set(val *UserOAuthSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserOAuthSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserOAuthSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserOAuthSourceConnectionRequest(val *UserOAuthSourceConnectionRequest) *NullableUserOAuthSourceConnectionRequest {
	return &NullableUserOAuthSourceConnectionRequest{value: val, isSet: true}
}

func (v NullableUserOAuthSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserOAuthSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
