/*
authentik

Making authentication simple.

API version: 2024.12.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserGroupRequest Simplified Group Serializer for user's groups
type UserGroupRequest struct {
	Name string `json:"name"`
	// Users added to this group will be superusers.
	IsSuperuser *bool                  `json:"is_superuser,omitempty"`
	Parent      NullableString         `json:"parent,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
}

// NewUserGroupRequest instantiates a new UserGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupRequest(name string) *UserGroupRequest {
	this := UserGroupRequest{}
	this.Name = name
	return &this
}

// NewUserGroupRequestWithDefaults instantiates a new UserGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupRequestWithDefaults() *UserGroupRequest {
	this := UserGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UserGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserGroupRequest) SetName(v string) {
	o.Name = v
}

// GetIsSuperuser returns the IsSuperuser field value if set, zero value otherwise.
func (o *UserGroupRequest) GetIsSuperuser() bool {
	if o == nil || o.IsSuperuser == nil {
		var ret bool
		return ret
	}
	return *o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupRequest) GetIsSuperuserOk() (*bool, bool) {
	if o == nil || o.IsSuperuser == nil {
		return nil, false
	}
	return o.IsSuperuser, true
}

// HasIsSuperuser returns a boolean if a field has been set.
func (o *UserGroupRequest) HasIsSuperuser() bool {
	if o != nil && o.IsSuperuser != nil {
		return true
	}

	return false
}

// SetIsSuperuser gets a reference to the given bool and assigns it to the IsSuperuser field.
func (o *UserGroupRequest) SetIsSuperuser(v bool) {
	o.IsSuperuser = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserGroupRequest) GetParent() string {
	if o == nil || o.Parent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroupRequest) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *UserGroupRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *UserGroupRequest) SetParent(v string) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *UserGroupRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *UserGroupRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UserGroupRequest) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupRequest) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UserGroupRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *UserGroupRequest) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o UserGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.IsSuperuser != nil {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableUserGroupRequest struct {
	value *UserGroupRequest
	isSet bool
}

func (v NullableUserGroupRequest) Get() *UserGroupRequest {
	return v.value
}

func (v *NullableUserGroupRequest) Set(val *UserGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupRequest(val *UserGroupRequest) *NullableUserGroupRequest {
	return &NullableUserGroupRequest{value: val, isSet: true}
}

func (v NullableUserGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
