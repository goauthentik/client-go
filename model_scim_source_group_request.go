/*
authentik

Making authentication simple.

API version: 2024.6.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMSourceGroupRequest SCIMSourceGroup Serializer
type SCIMSourceGroupRequest struct {
	Id         string      `json:"id"`
	Group      string      `json:"group"`
	Source     string      `json:"source"`
	Attributes interface{} `json:"attributes,omitempty"`
}

// NewSCIMSourceGroupRequest instantiates a new SCIMSourceGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMSourceGroupRequest(id string, group string, source string) *SCIMSourceGroupRequest {
	this := SCIMSourceGroupRequest{}
	this.Id = id
	this.Group = group
	this.Source = source
	return &this
}

// NewSCIMSourceGroupRequestWithDefaults instantiates a new SCIMSourceGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMSourceGroupRequestWithDefaults() *SCIMSourceGroupRequest {
	this := SCIMSourceGroupRequest{}
	return &this
}

// GetId returns the Id field value
func (o *SCIMSourceGroupRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceGroupRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SCIMSourceGroupRequest) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *SCIMSourceGroupRequest) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceGroupRequest) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SCIMSourceGroupRequest) SetGroup(v string) {
	o.Group = v
}

// GetSource returns the Source field value
func (o *SCIMSourceGroupRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceGroupRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SCIMSourceGroupRequest) SetSource(v string) {
	o.Source = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SCIMSourceGroupRequest) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SCIMSourceGroupRequest) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SCIMSourceGroupRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *SCIMSourceGroupRequest) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o SCIMSourceGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMSourceGroupRequest struct {
	value *SCIMSourceGroupRequest
	isSet bool
}

func (v NullableSCIMSourceGroupRequest) Get() *SCIMSourceGroupRequest {
	return v.value
}

func (v *NullableSCIMSourceGroupRequest) Set(val *SCIMSourceGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMSourceGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMSourceGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMSourceGroupRequest(val *SCIMSourceGroupRequest) *NullableSCIMSourceGroupRequest {
	return &NullableSCIMSourceGroupRequest{value: val, isSet: true}
}

func (v NullableSCIMSourceGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMSourceGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
