/*
authentik

Making authentication simple.

API version: 2022.7.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// PolicyBindingUserObj struct for PolicyBindingUserObj
type PolicyBindingUserObj struct {
	Pk       int32  `json:"pk"`
	Username string `json:"username"`
	// User's display name.
	Name string `json:"name"`
	// Designates whether this user should be treated as active. Unselect this instead of deleting accounts.
	IsActive    *bool                  `json:"is_active,omitempty"`
	LastLogin   NullableTime           `json:"last_login,omitempty"`
	IsSuperuser bool                   `json:"is_superuser"`
	Groups      []string               `json:"groups"`
	GroupsObj   []Group                `json:"groups_obj"`
	Email       *string                `json:"email,omitempty"`
	Avatar      string                 `json:"avatar"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	Uid         string                 `json:"uid"`
	Path        *string                `json:"path,omitempty"`
}

// NewPolicyBindingUserObj instantiates a new PolicyBindingUserObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyBindingUserObj(pk int32, username string, name string, isSuperuser bool, groups []string, groupsObj []Group, avatar string, uid string) *PolicyBindingUserObj {
	this := PolicyBindingUserObj{}
	this.Pk = pk
	this.Username = username
	this.Name = name
	this.IsSuperuser = isSuperuser
	this.Groups = groups
	this.GroupsObj = groupsObj
	this.Avatar = avatar
	this.Uid = uid
	return &this
}

// NewPolicyBindingUserObjWithDefaults instantiates a new PolicyBindingUserObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyBindingUserObjWithDefaults() *PolicyBindingUserObj {
	this := PolicyBindingUserObj{}
	return &this
}

// GetPk returns the Pk field value
func (o *PolicyBindingUserObj) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *PolicyBindingUserObj) SetPk(v int32) {
	o.Pk = v
}

// GetUsername returns the Username field value
func (o *PolicyBindingUserObj) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *PolicyBindingUserObj) SetUsername(v string) {
	o.Username = v
}

// GetName returns the Name field value
func (o *PolicyBindingUserObj) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PolicyBindingUserObj) SetName(v string) {
	o.Name = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PolicyBindingUserObj) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PolicyBindingUserObj) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PolicyBindingUserObj) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyBindingUserObj) GetLastLogin() time.Time {
	if o == nil || o.LastLogin.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLogin.Get()
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyBindingUserObj) GetLastLoginOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLogin.Get(), o.LastLogin.IsSet()
}

// HasLastLogin returns a boolean if a field has been set.
func (o *PolicyBindingUserObj) HasLastLogin() bool {
	if o != nil && o.LastLogin.IsSet() {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given NullableTime and assigns it to the LastLogin field.
func (o *PolicyBindingUserObj) SetLastLogin(v time.Time) {
	o.LastLogin.Set(&v)
}

// SetLastLoginNil sets the value for LastLogin to be an explicit nil
func (o *PolicyBindingUserObj) SetLastLoginNil() {
	o.LastLogin.Set(nil)
}

// UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
func (o *PolicyBindingUserObj) UnsetLastLogin() {
	o.LastLogin.Unset()
}

// GetIsSuperuser returns the IsSuperuser field value
func (o *PolicyBindingUserObj) GetIsSuperuser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetIsSuperuserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSuperuser, true
}

// SetIsSuperuser sets field value
func (o *PolicyBindingUserObj) SetIsSuperuser(v bool) {
	o.IsSuperuser = v
}

// GetGroups returns the Groups field value
func (o *PolicyBindingUserObj) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetGroupsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *PolicyBindingUserObj) SetGroups(v []string) {
	o.Groups = v
}

// GetGroupsObj returns the GroupsObj field value
func (o *PolicyBindingUserObj) GetGroupsObj() []Group {
	if o == nil {
		var ret []Group
		return ret
	}

	return o.GroupsObj
}

// GetGroupsObjOk returns a tuple with the GroupsObj field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetGroupsObjOk() ([]Group, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupsObj, true
}

// SetGroupsObj sets field value
func (o *PolicyBindingUserObj) SetGroupsObj(v []Group) {
	o.GroupsObj = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PolicyBindingUserObj) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PolicyBindingUserObj) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PolicyBindingUserObj) SetEmail(v string) {
	o.Email = &v
}

// GetAvatar returns the Avatar field value
func (o *PolicyBindingUserObj) GetAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avatar, true
}

// SetAvatar sets field value
func (o *PolicyBindingUserObj) SetAvatar(v string) {
	o.Avatar = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PolicyBindingUserObj) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PolicyBindingUserObj) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *PolicyBindingUserObj) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetUid returns the Uid field value
func (o *PolicyBindingUserObj) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *PolicyBindingUserObj) SetUid(v string) {
	o.Uid = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PolicyBindingUserObj) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBindingUserObj) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PolicyBindingUserObj) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PolicyBindingUserObj) SetPath(v string) {
	o.Path = &v
}

func (o PolicyBindingUserObj) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.LastLogin.IsSet() {
		toSerialize["last_login"] = o.LastLogin.Get()
	}
	if true {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	if true {
		toSerialize["groups"] = o.Groups
	}
	if true {
		toSerialize["groups_obj"] = o.GroupsObj
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["avatar"] = o.Avatar
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["uid"] = o.Uid
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyBindingUserObj struct {
	value *PolicyBindingUserObj
	isSet bool
}

func (v NullablePolicyBindingUserObj) Get() *PolicyBindingUserObj {
	return v.value
}

func (v *NullablePolicyBindingUserObj) Set(val *PolicyBindingUserObj) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyBindingUserObj) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyBindingUserObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyBindingUserObj(val *PolicyBindingUserObj) *NullablePolicyBindingUserObj {
	return &NullablePolicyBindingUserObj{value: val, isSet: true}
}

func (v NullablePolicyBindingUserObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyBindingUserObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
