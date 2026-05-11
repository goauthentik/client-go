# EndpointStageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Connector** | **string** |  | 
**Mode** | Pointer to [**StageModeEnum**](StageModeEnum.md) |  | [optional] 

## Methods

### NewEndpointStageRequest

`func NewEndpointStageRequest(name string, connector string, ) *EndpointStageRequest`

NewEndpointStageRequest instantiates a new EndpointStageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointStageRequestWithDefaults

`func NewEndpointStageRequestWithDefaults() *EndpointStageRequest`

NewEndpointStageRequestWithDefaults instantiates a new EndpointStageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EndpointStageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointStageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointStageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConnector

`func (o *EndpointStageRequest) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *EndpointStageRequest) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *EndpointStageRequest) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetMode

`func (o *EndpointStageRequest) GetMode() StageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EndpointStageRequest) GetModeOk() (*StageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EndpointStageRequest) SetMode(v StageModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EndpointStageRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


