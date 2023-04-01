# PatchedOAuth2ProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | Pointer to **string** | Flow used when authorizing this provider. | [optional] 
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

### NewPatchedOAuth2ProviderRequest

`func NewPatchedOAuth2ProviderRequest() *PatchedOAuth2ProviderRequest`

NewPatchedOAuth2ProviderRequest instantiates a new PatchedOAuth2ProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedOAuth2ProviderRequestWithDefaults

`func NewPatchedOAuth2ProviderRequestWithDefaults() *PatchedOAuth2ProviderRequest`

NewPatchedOAuth2ProviderRequestWithDefaults instantiates a new PatchedOAuth2ProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedOAuth2ProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedOAuth2ProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedOAuth2ProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedOAuth2ProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PatchedOAuth2ProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedOAuth2ProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedOAuth2ProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedOAuth2ProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedOAuth2ProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedOAuth2ProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *PatchedOAuth2ProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *PatchedOAuth2ProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *PatchedOAuth2ProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.

### HasAuthorizationFlow

`func (o *PatchedOAuth2ProviderRequest) HasAuthorizationFlow() bool`

HasAuthorizationFlow returns a boolean if a field has been set.

### GetPropertyMappings

`func (o *PatchedOAuth2ProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *PatchedOAuth2ProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *PatchedOAuth2ProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *PatchedOAuth2ProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetClientType

`func (o *PatchedOAuth2ProviderRequest) GetClientType() ClientTypeEnum`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *PatchedOAuth2ProviderRequest) GetClientTypeOk() (*ClientTypeEnum, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *PatchedOAuth2ProviderRequest) SetClientType(v ClientTypeEnum)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *PatchedOAuth2ProviderRequest) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetClientId

`func (o *PatchedOAuth2ProviderRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PatchedOAuth2ProviderRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PatchedOAuth2ProviderRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PatchedOAuth2ProviderRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *PatchedOAuth2ProviderRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PatchedOAuth2ProviderRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PatchedOAuth2ProviderRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PatchedOAuth2ProviderRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetAccessCodeValidity

`func (o *PatchedOAuth2ProviderRequest) GetAccessCodeValidity() string`

GetAccessCodeValidity returns the AccessCodeValidity field if non-nil, zero value otherwise.

### GetAccessCodeValidityOk

`func (o *PatchedOAuth2ProviderRequest) GetAccessCodeValidityOk() (*string, bool)`

GetAccessCodeValidityOk returns a tuple with the AccessCodeValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCodeValidity

`func (o *PatchedOAuth2ProviderRequest) SetAccessCodeValidity(v string)`

SetAccessCodeValidity sets AccessCodeValidity field to given value.

### HasAccessCodeValidity

`func (o *PatchedOAuth2ProviderRequest) HasAccessCodeValidity() bool`

HasAccessCodeValidity returns a boolean if a field has been set.

### GetAccessTokenValidity

`func (o *PatchedOAuth2ProviderRequest) GetAccessTokenValidity() string`

GetAccessTokenValidity returns the AccessTokenValidity field if non-nil, zero value otherwise.

### GetAccessTokenValidityOk

`func (o *PatchedOAuth2ProviderRequest) GetAccessTokenValidityOk() (*string, bool)`

GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidity

`func (o *PatchedOAuth2ProviderRequest) SetAccessTokenValidity(v string)`

SetAccessTokenValidity sets AccessTokenValidity field to given value.

### HasAccessTokenValidity

`func (o *PatchedOAuth2ProviderRequest) HasAccessTokenValidity() bool`

HasAccessTokenValidity returns a boolean if a field has been set.

### GetRefreshTokenValidity

`func (o *PatchedOAuth2ProviderRequest) GetRefreshTokenValidity() string`

GetRefreshTokenValidity returns the RefreshTokenValidity field if non-nil, zero value otherwise.

### GetRefreshTokenValidityOk

`func (o *PatchedOAuth2ProviderRequest) GetRefreshTokenValidityOk() (*string, bool)`

GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValidity

`func (o *PatchedOAuth2ProviderRequest) SetRefreshTokenValidity(v string)`

SetRefreshTokenValidity sets RefreshTokenValidity field to given value.

### HasRefreshTokenValidity

`func (o *PatchedOAuth2ProviderRequest) HasRefreshTokenValidity() bool`

HasRefreshTokenValidity returns a boolean if a field has been set.

### GetIncludeClaimsInIdToken

`func (o *PatchedOAuth2ProviderRequest) GetIncludeClaimsInIdToken() bool`

GetIncludeClaimsInIdToken returns the IncludeClaimsInIdToken field if non-nil, zero value otherwise.

### GetIncludeClaimsInIdTokenOk

`func (o *PatchedOAuth2ProviderRequest) GetIncludeClaimsInIdTokenOk() (*bool, bool)`

GetIncludeClaimsInIdTokenOk returns a tuple with the IncludeClaimsInIdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeClaimsInIdToken

`func (o *PatchedOAuth2ProviderRequest) SetIncludeClaimsInIdToken(v bool)`

SetIncludeClaimsInIdToken sets IncludeClaimsInIdToken field to given value.

### HasIncludeClaimsInIdToken

`func (o *PatchedOAuth2ProviderRequest) HasIncludeClaimsInIdToken() bool`

HasIncludeClaimsInIdToken returns a boolean if a field has been set.

### GetSigningKey

`func (o *PatchedOAuth2ProviderRequest) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *PatchedOAuth2ProviderRequest) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *PatchedOAuth2ProviderRequest) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.

### HasSigningKey

`func (o *PatchedOAuth2ProviderRequest) HasSigningKey() bool`

HasSigningKey returns a boolean if a field has been set.

### SetSigningKeyNil

`func (o *PatchedOAuth2ProviderRequest) SetSigningKeyNil(b bool)`

 SetSigningKeyNil sets the value for SigningKey to be an explicit nil

### UnsetSigningKey
`func (o *PatchedOAuth2ProviderRequest) UnsetSigningKey()`

UnsetSigningKey ensures that no value is present for SigningKey, not even an explicit nil
### GetRedirectUris

`func (o *PatchedOAuth2ProviderRequest) GetRedirectUris() string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *PatchedOAuth2ProviderRequest) GetRedirectUrisOk() (*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *PatchedOAuth2ProviderRequest) SetRedirectUris(v string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *PatchedOAuth2ProviderRequest) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetSubMode

`func (o *PatchedOAuth2ProviderRequest) GetSubMode() SubModeEnum`

GetSubMode returns the SubMode field if non-nil, zero value otherwise.

### GetSubModeOk

`func (o *PatchedOAuth2ProviderRequest) GetSubModeOk() (*SubModeEnum, bool)`

GetSubModeOk returns a tuple with the SubMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMode

`func (o *PatchedOAuth2ProviderRequest) SetSubMode(v SubModeEnum)`

SetSubMode sets SubMode field to given value.

### HasSubMode

`func (o *PatchedOAuth2ProviderRequest) HasSubMode() bool`

HasSubMode returns a boolean if a field has been set.

### GetIssuerMode

`func (o *PatchedOAuth2ProviderRequest) GetIssuerMode() IssuerModeEnum`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *PatchedOAuth2ProviderRequest) GetIssuerModeOk() (*IssuerModeEnum, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *PatchedOAuth2ProviderRequest) SetIssuerMode(v IssuerModeEnum)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *PatchedOAuth2ProviderRequest) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetJwksSources

`func (o *PatchedOAuth2ProviderRequest) GetJwksSources() []string`

GetJwksSources returns the JwksSources field if non-nil, zero value otherwise.

### GetJwksSourcesOk

`func (o *PatchedOAuth2ProviderRequest) GetJwksSourcesOk() (*[]string, bool)`

GetJwksSourcesOk returns a tuple with the JwksSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksSources

`func (o *PatchedOAuth2ProviderRequest) SetJwksSources(v []string)`

SetJwksSources sets JwksSources field to given value.

### HasJwksSources

`func (o *PatchedOAuth2ProviderRequest) HasJwksSources() bool`

HasJwksSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


