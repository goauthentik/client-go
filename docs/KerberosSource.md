# KerberosSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** | Source&#39;s display Name. | 
**Slug** | **string** | Internal source name, used in URLs. | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Promoted** | Pointer to **bool** | When enabled, this source will be displayed as a prominent button on the login page, instead of a small icon. | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow to use when authenticating existing users. | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Flow to use when enrolling new users. | [optional] 
**UserPropertyMappings** | Pointer to **[]string** |  | [optional] 
**GroupPropertyMappings** | Pointer to **[]string** |  | [optional] 
**Component** | **string** | Get object component so that we know how to edit the object | [readonly] 
**VerboseName** | **string** | Return object&#39;s verbose_name | [readonly] 
**VerboseNamePlural** | **string** | Return object&#39;s plural verbose_name | [readonly] 
**MetaModelName** | **string** | Return internal model name | [readonly] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**UserMatchingMode** | Pointer to [**UserMatchingModeEnum**](UserMatchingModeEnum.md) | How the source determines if an existing user should be authenticated or a new user enrolled. | [optional] 
**Managed** | **NullableString** | Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update. | [readonly] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**IconUrl** | **string** |  | [readonly] 
**GroupMatchingMode** | Pointer to [**GroupMatchingModeEnum**](GroupMatchingModeEnum.md) | How the source determines if an existing group should be used or a new group created. | [optional] 
**Realm** | **string** | Kerberos realm | 
**Krb5Conf** | Pointer to **string** | Custom krb5.conf to use. Uses the system one by default | [optional] 
**KadminType** | Pointer to [**KadminTypeEnum**](KadminTypeEnum.md) | KAdmin server type | [optional] 
**SyncUsers** | Pointer to **bool** | Sync users from Kerberos into authentik | [optional] 
**SyncUsersPassword** | Pointer to **bool** | When a user changes their password, sync it back to Kerberos | [optional] 
**SyncPrincipal** | Pointer to **string** | Principal to authenticate to kadmin for sync. | [optional] 
**SyncCcache** | Pointer to **string** | Credentials cache to authenticate to kadmin for sync. Must be in the form TYPE:residual | [optional] 
**Connectivity** | **map[string]string** | Get cached source connectivity | [readonly] 
**SpnegoServerName** | Pointer to **string** | Force the use of a specific server name for SPNEGO. Must be in the form HTTP@hostname | [optional] 
**SpnegoCcache** | Pointer to **string** | Credential cache to use for SPNEGO in form type:residual | [optional] 
**PasswordLoginUpdateInternalPassword** | Pointer to **bool** | If enabled, the authentik-stored password will be updated upon login with the Kerberos password backend | [optional] 

## Methods

### NewKerberosSource

`func NewKerberosSource(pk string, name string, slug string, component string, verboseName string, verboseNamePlural string, metaModelName string, managed NullableString, iconUrl string, realm string, connectivity map[string]string, ) *KerberosSource`

NewKerberosSource instantiates a new KerberosSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosSourceWithDefaults

`func NewKerberosSourceWithDefaults() *KerberosSource`

NewKerberosSourceWithDefaults instantiates a new KerberosSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *KerberosSource) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *KerberosSource) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *KerberosSource) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *KerberosSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KerberosSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KerberosSource) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *KerberosSource) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *KerberosSource) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *KerberosSource) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *KerberosSource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KerberosSource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KerberosSource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KerberosSource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPromoted

`func (o *KerberosSource) GetPromoted() bool`

GetPromoted returns the Promoted field if non-nil, zero value otherwise.

### GetPromotedOk

`func (o *KerberosSource) GetPromotedOk() (*bool, bool)`

GetPromotedOk returns a tuple with the Promoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoted

`func (o *KerberosSource) SetPromoted(v bool)`

SetPromoted sets Promoted field to given value.

### HasPromoted

`func (o *KerberosSource) HasPromoted() bool`

HasPromoted returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *KerberosSource) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *KerberosSource) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *KerberosSource) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *KerberosSource) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *KerberosSource) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *KerberosSource) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *KerberosSource) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *KerberosSource) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *KerberosSource) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *KerberosSource) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *KerberosSource) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *KerberosSource) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetUserPropertyMappings

`func (o *KerberosSource) GetUserPropertyMappings() []string`

GetUserPropertyMappings returns the UserPropertyMappings field if non-nil, zero value otherwise.

### GetUserPropertyMappingsOk

`func (o *KerberosSource) GetUserPropertyMappingsOk() (*[]string, bool)`

GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertyMappings

`func (o *KerberosSource) SetUserPropertyMappings(v []string)`

SetUserPropertyMappings sets UserPropertyMappings field to given value.

### HasUserPropertyMappings

`func (o *KerberosSource) HasUserPropertyMappings() bool`

HasUserPropertyMappings returns a boolean if a field has been set.

### GetGroupPropertyMappings

`func (o *KerberosSource) GetGroupPropertyMappings() []string`

GetGroupPropertyMappings returns the GroupPropertyMappings field if non-nil, zero value otherwise.

### GetGroupPropertyMappingsOk

`func (o *KerberosSource) GetGroupPropertyMappingsOk() (*[]string, bool)`

GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPropertyMappings

`func (o *KerberosSource) SetGroupPropertyMappings(v []string)`

SetGroupPropertyMappings sets GroupPropertyMappings field to given value.

### HasGroupPropertyMappings

`func (o *KerberosSource) HasGroupPropertyMappings() bool`

HasGroupPropertyMappings returns a boolean if a field has been set.

### GetComponent

`func (o *KerberosSource) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *KerberosSource) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *KerberosSource) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *KerberosSource) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *KerberosSource) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *KerberosSource) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *KerberosSource) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *KerberosSource) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *KerberosSource) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *KerberosSource) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *KerberosSource) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *KerberosSource) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetPolicyEngineMode

`func (o *KerberosSource) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *KerberosSource) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *KerberosSource) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *KerberosSource) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *KerberosSource) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *KerberosSource) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *KerberosSource) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *KerberosSource) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetManaged

`func (o *KerberosSource) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *KerberosSource) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *KerberosSource) SetManaged(v string)`

SetManaged sets Managed field to given value.


### SetManagedNil

`func (o *KerberosSource) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *KerberosSource) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetUserPathTemplate

`func (o *KerberosSource) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *KerberosSource) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *KerberosSource) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *KerberosSource) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetIcon

`func (o *KerberosSource) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *KerberosSource) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *KerberosSource) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *KerberosSource) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetIconUrl

`func (o *KerberosSource) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *KerberosSource) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *KerberosSource) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### GetGroupMatchingMode

`func (o *KerberosSource) GetGroupMatchingMode() GroupMatchingModeEnum`

GetGroupMatchingMode returns the GroupMatchingMode field if non-nil, zero value otherwise.

### GetGroupMatchingModeOk

`func (o *KerberosSource) GetGroupMatchingModeOk() (*GroupMatchingModeEnum, bool)`

GetGroupMatchingModeOk returns a tuple with the GroupMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMatchingMode

`func (o *KerberosSource) SetGroupMatchingMode(v GroupMatchingModeEnum)`

SetGroupMatchingMode sets GroupMatchingMode field to given value.

### HasGroupMatchingMode

`func (o *KerberosSource) HasGroupMatchingMode() bool`

HasGroupMatchingMode returns a boolean if a field has been set.

### GetRealm

`func (o *KerberosSource) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *KerberosSource) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *KerberosSource) SetRealm(v string)`

SetRealm sets Realm field to given value.


### GetKrb5Conf

`func (o *KerberosSource) GetKrb5Conf() string`

GetKrb5Conf returns the Krb5Conf field if non-nil, zero value otherwise.

### GetKrb5ConfOk

`func (o *KerberosSource) GetKrb5ConfOk() (*string, bool)`

GetKrb5ConfOk returns a tuple with the Krb5Conf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrb5Conf

`func (o *KerberosSource) SetKrb5Conf(v string)`

SetKrb5Conf sets Krb5Conf field to given value.

### HasKrb5Conf

`func (o *KerberosSource) HasKrb5Conf() bool`

HasKrb5Conf returns a boolean if a field has been set.

### GetKadminType

`func (o *KerberosSource) GetKadminType() KadminTypeEnum`

GetKadminType returns the KadminType field if non-nil, zero value otherwise.

### GetKadminTypeOk

`func (o *KerberosSource) GetKadminTypeOk() (*KadminTypeEnum, bool)`

GetKadminTypeOk returns a tuple with the KadminType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKadminType

`func (o *KerberosSource) SetKadminType(v KadminTypeEnum)`

SetKadminType sets KadminType field to given value.

### HasKadminType

`func (o *KerberosSource) HasKadminType() bool`

HasKadminType returns a boolean if a field has been set.

### GetSyncUsers

`func (o *KerberosSource) GetSyncUsers() bool`

GetSyncUsers returns the SyncUsers field if non-nil, zero value otherwise.

### GetSyncUsersOk

`func (o *KerberosSource) GetSyncUsersOk() (*bool, bool)`

GetSyncUsersOk returns a tuple with the SyncUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsers

`func (o *KerberosSource) SetSyncUsers(v bool)`

SetSyncUsers sets SyncUsers field to given value.

### HasSyncUsers

`func (o *KerberosSource) HasSyncUsers() bool`

HasSyncUsers returns a boolean if a field has been set.

### GetSyncUsersPassword

`func (o *KerberosSource) GetSyncUsersPassword() bool`

GetSyncUsersPassword returns the SyncUsersPassword field if non-nil, zero value otherwise.

### GetSyncUsersPasswordOk

`func (o *KerberosSource) GetSyncUsersPasswordOk() (*bool, bool)`

GetSyncUsersPasswordOk returns a tuple with the SyncUsersPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsersPassword

`func (o *KerberosSource) SetSyncUsersPassword(v bool)`

SetSyncUsersPassword sets SyncUsersPassword field to given value.

### HasSyncUsersPassword

`func (o *KerberosSource) HasSyncUsersPassword() bool`

HasSyncUsersPassword returns a boolean if a field has been set.

### GetSyncPrincipal

`func (o *KerberosSource) GetSyncPrincipal() string`

GetSyncPrincipal returns the SyncPrincipal field if non-nil, zero value otherwise.

### GetSyncPrincipalOk

`func (o *KerberosSource) GetSyncPrincipalOk() (*string, bool)`

GetSyncPrincipalOk returns a tuple with the SyncPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPrincipal

`func (o *KerberosSource) SetSyncPrincipal(v string)`

SetSyncPrincipal sets SyncPrincipal field to given value.

### HasSyncPrincipal

`func (o *KerberosSource) HasSyncPrincipal() bool`

HasSyncPrincipal returns a boolean if a field has been set.

### GetSyncCcache

`func (o *KerberosSource) GetSyncCcache() string`

GetSyncCcache returns the SyncCcache field if non-nil, zero value otherwise.

### GetSyncCcacheOk

`func (o *KerberosSource) GetSyncCcacheOk() (*string, bool)`

GetSyncCcacheOk returns a tuple with the SyncCcache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCcache

`func (o *KerberosSource) SetSyncCcache(v string)`

SetSyncCcache sets SyncCcache field to given value.

### HasSyncCcache

`func (o *KerberosSource) HasSyncCcache() bool`

HasSyncCcache returns a boolean if a field has been set.

### GetConnectivity

`func (o *KerberosSource) GetConnectivity() map[string]string`

GetConnectivity returns the Connectivity field if non-nil, zero value otherwise.

### GetConnectivityOk

`func (o *KerberosSource) GetConnectivityOk() (*map[string]string, bool)`

GetConnectivityOk returns a tuple with the Connectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivity

`func (o *KerberosSource) SetConnectivity(v map[string]string)`

SetConnectivity sets Connectivity field to given value.


### SetConnectivityNil

`func (o *KerberosSource) SetConnectivityNil(b bool)`

 SetConnectivityNil sets the value for Connectivity to be an explicit nil

### UnsetConnectivity
`func (o *KerberosSource) UnsetConnectivity()`

UnsetConnectivity ensures that no value is present for Connectivity, not even an explicit nil
### GetSpnegoServerName

`func (o *KerberosSource) GetSpnegoServerName() string`

GetSpnegoServerName returns the SpnegoServerName field if non-nil, zero value otherwise.

### GetSpnegoServerNameOk

`func (o *KerberosSource) GetSpnegoServerNameOk() (*string, bool)`

GetSpnegoServerNameOk returns a tuple with the SpnegoServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpnegoServerName

`func (o *KerberosSource) SetSpnegoServerName(v string)`

SetSpnegoServerName sets SpnegoServerName field to given value.

### HasSpnegoServerName

`func (o *KerberosSource) HasSpnegoServerName() bool`

HasSpnegoServerName returns a boolean if a field has been set.

### GetSpnegoCcache

`func (o *KerberosSource) GetSpnegoCcache() string`

GetSpnegoCcache returns the SpnegoCcache field if non-nil, zero value otherwise.

### GetSpnegoCcacheOk

`func (o *KerberosSource) GetSpnegoCcacheOk() (*string, bool)`

GetSpnegoCcacheOk returns a tuple with the SpnegoCcache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpnegoCcache

`func (o *KerberosSource) SetSpnegoCcache(v string)`

SetSpnegoCcache sets SpnegoCcache field to given value.

### HasSpnegoCcache

`func (o *KerberosSource) HasSpnegoCcache() bool`

HasSpnegoCcache returns a boolean if a field has been set.

### GetPasswordLoginUpdateInternalPassword

`func (o *KerberosSource) GetPasswordLoginUpdateInternalPassword() bool`

GetPasswordLoginUpdateInternalPassword returns the PasswordLoginUpdateInternalPassword field if non-nil, zero value otherwise.

### GetPasswordLoginUpdateInternalPasswordOk

`func (o *KerberosSource) GetPasswordLoginUpdateInternalPasswordOk() (*bool, bool)`

GetPasswordLoginUpdateInternalPasswordOk returns a tuple with the PasswordLoginUpdateInternalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLoginUpdateInternalPassword

`func (o *KerberosSource) SetPasswordLoginUpdateInternalPassword(v bool)`

SetPasswordLoginUpdateInternalPassword sets PasswordLoginUpdateInternalPassword field to given value.

### HasPasswordLoginUpdateInternalPassword

`func (o *KerberosSource) HasPasswordLoginUpdateInternalPassword() bool`

HasPasswordLoginUpdateInternalPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


