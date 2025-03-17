/*
authentik

Making authentication simple.

API version: 2025.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MicrosoftEntraProviderUser MicrosoftEntraProviderUser Serializer
type MicrosoftEntraProviderUser struct {
	Id          string      `json:"id"`
	MicrosoftId string      `json:"microsoft_id"`
	User        int32       `json:"user"`
	UserObj     GroupMember `json:"user_obj"`
	Provider    int32       `json:"provider"`
	Attributes  interface{} `json:"attributes"`
}

// NewMicrosoftEntraProviderUser instantiates a new MicrosoftEntraProviderUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftEntraProviderUser(id string, microsoftId string, user int32, userObj GroupMember, provider int32, attributes interface{}) *MicrosoftEntraProviderUser {
	this := MicrosoftEntraProviderUser{}
	this.Id = id
	this.MicrosoftId = microsoftId
	this.User = user
	this.UserObj = userObj
	this.Provider = provider
	this.Attributes = attributes
	return &this
}

// NewMicrosoftEntraProviderUserWithDefaults instantiates a new MicrosoftEntraProviderUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftEntraProviderUserWithDefaults() *MicrosoftEntraProviderUser {
	this := MicrosoftEntraProviderUser{}
	return &this
}

// GetId returns the Id field value
func (o *MicrosoftEntraProviderUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MicrosoftEntraProviderUser) SetId(v string) {
	o.Id = v
}

// GetMicrosoftId returns the MicrosoftId field value
func (o *MicrosoftEntraProviderUser) GetMicrosoftId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicrosoftId
}

// GetMicrosoftIdOk returns a tuple with the MicrosoftId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderUser) GetMicrosoftIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicrosoftId, true
}

// SetMicrosoftId sets field value
func (o *MicrosoftEntraProviderUser) SetMicrosoftId(v string) {
	o.MicrosoftId = v
}

// GetUser returns the User field value
func (o *MicrosoftEntraProviderUser) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderUser) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *MicrosoftEntraProviderUser) SetUser(v int32) {
	o.User = v
}

// GetUserObj returns the UserObj field value
func (o *MicrosoftEntraProviderUser) GetUserObj() GroupMember {
	if o == nil {
		var ret GroupMember
		return ret
	}

	return o.UserObj
}

// GetUserObjOk returns a tuple with the UserObj field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderUser) GetUserObjOk() (*GroupMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserObj, true
}

// SetUserObj sets field value
func (o *MicrosoftEntraProviderUser) SetUserObj(v GroupMember) {
	o.UserObj = v
}

// GetProvider returns the Provider field value
func (o *MicrosoftEntraProviderUser) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderUser) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *MicrosoftEntraProviderUser) SetProvider(v int32) {
	o.Provider = v
}

// GetAttributes returns the Attributes field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *MicrosoftEntraProviderUser) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftEntraProviderUser) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *MicrosoftEntraProviderUser) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o MicrosoftEntraProviderUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["microsoft_id"] = o.MicrosoftId
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["user_obj"] = o.UserObj
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftEntraProviderUser struct {
	value *MicrosoftEntraProviderUser
	isSet bool
}

func (v NullableMicrosoftEntraProviderUser) Get() *MicrosoftEntraProviderUser {
	return v.value
}

func (v *NullableMicrosoftEntraProviderUser) Set(val *MicrosoftEntraProviderUser) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftEntraProviderUser) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftEntraProviderUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftEntraProviderUser(val *MicrosoftEntraProviderUser) *NullableMicrosoftEntraProviderUser {
	return &NullableMicrosoftEntraProviderUser{value: val, isSet: true}
}

func (v NullableMicrosoftEntraProviderUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftEntraProviderUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
