/*
authentik

Making authentication simple.

API version: 2023.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PolicyBindingPolicyObj struct for PolicyBindingPolicyObj
type PolicyBindingPolicyObj struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// When this option is enabled, all executions of this policy will be logged. By default, only execution errors are logged.
	ExecutionLogging *bool `json:"execution_logging,omitempty"`
	// Get object component so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string `json:"meta_model_name"`
	// Return objects policy is bound to
	BoundTo int32 `json:"bound_to"`
}

// NewPolicyBindingPolicyObj instantiates a new PolicyBindingPolicyObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyBindingPolicyObj(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32) *PolicyBindingPolicyObj {
	this := PolicyBindingPolicyObj{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.BoundTo = boundTo
	return &this
}

// NewPolicyBindingPolicyObjWithDefaults instantiates a new PolicyBindingPolicyObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyBindingPolicyObjWithDefaults() *PolicyBindingPolicyObj {
	this := PolicyBindingPolicyObj{}
	return &this
}

// GetPk returns the Pk field value
func (o *PolicyBindingPolicyObj) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingPolicyObj) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *PolicyBindingPolicyObj) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *PolicyBindingPolicyObj) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingPolicyObj) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PolicyBindingPolicyObj) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *PolicyBindingPolicyObj) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyBindingPolicyObj) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *PolicyBindingPolicyObj) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *PolicyBindingPolicyObj) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetComponent returns the Component field value
func (o *PolicyBindingPolicyObj) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingPolicyObj) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *PolicyBindingPolicyObj) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *PolicyBindingPolicyObj) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingPolicyObj) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *PolicyBindingPolicyObj) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *PolicyBindingPolicyObj) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingPolicyObj) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *PolicyBindingPolicyObj) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *PolicyBindingPolicyObj) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingPolicyObj) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *PolicyBindingPolicyObj) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetBoundTo returns the BoundTo field value
func (o *PolicyBindingPolicyObj) GetBoundTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BoundTo
}

// GetBoundToOk returns a tuple with the BoundTo field value
// and a boolean to check if the value has been set.
func (o *PolicyBindingPolicyObj) GetBoundToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundTo, true
}

// SetBoundTo sets field value
func (o *PolicyBindingPolicyObj) SetBoundTo(v int32) {
	o.BoundTo = v
}

func (o PolicyBindingPolicyObj) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExecutionLogging != nil {
		toSerialize["execution_logging"] = o.ExecutionLogging
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
	if true {
		toSerialize["bound_to"] = o.BoundTo
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyBindingPolicyObj struct {
	value *PolicyBindingPolicyObj
	isSet bool
}

func (v NullablePolicyBindingPolicyObj) Get() *PolicyBindingPolicyObj {
	return v.value
}

func (v *NullablePolicyBindingPolicyObj) Set(val *PolicyBindingPolicyObj) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyBindingPolicyObj) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyBindingPolicyObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyBindingPolicyObj(val *PolicyBindingPolicyObj) *NullablePolicyBindingPolicyObj {
	return &NullablePolicyBindingPolicyObj{value: val, isSet: true}
}

func (v NullablePolicyBindingPolicyObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyBindingPolicyObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
