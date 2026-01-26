# PatchedSCIMProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**PropertyMappingsGroup** | Pointer to **[]string** | Property mappings used for group creation/updating. | [optional] 
**Url** | Pointer to **string** | Base URL to SCIM requests, usually ends in /v2 | [optional] 
**VerifyCertificates** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** | Authentication token | [optional] 
**AuthMode** | Pointer to [**SCIMAuthenticationModeEnum**](SCIMAuthenticationModeEnum.md) |  | [optional] 
**AuthOauth** | Pointer to **NullableString** | OAuth Source used for authentication | [optional] 
**AuthOauthParams** | Pointer to **map[string]interface{}** | Additional OAuth parameters, such as grant_type | [optional] 
**CompatibilityMode** | Pointer to [**CompatibilityModeEnum**](CompatibilityModeEnum.md) | Alter authentik behavior for vendor-specific SCIM implementations. | [optional] 
**ExcludeUsersServiceAccount** | Pointer to **bool** |  | [optional] 
**FilterGroup** | Pointer to **NullableString** |  | [optional] 
**DryRun** | Pointer to **bool** | When enabled, provider will not modify or create objects in the remote system. | [optional] 

## Methods

### NewPatchedSCIMProviderRequest

`func NewPatchedSCIMProviderRequest() *PatchedSCIMProviderRequest`

NewPatchedSCIMProviderRequest instantiates a new PatchedSCIMProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSCIMProviderRequestWithDefaults

`func NewPatchedSCIMProviderRequestWithDefaults() *PatchedSCIMProviderRequest`

NewPatchedSCIMProviderRequestWithDefaults instantiates a new PatchedSCIMProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedSCIMProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSCIMProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSCIMProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSCIMProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPropertyMappings

`func (o *PatchedSCIMProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *PatchedSCIMProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *PatchedSCIMProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *PatchedSCIMProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetPropertyMappingsGroup

`func (o *PatchedSCIMProviderRequest) GetPropertyMappingsGroup() []string`

GetPropertyMappingsGroup returns the PropertyMappingsGroup field if non-nil, zero value otherwise.

### GetPropertyMappingsGroupOk

`func (o *PatchedSCIMProviderRequest) GetPropertyMappingsGroupOk() (*[]string, bool)`

GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappingsGroup

`func (o *PatchedSCIMProviderRequest) SetPropertyMappingsGroup(v []string)`

SetPropertyMappingsGroup sets PropertyMappingsGroup field to given value.

### HasPropertyMappingsGroup

`func (o *PatchedSCIMProviderRequest) HasPropertyMappingsGroup() bool`

HasPropertyMappingsGroup returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedSCIMProviderRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedSCIMProviderRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedSCIMProviderRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedSCIMProviderRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVerifyCertificates

`func (o *PatchedSCIMProviderRequest) GetVerifyCertificates() bool`

GetVerifyCertificates returns the VerifyCertificates field if non-nil, zero value otherwise.

### GetVerifyCertificatesOk

`func (o *PatchedSCIMProviderRequest) GetVerifyCertificatesOk() (*bool, bool)`

GetVerifyCertificatesOk returns a tuple with the VerifyCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificates

`func (o *PatchedSCIMProviderRequest) SetVerifyCertificates(v bool)`

SetVerifyCertificates sets VerifyCertificates field to given value.

### HasVerifyCertificates

`func (o *PatchedSCIMProviderRequest) HasVerifyCertificates() bool`

HasVerifyCertificates returns a boolean if a field has been set.

### GetToken

`func (o *PatchedSCIMProviderRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PatchedSCIMProviderRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PatchedSCIMProviderRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PatchedSCIMProviderRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAuthMode

`func (o *PatchedSCIMProviderRequest) GetAuthMode() SCIMAuthenticationModeEnum`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *PatchedSCIMProviderRequest) GetAuthModeOk() (*SCIMAuthenticationModeEnum, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *PatchedSCIMProviderRequest) SetAuthMode(v SCIMAuthenticationModeEnum)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *PatchedSCIMProviderRequest) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetAuthOauth

`func (o *PatchedSCIMProviderRequest) GetAuthOauth() string`

GetAuthOauth returns the AuthOauth field if non-nil, zero value otherwise.

### GetAuthOauthOk

`func (o *PatchedSCIMProviderRequest) GetAuthOauthOk() (*string, bool)`

GetAuthOauthOk returns a tuple with the AuthOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOauth

`func (o *PatchedSCIMProviderRequest) SetAuthOauth(v string)`

SetAuthOauth sets AuthOauth field to given value.

### HasAuthOauth

`func (o *PatchedSCIMProviderRequest) HasAuthOauth() bool`

HasAuthOauth returns a boolean if a field has been set.

### SetAuthOauthNil

`func (o *PatchedSCIMProviderRequest) SetAuthOauthNil(b bool)`

 SetAuthOauthNil sets the value for AuthOauth to be an explicit nil

### UnsetAuthOauth
`func (o *PatchedSCIMProviderRequest) UnsetAuthOauth()`

UnsetAuthOauth ensures that no value is present for AuthOauth, not even an explicit nil
### GetAuthOauthParams

`func (o *PatchedSCIMProviderRequest) GetAuthOauthParams() map[string]interface{}`

GetAuthOauthParams returns the AuthOauthParams field if non-nil, zero value otherwise.

### GetAuthOauthParamsOk

`func (o *PatchedSCIMProviderRequest) GetAuthOauthParamsOk() (*map[string]interface{}, bool)`

GetAuthOauthParamsOk returns a tuple with the AuthOauthParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOauthParams

`func (o *PatchedSCIMProviderRequest) SetAuthOauthParams(v map[string]interface{})`

SetAuthOauthParams sets AuthOauthParams field to given value.

### HasAuthOauthParams

`func (o *PatchedSCIMProviderRequest) HasAuthOauthParams() bool`

HasAuthOauthParams returns a boolean if a field has been set.

### GetCompatibilityMode

`func (o *PatchedSCIMProviderRequest) GetCompatibilityMode() CompatibilityModeEnum`

GetCompatibilityMode returns the CompatibilityMode field if non-nil, zero value otherwise.

### GetCompatibilityModeOk

`func (o *PatchedSCIMProviderRequest) GetCompatibilityModeOk() (*CompatibilityModeEnum, bool)`

GetCompatibilityModeOk returns a tuple with the CompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityMode

`func (o *PatchedSCIMProviderRequest) SetCompatibilityMode(v CompatibilityModeEnum)`

SetCompatibilityMode sets CompatibilityMode field to given value.

### HasCompatibilityMode

`func (o *PatchedSCIMProviderRequest) HasCompatibilityMode() bool`

HasCompatibilityMode returns a boolean if a field has been set.

### GetExcludeUsersServiceAccount

`func (o *PatchedSCIMProviderRequest) GetExcludeUsersServiceAccount() bool`

GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field if non-nil, zero value otherwise.

### GetExcludeUsersServiceAccountOk

`func (o *PatchedSCIMProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool)`

GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsersServiceAccount

`func (o *PatchedSCIMProviderRequest) SetExcludeUsersServiceAccount(v bool)`

SetExcludeUsersServiceAccount sets ExcludeUsersServiceAccount field to given value.

### HasExcludeUsersServiceAccount

`func (o *PatchedSCIMProviderRequest) HasExcludeUsersServiceAccount() bool`

HasExcludeUsersServiceAccount returns a boolean if a field has been set.

### GetFilterGroup

`func (o *PatchedSCIMProviderRequest) GetFilterGroup() string`

GetFilterGroup returns the FilterGroup field if non-nil, zero value otherwise.

### GetFilterGroupOk

`func (o *PatchedSCIMProviderRequest) GetFilterGroupOk() (*string, bool)`

GetFilterGroupOk returns a tuple with the FilterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroup

`func (o *PatchedSCIMProviderRequest) SetFilterGroup(v string)`

SetFilterGroup sets FilterGroup field to given value.

### HasFilterGroup

`func (o *PatchedSCIMProviderRequest) HasFilterGroup() bool`

HasFilterGroup returns a boolean if a field has been set.

### SetFilterGroupNil

`func (o *PatchedSCIMProviderRequest) SetFilterGroupNil(b bool)`

 SetFilterGroupNil sets the value for FilterGroup to be an explicit nil

### UnsetFilterGroup
`func (o *PatchedSCIMProviderRequest) UnsetFilterGroup()`

UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
### GetDryRun

`func (o *PatchedSCIMProviderRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *PatchedSCIMProviderRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *PatchedSCIMProviderRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *PatchedSCIMProviderRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


