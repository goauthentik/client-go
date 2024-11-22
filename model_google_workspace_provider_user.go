/*
authentik

Making authentication simple.

API version: 2024.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GoogleWorkspaceProviderUser GoogleWorkspaceProviderUser Serializer
type GoogleWorkspaceProviderUser struct {
	Id         string      `json:"id"`
	GoogleId   string      `json:"google_id"`
	User       int32       `json:"user"`
	UserObj    GroupMember `json:"user_obj"`
	Provider   int32       `json:"provider"`
	Attributes interface{} `json:"attributes"`
}

// NewGoogleWorkspaceProviderUser instantiates a new GoogleWorkspaceProviderUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleWorkspaceProviderUser(id string, googleId string, user int32, userObj GroupMember, provider int32, attributes interface{}) *GoogleWorkspaceProviderUser {
	this := GoogleWorkspaceProviderUser{}
	this.Id = id
	this.GoogleId = googleId
	this.User = user
	this.UserObj = userObj
	this.Provider = provider
	this.Attributes = attributes
	return &this
}

// NewGoogleWorkspaceProviderUserWithDefaults instantiates a new GoogleWorkspaceProviderUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleWorkspaceProviderUserWithDefaults() *GoogleWorkspaceProviderUser {
	this := GoogleWorkspaceProviderUser{}
	return &this
}

// GetId returns the Id field value
func (o *GoogleWorkspaceProviderUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GoogleWorkspaceProviderUser) SetId(v string) {
	o.Id = v
}

// GetGoogleId returns the GoogleId field value
func (o *GoogleWorkspaceProviderUser) GetGoogleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoogleId
}

// GetGoogleIdOk returns a tuple with the GoogleId field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderUser) GetGoogleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoogleId, true
}

// SetGoogleId sets field value
func (o *GoogleWorkspaceProviderUser) SetGoogleId(v string) {
	o.GoogleId = v
}

// GetUser returns the User field value
func (o *GoogleWorkspaceProviderUser) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderUser) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GoogleWorkspaceProviderUser) SetUser(v int32) {
	o.User = v
}

// GetUserObj returns the UserObj field value
func (o *GoogleWorkspaceProviderUser) GetUserObj() GroupMember {
	if o == nil {
		var ret GroupMember
		return ret
	}

	return o.UserObj
}

// GetUserObjOk returns a tuple with the UserObj field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderUser) GetUserObjOk() (*GroupMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserObj, true
}

// SetUserObj sets field value
func (o *GoogleWorkspaceProviderUser) SetUserObj(v GroupMember) {
	o.UserObj = v
}

// GetProvider returns the Provider field value
func (o *GoogleWorkspaceProviderUser) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *GoogleWorkspaceProviderUser) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *GoogleWorkspaceProviderUser) SetProvider(v int32) {
	o.Provider = v
}

// GetAttributes returns the Attributes field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GoogleWorkspaceProviderUser) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleWorkspaceProviderUser) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GoogleWorkspaceProviderUser) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o GoogleWorkspaceProviderUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["google_id"] = o.GoogleId
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

type NullableGoogleWorkspaceProviderUser struct {
	value *GoogleWorkspaceProviderUser
	isSet bool
}

func (v NullableGoogleWorkspaceProviderUser) Get() *GoogleWorkspaceProviderUser {
	return v.value
}

func (v *NullableGoogleWorkspaceProviderUser) Set(val *GoogleWorkspaceProviderUser) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleWorkspaceProviderUser) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleWorkspaceProviderUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleWorkspaceProviderUser(val *GoogleWorkspaceProviderUser) *NullableGoogleWorkspaceProviderUser {
	return &NullableGoogleWorkspaceProviderUser{value: val, isSet: true}
}

func (v NullableGoogleWorkspaceProviderUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleWorkspaceProviderUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
