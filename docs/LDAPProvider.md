# LDAPProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**AuthenticationFlow** | Pointer to **NullableString** | Flow used for authentication when the associated application is accessed by an un-authenticated user. | [optional] 
**AuthorizationFlow** | **string** | Flow used when authorizing this provider. | 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**AssignedApplicationSlug** | **string** | Internal application name, used in URLs. | [readonly] 
**AssignedApplicationName** | **string** | Application&#39;s display Name. | [readonly] 
**AssignedBackchannelApplicationSlug** | **string** | Internal application name, used in URLs. | [readonly] 
**AssignedBackchannelApplicationName** | **string** | Application&#39;s display Name. | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**BaseDn** | Pointer to **string** | DN under which objects are accessible. | [optional] 
**SearchGroup** | Pointer to **NullableString** | Users in this group can do search queries. If not set, every user can execute search queries. | [optional] 
**Certificate** | Pointer to **NullableString** |  | [optional] 
**TlsServerName** | Pointer to **string** |  | [optional] 
**UidStartNumber** | Pointer to **int32** | The start for uidNumbers, this number is added to the user.pk to make sure that the numbers aren&#39;t too low for POSIX users. Default is 2000 to ensure that we don&#39;t collide with local users uidNumber | [optional] 
**GidStartNumber** | Pointer to **int32** | The start for gidNumbers, this number is added to a number generated from the group.pk to make sure that the numbers aren&#39;t too low for POSIX groups. Default is 4000 to ensure that we don&#39;t collide with local groups or users primary groups gidNumber | [optional] 
**OutpostSet** | **[]string** |  | [readonly] 
**SearchMode** | Pointer to [**LDAPAPIAccessMode**](LDAPAPIAccessMode.md) |  | [optional] 
**BindMode** | Pointer to [**LDAPAPIAccessMode**](LDAPAPIAccessMode.md) |  | [optional] 
**MfaSupport** | Pointer to **bool** | When enabled, code-based multi-factor authentication can be used by appending a semicolon and the TOTP code to the password. This should only be enabled if all users that will bind to this provider have a TOTP device configured, as otherwise a password may incorrectly be rejected if it contains a semicolon. | [optional] 

## Methods

### NewLDAPProvider

`func NewLDAPProvider(pk int32, name string, authorizationFlow string, component string, assignedApplicationSlug string, assignedApplicationName string, assignedBackchannelApplicationSlug string, assignedBackchannelApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, outpostSet []string, ) *LDAPProvider`

NewLDAPProvider instantiates a new LDAPProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPProviderWithDefaults

`func NewLDAPProviderWithDefaults() *LDAPProvider`

NewLDAPProviderWithDefaults instantiates a new LDAPProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *LDAPProvider) GetPk() int32`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *LDAPProvider) GetPkOk() (*int32, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *LDAPProvider) SetPk(v int32)`

SetPk sets Pk field to given value.


### GetName

`func (o *LDAPProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPProvider) SetName(v string)`

SetName sets Name field to given value.


### GetAuthenticationFlow

`func (o *LDAPProvider) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *LDAPProvider) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *LDAPProvider) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *LDAPProvider) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *LDAPProvider) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *LDAPProvider) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetAuthorizationFlow

`func (o *LDAPProvider) GetAuthorizationFlow() string`

GetAuthorizationFlow returns the AuthorizationFlow field if non-nil, zero value otherwise.

### GetAuthorizationFlowOk

`func (o *LDAPProvider) GetAuthorizationFlowOk() (*string, bool)`

GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationFlow

`func (o *LDAPProvider) SetAuthorizationFlow(v string)`

SetAuthorizationFlow sets AuthorizationFlow field to given value.


### GetPropertyMappings

`func (o *LDAPProvider) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *LDAPProvider) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *LDAPProvider) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *LDAPProvider) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetComponent

`func (o *LDAPProvider) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *LDAPProvider) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *LDAPProvider) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetAssignedApplicationSlug

`func (o *LDAPProvider) GetAssignedApplicationSlug() string`

GetAssignedApplicationSlug returns the AssignedApplicationSlug field if non-nil, zero value otherwise.

### GetAssignedApplicationSlugOk

`func (o *LDAPProvider) GetAssignedApplicationSlugOk() (*string, bool)`

GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedApplicationSlug

`func (o *LDAPProvider) SetAssignedApplicationSlug(v string)`

SetAssignedApplicationSlug sets AssignedApplicationSlug field to given value.


### GetAssignedApplicationName

`func (o *LDAPProvider) GetAssignedApplicationName() string`

GetAssignedApplicationName returns the AssignedApplicationName field if non-nil, zero value otherwise.

### GetAssignedApplicationNameOk

`func (o *LDAPProvider) GetAssignedApplicationNameOk() (*string, bool)`

GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedApplicationName

`func (o *LDAPProvider) SetAssignedApplicationName(v string)`

SetAssignedApplicationName sets AssignedApplicationName field to given value.


### GetAssignedBackchannelApplicationSlug

`func (o *LDAPProvider) GetAssignedBackchannelApplicationSlug() string`

GetAssignedBackchannelApplicationSlug returns the AssignedBackchannelApplicationSlug field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationSlugOk

`func (o *LDAPProvider) GetAssignedBackchannelApplicationSlugOk() (*string, bool)`

GetAssignedBackchannelApplicationSlugOk returns a tuple with the AssignedBackchannelApplicationSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationSlug

`func (o *LDAPProvider) SetAssignedBackchannelApplicationSlug(v string)`

SetAssignedBackchannelApplicationSlug sets AssignedBackchannelApplicationSlug field to given value.


### GetAssignedBackchannelApplicationName

`func (o *LDAPProvider) GetAssignedBackchannelApplicationName() string`

GetAssignedBackchannelApplicationName returns the AssignedBackchannelApplicationName field if non-nil, zero value otherwise.

### GetAssignedBackchannelApplicationNameOk

`func (o *LDAPProvider) GetAssignedBackchannelApplicationNameOk() (*string, bool)`

GetAssignedBackchannelApplicationNameOk returns a tuple with the AssignedBackchannelApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedBackchannelApplicationName

`func (o *LDAPProvider) SetAssignedBackchannelApplicationName(v string)`

SetAssignedBackchannelApplicationName sets AssignedBackchannelApplicationName field to given value.


### GetVerboseName

`func (o *LDAPProvider) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *LDAPProvider) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *LDAPProvider) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *LDAPProvider) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *LDAPProvider) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *LDAPProvider) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *LDAPProvider) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *LDAPProvider) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *LDAPProvider) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetBaseDn

`func (o *LDAPProvider) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LDAPProvider) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LDAPProvider) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *LDAPProvider) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetSearchGroup

`func (o *LDAPProvider) GetSearchGroup() string`

GetSearchGroup returns the SearchGroup field if non-nil, zero value otherwise.

### GetSearchGroupOk

`func (o *LDAPProvider) GetSearchGroupOk() (*string, bool)`

GetSearchGroupOk returns a tuple with the SearchGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchGroup

`func (o *LDAPProvider) SetSearchGroup(v string)`

SetSearchGroup sets SearchGroup field to given value.

### HasSearchGroup

`func (o *LDAPProvider) HasSearchGroup() bool`

HasSearchGroup returns a boolean if a field has been set.

### SetSearchGroupNil

`func (o *LDAPProvider) SetSearchGroupNil(b bool)`

 SetSearchGroupNil sets the value for SearchGroup to be an explicit nil

### UnsetSearchGroup
`func (o *LDAPProvider) UnsetSearchGroup()`

UnsetSearchGroup ensures that no value is present for SearchGroup, not even an explicit nil
### GetCertificate

`func (o *LDAPProvider) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LDAPProvider) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LDAPProvider) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *LDAPProvider) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *LDAPProvider) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *LDAPProvider) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetTlsServerName

`func (o *LDAPProvider) GetTlsServerName() string`

GetTlsServerName returns the TlsServerName field if non-nil, zero value otherwise.

### GetTlsServerNameOk

`func (o *LDAPProvider) GetTlsServerNameOk() (*string, bool)`

GetTlsServerNameOk returns a tuple with the TlsServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerName

`func (o *LDAPProvider) SetTlsServerName(v string)`

SetTlsServerName sets TlsServerName field to given value.

### HasTlsServerName

`func (o *LDAPProvider) HasTlsServerName() bool`

HasTlsServerName returns a boolean if a field has been set.

### GetUidStartNumber

`func (o *LDAPProvider) GetUidStartNumber() int32`

GetUidStartNumber returns the UidStartNumber field if non-nil, zero value otherwise.

### GetUidStartNumberOk

`func (o *LDAPProvider) GetUidStartNumberOk() (*int32, bool)`

GetUidStartNumberOk returns a tuple with the UidStartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidStartNumber

`func (o *LDAPProvider) SetUidStartNumber(v int32)`

SetUidStartNumber sets UidStartNumber field to given value.

### HasUidStartNumber

`func (o *LDAPProvider) HasUidStartNumber() bool`

HasUidStartNumber returns a boolean if a field has been set.

### GetGidStartNumber

`func (o *LDAPProvider) GetGidStartNumber() int32`

GetGidStartNumber returns the GidStartNumber field if non-nil, zero value otherwise.

### GetGidStartNumberOk

`func (o *LDAPProvider) GetGidStartNumberOk() (*int32, bool)`

GetGidStartNumberOk returns a tuple with the GidStartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGidStartNumber

`func (o *LDAPProvider) SetGidStartNumber(v int32)`

SetGidStartNumber sets GidStartNumber field to given value.

### HasGidStartNumber

`func (o *LDAPProvider) HasGidStartNumber() bool`

HasGidStartNumber returns a boolean if a field has been set.

### GetOutpostSet

`func (o *LDAPProvider) GetOutpostSet() []string`

GetOutpostSet returns the OutpostSet field if non-nil, zero value otherwise.

### GetOutpostSetOk

`func (o *LDAPProvider) GetOutpostSetOk() (*[]string, bool)`

GetOutpostSetOk returns a tuple with the OutpostSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutpostSet

`func (o *LDAPProvider) SetOutpostSet(v []string)`

SetOutpostSet sets OutpostSet field to given value.


### GetSearchMode

`func (o *LDAPProvider) GetSearchMode() LDAPAPIAccessMode`

GetSearchMode returns the SearchMode field if non-nil, zero value otherwise.

### GetSearchModeOk

`func (o *LDAPProvider) GetSearchModeOk() (*LDAPAPIAccessMode, bool)`

GetSearchModeOk returns a tuple with the SearchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMode

`func (o *LDAPProvider) SetSearchMode(v LDAPAPIAccessMode)`

SetSearchMode sets SearchMode field to given value.

### HasSearchMode

`func (o *LDAPProvider) HasSearchMode() bool`

HasSearchMode returns a boolean if a field has been set.

### GetBindMode

`func (o *LDAPProvider) GetBindMode() LDAPAPIAccessMode`

GetBindMode returns the BindMode field if non-nil, zero value otherwise.

### GetBindModeOk

`func (o *LDAPProvider) GetBindModeOk() (*LDAPAPIAccessMode, bool)`

GetBindModeOk returns a tuple with the BindMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMode

`func (o *LDAPProvider) SetBindMode(v LDAPAPIAccessMode)`

SetBindMode sets BindMode field to given value.

### HasBindMode

`func (o *LDAPProvider) HasBindMode() bool`

HasBindMode returns a boolean if a field has been set.

### GetMfaSupport

`func (o *LDAPProvider) GetMfaSupport() bool`

GetMfaSupport returns the MfaSupport field if non-nil, zero value otherwise.

### GetMfaSupportOk

`func (o *LDAPProvider) GetMfaSupportOk() (*bool, bool)`

GetMfaSupportOk returns a tuple with the MfaSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaSupport

`func (o *LDAPProvider) SetMfaSupport(v bool)`

SetMfaSupport sets MfaSupport field to given value.

### HasMfaSupport

`func (o *LDAPProvider) HasMfaSupport() bool`

HasMfaSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


