/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RedirectStage RedirectStage Serializer
type RedirectStage struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// Get object type so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string                `json:"meta_model_name"`
	FlowSet       []FlowSet             `json:"flow_set,omitempty"`
	KeepContext   *bool                 `json:"keep_context,omitempty"`
	Mode          RedirectStageModeEnum `json:"mode"`
	TargetStatic  *string               `json:"target_static,omitempty"`
	TargetFlow    NullableString        `json:"target_flow,omitempty"`
}

// NewRedirectStage instantiates a new RedirectStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, mode RedirectStageModeEnum) *RedirectStage {
	this := RedirectStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Mode = mode
	return &this
}

// NewRedirectStageWithDefaults instantiates a new RedirectStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectStageWithDefaults() *RedirectStage {
	this := RedirectStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *RedirectStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *RedirectStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *RedirectStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *RedirectStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RedirectStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RedirectStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *RedirectStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *RedirectStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *RedirectStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *RedirectStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *RedirectStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *RedirectStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *RedirectStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *RedirectStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *RedirectStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *RedirectStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *RedirectStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *RedirectStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *RedirectStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *RedirectStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *RedirectStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetKeepContext returns the KeepContext field value if set, zero value otherwise.
func (o *RedirectStage) GetKeepContext() bool {
	if o == nil || o.KeepContext == nil {
		var ret bool
		return ret
	}
	return *o.KeepContext
}

// GetKeepContextOk returns a tuple with the KeepContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectStage) GetKeepContextOk() (*bool, bool) {
	if o == nil || o.KeepContext == nil {
		return nil, false
	}
	return o.KeepContext, true
}

// HasKeepContext returns a boolean if a field has been set.
func (o *RedirectStage) HasKeepContext() bool {
	if o != nil && o.KeepContext != nil {
		return true
	}

	return false
}

// SetKeepContext gets a reference to the given bool and assigns it to the KeepContext field.
func (o *RedirectStage) SetKeepContext(v bool) {
	o.KeepContext = &v
}

// GetMode returns the Mode field value
func (o *RedirectStage) GetMode() RedirectStageModeEnum {
	if o == nil {
		var ret RedirectStageModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *RedirectStage) GetModeOk() (*RedirectStageModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *RedirectStage) SetMode(v RedirectStageModeEnum) {
	o.Mode = v
}

// GetTargetStatic returns the TargetStatic field value if set, zero value otherwise.
func (o *RedirectStage) GetTargetStatic() string {
	if o == nil || o.TargetStatic == nil {
		var ret string
		return ret
	}
	return *o.TargetStatic
}

// GetTargetStaticOk returns a tuple with the TargetStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectStage) GetTargetStaticOk() (*string, bool) {
	if o == nil || o.TargetStatic == nil {
		return nil, false
	}
	return o.TargetStatic, true
}

// HasTargetStatic returns a boolean if a field has been set.
func (o *RedirectStage) HasTargetStatic() bool {
	if o != nil && o.TargetStatic != nil {
		return true
	}

	return false
}

// SetTargetStatic gets a reference to the given string and assigns it to the TargetStatic field.
func (o *RedirectStage) SetTargetStatic(v string) {
	o.TargetStatic = &v
}

// GetTargetFlow returns the TargetFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RedirectStage) GetTargetFlow() string {
	if o == nil || o.TargetFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetFlow.Get()
}

// GetTargetFlowOk returns a tuple with the TargetFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RedirectStage) GetTargetFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetFlow.Get(), o.TargetFlow.IsSet()
}

// HasTargetFlow returns a boolean if a field has been set.
func (o *RedirectStage) HasTargetFlow() bool {
	if o != nil && o.TargetFlow.IsSet() {
		return true
	}

	return false
}

// SetTargetFlow gets a reference to the given NullableString and assigns it to the TargetFlow field.
func (o *RedirectStage) SetTargetFlow(v string) {
	o.TargetFlow.Set(&v)
}

// SetTargetFlowNil sets the value for TargetFlow to be an explicit nil
func (o *RedirectStage) SetTargetFlowNil() {
	o.TargetFlow.Set(nil)
}

// UnsetTargetFlow ensures that no value is present for TargetFlow, not even an explicit nil
func (o *RedirectStage) UnsetTargetFlow() {
	o.TargetFlow.Unset()
}

func (o RedirectStage) MarshalJSON() ([]byte, error) {
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
	if o.KeepContext != nil {
		toSerialize["keep_context"] = o.KeepContext
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.TargetStatic != nil {
		toSerialize["target_static"] = o.TargetStatic
	}
	if o.TargetFlow.IsSet() {
		toSerialize["target_flow"] = o.TargetFlow.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRedirectStage struct {
	value *RedirectStage
	isSet bool
}

func (v NullableRedirectStage) Get() *RedirectStage {
	return v.value
}

func (v *NullableRedirectStage) Set(val *RedirectStage) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectStage) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectStage(val *RedirectStage) *NullableRedirectStage {
	return &NullableRedirectStage{value: val, isSet: true}
}

func (v NullableRedirectStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
