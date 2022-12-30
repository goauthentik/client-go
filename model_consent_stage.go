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

// ConsentStage ConsentStage Serializer
type ConsentStage struct {
	Pk                string                `json:"pk"`
	Name              string                `json:"name"`
	Component         string                `json:"component"`
	VerboseName       string                `json:"verbose_name"`
	VerboseNamePlural string                `json:"verbose_name_plural"`
	MetaModelName     string                `json:"meta_model_name"`
	FlowSet           []FlowSet             `json:"flow_set,omitempty"`
	Mode              *ConsentStageModeEnum `json:"mode,omitempty"`
	// Offset after which consent expires. (Format: hours=1;minutes=2;seconds=3).
	ConsentExpireIn *string `json:"consent_expire_in,omitempty"`
}

// NewConsentStage instantiates a new ConsentStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string) *ConsentStage {
	this := ConsentStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewConsentStageWithDefaults instantiates a new ConsentStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentStageWithDefaults() *ConsentStage {
	this := ConsentStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *ConsentStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *ConsentStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *ConsentStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *ConsentStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConsentStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsentStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *ConsentStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ConsentStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *ConsentStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *ConsentStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *ConsentStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *ConsentStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *ConsentStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *ConsentStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *ConsentStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *ConsentStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *ConsentStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *ConsentStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *ConsentStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *ConsentStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *ConsentStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ConsentStage) GetMode() ConsentStageModeEnum {
	if o == nil || o.Mode == nil {
		var ret ConsentStageModeEnum
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentStage) GetModeOk() (*ConsentStageModeEnum, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ConsentStage) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given ConsentStageModeEnum and assigns it to the Mode field.
func (o *ConsentStage) SetMode(v ConsentStageModeEnum) {
	o.Mode = &v
}

// GetConsentExpireIn returns the ConsentExpireIn field value if set, zero value otherwise.
func (o *ConsentStage) GetConsentExpireIn() string {
	if o == nil || o.ConsentExpireIn == nil {
		var ret string
		return ret
	}
	return *o.ConsentExpireIn
}

// GetConsentExpireInOk returns a tuple with the ConsentExpireIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentStage) GetConsentExpireInOk() (*string, bool) {
	if o == nil || o.ConsentExpireIn == nil {
		return nil, false
	}
	return o.ConsentExpireIn, true
}

// HasConsentExpireIn returns a boolean if a field has been set.
func (o *ConsentStage) HasConsentExpireIn() bool {
	if o != nil && o.ConsentExpireIn != nil {
		return true
	}

	return false
}

// SetConsentExpireIn gets a reference to the given string and assigns it to the ConsentExpireIn field.
func (o *ConsentStage) SetConsentExpireIn(v string) {
	o.ConsentExpireIn = &v
}

func (o ConsentStage) MarshalJSON() ([]byte, error) {
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
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.ConsentExpireIn != nil {
		toSerialize["consent_expire_in"] = o.ConsentExpireIn
	}
	return json.Marshal(toSerialize)
}

type NullableConsentStage struct {
	value *ConsentStage
	isSet bool
}

func (v NullableConsentStage) Get() *ConsentStage {
	return v.value
}

func (v *NullableConsentStage) Set(val *ConsentStage) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentStage) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentStage(val *ConsentStage) *NullableConsentStage {
	return &NullableConsentStage{value: val, isSet: true}
}

func (v NullableConsentStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
