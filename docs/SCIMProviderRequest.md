# SCIMProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**PropertyMappingsGroup** | Pointer to **[]string** | Property mappings used for group creation/updating. | [optional] 
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

### NewSCIMProviderRequest

`func NewSCIMProviderRequest(name string, url string, ) *SCIMProviderRequest`

NewSCIMProviderRequest instantiates a new SCIMProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCIMProviderRequestWithDefaults

`func NewSCIMProviderRequestWithDefaults() *SCIMProviderRequest`

NewSCIMProviderRequestWithDefaults instantiates a new SCIMProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SCIMProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SCIMProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SCIMProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPropertyMappings

`func (o *SCIMProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *SCIMProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *SCIMProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *SCIMProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetPropertyMappingsGroup

`func (o *SCIMProviderRequest) GetPropertyMappingsGroup() []string`

GetPropertyMappingsGroup returns the PropertyMappingsGroup field if non-nil, zero value otherwise.

### GetPropertyMappingsGroupOk

`func (o *SCIMProviderRequest) GetPropertyMappingsGroupOk() (*[]string, bool)`

GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappingsGroup

`func (o *SCIMProviderRequest) SetPropertyMappingsGroup(v []string)`

SetPropertyMappingsGroup sets PropertyMappingsGroup field to given value.

### HasPropertyMappingsGroup

`func (o *SCIMProviderRequest) HasPropertyMappingsGroup() bool`

HasPropertyMappingsGroup returns a boolean if a field has been set.

### GetUrl

`func (o *SCIMProviderRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SCIMProviderRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SCIMProviderRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVerifyCertificates

`func (o *SCIMProviderRequest) GetVerifyCertificates() bool`

GetVerifyCertificates returns the VerifyCertificates field if non-nil, zero value otherwise.

### GetVerifyCertificatesOk

`func (o *SCIMProviderRequest) GetVerifyCertificatesOk() (*bool, bool)`

GetVerifyCertificatesOk returns a tuple with the VerifyCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificates

`func (o *SCIMProviderRequest) SetVerifyCertificates(v bool)`

SetVerifyCertificates sets VerifyCertificates field to given value.

### HasVerifyCertificates

`func (o *SCIMProviderRequest) HasVerifyCertificates() bool`

HasVerifyCertificates returns a boolean if a field has been set.

### GetToken

`func (o *SCIMProviderRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SCIMProviderRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SCIMProviderRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SCIMProviderRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAuthMode

`func (o *SCIMProviderRequest) GetAuthMode() SCIMAuthenticationModeEnum`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SCIMProviderRequest) GetAuthModeOk() (*SCIMAuthenticationModeEnum, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SCIMProviderRequest) SetAuthMode(v SCIMAuthenticationModeEnum)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *SCIMProviderRequest) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthOauth

`func (o *SCIMProviderRequest) GetAuthOauth() string`

GetAuthOauth returns the AuthOauth field if non-nil, zero value otherwise.

### GetAuthOauthOk

`func (o *SCIMProviderRequest) GetAuthOauthOk() (*string, bool)`

GetAuthOauthOk returns a tuple with the AuthOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOauth

`func (o *SCIMProviderRequest) SetAuthOauth(v string)`

SetAuthOauth sets AuthOauth field to given value.

### HasAuthOauth

`func (o *SCIMProviderRequest) HasAuthOauth() bool`

HasAuthOauth returns a boolean if a field has been set.

### SetAuthOauthNil

`func (o *SCIMProviderRequest) SetAuthOauthNil(b bool)`

 SetAuthOauthNil sets the value for AuthOauth to be an explicit nil

### UnsetAuthOauth
`func (o *SCIMProviderRequest) UnsetAuthOauth()`

UnsetAuthOauth ensures that no value is present for AuthOauth, not even an explicit nil
### GetAuthOauthParams

`func (o *SCIMProviderRequest) GetAuthOauthParams() map[string]interface{}`

GetAuthOauthParams returns the AuthOauthParams field if non-nil, zero value otherwise.

### GetAuthOauthParamsOk

`func (o *SCIMProviderRequest) GetAuthOauthParamsOk() (*map[string]interface{}, bool)`

GetAuthOauthParamsOk returns a tuple with the AuthOauthParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOauthParams

`func (o *SCIMProviderRequest) SetAuthOauthParams(v map[string]interface{})`

SetAuthOauthParams sets AuthOauthParams field to given value.

### HasAuthOauthParams

`func (o *SCIMProviderRequest) HasAuthOauthParams() bool`

HasAuthOauthParams returns a boolean if a field has been set.

### GetCompatibilityMode

`func (o *SCIMProviderRequest) GetCompatibilityMode() CompatibilityModeEnum`

GetCompatibilityMode returns the CompatibilityMode field if non-nil, zero value otherwise.

### GetCompatibilityModeOk

`func (o *SCIMProviderRequest) GetCompatibilityModeOk() (*CompatibilityModeEnum, bool)`

GetCompatibilityModeOk returns a tuple with the CompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityMode

`func (o *SCIMProviderRequest) SetCompatibilityMode(v CompatibilityModeEnum)`

SetCompatibilityMode sets CompatibilityMode field to given value.

### HasCompatibilityMode

`func (o *SCIMProviderRequest) HasCompatibilityMode() bool`

HasCompatibilityMode returns a boolean if a field has been set.

### GetExcludeUsersServiceAccount

`func (o *SCIMProviderRequest) GetExcludeUsersServiceAccount() bool`

GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field if non-nil, zero value otherwise.

### GetExcludeUsersServiceAccountOk

`func (o *SCIMProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool)`

GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsersServiceAccount

`func (o *SCIMProviderRequest) SetExcludeUsersServiceAccount(v bool)`

SetExcludeUsersServiceAccount sets ExcludeUsersServiceAccount field to given value.

### HasExcludeUsersServiceAccount

`func (o *SCIMProviderRequest) HasExcludeUsersServiceAccount() bool`

HasExcludeUsersServiceAccount returns a boolean if a field has been set.

### GetFilterGroup

`func (o *SCIMProviderRequest) GetFilterGroup() string`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *SCIMProviderRequest) GetFilterGroupOk() (*string, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *SCIMProviderRequest) SetFilterGroup(v string)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *SCIMProviderRequest) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### SetFilterGroupNil

`func (o *SCIMProviderRequest) SetFilterGroupNil(b bool)`

 SetFilterGroupNil sets the value for FilterGroup to be an explicit nil

### UnsetFilterGroup
`func (o *SCIMProviderRequest) UnsetFilterGroup()`

UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
### GetSyncPageSize

`func (o *SCIMProviderRequest) GetSyncPageSize() int32`

GetSyncPageSize returns the SyncPageSize field if non-nil, zero value otherwise.

### GetSyncPageSizeOk

`func (o *SCIMProviderRequest) GetSyncPageSizeOk() (*int32, bool)`

GetSyncPageSizeOk returns a tuple with the SyncPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPageSize

`func (o *SCIMProviderRequest) SetSyncPageSize(v int32)`

SetSyncPageSize sets SyncPageSize field to given value.

### HasSyncPageSize

`func (o *SCIMProviderRequest) HasSyncPageSize() bool`

HasSyncPageSize returns a boolean if a field has been set.

### GetSyncPageTimeout

`func (o *SCIMProviderRequest) GetSyncPageTimeout() string`

GetSyncPageTimeout returns the SyncPageTimeout field if non-nil, zero value otherwise.

### GetSyncPageTimeoutOk

`func (o *SCIMProviderRequest) GetSyncPageTimeoutOk() (*string, bool)`

GetSyncPageTimeoutOk returns a tuple with the SyncPageTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPageTimeout

`func (o *SCIMProviderRequest) SetSyncPageTimeout(v string)`

SetSyncPageTimeout sets SyncPageTimeout field to given value.

### HasSyncPageTimeout

`func (o *SCIMProviderRequest) HasSyncPageTimeout() bool`

HasSyncPageTimeout returns a boolean if a field has been set.

### GetDryRun

`func (o *SCIMProviderRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *SCIMProviderRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *SCIMProviderRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *SCIMProviderRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


