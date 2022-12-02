/*
authentik

Making authentication simple.

API version: 2022.11.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OutpostServiceConnectionObj struct for OutpostServiceConnectionObj
type OutpostServiceConnectionObj struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// If enabled, use the local connection. Required Docker socket/Kubernetes Integration
	Local             *bool  `json:"local,omitempty"`
	Component         string `json:"component"`
	VerboseName       string `json:"verbose_name"`
	VerboseNamePlural string `json:"verbose_name_plural"`
	MetaModelName     string `json:"meta_model_name"`
}

// NewOutpostServiceConnectionObj instantiates a new OutpostServiceConnectionObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutpostServiceConnectionObj(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string) *OutpostServiceConnectionObj {
	this := OutpostServiceConnectionObj{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewOutpostServiceConnectionObjWithDefaults instantiates a new OutpostServiceConnectionObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutpostServiceConnectionObjWithDefaults() *OutpostServiceConnectionObj {
	this := OutpostServiceConnectionObj{}
	return &this
}

// GetPk returns the Pk field value
func (o *OutpostServiceConnectionObj) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *OutpostServiceConnectionObj) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *OutpostServiceConnectionObj) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *OutpostServiceConnectionObj) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OutpostServiceConnectionObj) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OutpostServiceConnectionObj) SetName(v string) {
	o.Name = v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *OutpostServiceConnectionObj) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutpostServiceConnectionObj) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *OutpostServiceConnectionObj) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *OutpostServiceConnectionObj) SetLocal(v bool) {
	o.Local = &v
}

// GetComponent returns the Component field value
func (o *OutpostServiceConnectionObj) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *OutpostServiceConnectionObj) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *OutpostServiceConnectionObj) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *OutpostServiceConnectionObj) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *OutpostServiceConnectionObj) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *OutpostServiceConnectionObj) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *OutpostServiceConnectionObj) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *OutpostServiceConnectionObj) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *OutpostServiceConnectionObj) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *OutpostServiceConnectionObj) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *OutpostServiceConnectionObj) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *OutpostServiceConnectionObj) SetMetaModelName(v string) {
	o.MetaModelName = v
}

func (o OutpostServiceConnectionObj) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
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
	return json.Marshal(toSerialize)
}

type NullableOutpostServiceConnectionObj struct {
	value *OutpostServiceConnectionObj
	isSet bool
}

func (v NullableOutpostServiceConnectionObj) Get() *OutpostServiceConnectionObj {
	return v.value
}

func (v *NullableOutpostServiceConnectionObj) Set(val *OutpostServiceConnectionObj) {
	v.value = val
	v.isSet = true
}

func (v NullableOutpostServiceConnectionObj) IsSet() bool {
	return v.isSet
}

func (v *NullableOutpostServiceConnectionObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutpostServiceConnectionObj(val *OutpostServiceConnectionObj) *NullableOutpostServiceConnectionObj {
	return &NullableOutpostServiceConnectionObj{value: val, isSet: true}
}

func (v NullableOutpostServiceConnectionObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutpostServiceConnectionObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
