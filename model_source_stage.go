/*
authentik

Making authentication simple.

API version: 2024.6.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SourceStage SourceStage Serializer
type SourceStage struct {
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
	Source        string    `json:"source"`
	// Amount of time a user can take to return from the source to continue the flow (Format: hours=-1;minutes=-2;seconds=-3)
	ResumeTimeout *string `json:"resume_timeout,omitempty"`
}

// NewSourceStage instantiates a new SourceStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, source string) *SourceStage {
	this := SourceStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Source = source
	return &this
}

// NewSourceStageWithDefaults instantiates a new SourceStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceStageWithDefaults() *SourceStage {
	this := SourceStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *SourceStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *SourceStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *SourceStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *SourceStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *SourceStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *SourceStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *SourceStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *SourceStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *SourceStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *SourceStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *SourceStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *SourceStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *SourceStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *SourceStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *SourceStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *SourceStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *SourceStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *SourceStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *SourceStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetSource returns the Source field value
func (o *SourceStage) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SourceStage) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SourceStage) SetSource(v string) {
	o.Source = v
}

// GetResumeTimeout returns the ResumeTimeout field value if set, zero value otherwise.
func (o *SourceStage) GetResumeTimeout() string {
	if o == nil || o.ResumeTimeout == nil {
		var ret string
		return ret
	}
	return *o.ResumeTimeout
}

// GetResumeTimeoutOk returns a tuple with the ResumeTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceStage) GetResumeTimeoutOk() (*string, bool) {
	if o == nil || o.ResumeTimeout == nil {
		return nil, false
	}
	return o.ResumeTimeout, true
}

// HasResumeTimeout returns a boolean if a field has been set.
func (o *SourceStage) HasResumeTimeout() bool {
	if o != nil && o.ResumeTimeout != nil {
		return true
	}

	return false
}

// SetResumeTimeout gets a reference to the given string and assigns it to the ResumeTimeout field.
func (o *SourceStage) SetResumeTimeout(v string) {
	o.ResumeTimeout = &v
}

func (o SourceStage) MarshalJSON() ([]byte, error) {
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
		toSerialize["source"] = o.Source
	}
	if o.ResumeTimeout != nil {
		toSerialize["resume_timeout"] = o.ResumeTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableSourceStage struct {
	value *SourceStage
	isSet bool
}

func (v NullableSourceStage) Get() *SourceStage {
	return v.value
}

func (v *NullableSourceStage) Set(val *SourceStage) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceStage) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceStage(val *SourceStage) *NullableSourceStage {
	return &NullableSourceStage{value: val, isSet: true}
}

func (v NullableSourceStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
