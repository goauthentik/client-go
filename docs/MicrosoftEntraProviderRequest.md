# MicrosoftEntraProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**PropertyMappingsGroup** | Pointer to **[]string** | Property mappings used for group creation/updating. | [optional] 
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**TenantId** | **string** |  | 
**ExcludeUsersServiceAccount** | Pointer to **bool** |  | [optional] 
**FilterGroup** | Pointer to **NullableString** |  | [optional] 
**UserDeleteAction** | Pointer to [**OutgoingSyncDeleteAction**](OutgoingSyncDeleteAction.md) |  | [optional] 
**GroupDeleteAction** | Pointer to [**OutgoingSyncDeleteAction**](OutgoingSyncDeleteAction.md) |  | [optional] 
**SyncPageSize** | Pointer to **int32** | Controls the number of objects synced in a single task | [optional] 
**SyncPageTimeout** | Pointer to **string** | Timeout for synchronization of a single page | [optional] 
**DryRun** | Pointer to **bool** | When enabled, provider will not modify or create objects in the remote system. | [optional] 

## Methods

### NewMicrosoftEntraProviderRequest

`func NewMicrosoftEntraProviderRequest(name string, clientId string, clientSecret string, tenantId string, ) *MicrosoftEntraProviderRequest`

NewMicrosoftEntraProviderRequest instantiates a new MicrosoftEntraProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftEntraProviderRequestWithDefaults

`func NewMicrosoftEntraProviderRequestWithDefaults() *MicrosoftEntraProviderRequest`

NewMicrosoftEntraProviderRequestWithDefaults instantiates a new MicrosoftEntraProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MicrosoftEntraProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftEntraProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftEntraProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPropertyMappings

`func (o *MicrosoftEntraProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *MicrosoftEntraProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *MicrosoftEntraProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *MicrosoftEntraProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetPropertyMappingsGroup

`func (o *MicrosoftEntraProviderRequest) GetPropertyMappingsGroup() []string`

GetPropertyMappingsGroup returns the PropertyMappingsGroup field if non-nil, zero value otherwise.

### GetPropertyMappingsGroupOk

`func (o *MicrosoftEntraProviderRequest) GetPropertyMappingsGroupOk() (*[]string, bool)`

GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappingsGroup

`func (o *MicrosoftEntraProviderRequest) SetPropertyMappingsGroup(v []string)`

SetPropertyMappingsGroup sets PropertyMappingsGroup field to given value.

### HasPropertyMappingsGroup

`func (o *MicrosoftEntraProviderRequest) HasPropertyMappingsGroup() bool`

HasPropertyMappingsGroup returns a boolean if a field has been set.

### GetClientId

`func (o *MicrosoftEntraProviderRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MicrosoftEntraProviderRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MicrosoftEntraProviderRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *MicrosoftEntraProviderRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *MicrosoftEntraProviderRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *MicrosoftEntraProviderRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTenantId

`func (o *MicrosoftEntraProviderRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MicrosoftEntraProviderRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MicrosoftEntraProviderRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetExcludeUsersServiceAccount

`func (o *MicrosoftEntraProviderRequest) GetExcludeUsersServiceAccount() bool`

GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field if non-nil, zero value otherwise.

### GetExcludeUsersServiceAccountOk

`func (o *MicrosoftEntraProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool)`

GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsersServiceAccount

`func (o *MicrosoftEntraProviderRequest) SetExcludeUsersServiceAccount(v bool)`

SetExcludeUsersServiceAccount sets ExcludeUsersServiceAccount field to given value.

### HasExcludeUsersServiceAccount

`func (o *MicrosoftEntraProviderRequest) HasExcludeUsersServiceAccount() bool`

HasExcludeUsersServiceAccount returns a boolean if a field has been set.

### GetFilterGroup

`func (o *MicrosoftEntraProviderRequest) GetFilterGroup() string`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *MicrosoftEntraProviderRequest) GetFilterGroupOk() (*string, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *MicrosoftEntraProviderRequest) SetFilterGroup(v string)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *MicrosoftEntraProviderRequest) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### SetFilterGroupNil

`func (o *MicrosoftEntraProviderRequest) SetFilterGroupNil(b bool)`

 SetFilterGroupNil sets the value for FilterGroup to be an explicit nil

### UnsetFilterGroup
`func (o *MicrosoftEntraProviderRequest) UnsetFilterGroup()`

UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
### GetUserDeleteAction

`func (o *MicrosoftEntraProviderRequest) GetUserDeleteAction() OutgoingSyncDeleteAction`

GetUserDeleteAction returns the UserDeleteAction field if non-nil, zero value otherwise.

### GetUserDeleteActionOk

`func (o *MicrosoftEntraProviderRequest) GetUserDeleteActionOk() (*OutgoingSyncDeleteAction, bool)`

GetUserDeleteActionOk returns a tuple with the UserDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDeleteAction

`func (o *MicrosoftEntraProviderRequest) SetUserDeleteAction(v OutgoingSyncDeleteAction)`

SetUserDeleteAction sets UserDeleteAction field to given value.

### HasUserDeleteAction

`func (o *MicrosoftEntraProviderRequest) HasUserDeleteAction() bool`

HasUserDeleteAction returns a boolean if a field has been set.

### GetGroupDeleteAction

`func (o *MicrosoftEntraProviderRequest) GetGroupDeleteAction() OutgoingSyncDeleteAction`

GetGroupDeleteAction returns the GroupDeleteAction field if non-nil, zero value otherwise.

### GetGroupDeleteActionOk

`func (o *MicrosoftEntraProviderRequest) GetGroupDeleteActionOk() (*OutgoingSyncDeleteAction, bool)`

GetGroupDeleteActionOk returns a tuple with the GroupDeleteAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDeleteAction

`func (o *MicrosoftEntraProviderRequest) SetGroupDeleteAction(v OutgoingSyncDeleteAction)`

SetGroupDeleteAction sets GroupDeleteAction field to given value.

### HasGroupDeleteAction

`func (o *MicrosoftEntraProviderRequest) HasGroupDeleteAction() bool`

HasGroupDeleteAction returns a boolean if a field has been set.

### GetSyncPageSize

`func (o *MicrosoftEntraProviderRequest) GetSyncPageSize() int32`

GetSyncPageSize returns the SyncPageSize field if non-nil, zero value otherwise.

### GetSyncPageSizeOk

`func (o *MicrosoftEntraProviderRequest) GetSyncPageSizeOk() (*int32, bool)`

GetSyncPageSizeOk returns a tuple with the SyncPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPageSize

`func (o *MicrosoftEntraProviderRequest) SetSyncPageSize(v int32)`

SetSyncPageSize sets SyncPageSize field to given value.

### HasSyncPageSize

`func (o *MicrosoftEntraProviderRequest) HasSyncPageSize() bool`

HasSyncPageSize returns a boolean if a field has been set.

### GetSyncPageTimeout

`func (o *MicrosoftEntraProviderRequest) GetSyncPageTimeout() string`

GetSyncPageTimeout returns the SyncPageTimeout field if non-nil, zero value otherwise.

### GetSyncPageTimeoutOk

`func (o *MicrosoftEntraProviderRequest) GetSyncPageTimeoutOk() (*string, bool)`

GetSyncPageTimeoutOk returns a tuple with the SyncPageTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPageTimeout

`func (o *MicrosoftEntraProviderRequest) SetSyncPageTimeout(v string)`

SetSyncPageTimeout sets SyncPageTimeout field to given value.

### HasSyncPageTimeout

`func (o *MicrosoftEntraProviderRequest) HasSyncPageTimeout() bool`

HasSyncPageTimeout returns a boolean if a field has been set.

### GetDryRun

`func (o *MicrosoftEntraProviderRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *MicrosoftEntraProviderRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *MicrosoftEntraProviderRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *MicrosoftEntraProviderRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


