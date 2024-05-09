/*
authentik

Making authentication simple.

API version: 2024.4.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GoogleWorkspaceProviderGroup GoogleWorkspaceProviderGroup Serializer
type GoogleWorkspaceProviderGroup struct {
	Id       string    `json:"id"`
	Group    string    `json:"group"`
	GroupObj UserGroup `json:"group_obj"`
}

// NewGoogleWorkspaceProviderGroup instantiates a new GoogleWorkspaceProviderGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleWorkspaceProviderGroup(id string, group string, groupObj UserGroup) *GoogleWorkspaceProviderGroup {
	this := GoogleWorkspaceProviderGroup{}
	this.Id = id
	this.Group = group
	this.GroupObj = groupObj
	return &this
}

// NewGoogleWorkspaceProviderGroupWithDefaults instantiates a new GoogleWorkspaceProviderGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleWorkspaceProviderGroupWithDefaults() *GoogleWorkspaceProviderGroup {
	this := GoogleWorkspaceProviderGroup{}
	return &this
}

// GetId returns the Id field value
func (o *GoogleWorkspaceProviderGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GoogleWorkspaceProviderGroup) SetId(v string) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *GoogleWorkspaceProviderGroup) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderGroup) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *GoogleWorkspaceProviderGroup) SetGroup(v string) {
	o.Group = v
}

// GetGroupObj returns the GroupObj field value
func (o *GoogleWorkspaceProviderGroup) GetGroupObj() UserGroup {
	if o == nil {
		var ret UserGroup
		return ret
	}

	return o.GroupObj
}

// GetGroupObjOk returns a tuple with the GroupObj field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderGroup) GetGroupObjOk() (*UserGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupObj, true
}

// SetGroupObj sets field value
func (o *GoogleWorkspaceProviderGroup) SetGroupObj(v UserGroup) {
	o.GroupObj = v
}

func (o GoogleWorkspaceProviderGroup) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableGoogleWorkspaceProviderGroup struct {
	value *GoogleWorkspaceProviderGroup
	isSet bool
}

func (v NullableGoogleWorkspaceProviderGroup) Get() *GoogleWorkspaceProviderGroup {
	return v.value
}

func (v *NullableGoogleWorkspaceProviderGroup) Set(val *GoogleWorkspaceProviderGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleWorkspaceProviderGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleWorkspaceProviderGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleWorkspaceProviderGroup(val *GoogleWorkspaceProviderGroup) *NullableGoogleWorkspaceProviderGroup {
	return &NullableGoogleWorkspaceProviderGroup{value: val, isSet: true}
}

func (v NullableGoogleWorkspaceProviderGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleWorkspaceProviderGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}