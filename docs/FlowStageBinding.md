# FlowStageBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**PolicybindingmodelPtrId** | **string** |  | [readonly] 
**Target** | **string** |  | 
**Stage** | **string** |  | 
**StageObj** | [**Stage**](Stage.md) |  | [readonly] 
**EvaluateOnPlan** | Pointer to **bool** | Evaluate policies during the Flow planning process. | [optional] 
**ReEvaluatePolicies** | Pointer to **bool** | Evaluate policies when the Stage is present to the user. | [optional] 
**Order** | **int32** |  | 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**InvalidResponseAction** | Pointer to [**InvalidResponseActionEnum**](InvalidResponseActionEnum.md) |  | [optional] 

## Methods

### NewFlowStageBinding

`func NewFlowStageBinding(pk string, policybindingmodelPtrId string, target string, stage string, stageObj Stage, order int32, ) *FlowStageBinding`

NewFlowStageBinding instantiates a new FlowStageBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowStageBindingWithDefaults

`func NewFlowStageBindingWithDefaults() *FlowStageBinding`

NewFlowStageBindingWithDefaults instantiates a new FlowStageBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *FlowStageBinding) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *FlowStageBinding) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *FlowStageBinding) SetPk(v string)`

SetPk sets Pk field to given value.


### GetPolicybindingmodelPtrId

`func (o *FlowStageBinding) GetPolicybindingmodelPtrId() string`

GetPolicybindingmodelPtrId returns the PolicybindingmodelPtrId field if non-nil, zero value otherwise.

### GetPolicybindingmodelPtrIdOk

`func (o *FlowStageBinding) GetPolicybindingmodelPtrIdOk() (*string, bool)`

GetPolicybindingmodelPtrIdOk returns a tuple with the PolicybindingmodelPtrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicybindingmodelPtrId

`func (o *FlowStageBinding) SetPolicybindingmodelPtrId(v string)`

SetPolicybindingmodelPtrId sets PolicybindingmodelPtrId field to given value.


### GetTarget

`func (o *FlowStageBinding) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *FlowStageBinding) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *FlowStageBinding) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetStage

`func (o *FlowStageBinding) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *FlowStageBinding) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *FlowStageBinding) SetStage(v string)`

SetStage sets Stage field to given value.


### GetStageObj

`func (o *FlowStageBinding) GetStageObj() Stage`

GetStageObj returns the StageObj field if non-nil, zero value otherwise.

### GetStageObjOk

`func (o *FlowStageBinding) GetStageObjOk() (*Stage, bool)`

GetStageObjOk returns a tuple with the StageObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageObj

`func (o *FlowStageBinding) SetStageObj(v Stage)`

SetStageObj sets StageObj field to given value.


### GetEvaluateOnPlan

`func (o *FlowStageBinding) GetEvaluateOnPlan() bool`

GetEvaluateOnPlan returns the EvaluateOnPlan field if non-nil, zero value otherwise.

### GetEvaluateOnPlanOk

`func (o *FlowStageBinding) GetEvaluateOnPlanOk() (*bool, bool)`

GetEvaluateOnPlanOk returns a tuple with the EvaluateOnPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluateOnPlan

`func (o *FlowStageBinding) SetEvaluateOnPlan(v bool)`

SetEvaluateOnPlan sets EvaluateOnPlan field to given value.

### HasEvaluateOnPlan

`func (o *FlowStageBinding) HasEvaluateOnPlan() bool`

HasEvaluateOnPlan returns a boolean if a field has been set.

### GetReEvaluatePolicies

`func (o *FlowStageBinding) GetReEvaluatePolicies() bool`

GetReEvaluatePolicies returns the ReEvaluatePolicies field if non-nil, zero value otherwise.

### GetReEvaluatePoliciesOk

`func (o *FlowStageBinding) GetReEvaluatePoliciesOk() (*bool, bool)`

GetReEvaluatePoliciesOk returns a tuple with the ReEvaluatePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReEvaluatePolicies

`func (o *FlowStageBinding) SetReEvaluatePolicies(v bool)`

SetReEvaluatePolicies sets ReEvaluatePolicies field to given value.

### HasReEvaluatePolicies

`func (o *FlowStageBinding) HasReEvaluatePolicies() bool`

HasReEvaluatePolicies returns a boolean if a field has been set.

### GetOrder

`func (o *FlowStageBinding) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FlowStageBinding) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FlowStageBinding) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetPolicyEngineMode

`func (o *FlowStageBinding) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *FlowStageBinding) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *FlowStageBinding) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *FlowStageBinding) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetInvalidResponseAction

`func (o *FlowStageBinding) GetInvalidResponseAction() InvalidResponseActionEnum`

GetInvalidResponseAction returns the InvalidResponseAction field if non-nil, zero value otherwise.

### GetInvalidResponseActionOk

`func (o *FlowStageBinding) GetInvalidResponseActionOk() (*InvalidResponseActionEnum, bool)`

GetInvalidResponseActionOk returns a tuple with the InvalidResponseAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidResponseAction

`func (o *FlowStageBinding) SetInvalidResponseAction(v InvalidResponseActionEnum)`

SetInvalidResponseAction sets InvalidResponseAction field to given value.

### HasInvalidResponseAction

`func (o *FlowStageBinding) HasInvalidResponseAction() bool`

HasInvalidResponseAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


