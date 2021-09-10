# PatchedOAuth2ProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AuthorizationFlow** | Pointer to **string** | Flow used when authorizing this provider. | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**ClientType** | Pointer to [**ClientTypeEnum**](ClientTypeEnum.md) | Confidential clients are capable of maintaining the confidentiality     of their credentials. Public clients are incapable. | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**AccessCodeValidity** | Pointer to **string** | Access codes not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**TokenValidity** | Pointer to **string** | Tokens not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**IncludeClaimsInIdToken** | Pointer to **bool** | Include User claims from scopes in the id_token, for applications that don&#39;t access the userinfo endpoint. | [optional] 
**JwtAlg** | Pointer to [**JwtAlgEnum**](JwtAlgEnum.md) | Algorithm used to sign the JWT Token | [optional] 
**RsaKey** | Pointer to **NullableString** | Key used to sign the tokens. Only required when JWT Algorithm is set to RS256. | [optional] 
**RedirectUris** | Pointer to **string** | Enter each URI on a new line. | [optional] 
**SubMode** | Pointer to [**SubModeEnum**](SubModeEnum.md) | Configure what data should be used as unique User Identifier. For most cases, the default should be fine. | [optional] 
**IssuerMode** | Pointer to [**IssuerModeEnum**](IssuerModeEnum.md) | Configure how the issuer field of the ID Token should be filled. | [optional] 

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

### GetTokenValidity

`func (o *PatchedOAuth2ProviderRequest) GetTokenValidity() string`

GetTokenValidity returns the TokenValidity field if non-nil, zero value otherwise.

### GetTokenValidityOk

`func (o *PatchedOAuth2ProviderRequest) GetTokenValidityOk() (*string, bool)`

GetTokenValidityOk returns a tuple with the TokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenValidity

`func (o *PatchedOAuth2ProviderRequest) SetTokenValidity(v string)`

SetTokenValidity sets TokenValidity field to given value.

### HasTokenValidity

`func (o *PatchedOAuth2ProviderRequest) HasTokenValidity() bool`

HasTokenValidity returns a boolean if a field has been set.

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

### GetJwtAlg

`func (o *PatchedOAuth2ProviderRequest) GetJwtAlg() JwtAlgEnum`

GetJwtAlg returns the JwtAlg field if non-nil, zero value otherwise.

### GetJwtAlgOk

`func (o *PatchedOAuth2ProviderRequest) GetJwtAlgOk() (*JwtAlgEnum, bool)`

GetJwtAlgOk returns a tuple with the JwtAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAlg

`func (o *PatchedOAuth2ProviderRequest) SetJwtAlg(v JwtAlgEnum)`

SetJwtAlg sets JwtAlg field to given value.

### HasJwtAlg

`func (o *PatchedOAuth2ProviderRequest) HasJwtAlg() bool`

HasJwtAlg returns a boolean if a field has been set.

### GetRsaKey

`func (o *PatchedOAuth2ProviderRequest) GetRsaKey() string`

GetRsaKey returns the RsaKey field if non-nil, zero value otherwise.

### GetRsaKeyOk

`func (o *PatchedOAuth2ProviderRequest) GetRsaKeyOk() (*string, bool)`

GetRsaKeyOk returns a tuple with the RsaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaKey

`func (o *PatchedOAuth2ProviderRequest) SetRsaKey(v string)`

SetRsaKey sets RsaKey field to given value.

### HasRsaKey

`func (o *PatchedOAuth2ProviderRequest) HasRsaKey() bool`

HasRsaKey returns a boolean if a field has been set.

### SetRsaKeyNil

`func (o *PatchedOAuth2ProviderRequest) SetRsaKeyNil(b bool)`

 SetRsaKeyNil sets the value for RsaKey to be an explicit nil

### UnsetRsaKey
`func (o *PatchedOAuth2ProviderRequest) UnsetRsaKey()`

UnsetRsaKey ensures that no value is present for RsaKey, not even an explicit nil
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


