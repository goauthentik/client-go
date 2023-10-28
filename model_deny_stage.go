/*
authentik

Making authentication simple.

API version: 2023.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DenyStage DenyStage Serializer
type DenyStage struct {
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
	DenyMessage   *string   `json:"deny_message,omitempty"`
}

// NewDenyStage instantiates a new DenyStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDenyStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string) *DenyStage {
	this := DenyStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewDenyStageWithDefaults instantiates a new DenyStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDenyStageWithDefaults() *DenyStage {
	this := DenyStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *DenyStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *DenyStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *DenyStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *DenyStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DenyStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DenyStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *DenyStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DenyStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *DenyStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *DenyStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *DenyStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *DenyStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *DenyStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *DenyStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *DenyStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *DenyStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *DenyStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *DenyStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *DenyStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DenyStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *DenyStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *DenyStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetDenyMessage returns the DenyMessage field value if set, zero value otherwise.
func (o *DenyStage) GetDenyMessage() string {
	if o == nil || o.DenyMessage == nil {
		var ret string
		return ret
	}
	return *o.DenyMessage
}

// GetDenyMessageOk returns a tuple with the DenyMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DenyStage) GetDenyMessageOk() (*string, bool) {
	if o == nil || o.DenyMessage == nil {
		return nil, false
	}
	return o.DenyMessage, true
}

// HasDenyMessage returns a boolean if a field has been set.
func (o *DenyStage) HasDenyMessage() bool {
	if o != nil && o.DenyMessage != nil {
		return true
	}

	return false
}

// SetDenyMessage gets a reference to the given string and assigns it to the DenyMessage field.
func (o *DenyStage) SetDenyMessage(v string) {
	o.DenyMessage = &v
}

func (o DenyStage) MarshalJSON() ([]byte, error) {
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
	if o.DenyMessage != nil {
		toSerialize["deny_message"] = o.DenyMessage
	}
	return json.Marshal(toSerialize)
}

type NullableDenyStage struct {
	value *DenyStage
	isSet bool
}

func (v NullableDenyStage) Get() *DenyStage {
	return v.value
}

func (v *NullableDenyStage) Set(val *DenyStage) {
	v.value = val
	v.isSet = true
}

func (v NullableDenyStage) IsSet() bool {
	return v.isSet
}

func (v *NullableDenyStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDenyStage(val *DenyStage) *NullableDenyStage {
	return &NullableDenyStage{value: val, isSet: true}
}

func (v NullableDenyStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDenyStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
