/*
authentik

Making authentication simple.

API version: 2022.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserLoginStage UserLoginStage Serializer
type UserLoginStage struct {
	Pk                string    `json:"pk"`
	Name              string    `json:"name"`
	Component         string    `json:"component"`
	VerboseName       string    `json:"verbose_name"`
	VerboseNamePlural string    `json:"verbose_name_plural"`
	MetaModelName     string    `json:"meta_model_name"`
	FlowSet           []FlowSet `json:"flow_set,omitempty"`
	// Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours=-1;minutes=-2;seconds=-3)
	SessionDuration *string `json:"session_duration,omitempty"`
}

// NewUserLoginStage instantiates a new UserLoginStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLoginStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string) *UserLoginStage {
	this := UserLoginStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewUserLoginStageWithDefaults instantiates a new UserLoginStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLoginStageWithDefaults() *UserLoginStage {
	this := UserLoginStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *UserLoginStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *UserLoginStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *UserLoginStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserLoginStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *UserLoginStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *UserLoginStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *UserLoginStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *UserLoginStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *UserLoginStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *UserLoginStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *UserLoginStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *UserLoginStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *UserLoginStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *UserLoginStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *UserLoginStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetSessionDuration returns the SessionDuration field value if set, zero value otherwise.
func (o *UserLoginStage) GetSessionDuration() string {
	if o == nil || o.SessionDuration == nil {
		var ret string
		return ret
	}
	return *o.SessionDuration
}

// GetSessionDurationOk returns a tuple with the SessionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLoginStage) GetSessionDurationOk() (*string, bool) {
	if o == nil || o.SessionDuration == nil {
		return nil, false
	}
	return o.SessionDuration, true
}

// HasSessionDuration returns a boolean if a field has been set.
func (o *UserLoginStage) HasSessionDuration() bool {
	if o != nil && o.SessionDuration != nil {
		return true
	}

	return false
}

// SetSessionDuration gets a reference to the given string and assigns it to the SessionDuration field.
func (o *UserLoginStage) SetSessionDuration(v string) {
	o.SessionDuration = &v
}

func (o UserLoginStage) MarshalJSON() ([]byte, error) {
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
	if o.SessionDuration != nil {
		toSerialize["session_duration"] = o.SessionDuration
	}
	return json.Marshal(toSerialize)
}

type NullableUserLoginStage struct {
	value *UserLoginStage
	isSet bool
}

func (v NullableUserLoginStage) Get() *UserLoginStage {
	return v.value
}

func (v *NullableUserLoginStage) Set(val *UserLoginStage) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLoginStage) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLoginStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLoginStage(val *UserLoginStage) *NullableUserLoginStage {
	return &NullableUserLoginStage{value: val, isSet: true}
}

func (v NullableUserLoginStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLoginStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
