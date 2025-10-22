# GoogleWorkspaceProvider

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
**DelegatedSubject** | **string** |  | 
**Credentials** | **map[string]interface{}** |  | 
**Scopes** | Pointer to **string** |  | [optional] 
**ExcludeUsersServiceAccount** | Pointer to **bool** |  | [optional] 
**FilterGroup** | Pointer to **NullableString** |  | [optional] 
**UserDeleteAction** | Pointer to [**OutgoingSyncDeleteAction**](OutgoingSyncDeleteAction.md) |  | [optional] 
**GroupDeleteAction** | Pointer to [**OutgoingSyncDeleteAction**](OutgoingSyncDeleteAction.md) |  | [optional] 
**DefaultGroupEmailDomain** | **string** |  | 
**SyncPageSize** | Pointer to **int32** | Controls the number of objects synced in a single task | [optional] 
**SyncPageTimeout** | Pointer to **string** | Timeout for synchronization of a single page | [optional] 
**DryRun** | Pointer to **bool** | When enabled, provider will not modify or create objects in the remote system. | [optional] 

## Methods

### NewGoogleWorkspaceProvider

`func NewGoogleWorkspaceProvider(pk int32, name string, component string, assignedBackchannelApplicationSlug string, assignedBackchannelApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, delegatedSubject string, credentials map[string]interface{}, defaultGroupEmailDomain string, ) *GoogleWorkspaceProvider`

NewGoogleWorkspaceProvider instantiates a new GoogleWorkspaceProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleWorkspaceProviderWithDefaults

`func NewGoogleWorkspaceProviderWithDefaults() *GoogleWorkspaceProvider`

NewGoogleWorkspaceProviderWithDefaults instantiates a new GoogleWorkspaceProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *GoogleWorkspaceProvider) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *GoogleWorkspaceProvider) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *GoogleWorkspaceProvider) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *GoogleWorkspaceProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleWorkspaceProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleWorkspaceProvider) SetName(v string)`

SetName sets Name field to given value.


### GetPropertyMappings

`func (o *GoogleWorkspaceProvider) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *GoogleWorkspaceProvider) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *GoogleWorkspaceProvider) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *GoogleWorkspaceProvider) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetPropertyMappingsGroup

`func (o *GoogleWorkspaceProvider) GetPropertyMappingsGroup() []string`

GetPropertyMappingsGroup returns the PropertyMappingsGroup field if non-nil, zero value otherwise.

### GetPropertyMappingsGroupOk

`func (o *GoogleWorkspaceProvider) GetPropertyMappingsGroupOk() (*[]string, bool)`

GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappingsGroup

`func (o *GoogleWorkspaceProvider) SetPropertyMappingsGroup(v []string)`

SetPropertyMappingsGroup sets PropertyMappingsGroup field to given value.

### HasPropertyMappingsGroup

`func (o *GoogleWorkspaceProvider) HasPropertyMappingsGroup() bool`

HasPropertyMappingsGroup returns a boolean if a field has been set.

### GetComponent

`func (o *GoogleWorkspaceProvider) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *GoogleWorkspaceProvider) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *GoogleWorkspaceProvider) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetAssignedBackchannelApplicationSlug

`func (o *GoogleWorkspaceProvider) GetAssignedBackchannelApplicationSlug() string`

GetAssignedBackchannelApplicationSlug returns the AssignedBackchannelApplicationSlug field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationSlugOk

`func (o *GoogleWorkspaceProvider) GetAssignedBackchannelApplicationSlugOk() (*string, bool)`

GetAssignedBackchannelApplicationSlugOk returns a tuple with the AssignedBackchannelApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationSlug

`func (o *GoogleWorkspaceProvider) SetAssignedBackchannelApplicationSlug(v string)`

SetAssignedBackchannelApplicationSlug sets AssignedBackchannelApplicationSlug field to given value.


### GetAssignedBackchannelApplicationName

`func (o *GoogleWorkspaceProvider) GetAssignedBackchannelApplicationName() string`

GetAssignedBackchannelApplicationName returns the AssignedBackchannelApplicationName field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationNameOk

`func (o *GoogleWorkspaceProvider) GetAssignedBackchannelApplicationNameOk() (*string, bool)`

GetAssignedBackchannelApplicationNameOk returns a tuple with the AssignedBackchannelApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationName

`func (o *GoogleWorkspaceProvider) SetAssignedBackchannelApplicationName(v string)`

SetAssignedBackchannelApplicationName sets AssignedBackchannelApplicationName field to given value.


### GetVerboseName

`func (o *GoogleWorkspaceProvider) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *GoogleWorkspaceProvider) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *GoogleWorkspaceProvider) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *GoogleWorkspaceProvider) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *GoogleWorkspaceProvider) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *GoogleWorkspaceProvider) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *GoogleWorkspaceProvider) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *GoogleWorkspaceProvider) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *GoogleWorkspaceProvider) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetDelegatedSubject

`func (o *GoogleWorkspaceProvider) GetDelegatedSubject() string`

GetDelegatedSubject returns the DelegatedSubject field if non-nil, zero value otherwise.

### GetDelegatedSubjectOk

`func (o *GoogleWorkspaceProvider) GetDelegatedSubjectOk() (*string, bool)`

GetDelegatedSubjectOk returns a tuple with the DelegatedSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedSubject

`func (o *GoogleWorkspaceProvider) SetDelegatedSubject(v string)`

SetDelegatedSubject sets DelegatedSubject field to given value.


### GetCredentials

`func (o *GoogleWorkspaceProvider) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GoogleWorkspaceProvider) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GoogleWorkspaceProvider) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.


### GetScopes

`func (o *GoogleWorkspaceProvider) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GoogleWorkspaceProvider) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GoogleWorkspaceProvider) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GoogleWorkspaceProvider) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetExcludeUsersServiceAccount

`func (o *GoogleWorkspaceProvider) GetExcludeUsersServiceAccount() bool`

GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field if non-nil, zero value otherwise.

### GetExcludeUsersServiceAccountOk

`func (o *GoogleWorkspaceProvider) GetExcludeUsersServiceAccountOk() (*bool, bool)`

GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsersServiceAccount

`func (o *GoogleWorkspaceProvider) SetExcludeUsersServiceAccount(v bool)`

SetExcludeUsersServiceAccount sets ExcludeUsersServiceAccount field to given value.

### HasExcludeUsersServiceAccount

`func (o *GoogleWorkspaceProvider) HasExcludeUsersServiceAccount() bool`

HasExcludeUsersServiceAccount returns a boolean if a field has been set.

### GetFilterGroup

`func (o *GoogleWorkspaceProvider) GetFilterGroup() string`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *GoogleWorkspaceProvider) GetFilterGroupOk() (*string, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *GoogleWorkspaceProvider) SetFilterGroup(v string)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *GoogleWorkspaceProvider) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### SetFilterGroupNil

`func (o *GoogleWorkspaceProvider) SetFilterGroupNil(b bool)`

 SetFilterGroupNil sets the value for FilterGroup to be an explicit nil

### UnsetFilterGroup
`func (o *GoogleWorkspaceProvider) UnsetFilterGroup()`

UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
### GetUserDeleteAction

`func (o *GoogleWorkspaceProvider) GetUserDeleteAction() OutgoingSyncDeleteAction`

GetUserDeleteAction returns the UserDeleteAction field if non-nil, zero value otherwise.

### GetUserDeleteActionOk

`func (o *GoogleWorkspaceProvider) GetUserDeleteActionOk() (*OutgoingSyncDeleteAction, bool)`

GetUserDeleteActionOk returns a tuple with the UserDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDeleteAction

`func (o *GoogleWorkspaceProvider) SetUserDeleteAction(v OutgoingSyncDeleteAction)`

SetUserDeleteAction sets UserDeleteAction field to given value.

### HasUserDeleteAction

`func (o *GoogleWorkspaceProvider) HasUserDeleteAction() bool`

HasUserDeleteAction returns a boolean if a field has been set.

### GetGroupDeleteAction

`func (o *GoogleWorkspaceProvider) GetGroupDeleteAction() OutgoingSyncDeleteAction`

GetGroupDeleteAction returns the GroupDeleteAction field if non-nil, zero value otherwise.

### GetGroupDeleteActionOk

`func (o *GoogleWorkspaceProvider) GetGroupDeleteActionOk() (*OutgoingSyncDeleteAction, bool)`

GetGroupDeleteActionOk returns a tuple with the GroupDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDeleteAction

`func (o *GoogleWorkspaceProvider) SetGroupDeleteAction(v OutgoingSyncDeleteAction)`

SetGroupDeleteAction sets GroupDeleteAction field to given value.

### HasGroupDeleteAction

`func (o *GoogleWorkspaceProvider) HasGroupDeleteAction() bool`

HasGroupDeleteAction returns a boolean if a field has been set.

### GetDefaultGroupEmailDomain

`func (o *GoogleWorkspaceProvider) GetDefaultGroupEmailDomain() string`

GetDefaultGroupEmailDomain returns the DefaultGroupEmailDomain field if non-nil, zero value otherwise.

### GetDefaultGroupEmailDomainOk

`func (o *GoogleWorkspaceProvider) GetDefaultGroupEmailDomainOk() (*string, bool)`

GetDefaultGroupEmailDomainOk returns a tuple with the DefaultGroupEmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupEmailDomain

`func (o *GoogleWorkspaceProvider) SetDefaultGroupEmailDomain(v string)`

SetDefaultGroupEmailDomain sets DefaultGroupEmailDomain field to given value.


### GetSyncPageSize

`func (o *GoogleWorkspaceProvider) GetSyncPageSize() int32`

GetSyncPageSize returns the SyncPageSize field if non-nil, zero value otherwise.

### GetSyncPageSizeOk

`func (o *GoogleWorkspaceProvider) GetSyncPageSizeOk() (*int32, bool)`

GetSyncPageSizeOk returns a tuple with the SyncPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPageSize

`func (o *GoogleWorkspaceProvider) SetSyncPageSize(v int32)`

SetSyncPageSize sets SyncPageSize field to given value.

### HasSyncPageSize

`func (o *GoogleWorkspaceProvider) HasSyncPageSize() bool`

HasSyncPageSize returns a boolean if a field has been set.

### GetSyncPageTimeout

`func (o *GoogleWorkspaceProvider) GetSyncPageTimeout() string`

GetSyncPageTimeout returns the SyncPageTimeout field if non-nil, zero value otherwise.

### GetSyncPageTimeoutOk

`func (o *GoogleWorkspaceProvider) GetSyncPageTimeoutOk() (*string, bool)`

GetSyncPageTimeoutOk returns a tuple with the SyncPageTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPageTimeout

`func (o *GoogleWorkspaceProvider) SetSyncPageTimeout(v string)`

SetSyncPageTimeout sets SyncPageTimeout field to given value.

### HasSyncPageTimeout

`func (o *GoogleWorkspaceProvider) HasSyncPageTimeout() bool`

HasSyncPageTimeout returns a boolean if a field has been set.

### GetDryRun

`func (o *GoogleWorkspaceProvider) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *GoogleWorkspaceProvider) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *GoogleWorkspaceProvider) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *GoogleWorkspaceProvider) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


