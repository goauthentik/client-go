/*
authentik

Making authentication simple.

API version: 2024.12.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserSourceConnectionRequest User source connection
type UserSourceConnectionRequest struct {
	User   int32  `json:"user"`
	Source string `json:"source"`
}

// NewUserSourceConnectionRequest instantiates a new UserSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSourceConnectionRequest(user int32, source string) *UserSourceConnectionRequest {
	this := UserSourceConnectionRequest{}
	this.User = user
	this.Source = source
	return &this
}

// NewUserSourceConnectionRequestWithDefaults instantiates a new UserSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSourceConnectionRequestWithDefaults() *UserSourceConnectionRequest {
	this := UserSourceConnectionRequest{}
	return &this
}

// GetUser returns the User field value
func (o *UserSourceConnectionRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserSourceConnectionRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserSourceConnectionRequest) SetUser(v int32) {
	o.User = v
}

// GetSource returns the Source field value
func (o *UserSourceConnectionRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UserSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UserSourceConnectionRequest) SetSource(v string) {
	o.Source = v
}

func (o UserSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableUserSourceConnectionRequest struct {
	value *UserSourceConnectionRequest
	isSet bool
}

func (v NullableUserSourceConnectionRequest) Get() *UserSourceConnectionRequest {
	return v.value
}

func (v *NullableUserSourceConnectionRequest) Set(val *UserSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSourceConnectionRequest(val *UserSourceConnectionRequest) *NullableUserSourceConnectionRequest {
	return &NullableUserSourceConnectionRequest{value: val, isSet: true}
}

func (v NullableUserSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
