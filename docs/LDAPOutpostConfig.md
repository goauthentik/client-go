# LDAPOutpostConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**BaseDn** | Pointer to **string** | DN under which objects are accessible. | [optional] 
**BindFlowSlug** | **string** |  | 
**ApplicationSlug** | **string** | Prioritise backchannel slug over direct application slug | [readonly] 
**SearchGroup** | Pointer to **NullableString** | Users in this group can do search queries. If not set, every user can execute search queries. | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**TlsServerName** | Pointer to **string** |  | [optional] 
**UidStartNumber** | Pointer to **int32** | The start for uidNumbers, this number is added to the user.pk to make sure that the numbers aren&#39;t too low for POSIX users. Default is 2000 to ensure that we don&#39;t collide with local users uidNumber | [optional] 
**GidStartNumber** | Pointer to **int32** | The start for gidNumbers, this number is added to a number generated from the group.pk to make sure that the numbers aren&#39;t too low for POSIX groups. Default is 4000 to ensure that we don&#39;t collide with local groups or users primary groups gidNumber | [optional] 
**SearchMode** | Pointer to [**LDAPAPIAccessMode**](LDAPAPIAccessMode.md) |  | [optional] 
**BindMode** | Pointer to [**LDAPAPIAccessMode**](LDAPAPIAccessMode.md) |  | [optional] 
**MfaSupport** | Pointer to **bool** | When enabled, code-based multi-factor authentication can be used by appending a semicolon and the TOTP code to the password. This should only be enabled if all users that will bind to this provider have a TOTP device configured, as otherwise a password may incorrectly be rejected if it contains a semicolon. | [optional] 

## Methods

### NewLDAPOutpostConfig

`func NewLDAPOutpostConfig(pk int32, name string, bindFlowSlug string, applicationSlug string, ) *LDAPOutpostConfig`

NewLDAPOutpostConfig instantiates a new LDAPOutpostConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPOutpostConfigWithDefaults

`func NewLDAPOutpostConfigWithDefaults() *LDAPOutpostConfig`

NewLDAPOutpostConfigWithDefaults instantiates a new LDAPOutpostConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *LDAPOutpostConfig) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *LDAPOutpostConfig) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *LDAPOutpostConfig) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *LDAPOutpostConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPOutpostConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPOutpostConfig) SetName(v string)`

SetName sets Name field to given value.


### GetBaseDn

`func (o *LDAPOutpostConfig) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LDAPOutpostConfig) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LDAPOutpostConfig) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *LDAPOutpostConfig) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetBindFlowSlug

`func (o *LDAPOutpostConfig) GetBindFlowSlug() string`

GetBindFlowSlug returns the BindFlowSlug field if non-nil, zero value otherwise.

### GetBindFlowSlugOk

`func (o *LDAPOutpostConfig) GetBindFlowSlugOk() (*string, bool)`

GetBindFlowSlugOk returns a tuple with the BindFlowSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindFlowSlug

`func (o *LDAPOutpostConfig) SetBindFlowSlug(v string)`

SetBindFlowSlug sets BindFlowSlug field to given value.


### GetApplicationSlug

`func (o *LDAPOutpostConfig) GetApplicationSlug() string`

GetApplicationSlug returns the ApplicationSlug field if non-nil, zero value otherwise.

### GetApplicationSlugOk

`func (o *LDAPOutpostConfig) GetApplicationSlugOk() (*string, bool)`

GetApplicationSlugOk returns a tuple with the ApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSlug

`func (o *LDAPOutpostConfig) SetApplicationSlug(v string)`

SetApplicationSlug sets ApplicationSlug field to given value.


### GetSearchGroup

`func (o *LDAPOutpostConfig) GetSearchGroup() string`

GetSearchGroup returns the SearchGroup field if non-nil, zero value otherwise.

### GetSearchGroupOk

`func (o *LDAPOutpostConfig) GetSearchGroupOk() (*string, bool)`

GetSearchGroupOk returns a tuple with the SearchGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchGroup

`func (o *LDAPOutpostConfig) SetSearchGroup(v string)`

SetSearchGroup sets SearchGroup field to given value.

### HasSearchGroup

`func (o *LDAPOutpostConfig) HasSearchGroup() bool`

HasSearchGroup returns a boolean if a field has been set.

### SetSearchGroupNil

`func (o *LDAPOutpostConfig) SetSearchGroupNil(b bool)`

 SetSearchGroupNil sets the value for SearchGroup to be an explicit nil

### UnsetSearchGroup
`func (o *LDAPOutpostConfig) UnsetSearchGroup()`

UnsetSearchGroup ensures that no value is present for SearchGroup, not even an explicit nil
### GetCertificate

`func (o *LDAPOutpostConfig) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LDAPOutpostConfig) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LDAPOutpostConfig) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *LDAPOutpostConfig) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *LDAPOutpostConfig) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *LDAPOutpostConfig) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetTlsServerName

`func (o *LDAPOutpostConfig) GetTlsServerName() string`

GetTlsServerName returns the TlsServerName field if non-nil, zero value otherwise.

### GetTlsServerNameOk

`func (o *LDAPOutpostConfig) GetTlsServerNameOk() (*string, bool)`

GetTlsServerNameOk returns a tuple with the TlsServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerName

`func (o *LDAPOutpostConfig) SetTlsServerName(v string)`

SetTlsServerName sets TlsServerName field to given value.

### HasTlsServerName

`func (o *LDAPOutpostConfig) HasTlsServerName() bool`

HasTlsServerName returns a boolean if a field has been set.

### GetUidStartNumber

`func (o *LDAPOutpostConfig) GetUidStartNumber() int32`

GetUidStartNumber returns the UidStartNumber field if non-nil, zero value otherwise.

### GetUidStartNumberOk

`func (o *LDAPOutpostConfig) GetUidStartNumberOk() (*int32, bool)`

GetUidStartNumberOk returns a tuple with the UidStartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidStartNumber

`func (o *LDAPOutpostConfig) SetUidStartNumber(v int32)`

SetUidStartNumber sets UidStartNumber field to given value.

### HasUidStartNumber

`func (o *LDAPOutpostConfig) HasUidStartNumber() bool`

HasUidStartNumber returns a boolean if a field has been set.

### GetGidStartNumber

`func (o *LDAPOutpostConfig) GetGidStartNumber() int32`

GetGidStartNumber returns the GidStartNumber field if non-nil, zero value otherwise.

### GetGidStartNumberOk

`func (o *LDAPOutpostConfig) GetGidStartNumberOk() (*int32, bool)`

GetGidStartNumberOk returns a tuple with the GidStartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidStartNumber

`func (o *LDAPOutpostConfig) SetGidStartNumber(v int32)`

SetGidStartNumber sets GidStartNumber field to given value.

### HasGidStartNumber

`func (o *LDAPOutpostConfig) HasGidStartNumber() bool`

HasGidStartNumber returns a boolean if a field has been set.

### GetSearchMode

`func (o *LDAPOutpostConfig) GetSearchMode() LDAPAPIAccessMode`

GetSearchMode returns the SearchMode field if non-nil, zero value otherwise.

### GetSearchModeOk

`func (o *LDAPOutpostConfig) GetSearchModeOk() (*LDAPAPIAccessMode, bool)`

GetSearchModeOk returns a tuple with the SearchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMode

`func (o *LDAPOutpostConfig) SetSearchMode(v LDAPAPIAccessMode)`

SetSearchMode sets SearchMode field to given value.

### HasSearchMode

`func (o *LDAPOutpostConfig) HasSearchMode() bool`

HasSearchMode returns a boolean if a field has been set.

### GetBindMode

`func (o *LDAPOutpostConfig) GetBindMode() LDAPAPIAccessMode`

GetBindMode returns the BindMode field if non-nil, zero value otherwise.

### GetBindModeOk

`func (o *LDAPOutpostConfig) GetBindModeOk() (*LDAPAPIAccessMode, bool)`

GetBindModeOk returns a tuple with the BindMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMode

`func (o *LDAPOutpostConfig) SetBindMode(v LDAPAPIAccessMode)`

SetBindMode sets BindMode field to given value.

### HasBindMode

`func (o *LDAPOutpostConfig) HasBindMode() bool`

HasBindMode returns a boolean if a field has been set.

### GetMfaSupport

`func (o *LDAPOutpostConfig) GetMfaSupport() bool`

GetMfaSupport returns the MfaSupport field if non-nil, zero value otherwise.

### GetMfaSupportOk

`func (o *LDAPOutpostConfig) GetMfaSupportOk() (*bool, bool)`

GetMfaSupportOk returns a tuple with the MfaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaSupport

`func (o *LDAPOutpostConfig) SetMfaSupport(v bool)`

SetMfaSupport sets MfaSupport field to given value.

### HasMfaSupport

`func (o *LDAPOutpostConfig) HasMfaSupport() bool`

HasMfaSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


