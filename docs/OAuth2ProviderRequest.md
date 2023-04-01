# OAuth2ProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**ClientType** | Pointer to [**ClientTypeEnum**](ClientTypeEnum.md) |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**AccessCodeValidity** | Pointer to **string** | Access codes not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**AccessTokenValidity** | Pointer to **string** | Tokens not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**RefreshTokenValidity** | Pointer to **string** | Tokens not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**IncludeClaimsInIdToken** | Pointer to **bool** | Include User claims from scopes in the id_token, for applications that don&#39;t access the userinfo endpoint. | [optional] 
**SigningKey** | Pointer to **NullableString** | Key used to sign the tokens. Only required when JWT Algorithm is set to RS256. | [optional] 
**RedirectUris** | Pointer to **string** | Enter each URI on a new line. | [optional] 
**SubMode** | Pointer to [**SubModeEnum**](SubModeEnum.md) |  | [optional] 
**IssuerMode** | Pointer to [**IssuerModeEnum**](IssuerModeEnum.md) |  | [optional] 
**JwksSources** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOAuth2ProviderRequest

`func NewOAuth2ProviderRequest(name string, authorizationFlow string, ) *OAuth2ProviderRequest`

NewOAuth2ProviderRequest instantiates a new OAuth2ProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ProviderRequestWithDefaults

`func NewOAuth2ProviderRequestWithDefaults() *OAuth2ProviderRequest`

NewOAuth2ProviderRequestWithDefaults instantiates a new OAuth2ProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OAuth2ProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuth2ProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuth2ProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticationFlow

`func (o *OAuth2ProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *OAuth2ProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *OAuth2ProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *OAuth2ProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *OAuth2ProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *OAuth2ProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *OAuth2ProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *OAuth2ProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *OAuth2ProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetPropertyMappings

`func (o *OAuth2ProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *OAuth2ProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *OAuth2ProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *OAuth2ProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetClientType

`func (o *OAuth2ProviderRequest) GetClientType() ClientTypeEnum`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *OAuth2ProviderRequest) GetClientTypeOk() (*ClientTypeEnum, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *OAuth2ProviderRequest) SetClientType(v ClientTypeEnum)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *OAuth2ProviderRequest) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetClientId

`func (o *OAuth2ProviderRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2ProviderRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2ProviderRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuth2ProviderRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OAuth2ProviderRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2ProviderRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2ProviderRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuth2ProviderRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetAccessCodeValidity

`func (o *OAuth2ProviderRequest) GetAccessCodeValidity() string`

GetAccessCodeValidity returns the AccessCodeValidity field if non-nil, zero value otherwise.

### GetAccessCodeValidityOk

`func (o *OAuth2ProviderRequest) GetAccessCodeValidityOk() (*string, bool)`

GetAccessCodeValidityOk returns a tuple with the AccessCodeValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCodeValidity

`func (o *OAuth2ProviderRequest) SetAccessCodeValidity(v string)`

SetAccessCodeValidity sets AccessCodeValidity field to given value.

### HasAccessCodeValidity

`func (o *OAuth2ProviderRequest) HasAccessCodeValidity() bool`

HasAccessCodeValidity returns a boolean if a field has been set.

### GetAccessTokenValidity

`func (o *OAuth2ProviderRequest) GetAccessTokenValidity() string`

GetAccessTokenValidity returns the AccessTokenValidity field if non-nil, zero value otherwise.

### GetAccessTokenValidityOk

`func (o *OAuth2ProviderRequest) GetAccessTokenValidityOk() (*string, bool)`

GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidity

`func (o *OAuth2ProviderRequest) SetAccessTokenValidity(v string)`

SetAccessTokenValidity sets AccessTokenValidity field to given value.

### HasAccessTokenValidity

`func (o *OAuth2ProviderRequest) HasAccessTokenValidity() bool`

HasAccessTokenValidity returns a boolean if a field has been set.

### GetRefreshTokenValidity

`func (o *OAuth2ProviderRequest) GetRefreshTokenValidity() string`

GetRefreshTokenValidity returns the RefreshTokenValidity field if non-nil, zero value otherwise.

### GetRefreshTokenValidityOk

`func (o *OAuth2ProviderRequest) GetRefreshTokenValidityOk() (*string, bool)`

GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValidity

`func (o *OAuth2ProviderRequest) SetRefreshTokenValidity(v string)`

SetRefreshTokenValidity sets RefreshTokenValidity field to given value.

### HasRefreshTokenValidity

`func (o *OAuth2ProviderRequest) HasRefreshTokenValidity() bool`

HasRefreshTokenValidity returns a boolean if a field has been set.

### GetIncludeClaimsInIdToken

`func (o *OAuth2ProviderRequest) GetIncludeClaimsInIdToken() bool`

GetIncludeClaimsInIdToken returns the IncludeClaimsInIdToken field if non-nil, zero value otherwise.

### GetIncludeClaimsInIdTokenOk

`func (o *OAuth2ProviderRequest) GetIncludeClaimsInIdTokenOk() (*bool, bool)`

GetIncludeClaimsInIdTokenOk returns a tuple with the IncludeClaimsInIdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeClaimsInIdToken

`func (o *OAuth2ProviderRequest) SetIncludeClaimsInIdToken(v bool)`

SetIncludeClaimsInIdToken sets IncludeClaimsInIdToken field to given value.

### HasIncludeClaimsInIdToken

`func (o *OAuth2ProviderRequest) HasIncludeClaimsInIdToken() bool`

HasIncludeClaimsInIdToken returns a boolean if a field has been set.

### GetSigningKey

`func (o *OAuth2ProviderRequest) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *OAuth2ProviderRequest) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *OAuth2ProviderRequest) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.

### HasSigningKey

`func (o *OAuth2ProviderRequest) HasSigningKey() bool`

HasSigningKey returns a boolean if a field has been set.

### SetSigningKeyNil

`func (o *OAuth2ProviderRequest) SetSigningKeyNil(b bool)`

 SetSigningKeyNil sets the value for SigningKey to be an explicit nil

### UnsetSigningKey
`func (o *OAuth2ProviderRequest) UnsetSigningKey()`

UnsetSigningKey ensures that no value is present for SigningKey, not even an explicit nil
### GetRedirectUris

`func (o *OAuth2ProviderRequest) GetRedirectUris() string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OAuth2ProviderRequest) GetRedirectUrisOk() (*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OAuth2ProviderRequest) SetRedirectUris(v string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *OAuth2ProviderRequest) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetSubMode

`func (o *OAuth2ProviderRequest) GetSubMode() SubModeEnum`

GetSubMode returns the SubMode field if non-nil, zero value otherwise.

### GetSubModeOk

`func (o *OAuth2ProviderRequest) GetSubModeOk() (*SubModeEnum, bool)`

GetSubModeOk returns a tuple with the SubMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMode

`func (o *OAuth2ProviderRequest) SetSubMode(v SubModeEnum)`

SetSubMode sets SubMode field to given value.

### HasSubMode

`func (o *OAuth2ProviderRequest) HasSubMode() bool`

HasSubMode returns a boolean if a field has been set.

### GetIssuerMode

`func (o *OAuth2ProviderRequest) GetIssuerMode() IssuerModeEnum`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *OAuth2ProviderRequest) GetIssuerModeOk() (*IssuerModeEnum, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *OAuth2ProviderRequest) SetIssuerMode(v IssuerModeEnum)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *OAuth2ProviderRequest) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetJwksSources

`func (o *OAuth2ProviderRequest) GetJwksSources() []string`

GetJwksSources returns the JwksSources field if non-nil, zero value otherwise.

### GetJwksSourcesOk

`func (o *OAuth2ProviderRequest) GetJwksSourcesOk() (*[]string, bool)`

GetJwksSourcesOk returns a tuple with the JwksSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksSources

`func (o *OAuth2ProviderRequest) SetJwksSources(v []string)`

SetJwksSources sets JwksSources field to given value.

### HasJwksSources

`func (o *OAuth2ProviderRequest) HasJwksSources() bool`

HasJwksSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


