# PatchedConsentStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**FlowSet** | Pointer to [**[]FlowSetRequest**](FlowSetRequest.md) |  | [optional] 
**Mode** | Pointer to [**ConsentStageModeEnum**](ConsentStageModeEnum.md) |  | [optional] 
**ConsentExpireIn** | Pointer to **string** | Offset after which consent expires. (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 

## Methods

### NewPatchedConsentStageRequest

`func NewPatchedConsentStageRequest() *PatchedConsentStageRequest`

NewPatchedConsentStageRequest instantiates a new PatchedConsentStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedConsentStageRequestWithDefaults

`func NewPatchedConsentStageRequestWithDefaults() *PatchedConsentStageRequest`

NewPatchedConsentStageRequestWithDefaults instantiates a new PatchedConsentStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedConsentStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedConsentStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedConsentStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedConsentStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFlowSet

`func (o *PatchedConsentStageRequest) GetFlowSet() []FlowSetRequest`

GetFlowSet returns the FlowSet field if non-nil, zero value otherwise.

### GetFlowSetOk

`func (o *PatchedConsentStageRequest) GetFlowSetOk() (*[]FlowSetRequest, bool)`

GetFlowSetOk returns a tuple with the FlowSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSet

`func (o *PatchedConsentStageRequest) SetFlowSet(v []FlowSetRequest)`

SetFlowSet sets FlowSet field to given value.

### HasFlowSet

`func (o *PatchedConsentStageRequest) HasFlowSet() bool`

HasFlowSet returns a boolean if a field has been set.

### GetMode

`func (o *PatchedConsentStageRequest) GetMode() ConsentStageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedConsentStageRequest) GetModeOk() (*ConsentStageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedConsentStageRequest) SetMode(v ConsentStageModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedConsentStageRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetConsentExpireIn

`func (o *PatchedConsentStageRequest) GetConsentExpireIn() string`

GetConsentExpireIn returns the ConsentExpireIn field if non-nil, zero value otherwise.

### GetConsentExpireInOk

`func (o *PatchedConsentStageRequest) GetConsentExpireInOk() (*string, bool)`

GetConsentExpireInOk returns a tuple with the ConsentExpireIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentExpireIn

`func (o *PatchedConsentStageRequest) SetConsentExpireIn(v string)`

SetConsentExpireIn sets ConsentExpireIn field to given value.

### HasConsentExpireIn

`func (o *PatchedConsentStageRequest) HasConsentExpireIn() bool`

HasConsentExpireIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


