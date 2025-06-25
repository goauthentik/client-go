# PatchedGoogleWorkspaceProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**PropertyMappingsGroup** | Pointer to **[]string** | Property mappings used for group creation/updating. | [optional] 
**DelegatedSubject** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to **map[string]interface{}** |  | [optional] 
**Scopes** | Pointer to **string** |  | [optional] 
**ExcludeUsersServiceAccount** | Pointer to **bool** |  | [optional] 
**FilterGroup** | Pointer to **NullableString** |  | [optional] 
**UserDeleteAction** | Pointer to [**OutgoingSyncDeleteAction**](OutgoingSyncDeleteAction.md) |  | [optional] 
**GroupDeleteAction** | Pointer to [**OutgoingSyncDeleteAction**](OutgoingSyncDeleteAction.md) |  | [optional] 
**DefaultGroupEmailDomain** | Pointer to **string** |  | [optional] 
**DryRun** | Pointer to **bool** | When enabled, provider will not modify or create objects in the remote system. | [optional] 

## Methods

### NewPatchedGoogleWorkspaceProviderRequest

`func NewPatchedGoogleWorkspaceProviderRequest() *PatchedGoogleWorkspaceProviderRequest`

NewPatchedGoogleWorkspaceProviderRequest instantiates a new PatchedGoogleWorkspaceProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGoogleWorkspaceProviderRequestWithDefaults

`func NewPatchedGoogleWorkspaceProviderRequestWithDefaults() *PatchedGoogleWorkspaceProviderRequest`

NewPatchedGoogleWorkspaceProviderRequestWithDefaults instantiates a new PatchedGoogleWorkspaceProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedGoogleWorkspaceProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedGoogleWorkspaceProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedGoogleWorkspaceProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropertyMappings

`func (o *PatchedGoogleWorkspaceProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *PatchedGoogleWorkspaceProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *PatchedGoogleWorkspaceProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetPropertyMappingsGroup

`func (o *PatchedGoogleWorkspaceProviderRequest) GetPropertyMappingsGroup() []string`

GetPropertyMappingsGroup returns the PropertyMappingsGroup field if non-nil, zero value otherwise.

### GetPropertyMappingsGroupOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetPropertyMappingsGroupOk() (*[]string, bool)`

GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappingsGroup

`func (o *PatchedGoogleWorkspaceProviderRequest) SetPropertyMappingsGroup(v []string)`

SetPropertyMappingsGroup sets PropertyMappingsGroup field to given value.

### HasPropertyMappingsGroup

`func (o *PatchedGoogleWorkspaceProviderRequest) HasPropertyMappingsGroup() bool`

HasPropertyMappingsGroup returns a boolean if a field has been set.

### GetDelegatedSubject

`func (o *PatchedGoogleWorkspaceProviderRequest) GetDelegatedSubject() string`

GetDelegatedSubject returns the DelegatedSubject field if non-nil, zero value otherwise.

### GetDelegatedSubjectOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetDelegatedSubjectOk() (*string, bool)`

GetDelegatedSubjectOk returns a tuple with the DelegatedSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedSubject

`func (o *PatchedGoogleWorkspaceProviderRequest) SetDelegatedSubject(v string)`

SetDelegatedSubject sets DelegatedSubject field to given value.

### HasDelegatedSubject

`func (o *PatchedGoogleWorkspaceProviderRequest) HasDelegatedSubject() bool`

HasDelegatedSubject returns a boolean if a field has been set.

### GetCredentials

`func (o *PatchedGoogleWorkspaceProviderRequest) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PatchedGoogleWorkspaceProviderRequest) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *PatchedGoogleWorkspaceProviderRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetScopes

`func (o *PatchedGoogleWorkspaceProviderRequest) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PatchedGoogleWorkspaceProviderRequest) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PatchedGoogleWorkspaceProviderRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetExcludeUsersServiceAccount

`func (o *PatchedGoogleWorkspaceProviderRequest) GetExcludeUsersServiceAccount() bool`

GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field if non-nil, zero value otherwise.

### GetExcludeUsersServiceAccountOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool)`

GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsersServiceAccount

`func (o *PatchedGoogleWorkspaceProviderRequest) SetExcludeUsersServiceAccount(v bool)`

SetExcludeUsersServiceAccount sets ExcludeUsersServiceAccount field to given value.

### HasExcludeUsersServiceAccount

`func (o *PatchedGoogleWorkspaceProviderRequest) HasExcludeUsersServiceAccount() bool`

HasExcludeUsersServiceAccount returns a boolean if a field has been set.

### GetFilterGroup

`func (o *PatchedGoogleWorkspaceProviderRequest) GetFilterGroup() string`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetFilterGroupOk() (*string, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *PatchedGoogleWorkspaceProviderRequest) SetFilterGroup(v string)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *PatchedGoogleWorkspaceProviderRequest) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### SetFilterGroupNil

`func (o *PatchedGoogleWorkspaceProviderRequest) SetFilterGroupNil(b bool)`

 SetFilterGroupNil sets the value for FilterGroup to be an explicit nil

### UnsetFilterGroup
`func (o *PatchedGoogleWorkspaceProviderRequest) UnsetFilterGroup()`

UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
### GetUserDeleteAction

`func (o *PatchedGoogleWorkspaceProviderRequest) GetUserDeleteAction() OutgoingSyncDeleteAction`

GetUserDeleteAction returns the UserDeleteAction field if non-nil, zero value otherwise.

### GetUserDeleteActionOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetUserDeleteActionOk() (*OutgoingSyncDeleteAction, bool)`

GetUserDeleteActionOk returns a tuple with the UserDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDeleteAction

`func (o *PatchedGoogleWorkspaceProviderRequest) SetUserDeleteAction(v OutgoingSyncDeleteAction)`

SetUserDeleteAction sets UserDeleteAction field to given value.

### HasUserDeleteAction

`func (o *PatchedGoogleWorkspaceProviderRequest) HasUserDeleteAction() bool`

HasUserDeleteAction returns a boolean if a field has been set.

### GetGroupDeleteAction

`func (o *PatchedGoogleWorkspaceProviderRequest) GetGroupDeleteAction() OutgoingSyncDeleteAction`

GetGroupDeleteAction returns the GroupDeleteAction field if non-nil, zero value otherwise.

### GetGroupDeleteActionOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetGroupDeleteActionOk() (*OutgoingSyncDeleteAction, bool)`

GetGroupDeleteActionOk returns a tuple with the GroupDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDeleteAction

`func (o *PatchedGoogleWorkspaceProviderRequest) SetGroupDeleteAction(v OutgoingSyncDeleteAction)`

SetGroupDeleteAction sets GroupDeleteAction field to given value.

### HasGroupDeleteAction

`func (o *PatchedGoogleWorkspaceProviderRequest) HasGroupDeleteAction() bool`

HasGroupDeleteAction returns a boolean if a field has been set.

### GetDefaultGroupEmailDomain

`func (o *PatchedGoogleWorkspaceProviderRequest) GetDefaultGroupEmailDomain() string`

GetDefaultGroupEmailDomain returns the DefaultGroupEmailDomain field if non-nil, zero value otherwise.

### GetDefaultGroupEmailDomainOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetDefaultGroupEmailDomainOk() (*string, bool)`

GetDefaultGroupEmailDomainOk returns a tuple with the DefaultGroupEmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupEmailDomain

`func (o *PatchedGoogleWorkspaceProviderRequest) SetDefaultGroupEmailDomain(v string)`

SetDefaultGroupEmailDomain sets DefaultGroupEmailDomain field to given value.

### HasDefaultGroupEmailDomain

`func (o *PatchedGoogleWorkspaceProviderRequest) HasDefaultGroupEmailDomain() bool`

HasDefaultGroupEmailDomain returns a boolean if a field has been set.

### GetDryRun

`func (o *PatchedGoogleWorkspaceProviderRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PatchedGoogleWorkspaceProviderRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PatchedGoogleWorkspaceProviderRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PatchedGoogleWorkspaceProviderRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


