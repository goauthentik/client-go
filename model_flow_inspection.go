/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the FlowInspection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowInspection{}

// FlowInspection Serializer for inspect endpoint
type FlowInspection struct {
	Plans       []FlowInspectorPlan `json:"plans"`
	CurrentPlan *FlowInspectorPlan  `json:"current_plan,omitempty"`
	IsCompleted bool                `json:"is_completed"`
}

type _FlowInspection FlowInspection

// NewFlowInspection instantiates a new FlowInspection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowInspection(plans []FlowInspectorPlan, isCompleted bool) *FlowInspection {
	this := FlowInspection{}
	this.Plans = plans
	this.IsCompleted = isCompleted
	return &this
}

// NewFlowInspectionWithDefaults instantiates a new FlowInspection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowInspectionWithDefaults() *FlowInspection {
	this := FlowInspection{}
	return &this
}

// GetPlans returns the Plans field value
func (o *FlowInspection) GetPlans() []FlowInspectorPlan {
	if o == nil {
		var ret []FlowInspectorPlan
		return ret
	}

	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value
// and a boolean to check if the value has been set.
func (o *FlowInspection) GetPlansOk() ([]FlowInspectorPlan, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plans, true
}

// SetPlans sets field value
func (o *FlowInspection) SetPlans(v []FlowInspectorPlan) {
	o.Plans = v
}

// GetCurrentPlan returns the CurrentPlan field value if set, zero value otherwise.
func (o *FlowInspection) GetCurrentPlan() FlowInspectorPlan {
	if o == nil || IsNil(o.CurrentPlan) {
		var ret FlowInspectorPlan
		return ret
	}
	return *o.CurrentPlan
}

// GetCurrentPlanOk returns a tuple with the CurrentPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInspection) GetCurrentPlanOk() (*FlowInspectorPlan, bool) {
	if o == nil || IsNil(o.CurrentPlan) {
		return nil, false
	}
	return o.CurrentPlan, true
}

// HasCurrentPlan returns a boolean if a field has been set.
func (o *FlowInspection) HasCurrentPlan() bool {
	if o != nil && !IsNil(o.CurrentPlan) {
		return true
	}

	return false
}

// SetCurrentPlan gets a reference to the given FlowInspectorPlan and assigns it to the CurrentPlan field.
func (o *FlowInspection) SetCurrentPlan(v FlowInspectorPlan) {
	o.CurrentPlan = &v
}

// GetIsCompleted returns the IsCompleted field value
func (o *FlowInspection) GetIsCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value
// and a boolean to check if the value has been set.
func (o *FlowInspection) GetIsCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCompleted, true
}

// SetIsCompleted sets field value
func (o *FlowInspection) SetIsCompleted(v bool) {
	o.IsCompleted = v
}

func (o FlowInspection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowInspection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plans"] = o.Plans
	if !IsNil(o.CurrentPlan) {
		toSerialize["current_plan"] = o.CurrentPlan
	}
	toSerialize["is_completed"] = o.IsCompleted
	return toSerialize, nil
}

func (o *FlowInspection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plans",
		"is_completed",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFlowInspection := _FlowInspection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlowInspection)

	if err != nil {
		return err
	}

	*o = FlowInspection(varFlowInspection)

	return err
}

type NullableFlowInspection struct {
	value *FlowInspection
	isSet bool
}

func (v NullableFlowInspection) Get() *FlowInspection {
	return v.value
}

func (v *NullableFlowInspection) Set(val *FlowInspection) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowInspection) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowInspection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowInspection(val *FlowInspection) *NullableFlowInspection {
	return &NullableFlowInspection{value: val, isSet: true}
}

func (v NullableFlowInspection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowInspection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
