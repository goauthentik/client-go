# UserLoginStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowRequest**](FlowRequest.md) |  | [optional] 
**SessionDuration** | Pointer to **string** | Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3) | [optional] 

## Methods

### NewUserLoginStageRequest

`func NewUserLoginStageRequest(name string, ) *UserLoginStageRequest`

NewUserLoginStageRequest instantiates a new UserLoginStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginStageRequestWithDefaults

`func NewUserLoginStageRequestWithDefaults() *UserLoginStageRequest`

NewUserLoginStageRequestWithDefaults instantiates a new UserLoginStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserLoginStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserLoginStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserLoginStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *UserLoginStageRequest) GetFlowSet() []FlowRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *UserLoginStageRequest) GetFlowSetOk() (*[]FlowRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *UserLoginStageRequest) SetFlowSet(v []FlowRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *UserLoginStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetSessionDuration

`func (o *UserLoginStageRequest) GetSessionDuration() string`

GetSessionDuration returns the SessionDuration field if non-nil, zero value otherwise.

### GetSessionDurationOk

`func (o *UserLoginStageRequest) GetSessionDurationOk() (*string, bool)`

GetSessionDurationOk returns a tuple with the SessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDuration

`func (o *UserLoginStageRequest) SetSessionDuration(v string)`

SetSessionDuration sets SessionDuration field to given value.

### HasSessionDuration

`func (o *UserLoginStageRequest) HasSessionDuration() bool`

HasSessionDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


