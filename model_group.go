/*
authentik

Making authentication simple.

API version: 2023.8.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Group Group Serializer
type Group struct {
	Pk    string `json:"pk"`
	NumPk int32  `json:"num_pk"`
	Name  string `json:"name"`
	// Users added to this group will be superusers.
	IsSuperuser *bool                  `json:"is_superuser,omitempty"`
	Parent      NullableString         `json:"parent,omitempty"`
	ParentName  string                 `json:"parent_name"`
	Users       []int32                `json:"users,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	UsersObj    []GroupMember          `json:"users_obj"`
}

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup(pk string, numPk int32, name string, parentName string, usersObj []GroupMember) *Group {
	this := Group{}
	this.Pk = pk
	this.NumPk = numPk
	this.Name = name
	this.ParentName = parentName
	this.UsersObj = usersObj
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetPk returns the Pk field value
func (o *Group) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Group) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Group) SetPk(v string) {
	o.Pk = v
}

// GetNumPk returns the NumPk field value
func (o *Group) GetNumPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPk
}

// GetNumPkOk returns a tuple with the NumPk field value
// and a boolean to check if the value has been set.
func (o *Group) GetNumPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPk, true
}

// SetNumPk sets field value
func (o *Group) SetNumPk(v int32) {
	o.NumPk = v
}

// GetName returns the Name field value
func (o *Group) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Group) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Group) SetName(v string) {
	o.Name = v
}

// GetIsSuperuser returns the IsSuperuser field value if set, zero value otherwise.
func (o *Group) GetIsSuperuser() bool {
	if o == nil || o.IsSuperuser == nil {
		var ret bool
		return ret
	}
	return *o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIsSuperuserOk() (*bool, bool) {
	if o == nil || o.IsSuperuser == nil {
		return nil, false
	}
	return o.IsSuperuser, true
}

// HasIsSuperuser returns a boolean if a field has been set.
func (o *Group) HasIsSuperuser() bool {
	if o != nil && o.IsSuperuser != nil {
		return true
	}

	return false
}

// SetIsSuperuser gets a reference to the given bool and assigns it to the IsSuperuser field.
func (o *Group) SetIsSuperuser(v bool) {
	o.IsSuperuser = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetParent() string {
	if o == nil || o.Parent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *Group) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *Group) SetParent(v string) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *Group) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *Group) UnsetParent() {
	o.Parent.Unset()
}

// GetParentName returns the ParentName field value
func (o *Group) GetParentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value
// and a boolean to check if the value has been set.
func (o *Group) GetParentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentName, true
}

// SetParentName sets field value
func (o *Group) SetParentName(v string) {
	o.ParentName = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Group) GetUsers() []int32 {
	if o == nil || o.Users == nil {
		var ret []int32
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetUsersOk() ([]int32, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Group) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []int32 and assigns it to the Users field.
func (o *Group) SetUsers(v []int32) {
	o.Users = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Group) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Group) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Group) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetUsersObj returns the UsersObj field value
func (o *Group) GetUsersObj() []GroupMember {
	if o == nil {
		var ret []GroupMember
		return ret
	}

	return o.UsersObj
}

// GetUsersObjOk returns a tuple with the UsersObj field value
// and a boolean to check if the value has been set.
func (o *Group) GetUsersObjOk() ([]GroupMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsersObj, true
}

// SetUsersObj sets field value
func (o *Group) SetUsersObj(v []GroupMember) {
	o.UsersObj = v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["num_pk"] = o.NumPk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.IsSuperuser != nil {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if true {
		toSerialize["parent_name"] = o.ParentName
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["users_obj"] = o.UsersObj
	}
	return json.Marshal(toSerialize)
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
