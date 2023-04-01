# PatchedOAuthSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Source&#39;s display Name. | [optional] 
**Slug** | Pointer to **string** | Internal source name, used in URLs. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow to use when authenticating existing users. | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Flow to use when enrolling new users. | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**UserMatchingMode** | Pointer to [**UserMatchingModeEnum**](UserMatchingModeEnum.md) |  | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 
**ProviderType** | Pointer to [**ProviderTypeEnum**](ProviderTypeEnum.md) |  | [optional] 
**RequestTokenUrl** | Pointer to **NullableString** | URL used to request the initial token. This URL is only required for OAuth 1. | [optional] 
**AuthorizationUrl** | Pointer to **NullableString** | URL the user is redirect to to conest the flow. | [optional] 
**AccessTokenUrl** | Pointer to **NullableString** | URL used by authentik to retrieve tokens. | [optional] 
**ProfileUrl** | Pointer to **NullableString** | URL used by authentik to get user information. | [optional] 
**ConsumerKey** | Pointer to **string** |  | [optional] 
**ConsumerSecret** | Pointer to **string** |  | [optional] 
**AdditionalScopes** | Pointer to **string** |  | [optional] 
**OidcWellKnownUrl** | Pointer to **string** |  | [optional] 
**OidcJwksUrl** | Pointer to **string** |  | [optional] 
**OidcJwks** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedOAuthSourceRequest

`func NewPatchedOAuthSourceRequest() *PatchedOAuthSourceRequest`

NewPatchedOAuthSourceRequest instantiates a new PatchedOAuthSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedOAuthSourceRequestWithDefaults

`func NewPatchedOAuthSourceRequestWithDefaults() *PatchedOAuthSourceRequest`

NewPatchedOAuthSourceRequestWithDefaults instantiates a new PatchedOAuthSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedOAuthSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedOAuthSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedOAuthSourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedOAuthSourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedOAuthSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedOAuthSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedOAuthSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedOAuthSourceRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedOAuthSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedOAuthSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedOAuthSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedOAuthSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PatchedOAuthSourceRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedOAuthSourceRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedOAuthSourceRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedOAuthSourceRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedOAuthSourceRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedOAuthSourceRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *PatchedOAuthSourceRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *PatchedOAuthSourceRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *PatchedOAuthSourceRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *PatchedOAuthSourceRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *PatchedOAuthSourceRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *PatchedOAuthSourceRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetPolicyEngineMode

`func (o *PatchedOAuthSourceRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PatchedOAuthSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PatchedOAuthSourceRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PatchedOAuthSourceRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *PatchedOAuthSourceRequest) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *PatchedOAuthSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *PatchedOAuthSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *PatchedOAuthSourceRequest) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *PatchedOAuthSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *PatchedOAuthSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *PatchedOAuthSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *PatchedOAuthSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetProviderType

`func (o *PatchedOAuthSourceRequest) GetProviderType() ProviderTypeEnum`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *PatchedOAuthSourceRequest) GetProviderTypeOk() (*ProviderTypeEnum, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *PatchedOAuthSourceRequest) SetProviderType(v ProviderTypeEnum)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *PatchedOAuthSourceRequest) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetRequestTokenUrl

`func (o *PatchedOAuthSourceRequest) GetRequestTokenUrl() string`

GetRequestTokenUrl returns the RequestTokenUrl field if non-nil, zero value otherwise.

### GetRequestTokenUrlOk

`func (o *PatchedOAuthSourceRequest) GetRequestTokenUrlOk() (*string, bool)`

GetRequestTokenUrlOk returns a tuple with the RequestTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTokenUrl

`func (o *PatchedOAuthSourceRequest) SetRequestTokenUrl(v string)`

SetRequestTokenUrl sets RequestTokenUrl field to given value.

### HasRequestTokenUrl

`func (o *PatchedOAuthSourceRequest) HasRequestTokenUrl() bool`

HasRequestTokenUrl returns a boolean if a field has been set.

### SetRequestTokenUrlNil

`func (o *PatchedOAuthSourceRequest) SetRequestTokenUrlNil(b bool)`

 SetRequestTokenUrlNil sets the value for RequestTokenUrl to be an explicit nil

### UnsetRequestTokenUrl
`func (o *PatchedOAuthSourceRequest) UnsetRequestTokenUrl()`

UnsetRequestTokenUrl ensures that no value is present for RequestTokenUrl, not even an explicit nil
### GetAuthorizationUrl

`func (o *PatchedOAuthSourceRequest) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *PatchedOAuthSourceRequest) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *PatchedOAuthSourceRequest) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.

### HasAuthorizationUrl

`func (o *PatchedOAuthSourceRequest) HasAuthorizationUrl() bool`

HasAuthorizationUrl returns a boolean if a field has been set.

### SetAuthorizationUrlNil

`func (o *PatchedOAuthSourceRequest) SetAuthorizationUrlNil(b bool)`

 SetAuthorizationUrlNil sets the value for AuthorizationUrl to be an explicit nil

### UnsetAuthorizationUrl
`func (o *PatchedOAuthSourceRequest) UnsetAuthorizationUrl()`

UnsetAuthorizationUrl ensures that no value is present for AuthorizationUrl, not even an explicit nil
### GetAccessTokenUrl

`func (o *PatchedOAuthSourceRequest) GetAccessTokenUrl() string`

GetAccessTokenUrl returns the AccessTokenUrl field if non-nil, zero value otherwise.

### GetAccessTokenUrlOk

`func (o *PatchedOAuthSourceRequest) GetAccessTokenUrlOk() (*string, bool)`

GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenUrl

`func (o *PatchedOAuthSourceRequest) SetAccessTokenUrl(v string)`

SetAccessTokenUrl sets AccessTokenUrl field to given value.

### HasAccessTokenUrl

`func (o *PatchedOAuthSourceRequest) HasAccessTokenUrl() bool`

HasAccessTokenUrl returns a boolean if a field has been set.

### SetAccessTokenUrlNil

`func (o *PatchedOAuthSourceRequest) SetAccessTokenUrlNil(b bool)`

 SetAccessTokenUrlNil sets the value for AccessTokenUrl to be an explicit nil

### UnsetAccessTokenUrl
`func (o *PatchedOAuthSourceRequest) UnsetAccessTokenUrl()`

UnsetAccessTokenUrl ensures that no value is present for AccessTokenUrl, not even an explicit nil
### GetProfileUrl

`func (o *PatchedOAuthSourceRequest) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *PatchedOAuthSourceRequest) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *PatchedOAuthSourceRequest) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *PatchedOAuthSourceRequest) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### SetProfileUrlNil

`func (o *PatchedOAuthSourceRequest) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *PatchedOAuthSourceRequest) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
### GetConsumerKey

`func (o *PatchedOAuthSourceRequest) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *PatchedOAuthSourceRequest) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *PatchedOAuthSourceRequest) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.

### HasConsumerKey

`func (o *PatchedOAuthSourceRequest) HasConsumerKey() bool`

HasConsumerKey returns a boolean if a field has been set.

### GetConsumerSecret

`func (o *PatchedOAuthSourceRequest) GetConsumerSecret() string`

GetConsumerSecret returns the ConsumerSecret field if non-nil, zero value otherwise.

### GetConsumerSecretOk

`func (o *PatchedOAuthSourceRequest) GetConsumerSecretOk() (*string, bool)`

GetConsumerSecretOk returns a tuple with the ConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSecret

`func (o *PatchedOAuthSourceRequest) SetConsumerSecret(v string)`

SetConsumerSecret sets ConsumerSecret field to given value.

### HasConsumerSecret

`func (o *PatchedOAuthSourceRequest) HasConsumerSecret() bool`

HasConsumerSecret returns a boolean if a field has been set.

### GetAdditionalScopes

`func (o *PatchedOAuthSourceRequest) GetAdditionalScopes() string`

GetAdditionalScopes returns the AdditionalScopes field if non-nil, zero value otherwise.

### GetAdditionalScopesOk

`func (o *PatchedOAuthSourceRequest) GetAdditionalScopesOk() (*string, bool)`

GetAdditionalScopesOk returns a tuple with the AdditionalScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalScopes

`func (o *PatchedOAuthSourceRequest) SetAdditionalScopes(v string)`

SetAdditionalScopes sets AdditionalScopes field to given value.

### HasAdditionalScopes

`func (o *PatchedOAuthSourceRequest) HasAdditionalScopes() bool`

HasAdditionalScopes returns a boolean if a field has been set.

### GetOidcWellKnownUrl

`func (o *PatchedOAuthSourceRequest) GetOidcWellKnownUrl() string`

GetOidcWellKnownUrl returns the OidcWellKnownUrl field if non-nil, zero value otherwise.

### GetOidcWellKnownUrlOk

`func (o *PatchedOAuthSourceRequest) GetOidcWellKnownUrlOk() (*string, bool)`

GetOidcWellKnownUrlOk returns a tuple with the OidcWellKnownUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcWellKnownUrl

`func (o *PatchedOAuthSourceRequest) SetOidcWellKnownUrl(v string)`

SetOidcWellKnownUrl sets OidcWellKnownUrl field to given value.

### HasOidcWellKnownUrl

`func (o *PatchedOAuthSourceRequest) HasOidcWellKnownUrl() bool`

HasOidcWellKnownUrl returns a boolean if a field has been set.

### GetOidcJwksUrl

`func (o *PatchedOAuthSourceRequest) GetOidcJwksUrl() string`

GetOidcJwksUrl returns the OidcJwksUrl field if non-nil, zero value otherwise.

### GetOidcJwksUrlOk

`func (o *PatchedOAuthSourceRequest) GetOidcJwksUrlOk() (*string, bool)`

GetOidcJwksUrlOk returns a tuple with the OidcJwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcJwksUrl

`func (o *PatchedOAuthSourceRequest) SetOidcJwksUrl(v string)`

SetOidcJwksUrl sets OidcJwksUrl field to given value.

### HasOidcJwksUrl

`func (o *PatchedOAuthSourceRequest) HasOidcJwksUrl() bool`

HasOidcJwksUrl returns a boolean if a field has been set.

### GetOidcJwks

`func (o *PatchedOAuthSourceRequest) GetOidcJwks() map[string]interface{}`

GetOidcJwks returns the OidcJwks field if non-nil, zero value otherwise.

### GetOidcJwksOk

`func (o *PatchedOAuthSourceRequest) GetOidcJwksOk() (*map[string]interface{}, bool)`

GetOidcJwksOk returns a tuple with the OidcJwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcJwks

`func (o *PatchedOAuthSourceRequest) SetOidcJwks(v map[string]interface{})`

SetOidcJwks sets OidcJwks field to given value.

### HasOidcJwks

`func (o *PatchedOAuthSourceRequest) HasOidcJwks() bool`

HasOidcJwks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


