# UserWriteStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowRequest**](FlowRequest.md) |  | [optional] 
**CreateUsersAsInactive** | Pointer to **bool** | When set, newly created users are inactive and cannot login. | [optional] 

## Methods

### NewUserWriteStageRequest

`func NewUserWriteStageRequest(name string, ) *UserWriteStageRequest`

NewUserWriteStageRequest instantiates a new UserWriteStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWriteStageRequestWithDefaults

`func NewUserWriteStageRequestWithDefaults() *UserWriteStageRequest`

NewUserWriteStageRequestWithDefaults instantiates a new UserWriteStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserWriteStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserWriteStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserWriteStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *UserWriteStageRequest) GetFlowSet() []FlowRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *UserWriteStageRequest) GetFlowSetOk() (*[]FlowRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *UserWriteStageRequest) SetFlowSet(v []FlowRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *UserWriteStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetCreateUsersAsInactive

`func (o *UserWriteStageRequest) GetCreateUsersAsInactive() bool`

GetCreateUsersAsInactive returns the CreateUsersAsInactive field if non-nil, zero value otherwise.

### GetCreateUsersAsInactiveOk

`func (o *UserWriteStageRequest) GetCreateUsersAsInactiveOk() (*bool, bool)`

GetCreateUsersAsInactiveOk returns a tuple with the CreateUsersAsInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUsersAsInactive

`func (o *UserWriteStageRequest) SetCreateUsersAsInactive(v bool)`

SetCreateUsersAsInactive sets CreateUsersAsInactive field to given value.

### HasCreateUsersAsInactive

`func (o *UserWriteStageRequest) HasCreateUsersAsInactive() bool`

HasCreateUsersAsInactive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


