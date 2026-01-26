# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Provider** | **int32** |  | 
**ProviderObj** | [**RACProvider**](RACProvider.md) |  | [readonly] 
**Protocol** | [**ProtocolEnum**](ProtocolEnum.md) |  | 
**Host** | **string** |  | 
**Settings** | Pointer to **interface{}** |  | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**AuthMode** | [**AuthModeEnum**](AuthModeEnum.md) |  | 
**LaunchUrl** | **NullableString** | Build actual launch URL (the provider itself does not have one, just individual endpoints) | [readonly] 
**MaximumConnections** | Pointer to **int32** |  | [optional] 

## Methods

### NewEndpoint

`func NewEndpoint(pk string, name string, provider int32, providerObj RACProvider, protocol ProtocolEnum, host string, authMode AuthModeEnum, launchUrl NullableString, ) *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *Endpoint) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *Endpoint) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *Endpoint) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *Endpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Endpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Endpoint) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *Endpoint) GetProvider() int32`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Endpoint) GetProviderOk() (*int32, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Endpoint) SetProvider(v int32)`

SetProvider sets Provider field to given value.


### GetProviderObj

`func (o *Endpoint) GetProviderObj() RACProvider`

GetProviderObj returns the ProviderObj field if non-nil, zero value otherwise.

### GetProviderObjOk

`func (o *Endpoint) GetProviderObjOk() (*RACProvider, bool)`

GetProviderObjOk returns a tuple with the ProviderObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderObj

`func (o *Endpoint) SetProviderObj(v RACProvider)`

SetProviderObj sets ProviderObj field to given value.


### GetProtocol

`func (o *Endpoint) GetProtocol() ProtocolEnum`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Endpoint) GetProtocolOk() (*ProtocolEnum, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Endpoint) SetProtocol(v ProtocolEnum)`

SetProtocol sets Protocol field to given value.


### GetHost

`func (o *Endpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Endpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Endpoint) SetHost(v string)`

SetHost sets Host field to given value.


### GetSettings

`func (o *Endpoint) GetSettings() interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Endpoint) GetSettingsOk() (*interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Endpoint) SetSettings(v interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Endpoint) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *Endpoint) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *Endpoint) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetPropertyMappings

`func (o *Endpoint) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *Endpoint) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *Endpoint) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *Endpoint) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetAuthMode

`func (o *Endpoint) GetAuthMode() AuthModeEnum`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *Endpoint) GetAuthModeOk() (*AuthModeEnum, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *Endpoint) SetAuthMode(v AuthModeEnum)`

SetAuthMode sets AuthMode field to given value.


### GetLaunchUrl

`func (o *Endpoint) GetLaunchUrl() string`

GetLaunchUrl returns the LaunchUrl field if non-nil, zero value otherwise.

### GetLaunchUrlOk

`func (o *Endpoint) GetLaunchUrlOk() (*string, bool)`

GetLaunchUrlOk returns a tuple with the LaunchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchUrl

`func (o *Endpoint) SetLaunchUrl(v string)`

SetLaunchUrl sets LaunchUrl field to given value.


### SetLaunchUrlNil

`func (o *Endpoint) SetLaunchUrlNil(b bool)`

 SetLaunchUrlNil sets the value for LaunchUrl to be an explicit nil

### UnsetLaunchUrl
`func (o *Endpoint) UnsetLaunchUrl()`

UnsetLaunchUrl ensures that no value is present for LaunchUrl, not even an explicit nil
### GetMaximumConnections

`func (o *Endpoint) GetMaximumConnections() int32`

GetMaximumConnections returns the MaximumConnections field if non-nil, zero value otherwise.

### GetMaximumConnectionsOk

`func (o *Endpoint) GetMaximumConnectionsOk() (*int32, bool)`

GetMaximumConnectionsOk returns a tuple with the MaximumConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConnections

`func (o *Endpoint) SetMaximumConnections(v int32)`

SetMaximumConnections sets MaximumConnections field to given value.

### HasMaximumConnections

`func (o *Endpoint) HasMaximumConnections() bool`

HasMaximumConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


