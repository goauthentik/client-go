# RedirectStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**KeepContext** | Pointer to **bool** |  | [optional] 
**Mode** | [**RedirectStageModeEnum**](RedirectStageModeEnum.md) |  | 
**TargetStatic** | Pointer to **string** |  | [optional] 
**TargetFlow** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRedirectStageRequest

`func NewRedirectStageRequest(name string, mode RedirectStageModeEnum, ) *RedirectStageRequest`

NewRedirectStageRequest instantiates a new RedirectStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectStageRequestWithDefaults

`func NewRedirectStageRequestWithDefaults() *RedirectStageRequest`

NewRedirectStageRequestWithDefaults instantiates a new RedirectStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RedirectStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RedirectStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RedirectStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetKeepContext

`func (o *RedirectStageRequest) GetKeepContext() bool`

GetKeepContext returns the KeepContext field if non-nil, zero value otherwise.

### GetKeepContextOk

`func (o *RedirectStageRequest) GetKeepContextOk() (*bool, bool)`

GetKeepContextOk returns a tuple with the KeepContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepContext

`func (o *RedirectStageRequest) SetKeepContext(v bool)`

SetKeepContext sets KeepContext field to given value.

### HasKeepContext

`func (o *RedirectStageRequest) HasKeepContext() bool`

HasKeepContext returns a boolean if a field has been set.

### GetMode

`func (o *RedirectStageRequest) GetMode() RedirectStageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RedirectStageRequest) GetModeOk() (*RedirectStageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RedirectStageRequest) SetMode(v RedirectStageModeEnum)`

SetMode sets Mode field to given value.


### GetTargetStatic

`func (o *RedirectStageRequest) GetTargetStatic() string`

GetTargetStatic returns the TargetStatic field if non-nil, zero value otherwise.

### GetTargetStaticOk

`func (o *RedirectStageRequest) GetTargetStaticOk() (*string, bool)`

GetTargetStaticOk returns a tuple with the TargetStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStatic

`func (o *RedirectStageRequest) SetTargetStatic(v string)`

SetTargetStatic sets TargetStatic field to given value.

### HasTargetStatic

`func (o *RedirectStageRequest) HasTargetStatic() bool`

HasTargetStatic returns a boolean if a field has been set.

### GetTargetFlow

`func (o *RedirectStageRequest) GetTargetFlow() string`

GetTargetFlow returns the TargetFlow field if non-nil, zero value otherwise.

### GetTargetFlowOk

`func (o *RedirectStageRequest) GetTargetFlowOk() (*string, bool)`

GetTargetFlowOk returns a tuple with the TargetFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFlow

`func (o *RedirectStageRequest) SetTargetFlow(v string)`

SetTargetFlow sets TargetFlow field to given value.

### HasTargetFlow

`func (o *RedirectStageRequest) HasTargetFlow() bool`

HasTargetFlow returns a boolean if a field has been set.

### SetTargetFlowNil

`func (o *RedirectStageRequest) SetTargetFlowNil(b bool)`

 SetTargetFlowNil sets the value for TargetFlow to be an explicit nil

### UnsetTargetFlow
`func (o *RedirectStageRequest) UnsetTargetFlow()`

UnsetTargetFlow ensures that no value is present for TargetFlow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


