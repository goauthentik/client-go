/*
authentik

Making authentication simple.

API version: 2024.6.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GoogleWorkspaceProviderGroup GoogleWorkspaceProviderGroup Serializer
type GoogleWorkspaceProviderGroup struct {
	Id         string      `json:"id"`
	GoogleId   string      `json:"google_id"`
	Group      string      `json:"group"`
	GroupObj   UserGroup   `json:"group_obj"`
	Provider   int32       `json:"provider"`
	Attributes interface{} `json:"attributes"`
}

// NewGoogleWorkspaceProviderGroup instantiates a new GoogleWorkspaceProviderGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleWorkspaceProviderGroup(id string, googleId string, group string, groupObj UserGroup, provider int32, attributes interface{}) *GoogleWorkspaceProviderGroup {
	this := GoogleWorkspaceProviderGroup{}
	this.Id = id
	this.GoogleId = googleId
	this.Group = group
	this.GroupObj = groupObj
	this.Provider = provider
	this.Attributes = attributes
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

// GetGoogleId returns the GoogleId field value
func (o *GoogleWorkspaceProviderGroup) GetGoogleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoogleId
}

// GetGoogleIdOk returns a tuple with the GoogleId field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderGroup) GetGoogleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoogleId, true
}

// SetGoogleId sets field value
func (o *GoogleWorkspaceProviderGroup) SetGoogleId(v string) {
	o.GoogleId = v
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

// GetProvider returns the Provider field value
func (o *GoogleWorkspaceProviderGroup) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderGroup) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *GoogleWorkspaceProviderGroup) SetProvider(v int32) {
	o.Provider = v
}

// GetAttributes returns the Attributes field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GoogleWorkspaceProviderGroup) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleWorkspaceProviderGroup) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GoogleWorkspaceProviderGroup) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o GoogleWorkspaceProviderGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["google_id"] = o.GoogleId
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
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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
