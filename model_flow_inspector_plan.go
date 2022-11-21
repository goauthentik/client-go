/*
authentik

Making authentication simple.

API version: 2022.11.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FlowInspectorPlan Serializer for an active FlowPlan
type FlowInspectorPlan struct {
	CurrentStage     FlowInspectorPlanCurrentStage `json:"current_stage"`
	NextPlannedStage FlowInspectorPlanCurrentStage `json:"next_planned_stage"`
	PlanContext      map[string]interface{}        `json:"plan_context"`
	SessionId        string                        `json:"session_id"`
}

// NewFlowInspectorPlan instantiates a new FlowInspectorPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowInspectorPlan(currentStage FlowInspectorPlanCurrentStage, nextPlannedStage FlowInspectorPlanCurrentStage, planContext map[string]interface{}, sessionId string) *FlowInspectorPlan {
	this := FlowInspectorPlan{}
	this.CurrentStage = currentStage
	this.NextPlannedStage = nextPlannedStage
	this.PlanContext = planContext
	this.SessionId = sessionId
	return &this
}

// NewFlowInspectorPlanWithDefaults instantiates a new FlowInspectorPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowInspectorPlanWithDefaults() *FlowInspectorPlan {
	this := FlowInspectorPlan{}
	return &this
}

// GetCurrentStage returns the CurrentStage field value
func (o *FlowInspectorPlan) GetCurrentStage() FlowInspectorPlanCurrentStage {
	if o == nil {
		var ret FlowInspectorPlanCurrentStage
		return ret
	}

	return o.CurrentStage
}

// GetCurrentStageOk returns a tuple with the CurrentStage field value
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlan) GetCurrentStageOk() (*FlowInspectorPlanCurrentStage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentStage, true
}

// SetCurrentStage sets field value
func (o *FlowInspectorPlan) SetCurrentStage(v FlowInspectorPlanCurrentStage) {
	o.CurrentStage = v
}

// GetNextPlannedStage returns the NextPlannedStage field value
func (o *FlowInspectorPlan) GetNextPlannedStage() FlowInspectorPlanCurrentStage {
	if o == nil {
		var ret FlowInspectorPlanCurrentStage
		return ret
	}

	return o.NextPlannedStage
}

// GetNextPlannedStageOk returns a tuple with the NextPlannedStage field value
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlan) GetNextPlannedStageOk() (*FlowInspectorPlanCurrentStage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPlannedStage, true
}

// SetNextPlannedStage sets field value
func (o *FlowInspectorPlan) SetNextPlannedStage(v FlowInspectorPlanCurrentStage) {
	o.NextPlannedStage = v
}

// GetPlanContext returns the PlanContext field value
func (o *FlowInspectorPlan) GetPlanContext() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.PlanContext
}

// GetPlanContextOk returns a tuple with the PlanContext field value
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlan) GetPlanContextOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanContext, true
}

// SetPlanContext sets field value
func (o *FlowInspectorPlan) SetPlanContext(v map[string]interface{}) {
	o.PlanContext = v
}

// GetSessionId returns the SessionId field value
func (o *FlowInspectorPlan) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlan) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *FlowInspectorPlan) SetSessionId(v string) {
	o.SessionId = v
}

func (o FlowInspectorPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["current_stage"] = o.CurrentStage
	}
	if true {
		toSerialize["next_planned_stage"] = o.NextPlannedStage
	}
	if true {
		toSerialize["plan_context"] = o.PlanContext
	}
	if true {
		toSerialize["session_id"] = o.SessionId
	}
	return json.Marshal(toSerialize)
}

type NullableFlowInspectorPlan struct {
	value *FlowInspectorPlan
	isSet bool
}

func (v NullableFlowInspectorPlan) Get() *FlowInspectorPlan {
	return v.value
}

func (v *NullableFlowInspectorPlan) Set(val *FlowInspectorPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowInspectorPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowInspectorPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowInspectorPlan(val *FlowInspectorPlan) *NullableFlowInspectorPlan {
	return &NullableFlowInspectorPlan{value: val, isSet: true}
}

func (v NullableFlowInspectorPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowInspectorPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
