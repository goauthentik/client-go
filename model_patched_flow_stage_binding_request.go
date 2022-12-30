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

// PatchedFlowStageBindingRequest FlowStageBinding Serializer
type PatchedFlowStageBindingRequest struct {
	Target *string `json:"target,omitempty"`
	Stage  *string `json:"stage,omitempty"`
	// Evaluate policies during the Flow planning process. Disable this for input-based policies.
	EvaluateOnPlan *bool `json:"evaluate_on_plan,omitempty"`
	// Evaluate policies when the Stage is present to the user.
	ReEvaluatePolicies *bool             `json:"re_evaluate_policies,omitempty"`
	Order              *int32            `json:"order,omitempty"`
	PolicyEngineMode   *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// Configure how the flow executor should handle an invalid response to a challenge. RETRY returns the error message and a similar challenge to the executor. RESTART restarts the flow from the beginning, and RESTART_WITH_CONTEXT restarts the flow while keeping the current context.
	InvalidResponseAction NullableInvalidResponseActionEnum `json:"invalid_response_action,omitempty"`
}

// NewPatchedFlowStageBindingRequest instantiates a new PatchedFlowStageBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedFlowStageBindingRequest() *PatchedFlowStageBindingRequest {
	this := PatchedFlowStageBindingRequest{}
	return &this
}

// NewPatchedFlowStageBindingRequestWithDefaults instantiates a new PatchedFlowStageBindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedFlowStageBindingRequestWithDefaults() *PatchedFlowStageBindingRequest {
	this := PatchedFlowStageBindingRequest{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *PatchedFlowStageBindingRequest) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowStageBindingRequest) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *PatchedFlowStageBindingRequest) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *PatchedFlowStageBindingRequest) SetTarget(v string) {
	o.Target = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *PatchedFlowStageBindingRequest) GetStage() string {
	if o == nil || o.Stage == nil {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowStageBindingRequest) GetStageOk() (*string, bool) {
	if o == nil || o.Stage == nil {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *PatchedFlowStageBindingRequest) HasStage() bool {
	if o != nil && o.Stage != nil {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *PatchedFlowStageBindingRequest) SetStage(v string) {
	o.Stage = &v
}

// GetEvaluateOnPlan returns the EvaluateOnPlan field value if set, zero value otherwise.
func (o *PatchedFlowStageBindingRequest) GetEvaluateOnPlan() bool {
	if o == nil || o.EvaluateOnPlan == nil {
		var ret bool
		return ret
	}
	return *o.EvaluateOnPlan
}

// GetEvaluateOnPlanOk returns a tuple with the EvaluateOnPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowStageBindingRequest) GetEvaluateOnPlanOk() (*bool, bool) {
	if o == nil || o.EvaluateOnPlan == nil {
		return nil, false
	}
	return o.EvaluateOnPlan, true
}

// HasEvaluateOnPlan returns a boolean if a field has been set.
func (o *PatchedFlowStageBindingRequest) HasEvaluateOnPlan() bool {
	if o != nil && o.EvaluateOnPlan != nil {
		return true
	}

	return false
}

// SetEvaluateOnPlan gets a reference to the given bool and assigns it to the EvaluateOnPlan field.
func (o *PatchedFlowStageBindingRequest) SetEvaluateOnPlan(v bool) {
	o.EvaluateOnPlan = &v
}

// GetReEvaluatePolicies returns the ReEvaluatePolicies field value if set, zero value otherwise.
func (o *PatchedFlowStageBindingRequest) GetReEvaluatePolicies() bool {
	if o == nil || o.ReEvaluatePolicies == nil {
		var ret bool
		return ret
	}
	return *o.ReEvaluatePolicies
}

// GetReEvaluatePoliciesOk returns a tuple with the ReEvaluatePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowStageBindingRequest) GetReEvaluatePoliciesOk() (*bool, bool) {
	if o == nil || o.ReEvaluatePolicies == nil {
		return nil, false
	}
	return o.ReEvaluatePolicies, true
}

// HasReEvaluatePolicies returns a boolean if a field has been set.
func (o *PatchedFlowStageBindingRequest) HasReEvaluatePolicies() bool {
	if o != nil && o.ReEvaluatePolicies != nil {
		return true
	}

	return false
}

// SetReEvaluatePolicies gets a reference to the given bool and assigns it to the ReEvaluatePolicies field.
func (o *PatchedFlowStageBindingRequest) SetReEvaluatePolicies(v bool) {
	o.ReEvaluatePolicies = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PatchedFlowStageBindingRequest) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowStageBindingRequest) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PatchedFlowStageBindingRequest) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *PatchedFlowStageBindingRequest) SetOrder(v int32) {
	o.Order = &v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *PatchedFlowStageBindingRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFlowStageBindingRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *PatchedFlowStageBindingRequest) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *PatchedFlowStageBindingRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetInvalidResponseAction returns the InvalidResponseAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedFlowStageBindingRequest) GetInvalidResponseAction() InvalidResponseActionEnum {
	if o == nil || o.InvalidResponseAction.Get() == nil {
		var ret InvalidResponseActionEnum
		return ret
	}
	return *o.InvalidResponseAction.Get()
}

// GetInvalidResponseActionOk returns a tuple with the InvalidResponseAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedFlowStageBindingRequest) GetInvalidResponseActionOk() (*InvalidResponseActionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvalidResponseAction.Get(), o.InvalidResponseAction.IsSet()
}

// HasInvalidResponseAction returns a boolean if a field has been set.
func (o *PatchedFlowStageBindingRequest) HasInvalidResponseAction() bool {
	if o != nil && o.InvalidResponseAction.IsSet() {
		return true
	}

	return false
}

// SetInvalidResponseAction gets a reference to the given NullableInvalidResponseActionEnum and assigns it to the InvalidResponseAction field.
func (o *PatchedFlowStageBindingRequest) SetInvalidResponseAction(v InvalidResponseActionEnum) {
	o.InvalidResponseAction.Set(&v)
}

// SetInvalidResponseActionNil sets the value for InvalidResponseAction to be an explicit nil
func (o *PatchedFlowStageBindingRequest) SetInvalidResponseActionNil() {
	o.InvalidResponseAction.Set(nil)
}

// UnsetInvalidResponseAction ensures that no value is present for InvalidResponseAction, not even an explicit nil
func (o *PatchedFlowStageBindingRequest) UnsetInvalidResponseAction() {
	o.InvalidResponseAction.Unset()
}

func (o PatchedFlowStageBindingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Stage != nil {
		toSerialize["stage"] = o.Stage
	}
	if o.EvaluateOnPlan != nil {
		toSerialize["evaluate_on_plan"] = o.EvaluateOnPlan
	}
	if o.ReEvaluatePolicies != nil {
		toSerialize["re_evaluate_policies"] = o.ReEvaluatePolicies
	}
	if o.Order != nil {
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

type NullablePatchedFlowStageBindingRequest struct {
	value *PatchedFlowStageBindingRequest
	isSet bool
}

func (v NullablePatchedFlowStageBindingRequest) Get() *PatchedFlowStageBindingRequest {
	return v.value
}

func (v *NullablePatchedFlowStageBindingRequest) Set(val *PatchedFlowStageBindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedFlowStageBindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedFlowStageBindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedFlowStageBindingRequest(val *PatchedFlowStageBindingRequest) *NullablePatchedFlowStageBindingRequest {
	return &NullablePatchedFlowStageBindingRequest{value: val, isSet: true}
}

func (v NullablePatchedFlowStageBindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedFlowStageBindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
