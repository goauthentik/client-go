# SCIMProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**PropertyMappingsGroup** | Pointer to **[]string** | Property mappings used for group creation/updating. | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**AssignedBackchannelApplicationSlug** | **string** | Internal application name, used in URLs. | [readonly] 
**AssignedBackchannelApplicationName** | **string** | Application&#39;s display Name. | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**Url** | **string** | Base URL to SCIM requests, usually ends in /v2 | 
**VerifyCertificates** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** | Authentication token | [optional] 
**AuthMode** | Pointer to [**SCIMAuthenticationModeEnum**](SCIMAuthenticationModeEnum.md) |  | [optional] 
**AuthOauth** | Pointer to **NullableString** | OAuth Source used for authentication | [optional] 
**AuthOauthParams** | Pointer to **map[string]interface{}** | Additional OAuth parameters, such as grant_type | [optional] 
**CompatibilityMode** | Pointer to [**CompatibilityModeEnum**](CompatibilityModeEnum.md) | Alter authentik behavior for vendor-specific SCIM implementations. | [optional] 
**ExcludeUsersServiceAccount** | Pointer to **bool** |  | [optional] 
**FilterGroup** | Pointer to **NullableString** |  | [optional] 
**SyncPageSize** | Pointer to **int32** | Controls the number of objects synced in a single task | [optional] 
**SyncPageTimeout** | Pointer to **string** | Timeout for synchronization of a single page | [optional] 
**DryRun** | Pointer to **bool** | When enabled, provider will not modify or create objects in the remote system. | [optional] 

## Methods

### NewSCIMProvider

`func NewSCIMProvider(pk int32, name string, component string, assignedBackchannelApplicationSlug string, assignedBackchannelApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, url string, ) *SCIMProvider`

NewSCIMProvider instantiates a new SCIMProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMProviderWithDefaults

`func NewSCIMProviderWithDefaults() *SCIMProvider`

NewSCIMProviderWithDefaults instantiates a new SCIMProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *SCIMProvider) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *SCIMProvider) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *SCIMProvider) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *SCIMProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SCIMProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SCIMProvider) SetName(v string)`

SetName sets Name field to given value.


### GetPropertyMappings

`func (o *SCIMProvider) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *SCIMProvider) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *SCIMProvider) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *SCIMProvider) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetPropertyMappingsGroup

`func (o *SCIMProvider) GetPropertyMappingsGroup() []string`

GetPropertyMappingsGroup returns the PropertyMappingsGroup field if non-nil, zero value otherwise.

### GetPropertyMappingsGroupOk

`func (o *SCIMProvider) GetPropertyMappingsGroupOk() (*[]string, bool)`

GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappingsGroup

`func (o *SCIMProvider) SetPropertyMappingsGroup(v []string)`

SetPropertyMappingsGroup sets PropertyMappingsGroup field to given value.

### HasPropertyMappingsGroup

`func (o *SCIMProvider) HasPropertyMappingsGroup() bool`

HasPropertyMappingsGroup returns a boolean if a field has been set.

### GetComponent

`func (o *SCIMProvider) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *SCIMProvider) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *SCIMProvider) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetAssignedBackchannelApplicationSlug

`func (o *SCIMProvider) GetAssignedBackchannelApplicationSlug() string`

GetAssignedBackchannelApplicationSlug returns the AssignedBackchannelApplicationSlug field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationSlugOk

`func (o *SCIMProvider) GetAssignedBackchannelApplicationSlugOk() (*string, bool)`

GetAssignedBackchannelApplicationSlugOk returns a tuple with the AssignedBackchannelApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationSlug

`func (o *SCIMProvider) SetAssignedBackchannelApplicationSlug(v string)`

SetAssignedBackchannelApplicationSlug sets AssignedBackchannelApplicationSlug field to given value.


### GetAssignedBackchannelApplicationName

`func (o *SCIMProvider) GetAssignedBackchannelApplicationName() string`

GetAssignedBackchannelApplicationName returns the AssignedBackchannelApplicationName field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationNameOk

`func (o *SCIMProvider) GetAssignedBackchannelApplicationNameOk() (*string, bool)`

GetAssignedBackchannelApplicationNameOk returns a tuple with the AssignedBackchannelApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationName

`func (o *SCIMProvider) SetAssignedBackchannelApplicationName(v string)`

SetAssignedBackchannelApplicationName sets AssignedBackchannelApplicationName field to given value.


### GetVerboseName

`func (o *SCIMProvider) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *SCIMProvider) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *SCIMProvider) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *SCIMProvider) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *SCIMProvider) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *SCIMProvider) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *SCIMProvider) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *SCIMProvider) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *SCIMProvider) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetUrl

`func (o *SCIMProvider) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SCIMProvider) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SCIMProvider) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVerifyCertificates

`func (o *SCIMProvider) GetVerifyCertificates() bool`

GetVerifyCertificates returns the VerifyCertificates field if non-nil, zero value otherwise.

### GetVerifyCertificatesOk

`func (o *SCIMProvider) GetVerifyCertificatesOk() (*bool, bool)`

GetVerifyCertificatesOk returns a tuple with the VerifyCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificates

`func (o *SCIMProvider) SetVerifyCertificates(v bool)`

SetVerifyCertificates sets VerifyCertificates field to given value.

### HasVerifyCertificates

`func (o *SCIMProvider) HasVerifyCertificates() bool`

HasVerifyCertificates returns a boolean if a field has been set.

### GetToken

`func (o *SCIMProvider) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SCIMProvider) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SCIMProvider) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SCIMProvider) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAuthMode

`func (o *SCIMProvider) GetAuthMode() SCIMAuthenticationModeEnum`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SCIMProvider) GetAuthModeOk() (*SCIMAuthenticationModeEnum, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SCIMProvider) SetAuthMode(v SCIMAuthenticationModeEnum)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *SCIMProvider) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthOauth

`func (o *SCIMProvider) GetAuthOauth() string`

GetAuthOauth returns the AuthOauth field if non-nil, zero value otherwise.

### GetAuthOauthOk

`func (o *SCIMProvider) GetAuthOauthOk() (*string, bool)`

GetAuthOauthOk returns a tuple with the AuthOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOauth

`func (o *SCIMProvider) SetAuthOauth(v string)`

SetAuthOauth sets AuthOauth field to given value.

### HasAuthOauth

`func (o *SCIMProvider) HasAuthOauth() bool`

HasAuthOauth returns a boolean if a field has been set.

### SetAuthOauthNil

`func (o *SCIMProvider) SetAuthOauthNil(b bool)`

 SetAuthOauthNil sets the value for AuthOauth to be an explicit nil

### UnsetAuthOauth
`func (o *SCIMProvider) UnsetAuthOauth()`

UnsetAuthOauth ensures that no value is present for AuthOauth, not even an explicit nil
### GetAuthOauthParams

`func (o *SCIMProvider) GetAuthOauthParams() map[string]interface{}`

GetAuthOauthParams returns the AuthOauthParams field if non-nil, zero value otherwise.

### GetAuthOauthParamsOk

`func (o *SCIMProvider) GetAuthOauthParamsOk() (*map[string]interface{}, bool)`

GetAuthOauthParamsOk returns a tuple with the AuthOauthParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOauthParams

`func (o *SCIMProvider) SetAuthOauthParams(v map[string]interface{})`

SetAuthOauthParams sets AuthOauthParams field to given value.

### HasAuthOauthParams

`func (o *SCIMProvider) HasAuthOauthParams() bool`

HasAuthOauthParams returns a boolean if a field has been set.

### GetCompatibilityMode

`func (o *SCIMProvider) GetCompatibilityMode() CompatibilityModeEnum`

GetCompatibilityMode returns the CompatibilityMode field if non-nil, zero value otherwise.

### GetCompatibilityModeOk

`func (o *SCIMProvider) GetCompatibilityModeOk() (*CompatibilityModeEnum, bool)`

GetCompatibilityModeOk returns a tuple with the CompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityMode

`func (o *SCIMProvider) SetCompatibilityMode(v CompatibilityModeEnum)`

SetCompatibilityMode sets CompatibilityMode field to given value.

### HasCompatibilityMode

`func (o *SCIMProvider) HasCompatibilityMode() bool`

HasCompatibilityMode returns a boolean if a field has been set.

### GetExcludeUsersServiceAccount

`func (o *SCIMProvider) GetExcludeUsersServiceAccount() bool`

GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field if non-nil, zero value otherwise.

### GetExcludeUsersServiceAccountOk

`func (o *SCIMProvider) GetExcludeUsersServiceAccountOk() (*bool, bool)`

GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsersServiceAccount

`func (o *SCIMProvider) SetExcludeUsersServiceAccount(v bool)`

SetExcludeUsersServiceAccount sets ExcludeUsersServiceAccount field to given value.

### HasExcludeUsersServiceAccount

`func (o *SCIMProvider) HasExcludeUsersServiceAccount() bool`

HasExcludeUsersServiceAccount returns a boolean if a field has been set.

### GetFilterGroup

`func (o *SCIMProvider) GetFilterGroup() string`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *SCIMProvider) GetFilterGroupOk() (*string, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *SCIMProvider) SetFilterGroup(v string)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *SCIMProvider) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### SetFilterGroupNil

`func (o *SCIMProvider) SetFilterGroupNil(b bool)`

 SetFilterGroupNil sets the value for FilterGroup to be an explicit nil

### UnsetFilterGroup
`func (o *SCIMProvider) UnsetFilterGroup()`

UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
### GetSyncPageSize

`func (o *SCIMProvider) GetSyncPageSize() int32`

GetSyncPageSize returns the SyncPageSize field if non-nil, zero value otherwise.

### GetSyncPageSizeOk

`func (o *SCIMProvider) GetSyncPageSizeOk() (*int32, bool)`

GetSyncPageSizeOk returns a tuple with the SyncPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPageSize

`func (o *SCIMProvider) SetSyncPageSize(v int32)`

SetSyncPageSize sets SyncPageSize field to given value.

### HasSyncPageSize

`func (o *SCIMProvider) HasSyncPageSize() bool`

HasSyncPageSize returns a boolean if a field has been set.

### GetSyncPageTimeout

`func (o *SCIMProvider) GetSyncPageTimeout() string`

GetSyncPageTimeout returns the SyncPageTimeout field if non-nil, zero value otherwise.

### GetSyncPageTimeoutOk

`func (o *SCIMProvider) GetSyncPageTimeoutOk() (*string, bool)`

GetSyncPageTimeoutOk returns a tuple with the SyncPageTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPageTimeout

`func (o *SCIMProvider) SetSyncPageTimeout(v string)`

SetSyncPageTimeout sets SyncPageTimeout field to given value.

### HasSyncPageTimeout

`func (o *SCIMProvider) HasSyncPageTimeout() bool`

HasSyncPageTimeout returns a boolean if a field has been set.

### GetDryRun

`func (o *SCIMProvider) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *SCIMProvider) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *SCIMProvider) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *SCIMProvider) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


