# FlowInspection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plans** | [**[]FlowInspectorPlan**](FlowInspectorPlan.md) |  | 
**CurrentPlan** | Pointer to [**FlowInspectorPlan**](FlowInspectorPlan.md) |  | [optional] 
**IsCompleted** | **bool** |  | 

## Methods

### NewFlowInspection

`func NewFlowInspection(plans []FlowInspectorPlan, isCompleted bool, ) *FlowInspection`

NewFlowInspection instantiates a new FlowInspection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowInspectionWithDefaults

`func NewFlowInspectionWithDefaults() *FlowInspection`

NewFlowInspectionWithDefaults instantiates a new FlowInspection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlans

`func (o *FlowInspection) GetPlans() []FlowInspectorPlan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *FlowInspection) GetPlansOk() (*[]FlowInspectorPlan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *FlowInspection) SetPlans(v []FlowInspectorPlan)`

SetPlans sets Plans field to given value.


### GetCurrentPlan

`func (o *FlowInspection) GetCurrentPlan() FlowInspectorPlan`

GetCurrentPlan returns the CurrentPlan field if non-nil, zero value otherwise.

### GetCurrentPlanOk

`func (o *FlowInspection) GetCurrentPlanOk() (*FlowInspectorPlan, bool)`

GetCurrentPlanOk returns a tuple with the CurrentPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPlan

`func (o *FlowInspection) SetCurrentPlan(v FlowInspectorPlan)`

SetCurrentPlan sets CurrentPlan field to given value.

### HasCurrentPlan

`func (o *FlowInspection) HasCurrentPlan() bool`

HasCurrentPlan returns a boolean if a field has been set.

### GetIsCompleted

`func (o *FlowInspection) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *FlowInspection) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *FlowInspection) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


