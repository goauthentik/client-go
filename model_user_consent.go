/*
authentik

Making authentication simple.

API version: 2024.6.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// UserConsent UserConsent Serializer
type UserConsent struct {
	Pk          int32        `json:"pk"`
	Expires     NullableTime `json:"expires,omitempty"`
	Expiring    *bool        `json:"expiring,omitempty"`
	User        User         `json:"user"`
	Application Application  `json:"application"`
	Permissions *string      `json:"permissions,omitempty"`
}

// NewUserConsent instantiates a new UserConsent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserConsent(pk int32, user User, application Application) *UserConsent {
	this := UserConsent{}
	this.Pk = pk
	this.User = user
	this.Application = application
	var permissions string = ""
	this.Permissions = &permissions
	return &this
}

// NewUserConsentWithDefaults instantiates a new UserConsent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserConsentWithDefaults() *UserConsent {
	this := UserConsent{}
	var permissions string = ""
	this.Permissions = &permissions
	return &this
}

// GetPk returns the Pk field value
func (o *UserConsent) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserConsent) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserConsent) SetPk(v int32) {
	o.Pk = v
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserConsent) GetExpires() time.Time {
	if o == nil || o.Expires.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserConsent) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *UserConsent) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *UserConsent) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}

// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *UserConsent) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *UserConsent) UnsetExpires() {
	o.Expires.Unset()
}

// GetExpiring returns the Expiring field value if set, zero value otherwise.
func (o *UserConsent) GetExpiring() bool {
	if o == nil || o.Expiring == nil {
		var ret bool
		return ret
	}
	return *o.Expiring
}

// GetExpiringOk returns a tuple with the Expiring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConsent) GetExpiringOk() (*bool, bool) {
	if o == nil || o.Expiring == nil {
		return nil, false
	}
	return o.Expiring, true
}

// HasExpiring returns a boolean if a field has been set.
func (o *UserConsent) HasExpiring() bool {
	if o != nil && o.Expiring != nil {
		return true
	}

	return false
}

// SetExpiring gets a reference to the given bool and assigns it to the Expiring field.
func (o *UserConsent) SetExpiring(v bool) {
	o.Expiring = &v
}

// GetUser returns the User field value
func (o *UserConsent) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserConsent) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserConsent) SetUser(v User) {
	o.User = v
}

// GetApplication returns the Application field value
func (o *UserConsent) GetApplication() Application {
	if o == nil {
		var ret Application
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *UserConsent) GetApplicationOk() (*Application, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *UserConsent) SetApplication(v Application) {
	o.Application = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UserConsent) GetPermissions() string {
	if o == nil || o.Permissions == nil {
		var ret string
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConsent) GetPermissionsOk() (*string, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UserConsent) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given string and assigns it to the Permissions field.
func (o *UserConsent) SetPermissions(v string) {
	o.Permissions = &v
}

func (o UserConsent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	if o.Expiring != nil {
		toSerialize["expiring"] = o.Expiring
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["application"] = o.Application
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableUserConsent struct {
	value *UserConsent
	isSet bool
}

func (v NullableUserConsent) Get() *UserConsent {
	return v.value
}

func (v *NullableUserConsent) Set(val *UserConsent) {
	v.value = val
	v.isSet = true
}

func (v NullableUserConsent) IsSet() bool {
	return v.isSet
}

func (v *NullableUserConsent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserConsent(val *UserConsent) *NullableUserConsent {
	return &NullableUserConsent{value: val, isSet: true}
}

func (v NullableUserConsent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserConsent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
