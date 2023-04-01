# OAuthSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Source&#39;s display Name. | 
**Slug** | **string** | Internal source name, used in URLs. | 
**Enabled** | Pointer to **bool** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow to use when authenticating existing users. | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Flow to use when enrolling new users. | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**UserMatchingMode** | Pointer to [**UserMatchingModeEnum**](UserMatchingModeEnum.md) |  | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 
**ProviderType** | [**ProviderTypeEnum**](ProviderTypeEnum.md) |  | 
**RequestTokenUrl** | Pointer to **NullableString** | URL used to request the initial token. This URL is only required for OAuth 1. | [optional] 
**AuthorizationUrl** | Pointer to **NullableString** | URL the user is redirect to to conest the flow. | [optional] 
**AccessTokenUrl** | Pointer to **NullableString** | URL used by authentik to retrieve tokens. | [optional] 
**ProfileUrl** | Pointer to **NullableString** | URL used by authentik to get user information. | [optional] 
**ConsumerKey** | **string** |  | 
**ConsumerSecret** | **string** |  | 
**AdditionalScopes** | Pointer to **string** |  | [optional] 
**OidcWellKnownUrl** | Pointer to **string** |  | [optional] 
**OidcJwksUrl** | Pointer to **string** |  | [optional] 
**OidcJwks** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewOAuthSourceRequest

`func NewOAuthSourceRequest(name string, slug string, providerType ProviderTypeEnum, consumerKey string, consumerSecret string, ) *OAuthSourceRequest`

NewOAuthSourceRequest instantiates a new OAuthSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthSourceRequestWithDefaults

`func NewOAuthSourceRequestWithDefaults() *OAuthSourceRequest`

NewOAuthSourceRequestWithDefaults instantiates a new OAuthSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OAuthSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuthSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuthSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *OAuthSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OAuthSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OAuthSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *OAuthSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OAuthSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OAuthSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OAuthSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *OAuthSourceRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *OAuthSourceRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *OAuthSourceRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *OAuthSourceRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *OAuthSourceRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *OAuthSourceRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *OAuthSourceRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *OAuthSourceRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *OAuthSourceRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *OAuthSourceRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *OAuthSourceRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *OAuthSourceRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetPolicyEngineMode

`func (o *OAuthSourceRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *OAuthSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *OAuthSourceRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *OAuthSourceRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *OAuthSourceRequest) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *OAuthSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *OAuthSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *OAuthSourceRequest) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *OAuthSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *OAuthSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *OAuthSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *OAuthSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetProviderType

`func (o *OAuthSourceRequest) GetProviderType() ProviderTypeEnum`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *OAuthSourceRequest) GetProviderTypeOk() (*ProviderTypeEnum, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *OAuthSourceRequest) SetProviderType(v ProviderTypeEnum)`

SetProviderType sets ProviderType field to given value.


### GetRequestTokenUrl

`func (o *OAuthSourceRequest) GetRequestTokenUrl() string`

GetRequestTokenUrl returns the RequestTokenUrl field if non-nil, zero value otherwise.

### GetRequestTokenUrlOk

`func (o *OAuthSourceRequest) GetRequestTokenUrlOk() (*string, bool)`

GetRequestTokenUrlOk returns a tuple with the RequestTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTokenUrl

`func (o *OAuthSourceRequest) SetRequestTokenUrl(v string)`

SetRequestTokenUrl sets RequestTokenUrl field to given value.

### HasRequestTokenUrl

`func (o *OAuthSourceRequest) HasRequestTokenUrl() bool`

HasRequestTokenUrl returns a boolean if a field has been set.

### SetRequestTokenUrlNil

`func (o *OAuthSourceRequest) SetRequestTokenUrlNil(b bool)`

 SetRequestTokenUrlNil sets the value for RequestTokenUrl to be an explicit nil

### UnsetRequestTokenUrl
`func (o *OAuthSourceRequest) UnsetRequestTokenUrl()`

UnsetRequestTokenUrl ensures that no value is present for RequestTokenUrl, not even an explicit nil
### GetAuthorizationUrl

`func (o *OAuthSourceRequest) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *OAuthSourceRequest) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *OAuthSourceRequest) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.

### HasAuthorizationUrl

`func (o *OAuthSourceRequest) HasAuthorizationUrl() bool`

HasAuthorizationUrl returns a boolean if a field has been set.

### SetAuthorizationUrlNil

`func (o *OAuthSourceRequest) SetAuthorizationUrlNil(b bool)`

 SetAuthorizationUrlNil sets the value for AuthorizationUrl to be an explicit nil

### UnsetAuthorizationUrl
`func (o *OAuthSourceRequest) UnsetAuthorizationUrl()`

UnsetAuthorizationUrl ensures that no value is present for AuthorizationUrl, not even an explicit nil
### GetAccessTokenUrl

`func (o *OAuthSourceRequest) GetAccessTokenUrl() string`

GetAccessTokenUrl returns the AccessTokenUrl field if non-nil, zero value otherwise.

### GetAccessTokenUrlOk

`func (o *OAuthSourceRequest) GetAccessTokenUrlOk() (*string, bool)`

GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenUrl

`func (o *OAuthSourceRequest) SetAccessTokenUrl(v string)`

SetAccessTokenUrl sets AccessTokenUrl field to given value.

### HasAccessTokenUrl

`func (o *OAuthSourceRequest) HasAccessTokenUrl() bool`

HasAccessTokenUrl returns a boolean if a field has been set.

### SetAccessTokenUrlNil

`func (o *OAuthSourceRequest) SetAccessTokenUrlNil(b bool)`

 SetAccessTokenUrlNil sets the value for AccessTokenUrl to be an explicit nil

### UnsetAccessTokenUrl
`func (o *OAuthSourceRequest) UnsetAccessTokenUrl()`

UnsetAccessTokenUrl ensures that no value is present for AccessTokenUrl, not even an explicit nil
### GetProfileUrl

`func (o *OAuthSourceRequest) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *OAuthSourceRequest) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *OAuthSourceRequest) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *OAuthSourceRequest) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### SetProfileUrlNil

`func (o *OAuthSourceRequest) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *OAuthSourceRequest) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
### GetConsumerKey

`func (o *OAuthSourceRequest) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *OAuthSourceRequest) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *OAuthSourceRequest) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.


### GetConsumerSecret

`func (o *OAuthSourceRequest) GetConsumerSecret() string`

GetConsumerSecret returns the ConsumerSecret field if non-nil, zero value otherwise.

### GetConsumerSecretOk

`func (o *OAuthSourceRequest) GetConsumerSecretOk() (*string, bool)`

GetConsumerSecretOk returns a tuple with the ConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSecret

`func (o *OAuthSourceRequest) SetConsumerSecret(v string)`

SetConsumerSecret sets ConsumerSecret field to given value.


### GetAdditionalScopes

`func (o *OAuthSourceRequest) GetAdditionalScopes() string`

GetAdditionalScopes returns the AdditionalScopes field if non-nil, zero value otherwise.

### GetAdditionalScopesOk

`func (o *OAuthSourceRequest) GetAdditionalScopesOk() (*string, bool)`

GetAdditionalScopesOk returns a tuple with the AdditionalScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalScopes

`func (o *OAuthSourceRequest) SetAdditionalScopes(v string)`

SetAdditionalScopes sets AdditionalScopes field to given value.

### HasAdditionalScopes

`func (o *OAuthSourceRequest) HasAdditionalScopes() bool`

HasAdditionalScopes returns a boolean if a field has been set.

### GetOidcWellKnownUrl

`func (o *OAuthSourceRequest) GetOidcWellKnownUrl() string`

GetOidcWellKnownUrl returns the OidcWellKnownUrl field if non-nil, zero value otherwise.

### GetOidcWellKnownUrlOk

`func (o *OAuthSourceRequest) GetOidcWellKnownUrlOk() (*string, bool)`

GetOidcWellKnownUrlOk returns a tuple with the OidcWellKnownUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcWellKnownUrl

`func (o *OAuthSourceRequest) SetOidcWellKnownUrl(v string)`

SetOidcWellKnownUrl sets OidcWellKnownUrl field to given value.

### HasOidcWellKnownUrl

`func (o *OAuthSourceRequest) HasOidcWellKnownUrl() bool`

HasOidcWellKnownUrl returns a boolean if a field has been set.

### GetOidcJwksUrl

`func (o *OAuthSourceRequest) GetOidcJwksUrl() string`

GetOidcJwksUrl returns the OidcJwksUrl field if non-nil, zero value otherwise.

### GetOidcJwksUrlOk

`func (o *OAuthSourceRequest) GetOidcJwksUrlOk() (*string, bool)`

GetOidcJwksUrlOk returns a tuple with the OidcJwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcJwksUrl

`func (o *OAuthSourceRequest) SetOidcJwksUrl(v string)`

SetOidcJwksUrl sets OidcJwksUrl field to given value.

### HasOidcJwksUrl

`func (o *OAuthSourceRequest) HasOidcJwksUrl() bool`

HasOidcJwksUrl returns a boolean if a field has been set.

### GetOidcJwks

`func (o *OAuthSourceRequest) GetOidcJwks() map[string]interface{}`

GetOidcJwks returns the OidcJwks field if non-nil, zero value otherwise.

### GetOidcJwksOk

`func (o *OAuthSourceRequest) GetOidcJwksOk() (*map[string]interface{}, bool)`

GetOidcJwksOk returns a tuple with the OidcJwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcJwks

`func (o *OAuthSourceRequest) SetOidcJwks(v map[string]interface{})`

SetOidcJwks sets OidcJwks field to given value.

### HasOidcJwks

`func (o *OAuthSourceRequest) HasOidcJwks() bool`

HasOidcJwks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


