# PatchedRadiusProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | Pointer to **string** | Flow used when authorizing this provider. | [optional] 
**InvalidationFlow** | Pointer to **string** | Flow used ending the session from a provider. | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**ClientNetworks** | Pointer to **string** | List of CIDRs (comma-separated) that clients can connect from. A more specific CIDR will match before a looser one. Clients connecting from a non-specified CIDR will be dropped. | [optional] 
**SharedSecret** | Pointer to **string** | Shared secret between clients and server to hash packets. | [optional] 
**MfaSupport** | Pointer to **bool** | When enabled, code-based multi-factor authentication can be used by appending a semicolon and the TOTP code to the password. This should only be enabled if all users that will bind to this provider have a TOTP device configured, as otherwise a password may incorrectly be rejected if it contains a semicolon. | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 

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

### GetAuthenticationFlow

`func (o *PatchedRadiusProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedRadiusProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedRadiusProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedRadiusProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedRadiusProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedRadiusProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
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

### GetInvalidationFlow

`func (o *PatchedRadiusProviderRequest) GetInvalidationFlow() string`

GetInvalidationFlow returns the InvalidationFlow field if non-nil, zero value otherwise.

### GetInvalidationFlowOk

`func (o *PatchedRadiusProviderRequest) GetInvalidationFlowOk() (*string, bool)`

GetInvalidationFlowOk returns a tuple with the InvalidationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidationFlow

`func (o *PatchedRadiusProviderRequest) SetInvalidationFlow(v string)`

SetInvalidationFlow sets InvalidationFlow field to given value.

### HasInvalidationFlow

`func (o *PatchedRadiusProviderRequest) HasInvalidationFlow() bool`

HasInvalidationFlow returns a boolean if a field has been set.

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

### GetMfaSupport

`func (o *PatchedRadiusProviderRequest) GetMfaSupport() bool`

GetMfaSupport returns the MfaSupport field if non-nil, zero value otherwise.

### GetMfaSupportOk

`func (o *PatchedRadiusProviderRequest) GetMfaSupportOk() (*bool, bool)`

GetMfaSupportOk returns a tuple with the MfaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaSupport

`func (o *PatchedRadiusProviderRequest) SetMfaSupport(v bool)`

SetMfaSupport sets MfaSupport field to given value.

### HasMfaSupport

`func (o *PatchedRadiusProviderRequest) HasMfaSupport() bool`

HasMfaSupport returns a boolean if a field has been set.

### GetCertificate

`func (o *PatchedRadiusProviderRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PatchedRadiusProviderRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PatchedRadiusProviderRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *PatchedRadiusProviderRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *PatchedRadiusProviderRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *PatchedRadiusProviderRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


