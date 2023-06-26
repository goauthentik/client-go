/*
authentik

Making authentication simple.

API version: 2023.5.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// PatchedUserRequest User Serializer
type PatchedUserRequest struct {
	Username *string `json:"username,omitempty"`
	// User's display name.
	Name *string `json:"name,omitempty"`
	// Designates whether this user should be treated as active. Unselect this instead of deleting accounts.
	IsActive   *bool                  `json:"is_active,omitempty"`
	LastLogin  NullableTime           `json:"last_login,omitempty"`
	Groups     []string               `json:"groups,omitempty"`
	Email      *string                `json:"email,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Path       *string                `json:"path,omitempty"`
}

// NewPatchedUserRequest instantiates a new PatchedUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserRequest() *PatchedUserRequest {
	this := PatchedUserRequest{}
	return &this
}

// NewPatchedUserRequestWithDefaults instantiates a new PatchedUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserRequestWithDefaults() *PatchedUserRequest {
	this := PatchedUserRequest{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PatchedUserRequest) SetUsername(v string) {
	o.Username = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedUserRequest) SetName(v string) {
	o.Name = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PatchedUserRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUserRequest) GetLastLogin() time.Time {
	if o == nil || o.LastLogin.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLogin.Get()
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUserRequest) GetLastLoginOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLogin.Get(), o.LastLogin.IsSet()
}

// HasLastLogin returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasLastLogin() bool {
	if o != nil && o.LastLogin.IsSet() {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given NullableTime and assigns it to the LastLogin field.
func (o *PatchedUserRequest) SetLastLogin(v time.Time) {
	o.LastLogin.Set(&v)
}

// SetLastLoginNil sets the value for LastLogin to be an explicit nil
func (o *PatchedUserRequest) SetLastLoginNil() {
	o.LastLogin.Set(nil)
}

// UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
func (o *PatchedUserRequest) UnsetLastLogin() {
	o.LastLogin.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetGroupsOk() ([]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *PatchedUserRequest) SetGroups(v []string) {
	o.Groups = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PatchedUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *PatchedUserRequest) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PatchedUserRequest) SetPath(v string) {
	o.Path = &v
}

func (o PatchedUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.LastLogin.IsSet() {
		toSerialize["last_login"] = o.LastLogin.Get()
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedUserRequest struct {
	value *PatchedUserRequest
	isSet bool
}

func (v NullablePatchedUserRequest) Get() *PatchedUserRequest {
	return v.value
}

func (v *NullablePatchedUserRequest) Set(val *PatchedUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserRequest(val *PatchedUserRequest) *NullablePatchedUserRequest {
	return &NullablePatchedUserRequest{value: val, isSet: true}
}

func (v NullablePatchedUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
