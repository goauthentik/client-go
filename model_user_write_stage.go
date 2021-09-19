/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserWriteStage UserWriteStage Serializer
type UserWriteStage struct {
	Pk string `json:"pk"`
	Name string `json:"name"`
	Component string `json:"component"`
	VerboseName string `json:"verbose_name"`
	VerboseNamePlural string `json:"verbose_name_plural"`
	FlowSet *[]Flow `json:"flow_set,omitempty"`
	// When set, newly created users are inactive and cannot login.
	CreateUsersAsInactive *bool `json:"create_users_as_inactive,omitempty"`
	// Optionally add newly created users to this group.
	CreateUsersGroup NullableString `json:"create_users_group,omitempty"`
}

// NewUserWriteStage instantiates a new UserWriteStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWriteStage(pk string, name string, component string, verboseName string, verboseNamePlural string) *UserWriteStage {
	this := UserWriteStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	return &this
}

// NewUserWriteStageWithDefaults instantiates a new UserWriteStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWriteStageWithDefaults() *UserWriteStage {
	this := UserWriteStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserWriteStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserWriteStage) GetPkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserWriteStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *UserWriteStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserWriteStage) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserWriteStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *UserWriteStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *UserWriteStage) GetComponentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *UserWriteStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *UserWriteStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *UserWriteStage) GetVerboseNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *UserWriteStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *UserWriteStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *UserWriteStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *UserWriteStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *UserWriteStage) GetFlowSet() []Flow {
	if o == nil || o.FlowSet == nil {
		var ret []Flow
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWriteStage) GetFlowSetOk() (*[]Flow, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *UserWriteStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []Flow and assigns it to the FlowSet field.
func (o *UserWriteStage) SetFlowSet(v []Flow) {
	o.FlowSet = &v
}

// GetCreateUsersAsInactive returns the CreateUsersAsInactive field value if set, zero value otherwise.
func (o *UserWriteStage) GetCreateUsersAsInactive() bool {
	if o == nil || o.CreateUsersAsInactive == nil {
		var ret bool
		return ret
	}
	return *o.CreateUsersAsInactive
}

// GetCreateUsersAsInactiveOk returns a tuple with the CreateUsersAsInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWriteStage) GetCreateUsersAsInactiveOk() (*bool, bool) {
	if o == nil || o.CreateUsersAsInactive == nil {
		return nil, false
	}
	return o.CreateUsersAsInactive, true
}

// HasCreateUsersAsInactive returns a boolean if a field has been set.
func (o *UserWriteStage) HasCreateUsersAsInactive() bool {
	if o != nil && o.CreateUsersAsInactive != nil {
		return true
	}

	return false
}

// SetCreateUsersAsInactive gets a reference to the given bool and assigns it to the CreateUsersAsInactive field.
func (o *UserWriteStage) SetCreateUsersAsInactive(v bool) {
	o.CreateUsersAsInactive = &v
}

// GetCreateUsersGroup returns the CreateUsersGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserWriteStage) GetCreateUsersGroup() string {
	if o == nil || o.CreateUsersGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreateUsersGroup.Get()
}

// GetCreateUsersGroupOk returns a tuple with the CreateUsersGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserWriteStage) GetCreateUsersGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreateUsersGroup.Get(), o.CreateUsersGroup.IsSet()
}

// HasCreateUsersGroup returns a boolean if a field has been set.
func (o *UserWriteStage) HasCreateUsersGroup() bool {
	if o != nil && o.CreateUsersGroup.IsSet() {
		return true
	}

	return false
}

// SetCreateUsersGroup gets a reference to the given NullableString and assigns it to the CreateUsersGroup field.
func (o *UserWriteStage) SetCreateUsersGroup(v string) {
	o.CreateUsersGroup.Set(&v)
}
// SetCreateUsersGroupNil sets the value for CreateUsersGroup to be an explicit nil
func (o *UserWriteStage) SetCreateUsersGroupNil() {
	o.CreateUsersGroup.Set(nil)
}

// UnsetCreateUsersGroup ensures that no value is present for CreateUsersGroup, not even an explicit nil
func (o *UserWriteStage) UnsetCreateUsersGroup() {
	o.CreateUsersGroup.Unset()
}

func (o UserWriteStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.CreateUsersAsInactive != nil {
		toSerialize["create_users_as_inactive"] = o.CreateUsersAsInactive
	}
	if o.CreateUsersGroup.IsSet() {
		toSerialize["create_users_group"] = o.CreateUsersGroup.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserWriteStage struct {
	value *UserWriteStage
	isSet bool
}

func (v NullableUserWriteStage) Get() *UserWriteStage {
	return v.value
}

func (v *NullableUserWriteStage) Set(val *UserWriteStage) {
	v.value = val
	v.isSet = true
}

func (v NullableUserWriteStage) IsSet() bool {
	return v.isSet
}

func (v *NullableUserWriteStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserWriteStage(val *UserWriteStage) *NullableUserWriteStage {
	return &NullableUserWriteStage{value: val, isSet: true}
}

func (v NullableUserWriteStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserWriteStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


