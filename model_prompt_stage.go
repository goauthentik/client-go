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

// PromptStage PromptStage Serializer
type PromptStage struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// Get object type so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName      string    `json:"meta_model_name"`
	FlowSet            []FlowSet `json:"flow_set,omitempty"`
	Fields             []string  `json:"fields"`
	ValidationPolicies []string  `json:"validation_policies,omitempty"`
}

// NewPromptStage instantiates a new PromptStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, fields []string) *PromptStage {
	this := PromptStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Fields = fields
	return &this
}

// NewPromptStageWithDefaults instantiates a new PromptStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptStageWithDefaults() *PromptStage {
	this := PromptStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *PromptStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *PromptStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *PromptStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *PromptStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PromptStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PromptStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *PromptStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *PromptStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *PromptStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *PromptStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *PromptStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *PromptStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *PromptStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *PromptStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *PromptStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *PromptStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *PromptStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *PromptStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PromptStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PromptStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *PromptStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetFields returns the Fields field value
func (o *PromptStage) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *PromptStage) GetFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *PromptStage) SetFields(v []string) {
	o.Fields = v
}

// GetValidationPolicies returns the ValidationPolicies field value if set, zero value otherwise.
func (o *PromptStage) GetValidationPolicies() []string {
	if o == nil || o.ValidationPolicies == nil {
		var ret []string
		return ret
	}
	return o.ValidationPolicies
}

// GetValidationPoliciesOk returns a tuple with the ValidationPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptStage) GetValidationPoliciesOk() ([]string, bool) {
	if o == nil || o.ValidationPolicies == nil {
		return nil, false
	}
	return o.ValidationPolicies, true
}

// HasValidationPolicies returns a boolean if a field has been set.
func (o *PromptStage) HasValidationPolicies() bool {
	if o != nil && o.ValidationPolicies != nil {
		return true
	}

	return false
}

// SetValidationPolicies gets a reference to the given []string and assigns it to the ValidationPolicies field.
func (o *PromptStage) SetValidationPolicies(v []string) {
	o.ValidationPolicies = v
}

func (o PromptStage) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["fields"] = o.Fields
	}
	if o.ValidationPolicies != nil {
		toSerialize["validation_policies"] = o.ValidationPolicies
	}
	return json.Marshal(toSerialize)
}

type NullablePromptStage struct {
	value *PromptStage
	isSet bool
}

func (v NullablePromptStage) Get() *PromptStage {
	return v.value
}

func (v *NullablePromptStage) Set(val *PromptStage) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptStage) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptStage(val *PromptStage) *NullablePromptStage {
	return &NullablePromptStage{value: val, isSet: true}
}

func (v NullablePromptStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
