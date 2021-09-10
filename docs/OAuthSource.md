# OAuthSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** | Source&#39;s display Name. | 
**Slug** | **string** | Internal source name, used in URLs. | 
**Enabled** | Pointer to **bool** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow to use when authenticating existing users. | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Flow to use when enrolling new users. | [optional] 
**Component** | **string** |  | [readonly] 
**VerboseName** | **string** |  | [readonly] 
**VerboseNamePlural** | **string** |  | [readonly] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**UserMatchingMode** | Pointer to [**UserMatchingModeEnum**](UserMatchingModeEnum.md) | How the source determines if an existing user should be authenticated or a new user enrolled. | [optional] 
**ProviderType** | **string** |  | 
**RequestTokenUrl** | Pointer to **NullableString** | URL used to request the initial token. This URL is only required for OAuth 1. | [optional] 
**AuthorizationUrl** | Pointer to **NullableString** | URL the user is redirect to to conest the flow. | [optional] 
**AccessTokenUrl** | Pointer to **NullableString** | URL used by authentik to retrive tokens. | [optional] 
**ProfileUrl** | Pointer to **NullableString** | URL used by authentik to get user information. | [optional] 
**ConsumerKey** | **string** |  | 
**CallbackUrl** | **string** |  | [readonly] 
**Type** | [**SourceType**](SourceType.md) |  | [readonly] 

## Methods

### NewOAuthSource

`func NewOAuthSource(pk string, name string, slug string, component string, verboseName string, verboseNamePlural string, providerType string, consumerKey string, callbackUrl string, type_ SourceType, ) *OAuthSource`

NewOAuthSource instantiates a new OAuthSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthSourceWithDefaults

`func NewOAuthSourceWithDefaults() *OAuthSource`

NewOAuthSourceWithDefaults instantiates a new OAuthSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *OAuthSource) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *OAuthSource) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *OAuthSource) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *OAuthSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OAuthSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OAuthSource) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *OAuthSource) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OAuthSource) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OAuthSource) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *OAuthSource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OAuthSource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OAuthSource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OAuthSource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *OAuthSource) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *OAuthSource) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *OAuthSource) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *OAuthSource) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *OAuthSource) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *OAuthSource) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *OAuthSource) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *OAuthSource) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *OAuthSource) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *OAuthSource) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *OAuthSource) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *OAuthSource) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetComponent

`func (o *OAuthSource) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *OAuthSource) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *OAuthSource) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *OAuthSource) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *OAuthSource) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *OAuthSource) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *OAuthSource) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *OAuthSource) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *OAuthSource) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetPolicyEngineMode

`func (o *OAuthSource) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *OAuthSource) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *OAuthSource) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *OAuthSource) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *OAuthSource) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *OAuthSource) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *OAuthSource) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *OAuthSource) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetProviderType

`func (o *OAuthSource) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *OAuthSource) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *OAuthSource) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetRequestTokenUrl

`func (o *OAuthSource) GetRequestTokenUrl() string`

GetRequestTokenUrl returns the RequestTokenUrl field if non-nil, zero value otherwise.

### GetRequestTokenUrlOk

`func (o *OAuthSource) GetRequestTokenUrlOk() (*string, bool)`

GetRequestTokenUrlOk returns a tuple with the RequestTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTokenUrl

`func (o *OAuthSource) SetRequestTokenUrl(v string)`

SetRequestTokenUrl sets RequestTokenUrl field to given value.

### HasRequestTokenUrl

`func (o *OAuthSource) HasRequestTokenUrl() bool`

HasRequestTokenUrl returns a boolean if a field has been set.

### SetRequestTokenUrlNil

`func (o *OAuthSource) SetRequestTokenUrlNil(b bool)`

 SetRequestTokenUrlNil sets the value for RequestTokenUrl to be an explicit nil

### UnsetRequestTokenUrl
`func (o *OAuthSource) UnsetRequestTokenUrl()`

UnsetRequestTokenUrl ensures that no value is present for RequestTokenUrl, not even an explicit nil
### GetAuthorizationUrl

`func (o *OAuthSource) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *OAuthSource) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *OAuthSource) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.

### HasAuthorizationUrl

`func (o *OAuthSource) HasAuthorizationUrl() bool`

HasAuthorizationUrl returns a boolean if a field has been set.

### SetAuthorizationUrlNil

`func (o *OAuthSource) SetAuthorizationUrlNil(b bool)`

 SetAuthorizationUrlNil sets the value for AuthorizationUrl to be an explicit nil

### UnsetAuthorizationUrl
`func (o *OAuthSource) UnsetAuthorizationUrl()`

UnsetAuthorizationUrl ensures that no value is present for AuthorizationUrl, not even an explicit nil
### GetAccessTokenUrl

`func (o *OAuthSource) GetAccessTokenUrl() string`

GetAccessTokenUrl returns the AccessTokenUrl field if non-nil, zero value otherwise.

### GetAccessTokenUrlOk

`func (o *OAuthSource) GetAccessTokenUrlOk() (*string, bool)`

GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenUrl

`func (o *OAuthSource) SetAccessTokenUrl(v string)`

SetAccessTokenUrl sets AccessTokenUrl field to given value.

### HasAccessTokenUrl

`func (o *OAuthSource) HasAccessTokenUrl() bool`

HasAccessTokenUrl returns a boolean if a field has been set.

### SetAccessTokenUrlNil

`func (o *OAuthSource) SetAccessTokenUrlNil(b bool)`

 SetAccessTokenUrlNil sets the value for AccessTokenUrl to be an explicit nil

### UnsetAccessTokenUrl
`func (o *OAuthSource) UnsetAccessTokenUrl()`

UnsetAccessTokenUrl ensures that no value is present for AccessTokenUrl, not even an explicit nil
### GetProfileUrl

`func (o *OAuthSource) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *OAuthSource) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *OAuthSource) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *OAuthSource) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### SetProfileUrlNil

`func (o *OAuthSource) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *OAuthSource) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
### GetConsumerKey

`func (o *OAuthSource) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *OAuthSource) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *OAuthSource) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.


### GetCallbackUrl

`func (o *OAuthSource) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *OAuthSource) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *OAuthSource) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetType

`func (o *OAuthSource) GetType() SourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuthSource) GetTypeOk() (*SourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuthSource) SetType(v SourceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


