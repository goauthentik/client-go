# PatchedDenyStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowRequest**](FlowRequest.md) |  | [optional] 

## Methods

### NewPatchedDenyStageRequest

`func NewPatchedDenyStageRequest() *PatchedDenyStageRequest`

NewPatchedDenyStageRequest instantiates a new PatchedDenyStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDenyStageRequestWithDefaults

`func NewPatchedDenyStageRequestWithDefaults() *PatchedDenyStageRequest`

NewPatchedDenyStageRequestWithDefaults instantiates a new PatchedDenyStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedDenyStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDenyStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDenyStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDenyStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedDenyStageRequest) GetFlowSet() []FlowRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedDenyStageRequest) GetFlowSetOk() (*[]FlowRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedDenyStageRequest) SetFlowSet(v []FlowRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedDenyStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


