# ConnectionTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | Pointer to **string** |  | [optional] 
**Provider** | **int32** |  | 
**Endpoint** | **string** |  | 

## Methods

### NewConnectionTokenRequest

`func NewConnectionTokenRequest(provider int32, endpoint string, ) *ConnectionTokenRequest`

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


### GetEndpoint

`func (o *ConnectionTokenRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ConnectionTokenRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ConnectionTokenRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


