/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMProviderUser SCIMProviderUser Serializer
type SCIMProviderUser struct {
	Id       string      `json:"id"`
	ScimId   string      `json:"scim_id"`
	User     int32       `json:"user"`
	UserObj  GroupMember `json:"user_obj"`
	Provider int32       `json:"provider"`
}

// NewSCIMProviderUser instantiates a new SCIMProviderUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMProviderUser(id string, scimId string, user int32, userObj GroupMember, provider int32) *SCIMProviderUser {
	this := SCIMProviderUser{}
	this.Id = id
	this.ScimId = scimId
	this.User = user
	this.UserObj = userObj
	this.Provider = provider
	return &this
}

// NewSCIMProviderUserWithDefaults instantiates a new SCIMProviderUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMProviderUserWithDefaults() *SCIMProviderUser {
	this := SCIMProviderUser{}
	return &this
}

// GetId returns the Id field value
func (o *SCIMProviderUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SCIMProviderUser) SetId(v string) {
	o.Id = v
}

// GetScimId returns the ScimId field value
func (o *SCIMProviderUser) GetScimId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScimId
}

// GetScimIdOk returns a tuple with the ScimId field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderUser) GetScimIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScimId, true
}

// SetScimId sets field value
func (o *SCIMProviderUser) SetScimId(v string) {
	o.ScimId = v
}

// GetUser returns the User field value
func (o *SCIMProviderUser) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderUser) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *SCIMProviderUser) SetUser(v int32) {
	o.User = v
}

// GetUserObj returns the UserObj field value
func (o *SCIMProviderUser) GetUserObj() GroupMember {
	if o == nil {
		var ret GroupMember
		return ret
	}

	return o.UserObj
}

// GetUserObjOk returns a tuple with the UserObj field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderUser) GetUserObjOk() (*GroupMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserObj, true
}

// SetUserObj sets field value
func (o *SCIMProviderUser) SetUserObj(v GroupMember) {
	o.UserObj = v
}

// GetProvider returns the Provider field value
func (o *SCIMProviderUser) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderUser) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *SCIMProviderUser) SetProvider(v int32) {
	o.Provider = v
}

func (o SCIMProviderUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["scim_id"] = o.ScimId
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
	return json.Marshal(toSerialize)
}

type NullableSCIMProviderUser struct {
	value *SCIMProviderUser
	isSet bool
}

func (v NullableSCIMProviderUser) Get() *SCIMProviderUser {
	return v.value
}

func (v *NullableSCIMProviderUser) Set(val *SCIMProviderUser) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMProviderUser) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMProviderUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMProviderUser(val *SCIMProviderUser) *NullableSCIMProviderUser {
	return &NullableSCIMProviderUser{value: val, isSet: true}
}

func (v NullableSCIMProviderUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMProviderUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
