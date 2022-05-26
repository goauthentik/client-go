# FlowInspectorPlanCurrentStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**PolicybindingmodelPtrId** | **string** |  | [readonly] 
**Target** | **string** |  | 
**Stage** | **string** |  | 
**StageObj** | [**FlowStageBindingStageObj**](FlowStageBindingStageObj.md) |  | 
**EvaluateOnPlan** | Pointer to **bool** | Evaluate policies during the Flow planning process. Disable this for input-based policies. | [optional] 
**ReEvaluatePolicies** | Pointer to **bool** | Evaluate policies when the Stage is present to the user. | [optional] 
**Order** | **int32** |  | 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**InvalidResponseAction** | Pointer to [**NullableInvalidResponseActionEnum**](InvalidResponseActionEnum.md) | Configure how the flow executor should handle an invalid response to a challenge. RETRY returns the error message and a similar challenge to the executor. RESTART restarts the flow from the beginning, and RESTART_WITH_CONTEXT restarts the flow while keeping the current context. | [optional] 

## Methods

### NewFlowInspectorPlanCurrentStage

`func NewFlowInspectorPlanCurrentStage(pk string, policybindingmodelPtrId string, target string, stage string, stageObj FlowStageBindingStageObj, order int32, ) *FlowInspectorPlanCurrentStage`

NewFlowInspectorPlanCurrentStage instantiates a new FlowInspectorPlanCurrentStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowInspectorPlanCurrentStageWithDefaults

`func NewFlowInspectorPlanCurrentStageWithDefaults() *FlowInspectorPlanCurrentStage`

NewFlowInspectorPlanCurrentStageWithDefaults instantiates a new FlowInspectorPlanCurrentStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *FlowInspectorPlanCurrentStage) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *FlowInspectorPlanCurrentStage) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *FlowInspectorPlanCurrentStage) SetPk(v string)`

SetPk sets Pk field to given value.


### GetPolicybindingmodelPtrId

`func (o *FlowInspectorPlanCurrentStage) GetPolicybindingmodelPtrId() string`

GetPolicybindingmodelPtrId returns the PolicybindingmodelPtrId field if non-nil, zero value otherwise.

### GetPolicybindingmodelPtrIdOk

`func (o *FlowInspectorPlanCurrentStage) GetPolicybindingmodelPtrIdOk() (*string, bool)`

GetPolicybindingmodelPtrIdOk returns a tuple with the PolicybindingmodelPtrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicybindingmodelPtrId

`func (o *FlowInspectorPlanCurrentStage) SetPolicybindingmodelPtrId(v string)`

SetPolicybindingmodelPtrId sets PolicybindingmodelPtrId field to given value.


### GetTarget

`func (o *FlowInspectorPlanCurrentStage) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *FlowInspectorPlanCurrentStage) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *FlowInspectorPlanCurrentStage) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetStage

`func (o *FlowInspectorPlanCurrentStage) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *FlowInspectorPlanCurrentStage) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *FlowInspectorPlanCurrentStage) SetStage(v string)`

SetStage sets Stage field to given value.


### GetStageObj

`func (o *FlowInspectorPlanCurrentStage) GetStageObj() FlowStageBindingStageObj`

GetStageObj returns the StageObj field if non-nil, zero value otherwise.

### GetStageObjOk

`func (o *FlowInspectorPlanCurrentStage) GetStageObjOk() (*FlowStageBindingStageObj, bool)`

GetStageObjOk returns a tuple with the StageObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageObj

`func (o *FlowInspectorPlanCurrentStage) SetStageObj(v FlowStageBindingStageObj)`

SetStageObj sets StageObj field to given value.


### GetEvaluateOnPlan

`func (o *FlowInspectorPlanCurrentStage) GetEvaluateOnPlan() bool`

GetEvaluateOnPlan returns the EvaluateOnPlan field if non-nil, zero value otherwise.

### GetEvaluateOnPlanOk

`func (o *FlowInspectorPlanCurrentStage) GetEvaluateOnPlanOk() (*bool, bool)`

GetEvaluateOnPlanOk returns a tuple with the EvaluateOnPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluateOnPlan

`func (o *FlowInspectorPlanCurrentStage) SetEvaluateOnPlan(v bool)`

SetEvaluateOnPlan sets EvaluateOnPlan field to given value.

### HasEvaluateOnPlan

`func (o *FlowInspectorPlanCurrentStage) HasEvaluateOnPlan() bool`

HasEvaluateOnPlan returns a boolean if a field has been set.

### GetReEvaluatePolicies

`func (o *FlowInspectorPlanCurrentStage) GetReEvaluatePolicies() bool`

GetReEvaluatePolicies returns the ReEvaluatePolicies field if non-nil, zero value otherwise.

### GetReEvaluatePoliciesOk

`func (o *FlowInspectorPlanCurrentStage) GetReEvaluatePoliciesOk() (*bool, bool)`

GetReEvaluatePoliciesOk returns a tuple with the ReEvaluatePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReEvaluatePolicies

`func (o *FlowInspectorPlanCurrentStage) SetReEvaluatePolicies(v bool)`

SetReEvaluatePolicies sets ReEvaluatePolicies field to given value.

### HasReEvaluatePolicies

`func (o *FlowInspectorPlanCurrentStage) HasReEvaluatePolicies() bool`

HasReEvaluatePolicies returns a boolean if a field has been set.

### GetOrder

`func (o *FlowInspectorPlanCurrentStage) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FlowInspectorPlanCurrentStage) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FlowInspectorPlanCurrentStage) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetPolicyEngineMode

`func (o *FlowInspectorPlanCurrentStage) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *FlowInspectorPlanCurrentStage) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *FlowInspectorPlanCurrentStage) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *FlowInspectorPlanCurrentStage) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetInvalidResponseAction

`func (o *FlowInspectorPlanCurrentStage) GetInvalidResponseAction() InvalidResponseActionEnum`

GetInvalidResponseAction returns the InvalidResponseAction field if non-nil, zero value otherwise.

### GetInvalidResponseActionOk

`func (o *FlowInspectorPlanCurrentStage) GetInvalidResponseActionOk() (*InvalidResponseActionEnum, bool)`

GetInvalidResponseActionOk returns a tuple with the InvalidResponseAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidResponseAction

`func (o *FlowInspectorPlanCurrentStage) SetInvalidResponseAction(v InvalidResponseActionEnum)`

SetInvalidResponseAction sets InvalidResponseAction field to given value.

### HasInvalidResponseAction

`func (o *FlowInspectorPlanCurrentStage) HasInvalidResponseAction() bool`

HasInvalidResponseAction returns a boolean if a field has been set.

### SetInvalidResponseActionNil

`func (o *FlowInspectorPlanCurrentStage) SetInvalidResponseActionNil(b bool)`

 SetInvalidResponseActionNil sets the value for InvalidResponseAction to be an explicit nil

### UnsetInvalidResponseAction
`func (o *FlowInspectorPlanCurrentStage) UnsetInvalidResponseAction()`

UnsetInvalidResponseAction ensures that no value is present for InvalidResponseAction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


