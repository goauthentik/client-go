# FlowStageBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **string** |  | 
**Stage** | **string** |  | 
**EvaluateOnPlan** | Pointer to **bool** | Evaluate policies during the Flow planning process. | [optional] 
**ReEvaluatePolicies** | Pointer to **bool** | Evaluate policies when the Stage is present to the user. | [optional] 
**Order** | **int32** |  | 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**InvalidResponseAction** | Pointer to [**InvalidResponseActionEnum**](InvalidResponseActionEnum.md) |  | [optional] 

## Methods

### NewFlowStageBindingRequest

`func NewFlowStageBindingRequest(target string, stage string, order int32, ) *FlowStageBindingRequest`

NewFlowStageBindingRequest instantiates a new FlowStageBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowStageBindingRequestWithDefaults

`func NewFlowStageBindingRequestWithDefaults() *FlowStageBindingRequest`

NewFlowStageBindingRequestWithDefaults instantiates a new FlowStageBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *FlowStageBindingRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *FlowStageBindingRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *FlowStageBindingRequest) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetStage

`func (o *FlowStageBindingRequest) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *FlowStageBindingRequest) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *FlowStageBindingRequest) SetStage(v string)`

SetStage sets Stage field to given value.


### GetEvaluateOnPlan

`func (o *FlowStageBindingRequest) GetEvaluateOnPlan() bool`

GetEvaluateOnPlan returns the EvaluateOnPlan field if non-nil, zero value otherwise.

### GetEvaluateOnPlanOk

`func (o *FlowStageBindingRequest) GetEvaluateOnPlanOk() (*bool, bool)`

GetEvaluateOnPlanOk returns a tuple with the EvaluateOnPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluateOnPlan

`func (o *FlowStageBindingRequest) SetEvaluateOnPlan(v bool)`

SetEvaluateOnPlan sets EvaluateOnPlan field to given value.

### HasEvaluateOnPlan

`func (o *FlowStageBindingRequest) HasEvaluateOnPlan() bool`

HasEvaluateOnPlan returns a boolean if a field has been set.

### GetReEvaluatePolicies

`func (o *FlowStageBindingRequest) GetReEvaluatePolicies() bool`

GetReEvaluatePolicies returns the ReEvaluatePolicies field if non-nil, zero value otherwise.

### GetReEvaluatePoliciesOk

`func (o *FlowStageBindingRequest) GetReEvaluatePoliciesOk() (*bool, bool)`

GetReEvaluatePoliciesOk returns a tuple with the ReEvaluatePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReEvaluatePolicies

`func (o *FlowStageBindingRequest) SetReEvaluatePolicies(v bool)`

SetReEvaluatePolicies sets ReEvaluatePolicies field to given value.

### HasReEvaluatePolicies

`func (o *FlowStageBindingRequest) HasReEvaluatePolicies() bool`

HasReEvaluatePolicies returns a boolean if a field has been set.

### GetOrder

`func (o *FlowStageBindingRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FlowStageBindingRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FlowStageBindingRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetPolicyEngineMode

`func (o *FlowStageBindingRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *FlowStageBindingRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *FlowStageBindingRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *FlowStageBindingRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetInvalidResponseAction

`func (o *FlowStageBindingRequest) GetInvalidResponseAction() InvalidResponseActionEnum`

GetInvalidResponseAction returns the InvalidResponseAction field if non-nil, zero value otherwise.

### GetInvalidResponseActionOk

`func (o *FlowStageBindingRequest) GetInvalidResponseActionOk() (*InvalidResponseActionEnum, bool)`

GetInvalidResponseActionOk returns a tuple with the InvalidResponseAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidResponseAction

`func (o *FlowStageBindingRequest) SetInvalidResponseAction(v InvalidResponseActionEnum)`

SetInvalidResponseAction sets InvalidResponseAction field to given value.

### HasInvalidResponseAction

`func (o *FlowStageBindingRequest) HasInvalidResponseAction() bool`

HasInvalidResponseAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


