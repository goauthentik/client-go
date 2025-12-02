# PlexSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** | Source&#39;s display Name. | 
**Slug** | **string** | Internal source name, used in URLs. | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Promoted** | Pointer to **bool** | When enabled, this source will be displayed as a prominent button on the login page, instead of a small icon. | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow to use when authenticating existing users. | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Flow to use when enrolling new users. | [optional] 
**UserPropertyMappings** | Pointer to **[]string** |  | [optional] 
**GroupPropertyMappings** | Pointer to **[]string** |  | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**UserMatchingMode** | Pointer to [**UserMatchingModeEnum**](UserMatchingModeEnum.md) | How the source determines if an existing user should be authenticated or a new user enrolled. | [optional] 
**Managed** | **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [readonly] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**IconUrl** | **string** |  | [readonly] 
**GroupMatchingMode** | Pointer to [**GroupMatchingModeEnum**](GroupMatchingModeEnum.md) | How the source determines if an existing group should be used or a new group created. | [optional] 
**ClientId** | Pointer to **string** | Client identifier used to talk to Plex. | [optional] 
**AllowedServers** | Pointer to **[]string** | Which servers a user has to be a member of to be granted access. Empty list allows every server. | [optional] 
**AllowFriends** | Pointer to **bool** | Allow friends to authenticate, even if you don&#39;t share a server. | [optional] 
**PlexToken** | **string** | Plex token used to check friends | 

## Methods

### NewPlexSource

`func NewPlexSource(pk string, name string, slug string, component string, verboseName string, verboseNamePlural string, metaModelName string, managed NullableString, iconUrl string, plexToken string, ) *PlexSource`

NewPlexSource instantiates a new PlexSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlexSourceWithDefaults

`func NewPlexSourceWithDefaults() *PlexSource`

NewPlexSourceWithDefaults instantiates a new PlexSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *PlexSource) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *PlexSource) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *PlexSource) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *PlexSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlexSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlexSource) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *PlexSource) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PlexSource) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PlexSource) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *PlexSource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PlexSource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PlexSource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PlexSource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPromoted

`func (o *PlexSource) GetPromoted() bool`

GetPromoted returns the Promoted field if non-nil, zero value otherwise.

### GetPromotedOk

`func (o *PlexSource) GetPromotedOk() (*bool, bool)`

GetPromotedOk returns a tuple with the Promoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoted

`func (o *PlexSource) SetPromoted(v bool)`

SetPromoted sets Promoted field to given value.

### HasPromoted

`func (o *PlexSource) HasPromoted() bool`

HasPromoted returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PlexSource) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PlexSource) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PlexSource) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PlexSource) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PlexSource) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PlexSource) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *PlexSource) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *PlexSource) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *PlexSource) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *PlexSource) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *PlexSource) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *PlexSource) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetUserPropertyMappings

`func (o *PlexSource) GetUserPropertyMappings() []string`

GetUserPropertyMappings returns the UserPropertyMappings field if non-nil, zero value otherwise.

### GetUserPropertyMappingsOk

`func (o *PlexSource) GetUserPropertyMappingsOk() (*[]string, bool)`

GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertyMappings

`func (o *PlexSource) SetUserPropertyMappings(v []string)`

SetUserPropertyMappings sets UserPropertyMappings field to given value.

### HasUserPropertyMappings

`func (o *PlexSource) HasUserPropertyMappings() bool`

HasUserPropertyMappings returns a boolean if a field has been set.

### GetGroupPropertyMappings

`func (o *PlexSource) GetGroupPropertyMappings() []string`

GetGroupPropertyMappings returns the GroupPropertyMappings field if non-nil, zero value otherwise.

### GetGroupPropertyMappingsOk

`func (o *PlexSource) GetGroupPropertyMappingsOk() (*[]string, bool)`

GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPropertyMappings

`func (o *PlexSource) SetGroupPropertyMappings(v []string)`

SetGroupPropertyMappings sets GroupPropertyMappings field to given value.

### HasGroupPropertyMappings

`func (o *PlexSource) HasGroupPropertyMappings() bool`

HasGroupPropertyMappings returns a boolean if a field has been set.

### GetComponent

`func (o *PlexSource) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PlexSource) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PlexSource) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *PlexSource) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *PlexSource) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *PlexSource) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *PlexSource) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *PlexSource) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *PlexSource) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *PlexSource) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *PlexSource) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *PlexSource) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetPolicyEngineMode

`func (o *PlexSource) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PlexSource) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PlexSource) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PlexSource) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *PlexSource) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *PlexSource) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *PlexSource) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *PlexSource) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetManaged

`func (o *PlexSource) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PlexSource) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PlexSource) SetManaged(v string)`

SetManaged sets Managed field to given value.


### SetManagedNil

`func (o *PlexSource) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *PlexSource) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetUserPathTemplate

`func (o *PlexSource) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *PlexSource) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *PlexSource) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *PlexSource) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetIcon

`func (o *PlexSource) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PlexSource) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PlexSource) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PlexSource) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetIconUrl

`func (o *PlexSource) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *PlexSource) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *PlexSource) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### GetGroupMatchingMode

`func (o *PlexSource) GetGroupMatchingMode() GroupMatchingModeEnum`

GetGroupMatchingMode returns the GroupMatchingMode field if non-nil, zero value otherwise.

### GetGroupMatchingModeOk

`func (o *PlexSource) GetGroupMatchingModeOk() (*GroupMatchingModeEnum, bool)`

GetGroupMatchingModeOk returns a tuple with the GroupMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMatchingMode

`func (o *PlexSource) SetGroupMatchingMode(v GroupMatchingModeEnum)`

SetGroupMatchingMode sets GroupMatchingMode field to given value.

### HasGroupMatchingMode

`func (o *PlexSource) HasGroupMatchingMode() bool`

HasGroupMatchingMode returns a boolean if a field has been set.

### GetClientId

`func (o *PlexSource) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PlexSource) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PlexSource) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PlexSource) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAllowedServers

`func (o *PlexSource) GetAllowedServers() []string`

GetAllowedServers returns the AllowedServers field if non-nil, zero value otherwise.

### GetAllowedServersOk

`func (o *PlexSource) GetAllowedServersOk() (*[]string, bool)`

GetAllowedServersOk returns a tuple with the AllowedServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedServers

`func (o *PlexSource) SetAllowedServers(v []string)`

SetAllowedServers sets AllowedServers field to given value.

### HasAllowedServers

`func (o *PlexSource) HasAllowedServers() bool`

HasAllowedServers returns a boolean if a field has been set.

### GetAllowFriends

`func (o *PlexSource) GetAllowFriends() bool`

GetAllowFriends returns the AllowFriends field if non-nil, zero value otherwise.

### GetAllowFriendsOk

`func (o *PlexSource) GetAllowFriendsOk() (*bool, bool)`

GetAllowFriendsOk returns a tuple with the AllowFriends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFriends

`func (o *PlexSource) SetAllowFriends(v bool)`

SetAllowFriends sets AllowFriends field to given value.

### HasAllowFriends

`func (o *PlexSource) HasAllowFriends() bool`

HasAllowFriends returns a boolean if a field has been set.

### GetPlexToken

`func (o *PlexSource) GetPlexToken() string`

GetPlexToken returns the PlexToken field if non-nil, zero value otherwise.

### GetPlexTokenOk

`func (o *PlexSource) GetPlexTokenOk() (*string, bool)`

GetPlexTokenOk returns a tuple with the PlexToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlexToken

`func (o *PlexSource) SetPlexToken(v string)`

SetPlexToken sets PlexToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


