# PlexSourceRequest

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
**ClientId** | Pointer to **string** | Client identifier used to talk to Plex. | [optional] 
**AllowedServers** | Pointer to **[]string** | Which servers a user has to be a member of to be granted access. Empty list allows every server. | [optional] 
**AllowFriends** | Pointer to **bool** | Allow friends to authenticate, even if you don&#39;t share a server. | [optional] 
**PlexToken** | **string** | Plex token used to check friends | 

## Methods

### NewPlexSourceRequest

`func NewPlexSourceRequest(name string, slug string, plexToken string, ) *PlexSourceRequest`

NewPlexSourceRequest instantiates a new PlexSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlexSourceRequestWithDefaults

`func NewPlexSourceRequestWithDefaults() *PlexSourceRequest`

NewPlexSourceRequestWithDefaults instantiates a new PlexSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlexSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlexSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlexSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *PlexSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PlexSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PlexSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *PlexSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PlexSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PlexSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PlexSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PlexSourceRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PlexSourceRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PlexSourceRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PlexSourceRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PlexSourceRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PlexSourceRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *PlexSourceRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *PlexSourceRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *PlexSourceRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *PlexSourceRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *PlexSourceRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *PlexSourceRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetPolicyEngineMode

`func (o *PlexSourceRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PlexSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PlexSourceRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PlexSourceRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *PlexSourceRequest) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *PlexSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *PlexSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *PlexSourceRequest) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *PlexSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *PlexSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *PlexSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *PlexSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetClientId

`func (o *PlexSourceRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PlexSourceRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PlexSourceRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PlexSourceRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAllowedServers

`func (o *PlexSourceRequest) GetAllowedServers() []string`

GetAllowedServers returns the AllowedServers field if non-nil, zero value otherwise.

### GetAllowedServersOk

`func (o *PlexSourceRequest) GetAllowedServersOk() (*[]string, bool)`

GetAllowedServersOk returns a tuple with the AllowedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServers

`func (o *PlexSourceRequest) SetAllowedServers(v []string)`

SetAllowedServers sets AllowedServers field to given value.

### HasAllowedServers

`func (o *PlexSourceRequest) HasAllowedServers() bool`

HasAllowedServers returns a boolean if a field has been set.

### GetAllowFriends

`func (o *PlexSourceRequest) GetAllowFriends() bool`

GetAllowFriends returns the AllowFriends field if non-nil, zero value otherwise.

### GetAllowFriendsOk

`func (o *PlexSourceRequest) GetAllowFriendsOk() (*bool, bool)`

GetAllowFriendsOk returns a tuple with the AllowFriends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFriends

`func (o *PlexSourceRequest) SetAllowFriends(v bool)`

SetAllowFriends sets AllowFriends field to given value.

### HasAllowFriends

`func (o *PlexSourceRequest) HasAllowFriends() bool`

HasAllowFriends returns a boolean if a field has been set.

### GetPlexToken

`func (o *PlexSourceRequest) GetPlexToken() string`

GetPlexToken returns the PlexToken field if non-nil, zero value otherwise.

### GetPlexTokenOk

`func (o *PlexSourceRequest) GetPlexTokenOk() (*string, bool)`

GetPlexTokenOk returns a tuple with the PlexToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlexToken

`func (o *PlexSourceRequest) SetPlexToken(v string)`

SetPlexToken sets PlexToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


