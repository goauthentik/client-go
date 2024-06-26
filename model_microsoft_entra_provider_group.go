/*
authentik

Making authentication simple.

API version: 2024.6.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MicrosoftEntraProviderGroup MicrosoftEntraProviderGroup Serializer
type MicrosoftEntraProviderGroup struct {
	Id          string      `json:"id"`
	MicrosoftId string      `json:"microsoft_id"`
	Group       string      `json:"group"`
	GroupObj    UserGroup   `json:"group_obj"`
	Provider    int32       `json:"provider"`
	Attributes  interface{} `json:"attributes"`
}

// NewMicrosoftEntraProviderGroup instantiates a new MicrosoftEntraProviderGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftEntraProviderGroup(id string, microsoftId string, group string, groupObj UserGroup, provider int32, attributes interface{}) *MicrosoftEntraProviderGroup {
	this := MicrosoftEntraProviderGroup{}
	this.Id = id
	this.MicrosoftId = microsoftId
	this.Group = group
	this.GroupObj = groupObj
	this.Provider = provider
	this.Attributes = attributes
	return &this
}

// NewMicrosoftEntraProviderGroupWithDefaults instantiates a new MicrosoftEntraProviderGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftEntraProviderGroupWithDefaults() *MicrosoftEntraProviderGroup {
	this := MicrosoftEntraProviderGroup{}
	return &this
}

// GetId returns the Id field value
func (o *MicrosoftEntraProviderGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MicrosoftEntraProviderGroup) SetId(v string) {
	o.Id = v
}

// GetMicrosoftId returns the MicrosoftId field value
func (o *MicrosoftEntraProviderGroup) GetMicrosoftId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicrosoftId
}

// GetMicrosoftIdOk returns a tuple with the MicrosoftId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderGroup) GetMicrosoftIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicrosoftId, true
}

// SetMicrosoftId sets field value
func (o *MicrosoftEntraProviderGroup) SetMicrosoftId(v string) {
	o.MicrosoftId = v
}

// GetGroup returns the Group field value
func (o *MicrosoftEntraProviderGroup) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderGroup) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *MicrosoftEntraProviderGroup) SetGroup(v string) {
	o.Group = v
}

// GetGroupObj returns the GroupObj field value
func (o *MicrosoftEntraProviderGroup) GetGroupObj() UserGroup {
	if o == nil {
		var ret UserGroup
		return ret
	}

	return o.GroupObj
}

// GetGroupObjOk returns a tuple with the GroupObj field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderGroup) GetGroupObjOk() (*UserGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupObj, true
}

// SetGroupObj sets field value
func (o *MicrosoftEntraProviderGroup) SetGroupObj(v UserGroup) {
	o.GroupObj = v
}

// GetProvider returns the Provider field value
func (o *MicrosoftEntraProviderGroup) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderGroup) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *MicrosoftEntraProviderGroup) SetProvider(v int32) {
	o.Provider = v
}

// GetAttributes returns the Attributes field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *MicrosoftEntraProviderGroup) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftEntraProviderGroup) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *MicrosoftEntraProviderGroup) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o MicrosoftEntraProviderGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["microsoft_id"] = o.MicrosoftId
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

type NullableMicrosoftEntraProviderGroup struct {
	value *MicrosoftEntraProviderGroup
	isSet bool
}

func (v NullableMicrosoftEntraProviderGroup) Get() *MicrosoftEntraProviderGroup {
	return v.value
}

func (v *NullableMicrosoftEntraProviderGroup) Set(val *MicrosoftEntraProviderGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftEntraProviderGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftEntraProviderGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftEntraProviderGroup(val *MicrosoftEntraProviderGroup) *NullableMicrosoftEntraProviderGroup {
	return &NullableMicrosoftEntraProviderGroup{value: val, isSet: true}
}

func (v NullableMicrosoftEntraProviderGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftEntraProviderGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
