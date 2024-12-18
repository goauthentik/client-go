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

// GroupKerberosSourceConnectionRequest OAuth Group-Source connection Serializer
type GroupKerberosSourceConnectionRequest struct {
	Group      string `json:"group"`
	Source     string `json:"source"`
	Identifier string `json:"identifier"`
}

// NewGroupKerberosSourceConnectionRequest instantiates a new GroupKerberosSourceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupKerberosSourceConnectionRequest(group string, source string, identifier string) *GroupKerberosSourceConnectionRequest {
	this := GroupKerberosSourceConnectionRequest{}
	this.Group = group
	this.Source = source
	this.Identifier = identifier
	return &this
}

// NewGroupKerberosSourceConnectionRequestWithDefaults instantiates a new GroupKerberosSourceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupKerberosSourceConnectionRequestWithDefaults() *GroupKerberosSourceConnectionRequest {
	this := GroupKerberosSourceConnectionRequest{}
	return &this
}

// GetGroup returns the Group field value
func (o *GroupKerberosSourceConnectionRequest) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *GroupKerberosSourceConnectionRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *GroupKerberosSourceConnectionRequest) SetGroup(v string) {
	o.Group = v
}

// GetSource returns the Source field value
func (o *GroupKerberosSourceConnectionRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GroupKerberosSourceConnectionRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GroupKerberosSourceConnectionRequest) SetSource(v string) {
	o.Source = v
}

// GetIdentifier returns the Identifier field value
func (o *GroupKerberosSourceConnectionRequest) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *GroupKerberosSourceConnectionRequest) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *GroupKerberosSourceConnectionRequest) SetIdentifier(v string) {
	o.Identifier = v
}

func (o GroupKerberosSourceConnectionRequest) MarshalJSON() ([]byte, error) {
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

type NullableGroupKerberosSourceConnectionRequest struct {
	value *GroupKerberosSourceConnectionRequest
	isSet bool
}

func (v NullableGroupKerberosSourceConnectionRequest) Get() *GroupKerberosSourceConnectionRequest {
	return v.value
}

func (v *NullableGroupKerberosSourceConnectionRequest) Set(val *GroupKerberosSourceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupKerberosSourceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupKerberosSourceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupKerberosSourceConnectionRequest(val *GroupKerberosSourceConnectionRequest) *NullableGroupKerberosSourceConnectionRequest {
	return &NullableGroupKerberosSourceConnectionRequest{value: val, isSet: true}
}

func (v NullableGroupKerberosSourceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupKerberosSourceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
