/*
authentik

Making authentication simple.

API version: 2023.1.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserWriteStage UserWriteStage Serializer
type UserWriteStage struct {
	Pk                string                `json:"pk"`
	Name              string                `json:"name"`
	Component         string                `json:"component"`
	VerboseName       string                `json:"verbose_name"`
	VerboseNamePlural string                `json:"verbose_name_plural"`
	MetaModelName     string                `json:"meta_model_name"`
	FlowSet           []FlowSet             `json:"flow_set,omitempty"`
	UserCreationMode  *UserCreationModeEnum `json:"user_creation_mode,omitempty"`
	// When set, newly created users are inactive and cannot login.
	CreateUsersAsInactive *bool `json:"create_users_as_inactive,omitempty"`
	// Optionally add newly created users to this group.
	CreateUsersGroup NullableString `json:"create_users_group,omitempty"`
	UserPathTemplate *string        `json:"user_path_template,omitempty"`
}

// NewUserWriteStage instantiates a new UserWriteStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWriteStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string) *UserWriteStage {
	this := UserWriteStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
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
	if o == nil {
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
	if o == nil {
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
	if o == nil {
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
	if o == nil {
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
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *UserWriteStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *UserWriteStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *UserWriteStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *UserWriteStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *UserWriteStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWriteStage) GetFlowSetOk() ([]FlowSet, bool) {
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

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *UserWriteStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetUserCreationMode returns the UserCreationMode field value if set, zero value otherwise.
func (o *UserWriteStage) GetUserCreationMode() UserCreationModeEnum {
	if o == nil || o.UserCreationMode == nil {
		var ret UserCreationModeEnum
		return ret
	}
	return *o.UserCreationMode
}

// GetUserCreationModeOk returns a tuple with the UserCreationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWriteStage) GetUserCreationModeOk() (*UserCreationModeEnum, bool) {
	if o == nil || o.UserCreationMode == nil {
		return nil, false
	}
	return o.UserCreationMode, true
}

// HasUserCreationMode returns a boolean if a field has been set.
func (o *UserWriteStage) HasUserCreationMode() bool {
	if o != nil && o.UserCreationMode != nil {
		return true
	}

	return false
}

// SetUserCreationMode gets a reference to the given UserCreationModeEnum and assigns it to the UserCreationMode field.
func (o *UserWriteStage) SetUserCreationMode(v UserCreationModeEnum) {
	o.UserCreationMode = &v
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
	if o == nil {
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

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *UserWriteStage) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserWriteStage) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *UserWriteStage) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *UserWriteStage) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
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
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.UserCreationMode != nil {
		toSerialize["user_creation_mode"] = o.UserCreationMode
	}
	if o.CreateUsersAsInactive != nil {
		toSerialize["create_users_as_inactive"] = o.CreateUsersAsInactive
	}
	if o.CreateUsersGroup.IsSet() {
		toSerialize["create_users_group"] = o.CreateUsersGroup.Get()
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
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
