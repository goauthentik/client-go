/*
authentik

Making authentication simple.

API version: 2023.5.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedGroupRequest Group Serializer
type PatchedGroupRequest struct {
	Name *string `json:"name,omitempty"`
	// Users added to this group will be superusers.
	IsSuperuser *bool                  `json:"is_superuser,omitempty"`
	Parent      NullableString         `json:"parent,omitempty"`
	Users       []int32                `json:"users,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
}

// NewPatchedGroupRequest instantiates a new PatchedGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedGroupRequest() *PatchedGroupRequest {
	this := PatchedGroupRequest{}
	return &this
}

// NewPatchedGroupRequestWithDefaults instantiates a new PatchedGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedGroupRequestWithDefaults() *PatchedGroupRequest {
	this := PatchedGroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedGroupRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedGroupRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetIsSuperuser returns the IsSuperuser field value if set, zero value otherwise.
func (o *PatchedGroupRequest) GetIsSuperuser() bool {
	if o == nil || o.IsSuperuser == nil {
		var ret bool
		return ret
	}
	return *o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGroupRequest) GetIsSuperuserOk() (*bool, bool) {
	if o == nil || o.IsSuperuser == nil {
		return nil, false
	}
	return o.IsSuperuser, true
}

// HasIsSuperuser returns a boolean if a field has been set.
func (o *PatchedGroupRequest) HasIsSuperuser() bool {
	if o != nil && o.IsSuperuser != nil {
		return true
	}

	return false
}

// SetIsSuperuser gets a reference to the given bool and assigns it to the IsSuperuser field.
func (o *PatchedGroupRequest) SetIsSuperuser(v bool) {
	o.IsSuperuser = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedGroupRequest) GetParent() string {
	if o == nil || o.Parent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedGroupRequest) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedGroupRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *PatchedGroupRequest) SetParent(v string) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *PatchedGroupRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *PatchedGroupRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *PatchedGroupRequest) GetUsers() []int32 {
	if o == nil || o.Users == nil {
		var ret []int32
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGroupRequest) GetUsersOk() ([]int32, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *PatchedGroupRequest) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []int32 and assigns it to the Users field.
func (o *PatchedGroupRequest) SetUsers(v []int32) {
	o.Users = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PatchedGroupRequest) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGroupRequest) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PatchedGroupRequest) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *PatchedGroupRequest) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o PatchedGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.IsSuperuser != nil {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedGroupRequest struct {
	value *PatchedGroupRequest
	isSet bool
}

func (v NullablePatchedGroupRequest) Get() *PatchedGroupRequest {
	return v.value
}

func (v *NullablePatchedGroupRequest) Set(val *PatchedGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedGroupRequest(val *PatchedGroupRequest) *NullablePatchedGroupRequest {
	return &NullablePatchedGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
