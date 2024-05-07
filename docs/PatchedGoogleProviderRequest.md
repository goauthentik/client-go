# PatchedGoogleProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**PropertyMappingsGroup** | Pointer to **[]string** | Property mappings used for group creation/updating. | [optional] 
**DelegatedSubject** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to **interface{}** |  | [optional] 
**Scopes** | Pointer to **string** |  | [optional] 
**ExcludeUsersServiceAccount** | Pointer to **bool** |  | [optional] 
**FilterGroup** | Pointer to **NullableString** |  | [optional] 
**UserDeleteAction** | Pointer to [**GoogleWorkspaceDeleteAction**](GoogleWorkspaceDeleteAction.md) |  | [optional] 
**GroupDeleteAction** | Pointer to [**GoogleWorkspaceDeleteAction**](GoogleWorkspaceDeleteAction.md) |  | [optional] 
**DefaultGroupEmailDomain** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedGoogleProviderRequest

`func NewPatchedGoogleProviderRequest() *PatchedGoogleProviderRequest`

NewPatchedGoogleProviderRequest instantiates a new PatchedGoogleProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGoogleProviderRequestWithDefaults

`func NewPatchedGoogleProviderRequestWithDefaults() *PatchedGoogleProviderRequest`

NewPatchedGoogleProviderRequestWithDefaults instantiates a new PatchedGoogleProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedGoogleProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedGoogleProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedGoogleProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedGoogleProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropertyMappings

`func (o *PatchedGoogleProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *PatchedGoogleProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *PatchedGoogleProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *PatchedGoogleProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetPropertyMappingsGroup

`func (o *PatchedGoogleProviderRequest) GetPropertyMappingsGroup() []string`

GetPropertyMappingsGroup returns the PropertyMappingsGroup field if non-nil, zero value otherwise.

### GetPropertyMappingsGroupOk

`func (o *PatchedGoogleProviderRequest) GetPropertyMappingsGroupOk() (*[]string, bool)`

GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappingsGroup

`func (o *PatchedGoogleProviderRequest) SetPropertyMappingsGroup(v []string)`

SetPropertyMappingsGroup sets PropertyMappingsGroup field to given value.

### HasPropertyMappingsGroup

`func (o *PatchedGoogleProviderRequest) HasPropertyMappingsGroup() bool`

HasPropertyMappingsGroup returns a boolean if a field has been set.

### GetDelegatedSubject

`func (o *PatchedGoogleProviderRequest) GetDelegatedSubject() string`

GetDelegatedSubject returns the DelegatedSubject field if non-nil, zero value otherwise.

### GetDelegatedSubjectOk

`func (o *PatchedGoogleProviderRequest) GetDelegatedSubjectOk() (*string, bool)`

GetDelegatedSubjectOk returns a tuple with the DelegatedSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedSubject

`func (o *PatchedGoogleProviderRequest) SetDelegatedSubject(v string)`

SetDelegatedSubject sets DelegatedSubject field to given value.

### HasDelegatedSubject

`func (o *PatchedGoogleProviderRequest) HasDelegatedSubject() bool`

HasDelegatedSubject returns a boolean if a field has been set.

### GetCredentials

`func (o *PatchedGoogleProviderRequest) GetCredentials() interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PatchedGoogleProviderRequest) GetCredentialsOk() (*interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PatchedGoogleProviderRequest) SetCredentials(v interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *PatchedGoogleProviderRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *PatchedGoogleProviderRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *PatchedGoogleProviderRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetScopes

`func (o *PatchedGoogleProviderRequest) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PatchedGoogleProviderRequest) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PatchedGoogleProviderRequest) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PatchedGoogleProviderRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetExcludeUsersServiceAccount

`func (o *PatchedGoogleProviderRequest) GetExcludeUsersServiceAccount() bool`

GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field if non-nil, zero value otherwise.

### GetExcludeUsersServiceAccountOk

`func (o *PatchedGoogleProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool)`

GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsersServiceAccount

`func (o *PatchedGoogleProviderRequest) SetExcludeUsersServiceAccount(v bool)`

SetExcludeUsersServiceAccount sets ExcludeUsersServiceAccount field to given value.

### HasExcludeUsersServiceAccount

`func (o *PatchedGoogleProviderRequest) HasExcludeUsersServiceAccount() bool`

HasExcludeUsersServiceAccount returns a boolean if a field has been set.

### GetFilterGroup

`func (o *PatchedGoogleProviderRequest) GetFilterGroup() string`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *PatchedGoogleProviderRequest) GetFilterGroupOk() (*string, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *PatchedGoogleProviderRequest) SetFilterGroup(v string)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *PatchedGoogleProviderRequest) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### SetFilterGroupNil

`func (o *PatchedGoogleProviderRequest) SetFilterGroupNil(b bool)`

 SetFilterGroupNil sets the value for FilterGroup to be an explicit nil

### UnsetFilterGroup
`func (o *PatchedGoogleProviderRequest) UnsetFilterGroup()`

UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
### GetUserDeleteAction

`func (o *PatchedGoogleProviderRequest) GetUserDeleteAction() GoogleWorkspaceDeleteAction`

GetUserDeleteAction returns the UserDeleteAction field if non-nil, zero value otherwise.

### GetUserDeleteActionOk

`func (o *PatchedGoogleProviderRequest) GetUserDeleteActionOk() (*GoogleWorkspaceDeleteAction, bool)`

GetUserDeleteActionOk returns a tuple with the UserDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDeleteAction

`func (o *PatchedGoogleProviderRequest) SetUserDeleteAction(v GoogleWorkspaceDeleteAction)`

SetUserDeleteAction sets UserDeleteAction field to given value.

### HasUserDeleteAction

`func (o *PatchedGoogleProviderRequest) HasUserDeleteAction() bool`

HasUserDeleteAction returns a boolean if a field has been set.

### GetGroupDeleteAction

`func (o *PatchedGoogleProviderRequest) GetGroupDeleteAction() GoogleWorkspaceDeleteAction`

GetGroupDeleteAction returns the GroupDeleteAction field if non-nil, zero value otherwise.

### GetGroupDeleteActionOk

`func (o *PatchedGoogleProviderRequest) GetGroupDeleteActionOk() (*GoogleWorkspaceDeleteAction, bool)`

GetGroupDeleteActionOk returns a tuple with the GroupDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDeleteAction

`func (o *PatchedGoogleProviderRequest) SetGroupDeleteAction(v GoogleWorkspaceDeleteAction)`

SetGroupDeleteAction sets GroupDeleteAction field to given value.

### HasGroupDeleteAction

`func (o *PatchedGoogleProviderRequest) HasGroupDeleteAction() bool`

HasGroupDeleteAction returns a boolean if a field has been set.

### GetDefaultGroupEmailDomain

`func (o *PatchedGoogleProviderRequest) GetDefaultGroupEmailDomain() string`

GetDefaultGroupEmailDomain returns the DefaultGroupEmailDomain field if non-nil, zero value otherwise.

### GetDefaultGroupEmailDomainOk

`func (o *PatchedGoogleProviderRequest) GetDefaultGroupEmailDomainOk() (*string, bool)`

GetDefaultGroupEmailDomainOk returns a tuple with the DefaultGroupEmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupEmailDomain

`func (o *PatchedGoogleProviderRequest) SetDefaultGroupEmailDomain(v string)`

SetDefaultGroupEmailDomain sets DefaultGroupEmailDomain field to given value.

### HasDefaultGroupEmailDomain

`func (o *PatchedGoogleProviderRequest) HasDefaultGroupEmailDomain() bool`

HasDefaultGroupEmailDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


