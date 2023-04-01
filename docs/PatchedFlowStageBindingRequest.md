# PatchedFlowStageBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 
**EvaluateOnPlan** | Pointer to **bool** | Evaluate policies during the Flow planning process. | [optional] 
**ReEvaluatePolicies** | Pointer to **bool** | Evaluate policies when the Stage is present to the user. | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**InvalidResponseAction** | Pointer to [**InvalidResponseActionEnum**](InvalidResponseActionEnum.md) |  | [optional] 

## Methods

### NewPatchedFlowStageBindingRequest

`func NewPatchedFlowStageBindingRequest() *PatchedFlowStageBindingRequest`

NewPatchedFlowStageBindingRequest instantiates a new PatchedFlowStageBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFlowStageBindingRequestWithDefaults

`func NewPatchedFlowStageBindingRequestWithDefaults() *PatchedFlowStageBindingRequest`

NewPatchedFlowStageBindingRequestWithDefaults instantiates a new PatchedFlowStageBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *PatchedFlowStageBindingRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PatchedFlowStageBindingRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PatchedFlowStageBindingRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PatchedFlowStageBindingRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStage

`func (o *PatchedFlowStageBindingRequest) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *PatchedFlowStageBindingRequest) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *PatchedFlowStageBindingRequest) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *PatchedFlowStageBindingRequest) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetEvaluateOnPlan

`func (o *PatchedFlowStageBindingRequest) GetEvaluateOnPlan() bool`

GetEvaluateOnPlan returns the EvaluateOnPlan field if non-nil, zero value otherwise.

### GetEvaluateOnPlanOk

`func (o *PatchedFlowStageBindingRequest) GetEvaluateOnPlanOk() (*bool, bool)`

GetEvaluateOnPlanOk returns a tuple with the EvaluateOnPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluateOnPlan

`func (o *PatchedFlowStageBindingRequest) SetEvaluateOnPlan(v bool)`

SetEvaluateOnPlan sets EvaluateOnPlan field to given value.

### HasEvaluateOnPlan

`func (o *PatchedFlowStageBindingRequest) HasEvaluateOnPlan() bool`

HasEvaluateOnPlan returns a boolean if a field has been set.

### GetReEvaluatePolicies

`func (o *PatchedFlowStageBindingRequest) GetReEvaluatePolicies() bool`

GetReEvaluatePolicies returns the ReEvaluatePolicies field if non-nil, zero value otherwise.

### GetReEvaluatePoliciesOk

`func (o *PatchedFlowStageBindingRequest) GetReEvaluatePoliciesOk() (*bool, bool)`

GetReEvaluatePoliciesOk returns a tuple with the ReEvaluatePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReEvaluatePolicies

`func (o *PatchedFlowStageBindingRequest) SetReEvaluatePolicies(v bool)`

SetReEvaluatePolicies sets ReEvaluatePolicies field to given value.

### HasReEvaluatePolicies

`func (o *PatchedFlowStageBindingRequest) HasReEvaluatePolicies() bool`

HasReEvaluatePolicies returns a boolean if a field has been set.

### GetOrder

`func (o *PatchedFlowStageBindingRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PatchedFlowStageBindingRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PatchedFlowStageBindingRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PatchedFlowStageBindingRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *PatchedFlowStageBindingRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PatchedFlowStageBindingRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PatchedFlowStageBindingRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PatchedFlowStageBindingRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetInvalidResponseAction

`func (o *PatchedFlowStageBindingRequest) GetInvalidResponseAction() InvalidResponseActionEnum`

GetInvalidResponseAction returns the InvalidResponseAction field if non-nil, zero value otherwise.

### GetInvalidResponseActionOk

`func (o *PatchedFlowStageBindingRequest) GetInvalidResponseActionOk() (*InvalidResponseActionEnum, bool)`

GetInvalidResponseActionOk returns a tuple with the InvalidResponseAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidResponseAction

`func (o *PatchedFlowStageBindingRequest) SetInvalidResponseAction(v InvalidResponseActionEnum)`

SetInvalidResponseAction sets InvalidResponseAction field to given value.

### HasInvalidResponseAction

`func (o *PatchedFlowStageBindingRequest) HasInvalidResponseAction() bool`

HasInvalidResponseAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


