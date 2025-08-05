# OAuth2Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**InvalidationFlow** | **string** | Flow used ending the session from a provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**AssignedApplicationSlug** | **string** | Internal application name, used in URLs. | [readonly] 
**AssignedApplicationName** | **string** | Application&#39;s display Name. | [readonly] 
**AssignedBackchannelApplicationSlug** | **string** | Internal application name, used in URLs. | [readonly] 
**AssignedBackchannelApplicationName** | **string** | Application&#39;s display Name. | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**ClientType** | Pointer to [**ClientTypeEnum**](ClientTypeEnum.md) | Confidential clients are capable of maintaining the confidentiality of their credentials. Public clients are incapable | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**AccessCodeValidity** | Pointer to **string** | Access codes not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**AccessTokenValidity** | Pointer to **string** | Tokens not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**RefreshTokenValidity** | Pointer to **string** | Tokens not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**IncludeClaimsInIdToken** | Pointer to **bool** | Include User claims from scopes in the id_token, for applications that don&#39;t access the userinfo endpoint. | [optional] 
**SigningKey** | Pointer to **NullableString** | Key used to sign the tokens. | [optional] 
**EncryptionKey** | Pointer to **NullableString** | Key used to encrypt the tokens. When set, tokens will be encrypted and returned as JWEs. | [optional] 
**RedirectUris** | [**[]RedirectURI**](RedirectURI.md) |  | 
**BackchannelLogoutUri** | Pointer to **string** |  | [optional] 
**SubMode** | Pointer to [**SubModeEnum**](SubModeEnum.md) | Configure what data should be used as unique User Identifier. For most cases, the default should be fine. | [optional] 
**IssuerMode** | Pointer to [**IssuerModeEnum**](IssuerModeEnum.md) | Configure how the issuer field of the ID Token should be filled. | [optional] 
**JwtFederationSources** | Pointer to **[]string** |  | [optional] 
**JwtFederationProviders** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewOAuth2Provider

`func NewOAuth2Provider(pk int32, name string, authorizationFlow string, invalidationFlow string, component string, assignedApplicationSlug string, assignedApplicationName string, assignedBackchannelApplicationSlug string, assignedBackchannelApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, redirectUris []RedirectURI, ) *OAuth2Provider`

NewOAuth2Provider instantiates a new OAuth2Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ProviderWithDefaults

`func NewOAuth2ProviderWithDefaults() *OAuth2Provider`

NewOAuth2ProviderWithDefaults instantiates a new OAuth2Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *OAuth2Provider) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *OAuth2Provider) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *OAuth2Provider) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *OAuth2Provider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuth2Provider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuth2Provider) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticationFlow

`func (o *OAuth2Provider) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *OAuth2Provider) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *OAuth2Provider) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *OAuth2Provider) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *OAuth2Provider) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *OAuth2Provider) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *OAuth2Provider) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *OAuth2Provider) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *OAuth2Provider) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetInvalidationFlow

`func (o *OAuth2Provider) GetInvalidationFlow() string`

GetInvalidationFlow returns the InvalidationFlow field if non-nil, zero value otherwise.

### GetInvalidationFlowOk

`func (o *OAuth2Provider) GetInvalidationFlowOk() (*string, bool)`

GetInvalidationFlowOk returns a tuple with the InvalidationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidationFlow

`func (o *OAuth2Provider) SetInvalidationFlow(v string)`

SetInvalidationFlow sets InvalidationFlow field to given value.


### GetPropertyMappings

`func (o *OAuth2Provider) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *OAuth2Provider) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *OAuth2Provider) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *OAuth2Provider) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetComponent

`func (o *OAuth2Provider) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *OAuth2Provider) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *OAuth2Provider) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetAssignedApplicationSlug

`func (o *OAuth2Provider) GetAssignedApplicationSlug() string`

GetAssignedApplicationSlug returns the AssignedApplicationSlug field if non-nil, zero value otherwise.

### GetAssignedApplicationSlugOk

`func (o *OAuth2Provider) GetAssignedApplicationSlugOk() (*string, bool)`

GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedApplicationSlug

`func (o *OAuth2Provider) SetAssignedApplicationSlug(v string)`

SetAssignedApplicationSlug sets AssignedApplicationSlug field to given value.


### GetAssignedApplicationName

`func (o *OAuth2Provider) GetAssignedApplicationName() string`

GetAssignedApplicationName returns the AssignedApplicationName field if non-nil, zero value otherwise.

### GetAssignedApplicationNameOk

`func (o *OAuth2Provider) GetAssignedApplicationNameOk() (*string, bool)`

GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedApplicationName

`func (o *OAuth2Provider) SetAssignedApplicationName(v string)`

SetAssignedApplicationName sets AssignedApplicationName field to given value.


### GetAssignedBackchannelApplicationSlug

`func (o *OAuth2Provider) GetAssignedBackchannelApplicationSlug() string`

GetAssignedBackchannelApplicationSlug returns the AssignedBackchannelApplicationSlug field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationSlugOk

`func (o *OAuth2Provider) GetAssignedBackchannelApplicationSlugOk() (*string, bool)`

GetAssignedBackchannelApplicationSlugOk returns a tuple with the AssignedBackchannelApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationSlug

`func (o *OAuth2Provider) SetAssignedBackchannelApplicationSlug(v string)`

SetAssignedBackchannelApplicationSlug sets AssignedBackchannelApplicationSlug field to given value.


### GetAssignedBackchannelApplicationName

`func (o *OAuth2Provider) GetAssignedBackchannelApplicationName() string`

GetAssignedBackchannelApplicationName returns the AssignedBackchannelApplicationName field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationNameOk

`func (o *OAuth2Provider) GetAssignedBackchannelApplicationNameOk() (*string, bool)`

GetAssignedBackchannelApplicationNameOk returns a tuple with the AssignedBackchannelApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationName

`func (o *OAuth2Provider) SetAssignedBackchannelApplicationName(v string)`

SetAssignedBackchannelApplicationName sets AssignedBackchannelApplicationName field to given value.


### GetVerboseName

`func (o *OAuth2Provider) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *OAuth2Provider) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *OAuth2Provider) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *OAuth2Provider) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *OAuth2Provider) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *OAuth2Provider) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *OAuth2Provider) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *OAuth2Provider) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *OAuth2Provider) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetClientType

`func (o *OAuth2Provider) GetClientType() ClientTypeEnum`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *OAuth2Provider) GetClientTypeOk() (*ClientTypeEnum, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *OAuth2Provider) SetClientType(v ClientTypeEnum)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *OAuth2Provider) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetClientId

`func (o *OAuth2Provider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2Provider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2Provider) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuth2Provider) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OAuth2Provider) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2Provider) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2Provider) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuth2Provider) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetAccessCodeValidity

`func (o *OAuth2Provider) GetAccessCodeValidity() string`

GetAccessCodeValidity returns the AccessCodeValidity field if non-nil, zero value otherwise.

### GetAccessCodeValidityOk

`func (o *OAuth2Provider) GetAccessCodeValidityOk() (*string, bool)`

GetAccessCodeValidityOk returns a tuple with the AccessCodeValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCodeValidity

`func (o *OAuth2Provider) SetAccessCodeValidity(v string)`

SetAccessCodeValidity sets AccessCodeValidity field to given value.

### HasAccessCodeValidity

`func (o *OAuth2Provider) HasAccessCodeValidity() bool`

HasAccessCodeValidity returns a boolean if a field has been set.

### GetAccessTokenValidity

`func (o *OAuth2Provider) GetAccessTokenValidity() string`

GetAccessTokenValidity returns the AccessTokenValidity field if non-nil, zero value otherwise.

### GetAccessTokenValidityOk

`func (o *OAuth2Provider) GetAccessTokenValidityOk() (*string, bool)`

GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidity

`func (o *OAuth2Provider) SetAccessTokenValidity(v string)`

SetAccessTokenValidity sets AccessTokenValidity field to given value.

### HasAccessTokenValidity

`func (o *OAuth2Provider) HasAccessTokenValidity() bool`

HasAccessTokenValidity returns a boolean if a field has been set.

### GetRefreshTokenValidity

`func (o *OAuth2Provider) GetRefreshTokenValidity() string`

GetRefreshTokenValidity returns the RefreshTokenValidity field if non-nil, zero value otherwise.

### GetRefreshTokenValidityOk

`func (o *OAuth2Provider) GetRefreshTokenValidityOk() (*string, bool)`

GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValidity

`func (o *OAuth2Provider) SetRefreshTokenValidity(v string)`

SetRefreshTokenValidity sets RefreshTokenValidity field to given value.

### HasRefreshTokenValidity

`func (o *OAuth2Provider) HasRefreshTokenValidity() bool`

HasRefreshTokenValidity returns a boolean if a field has been set.

### GetIncludeClaimsInIdToken

`func (o *OAuth2Provider) GetIncludeClaimsInIdToken() bool`

GetIncludeClaimsInIdToken returns the IncludeClaimsInIdToken field if non-nil, zero value otherwise.

### GetIncludeClaimsInIdTokenOk

`func (o *OAuth2Provider) GetIncludeClaimsInIdTokenOk() (*bool, bool)`

GetIncludeClaimsInIdTokenOk returns a tuple with the IncludeClaimsInIdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeClaimsInIdToken

`func (o *OAuth2Provider) SetIncludeClaimsInIdToken(v bool)`

SetIncludeClaimsInIdToken sets IncludeClaimsInIdToken field to given value.

### HasIncludeClaimsInIdToken

`func (o *OAuth2Provider) HasIncludeClaimsInIdToken() bool`

HasIncludeClaimsInIdToken returns a boolean if a field has been set.

### GetSigningKey

`func (o *OAuth2Provider) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *OAuth2Provider) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *OAuth2Provider) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.

### HasSigningKey

`func (o *OAuth2Provider) HasSigningKey() bool`

HasSigningKey returns a boolean if a field has been set.

### SetSigningKeyNil

`func (o *OAuth2Provider) SetSigningKeyNil(b bool)`

 SetSigningKeyNil sets the value for SigningKey to be an explicit nil

### UnsetSigningKey
`func (o *OAuth2Provider) UnsetSigningKey()`

UnsetSigningKey ensures that no value is present for SigningKey, not even an explicit nil
### GetEncryptionKey

`func (o *OAuth2Provider) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *OAuth2Provider) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *OAuth2Provider) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *OAuth2Provider) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### SetEncryptionKeyNil

`func (o *OAuth2Provider) SetEncryptionKeyNil(b bool)`

 SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil

### UnsetEncryptionKey
`func (o *OAuth2Provider) UnsetEncryptionKey()`

UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
### GetRedirectUris

`func (o *OAuth2Provider) GetRedirectUris() []RedirectURI`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OAuth2Provider) GetRedirectUrisOk() (*[]RedirectURI, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OAuth2Provider) SetRedirectUris(v []RedirectURI)`

SetRedirectUris sets RedirectUris field to given value.


### GetBackchannelLogoutUri

`func (o *OAuth2Provider) GetBackchannelLogoutUri() string`

GetBackchannelLogoutUri returns the BackchannelLogoutUri field if non-nil, zero value otherwise.

### GetBackchannelLogoutUriOk

`func (o *OAuth2Provider) GetBackchannelLogoutUriOk() (*string, bool)`

GetBackchannelLogoutUriOk returns a tuple with the BackchannelLogoutUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelLogoutUri

`func (o *OAuth2Provider) SetBackchannelLogoutUri(v string)`

SetBackchannelLogoutUri sets BackchannelLogoutUri field to given value.

### HasBackchannelLogoutUri

`func (o *OAuth2Provider) HasBackchannelLogoutUri() bool`

HasBackchannelLogoutUri returns a boolean if a field has been set.

### GetSubMode

`func (o *OAuth2Provider) GetSubMode() SubModeEnum`

GetSubMode returns the SubMode field if non-nil, zero value otherwise.

### GetSubModeOk

`func (o *OAuth2Provider) GetSubModeOk() (*SubModeEnum, bool)`

GetSubModeOk returns a tuple with the SubMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMode

`func (o *OAuth2Provider) SetSubMode(v SubModeEnum)`

SetSubMode sets SubMode field to given value.

### HasSubMode

`func (o *OAuth2Provider) HasSubMode() bool`

HasSubMode returns a boolean if a field has been set.

### GetIssuerMode

`func (o *OAuth2Provider) GetIssuerMode() IssuerModeEnum`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *OAuth2Provider) GetIssuerModeOk() (*IssuerModeEnum, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *OAuth2Provider) SetIssuerMode(v IssuerModeEnum)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *OAuth2Provider) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetJwtFederationSources

`func (o *OAuth2Provider) GetJwtFederationSources() []string`

GetJwtFederationSources returns the JwtFederationSources field if non-nil, zero value otherwise.

### GetJwtFederationSourcesOk

`func (o *OAuth2Provider) GetJwtFederationSourcesOk() (*[]string, bool)`

GetJwtFederationSourcesOk returns a tuple with the JwtFederationSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtFederationSources

`func (o *OAuth2Provider) SetJwtFederationSources(v []string)`

SetJwtFederationSources sets JwtFederationSources field to given value.

### HasJwtFederationSources

`func (o *OAuth2Provider) HasJwtFederationSources() bool`

HasJwtFederationSources returns a boolean if a field has been set.

### GetJwtFederationProviders

`func (o *OAuth2Provider) GetJwtFederationProviders() []int32`

GetJwtFederationProviders returns the JwtFederationProviders field if non-nil, zero value otherwise.

### GetJwtFederationProvidersOk

`func (o *OAuth2Provider) GetJwtFederationProvidersOk() (*[]int32, bool)`

GetJwtFederationProvidersOk returns a tuple with the JwtFederationProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtFederationProviders

`func (o *OAuth2Provider) SetJwtFederationProviders(v []int32)`

SetJwtFederationProviders sets JwtFederationProviders field to given value.

### HasJwtFederationProviders

`func (o *OAuth2Provider) HasJwtFederationProviders() bool`

HasJwtFederationProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


