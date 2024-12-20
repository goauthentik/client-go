/*
authentik

Making authentication simple.

API version: 2024.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GroupPlexSourceConnectionRequest Plex Group-Source connection Serializer
type GroupPlexSourceConnectionRequest struct {
	Group      string `json:"group"`
	Source     string `json:"source"`
	Identifier string `json:"identifier"`
}

// NewGroupPlexSourceConnectionRequest instantiates a new GroupPlexSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPlexSourceConnectionRequest(group string, source string, identifier string) *GroupPlexSourceConnectionRequest {
	this := GroupPlexSourceConnectionRequest{}
	this.Group = group
	this.Source = source
	this.Identifier = identifier
	return &this
}

// NewGroupPlexSourceConnectionRequestWithDefaults instantiates a new GroupPlexSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPlexSourceConnectionRequestWithDefaults() *GroupPlexSourceConnectionRequest {
	this := GroupPlexSourceConnectionRequest{}
	return &this
}

// GetGroup returns the Group field value
func (o *GroupPlexSourceConnectionRequest) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *GroupPlexSourceConnectionRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *GroupPlexSourceConnectionRequest) SetGroup(v string) {
	o.Group = v
}

// GetSource returns the Source field value
func (o *GroupPlexSourceConnectionRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GroupPlexSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GroupPlexSourceConnectionRequest) SetSource(v string) {
	o.Source = v
}

// GetIdentifier returns the Identifier field value
func (o *GroupPlexSourceConnectionRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GroupPlexSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GroupPlexSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = v
}

func (o GroupPlexSourceConnectionRequest) MarshalJSON() ([]byte, error) {
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

type NullableGroupPlexSourceConnectionRequest struct {
	value *GroupPlexSourceConnectionRequest
	isSet bool
}

func (v NullableGroupPlexSourceConnectionRequest) Get() *GroupPlexSourceConnectionRequest {
	return v.value
}

func (v *NullableGroupPlexSourceConnectionRequest) Set(val *GroupPlexSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPlexSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPlexSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPlexSourceConnectionRequest(val *GroupPlexSourceConnectionRequest) *NullableGroupPlexSourceConnectionRequest {
	return &NullableGroupPlexSourceConnectionRequest{value: val, isSet: true}
}

func (v NullableGroupPlexSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPlexSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}