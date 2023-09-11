/*
authentik

Making authentication simple.

API version: 2023.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LDAPDebug struct for LDAPDebug
type LDAPDebug struct {
	User       []map[string]interface{} `json:"user"`
	Group      []map[string]interface{} `json:"group"`
	Membership []map[string]interface{} `json:"membership"`
}

// NewLDAPDebug instantiates a new LDAPDebug object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPDebug(user []map[string]interface{}, group []map[string]interface{}, membership []map[string]interface{}) *LDAPDebug {
	this := LDAPDebug{}
	this.User = user
	this.Group = group
	this.Membership = membership
	return &this
}

// NewLDAPDebugWithDefaults instantiates a new LDAPDebug object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPDebugWithDefaults() *LDAPDebug {
	this := LDAPDebug{}
	return &this
}

// GetUser returns the User field value
func (o *LDAPDebug) GetUser() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *LDAPDebug) GetUserOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.User, true
}

// SetUser sets field value
func (o *LDAPDebug) SetUser(v []map[string]interface{}) {
	o.User = v
}

// GetGroup returns the Group field value
func (o *LDAPDebug) GetGroup() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *LDAPDebug) GetGroupOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group, true
}

// SetGroup sets field value
func (o *LDAPDebug) SetGroup(v []map[string]interface{}) {
	o.Group = v
}

// GetMembership returns the Membership field value
func (o *LDAPDebug) GetMembership() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Membership
}

// GetMembershipOk returns a tuple with the Membership field value
// and a boolean to check if the value has been set.
func (o *LDAPDebug) GetMembershipOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Membership, true
}

// SetMembership sets field value
func (o *LDAPDebug) SetMembership(v []map[string]interface{}) {
	o.Membership = v
}

func (o LDAPDebug) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["membership"] = o.Membership
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPDebug struct {
	value *LDAPDebug
	isSet bool
}

func (v NullableLDAPDebug) Get() *LDAPDebug {
	return v.value
}

func (v *NullableLDAPDebug) Set(val *LDAPDebug) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPDebug) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPDebug) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPDebug(val *LDAPDebug) *NullableLDAPDebug {
	return &NullableLDAPDebug{value: val, isSet: true}
}

func (v NullableLDAPDebug) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPDebug) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
