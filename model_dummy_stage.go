/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DummyStage DummyStage Serializer
type DummyStage struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// Get object type so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string    `json:"meta_model_name"`
	FlowSet       []FlowSet `json:"flow_set,omitempty"`
	ThrowError    *bool     `json:"throw_error,omitempty"`
}

// NewDummyStage instantiates a new DummyStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDummyStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string) *DummyStage {
	this := DummyStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewDummyStageWithDefaults instantiates a new DummyStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDummyStageWithDefaults() *DummyStage {
	this := DummyStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *DummyStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *DummyStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *DummyStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *DummyStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DummyStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DummyStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *DummyStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DummyStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *DummyStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *DummyStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *DummyStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *DummyStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *DummyStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *DummyStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *DummyStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *DummyStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *DummyStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *DummyStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *DummyStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *DummyStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *DummyStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetThrowError returns the ThrowError field value if set, zero value otherwise.
func (o *DummyStage) GetThrowError() bool {
	if o == nil || o.ThrowError == nil {
		var ret bool
		return ret
	}
	return *o.ThrowError
}

// GetThrowErrorOk returns a tuple with the ThrowError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyStage) GetThrowErrorOk() (*bool, bool) {
	if o == nil || o.ThrowError == nil {
		return nil, false
	}
	return o.ThrowError, true
}

// HasThrowError returns a boolean if a field has been set.
func (o *DummyStage) HasThrowError() bool {
	if o != nil && o.ThrowError != nil {
		return true
	}

	return false
}

// SetThrowError gets a reference to the given bool and assigns it to the ThrowError field.
func (o *DummyStage) SetThrowError(v bool) {
	o.ThrowError = &v
}

func (o DummyStage) MarshalJSON() ([]byte, error) {
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
	if o.ThrowError != nil {
		toSerialize["throw_error"] = o.ThrowError
	}
	return json.Marshal(toSerialize)
}

type NullableDummyStage struct {
	value *DummyStage
	isSet bool
}

func (v NullableDummyStage) Get() *DummyStage {
	return v.value
}

func (v *NullableDummyStage) Set(val *DummyStage) {
	v.value = val
	v.isSet = true
}

func (v NullableDummyStage) IsSet() bool {
	return v.isSet
}

func (v *NullableDummyStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDummyStage(val *DummyStage) *NullableDummyStage {
	return &NullableDummyStage{value: val, isSet: true}
}

func (v NullableDummyStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDummyStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
