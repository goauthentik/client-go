/*
authentik

Making authentication simple.

API version: 2024.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GroupOAuthSourceConnectionRequest OAuth Group-Source connection Serializer
type GroupOAuthSourceConnectionRequest struct {
	Group      string `json:"group"`
	Source     string `json:"source"`
	Identifier string `json:"identifier"`
}

// NewGroupOAuthSourceConnectionRequest instantiates a new GroupOAuthSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupOAuthSourceConnectionRequest(group string, source string, identifier string) *GroupOAuthSourceConnectionRequest {
	this := GroupOAuthSourceConnectionRequest{}
	this.Group = group
	this.Source = source
	this.Identifier = identifier
	return &this
}

// NewGroupOAuthSourceConnectionRequestWithDefaults instantiates a new GroupOAuthSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupOAuthSourceConnectionRequestWithDefaults() *GroupOAuthSourceConnectionRequest {
	this := GroupOAuthSourceConnectionRequest{}
	return &this
}

// GetGroup returns the Group field value
func (o *GroupOAuthSourceConnectionRequest) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *GroupOAuthSourceConnectionRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *GroupOAuthSourceConnectionRequest) SetGroup(v string) {
	o.Group = v
}

// GetSource returns the Source field value
func (o *GroupOAuthSourceConnectionRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GroupOAuthSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GroupOAuthSourceConnectionRequest) SetSource(v string) {
	o.Source = v
}

// GetIdentifier returns the Identifier field value
func (o *GroupOAuthSourceConnectionRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GroupOAuthSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GroupOAuthSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = v
}

func (o GroupOAuthSourceConnectionRequest) MarshalJSON() ([]byte, error) {
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

type NullableGroupOAuthSourceConnectionRequest struct {
	value *GroupOAuthSourceConnectionRequest
	isSet bool
}

func (v NullableGroupOAuthSourceConnectionRequest) Get() *GroupOAuthSourceConnectionRequest {
	return v.value
}

func (v *NullableGroupOAuthSourceConnectionRequest) Set(val *GroupOAuthSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupOAuthSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupOAuthSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupOAuthSourceConnectionRequest(val *GroupOAuthSourceConnectionRequest) *NullableGroupOAuthSourceConnectionRequest {
	return &NullableGroupOAuthSourceConnectionRequest{value: val, isSet: true}
}

func (v NullableGroupOAuthSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupOAuthSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
