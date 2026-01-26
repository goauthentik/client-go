# PatchedDummyStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**ThrowError** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedDummyStageRequest

`func NewPatchedDummyStageRequest() *PatchedDummyStageRequest`

NewPatchedDummyStageRequest instantiates a new PatchedDummyStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDummyStageRequestWithDefaults

`func NewPatchedDummyStageRequestWithDefaults() *PatchedDummyStageRequest`

NewPatchedDummyStageRequestWithDefaults instantiates a new PatchedDummyStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedDummyStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDummyStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDummyStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDummyStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedDummyStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedDummyStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedDummyStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedDummyStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetThrowError

`func (o *PatchedDummyStageRequest) GetThrowError() bool`

GetThrowError returns the ThrowError field if non-nil, zero value otherwise.

### GetThrowErrorOk

`func (o *PatchedDummyStageRequest) GetThrowErrorOk() (*bool, bool)`

GetThrowErrorOk returns a tuple with the ThrowError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrowError

`func (o *PatchedDummyStageRequest) SetThrowError(v bool)`

SetThrowError sets ThrowError field to given value.

### HasThrowError

`func (o *PatchedDummyStageRequest) HasThrowError() bool`

HasThrowError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


