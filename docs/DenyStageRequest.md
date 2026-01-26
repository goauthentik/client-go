# DenyStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**DenyMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewDenyStageRequest

`func NewDenyStageRequest(name string, ) *DenyStageRequest`

NewDenyStageRequest instantiates a new DenyStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDenyStageRequestWithDefaults

`func NewDenyStageRequestWithDefaults() *DenyStageRequest`

NewDenyStageRequestWithDefaults instantiates a new DenyStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DenyStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DenyStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DenyStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *DenyStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *DenyStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *DenyStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *DenyStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetDenyMessage

`func (o *DenyStageRequest) GetDenyMessage() string`

GetDenyMessage returns the DenyMessage field if non-nil, zero value otherwise.

### GetDenyMessageOk

`func (o *DenyStageRequest) GetDenyMessageOk() (*string, bool)`

GetDenyMessageOk returns a tuple with the DenyMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyMessage

`func (o *DenyStageRequest) SetDenyMessage(v string)`

SetDenyMessage sets DenyMessage field to given value.

### HasDenyMessage

`func (o *DenyStageRequest) HasDenyMessage() bool`

HasDenyMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


