/*
authentik

Making authentication simple.

API version: 2022.7.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SelectableStage Serializer for stages which can be selected by users
type SelectableStage struct {
	Pk            string `json:"pk"`
	Name          string `json:"name"`
	VerboseName   string `json:"verbose_name"`
	MetaModelName string `json:"meta_model_name"`
}

// NewSelectableStage instantiates a new SelectableStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectableStage(pk string, name string, verboseName string, metaModelName string) *SelectableStage {
	this := SelectableStage{}
	this.Pk = pk
	this.Name = name
	this.VerboseName = verboseName
	this.MetaModelName = metaModelName
	return &this
}

// NewSelectableStageWithDefaults instantiates a new SelectableStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectableStageWithDefaults() *SelectableStage {
	this := SelectableStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *SelectableStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *SelectableStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *SelectableStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *SelectableStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SelectableStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SelectableStage) SetName(v string) {
	o.Name = v
}

// GetVerboseName returns the VerboseName field value
func (o *SelectableStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *SelectableStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *SelectableStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *SelectableStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *SelectableStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *SelectableStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

func (o SelectableStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	return json.Marshal(toSerialize)
}

type NullableSelectableStage struct {
	value *SelectableStage
	isSet bool
}

func (v NullableSelectableStage) Get() *SelectableStage {
	return v.value
}

func (v *NullableSelectableStage) Set(val *SelectableStage) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectableStage) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectableStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectableStage(val *SelectableStage) *NullableSelectableStage {
	return &NullableSelectableStage{value: val, isSet: true}
}

func (v NullableSelectableStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectableStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
