/*
authentik

Making authentication simple.

API version: 2024.8.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMSourceGroup SCIMSourceGroup Serializer
type SCIMSourceGroup struct {
	Id         string      `json:"id"`
	Group      string      `json:"group"`
	GroupObj   UserGroup   `json:"group_obj"`
	Source     string      `json:"source"`
	Attributes interface{} `json:"attributes,omitempty"`
}

// NewSCIMSourceGroup instantiates a new SCIMSourceGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMSourceGroup(id string, group string, groupObj UserGroup, source string) *SCIMSourceGroup {
	this := SCIMSourceGroup{}
	this.Id = id
	this.Group = group
	this.GroupObj = groupObj
	this.Source = source
	return &this
}

// NewSCIMSourceGroupWithDefaults instantiates a new SCIMSourceGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMSourceGroupWithDefaults() *SCIMSourceGroup {
	this := SCIMSourceGroup{}
	return &this
}

// GetId returns the Id field value
func (o *SCIMSourceGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SCIMSourceGroup) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *SCIMSourceGroup) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceGroup) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SCIMSourceGroup) SetGroup(v string) {
	o.Group = v
}

// GetGroupObj returns the GroupObj field value
func (o *SCIMSourceGroup) GetGroupObj() UserGroup {
	if o == nil {
		var ret UserGroup
		return ret
	}

	return o.GroupObj
}

// GetGroupObjOk returns a tuple with the GroupObj field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceGroup) GetGroupObjOk() (*UserGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupObj, true
}

// SetGroupObj sets field value
func (o *SCIMSourceGroup) SetGroupObj(v UserGroup) {
	o.GroupObj = v
}

// GetSource returns the Source field value
func (o *SCIMSourceGroup) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceGroup) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SCIMSourceGroup) SetSource(v string) {
	o.Source = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SCIMSourceGroup) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SCIMSourceGroup) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SCIMSourceGroup) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *SCIMSourceGroup) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o SCIMSourceGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["group_obj"] = o.GroupObj
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMSourceGroup struct {
	value *SCIMSourceGroup
	isSet bool
}

func (v NullableSCIMSourceGroup) Get() *SCIMSourceGroup {
	return v.value
}

func (v *NullableSCIMSourceGroup) Set(val *SCIMSourceGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMSourceGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMSourceGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMSourceGroup(val *SCIMSourceGroup) *NullableSCIMSourceGroup {
	return &NullableSCIMSourceGroup{value: val, isSet: true}
}

func (v NullableSCIMSourceGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMSourceGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
