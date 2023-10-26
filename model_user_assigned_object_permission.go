/*
authentik

Making authentication simple.

API version: 2023.10.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// UserAssignedObjectPermission Users assigned object permission serializer
type UserAssignedObjectPermission struct {
	Pk int32 `json:"pk"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
	// User's display name.
	Name string `json:"name"`
	// Designates whether this user should be treated as active. Unselect this instead of deleting accounts.
	IsActive    *bool                  `json:"is_active,omitempty"`
	LastLogin   NullableTime           `json:"last_login,omitempty"`
	Email       *string                `json:"email,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	Uid         string                 `json:"uid"`
	Permissions []UserObjectPermission `json:"permissions"`
	IsSuperuser bool                   `json:"is_superuser"`
}

// NewUserAssignedObjectPermission instantiates a new UserAssignedObjectPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAssignedObjectPermission(pk int32, username string, name string, uid string, permissions []UserObjectPermission, isSuperuser bool) *UserAssignedObjectPermission {
	this := UserAssignedObjectPermission{}
	this.Pk = pk
	this.Username = username
	this.Name = name
	this.Uid = uid
	this.Permissions = permissions
	this.IsSuperuser = isSuperuser
	return &this
}

// NewUserAssignedObjectPermissionWithDefaults instantiates a new UserAssignedObjectPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAssignedObjectPermissionWithDefaults() *UserAssignedObjectPermission {
	this := UserAssignedObjectPermission{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserAssignedObjectPermission) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserAssignedObjectPermission) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserAssignedObjectPermission) SetPk(v int32) {
	o.Pk = v
}

// GetUsername returns the Username field value
func (o *UserAssignedObjectPermission) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserAssignedObjectPermission) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserAssignedObjectPermission) SetUsername(v string) {
	o.Username = v
}

// GetName returns the Name field value
func (o *UserAssignedObjectPermission) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserAssignedObjectPermission) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserAssignedObjectPermission) SetName(v string) {
	o.Name = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UserAssignedObjectPermission) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssignedObjectPermission) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UserAssignedObjectPermission) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UserAssignedObjectPermission) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAssignedObjectPermission) GetLastLogin() time.Time {
	if o == nil || o.LastLogin.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLogin.Get()
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAssignedObjectPermission) GetLastLoginOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLogin.Get(), o.LastLogin.IsSet()
}

// HasLastLogin returns a boolean if a field has been set.
func (o *UserAssignedObjectPermission) HasLastLogin() bool {
	if o != nil && o.LastLogin.IsSet() {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given NullableTime and assigns it to the LastLogin field.
func (o *UserAssignedObjectPermission) SetLastLogin(v time.Time) {
	o.LastLogin.Set(&v)
}

// SetLastLoginNil sets the value for LastLogin to be an explicit nil
func (o *UserAssignedObjectPermission) SetLastLoginNil() {
	o.LastLogin.Set(nil)
}

// UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
func (o *UserAssignedObjectPermission) UnsetLastLogin() {
	o.LastLogin.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserAssignedObjectPermission) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssignedObjectPermission) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserAssignedObjectPermission) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserAssignedObjectPermission) SetEmail(v string) {
	o.Email = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UserAssignedObjectPermission) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAssignedObjectPermission) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UserAssignedObjectPermission) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *UserAssignedObjectPermission) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetUid returns the Uid field value
func (o *UserAssignedObjectPermission) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *UserAssignedObjectPermission) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *UserAssignedObjectPermission) SetUid(v string) {
	o.Uid = v
}

// GetPermissions returns the Permissions field value
func (o *UserAssignedObjectPermission) GetPermissions() []UserObjectPermission {
	if o == nil {
		var ret []UserObjectPermission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *UserAssignedObjectPermission) GetPermissionsOk() ([]UserObjectPermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *UserAssignedObjectPermission) SetPermissions(v []UserObjectPermission) {
	o.Permissions = v
}

// GetIsSuperuser returns the IsSuperuser field value
func (o *UserAssignedObjectPermission) GetIsSuperuser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value
// and a boolean to check if the value has been set.
func (o *UserAssignedObjectPermission) GetIsSuperuserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSuperuser, true
}

// SetIsSuperuser sets field value
func (o *UserAssignedObjectPermission) SetIsSuperuser(v bool) {
	o.IsSuperuser = v
}

func (o UserAssignedObjectPermission) MarshalJSON() ([]byte, error) {
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
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["uid"] = o.Uid
	}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	if true {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	return json.Marshal(toSerialize)
}

type NullableUserAssignedObjectPermission struct {
	value *UserAssignedObjectPermission
	isSet bool
}

func (v NullableUserAssignedObjectPermission) Get() *UserAssignedObjectPermission {
	return v.value
}

func (v *NullableUserAssignedObjectPermission) Set(val *UserAssignedObjectPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAssignedObjectPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAssignedObjectPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAssignedObjectPermission(val *UserAssignedObjectPermission) *NullableUserAssignedObjectPermission {
	return &NullableUserAssignedObjectPermission{value: val, isSet: true}
}

func (v NullableUserAssignedObjectPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAssignedObjectPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
