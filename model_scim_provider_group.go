/*
authentik

Making authentication simple.

API version: 2024.6.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMProviderGroup SCIMProviderGroup Serializer
type SCIMProviderGroup struct {
	Id       string    `json:"id"`
	ScimId   string    `json:"scim_id"`
	Group    string    `json:"group"`
	GroupObj UserGroup `json:"group_obj"`
	Provider int32     `json:"provider"`
}

// NewSCIMProviderGroup instantiates a new SCIMProviderGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMProviderGroup(id string, scimId string, group string, groupObj UserGroup, provider int32) *SCIMProviderGroup {
	this := SCIMProviderGroup{}
	this.Id = id
	this.ScimId = scimId
	this.Group = group
	this.GroupObj = groupObj
	this.Provider = provider
	return &this
}

// NewSCIMProviderGroupWithDefaults instantiates a new SCIMProviderGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMProviderGroupWithDefaults() *SCIMProviderGroup {
	this := SCIMProviderGroup{}
	return &this
}

// GetId returns the Id field value
func (o *SCIMProviderGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SCIMProviderGroup) SetId(v string) {
	o.Id = v
}

// GetScimId returns the ScimId field value
func (o *SCIMProviderGroup) GetScimId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScimId
}

// GetScimIdOk returns a tuple with the ScimId field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderGroup) GetScimIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScimId, true
}

// SetScimId sets field value
func (o *SCIMProviderGroup) SetScimId(v string) {
	o.ScimId = v
}

// GetGroup returns the Group field value
func (o *SCIMProviderGroup) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderGroup) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *SCIMProviderGroup) SetGroup(v string) {
	o.Group = v
}

// GetGroupObj returns the GroupObj field value
func (o *SCIMProviderGroup) GetGroupObj() UserGroup {
	if o == nil {
		var ret UserGroup
		return ret
	}

	return o.GroupObj
}

// GetGroupObjOk returns a tuple with the GroupObj field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderGroup) GetGroupObjOk() (*UserGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupObj, true
}

// SetGroupObj sets field value
func (o *SCIMProviderGroup) SetGroupObj(v UserGroup) {
	o.GroupObj = v
}

// GetProvider returns the Provider field value
func (o *SCIMProviderGroup) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderGroup) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SCIMProviderGroup) SetProvider(v int32) {
	o.Provider = v
}

func (o SCIMProviderGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["scim_id"] = o.ScimId
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["group_obj"] = o.GroupObj
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMProviderGroup struct {
	value *SCIMProviderGroup
	isSet bool
}

func (v NullableSCIMProviderGroup) Get() *SCIMProviderGroup {
	return v.value
}

func (v *NullableSCIMProviderGroup) Set(val *SCIMProviderGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMProviderGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMProviderGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMProviderGroup(val *SCIMProviderGroup) *NullableSCIMProviderGroup {
	return &NullableSCIMProviderGroup{value: val, isSet: true}
}

func (v NullableSCIMProviderGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMProviderGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
