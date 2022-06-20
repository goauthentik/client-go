/*
authentik

Making authentication simple.

API version: 2022.6.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FlowInspectorPlanCurrentStage struct for FlowInspectorPlanCurrentStage
type FlowInspectorPlanCurrentStage struct {
	Pk                      string                   `json:"pk"`
	PolicybindingmodelPtrId string                   `json:"policybindingmodel_ptr_id"`
	Target                  string                   `json:"target"`
	Stage                   string                   `json:"stage"`
	StageObj                FlowStageBindingStageObj `json:"stage_obj"`
	// Evaluate policies during the Flow planning process. Disable this for input-based policies.
	EvaluateOnPlan *bool `json:"evaluate_on_plan,omitempty"`
	// Evaluate policies when the Stage is present to the user.
	ReEvaluatePolicies *bool             `json:"re_evaluate_policies,omitempty"`
	Order              int32             `json:"order"`
	PolicyEngineMode   *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// Configure how the flow executor should handle an invalid response to a challenge. RETRY returns the error message and a similar challenge to the executor. RESTART restarts the flow from the beginning, and RESTART_WITH_CONTEXT restarts the flow while keeping the current context.
	InvalidResponseAction NullableInvalidResponseActionEnum `json:"invalid_response_action,omitempty"`
}

// NewFlowInspectorPlanCurrentStage instantiates a new FlowInspectorPlanCurrentStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowInspectorPlanCurrentStage(pk string, policybindingmodelPtrId string, target string, stage string, stageObj FlowStageBindingStageObj, order int32) *FlowInspectorPlanCurrentStage {
	this := FlowInspectorPlanCurrentStage{}
	this.Pk = pk
	this.PolicybindingmodelPtrId = policybindingmodelPtrId
	this.Target = target
	this.Stage = stage
	this.StageObj = stageObj
	this.Order = order
	return &this
}

// NewFlowInspectorPlanCurrentStageWithDefaults instantiates a new FlowInspectorPlanCurrentStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowInspectorPlanCurrentStageWithDefaults() *FlowInspectorPlanCurrentStage {
	this := FlowInspectorPlanCurrentStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *FlowInspectorPlanCurrentStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlanCurrentStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *FlowInspectorPlanCurrentStage) SetPk(v string) {
	o.Pk = v
}

// GetPolicybindingmodelPtrId returns the PolicybindingmodelPtrId field value
func (o *FlowInspectorPlanCurrentStage) GetPolicybindingmodelPtrId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicybindingmodelPtrId
}

// GetPolicybindingmodelPtrIdOk returns a tuple with the PolicybindingmodelPtrId field value
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlanCurrentStage) GetPolicybindingmodelPtrIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicybindingmodelPtrId, true
}

// SetPolicybindingmodelPtrId sets field value
func (o *FlowInspectorPlanCurrentStage) SetPolicybindingmodelPtrId(v string) {
	o.PolicybindingmodelPtrId = v
}

// GetTarget returns the Target field value
func (o *FlowInspectorPlanCurrentStage) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlanCurrentStage) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *FlowInspectorPlanCurrentStage) SetTarget(v string) {
	o.Target = v
}

// GetStage returns the Stage field value
func (o *FlowInspectorPlanCurrentStage) GetStage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlanCurrentStage) GetStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stage, true
}

// SetStage sets field value
func (o *FlowInspectorPlanCurrentStage) SetStage(v string) {
	o.Stage = v
}

// GetStageObj returns the StageObj field value
func (o *FlowInspectorPlanCurrentStage) GetStageObj() FlowStageBindingStageObj {
	if o == nil {
		var ret FlowStageBindingStageObj
		return ret
	}

	return o.StageObj
}

// GetStageObjOk returns a tuple with the StageObj field value
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlanCurrentStage) GetStageObjOk() (*FlowStageBindingStageObj, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StageObj, true
}

// SetStageObj sets field value
func (o *FlowInspectorPlanCurrentStage) SetStageObj(v FlowStageBindingStageObj) {
	o.StageObj = v
}

// GetEvaluateOnPlan returns the EvaluateOnPlan field value if set, zero value otherwise.
func (o *FlowInspectorPlanCurrentStage) GetEvaluateOnPlan() bool {
	if o == nil || o.EvaluateOnPlan == nil {
		var ret bool
		return ret
	}
	return *o.EvaluateOnPlan
}

// GetEvaluateOnPlanOk returns a tuple with the EvaluateOnPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlanCurrentStage) GetEvaluateOnPlanOk() (*bool, bool) {
	if o == nil || o.EvaluateOnPlan == nil {
		return nil, false
	}
	return o.EvaluateOnPlan, true
}

// HasEvaluateOnPlan returns a boolean if a field has been set.
func (o *FlowInspectorPlanCurrentStage) HasEvaluateOnPlan() bool {
	if o != nil && o.EvaluateOnPlan != nil {
		return true
	}

	return false
}

// SetEvaluateOnPlan gets a reference to the given bool and assigns it to the EvaluateOnPlan field.
func (o *FlowInspectorPlanCurrentStage) SetEvaluateOnPlan(v bool) {
	o.EvaluateOnPlan = &v
}

// GetReEvaluatePolicies returns the ReEvaluatePolicies field value if set, zero value otherwise.
func (o *FlowInspectorPlanCurrentStage) GetReEvaluatePolicies() bool {
	if o == nil || o.ReEvaluatePolicies == nil {
		var ret bool
		return ret
	}
	return *o.ReEvaluatePolicies
}

// GetReEvaluatePoliciesOk returns a tuple with the ReEvaluatePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlanCurrentStage) GetReEvaluatePoliciesOk() (*bool, bool) {
	if o == nil || o.ReEvaluatePolicies == nil {
		return nil, false
	}
	return o.ReEvaluatePolicies, true
}

// HasReEvaluatePolicies returns a boolean if a field has been set.
func (o *FlowInspectorPlanCurrentStage) HasReEvaluatePolicies() bool {
	if o != nil && o.ReEvaluatePolicies != nil {
		return true
	}

	return false
}

// SetReEvaluatePolicies gets a reference to the given bool and assigns it to the ReEvaluatePolicies field.
func (o *FlowInspectorPlanCurrentStage) SetReEvaluatePolicies(v bool) {
	o.ReEvaluatePolicies = &v
}

// GetOrder returns the Order field value
func (o *FlowInspectorPlanCurrentStage) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlanCurrentStage) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *FlowInspectorPlanCurrentStage) SetOrder(v int32) {
	o.Order = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *FlowInspectorPlanCurrentStage) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInspectorPlanCurrentStage) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *FlowInspectorPlanCurrentStage) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *FlowInspectorPlanCurrentStage) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetInvalidResponseAction returns the InvalidResponseAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowInspectorPlanCurrentStage) GetInvalidResponseAction() InvalidResponseActionEnum {
	if o == nil || o.InvalidResponseAction.Get() == nil {
		var ret InvalidResponseActionEnum
		return ret
	}
	return *o.InvalidResponseAction.Get()
}

// GetInvalidResponseActionOk returns a tuple with the InvalidResponseAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowInspectorPlanCurrentStage) GetInvalidResponseActionOk() (*InvalidResponseActionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvalidResponseAction.Get(), o.InvalidResponseAction.IsSet()
}

// HasInvalidResponseAction returns a boolean if a field has been set.
func (o *FlowInspectorPlanCurrentStage) HasInvalidResponseAction() bool {
	if o != nil && o.InvalidResponseAction.IsSet() {
		return true
	}

	return false
}

// SetInvalidResponseAction gets a reference to the given NullableInvalidResponseActionEnum and assigns it to the InvalidResponseAction field.
func (o *FlowInspectorPlanCurrentStage) SetInvalidResponseAction(v InvalidResponseActionEnum) {
	o.InvalidResponseAction.Set(&v)
}

// SetInvalidResponseActionNil sets the value for InvalidResponseAction to be an explicit nil
func (o *FlowInspectorPlanCurrentStage) SetInvalidResponseActionNil() {
	o.InvalidResponseAction.Set(nil)
}

// UnsetInvalidResponseAction ensures that no value is present for InvalidResponseAction, not even an explicit nil
func (o *FlowInspectorPlanCurrentStage) UnsetInvalidResponseAction() {
	o.InvalidResponseAction.Unset()
}

func (o FlowInspectorPlanCurrentStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["policybindingmodel_ptr_id"] = o.PolicybindingmodelPtrId
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["stage"] = o.Stage
	}
	if true {
		toSerialize["stage_obj"] = o.StageObj
	}
	if o.EvaluateOnPlan != nil {
		toSerialize["evaluate_on_plan"] = o.EvaluateOnPlan
	}
	if o.ReEvaluatePolicies != nil {
		toSerialize["re_evaluate_policies"] = o.ReEvaluatePolicies
	}
	if true {
		toSerialize["order"] = o.Order
	}
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.InvalidResponseAction.IsSet() {
		toSerialize["invalid_response_action"] = o.InvalidResponseAction.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFlowInspectorPlanCurrentStage struct {
	value *FlowInspectorPlanCurrentStage
	isSet bool
}

func (v NullableFlowInspectorPlanCurrentStage) Get() *FlowInspectorPlanCurrentStage {
	return v.value
}

func (v *NullableFlowInspectorPlanCurrentStage) Set(val *FlowInspectorPlanCurrentStage) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowInspectorPlanCurrentStage) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowInspectorPlanCurrentStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowInspectorPlanCurrentStage(val *FlowInspectorPlanCurrentStage) *NullableFlowInspectorPlanCurrentStage {
	return &NullableFlowInspectorPlanCurrentStage{value: val, isSet: true}
}

func (v NullableFlowInspectorPlanCurrentStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowInspectorPlanCurrentStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
