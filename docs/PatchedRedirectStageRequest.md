# PatchedRedirectStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**KeepContext** | Pointer to **bool** |  | [optional] 
**Mode** | Pointer to [**RedirectStageModeEnum**](RedirectStageModeEnum.md) |  | [optional] 
**TargetStatic** | Pointer to **string** |  | [optional] 
**TargetFlow** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPatchedRedirectStageRequest

`func NewPatchedRedirectStageRequest() *PatchedRedirectStageRequest`

NewPatchedRedirectStageRequest instantiates a new PatchedRedirectStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRedirectStageRequestWithDefaults

`func NewPatchedRedirectStageRequestWithDefaults() *PatchedRedirectStageRequest`

NewPatchedRedirectStageRequestWithDefaults instantiates a new PatchedRedirectStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedRedirectStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRedirectStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRedirectStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRedirectStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedRedirectStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedRedirectStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedRedirectStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedRedirectStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetKeepContext

`func (o *PatchedRedirectStageRequest) GetKeepContext() bool`

GetKeepContext returns the KeepContext field if non-nil, zero value otherwise.

### GetKeepContextOk

`func (o *PatchedRedirectStageRequest) GetKeepContextOk() (*bool, bool)`

GetKeepContextOk returns a tuple with the KeepContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepContext

`func (o *PatchedRedirectStageRequest) SetKeepContext(v bool)`

SetKeepContext sets KeepContext field to given value.

### HasKeepContext

`func (o *PatchedRedirectStageRequest) HasKeepContext() bool`

HasKeepContext returns a boolean if a field has been set.

### GetMode

`func (o *PatchedRedirectStageRequest) GetMode() RedirectStageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedRedirectStageRequest) GetModeOk() (*RedirectStageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedRedirectStageRequest) SetMode(v RedirectStageModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedRedirectStageRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTargetStatic

`func (o *PatchedRedirectStageRequest) GetTargetStatic() string`

GetTargetStatic returns the TargetStatic field if non-nil, zero value otherwise.

### GetTargetStaticOk

`func (o *PatchedRedirectStageRequest) GetTargetStaticOk() (*string, bool)`

GetTargetStaticOk returns a tuple with the TargetStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStatic

`func (o *PatchedRedirectStageRequest) SetTargetStatic(v string)`

SetTargetStatic sets TargetStatic field to given value.

### HasTargetStatic

`func (o *PatchedRedirectStageRequest) HasTargetStatic() bool`

HasTargetStatic returns a boolean if a field has been set.

### GetTargetFlow

`func (o *PatchedRedirectStageRequest) GetTargetFlow() string`

GetTargetFlow returns the TargetFlow field if non-nil, zero value otherwise.

### GetTargetFlowOk

`func (o *PatchedRedirectStageRequest) GetTargetFlowOk() (*string, bool)`

GetTargetFlowOk returns a tuple with the TargetFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFlow

`func (o *PatchedRedirectStageRequest) SetTargetFlow(v string)`

SetTargetFlow sets TargetFlow field to given value.

### HasTargetFlow

`func (o *PatchedRedirectStageRequest) HasTargetFlow() bool`

HasTargetFlow returns a boolean if a field has been set.

### SetTargetFlowNil

`func (o *PatchedRedirectStageRequest) SetTargetFlowNil(b bool)`

 SetTargetFlowNil sets the value for TargetFlow to be an explicit nil

### UnsetTargetFlow
`func (o *PatchedRedirectStageRequest) UnsetTargetFlow()`

UnsetTargetFlow ensures that no value is present for TargetFlow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


