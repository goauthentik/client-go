# LDAPProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
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

### NewLDAPProviderRequest

`func NewLDAPProviderRequest(name string, authorizationFlow string, ) *LDAPProviderRequest`

NewLDAPProviderRequest instantiates a new LDAPProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPProviderRequestWithDefaults

`func NewLDAPProviderRequestWithDefaults() *LDAPProviderRequest`

NewLDAPProviderRequestWithDefaults instantiates a new LDAPProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LDAPProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticationFlow

`func (o *LDAPProviderRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *LDAPProviderRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *LDAPProviderRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *LDAPProviderRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *LDAPProviderRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *LDAPProviderRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *LDAPProviderRequest) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *LDAPProviderRequest) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *LDAPProviderRequest) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetPropertyMappings

`func (o *LDAPProviderRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *LDAPProviderRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *LDAPProviderRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *LDAPProviderRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetBaseDn

`func (o *LDAPProviderRequest) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LDAPProviderRequest) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LDAPProviderRequest) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *LDAPProviderRequest) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetSearchGroup

`func (o *LDAPProviderRequest) GetSearchGroup() string`

GetSearchGroup returns the SearchGroup field if non-nil, zero value otherwise.

### GetSearchGroupOk

`func (o *LDAPProviderRequest) GetSearchGroupOk() (*string, bool)`

GetSearchGroupOk returns a tuple with the SearchGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchGroup

`func (o *LDAPProviderRequest) SetSearchGroup(v string)`

SetSearchGroup sets SearchGroup field to given value.

### HasSearchGroup

`func (o *LDAPProviderRequest) HasSearchGroup() bool`

HasSearchGroup returns a boolean if a field has been set.

### SetSearchGroupNil

`func (o *LDAPProviderRequest) SetSearchGroupNil(b bool)`

 SetSearchGroupNil sets the value for SearchGroup to be an explicit nil

### UnsetSearchGroup
`func (o *LDAPProviderRequest) UnsetSearchGroup()`

UnsetSearchGroup ensures that no value is present for SearchGroup, not even an explicit nil
### GetCertificate

`func (o *LDAPProviderRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LDAPProviderRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LDAPProviderRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *LDAPProviderRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *LDAPProviderRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *LDAPProviderRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetTlsServerName

`func (o *LDAPProviderRequest) GetTlsServerName() string`

GetTlsServerName returns the TlsServerName field if non-nil, zero value otherwise.

### GetTlsServerNameOk

`func (o *LDAPProviderRequest) GetTlsServerNameOk() (*string, bool)`

GetTlsServerNameOk returns a tuple with the TlsServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerName

`func (o *LDAPProviderRequest) SetTlsServerName(v string)`

SetTlsServerName sets TlsServerName field to given value.

### HasTlsServerName

`func (o *LDAPProviderRequest) HasTlsServerName() bool`

HasTlsServerName returns a boolean if a field has been set.

### GetUidStartNumber

`func (o *LDAPProviderRequest) GetUidStartNumber() int32`

GetUidStartNumber returns the UidStartNumber field if non-nil, zero value otherwise.

### GetUidStartNumberOk

`func (o *LDAPProviderRequest) GetUidStartNumberOk() (*int32, bool)`

GetUidStartNumberOk returns a tuple with the UidStartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidStartNumber

`func (o *LDAPProviderRequest) SetUidStartNumber(v int32)`

SetUidStartNumber sets UidStartNumber field to given value.

### HasUidStartNumber

`func (o *LDAPProviderRequest) HasUidStartNumber() bool`

HasUidStartNumber returns a boolean if a field has been set.

### GetGidStartNumber

`func (o *LDAPProviderRequest) GetGidStartNumber() int32`

GetGidStartNumber returns the GidStartNumber field if non-nil, zero value otherwise.

### GetGidStartNumberOk

`func (o *LDAPProviderRequest) GetGidStartNumberOk() (*int32, bool)`

GetGidStartNumberOk returns a tuple with the GidStartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidStartNumber

`func (o *LDAPProviderRequest) SetGidStartNumber(v int32)`

SetGidStartNumber sets GidStartNumber field to given value.

### HasGidStartNumber

`func (o *LDAPProviderRequest) HasGidStartNumber() bool`

HasGidStartNumber returns a boolean if a field has been set.

### GetSearchMode

`func (o *LDAPProviderRequest) GetSearchMode() LDAPAPIAccessMode`

GetSearchMode returns the SearchMode field if non-nil, zero value otherwise.

### GetSearchModeOk

`func (o *LDAPProviderRequest) GetSearchModeOk() (*LDAPAPIAccessMode, bool)`

GetSearchModeOk returns a tuple with the SearchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMode

`func (o *LDAPProviderRequest) SetSearchMode(v LDAPAPIAccessMode)`

SetSearchMode sets SearchMode field to given value.

### HasSearchMode

`func (o *LDAPProviderRequest) HasSearchMode() bool`

HasSearchMode returns a boolean if a field has been set.

### GetBindMode

`func (o *LDAPProviderRequest) GetBindMode() LDAPAPIAccessMode`

GetBindMode returns the BindMode field if non-nil, zero value otherwise.

### GetBindModeOk

`func (o *LDAPProviderRequest) GetBindModeOk() (*LDAPAPIAccessMode, bool)`

GetBindModeOk returns a tuple with the BindMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMode

`func (o *LDAPProviderRequest) SetBindMode(v LDAPAPIAccessMode)`

SetBindMode sets BindMode field to given value.

### HasBindMode

`func (o *LDAPProviderRequest) HasBindMode() bool`

HasBindMode returns a boolean if a field has been set.

### GetMfaSupport

`func (o *LDAPProviderRequest) GetMfaSupport() bool`

GetMfaSupport returns the MfaSupport field if non-nil, zero value otherwise.

### GetMfaSupportOk

`func (o *LDAPProviderRequest) GetMfaSupportOk() (*bool, bool)`

GetMfaSupportOk returns a tuple with the MfaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaSupport

`func (o *LDAPProviderRequest) SetMfaSupport(v bool)`

SetMfaSupport sets MfaSupport field to given value.

### HasMfaSupport

`func (o *LDAPProviderRequest) HasMfaSupport() bool`

HasMfaSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


