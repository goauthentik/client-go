# RadiusProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**InvalidationFlow** | **string** | Flow used ending the session from a provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**ClientNetworks** | Pointer to **string** | List of CIDRs (comma-separated) that clients can connect from. A more specific CIDR will match before a looser one. Clients connecting from a non-specified CIDR will be dropped. | [optional] 
**SharedSecret** | Pointer to **string** | Shared secret between clients and server to hash packets. | [optional] 
**MfaSupport** | Pointer to **bool** | When enabled, code-based multi-factor authentication can be used by appending a semicolon and the TOTP code to the password. This should only be enabled if all users that will bind to this provider have a TOTP device configured, as otherwise a password may incorrectly be rejected if it contains a semicolon. | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRadiusProviderRequest

`func NewRadiusProviderRequest(name string, authorizationFlow string, invalidationFlow string, ) *RadiusProviderRequest`

NewRadiusProviderRequest instantiates a new RadiusProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusProviderRequestWithDefaults

`func NewRadiusProviderRequestWithDefaults() *RadiusProviderRequest`

NewRadiusProviderRequestWithDefaults instantiates a new RadiusProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RadiusProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RadiusProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RadiusProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticationFlow

`func (o *RadiusProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *RadiusProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *RadiusProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *RadiusProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *RadiusProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *RadiusProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *RadiusProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *RadiusProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *RadiusProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetInvalidationFlow

`func (o *RadiusProviderRequest) GetInvalidationFlow() string`

GetInvalidationFlow returns the InvalidationFlow field if non-nil, zero value otherwise.

### GetInvalidationFlowOk

`func (o *RadiusProviderRequest) GetInvalidationFlowOk() (*string, bool)`

GetInvalidationFlowOk returns a tuple with the InvalidationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidationFlow

`func (o *RadiusProviderRequest) SetInvalidationFlow(v string)`

SetInvalidationFlow sets InvalidationFlow field to given value.


### GetPropertyMappings

`func (o *RadiusProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *RadiusProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *RadiusProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *RadiusProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetClientNetworks

`func (o *RadiusProviderRequest) GetClientNetworks() string`

GetClientNetworks returns the ClientNetworks field if non-nil, zero value otherwise.

### GetClientNetworksOk

`func (o *RadiusProviderRequest) GetClientNetworksOk() (*string, bool)`

GetClientNetworksOk returns a tuple with the ClientNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetworks

`func (o *RadiusProviderRequest) SetClientNetworks(v string)`

SetClientNetworks sets ClientNetworks field to given value.

### HasClientNetworks

`func (o *RadiusProviderRequest) HasClientNetworks() bool`

HasClientNetworks returns a boolean if a field has been set.

### GetSharedSecret

`func (o *RadiusProviderRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *RadiusProviderRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *RadiusProviderRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *RadiusProviderRequest) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetMfaSupport

`func (o *RadiusProviderRequest) GetMfaSupport() bool`

GetMfaSupport returns the MfaSupport field if non-nil, zero value otherwise.

### GetMfaSupportOk

`func (o *RadiusProviderRequest) GetMfaSupportOk() (*bool, bool)`

GetMfaSupportOk returns a tuple with the MfaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaSupport

`func (o *RadiusProviderRequest) SetMfaSupport(v bool)`

SetMfaSupport sets MfaSupport field to given value.

### HasMfaSupport

`func (o *RadiusProviderRequest) HasMfaSupport() bool`

HasMfaSupport returns a boolean if a field has been set.

### GetCertificate

`func (o *RadiusProviderRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *RadiusProviderRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *RadiusProviderRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *RadiusProviderRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *RadiusProviderRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *RadiusProviderRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


