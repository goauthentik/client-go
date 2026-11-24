# FlowInspectorPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentStage** | [**FlowStageBinding**](FlowStageBinding.md) |  | [readonly] 
**NextPlannedStage** | [**FlowStageBinding**](FlowStageBinding.md) |  | [readonly] 
**PlanContext** | **map[string]interface{}** | Get the plan&#39;s context, sanitized | [readonly] 
**SessionId** | **string** | Get a unique session ID | [readonly] 

## Methods

### NewFlowInspectorPlan

`func NewFlowInspectorPlan(currentStage FlowStageBinding, nextPlannedStage FlowStageBinding, planContext map[string]interface{}, sessionId string, ) *FlowInspectorPlan`

NewFlowInspectorPlan instantiates a new FlowInspectorPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowInspectorPlanWithDefaults

`func NewFlowInspectorPlanWithDefaults() *FlowInspectorPlan`

NewFlowInspectorPlanWithDefaults instantiates a new FlowInspectorPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentStage

`func (o *FlowInspectorPlan) GetCurrentStage() FlowStageBinding`

GetCurrentStage returns the CurrentStage field if non-nil, zero value otherwise.

### GetCurrentStageOk

`func (o *FlowInspectorPlan) GetCurrentStageOk() (*FlowStageBinding, bool)`

GetCurrentStageOk returns a tuple with the CurrentStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStage

`func (o *FlowInspectorPlan) SetCurrentStage(v FlowStageBinding)`

SetCurrentStage sets CurrentStage field to given value.


### GetNextPlannedStage

`func (o *FlowInspectorPlan) GetNextPlannedStage() FlowStageBinding`

GetNextPlannedStage returns the NextPlannedStage field if non-nil, zero value otherwise.

### GetNextPlannedStageOk

`func (o *FlowInspectorPlan) GetNextPlannedStageOk() (*FlowStageBinding, bool)`

GetNextPlannedStageOk returns a tuple with the NextPlannedStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPlannedStage

`func (o *FlowInspectorPlan) SetNextPlannedStage(v FlowStageBinding)`

SetNextPlannedStage sets NextPlannedStage field to given value.


### GetPlanContext

`func (o *FlowInspectorPlan) GetPlanContext() map[string]interface{}`

GetPlanContext returns the PlanContext field if non-nil, zero value otherwise.

### GetPlanContextOk

`func (o *FlowInspectorPlan) GetPlanContextOk() (*map[string]interface{}, bool)`

GetPlanContextOk returns a tuple with the PlanContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanContext

`func (o *FlowInspectorPlan) SetPlanContext(v map[string]interface{})`

SetPlanContext sets PlanContext field to given value.


### GetSessionId

`func (o *FlowInspectorPlan) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *FlowInspectorPlan) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *FlowInspectorPlan) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


