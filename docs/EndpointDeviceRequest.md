# EndpointDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | Pointer to **string** |  | [optional] 
**Name** | **string** | The human-readable name of this device. | 

## Methods

### NewEndpointDeviceRequest

`func NewEndpointDeviceRequest(name string, ) *EndpointDeviceRequest`

NewEndpointDeviceRequest instantiates a new EndpointDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointDeviceRequestWithDefaults

`func NewEndpointDeviceRequestWithDefaults() *EndpointDeviceRequest`

NewEndpointDeviceRequestWithDefaults instantiates a new EndpointDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *EndpointDeviceRequest) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *EndpointDeviceRequest) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *EndpointDeviceRequest) SetPk(v string)`

SetPk sets Pk field to given value.

### HasPk

`func (o *EndpointDeviceRequest) HasPk() bool`

HasPk returns a boolean if a field has been set.

### GetName

`func (o *EndpointDeviceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointDeviceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointDeviceRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


