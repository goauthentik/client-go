# ConnectionTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | Pointer to **string** |  | [optional] 
**Provider** | **int32** |  | 
**Device** | **string** |  | 

## Methods

### NewConnectionTokenRequest

`func NewConnectionTokenRequest(provider int32, device string, ) *ConnectionTokenRequest`

NewConnectionTokenRequest instantiates a new ConnectionTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionTokenRequestWithDefaults

`func NewConnectionTokenRequestWithDefaults() *ConnectionTokenRequest`

NewConnectionTokenRequestWithDefaults instantiates a new ConnectionTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *ConnectionTokenRequest) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *ConnectionTokenRequest) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *ConnectionTokenRequest) SetPk(v string)`

SetPk sets Pk field to given value.

### HasPk

`func (o *ConnectionTokenRequest) HasPk() bool`

HasPk returns a boolean if a field has been set.

### GetProvider

`func (o *ConnectionTokenRequest) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ConnectionTokenRequest) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ConnectionTokenRequest) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetDevice

`func (o *ConnectionTokenRequest) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ConnectionTokenRequest) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ConnectionTokenRequest) SetDevice(v string)`

SetDevice sets Device field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


