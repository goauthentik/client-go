# GoogleProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**PropertyMappingsGroup** | Pointer to **[]string** | Property mappings used for group creation/updating. | [optional] 
**DelegatedSubject** | **string** |  | 
**Credentials** | **interface{}** |  | 
**Scopes** | Pointer to **string** |  | [optional] 
**ExcludeUsersServiceAccount** | Pointer to **bool** |  | [optional] 
**FilterGroup** | Pointer to **NullableString** |  | [optional] 
**UserDeleteAction** | Pointer to [**GoogleWorkspaceDeleteAction**](GoogleWorkspaceDeleteAction.md) |  | [optional] 
**GroupDeleteAction** | Pointer to [**GoogleWorkspaceDeleteAction**](GoogleWorkspaceDeleteAction.md) |  | [optional] 
**DefaultGroupEmailDomain** | **string** |  | 

## Methods

### NewGoogleProviderRequest

`func NewGoogleProviderRequest(name string, delegatedSubject string, credentials interface{}, defaultGroupEmailDomain string, ) *GoogleProviderRequest`

NewGoogleProviderRequest instantiates a new GoogleProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleProviderRequestWithDefaults

`func NewGoogleProviderRequestWithDefaults() *GoogleProviderRequest`

NewGoogleProviderRequestWithDefaults instantiates a new GoogleProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GoogleProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPropertyMappings

`func (o *GoogleProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *GoogleProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *GoogleProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *GoogleProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetPropertyMappingsGroup

`func (o *GoogleProviderRequest) GetPropertyMappingsGroup() []string`

GetPropertyMappingsGroup returns the PropertyMappingsGroup field if non-nil, zero value otherwise.

### GetPropertyMappingsGroupOk

`func (o *GoogleProviderRequest) GetPropertyMappingsGroupOk() (*[]string, bool)`

GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappingsGroup

`func (o *GoogleProviderRequest) SetPropertyMappingsGroup(v []string)`

SetPropertyMappingsGroup sets PropertyMappingsGroup field to given value.

### HasPropertyMappingsGroup

`func (o *GoogleProviderRequest) HasPropertyMappingsGroup() bool`

HasPropertyMappingsGroup returns a boolean if a field has been set.

### GetDelegatedSubject

`func (o *GoogleProviderRequest) GetDelegatedSubject() string`

GetDelegatedSubject returns the DelegatedSubject field if non-nil, zero value otherwise.

### GetDelegatedSubjectOk

`func (o *GoogleProviderRequest) GetDelegatedSubjectOk() (*string, bool)`

GetDelegatedSubjectOk returns a tuple with the DelegatedSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedSubject

`func (o *GoogleProviderRequest) SetDelegatedSubject(v string)`

SetDelegatedSubject sets DelegatedSubject field to given value.


### GetCredentials

`func (o *GoogleProviderRequest) GetCredentials() interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GoogleProviderRequest) GetCredentialsOk() (*interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GoogleProviderRequest) SetCredentials(v interface{})`

SetCredentials sets Credentials field to given value.


### SetCredentialsNil

`func (o *GoogleProviderRequest) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *GoogleProviderRequest) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetScopes

`func (o *GoogleProviderRequest) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GoogleProviderRequest) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GoogleProviderRequest) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GoogleProviderRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetExcludeUsersServiceAccount

`func (o *GoogleProviderRequest) GetExcludeUsersServiceAccount() bool`

GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field if non-nil, zero value otherwise.

### GetExcludeUsersServiceAccountOk

`func (o *GoogleProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool)`

GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsersServiceAccount

`func (o *GoogleProviderRequest) SetExcludeUsersServiceAccount(v bool)`

SetExcludeUsersServiceAccount sets ExcludeUsersServiceAccount field to given value.

### HasExcludeUsersServiceAccount

`func (o *GoogleProviderRequest) HasExcludeUsersServiceAccount() bool`

HasExcludeUsersServiceAccount returns a boolean if a field has been set.

### GetFilterGroup

`func (o *GoogleProviderRequest) GetFilterGroup() string`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *GoogleProviderRequest) GetFilterGroupOk() (*string, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *GoogleProviderRequest) SetFilterGroup(v string)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *GoogleProviderRequest) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### SetFilterGroupNil

`func (o *GoogleProviderRequest) SetFilterGroupNil(b bool)`

 SetFilterGroupNil sets the value for FilterGroup to be an explicit nil

### UnsetFilterGroup
`func (o *GoogleProviderRequest) UnsetFilterGroup()`

UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
### GetUserDeleteAction

`func (o *GoogleProviderRequest) GetUserDeleteAction() GoogleWorkspaceDeleteAction`

GetUserDeleteAction returns the UserDeleteAction field if non-nil, zero value otherwise.

### GetUserDeleteActionOk

`func (o *GoogleProviderRequest) GetUserDeleteActionOk() (*GoogleWorkspaceDeleteAction, bool)`

GetUserDeleteActionOk returns a tuple with the UserDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDeleteAction

`func (o *GoogleProviderRequest) SetUserDeleteAction(v GoogleWorkspaceDeleteAction)`

SetUserDeleteAction sets UserDeleteAction field to given value.

### HasUserDeleteAction

`func (o *GoogleProviderRequest) HasUserDeleteAction() bool`

HasUserDeleteAction returns a boolean if a field has been set.

### GetGroupDeleteAction

`func (o *GoogleProviderRequest) GetGroupDeleteAction() GoogleWorkspaceDeleteAction`

GetGroupDeleteAction returns the GroupDeleteAction field if non-nil, zero value otherwise.

### GetGroupDeleteActionOk

`func (o *GoogleProviderRequest) GetGroupDeleteActionOk() (*GoogleWorkspaceDeleteAction, bool)`

GetGroupDeleteActionOk returns a tuple with the GroupDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDeleteAction

`func (o *GoogleProviderRequest) SetGroupDeleteAction(v GoogleWorkspaceDeleteAction)`

SetGroupDeleteAction sets GroupDeleteAction field to given value.

### HasGroupDeleteAction

`func (o *GoogleProviderRequest) HasGroupDeleteAction() bool`

HasGroupDeleteAction returns a boolean if a field has been set.

### GetDefaultGroupEmailDomain

`func (o *GoogleProviderRequest) GetDefaultGroupEmailDomain() string`

GetDefaultGroupEmailDomain returns the DefaultGroupEmailDomain field if non-nil, zero value otherwise.

### GetDefaultGroupEmailDomainOk

`func (o *GoogleProviderRequest) GetDefaultGroupEmailDomainOk() (*string, bool)`

GetDefaultGroupEmailDomainOk returns a tuple with the DefaultGroupEmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupEmailDomain

`func (o *GoogleProviderRequest) SetDefaultGroupEmailDomain(v string)`

SetDefaultGroupEmailDomain sets DefaultGroupEmailDomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


