# PatchedLDAPProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | Pointer to **string** | Flow used when authorizing this provider. | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**BaseDn** | Pointer to **string** | DN under which objects are accessible. | [optional] 
**SearchGroup** | Pointer to **NullableString** | Users in this group can do search queries. If not set, every user can execute search queries. | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**TlsServerName** | Pointer to **string** |  | [optional] 
**UidStartNumber** | Pointer to **int32** | The start for uidNumbers, this number is added to the user.pk to make sure that the numbers aren&#39;t too low for POSIX users. Default is 2000 to ensure that we don&#39;t collide with local users uidNumber | [optional] 
**GidStartNumber** | Pointer to **int32** | The start for gidNumbers, this number is added to a number generated from the group.pk to make sure that the numbers aren&#39;t too low for POSIX groups. Default is 4000 to ensure that we don&#39;t collide with local groups or users primary groups gidNumber | [optional] 
**SearchMode** | Pointer to [**LDAPAPIAccessMode**](LDAPAPIAccessMode.md) |  | [optional] 
**BindMode** | Pointer to [**LDAPAPIAccessMode**](LDAPAPIAccessMode.md) |  | [optional] 
**MfaSupport** | Pointer to **bool** | When enabled, code-based multi-factor authentication can be used by appending a semicolon and the TOTP code to the password. This should only be enabled if all users that will bind to this provider have a TOTP device configured, as otherwise a password may incorrectly be rejected if it contains a semicolon. | [optional] 

## Methods

### NewPatchedLDAPProviderRequest

`func NewPatchedLDAPProviderRequest() *PatchedLDAPProviderRequest`

NewPatchedLDAPProviderRequest instantiates a new PatchedLDAPProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLDAPProviderRequestWithDefaults

`func NewPatchedLDAPProviderRequestWithDefaults() *PatchedLDAPProviderRequest`

NewPatchedLDAPProviderRequestWithDefaults instantiates a new PatchedLDAPProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedLDAPProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedLDAPProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedLDAPProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedLDAPProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PatchedLDAPProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedLDAPProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedLDAPProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedLDAPProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedLDAPProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedLDAPProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *PatchedLDAPProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *PatchedLDAPProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *PatchedLDAPProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.

### HasAuthorizationFlow

`func (o *PatchedLDAPProviderRequest) HasAuthorizationFlow() bool`

HasAuthorizationFlow returns a boolean if a field has been set.

### GetPropertyMappings

`func (o *PatchedLDAPProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *PatchedLDAPProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *PatchedLDAPProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *PatchedLDAPProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetBaseDn

`func (o *PatchedLDAPProviderRequest) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *PatchedLDAPProviderRequest) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *PatchedLDAPProviderRequest) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *PatchedLDAPProviderRequest) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetSearchGroup

`func (o *PatchedLDAPProviderRequest) GetSearchGroup() string`

GetSearchGroup returns the SearchGroup field if non-nil, zero value otherwise.

### GetSearchGroupOk

`func (o *PatchedLDAPProviderRequest) GetSearchGroupOk() (*string, bool)`

GetSearchGroupOk returns a tuple with the SearchGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchGroup

`func (o *PatchedLDAPProviderRequest) SetSearchGroup(v string)`

SetSearchGroup sets SearchGroup field to given value.

### HasSearchGroup

`func (o *PatchedLDAPProviderRequest) HasSearchGroup() bool`

HasSearchGroup returns a boolean if a field has been set.

### SetSearchGroupNil

`func (o *PatchedLDAPProviderRequest) SetSearchGroupNil(b bool)`

 SetSearchGroupNil sets the value for SearchGroup to be an explicit nil

### UnsetSearchGroup
`func (o *PatchedLDAPProviderRequest) UnsetSearchGroup()`

UnsetSearchGroup ensures that no value is present for SearchGroup, not even an explicit nil
### GetCertificate

`func (o *PatchedLDAPProviderRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *PatchedLDAPProviderRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *PatchedLDAPProviderRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *PatchedLDAPProviderRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *PatchedLDAPProviderRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *PatchedLDAPProviderRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetTlsServerName

`func (o *PatchedLDAPProviderRequest) GetTlsServerName() string`

GetTlsServerName returns the TlsServerName field if non-nil, zero value otherwise.

### GetTlsServerNameOk

`func (o *PatchedLDAPProviderRequest) GetTlsServerNameOk() (*string, bool)`

GetTlsServerNameOk returns a tuple with the TlsServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerName

`func (o *PatchedLDAPProviderRequest) SetTlsServerName(v string)`

SetTlsServerName sets TlsServerName field to given value.

### HasTlsServerName

`func (o *PatchedLDAPProviderRequest) HasTlsServerName() bool`

HasTlsServerName returns a boolean if a field has been set.

### GetUidStartNumber

`func (o *PatchedLDAPProviderRequest) GetUidStartNumber() int32`

GetUidStartNumber returns the UidStartNumber field if non-nil, zero value otherwise.

### GetUidStartNumberOk

`func (o *PatchedLDAPProviderRequest) GetUidStartNumberOk() (*int32, bool)`

GetUidStartNumberOk returns a tuple with the UidStartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidStartNumber

`func (o *PatchedLDAPProviderRequest) SetUidStartNumber(v int32)`

SetUidStartNumber sets UidStartNumber field to given value.

### HasUidStartNumber

`func (o *PatchedLDAPProviderRequest) HasUidStartNumber() bool`

HasUidStartNumber returns a boolean if a field has been set.

### GetGidStartNumber

`func (o *PatchedLDAPProviderRequest) GetGidStartNumber() int32`

GetGidStartNumber returns the GidStartNumber field if non-nil, zero value otherwise.

### GetGidStartNumberOk

`func (o *PatchedLDAPProviderRequest) GetGidStartNumberOk() (*int32, bool)`

GetGidStartNumberOk returns a tuple with the GidStartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidStartNumber

`func (o *PatchedLDAPProviderRequest) SetGidStartNumber(v int32)`

SetGidStartNumber sets GidStartNumber field to given value.

### HasGidStartNumber

`func (o *PatchedLDAPProviderRequest) HasGidStartNumber() bool`

HasGidStartNumber returns a boolean if a field has been set.

### GetSearchMode

`func (o *PatchedLDAPProviderRequest) GetSearchMode() LDAPAPIAccessMode`

GetSearchMode returns the SearchMode field if non-nil, zero value otherwise.

### GetSearchModeOk

`func (o *PatchedLDAPProviderRequest) GetSearchModeOk() (*LDAPAPIAccessMode, bool)`

GetSearchModeOk returns a tuple with the SearchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMode

`func (o *PatchedLDAPProviderRequest) SetSearchMode(v LDAPAPIAccessMode)`

SetSearchMode sets SearchMode field to given value.

### HasSearchMode

`func (o *PatchedLDAPProviderRequest) HasSearchMode() bool`

HasSearchMode returns a boolean if a field has been set.

### GetBindMode

`func (o *PatchedLDAPProviderRequest) GetBindMode() LDAPAPIAccessMode`

GetBindMode returns the BindMode field if non-nil, zero value otherwise.

### GetBindModeOk

`func (o *PatchedLDAPProviderRequest) GetBindModeOk() (*LDAPAPIAccessMode, bool)`

GetBindModeOk returns a tuple with the BindMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMode

`func (o *PatchedLDAPProviderRequest) SetBindMode(v LDAPAPIAccessMode)`

SetBindMode sets BindMode field to given value.

### HasBindMode

`func (o *PatchedLDAPProviderRequest) HasBindMode() bool`

HasBindMode returns a boolean if a field has been set.

### GetMfaSupport

`func (o *PatchedLDAPProviderRequest) GetMfaSupport() bool`

GetMfaSupport returns the MfaSupport field if non-nil, zero value otherwise.

### GetMfaSupportOk

`func (o *PatchedLDAPProviderRequest) GetMfaSupportOk() (*bool, bool)`

GetMfaSupportOk returns a tuple with the MfaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaSupport

`func (o *PatchedLDAPProviderRequest) SetMfaSupport(v bool)`

SetMfaSupport sets MfaSupport field to given value.

### HasMfaSupport

`func (o *PatchedLDAPProviderRequest) HasMfaSupport() bool`

HasMfaSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


