# EndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Provider** | **int32** |  | 
**Protocol** | [**ProtocolEnum**](ProtocolEnum.md) |  | 
**Host** | **string** |  | 
**Settings** | Pointer to **interface{}** |  | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**AuthMode** | [**AuthModeEnum**](AuthModeEnum.md) |  | 
**MaximumConnections** | Pointer to **int32** |  | [optional] 

## Methods

### NewEndpointRequest

`func NewEndpointRequest(name string, provider int32, protocol ProtocolEnum, host string, authMode AuthModeEnum, ) *EndpointRequest`

NewEndpointRequest instantiates a new EndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointRequestWithDefaults

`func NewEndpointRequestWithDefaults() *EndpointRequest`

NewEndpointRequestWithDefaults instantiates a new EndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EndpointRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *EndpointRequest) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EndpointRequest) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EndpointRequest) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetProtocol

`func (o *EndpointRequest) GetProtocol() ProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *EndpointRequest) GetProtocolOk() (*ProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *EndpointRequest) SetProtocol(v ProtocolEnum)`

SetProtocol sets Protocol field to given value.


### GetHost

`func (o *EndpointRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EndpointRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EndpointRequest) SetHost(v string)`

SetHost sets Host field to given value.


### GetSettings

`func (o *EndpointRequest) GetSettings() interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *EndpointRequest) GetSettingsOk() (*interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *EndpointRequest) SetSettings(v interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *EndpointRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *EndpointRequest) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *EndpointRequest) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetPropertyMappings

`func (o *EndpointRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *EndpointRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *EndpointRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *EndpointRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetAuthMode

`func (o *EndpointRequest) GetAuthMode() AuthModeEnum`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *EndpointRequest) GetAuthModeOk() (*AuthModeEnum, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *EndpointRequest) SetAuthMode(v AuthModeEnum)`

SetAuthMode sets AuthMode field to given value.


### GetMaximumConnections

`func (o *EndpointRequest) GetMaximumConnections() int32`

GetMaximumConnections returns the MaximumConnections field if non-nil, zero value otherwise.

### GetMaximumConnectionsOk

`func (o *EndpointRequest) GetMaximumConnectionsOk() (*int32, bool)`

GetMaximumConnectionsOk returns a tuple with the MaximumConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConnections

`func (o *EndpointRequest) SetMaximumConnections(v int32)`

SetMaximumConnections sets MaximumConnections field to given value.

### HasMaximumConnections

`func (o *EndpointRequest) HasMaximumConnections() bool`

HasMaximumConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


