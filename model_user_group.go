/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the UserGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroup{}

// UserGroup Simplified Group Serializer for user's groups
type UserGroup struct {
	Pk string `json:"pk"`
	// Get a numerical, int32 ID for the group
	NumPk int32  `json:"num_pk"`
	Name  string `json:"name"`
	// Users added to this group will be superusers.
	IsSuperuser *bool                  `json:"is_superuser,omitempty"`
	Parent      NullableString         `json:"parent,omitempty"`
	ParentName  NullableString         `json:"parent_name"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
}

type _UserGroup UserGroup

// NewUserGroup instantiates a new UserGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroup(pk string, numPk int32, name string, parentName NullableString) *UserGroup {
	this := UserGroup{}
	this.Pk = pk
	this.NumPk = numPk
	this.Name = name
	this.ParentName = parentName
	return &this
}

// NewUserGroupWithDefaults instantiates a new UserGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupWithDefaults() *UserGroup {
	this := UserGroup{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserGroup) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserGroup) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserGroup) SetPk(v string) {
	o.Pk = v
}

// GetNumPk returns the NumPk field value
func (o *UserGroup) GetNumPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPk
}

// GetNumPkOk returns a tuple with the NumPk field value
// and a boolean to check if the value has been set.
func (o *UserGroup) GetNumPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPk, true
}

// SetNumPk sets field value
func (o *UserGroup) SetNumPk(v int32) {
	o.NumPk = v
}

// GetName returns the Name field value
func (o *UserGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserGroup) SetName(v string) {
	o.Name = v
}

// GetIsSuperuser returns the IsSuperuser field value if set, zero value otherwise.
func (o *UserGroup) GetIsSuperuser() bool {
	if o == nil || IsNil(o.IsSuperuser) {
		var ret bool
		return ret
	}
	return *o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetIsSuperuserOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuperuser) {
		return nil, false
	}
	return o.IsSuperuser, true
}

// HasIsSuperuser returns a boolean if a field has been set.
func (o *UserGroup) HasIsSuperuser() bool {
	if o != nil && !IsNil(o.IsSuperuser) {
		return true
	}

	return false
}

// SetIsSuperuser gets a reference to the given bool and assigns it to the IsSuperuser field.
func (o *UserGroup) SetIsSuperuser(v bool) {
	o.IsSuperuser = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserGroup) GetParent() string {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret string
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroup) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *UserGroup) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableString and assigns it to the Parent field.
func (o *UserGroup) SetParent(v string) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *UserGroup) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *UserGroup) UnsetParent() {
	o.Parent.Unset()
}

// GetParentName returns the ParentName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserGroup) GetParentName() string {
	if o == nil || o.ParentName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentName.Get()
}

// GetParentNameOk returns a tuple with the ParentName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserGroup) GetParentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentName.Get(), o.ParentName.IsSet()
}

// SetParentName sets field value
func (o *UserGroup) SetParentName(v string) {
	o.ParentName.Set(&v)
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UserGroup) GetAttributes() map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroup) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UserGroup) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *UserGroup) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o UserGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pk"] = o.Pk
	toSerialize["num_pk"] = o.NumPk
	toSerialize["name"] = o.Name
	if !IsNil(o.IsSuperuser) {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	toSerialize["parent_name"] = o.ParentName.Get()
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *UserGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pk",
		"num_pk",
		"name",
		"parent_name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserGroup := _UserGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGroup)

	if err != nil {
		return err
	}

	*o = UserGroup(varUserGroup)

	return err
}

type NullableUserGroup struct {
	value *UserGroup
	isSet bool
}

func (v NullableUserGroup) Get() *UserGroup {
	return v.value
}

func (v *NullableUserGroup) Set(val *UserGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroup(val *UserGroup) *NullableUserGroup {
	return &NullableUserGroup{value: val, isSet: true}
}

func (v NullableUserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
