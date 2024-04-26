/*
authentik

Making authentication simple.

API version: 2024.4.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMSourceUser SCIMSourceUser Serializer
type SCIMSourceUser struct {
	Id         string      `json:"id"`
	User       int32       `json:"user"`
	UserObj    GroupMember `json:"user_obj"`
	Source     string      `json:"source"`
	Attributes interface{} `json:"attributes,omitempty"`
}

// NewSCIMSourceUser instantiates a new SCIMSourceUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMSourceUser(id string, user int32, userObj GroupMember, source string) *SCIMSourceUser {
	this := SCIMSourceUser{}
	this.Id = id
	this.User = user
	this.UserObj = userObj
	this.Source = source
	return &this
}

// NewSCIMSourceUserWithDefaults instantiates a new SCIMSourceUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMSourceUserWithDefaults() *SCIMSourceUser {
	this := SCIMSourceUser{}
	return &this
}

// GetId returns the Id field value
func (o *SCIMSourceUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SCIMSourceUser) SetId(v string) {
	o.Id = v
}

// GetUser returns the User field value
func (o *SCIMSourceUser) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceUser) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *SCIMSourceUser) SetUser(v int32) {
	o.User = v
}

// GetUserObj returns the UserObj field value
func (o *SCIMSourceUser) GetUserObj() GroupMember {
	if o == nil {
		var ret GroupMember
		return ret
	}

	return o.UserObj
}

// GetUserObjOk returns a tuple with the UserObj field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceUser) GetUserObjOk() (*GroupMember, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserObj, true
}

// SetUserObj sets field value
func (o *SCIMSourceUser) SetUserObj(v GroupMember) {
	o.UserObj = v
}

// GetSource returns the Source field value
func (o *SCIMSourceUser) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SCIMSourceUser) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SCIMSourceUser) SetSource(v string) {
	o.Source = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SCIMSourceUser) GetAttributes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SCIMSourceUser) GetAttributesOk() (*interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SCIMSourceUser) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *SCIMSourceUser) SetAttributes(v interface{}) {
	o.Attributes = v
}

func (o SCIMSourceUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["user_obj"] = o.UserObj
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMSourceUser struct {
	value *SCIMSourceUser
	isSet bool
}

func (v NullableSCIMSourceUser) Get() *SCIMSourceUser {
	return v.value
}

func (v *NullableSCIMSourceUser) Set(val *SCIMSourceUser) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMSourceUser) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMSourceUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMSourceUser(val *SCIMSourceUser) *NullableSCIMSourceUser {
	return &NullableSCIMSourceUser{value: val, isSet: true}
}

func (v NullableSCIMSourceUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMSourceUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
