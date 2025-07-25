/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SessionUser Response for the /user/me endpoint, returns the currently active user (as `user` property) and, if this user is being impersonated, the original user in the `original` property.
type SessionUser struct {
	User     UserSelf  `json:"user"`
	Original *UserSelf `json:"original,omitempty"`
}

// NewSessionUser instantiates a new SessionUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionUser(user UserSelf) *SessionUser {
	this := SessionUser{}
	this.User = user
	return &this
}

// NewSessionUserWithDefaults instantiates a new SessionUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionUserWithDefaults() *SessionUser {
	this := SessionUser{}
	return &this
}

// GetUser returns the User field value
func (o *SessionUser) GetUser() UserSelf {
	if o == nil {
		var ret UserSelf
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *SessionUser) GetUserOk() (*UserSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *SessionUser) SetUser(v UserSelf) {
	o.User = v
}

// GetOriginal returns the Original field value if set, zero value otherwise.
func (o *SessionUser) GetOriginal() UserSelf {
	if o == nil || o.Original == nil {
		var ret UserSelf
		return ret
	}
	return *o.Original
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionUser) GetOriginalOk() (*UserSelf, bool) {
	if o == nil || o.Original == nil {
		return nil, false
	}
	return o.Original, true
}

// HasOriginal returns a boolean if a field has been set.
func (o *SessionUser) HasOriginal() bool {
	if o != nil && o.Original != nil {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given UserSelf and assigns it to the Original field.
func (o *SessionUser) SetOriginal(v UserSelf) {
	o.Original = &v
}

func (o SessionUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	if o.Original != nil {
		toSerialize["original"] = o.Original
	}
	return json.Marshal(toSerialize)
}

type NullableSessionUser struct {
	value *SessionUser
	isSet bool
}

func (v NullableSessionUser) Get() *SessionUser {
	return v.value
}

func (v *NullableSessionUser) Set(val *SessionUser) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionUser) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionUser(val *SessionUser) *NullableSessionUser {
	return &NullableSessionUser{value: val, isSet: true}
}

func (v NullableSessionUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
