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

// GroupLDAPSourceConnectionRequest Group Source Connection
type GroupLDAPSourceConnectionRequest struct {
	Group      string `json:"group"`
	Source     string `json:"source"`
	Identifier string `json:"identifier"`
}

// NewGroupLDAPSourceConnectionRequest instantiates a new GroupLDAPSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupLDAPSourceConnectionRequest(group string, source string, identifier string) *GroupLDAPSourceConnectionRequest {
	this := GroupLDAPSourceConnectionRequest{}
	this.Group = group
	this.Source = source
	this.Identifier = identifier
	return &this
}

// NewGroupLDAPSourceConnectionRequestWithDefaults instantiates a new GroupLDAPSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupLDAPSourceConnectionRequestWithDefaults() *GroupLDAPSourceConnectionRequest {
	this := GroupLDAPSourceConnectionRequest{}
	return &this
}

// GetGroup returns the Group field value
func (o *GroupLDAPSourceConnectionRequest) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *GroupLDAPSourceConnectionRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *GroupLDAPSourceConnectionRequest) SetGroup(v string) {
	o.Group = v
}

// GetSource returns the Source field value
func (o *GroupLDAPSourceConnectionRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GroupLDAPSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GroupLDAPSourceConnectionRequest) SetSource(v string) {
	o.Source = v
}

// GetIdentifier returns the Identifier field value
func (o *GroupLDAPSourceConnectionRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GroupLDAPSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GroupLDAPSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = v
}

func (o GroupLDAPSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	return json.Marshal(toSerialize)
}

type NullableGroupLDAPSourceConnectionRequest struct {
	value *GroupLDAPSourceConnectionRequest
	isSet bool
}

func (v NullableGroupLDAPSourceConnectionRequest) Get() *GroupLDAPSourceConnectionRequest {
	return v.value
}

func (v *NullableGroupLDAPSourceConnectionRequest) Set(val *GroupLDAPSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupLDAPSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupLDAPSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupLDAPSourceConnectionRequest(val *GroupLDAPSourceConnectionRequest) *NullableGroupLDAPSourceConnectionRequest {
	return &NullableGroupLDAPSourceConnectionRequest{value: val, isSet: true}
}

func (v NullableGroupLDAPSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupLDAPSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
