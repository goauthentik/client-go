/*
authentik

Making authentication simple.

API version: 2022.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// UserSourceConnection OAuth Source Serializer
type UserSourceConnection struct {
	Pk      int32                      `json:"pk"`
	User    int32                      `json:"user"`
	Source  PlexSourceConnectionSource `json:"source"`
	Created time.Time                  `json:"created"`
}

// NewUserSourceConnection instantiates a new UserSourceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSourceConnection(pk int32, user int32, source PlexSourceConnectionSource, created time.Time) *UserSourceConnection {
	this := UserSourceConnection{}
	this.Pk = pk
	this.User = user
	this.Source = source
	this.Created = created
	return &this
}

// NewUserSourceConnectionWithDefaults instantiates a new UserSourceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSourceConnectionWithDefaults() *UserSourceConnection {
	this := UserSourceConnection{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserSourceConnection) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserSourceConnection) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserSourceConnection) SetPk(v int32) {
	o.Pk = v
}

// GetUser returns the User field value
func (o *UserSourceConnection) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserSourceConnection) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserSourceConnection) SetUser(v int32) {
	o.User = v
}

// GetSource returns the Source field value
func (o *UserSourceConnection) GetSource() PlexSourceConnectionSource {
	if o == nil {
		var ret PlexSourceConnectionSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UserSourceConnection) GetSourceOk() (*PlexSourceConnectionSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UserSourceConnection) SetSource(v PlexSourceConnectionSource) {
	o.Source = v
}

// GetCreated returns the Created field value
func (o *UserSourceConnection) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *UserSourceConnection) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *UserSourceConnection) SetCreated(v time.Time) {
	o.Created = v
}

func (o UserSourceConnection) MarshalJSON() ([]byte, error) {
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
		toSerialize["created"] = o.Created
	}
	return json.Marshal(toSerialize)
}

type NullableUserSourceConnection struct {
	value *UserSourceConnection
	isSet bool
}

func (v NullableUserSourceConnection) Get() *UserSourceConnection {
	return v.value
}

func (v *NullableUserSourceConnection) Set(val *UserSourceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSourceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSourceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSourceConnection(val *UserSourceConnection) *NullableUserSourceConnection {
	return &NullableUserSourceConnection{value: val, isSet: true}
}

func (v NullableUserSourceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSourceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
