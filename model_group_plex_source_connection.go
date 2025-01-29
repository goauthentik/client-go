/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// GroupPlexSourceConnection Plex Group-Source connection Serializer
type GroupPlexSourceConnection struct {
	Pk         int32     `json:"pk"`
	Group      string    `json:"group"`
	Source     string    `json:"source"`
	SourceObj  Source    `json:"source_obj"`
	Identifier string    `json:"identifier"`
	Created    time.Time `json:"created"`
}

// NewGroupPlexSourceConnection instantiates a new GroupPlexSourceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPlexSourceConnection(pk int32, group string, source string, sourceObj Source, identifier string, created time.Time) *GroupPlexSourceConnection {
	this := GroupPlexSourceConnection{}
	this.Pk = pk
	this.Group = group
	this.Source = source
	this.SourceObj = sourceObj
	this.Identifier = identifier
	this.Created = created
	return &this
}

// NewGroupPlexSourceConnectionWithDefaults instantiates a new GroupPlexSourceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPlexSourceConnectionWithDefaults() *GroupPlexSourceConnection {
	this := GroupPlexSourceConnection{}
	return &this
}

// GetPk returns the Pk field value
func (o *GroupPlexSourceConnection) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *GroupPlexSourceConnection) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *GroupPlexSourceConnection) SetPk(v int32) {
	o.Pk = v
}

// GetGroup returns the Group field value
func (o *GroupPlexSourceConnection) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *GroupPlexSourceConnection) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *GroupPlexSourceConnection) SetGroup(v string) {
	o.Group = v
}

// GetSource returns the Source field value
func (o *GroupPlexSourceConnection) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GroupPlexSourceConnection) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GroupPlexSourceConnection) SetSource(v string) {
	o.Source = v
}

// GetSourceObj returns the SourceObj field value
func (o *GroupPlexSourceConnection) GetSourceObj() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.SourceObj
}

// GetSourceObjOk returns a tuple with the SourceObj field value
// and a boolean to check if the value has been set.
func (o *GroupPlexSourceConnection) GetSourceObjOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceObj, true
}

// SetSourceObj sets field value
func (o *GroupPlexSourceConnection) SetSourceObj(v Source) {
	o.SourceObj = v
}

// GetIdentifier returns the Identifier field value
func (o *GroupPlexSourceConnection) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GroupPlexSourceConnection) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GroupPlexSourceConnection) SetIdentifier(v string) {
	o.Identifier = v
}

// GetCreated returns the Created field value
func (o *GroupPlexSourceConnection) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GroupPlexSourceConnection) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *GroupPlexSourceConnection) SetCreated(v time.Time) {
	o.Created = v
}

func (o GroupPlexSourceConnection) MarshalJSON() ([]byte, error) {
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
		toSerialize["source_obj"] = o.SourceObj
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["created"] = o.Created
	}
	return json.Marshal(toSerialize)
}

type NullableGroupPlexSourceConnection struct {
	value *GroupPlexSourceConnection
	isSet bool
}

func (v NullableGroupPlexSourceConnection) Get() *GroupPlexSourceConnection {
	return v.value
}

func (v *NullableGroupPlexSourceConnection) Set(val *GroupPlexSourceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPlexSourceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPlexSourceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPlexSourceConnection(val *GroupPlexSourceConnection) *NullableGroupPlexSourceConnection {
	return &NullableGroupPlexSourceConnection{value: val, isSet: true}
}

func (v NullableGroupPlexSourceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPlexSourceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
