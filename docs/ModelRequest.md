# ModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**BaseDn** | Pointer to **string** | DN under which objects are accessible. | [optional] 
**SearchGroup** | Pointer to **NullableString** | Users in this group can do search queries. If not set, every user can execute search queries. | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**TlsServerName** | Pointer to **string** |  | [optional] 
**UidStartNumber** | Pointer to **int32** | The start for uidNumbers, this number is added to the user.pk to make sure that the numbers aren&#39;t too low for POSIX users. Default is 2000 to ensure that we don&#39;t collide with local users uidNumber | [optional] 
**GidStartNumber** | Pointer to **int32** | The start for gidNumbers, this number is added to a number generated from the group.pk to make sure that the numbers aren&#39;t too low for POSIX groups. Default is 4000 to ensure that we don&#39;t collide with local groups or users primary groups gidNumber | [optional] 
**SearchMode** | Pointer to [**LDAPAPIAccessMode**](LDAPAPIAccessMode.md) |  | [optional] 
**BindMode** | Pointer to [**LDAPAPIAccessMode**](LDAPAPIAccessMode.md) |  | [optional] 
**MfaSupport** | Pointer to **bool** | When enabled, code-based multi-factor authentication can be used by appending a semicolon and the TOTP code to the password. This should only be enabled if all users that will bind to this provider have a TOTP device configured, as otherwise a password may incorrectly be rejected if it contains a semicolon. | [optional] 
**ClientType** | Pointer to [**ClientTypeEnum**](ClientTypeEnum.md) | Confidential clients are capable of maintaining the confidentiality of their credentials. Public clients are incapable  * &#x60;confidential&#x60; - Confidential * &#x60;public&#x60; - Public | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**AccessCodeValidity** | Pointer to **string** | Access codes not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**AccessTokenValidity** | Pointer to **string** | Tokens not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**RefreshTokenValidity** | Pointer to **string** | Tokens not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**IncludeClaimsInIdToken** | Pointer to **bool** | Include User claims from scopes in the id_token, for applications that don&#39;t access the userinfo endpoint. | [optional] 
**SigningKey** | Pointer to **NullableString** | Key used to sign the tokens. Only required when JWT Algorithm is set to RS256. | [optional] 
**RedirectUris** | Pointer to **string** | Enter each URI on a new line. | [optional] 
**SubMode** | Pointer to [**SubModeEnum**](SubModeEnum.md) | Configure what data should be used as unique User Identifier. For most cases, the default should be fine.  * &#x60;hashed_user_id&#x60; - Based on the Hashed User ID * &#x60;user_id&#x60; - Based on user ID * &#x60;user_uuid&#x60; - Based on user UUID * &#x60;user_username&#x60; - Based on the username * &#x60;user_email&#x60; - Based on the User&#39;s Email. This is recommended over the UPN method. * &#x60;user_upn&#x60; - Based on the User&#39;s UPN, only works if user has a &#39;upn&#39; attribute set. Use this method only if you have different UPN and Mail domains. | [optional] 
**IssuerMode** | Pointer to [**IssuerModeEnum**](IssuerModeEnum.md) | Configure how the issuer field of the ID Token should be filled.  * &#x60;global&#x60; - Same identifier is used for all providers * &#x60;per_provider&#x60; - Each provider has a different issuer, based on the application slug. | [optional] 
**JwksSources** | Pointer to **[]string** |  | [optional] 
**InternalHost** | Pointer to **string** |  | [optional] 
**ExternalHost** | **string** |  | 
**InternalHostSslValidation** | Pointer to **bool** | Validate SSL Certificates of upstream servers | [optional] 
**SkipPathRegex** | Pointer to **string** | Regular expressions for which authentication is not required. Each new line is interpreted as a new Regular Expression. | [optional] 
**BasicAuthEnabled** | Pointer to **bool** | Set a custom HTTP-Basic Authentication header based on values from authentik. | [optional] 
**BasicAuthPasswordAttribute** | Pointer to **string** | User/Group Attribute used for the password part of the HTTP-Basic Header. | [optional] 
**BasicAuthUserAttribute** | Pointer to **string** | User/Group Attribute used for the user part of the HTTP-Basic Header. If not set, the user&#39;s Email address is used. | [optional] 
**Mode** | Pointer to [**ProxyMode**](ProxyMode.md) | Enable support for forwardAuth in traefik and nginx auth_request. Exclusive with internal_host.  * &#x60;proxy&#x60; - Proxy * &#x60;forward_single&#x60; - Forward Single * &#x60;forward_domain&#x60; - Forward Domain | [optional] 
**InterceptHeaderAuth** | Pointer to **bool** | When enabled, this provider will intercept the authorization header and authenticate requests based on its value. | [optional] 
**CookieDomain** | Pointer to **string** |  | [optional] 
**ClientNetworks** | Pointer to **string** | List of CIDRs (comma-separated) that clients can connect from. A more specific CIDR will match before a looser one. Clients connecting from a non-specified CIDR will be dropped. | [optional] 
**SharedSecret** | Pointer to **string** | Shared secret between clients and server to hash packets. | [optional] 
**AcsUrl** | **string** |  | 
**Audience** | Pointer to **string** | Value of the audience restriction field of the assertion. When left empty, no audience restriction will be added. | [optional] 
**Issuer** | Pointer to **string** | Also known as EntityID | [optional] 
**AssertionValidNotBefore** | Pointer to **string** | Assertion valid not before current time + this value (Format: hours&#x3D;-1;minutes&#x3D;-2;seconds&#x3D;-3). | [optional] 
**AssertionValidNotOnOrAfter** | Pointer to **string** | Assertion not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**SessionValidNotOnOrAfter** | Pointer to **string** | Session not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**NameIdMapping** | Pointer to **NullableString** | Configure how the NameID value will be created. When left empty, the NameIDPolicy of the incoming request will be considered | [optional] 
**DigestAlgorithm** | Pointer to [**DigestAlgorithmEnum**](DigestAlgorithmEnum.md) |  | [optional] 
**SignatureAlgorithm** | Pointer to [**SignatureAlgorithmEnum**](SignatureAlgorithmEnum.md) |  | [optional] 
**SigningKp** | Pointer to **NullableString** | Keypair used to sign outgoing Responses going to the Service Provider. | [optional] 
**VerificationKp** | Pointer to **NullableString** | When selected, incoming assertion&#39;s Signatures will be validated against this certificate. To allow unsigned Requests, leave on default. | [optional] 
**SpBinding** | Pointer to [**SpBindingEnum**](SpBindingEnum.md) | This determines how authentik sends the response back to the Service Provider.  * &#x60;redirect&#x60; - Redirect * &#x60;post&#x60; - Post | [optional] 
**PropertyMappingsGroup** | Pointer to **[]string** | Property mappings used for group creation/updating. | [optional] 
**Url** | **string** | Base URL to SCIM requests, usually ends in /v2 | 
**Token** | **string** | Authentication token | 
**ExcludeUsersServiceAccount** | Pointer to **bool** |  | [optional] 
**FilterGroup** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelRequest

`func NewModelRequest(name string, authorizationFlow string, externalHost string, acsUrl string, url string, token string, ) *ModelRequest`

NewModelRequest instantiates a new ModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRequestWithDefaults

`func NewModelRequestWithDefaults() *ModelRequest`

NewModelRequestWithDefaults instantiates a new ModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticationFlow

`func (o *ModelRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *ModelRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *ModelRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *ModelRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *ModelRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *ModelRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *ModelRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *ModelRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *ModelRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetPropertyMappings

`func (o *ModelRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *ModelRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *ModelRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *ModelRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetBaseDn

`func (o *ModelRequest) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *ModelRequest) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *ModelRequest) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *ModelRequest) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetSearchGroup

`func (o *ModelRequest) GetSearchGroup() string`

GetSearchGroup returns the SearchGroup field if non-nil, zero value otherwise.

### GetSearchGroupOk

`func (o *ModelRequest) GetSearchGroupOk() (*string, bool)`

GetSearchGroupOk returns a tuple with the SearchGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchGroup

`func (o *ModelRequest) SetSearchGroup(v string)`

SetSearchGroup sets SearchGroup field to given value.

### HasSearchGroup

`func (o *ModelRequest) HasSearchGroup() bool`

HasSearchGroup returns a boolean if a field has been set.

### SetSearchGroupNil

`func (o *ModelRequest) SetSearchGroupNil(b bool)`

 SetSearchGroupNil sets the value for SearchGroup to be an explicit nil

### UnsetSearchGroup
`func (o *ModelRequest) UnsetSearchGroup()`

UnsetSearchGroup ensures that no value is present for SearchGroup, not even an explicit nil
### GetCertificate

`func (o *ModelRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ModelRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ModelRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ModelRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *ModelRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *ModelRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetTlsServerName

`func (o *ModelRequest) GetTlsServerName() string`

GetTlsServerName returns the TlsServerName field if non-nil, zero value otherwise.

### GetTlsServerNameOk

`func (o *ModelRequest) GetTlsServerNameOk() (*string, bool)`

GetTlsServerNameOk returns a tuple with the TlsServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerName

`func (o *ModelRequest) SetTlsServerName(v string)`

SetTlsServerName sets TlsServerName field to given value.

### HasTlsServerName

`func (o *ModelRequest) HasTlsServerName() bool`

HasTlsServerName returns a boolean if a field has been set.

### GetUidStartNumber

`func (o *ModelRequest) GetUidStartNumber() int32`

GetUidStartNumber returns the UidStartNumber field if non-nil, zero value otherwise.

### GetUidStartNumberOk

`func (o *ModelRequest) GetUidStartNumberOk() (*int32, bool)`

GetUidStartNumberOk returns a tuple with the UidStartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidStartNumber

`func (o *ModelRequest) SetUidStartNumber(v int32)`

SetUidStartNumber sets UidStartNumber field to given value.

### HasUidStartNumber

`func (o *ModelRequest) HasUidStartNumber() bool`

HasUidStartNumber returns a boolean if a field has been set.

### GetGidStartNumber

`func (o *ModelRequest) GetGidStartNumber() int32`

GetGidStartNumber returns the GidStartNumber field if non-nil, zero value otherwise.

### GetGidStartNumberOk

`func (o *ModelRequest) GetGidStartNumberOk() (*int32, bool)`

GetGidStartNumberOk returns a tuple with the GidStartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidStartNumber

`func (o *ModelRequest) SetGidStartNumber(v int32)`

SetGidStartNumber sets GidStartNumber field to given value.

### HasGidStartNumber

`func (o *ModelRequest) HasGidStartNumber() bool`

HasGidStartNumber returns a boolean if a field has been set.

### GetSearchMode

`func (o *ModelRequest) GetSearchMode() LDAPAPIAccessMode`

GetSearchMode returns the SearchMode field if non-nil, zero value otherwise.

### GetSearchModeOk

`func (o *ModelRequest) GetSearchModeOk() (*LDAPAPIAccessMode, bool)`

GetSearchModeOk returns a tuple with the SearchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMode

`func (o *ModelRequest) SetSearchMode(v LDAPAPIAccessMode)`

SetSearchMode sets SearchMode field to given value.

### HasSearchMode

`func (o *ModelRequest) HasSearchMode() bool`

HasSearchMode returns a boolean if a field has been set.

### GetBindMode

`func (o *ModelRequest) GetBindMode() LDAPAPIAccessMode`

GetBindMode returns the BindMode field if non-nil, zero value otherwise.

### GetBindModeOk

`func (o *ModelRequest) GetBindModeOk() (*LDAPAPIAccessMode, bool)`

GetBindModeOk returns a tuple with the BindMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMode

`func (o *ModelRequest) SetBindMode(v LDAPAPIAccessMode)`

SetBindMode sets BindMode field to given value.

### HasBindMode

`func (o *ModelRequest) HasBindMode() bool`

HasBindMode returns a boolean if a field has been set.

### GetMfaSupport

`func (o *ModelRequest) GetMfaSupport() bool`

GetMfaSupport returns the MfaSupport field if non-nil, zero value otherwise.

### GetMfaSupportOk

`func (o *ModelRequest) GetMfaSupportOk() (*bool, bool)`

GetMfaSupportOk returns a tuple with the MfaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaSupport

`func (o *ModelRequest) SetMfaSupport(v bool)`

SetMfaSupport sets MfaSupport field to given value.

### HasMfaSupport

`func (o *ModelRequest) HasMfaSupport() bool`

HasMfaSupport returns a boolean if a field has been set.

### GetClientType

`func (o *ModelRequest) GetClientType() ClientTypeEnum`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *ModelRequest) GetClientTypeOk() (*ClientTypeEnum, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *ModelRequest) SetClientType(v ClientTypeEnum)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *ModelRequest) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetClientId

`func (o *ModelRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ModelRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ModelRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ModelRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ModelRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ModelRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ModelRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ModelRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetAccessCodeValidity

`func (o *ModelRequest) GetAccessCodeValidity() string`

GetAccessCodeValidity returns the AccessCodeValidity field if non-nil, zero value otherwise.

### GetAccessCodeValidityOk

`func (o *ModelRequest) GetAccessCodeValidityOk() (*string, bool)`

GetAccessCodeValidityOk returns a tuple with the AccessCodeValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCodeValidity

`func (o *ModelRequest) SetAccessCodeValidity(v string)`

SetAccessCodeValidity sets AccessCodeValidity field to given value.

### HasAccessCodeValidity

`func (o *ModelRequest) HasAccessCodeValidity() bool`

HasAccessCodeValidity returns a boolean if a field has been set.

### GetAccessTokenValidity

`func (o *ModelRequest) GetAccessTokenValidity() string`

GetAccessTokenValidity returns the AccessTokenValidity field if non-nil, zero value otherwise.

### GetAccessTokenValidityOk

`func (o *ModelRequest) GetAccessTokenValidityOk() (*string, bool)`

GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidity

`func (o *ModelRequest) SetAccessTokenValidity(v string)`

SetAccessTokenValidity sets AccessTokenValidity field to given value.

### HasAccessTokenValidity

`func (o *ModelRequest) HasAccessTokenValidity() bool`

HasAccessTokenValidity returns a boolean if a field has been set.

### GetRefreshTokenValidity

`func (o *ModelRequest) GetRefreshTokenValidity() string`

GetRefreshTokenValidity returns the RefreshTokenValidity field if non-nil, zero value otherwise.

### GetRefreshTokenValidityOk

`func (o *ModelRequest) GetRefreshTokenValidityOk() (*string, bool)`

GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValidity

`func (o *ModelRequest) SetRefreshTokenValidity(v string)`

SetRefreshTokenValidity sets RefreshTokenValidity field to given value.

### HasRefreshTokenValidity

`func (o *ModelRequest) HasRefreshTokenValidity() bool`

HasRefreshTokenValidity returns a boolean if a field has been set.

### GetIncludeClaimsInIdToken

`func (o *ModelRequest) GetIncludeClaimsInIdToken() bool`

GetIncludeClaimsInIdToken returns the IncludeClaimsInIdToken field if non-nil, zero value otherwise.

### GetIncludeClaimsInIdTokenOk

`func (o *ModelRequest) GetIncludeClaimsInIdTokenOk() (*bool, bool)`

GetIncludeClaimsInIdTokenOk returns a tuple with the IncludeClaimsInIdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeClaimsInIdToken

`func (o *ModelRequest) SetIncludeClaimsInIdToken(v bool)`

SetIncludeClaimsInIdToken sets IncludeClaimsInIdToken field to given value.

### HasIncludeClaimsInIdToken

`func (o *ModelRequest) HasIncludeClaimsInIdToken() bool`

HasIncludeClaimsInIdToken returns a boolean if a field has been set.

### GetSigningKey

`func (o *ModelRequest) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *ModelRequest) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *ModelRequest) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.

### HasSigningKey

`func (o *ModelRequest) HasSigningKey() bool`

HasSigningKey returns a boolean if a field has been set.

### SetSigningKeyNil

`func (o *ModelRequest) SetSigningKeyNil(b bool)`

 SetSigningKeyNil sets the value for SigningKey to be an explicit nil

### UnsetSigningKey
`func (o *ModelRequest) UnsetSigningKey()`

UnsetSigningKey ensures that no value is present for SigningKey, not even an explicit nil
### GetRedirectUris

`func (o *ModelRequest) GetRedirectUris() string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ModelRequest) GetRedirectUrisOk() (*string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ModelRequest) SetRedirectUris(v string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *ModelRequest) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetSubMode

`func (o *ModelRequest) GetSubMode() SubModeEnum`

GetSubMode returns the SubMode field if non-nil, zero value otherwise.

### GetSubModeOk

`func (o *ModelRequest) GetSubModeOk() (*SubModeEnum, bool)`

GetSubModeOk returns a tuple with the SubMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMode

`func (o *ModelRequest) SetSubMode(v SubModeEnum)`

SetSubMode sets SubMode field to given value.

### HasSubMode

`func (o *ModelRequest) HasSubMode() bool`

HasSubMode returns a boolean if a field has been set.

### GetIssuerMode

`func (o *ModelRequest) GetIssuerMode() IssuerModeEnum`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *ModelRequest) GetIssuerModeOk() (*IssuerModeEnum, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *ModelRequest) SetIssuerMode(v IssuerModeEnum)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *ModelRequest) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetJwksSources

`func (o *ModelRequest) GetJwksSources() []string`

GetJwksSources returns the JwksSources field if non-nil, zero value otherwise.

### GetJwksSourcesOk

`func (o *ModelRequest) GetJwksSourcesOk() (*[]string, bool)`

GetJwksSourcesOk returns a tuple with the JwksSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksSources

`func (o *ModelRequest) SetJwksSources(v []string)`

SetJwksSources sets JwksSources field to given value.

### HasJwksSources

`func (o *ModelRequest) HasJwksSources() bool`

HasJwksSources returns a boolean if a field has been set.

### GetInternalHost

`func (o *ModelRequest) GetInternalHost() string`

GetInternalHost returns the InternalHost field if non-nil, zero value otherwise.

### GetInternalHostOk

`func (o *ModelRequest) GetInternalHostOk() (*string, bool)`

GetInternalHostOk returns a tuple with the InternalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHost

`func (o *ModelRequest) SetInternalHost(v string)`

SetInternalHost sets InternalHost field to given value.

### HasInternalHost

`func (o *ModelRequest) HasInternalHost() bool`

HasInternalHost returns a boolean if a field has been set.

### GetExternalHost

`func (o *ModelRequest) GetExternalHost() string`

GetExternalHost returns the ExternalHost field if non-nil, zero value otherwise.

### GetExternalHostOk

`func (o *ModelRequest) GetExternalHostOk() (*string, bool)`

GetExternalHostOk returns a tuple with the ExternalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHost

`func (o *ModelRequest) SetExternalHost(v string)`

SetExternalHost sets ExternalHost field to given value.


### GetInternalHostSslValidation

`func (o *ModelRequest) GetInternalHostSslValidation() bool`

GetInternalHostSslValidation returns the InternalHostSslValidation field if non-nil, zero value otherwise.

### GetInternalHostSslValidationOk

`func (o *ModelRequest) GetInternalHostSslValidationOk() (*bool, bool)`

GetInternalHostSslValidationOk returns a tuple with the InternalHostSslValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostSslValidation

`func (o *ModelRequest) SetInternalHostSslValidation(v bool)`

SetInternalHostSslValidation sets InternalHostSslValidation field to given value.

### HasInternalHostSslValidation

`func (o *ModelRequest) HasInternalHostSslValidation() bool`

HasInternalHostSslValidation returns a boolean if a field has been set.

### GetSkipPathRegex

`func (o *ModelRequest) GetSkipPathRegex() string`

GetSkipPathRegex returns the SkipPathRegex field if non-nil, zero value otherwise.

### GetSkipPathRegexOk

`func (o *ModelRequest) GetSkipPathRegexOk() (*string, bool)`

GetSkipPathRegexOk returns a tuple with the SkipPathRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPathRegex

`func (o *ModelRequest) SetSkipPathRegex(v string)`

SetSkipPathRegex sets SkipPathRegex field to given value.

### HasSkipPathRegex

`func (o *ModelRequest) HasSkipPathRegex() bool`

HasSkipPathRegex returns a boolean if a field has been set.

### GetBasicAuthEnabled

`func (o *ModelRequest) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *ModelRequest) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *ModelRequest) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *ModelRequest) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetBasicAuthPasswordAttribute

`func (o *ModelRequest) GetBasicAuthPasswordAttribute() string`

GetBasicAuthPasswordAttribute returns the BasicAuthPasswordAttribute field if non-nil, zero value otherwise.

### GetBasicAuthPasswordAttributeOk

`func (o *ModelRequest) GetBasicAuthPasswordAttributeOk() (*string, bool)`

GetBasicAuthPasswordAttributeOk returns a tuple with the BasicAuthPasswordAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPasswordAttribute

`func (o *ModelRequest) SetBasicAuthPasswordAttribute(v string)`

SetBasicAuthPasswordAttribute sets BasicAuthPasswordAttribute field to given value.

### HasBasicAuthPasswordAttribute

`func (o *ModelRequest) HasBasicAuthPasswordAttribute() bool`

HasBasicAuthPasswordAttribute returns a boolean if a field has been set.

### GetBasicAuthUserAttribute

`func (o *ModelRequest) GetBasicAuthUserAttribute() string`

GetBasicAuthUserAttribute returns the BasicAuthUserAttribute field if non-nil, zero value otherwise.

### GetBasicAuthUserAttributeOk

`func (o *ModelRequest) GetBasicAuthUserAttributeOk() (*string, bool)`

GetBasicAuthUserAttributeOk returns a tuple with the BasicAuthUserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUserAttribute

`func (o *ModelRequest) SetBasicAuthUserAttribute(v string)`

SetBasicAuthUserAttribute sets BasicAuthUserAttribute field to given value.

### HasBasicAuthUserAttribute

`func (o *ModelRequest) HasBasicAuthUserAttribute() bool`

HasBasicAuthUserAttribute returns a boolean if a field has been set.

### GetMode

`func (o *ModelRequest) GetMode() ProxyMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ModelRequest) GetModeOk() (*ProxyMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ModelRequest) SetMode(v ProxyMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ModelRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetInterceptHeaderAuth

`func (o *ModelRequest) GetInterceptHeaderAuth() bool`

GetInterceptHeaderAuth returns the InterceptHeaderAuth field if non-nil, zero value otherwise.

### GetInterceptHeaderAuthOk

`func (o *ModelRequest) GetInterceptHeaderAuthOk() (*bool, bool)`

GetInterceptHeaderAuthOk returns a tuple with the InterceptHeaderAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterceptHeaderAuth

`func (o *ModelRequest) SetInterceptHeaderAuth(v bool)`

SetInterceptHeaderAuth sets InterceptHeaderAuth field to given value.

### HasInterceptHeaderAuth

`func (o *ModelRequest) HasInterceptHeaderAuth() bool`

HasInterceptHeaderAuth returns a boolean if a field has been set.

### GetCookieDomain

`func (o *ModelRequest) GetCookieDomain() string`

GetCookieDomain returns the CookieDomain field if non-nil, zero value otherwise.

### GetCookieDomainOk

`func (o *ModelRequest) GetCookieDomainOk() (*string, bool)`

GetCookieDomainOk returns a tuple with the CookieDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieDomain

`func (o *ModelRequest) SetCookieDomain(v string)`

SetCookieDomain sets CookieDomain field to given value.

### HasCookieDomain

`func (o *ModelRequest) HasCookieDomain() bool`

HasCookieDomain returns a boolean if a field has been set.

### GetClientNetworks

`func (o *ModelRequest) GetClientNetworks() string`

GetClientNetworks returns the ClientNetworks field if non-nil, zero value otherwise.

### GetClientNetworksOk

`func (o *ModelRequest) GetClientNetworksOk() (*string, bool)`

GetClientNetworksOk returns a tuple with the ClientNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNetworks

`func (o *ModelRequest) SetClientNetworks(v string)`

SetClientNetworks sets ClientNetworks field to given value.

### HasClientNetworks

`func (o *ModelRequest) HasClientNetworks() bool`

HasClientNetworks returns a boolean if a field has been set.

### GetSharedSecret

`func (o *ModelRequest) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *ModelRequest) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *ModelRequest) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *ModelRequest) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetAcsUrl

`func (o *ModelRequest) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *ModelRequest) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *ModelRequest) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.


### GetAudience

`func (o *ModelRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *ModelRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *ModelRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *ModelRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *ModelRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ModelRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ModelRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ModelRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAssertionValidNotBefore

`func (o *ModelRequest) GetAssertionValidNotBefore() string`

GetAssertionValidNotBefore returns the AssertionValidNotBefore field if non-nil, zero value otherwise.

### GetAssertionValidNotBeforeOk

`func (o *ModelRequest) GetAssertionValidNotBeforeOk() (*string, bool)`

GetAssertionValidNotBeforeOk returns a tuple with the AssertionValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotBefore

`func (o *ModelRequest) SetAssertionValidNotBefore(v string)`

SetAssertionValidNotBefore sets AssertionValidNotBefore field to given value.

### HasAssertionValidNotBefore

`func (o *ModelRequest) HasAssertionValidNotBefore() bool`

HasAssertionValidNotBefore returns a boolean if a field has been set.

### GetAssertionValidNotOnOrAfter

`func (o *ModelRequest) GetAssertionValidNotOnOrAfter() string`

GetAssertionValidNotOnOrAfter returns the AssertionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetAssertionValidNotOnOrAfterOk

`func (o *ModelRequest) GetAssertionValidNotOnOrAfterOk() (*string, bool)`

GetAssertionValidNotOnOrAfterOk returns a tuple with the AssertionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionValidNotOnOrAfter

`func (o *ModelRequest) SetAssertionValidNotOnOrAfter(v string)`

SetAssertionValidNotOnOrAfter sets AssertionValidNotOnOrAfter field to given value.

### HasAssertionValidNotOnOrAfter

`func (o *ModelRequest) HasAssertionValidNotOnOrAfter() bool`

HasAssertionValidNotOnOrAfter returns a boolean if a field has been set.

### GetSessionValidNotOnOrAfter

`func (o *ModelRequest) GetSessionValidNotOnOrAfter() string`

GetSessionValidNotOnOrAfter returns the SessionValidNotOnOrAfter field if non-nil, zero value otherwise.

### GetSessionValidNotOnOrAfterOk

`func (o *ModelRequest) GetSessionValidNotOnOrAfterOk() (*string, bool)`

GetSessionValidNotOnOrAfterOk returns a tuple with the SessionValidNotOnOrAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionValidNotOnOrAfter

`func (o *ModelRequest) SetSessionValidNotOnOrAfter(v string)`

SetSessionValidNotOnOrAfter sets SessionValidNotOnOrAfter field to given value.

### HasSessionValidNotOnOrAfter

`func (o *ModelRequest) HasSessionValidNotOnOrAfter() bool`

HasSessionValidNotOnOrAfter returns a boolean if a field has been set.

### GetNameIdMapping

`func (o *ModelRequest) GetNameIdMapping() string`

GetNameIdMapping returns the NameIdMapping field if non-nil, zero value otherwise.

### GetNameIdMappingOk

`func (o *ModelRequest) GetNameIdMappingOk() (*string, bool)`

GetNameIdMappingOk returns a tuple with the NameIdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdMapping

`func (o *ModelRequest) SetNameIdMapping(v string)`

SetNameIdMapping sets NameIdMapping field to given value.

### HasNameIdMapping

`func (o *ModelRequest) HasNameIdMapping() bool`

HasNameIdMapping returns a boolean if a field has been set.

### SetNameIdMappingNil

`func (o *ModelRequest) SetNameIdMappingNil(b bool)`

 SetNameIdMappingNil sets the value for NameIdMapping to be an explicit nil

### UnsetNameIdMapping
`func (o *ModelRequest) UnsetNameIdMapping()`

UnsetNameIdMapping ensures that no value is present for NameIdMapping, not even an explicit nil
### GetDigestAlgorithm

`func (o *ModelRequest) GetDigestAlgorithm() DigestAlgorithmEnum`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *ModelRequest) GetDigestAlgorithmOk() (*DigestAlgorithmEnum, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *ModelRequest) SetDigestAlgorithm(v DigestAlgorithmEnum)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *ModelRequest) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *ModelRequest) GetSignatureAlgorithm() SignatureAlgorithmEnum`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *ModelRequest) GetSignatureAlgorithmOk() (*SignatureAlgorithmEnum, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *ModelRequest) SetSignatureAlgorithm(v SignatureAlgorithmEnum)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *ModelRequest) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.

### GetSigningKp

`func (o *ModelRequest) GetSigningKp() string`

GetSigningKp returns the SigningKp field if non-nil, zero value otherwise.

### GetSigningKpOk

`func (o *ModelRequest) GetSigningKpOk() (*string, bool)`

GetSigningKpOk returns a tuple with the SigningKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKp

`func (o *ModelRequest) SetSigningKp(v string)`

SetSigningKp sets SigningKp field to given value.

### HasSigningKp

`func (o *ModelRequest) HasSigningKp() bool`

HasSigningKp returns a boolean if a field has been set.

### SetSigningKpNil

`func (o *ModelRequest) SetSigningKpNil(b bool)`

 SetSigningKpNil sets the value for SigningKp to be an explicit nil

### UnsetSigningKp
`func (o *ModelRequest) UnsetSigningKp()`

UnsetSigningKp ensures that no value is present for SigningKp, not even an explicit nil
### GetVerificationKp

`func (o *ModelRequest) GetVerificationKp() string`

GetVerificationKp returns the VerificationKp field if non-nil, zero value otherwise.

### GetVerificationKpOk

`func (o *ModelRequest) GetVerificationKpOk() (*string, bool)`

GetVerificationKpOk returns a tuple with the VerificationKp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationKp

`func (o *ModelRequest) SetVerificationKp(v string)`

SetVerificationKp sets VerificationKp field to given value.

### HasVerificationKp

`func (o *ModelRequest) HasVerificationKp() bool`

HasVerificationKp returns a boolean if a field has been set.

### SetVerificationKpNil

`func (o *ModelRequest) SetVerificationKpNil(b bool)`

 SetVerificationKpNil sets the value for VerificationKp to be an explicit nil

### UnsetVerificationKp
`func (o *ModelRequest) UnsetVerificationKp()`

UnsetVerificationKp ensures that no value is present for VerificationKp, not even an explicit nil
### GetSpBinding

`func (o *ModelRequest) GetSpBinding() SpBindingEnum`

GetSpBinding returns the SpBinding field if non-nil, zero value otherwise.

### GetSpBindingOk

`func (o *ModelRequest) GetSpBindingOk() (*SpBindingEnum, bool)`

GetSpBindingOk returns a tuple with the SpBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpBinding

`func (o *ModelRequest) SetSpBinding(v SpBindingEnum)`

SetSpBinding sets SpBinding field to given value.

### HasSpBinding

`func (o *ModelRequest) HasSpBinding() bool`

HasSpBinding returns a boolean if a field has been set.

### GetPropertyMappingsGroup

`func (o *ModelRequest) GetPropertyMappingsGroup() []string`

GetPropertyMappingsGroup returns the PropertyMappingsGroup field if non-nil, zero value otherwise.

### GetPropertyMappingsGroupOk

`func (o *ModelRequest) GetPropertyMappingsGroupOk() (*[]string, bool)`

GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappingsGroup

`func (o *ModelRequest) SetPropertyMappingsGroup(v []string)`

SetPropertyMappingsGroup sets PropertyMappingsGroup field to given value.

### HasPropertyMappingsGroup

`func (o *ModelRequest) HasPropertyMappingsGroup() bool`

HasPropertyMappingsGroup returns a boolean if a field has been set.

### GetUrl

`func (o *ModelRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetToken

`func (o *ModelRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ModelRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ModelRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetExcludeUsersServiceAccount

`func (o *ModelRequest) GetExcludeUsersServiceAccount() bool`

GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field if non-nil, zero value otherwise.

### GetExcludeUsersServiceAccountOk

`func (o *ModelRequest) GetExcludeUsersServiceAccountOk() (*bool, bool)`

GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsersServiceAccount

`func (o *ModelRequest) SetExcludeUsersServiceAccount(v bool)`

SetExcludeUsersServiceAccount sets ExcludeUsersServiceAccount field to given value.

### HasExcludeUsersServiceAccount

`func (o *ModelRequest) HasExcludeUsersServiceAccount() bool`

HasExcludeUsersServiceAccount returns a boolean if a field has been set.

### GetFilterGroup

`func (o *ModelRequest) GetFilterGroup() string`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *ModelRequest) GetFilterGroupOk() (*string, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *ModelRequest) SetFilterGroup(v string)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *ModelRequest) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### SetFilterGroupNil

`func (o *ModelRequest) SetFilterGroupNil(b bool)`

 SetFilterGroupNil sets the value for FilterGroup to be an explicit nil

### UnsetFilterGroup
`func (o *ModelRequest) UnsetFilterGroup()`

UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


