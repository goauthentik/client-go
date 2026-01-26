# DummyStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**ThrowError** | Pointer to **bool** |  | [optional] 

## Methods

### NewDummyStageRequest

`func NewDummyStageRequest(name string, ) *DummyStageRequest`

NewDummyStageRequest instantiates a new DummyStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDummyStageRequestWithDefaults

`func NewDummyStageRequestWithDefaults() *DummyStageRequest`

NewDummyStageRequestWithDefaults instantiates a new DummyStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DummyStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DummyStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DummyStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *DummyStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *DummyStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *DummyStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *DummyStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetThrowError

`func (o *DummyStageRequest) GetThrowError() bool`

GetThrowError returns the ThrowError field if non-nil, zero value otherwise.

### GetThrowErrorOk

`func (o *DummyStageRequest) GetThrowErrorOk() (*bool, bool)`

GetThrowErrorOk returns a tuple with the ThrowError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrowError

`func (o *DummyStageRequest) SetThrowError(v bool)`

SetThrowError sets ThrowError field to given value.

### HasThrowError

`func (o *DummyStageRequest) HasThrowError() bool`

HasThrowError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


