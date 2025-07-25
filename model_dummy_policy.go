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

// DummyPolicy Dummy Policy Serializer
type DummyPolicy struct {
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
	BoundTo int32  `json:"bound_to"`
	Result  *bool  `json:"result,omitempty"`
	WaitMin *int32 `json:"wait_min,omitempty"`
	WaitMax *int32 `json:"wait_max,omitempty"`
}

// NewDummyPolicy instantiates a new DummyPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDummyPolicy(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, boundTo int32) *DummyPolicy {
	this := DummyPolicy{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.BoundTo = boundTo
	return &this
}

// NewDummyPolicyWithDefaults instantiates a new DummyPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDummyPolicyWithDefaults() *DummyPolicy {
	this := DummyPolicy{}
	return &this
}

// GetPk returns the Pk field value
func (o *DummyPolicy) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *DummyPolicy) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *DummyPolicy) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *DummyPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DummyPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DummyPolicy) SetName(v string) {
	o.Name = v
}

// GetExecutionLogging returns the ExecutionLogging field value if set, zero value otherwise.
func (o *DummyPolicy) GetExecutionLogging() bool {
	if o == nil || o.ExecutionLogging == nil {
		var ret bool
		return ret
	}
	return *o.ExecutionLogging
}

// GetExecutionLoggingOk returns a tuple with the ExecutionLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicy) GetExecutionLoggingOk() (*bool, bool) {
	if o == nil || o.ExecutionLogging == nil {
		return nil, false
	}
	return o.ExecutionLogging, true
}

// HasExecutionLogging returns a boolean if a field has been set.
func (o *DummyPolicy) HasExecutionLogging() bool {
	if o != nil && o.ExecutionLogging != nil {
		return true
	}

	return false
}

// SetExecutionLogging gets a reference to the given bool and assigns it to the ExecutionLogging field.
func (o *DummyPolicy) SetExecutionLogging(v bool) {
	o.ExecutionLogging = &v
}

// GetComponent returns the Component field value
func (o *DummyPolicy) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DummyPolicy) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *DummyPolicy) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *DummyPolicy) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *DummyPolicy) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *DummyPolicy) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *DummyPolicy) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *DummyPolicy) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *DummyPolicy) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *DummyPolicy) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *DummyPolicy) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *DummyPolicy) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetBoundTo returns the BoundTo field value
func (o *DummyPolicy) GetBoundTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BoundTo
}

// GetBoundToOk returns a tuple with the BoundTo field value
// and a boolean to check if the value has been set.
func (o *DummyPolicy) GetBoundToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundTo, true
}

// SetBoundTo sets field value
func (o *DummyPolicy) SetBoundTo(v int32) {
	o.BoundTo = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DummyPolicy) GetResult() bool {
	if o == nil || o.Result == nil {
		var ret bool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicy) GetResultOk() (*bool, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DummyPolicy) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given bool and assigns it to the Result field.
func (o *DummyPolicy) SetResult(v bool) {
	o.Result = &v
}

// GetWaitMin returns the WaitMin field value if set, zero value otherwise.
func (o *DummyPolicy) GetWaitMin() int32 {
	if o == nil || o.WaitMin == nil {
		var ret int32
		return ret
	}
	return *o.WaitMin
}

// GetWaitMinOk returns a tuple with the WaitMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicy) GetWaitMinOk() (*int32, bool) {
	if o == nil || o.WaitMin == nil {
		return nil, false
	}
	return o.WaitMin, true
}

// HasWaitMin returns a boolean if a field has been set.
func (o *DummyPolicy) HasWaitMin() bool {
	if o != nil && o.WaitMin != nil {
		return true
	}

	return false
}

// SetWaitMin gets a reference to the given int32 and assigns it to the WaitMin field.
func (o *DummyPolicy) SetWaitMin(v int32) {
	o.WaitMin = &v
}

// GetWaitMax returns the WaitMax field value if set, zero value otherwise.
func (o *DummyPolicy) GetWaitMax() int32 {
	if o == nil || o.WaitMax == nil {
		var ret int32
		return ret
	}
	return *o.WaitMax
}

// GetWaitMaxOk returns a tuple with the WaitMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyPolicy) GetWaitMaxOk() (*int32, bool) {
	if o == nil || o.WaitMax == nil {
		return nil, false
	}
	return o.WaitMax, true
}

// HasWaitMax returns a boolean if a field has been set.
func (o *DummyPolicy) HasWaitMax() bool {
	if o != nil && o.WaitMax != nil {
		return true
	}

	return false
}

// SetWaitMax gets a reference to the given int32 and assigns it to the WaitMax field.
func (o *DummyPolicy) SetWaitMax(v int32) {
	o.WaitMax = &v
}

func (o DummyPolicy) MarshalJSON() ([]byte, error) {
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
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.WaitMin != nil {
		toSerialize["wait_min"] = o.WaitMin
	}
	if o.WaitMax != nil {
		toSerialize["wait_max"] = o.WaitMax
	}
	return json.Marshal(toSerialize)
}

type NullableDummyPolicy struct {
	value *DummyPolicy
	isSet bool
}

func (v NullableDummyPolicy) Get() *DummyPolicy {
	return v.value
}

func (v *NullableDummyPolicy) Set(val *DummyPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDummyPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDummyPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDummyPolicy(val *DummyPolicy) *NullableDummyPolicy {
	return &NullableDummyPolicy{value: val, isSet: true}
}

func (v NullableDummyPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDummyPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
