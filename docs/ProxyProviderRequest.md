# ProxyProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**InternalHost** | Pointer to **string** |  | [optional] 
**ExternalHost** | **string** |  | 
**InternalHostSslValidation** | Pointer to **bool** | Validate SSL Certificates of upstream servers | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**SkipPathRegex** | Pointer to **string** | Regular expressions for which authentication is not required. Each new line is interpreted as a new Regular Expression. | [optional] 
**BasicAuthEnabled** | Pointer to **bool** | Set a custom HTTP-Basic Authentication header based on values from authentik. | [optional] 
**BasicAuthPasswordAttribute** | Pointer to **string** | User/Group Attribute used for the password part of the HTTP-Basic Header. | [optional] 
**BasicAuthUserAttribute** | Pointer to **string** | User/Group Attribute used for the user part of the HTTP-Basic Header. If not set, the user&#39;s Email address is used. | [optional] 
**Mode** | Pointer to [**ProxyMode**](ProxyMode.md) |  | [optional] 
**InterceptHeaderAuth** | Pointer to **bool** | When enabled, this provider will intercept the authorization header and authenticate requests based on its value. | [optional] 
**CookieDomain** | Pointer to **string** |  | [optional] 
**JwksSources** | Pointer to **[]string** |  | [optional] 
**AccessTokenValidity** | Pointer to **string** | Tokens not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 
**RefreshTokenValidity** | Pointer to **string** | Tokens not valid on or after current time + this value (Format: hours&#x3D;1;minutes&#x3D;2;seconds&#x3D;3). | [optional] 

## Methods

### NewProxyProviderRequest

`func NewProxyProviderRequest(name string, authorizationFlow string, externalHost string, ) *ProxyProviderRequest`

NewProxyProviderRequest instantiates a new ProxyProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyProviderRequestWithDefaults

`func NewProxyProviderRequestWithDefaults() *ProxyProviderRequest`

NewProxyProviderRequestWithDefaults instantiates a new ProxyProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProxyProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProxyProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProxyProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticationFlow

`func (o *ProxyProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *ProxyProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *ProxyProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *ProxyProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *ProxyProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *ProxyProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *ProxyProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *ProxyProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *ProxyProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetPropertyMappings

`func (o *ProxyProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *ProxyProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *ProxyProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *ProxyProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetInternalHost

`func (o *ProxyProviderRequest) GetInternalHost() string`

GetInternalHost returns the InternalHost field if non-nil, zero value otherwise.

### GetInternalHostOk

`func (o *ProxyProviderRequest) GetInternalHostOk() (*string, bool)`

GetInternalHostOk returns a tuple with the InternalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHost

`func (o *ProxyProviderRequest) SetInternalHost(v string)`

SetInternalHost sets InternalHost field to given value.

### HasInternalHost

`func (o *ProxyProviderRequest) HasInternalHost() bool`

HasInternalHost returns a boolean if a field has been set.

### GetExternalHost

`func (o *ProxyProviderRequest) GetExternalHost() string`

GetExternalHost returns the ExternalHost field if non-nil, zero value otherwise.

### GetExternalHostOk

`func (o *ProxyProviderRequest) GetExternalHostOk() (*string, bool)`

GetExternalHostOk returns a tuple with the ExternalHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalHost

`func (o *ProxyProviderRequest) SetExternalHost(v string)`

SetExternalHost sets ExternalHost field to given value.


### GetInternalHostSslValidation

`func (o *ProxyProviderRequest) GetInternalHostSslValidation() bool`

GetInternalHostSslValidation returns the InternalHostSslValidation field if non-nil, zero value otherwise.

### GetInternalHostSslValidationOk

`func (o *ProxyProviderRequest) GetInternalHostSslValidationOk() (*bool, bool)`

GetInternalHostSslValidationOk returns a tuple with the InternalHostSslValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHostSslValidation

`func (o *ProxyProviderRequest) SetInternalHostSslValidation(v bool)`

SetInternalHostSslValidation sets InternalHostSslValidation field to given value.

### HasInternalHostSslValidation

`func (o *ProxyProviderRequest) HasInternalHostSslValidation() bool`

HasInternalHostSslValidation returns a boolean if a field has been set.

### GetCertificate

`func (o *ProxyProviderRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ProxyProviderRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ProxyProviderRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ProxyProviderRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *ProxyProviderRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *ProxyProviderRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetSkipPathRegex

`func (o *ProxyProviderRequest) GetSkipPathRegex() string`

GetSkipPathRegex returns the SkipPathRegex field if non-nil, zero value otherwise.

### GetSkipPathRegexOk

`func (o *ProxyProviderRequest) GetSkipPathRegexOk() (*string, bool)`

GetSkipPathRegexOk returns a tuple with the SkipPathRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPathRegex

`func (o *ProxyProviderRequest) SetSkipPathRegex(v string)`

SetSkipPathRegex sets SkipPathRegex field to given value.

### HasSkipPathRegex

`func (o *ProxyProviderRequest) HasSkipPathRegex() bool`

HasSkipPathRegex returns a boolean if a field has been set.

### GetBasicAuthEnabled

`func (o *ProxyProviderRequest) GetBasicAuthEnabled() bool`

GetBasicAuthEnabled returns the BasicAuthEnabled field if non-nil, zero value otherwise.

### GetBasicAuthEnabledOk

`func (o *ProxyProviderRequest) GetBasicAuthEnabledOk() (*bool, bool)`

GetBasicAuthEnabledOk returns a tuple with the BasicAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthEnabled

`func (o *ProxyProviderRequest) SetBasicAuthEnabled(v bool)`

SetBasicAuthEnabled sets BasicAuthEnabled field to given value.

### HasBasicAuthEnabled

`func (o *ProxyProviderRequest) HasBasicAuthEnabled() bool`

HasBasicAuthEnabled returns a boolean if a field has been set.

### GetBasicAuthPasswordAttribute

`func (o *ProxyProviderRequest) GetBasicAuthPasswordAttribute() string`

GetBasicAuthPasswordAttribute returns the BasicAuthPasswordAttribute field if non-nil, zero value otherwise.

### GetBasicAuthPasswordAttributeOk

`func (o *ProxyProviderRequest) GetBasicAuthPasswordAttributeOk() (*string, bool)`

GetBasicAuthPasswordAttributeOk returns a tuple with the BasicAuthPasswordAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPasswordAttribute

`func (o *ProxyProviderRequest) SetBasicAuthPasswordAttribute(v string)`

SetBasicAuthPasswordAttribute sets BasicAuthPasswordAttribute field to given value.

### HasBasicAuthPasswordAttribute

`func (o *ProxyProviderRequest) HasBasicAuthPasswordAttribute() bool`

HasBasicAuthPasswordAttribute returns a boolean if a field has been set.

### GetBasicAuthUserAttribute

`func (o *ProxyProviderRequest) GetBasicAuthUserAttribute() string`

GetBasicAuthUserAttribute returns the BasicAuthUserAttribute field if non-nil, zero value otherwise.

### GetBasicAuthUserAttributeOk

`func (o *ProxyProviderRequest) GetBasicAuthUserAttributeOk() (*string, bool)`

GetBasicAuthUserAttributeOk returns a tuple with the BasicAuthUserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUserAttribute

`func (o *ProxyProviderRequest) SetBasicAuthUserAttribute(v string)`

SetBasicAuthUserAttribute sets BasicAuthUserAttribute field to given value.

### HasBasicAuthUserAttribute

`func (o *ProxyProviderRequest) HasBasicAuthUserAttribute() bool`

HasBasicAuthUserAttribute returns a boolean if a field has been set.

### GetMode

`func (o *ProxyProviderRequest) GetMode() ProxyMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProxyProviderRequest) GetModeOk() (*ProxyMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProxyProviderRequest) SetMode(v ProxyMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ProxyProviderRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetInterceptHeaderAuth

`func (o *ProxyProviderRequest) GetInterceptHeaderAuth() bool`

GetInterceptHeaderAuth returns the InterceptHeaderAuth field if non-nil, zero value otherwise.

### GetInterceptHeaderAuthOk

`func (o *ProxyProviderRequest) GetInterceptHeaderAuthOk() (*bool, bool)`

GetInterceptHeaderAuthOk returns a tuple with the InterceptHeaderAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterceptHeaderAuth

`func (o *ProxyProviderRequest) SetInterceptHeaderAuth(v bool)`

SetInterceptHeaderAuth sets InterceptHeaderAuth field to given value.

### HasInterceptHeaderAuth

`func (o *ProxyProviderRequest) HasInterceptHeaderAuth() bool`

HasInterceptHeaderAuth returns a boolean if a field has been set.

### GetCookieDomain

`func (o *ProxyProviderRequest) GetCookieDomain() string`

GetCookieDomain returns the CookieDomain field if non-nil, zero value otherwise.

### GetCookieDomainOk

`func (o *ProxyProviderRequest) GetCookieDomainOk() (*string, bool)`

GetCookieDomainOk returns a tuple with the CookieDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieDomain

`func (o *ProxyProviderRequest) SetCookieDomain(v string)`

SetCookieDomain sets CookieDomain field to given value.

### HasCookieDomain

`func (o *ProxyProviderRequest) HasCookieDomain() bool`

HasCookieDomain returns a boolean if a field has been set.

### GetJwksSources

`func (o *ProxyProviderRequest) GetJwksSources() []string`

GetJwksSources returns the JwksSources field if non-nil, zero value otherwise.

### GetJwksSourcesOk

`func (o *ProxyProviderRequest) GetJwksSourcesOk() (*[]string, bool)`

GetJwksSourcesOk returns a tuple with the JwksSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksSources

`func (o *ProxyProviderRequest) SetJwksSources(v []string)`

SetJwksSources sets JwksSources field to given value.

### HasJwksSources

`func (o *ProxyProviderRequest) HasJwksSources() bool`

HasJwksSources returns a boolean if a field has been set.

### GetAccessTokenValidity

`func (o *ProxyProviderRequest) GetAccessTokenValidity() string`

GetAccessTokenValidity returns the AccessTokenValidity field if non-nil, zero value otherwise.

### GetAccessTokenValidityOk

`func (o *ProxyProviderRequest) GetAccessTokenValidityOk() (*string, bool)`

GetAccessTokenValidityOk returns a tuple with the AccessTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidity

`func (o *ProxyProviderRequest) SetAccessTokenValidity(v string)`

SetAccessTokenValidity sets AccessTokenValidity field to given value.

### HasAccessTokenValidity

`func (o *ProxyProviderRequest) HasAccessTokenValidity() bool`

HasAccessTokenValidity returns a boolean if a field has been set.

### GetRefreshTokenValidity

`func (o *ProxyProviderRequest) GetRefreshTokenValidity() string`

GetRefreshTokenValidity returns the RefreshTokenValidity field if non-nil, zero value otherwise.

### GetRefreshTokenValidityOk

`func (o *ProxyProviderRequest) GetRefreshTokenValidityOk() (*string, bool)`

GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValidity

`func (o *ProxyProviderRequest) SetRefreshTokenValidity(v string)`

SetRefreshTokenValidity sets RefreshTokenValidity field to given value.

### HasRefreshTokenValidity

`func (o *ProxyProviderRequest) HasRefreshTokenValidity() bool`

HasRefreshTokenValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


