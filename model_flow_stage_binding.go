/*
authentik

Making authentication simple.

API version: 2023.5.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FlowStageBinding FlowStageBinding Serializer
type FlowStageBinding struct {
	Pk                      string `json:"pk"`
	PolicybindingmodelPtrId string `json:"policybindingmodel_ptr_id"`
	Target                  string `json:"target"`
	Stage                   string `json:"stage"`
	StageObj                Stage  `json:"stage_obj"`
	// Evaluate policies during the Flow planning process.
	EvaluateOnPlan *bool `json:"evaluate_on_plan,omitempty"`
	// Evaluate policies when the Stage is present to the user.
	ReEvaluatePolicies *bool             `json:"re_evaluate_policies,omitempty"`
	Order              int32             `json:"order"`
	PolicyEngineMode   *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// Configure how the flow executor should handle an invalid response to a challenge. RETRY returns the error message and a similar challenge to the executor. RESTART restarts the flow from the beginning, and RESTART_WITH_CONTEXT restarts the flow while keeping the current context.  * `retry` - Retry * `restart` - Restart * `restart_with_context` - Restart With Context
	InvalidResponseAction *InvalidResponseActionEnum `json:"invalid_response_action,omitempty"`
}

// NewFlowStageBinding instantiates a new FlowStageBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowStageBinding(pk string, policybindingmodelPtrId string, target string, stage string, stageObj Stage, order int32) *FlowStageBinding {
	this := FlowStageBinding{}
	this.Pk = pk
	this.PolicybindingmodelPtrId = policybindingmodelPtrId
	this.Target = target
	this.Stage = stage
	this.StageObj = stageObj
	this.Order = order
	return &this
}

// NewFlowStageBindingWithDefaults instantiates a new FlowStageBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowStageBindingWithDefaults() *FlowStageBinding {
	this := FlowStageBinding{}
	return &this
}

// GetPk returns the Pk field value
func (o *FlowStageBinding) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *FlowStageBinding) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *FlowStageBinding) SetPk(v string) {
	o.Pk = v
}

// GetPolicybindingmodelPtrId returns the PolicybindingmodelPtrId field value
func (o *FlowStageBinding) GetPolicybindingmodelPtrId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicybindingmodelPtrId
}

// GetPolicybindingmodelPtrIdOk returns a tuple with the PolicybindingmodelPtrId field value
// and a boolean to check if the value has been set.
func (o *FlowStageBinding) GetPolicybindingmodelPtrIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicybindingmodelPtrId, true
}

// SetPolicybindingmodelPtrId sets field value
func (o *FlowStageBinding) SetPolicybindingmodelPtrId(v string) {
	o.PolicybindingmodelPtrId = v
}

// GetTarget returns the Target field value
func (o *FlowStageBinding) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *FlowStageBinding) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *FlowStageBinding) SetTarget(v string) {
	o.Target = v
}

// GetStage returns the Stage field value
func (o *FlowStageBinding) GetStage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value
// and a boolean to check if the value has been set.
func (o *FlowStageBinding) GetStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stage, true
}

// SetStage sets field value
func (o *FlowStageBinding) SetStage(v string) {
	o.Stage = v
}

// GetStageObj returns the StageObj field value
func (o *FlowStageBinding) GetStageObj() Stage {
	if o == nil {
		var ret Stage
		return ret
	}

	return o.StageObj
}

// GetStageObjOk returns a tuple with the StageObj field value
// and a boolean to check if the value has been set.
func (o *FlowStageBinding) GetStageObjOk() (*Stage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StageObj, true
}

// SetStageObj sets field value
func (o *FlowStageBinding) SetStageObj(v Stage) {
	o.StageObj = v
}

// GetEvaluateOnPlan returns the EvaluateOnPlan field value if set, zero value otherwise.
func (o *FlowStageBinding) GetEvaluateOnPlan() bool {
	if o == nil || o.EvaluateOnPlan == nil {
		var ret bool
		return ret
	}
	return *o.EvaluateOnPlan
}

// GetEvaluateOnPlanOk returns a tuple with the EvaluateOnPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStageBinding) GetEvaluateOnPlanOk() (*bool, bool) {
	if o == nil || o.EvaluateOnPlan == nil {
		return nil, false
	}
	return o.EvaluateOnPlan, true
}

// HasEvaluateOnPlan returns a boolean if a field has been set.
func (o *FlowStageBinding) HasEvaluateOnPlan() bool {
	if o != nil && o.EvaluateOnPlan != nil {
		return true
	}

	return false
}

// SetEvaluateOnPlan gets a reference to the given bool and assigns it to the EvaluateOnPlan field.
func (o *FlowStageBinding) SetEvaluateOnPlan(v bool) {
	o.EvaluateOnPlan = &v
}

// GetReEvaluatePolicies returns the ReEvaluatePolicies field value if set, zero value otherwise.
func (o *FlowStageBinding) GetReEvaluatePolicies() bool {
	if o == nil || o.ReEvaluatePolicies == nil {
		var ret bool
		return ret
	}
	return *o.ReEvaluatePolicies
}

// GetReEvaluatePoliciesOk returns a tuple with the ReEvaluatePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStageBinding) GetReEvaluatePoliciesOk() (*bool, bool) {
	if o == nil || o.ReEvaluatePolicies == nil {
		return nil, false
	}
	return o.ReEvaluatePolicies, true
}

// HasReEvaluatePolicies returns a boolean if a field has been set.
func (o *FlowStageBinding) HasReEvaluatePolicies() bool {
	if o != nil && o.ReEvaluatePolicies != nil {
		return true
	}

	return false
}

// SetReEvaluatePolicies gets a reference to the given bool and assigns it to the ReEvaluatePolicies field.
func (o *FlowStageBinding) SetReEvaluatePolicies(v bool) {
	o.ReEvaluatePolicies = &v
}

// GetOrder returns the Order field value
func (o *FlowStageBinding) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *FlowStageBinding) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *FlowStageBinding) SetOrder(v int32) {
	o.Order = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *FlowStageBinding) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStageBinding) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *FlowStageBinding) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *FlowStageBinding) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetInvalidResponseAction returns the InvalidResponseAction field value if set, zero value otherwise.
func (o *FlowStageBinding) GetInvalidResponseAction() InvalidResponseActionEnum {
	if o == nil || o.InvalidResponseAction == nil {
		var ret InvalidResponseActionEnum
		return ret
	}
	return *o.InvalidResponseAction
}

// GetInvalidResponseActionOk returns a tuple with the InvalidResponseAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowStageBinding) GetInvalidResponseActionOk() (*InvalidResponseActionEnum, bool) {
	if o == nil || o.InvalidResponseAction == nil {
		return nil, false
	}
	return o.InvalidResponseAction, true
}

// HasInvalidResponseAction returns a boolean if a field has been set.
func (o *FlowStageBinding) HasInvalidResponseAction() bool {
	if o != nil && o.InvalidResponseAction != nil {
		return true
	}

	return false
}

// SetInvalidResponseAction gets a reference to the given InvalidResponseActionEnum and assigns it to the InvalidResponseAction field.
func (o *FlowStageBinding) SetInvalidResponseAction(v InvalidResponseActionEnum) {
	o.InvalidResponseAction = &v
}

func (o FlowStageBinding) MarshalJSON() ([]byte, error) {
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
	if o.InvalidResponseAction != nil {
		toSerialize["invalid_response_action"] = o.InvalidResponseAction
	}
	return json.Marshal(toSerialize)
}

type NullableFlowStageBinding struct {
	value *FlowStageBinding
	isSet bool
}

func (v NullableFlowStageBinding) Get() *FlowStageBinding {
	return v.value
}

func (v *NullableFlowStageBinding) Set(val *FlowStageBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowStageBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowStageBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowStageBinding(val *FlowStageBinding) *NullableFlowStageBinding {
	return &NullableFlowStageBinding{value: val, isSet: true}
}

func (v NullableFlowStageBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowStageBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
