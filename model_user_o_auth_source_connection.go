/*
authentik

Making authentication simple.

API version: 2021.10.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserOAuthSourceConnection OAuth Source Serializer
type UserOAuthSourceConnection struct {
	Pk         int32  `json:"pk"`
	User       int32  `json:"user"`
	Source     string `json:"source"`
	Identifier string `json:"identifier"`
}

// NewUserOAuthSourceConnection instantiates a new UserOAuthSourceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserOAuthSourceConnection(pk int32, user int32, source string, identifier string) *UserOAuthSourceConnection {
	this := UserOAuthSourceConnection{}
	this.Pk = pk
	this.User = user
	this.Source = source
	this.Identifier = identifier
	return &this
}

// NewUserOAuthSourceConnectionWithDefaults instantiates a new UserOAuthSourceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserOAuthSourceConnectionWithDefaults() *UserOAuthSourceConnection {
	this := UserOAuthSourceConnection{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserOAuthSourceConnection) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserOAuthSourceConnection) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserOAuthSourceConnection) SetPk(v int32) {
	o.Pk = v
}

// GetUser returns the User field value
func (o *UserOAuthSourceConnection) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserOAuthSourceConnection) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserOAuthSourceConnection) SetUser(v int32) {
	o.User = v
}

// GetSource returns the Source field value
func (o *UserOAuthSourceConnection) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UserOAuthSourceConnection) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UserOAuthSourceConnection) SetSource(v string) {
	o.Source = v
}

// GetIdentifier returns the Identifier field value
func (o *UserOAuthSourceConnection) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *UserOAuthSourceConnection) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *UserOAuthSourceConnection) SetIdentifier(v string) {
	o.Identifier = v
}

func (o UserOAuthSourceConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
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

type NullableUserOAuthSourceConnection struct {
	value *UserOAuthSourceConnection
	isSet bool
}

func (v NullableUserOAuthSourceConnection) Get() *UserOAuthSourceConnection {
	return v.value
}

func (v *NullableUserOAuthSourceConnection) Set(val *UserOAuthSourceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableUserOAuthSourceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableUserOAuthSourceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserOAuthSourceConnection(val *UserOAuthSourceConnection) *NullableUserOAuthSourceConnection {
	return &NullableUserOAuthSourceConnection{value: val, isSet: true}
}

func (v NullableUserOAuthSourceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserOAuthSourceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
