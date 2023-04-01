# ProxyOutpostConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**InternalHost** | Pointer to **string** |  | [optional] 
**ExternalHost** | **string** |  | 
**InternalHostSslValidation** | Pointer to **bool** | Validate SSL Certificates of upstream servers | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**OidcConfiguration** | [**OpenIDConnectConfiguration**](OpenIDConnectConfiguration.md) |  | [readonly] 
**CookieSecret** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**SkipPathRegex** | Pointer to **string** | Regular expressions for which authentication is not required. Each new line is interpreted as a new Regular Expression. | [optional] 
**BasicAuthEnabled** | Pointer to **bool** | Set a custom HTTP-Basic Authentication header based on values from authentik. | [optional] 
**BasicAuthPasswordAttribute** | Pointer to **string** | User/Group Attribute used for the password part of the HTTP-Basic Header. | [optional] 
**BasicAuthUserAttribute** | Pointer to **string** | User/Group Attribute used for the user part of the HTTP-Basic Header. If not set, the user&#39;s Email address is used. | [optional] 
**Mode** | Pointer to [**ProxyMode**](ProxyMode.md) |  | [optional] 
**CookieDomain** | Pointer to **string** |  | [optional] 
**AccessTokenValidity** | **NullableFloat64** | Get token validity as second count | [readonly] 
**InterceptHeaderAuth** | Pointer to **bool** | When enabled, this provider will intercept the authorization header and authenticate requests based on its value. | [optional] 
**ScopesToRequest** | **[]string** | Get all the scope names the outpost should request, including custom-defined ones | [readonly] 
**AssignedApplicationSlug** | **string** | Internal application name, used in URLs. | [readonly] 
**AssignedApplicationName** | **string** | Application&#39;s display Name. | [readonly] 

## Methods

### NewProxyOutpostConfig

`func NewProxyOutpostConfig(pk int32, name string, externalHost string, oidcConfiguration OpenIDConnectConfiguration, accessTokenValidity NullableFloat64, scopesToRequest []string, assignedApplicationSlug string, assignedApplicationName string, ) *ProxyOutpostConfig`

NewProxyOutpostConfig instantiates a new ProxyOutpostConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyOutpostConfigWithDefaults

`func NewProxyOutpostConfigWithDefaults() *ProxyOutpostConfig`

NewProxyOutpostConfigWithDefaults instantiates a new ProxyOutpostConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *ProxyOutpostConfig) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *ProxyOutpostConfig) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *ProxyOutpostConfig) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *ProxyOutpostConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProxyOutpostConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProxyOutpostConfig) SetName(v string)`

SetName sets Name field to given value.


### GetInternalHost

`func (o *ProxyOutpostConfig) GetInternalHost() string`

GetInternalHost returns the InternalHost field if non-nil, zero value otherwise.

### GetInternalHostOk

`func (o *ProxyOutpostConfig) GetInternalHostOk() (*string, bool)`

GetInternalHostOk returns a tuple with the InternalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHost

`func (o *ProxyOutpostConfig) SetInternalHost(v string)`

SetInternalHost sets InternalHost field to given value.

### HasInternalHost

`func (o *ProxyOutpostConfig) HasInternalHost() bool`

HasInternalHost returns a boolean if a field has been set.

### GetExternalHost

`func (o *ProxyOutpostConfig) GetExternalHost() string`

GetExternalHost returns the ExternalHost field if non-nil, zero value otherwise.

### GetExternalHostOk

`func (o *ProxyOutpostConfig) GetExternalHostOk() (*string, bool)`

GetExternalHostOk returns a tuple with the ExternalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHost

`func (o *ProxyOutpostConfig) SetExternalHost(v string)`

SetExternalHost sets ExternalHost field to given value.


### GetInternalHostSslValidation

`func (o *ProxyOutpostConfig) GetInternalHostSslValidation() bool`

GetInternalHostSslValidation returns the InternalHostSslValidation field if non-nil, zero value otherwise.

### GetInternalHostSslValidationOk

`func (o *ProxyOutpostConfig) GetInternalHostSslValidationOk() (*bool, bool)`

GetInternalHostSslValidationOk returns a tuple with the InternalHostSslValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostSslValidation

`func (o *ProxyOutpostConfig) SetInternalHostSslValidation(v bool)`

SetInternalHostSslValidation sets InternalHostSslValidation field to given value.

### HasInternalHostSslValidation

`func (o *ProxyOutpostConfig) HasInternalHostSslValidation() bool`

HasInternalHostSslValidation returns a boolean if a field has been set.

### GetClientId

`func (o *ProxyOutpostConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ProxyOutpostConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ProxyOutpostConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ProxyOutpostConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ProxyOutpostConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ProxyOutpostConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ProxyOutpostConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ProxyOutpostConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetOidcConfiguration

`func (o *ProxyOutpostConfig) GetOidcConfiguration() OpenIDConnectConfiguration`

GetOidcConfiguration returns the OidcConfiguration field if non-nil, zero value otherwise.

### GetOidcConfigurationOk

`func (o *ProxyOutpostConfig) GetOidcConfigurationOk() (*OpenIDConnectConfiguration, bool)`

GetOidcConfigurationOk returns a tuple with the OidcConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcConfiguration

`func (o *ProxyOutpostConfig) SetOidcConfiguration(v OpenIDConnectConfiguration)`

SetOidcConfiguration sets OidcConfiguration field to given value.


### GetCookieSecret

`func (o *ProxyOutpostConfig) GetCookieSecret() string`

GetCookieSecret returns the CookieSecret field if non-nil, zero value otherwise.

### GetCookieSecretOk

`func (o *ProxyOutpostConfig) GetCookieSecretOk() (*string, bool)`

GetCookieSecretOk returns a tuple with the CookieSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieSecret

`func (o *ProxyOutpostConfig) SetCookieSecret(v string)`

SetCookieSecret sets CookieSecret field to given value.

### HasCookieSecret

`func (o *ProxyOutpostConfig) HasCookieSecret() bool`

HasCookieSecret returns a boolean if a field has been set.

### GetCertificate

`func (o *ProxyOutpostConfig) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ProxyOutpostConfig) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ProxyOutpostConfig) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ProxyOutpostConfig) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *ProxyOutpostConfig) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *ProxyOutpostConfig) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetSkipPathRegex

`func (o *ProxyOutpostConfig) GetSkipPathRegex() string`

GetSkipPathRegex returns the SkipPathRegex field if non-nil, zero value otherwise.

### GetSkipPathRegexOk

`func (o *ProxyOutpostConfig) GetSkipPathRegexOk() (*string, bool)`

GetSkipPathRegexOk returns a tuple with the SkipPathRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPathRegex

`func (o *ProxyOutpostConfig) SetSkipPathRegex(v string)`

SetSkipPathRegex sets SkipPathRegex field to given value.

### HasSkipPathRegex

`func (o *ProxyOutpostConfig) HasSkipPathRegex() bool`

HasSkipPathRegex returns a boolean if a field has been set.

### GetBasicAuthEnabled

`func (o *ProxyOutpostConfig) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *ProxyOutpostConfig) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *ProxyOutpostConfig) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *ProxyOutpostConfig) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetBasicAuthPasswordAttribute

`func (o *ProxyOutpostConfig) GetBasicAuthPasswordAttribute() string`

GetBasicAuthPasswordAttribute returns the BasicAuthPasswordAttribute field if non-nil, zero value otherwise.

### GetBasicAuthPasswordAttributeOk

`func (o *ProxyOutpostConfig) GetBasicAuthPasswordAttributeOk() (*string, bool)`

GetBasicAuthPasswordAttributeOk returns a tuple with the BasicAuthPasswordAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPasswordAttribute

`func (o *ProxyOutpostConfig) SetBasicAuthPasswordAttribute(v string)`

SetBasicAuthPasswordAttribute sets BasicAuthPasswordAttribute field to given value.

### HasBasicAuthPasswordAttribute

`func (o *ProxyOutpostConfig) HasBasicAuthPasswordAttribute() bool`

HasBasicAuthPasswordAttribute returns a boolean if a field has been set.

### GetBasicAuthUserAttribute

`func (o *ProxyOutpostConfig) GetBasicAuthUserAttribute() string`

GetBasicAuthUserAttribute returns the BasicAuthUserAttribute field if non-nil, zero value otherwise.

### GetBasicAuthUserAttributeOk

`func (o *ProxyOutpostConfig) GetBasicAuthUserAttributeOk() (*string, bool)`

GetBasicAuthUserAttributeOk returns a tuple with the BasicAuthUserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUserAttribute

`func (o *ProxyOutpostConfig) SetBasicAuthUserAttribute(v string)`

SetBasicAuthUserAttribute sets BasicAuthUserAttribute field to given value.

### HasBasicAuthUserAttribute

`func (o *ProxyOutpostConfig) HasBasicAuthUserAttribute() bool`

HasBasicAuthUserAttribute returns a boolean if a field has been set.

### GetMode

`func (o *ProxyOutpostConfig) GetMode() ProxyMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProxyOutpostConfig) GetModeOk() (*ProxyMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProxyOutpostConfig) SetMode(v ProxyMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ProxyOutpostConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCookieDomain

`func (o *ProxyOutpostConfig) GetCookieDomain() string`

GetCookieDomain returns the CookieDomain field if non-nil, zero value otherwise.

### GetCookieDomainOk

`func (o *ProxyOutpostConfig) GetCookieDomainOk() (*string, bool)`

GetCookieDomainOk returns a tuple with the CookieDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieDomain

`func (o *ProxyOutpostConfig) SetCookieDomain(v string)`

SetCookieDomain sets CookieDomain field to given value.

### HasCookieDomain

`func (o *ProxyOutpostConfig) HasCookieDomain() bool`

HasCookieDomain returns a boolean if a field has been set.

### GetAccessTokenValidity

`func (o *ProxyOutpostConfig) GetAccessTokenValidity() float64`

GetAccessTokenValidity returns the AccessTokenValidity field if non-nil, zero value otherwise.

### GetAccessTokenValidityOk

`func (o *ProxyOutpostConfig) GetAccessTokenValidityOk() (*float64, bool)`

GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidity

`func (o *ProxyOutpostConfig) SetAccessTokenValidity(v float64)`

SetAccessTokenValidity sets AccessTokenValidity field to given value.


### SetAccessTokenValidityNil

`func (o *ProxyOutpostConfig) SetAccessTokenValidityNil(b bool)`

 SetAccessTokenValidityNil sets the value for AccessTokenValidity to be an explicit nil

### UnsetAccessTokenValidity
`func (o *ProxyOutpostConfig) UnsetAccessTokenValidity()`

UnsetAccessTokenValidity ensures that no value is present for AccessTokenValidity, not even an explicit nil
### GetInterceptHeaderAuth

`func (o *ProxyOutpostConfig) GetInterceptHeaderAuth() bool`

GetInterceptHeaderAuth returns the InterceptHeaderAuth field if non-nil, zero value otherwise.

### GetInterceptHeaderAuthOk

`func (o *ProxyOutpostConfig) GetInterceptHeaderAuthOk() (*bool, bool)`

GetInterceptHeaderAuthOk returns a tuple with the InterceptHeaderAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterceptHeaderAuth

`func (o *ProxyOutpostConfig) SetInterceptHeaderAuth(v bool)`

SetInterceptHeaderAuth sets InterceptHeaderAuth field to given value.

### HasInterceptHeaderAuth

`func (o *ProxyOutpostConfig) HasInterceptHeaderAuth() bool`

HasInterceptHeaderAuth returns a boolean if a field has been set.

### GetScopesToRequest

`func (o *ProxyOutpostConfig) GetScopesToRequest() []string`

GetScopesToRequest returns the ScopesToRequest field if non-nil, zero value otherwise.

### GetScopesToRequestOk

`func (o *ProxyOutpostConfig) GetScopesToRequestOk() (*[]string, bool)`

GetScopesToRequestOk returns a tuple with the ScopesToRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesToRequest

`func (o *ProxyOutpostConfig) SetScopesToRequest(v []string)`

SetScopesToRequest sets ScopesToRequest field to given value.


### GetAssignedApplicationSlug

`func (o *ProxyOutpostConfig) GetAssignedApplicationSlug() string`

GetAssignedApplicationSlug returns the AssignedApplicationSlug field if non-nil, zero value otherwise.

### GetAssignedApplicationSlugOk

`func (o *ProxyOutpostConfig) GetAssignedApplicationSlugOk() (*string, bool)`

GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedApplicationSlug

`func (o *ProxyOutpostConfig) SetAssignedApplicationSlug(v string)`

SetAssignedApplicationSlug sets AssignedApplicationSlug field to given value.


### GetAssignedApplicationName

`func (o *ProxyOutpostConfig) GetAssignedApplicationName() string`

GetAssignedApplicationName returns the AssignedApplicationName field if non-nil, zero value otherwise.

### GetAssignedApplicationNameOk

`func (o *ProxyOutpostConfig) GetAssignedApplicationNameOk() (*string, bool)`

GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedApplicationName

`func (o *ProxyOutpostConfig) SetAssignedApplicationName(v string)`

SetAssignedApplicationName sets AssignedApplicationName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


