# OAuth2Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**Component** | **string** |  | [readonly] 
**AssignedApplicationSlug** | **string** | Internal application name, used in URLs. | [readonly] 
**AssignedApplicationName** | **string** | Application&#39;s display Name. | [readonly] 
**VerboseName** | **string** |  | [readonly] 
**VerboseNamePlural** | **string** |  | [readonly] 
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

### NewOAuth2Provider

`func NewOAuth2Provider(pk int32, name string, authorizationFlow string, component string, assignedApplicationSlug string, assignedApplicationName string, verboseName string, verboseNamePlural string, ) *OAuth2Provider`

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

### GetTokenValidity

`func (o *OAuth2Provider) GetTokenValidity() string`

GetTokenValidity returns the TokenValidity field if non-nil, zero value otherwise.

### GetTokenValidityOk

`func (o *OAuth2Provider) GetTokenValidityOk() (*string, bool)`

GetTokenValidityOk returns a tuple with the TokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenValidity

`func (o *OAuth2Provider) SetTokenValidity(v string)`

SetTokenValidity sets TokenValidity field to given value.

### HasTokenValidity

`func (o *OAuth2Provider) HasTokenValidity() bool`

HasTokenValidity returns a boolean if a field has been set.

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

### GetJwtAlg

`func (o *OAuth2Provider) GetJwtAlg() JwtAlgEnum`

GetJwtAlg returns the JwtAlg field if non-nil, zero value otherwise.

### GetJwtAlgOk

`func (o *OAuth2Provider) GetJwtAlgOk() (*JwtAlgEnum, bool)`

GetJwtAlgOk returns a tuple with the JwtAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAlg

`func (o *OAuth2Provider) SetJwtAlg(v JwtAlgEnum)`

SetJwtAlg sets JwtAlg field to given value.

### HasJwtAlg

`func (o *OAuth2Provider) HasJwtAlg() bool`

HasJwtAlg returns a boolean if a field has been set.

### GetRsaKey

`func (o *OAuth2Provider) GetRsaKey() string`

GetRsaKey returns the RsaKey field if non-nil, zero value otherwise.

### GetRsaKeyOk

`func (o *OAuth2Provider) GetRsaKeyOk() (*string, bool)`

GetRsaKeyOk returns a tuple with the RsaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaKey

`func (o *OAuth2Provider) SetRsaKey(v string)`

SetRsaKey sets RsaKey field to given value.

### HasRsaKey

`func (o *OAuth2Provider) HasRsaKey() bool`

HasRsaKey returns a boolean if a field has been set.

### SetRsaKeyNil

`func (o *OAuth2Provider) SetRsaKeyNil(b bool)`

 SetRsaKeyNil sets the value for RsaKey to be an explicit nil

### UnsetRsaKey
`func (o *OAuth2Provider) UnsetRsaKey()`

UnsetRsaKey ensures that no value is present for RsaKey, not even an explicit nil
### GetRedirectUris

`func (o *OAuth2Provider) GetRedirectUris() string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OAuth2Provider) GetRedirectUrisOk() (*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OAuth2Provider) SetRedirectUris(v string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *OAuth2Provider) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


