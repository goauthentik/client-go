# PatchedEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to [**ProtocolEnum**](ProtocolEnum.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to **interface{}** |  | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**AuthMode** | Pointer to [**AuthModeEnum**](AuthModeEnum.md) |  | [optional] 
**MaximumConnections** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedEndpointRequest

`func NewPatchedEndpointRequest() *PatchedEndpointRequest`

NewPatchedEndpointRequest instantiates a new PatchedEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEndpointRequestWithDefaults

`func NewPatchedEndpointRequestWithDefaults() *PatchedEndpointRequest`

NewPatchedEndpointRequestWithDefaults instantiates a new PatchedEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEndpointRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEndpointRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEndpointRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEndpointRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *PatchedEndpointRequest) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PatchedEndpointRequest) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PatchedEndpointRequest) SetProvider(v int32)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PatchedEndpointRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedEndpointRequest) GetProtocol() ProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedEndpointRequest) GetProtocolOk() (*ProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedEndpointRequest) SetProtocol(v ProtocolEnum)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedEndpointRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetHost

`func (o *PatchedEndpointRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PatchedEndpointRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PatchedEndpointRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PatchedEndpointRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetSettings

`func (o *PatchedEndpointRequest) GetSettings() interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PatchedEndpointRequest) GetSettingsOk() (*interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PatchedEndpointRequest) SetSettings(v interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PatchedEndpointRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *PatchedEndpointRequest) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *PatchedEndpointRequest) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetPropertyMappings

`func (o *PatchedEndpointRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *PatchedEndpointRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *PatchedEndpointRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *PatchedEndpointRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetAuthMode

`func (o *PatchedEndpointRequest) GetAuthMode() AuthModeEnum`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *PatchedEndpointRequest) GetAuthModeOk() (*AuthModeEnum, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *PatchedEndpointRequest) SetAuthMode(v AuthModeEnum)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *PatchedEndpointRequest) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetMaximumConnections

`func (o *PatchedEndpointRequest) GetMaximumConnections() int32`

GetMaximumConnections returns the MaximumConnections field if non-nil, zero value otherwise.

### GetMaximumConnectionsOk

`func (o *PatchedEndpointRequest) GetMaximumConnectionsOk() (*int32, bool)`

GetMaximumConnectionsOk returns a tuple with the MaximumConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConnections

`func (o *PatchedEndpointRequest) SetMaximumConnections(v int32)`

SetMaximumConnections sets MaximumConnections field to given value.

### HasMaximumConnections

`func (o *PatchedEndpointRequest) HasMaximumConnections() bool`

HasMaximumConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


