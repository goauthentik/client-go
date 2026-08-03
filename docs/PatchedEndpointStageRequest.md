# PatchedEndpointStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Connector** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**StageModeEnum**](StageModeEnum.md) |  | [optional] 

## Methods

### NewPatchedEndpointStageRequest

`func NewPatchedEndpointStageRequest() *PatchedEndpointStageRequest`

NewPatchedEndpointStageRequest instantiates a new PatchedEndpointStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEndpointStageRequestWithDefaults

`func NewPatchedEndpointStageRequestWithDefaults() *PatchedEndpointStageRequest`

NewPatchedEndpointStageRequestWithDefaults instantiates a new PatchedEndpointStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEndpointStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEndpointStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEndpointStageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEndpointStageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnector

`func (o *PatchedEndpointStageRequest) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *PatchedEndpointStageRequest) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *PatchedEndpointStageRequest) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *PatchedEndpointStageRequest) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetMode

`func (o *PatchedEndpointStageRequest) GetMode() StageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedEndpointStageRequest) GetModeOk() (*StageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedEndpointStageRequest) SetMode(v StageModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedEndpointStageRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


