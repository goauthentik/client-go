# GoogleProvider

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
**Credentials** | **interface{}** |  | 
**Scopes** | Pointer to **string** |  | [optional] 
**ExcludeUsersServiceAccount** | Pointer to **bool** |  | [optional] 
**FilterGroup** | Pointer to **NullableString** |  | [optional] 
**UserDeleteAction** | Pointer to [**GoogleWorkspaceDeleteAction**](GoogleWorkspaceDeleteAction.md) |  | [optional] 
**GroupDeleteAction** | Pointer to [**GoogleWorkspaceDeleteAction**](GoogleWorkspaceDeleteAction.md) |  | [optional] 
**DefaultGroupEmailDomain** | **string** |  | 

## Methods

### NewGoogleProvider

`func NewGoogleProvider(pk int32, name string, component string, assignedBackchannelApplicationSlug string, assignedBackchannelApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, delegatedSubject string, credentials interface{}, defaultGroupEmailDomain string, ) *GoogleProvider`

NewGoogleProvider instantiates a new GoogleProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleProviderWithDefaults

`func NewGoogleProviderWithDefaults() *GoogleProvider`

NewGoogleProviderWithDefaults instantiates a new GoogleProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *GoogleProvider) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *GoogleProvider) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *GoogleProvider) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *GoogleProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleProvider) SetName(v string)`

SetName sets Name field to given value.


### GetPropertyMappings

`func (o *GoogleProvider) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *GoogleProvider) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *GoogleProvider) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *GoogleProvider) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetPropertyMappingsGroup

`func (o *GoogleProvider) GetPropertyMappingsGroup() []string`

GetPropertyMappingsGroup returns the PropertyMappingsGroup field if non-nil, zero value otherwise.

### GetPropertyMappingsGroupOk

`func (o *GoogleProvider) GetPropertyMappingsGroupOk() (*[]string, bool)`

GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappingsGroup

`func (o *GoogleProvider) SetPropertyMappingsGroup(v []string)`

SetPropertyMappingsGroup sets PropertyMappingsGroup field to given value.

### HasPropertyMappingsGroup

`func (o *GoogleProvider) HasPropertyMappingsGroup() bool`

HasPropertyMappingsGroup returns a boolean if a field has been set.

### GetComponent

`func (o *GoogleProvider) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *GoogleProvider) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *GoogleProvider) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetAssignedBackchannelApplicationSlug

`func (o *GoogleProvider) GetAssignedBackchannelApplicationSlug() string`

GetAssignedBackchannelApplicationSlug returns the AssignedBackchannelApplicationSlug field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationSlugOk

`func (o *GoogleProvider) GetAssignedBackchannelApplicationSlugOk() (*string, bool)`

GetAssignedBackchannelApplicationSlugOk returns a tuple with the AssignedBackchannelApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationSlug

`func (o *GoogleProvider) SetAssignedBackchannelApplicationSlug(v string)`

SetAssignedBackchannelApplicationSlug sets AssignedBackchannelApplicationSlug field to given value.


### GetAssignedBackchannelApplicationName

`func (o *GoogleProvider) GetAssignedBackchannelApplicationName() string`

GetAssignedBackchannelApplicationName returns the AssignedBackchannelApplicationName field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationNameOk

`func (o *GoogleProvider) GetAssignedBackchannelApplicationNameOk() (*string, bool)`

GetAssignedBackchannelApplicationNameOk returns a tuple with the AssignedBackchannelApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationName

`func (o *GoogleProvider) SetAssignedBackchannelApplicationName(v string)`

SetAssignedBackchannelApplicationName sets AssignedBackchannelApplicationName field to given value.


### GetVerboseName

`func (o *GoogleProvider) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *GoogleProvider) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *GoogleProvider) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *GoogleProvider) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *GoogleProvider) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *GoogleProvider) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *GoogleProvider) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *GoogleProvider) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *GoogleProvider) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetDelegatedSubject

`func (o *GoogleProvider) GetDelegatedSubject() string`

GetDelegatedSubject returns the DelegatedSubject field if non-nil, zero value otherwise.

### GetDelegatedSubjectOk

`func (o *GoogleProvider) GetDelegatedSubjectOk() (*string, bool)`

GetDelegatedSubjectOk returns a tuple with the DelegatedSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedSubject

`func (o *GoogleProvider) SetDelegatedSubject(v string)`

SetDelegatedSubject sets DelegatedSubject field to given value.


### GetCredentials

`func (o *GoogleProvider) GetCredentials() interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GoogleProvider) GetCredentialsOk() (*interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GoogleProvider) SetCredentials(v interface{})`

SetCredentials sets Credentials field to given value.


### SetCredentialsNil

`func (o *GoogleProvider) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *GoogleProvider) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetScopes

`func (o *GoogleProvider) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GoogleProvider) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GoogleProvider) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GoogleProvider) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetExcludeUsersServiceAccount

`func (o *GoogleProvider) GetExcludeUsersServiceAccount() bool`

GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field if non-nil, zero value otherwise.

### GetExcludeUsersServiceAccountOk

`func (o *GoogleProvider) GetExcludeUsersServiceAccountOk() (*bool, bool)`

GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsersServiceAccount

`func (o *GoogleProvider) SetExcludeUsersServiceAccount(v bool)`

SetExcludeUsersServiceAccount sets ExcludeUsersServiceAccount field to given value.

### HasExcludeUsersServiceAccount

`func (o *GoogleProvider) HasExcludeUsersServiceAccount() bool`

HasExcludeUsersServiceAccount returns a boolean if a field has been set.

### GetFilterGroup

`func (o *GoogleProvider) GetFilterGroup() string`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *GoogleProvider) GetFilterGroupOk() (*string, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *GoogleProvider) SetFilterGroup(v string)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *GoogleProvider) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### SetFilterGroupNil

`func (o *GoogleProvider) SetFilterGroupNil(b bool)`

 SetFilterGroupNil sets the value for FilterGroup to be an explicit nil

### UnsetFilterGroup
`func (o *GoogleProvider) UnsetFilterGroup()`

UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
### GetUserDeleteAction

`func (o *GoogleProvider) GetUserDeleteAction() GoogleWorkspaceDeleteAction`

GetUserDeleteAction returns the UserDeleteAction field if non-nil, zero value otherwise.

### GetUserDeleteActionOk

`func (o *GoogleProvider) GetUserDeleteActionOk() (*GoogleWorkspaceDeleteAction, bool)`

GetUserDeleteActionOk returns a tuple with the UserDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDeleteAction

`func (o *GoogleProvider) SetUserDeleteAction(v GoogleWorkspaceDeleteAction)`

SetUserDeleteAction sets UserDeleteAction field to given value.

### HasUserDeleteAction

`func (o *GoogleProvider) HasUserDeleteAction() bool`

HasUserDeleteAction returns a boolean if a field has been set.

### GetGroupDeleteAction

`func (o *GoogleProvider) GetGroupDeleteAction() GoogleWorkspaceDeleteAction`

GetGroupDeleteAction returns the GroupDeleteAction field if non-nil, zero value otherwise.

### GetGroupDeleteActionOk

`func (o *GoogleProvider) GetGroupDeleteActionOk() (*GoogleWorkspaceDeleteAction, bool)`

GetGroupDeleteActionOk returns a tuple with the GroupDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDeleteAction

`func (o *GoogleProvider) SetGroupDeleteAction(v GoogleWorkspaceDeleteAction)`

SetGroupDeleteAction sets GroupDeleteAction field to given value.

### HasGroupDeleteAction

`func (o *GoogleProvider) HasGroupDeleteAction() bool`

HasGroupDeleteAction returns a boolean if a field has been set.

### GetDefaultGroupEmailDomain

`func (o *GoogleProvider) GetDefaultGroupEmailDomain() string`

GetDefaultGroupEmailDomain returns the DefaultGroupEmailDomain field if non-nil, zero value otherwise.

### GetDefaultGroupEmailDomainOk

`func (o *GoogleProvider) GetDefaultGroupEmailDomainOk() (*string, bool)`

GetDefaultGroupEmailDomainOk returns a tuple with the DefaultGroupEmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupEmailDomain

`func (o *GoogleProvider) SetDefaultGroupEmailDomain(v string)`

SetDefaultGroupEmailDomain sets DefaultGroupEmailDomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


