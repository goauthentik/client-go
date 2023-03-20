# PatchedRadiusProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AuthorizationFlow** | Pointer to **string** | Flow used when authorizing this provider. | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**ClientNetworks** | Pointer to **string** | List of CIDRs (comma-seperated) that clients can connect from. A more specific CIDR will match before a looser one. Clients connecting from a non-specified CIDR will be dropped. | [optional] 
**SharedSecret** | Pointer to **string** | Shared secret between clients and server to hash packets. | [optional] 

## Methods

### NewPatchedRadiusProviderRequest

`func NewPatchedRadiusProviderRequest() *PatchedRadiusProviderRequest`

NewPatchedRadiusProviderRequest instantiates a new PatchedRadiusProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRadiusProviderRequestWithDefaults

`func NewPatchedRadiusProviderRequestWithDefaults() *PatchedRadiusProviderRequest`

NewPatchedRadiusProviderRequestWithDefaults instantiates a new PatchedRadiusProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedRadiusProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRadiusProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRadiusProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRadiusProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthorizationFlow

`func (o *PatchedRadiusProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *PatchedRadiusProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *PatchedRadiusProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.

### HasAuthorizationFlow

`func (o *PatchedRadiusProviderRequest) HasAuthorizationFlow() bool`

HasAuthorizationFlow returns a boolean if a field has been set.

### GetPropertyMappings

`func (o *PatchedRadiusProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *PatchedRadiusProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *PatchedRadiusProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *PatchedRadiusProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetClientNetworks

`func (o *PatchedRadiusProviderRequest) GetClientNetworks() string`

GetClientNetworks returns the ClientNetworks field if non-nil, zero value otherwise.

### GetClientNetworksOk

`func (o *PatchedRadiusProviderRequest) GetClientNetworksOk() (*string, bool)`

GetClientNetworksOk returns a tuple with the ClientNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetworks

`func (o *PatchedRadiusProviderRequest) SetClientNetworks(v string)`

SetClientNetworks sets ClientNetworks field to given value.

### HasClientNetworks

`func (o *PatchedRadiusProviderRequest) HasClientNetworks() bool`

HasClientNetworks returns a boolean if a field has been set.

### GetSharedSecret

`func (o *PatchedRadiusProviderRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *PatchedRadiusProviderRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *PatchedRadiusProviderRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *PatchedRadiusProviderRequest) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


