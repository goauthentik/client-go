# ConsentStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**Mode** | Pointer to [**ConsentStageModeEnum**](ConsentStageModeEnum.md) |  | [optional] 
**ConsentExpireIn** | Pointer to **string** | Offset after which consent expires. (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 

## Methods

### NewConsentStageRequest

`func NewConsentStageRequest(name string, ) *ConsentStageRequest`

NewConsentStageRequest instantiates a new ConsentStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentStageRequestWithDefaults

`func NewConsentStageRequestWithDefaults() *ConsentStageRequest`

NewConsentStageRequestWithDefaults instantiates a new ConsentStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConsentStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsentStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsentStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowSet

`func (o *ConsentStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *ConsentStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *ConsentStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *ConsentStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetMode

`func (o *ConsentStageRequest) GetMode() ConsentStageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ConsentStageRequest) GetModeOk() (*ConsentStageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ConsentStageRequest) SetMode(v ConsentStageModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ConsentStageRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetConsentExpireIn

`func (o *ConsentStageRequest) GetConsentExpireIn() string`

GetConsentExpireIn returns the ConsentExpireIn field if non-nil, zero value otherwise.

### GetConsentExpireInOk

`func (o *ConsentStageRequest) GetConsentExpireInOk() (*string, bool)`

GetConsentExpireInOk returns a tuple with the ConsentExpireIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentExpireIn

`func (o *ConsentStageRequest) SetConsentExpireIn(v string)`

SetConsentExpireIn sets ConsentExpireIn field to given value.

### HasConsentExpireIn

`func (o *ConsentStageRequest) HasConsentExpireIn() bool`

HasConsentExpireIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


