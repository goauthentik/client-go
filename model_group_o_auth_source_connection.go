/*
authentik

Making authentication simple.

API version: 2024.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// GroupOAuthSourceConnection OAuth Group-Source connection Serializer
type GroupOAuthSourceConnection struct {
	Pk         int32     `json:"pk"`
	Group      string    `json:"group"`
	Source     Source    `json:"source"`
	Identifier string    `json:"identifier"`
	Created    time.Time `json:"created"`
}

// NewGroupOAuthSourceConnection instantiates a new GroupOAuthSourceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupOAuthSourceConnection(pk int32, group string, source Source, identifier string, created time.Time) *GroupOAuthSourceConnection {
	this := GroupOAuthSourceConnection{}
	this.Pk = pk
	this.Group = group
	this.Source = source
	this.Identifier = identifier
	this.Created = created
	return &this
}

// NewGroupOAuthSourceConnectionWithDefaults instantiates a new GroupOAuthSourceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupOAuthSourceConnectionWithDefaults() *GroupOAuthSourceConnection {
	this := GroupOAuthSourceConnection{}
	return &this
}

// GetPk returns the Pk field value
func (o *GroupOAuthSourceConnection) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *GroupOAuthSourceConnection) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *GroupOAuthSourceConnection) SetPk(v int32) {
	o.Pk = v
}

// GetGroup returns the Group field value
func (o *GroupOAuthSourceConnection) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *GroupOAuthSourceConnection) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *GroupOAuthSourceConnection) SetGroup(v string) {
	o.Group = v
}

// GetSource returns the Source field value
func (o *GroupOAuthSourceConnection) GetSource() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GroupOAuthSourceConnection) GetSourceOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GroupOAuthSourceConnection) SetSource(v Source) {
	o.Source = v
}

// GetIdentifier returns the Identifier field value
func (o *GroupOAuthSourceConnection) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GroupOAuthSourceConnection) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GroupOAuthSourceConnection) SetIdentifier(v string) {
	o.Identifier = v
}

// GetCreated returns the Created field value
func (o *GroupOAuthSourceConnection) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GroupOAuthSourceConnection) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *GroupOAuthSourceConnection) SetCreated(v time.Time) {
	o.Created = v
}

func (o GroupOAuthSourceConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["created"] = o.Created
	}
	return json.Marshal(toSerialize)
}

type NullableGroupOAuthSourceConnection struct {
	value *GroupOAuthSourceConnection
	isSet bool
}

func (v NullableGroupOAuthSourceConnection) Get() *GroupOAuthSourceConnection {
	return v.value
}

func (v *NullableGroupOAuthSourceConnection) Set(val *GroupOAuthSourceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupOAuthSourceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupOAuthSourceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupOAuthSourceConnection(val *GroupOAuthSourceConnection) *NullableGroupOAuthSourceConnection {
	return &NullableGroupOAuthSourceConnection{value: val, isSet: true}
}

func (v NullableGroupOAuthSourceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupOAuthSourceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
