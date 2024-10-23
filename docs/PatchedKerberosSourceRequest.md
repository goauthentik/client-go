# PatchedKerberosSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Source&#39;s display Name. | [optional] 
**Slug** | Pointer to **string** | Internal source name, used in URLs. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow to use when authenticating existing users. | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Flow to use when enrolling new users. | [optional] 
**UserPropertyMappings** | Pointer to **[]string** |  | [optional] 
**GroupPropertyMappings** | Pointer to **[]string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**UserMatchingMode** | Pointer to [**UserMatchingModeEnum**](UserMatchingModeEnum.md) | How the source determines if an existing user should be authenticated or a new user enrolled. | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 
**GroupMatchingMode** | Pointer to [**GroupMatchingModeEnum**](GroupMatchingModeEnum.md) | How the source determines if an existing group should be used or a new group created. | [optional] 
**Realm** | Pointer to **string** | Kerberos realm | [optional] 
**Krb5Conf** | Pointer to **string** | Custom krb5.conf to use. Uses the system one by default | [optional] 
**SyncUsers** | Pointer to **bool** | Sync users from Kerberos into authentik | [optional] 
**SyncUsersPassword** | Pointer to **bool** | When a user changes their password, sync it back to Kerberos | [optional] 
**SyncPrincipal** | Pointer to **string** | Principal to authenticate to kadmin for sync. | [optional] 
**SyncPassword** | Pointer to **string** | Password to authenticate to kadmin for sync | [optional] 
**SyncKeytab** | Pointer to **string** | Keytab to authenticate to kadmin for sync. Must be base64-encoded or in the form TYPE:residual | [optional] 
**SyncCcache** | Pointer to **string** | Credentials cache to authenticate to kadmin for sync. Must be in the form TYPE:residual | [optional] 
**SpnegoServerName** | Pointer to **string** | Force the use of a specific server name for SPNEGO | [optional] 
**SpnegoKeytab** | Pointer to **string** | SPNEGO keytab base64-encoded or path to keytab in the form FILE:path | [optional] 
**SpnegoCcache** | Pointer to **string** | Credential cache to use for SPNEGO in form type:residual | [optional] 
**PasswordLoginUpdateInternalPassword** | Pointer to **bool** | If enabled, the authentik-stored password will be updated upon login with the Kerberos password backend | [optional] 

## Methods

### NewPatchedKerberosSourceRequest

`func NewPatchedKerberosSourceRequest() *PatchedKerberosSourceRequest`

NewPatchedKerberosSourceRequest instantiates a new PatchedKerberosSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedKerberosSourceRequestWithDefaults

`func NewPatchedKerberosSourceRequestWithDefaults() *PatchedKerberosSourceRequest`

NewPatchedKerberosSourceRequestWithDefaults instantiates a new PatchedKerberosSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedKerberosSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedKerberosSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedKerberosSourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedKerberosSourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedKerberosSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedKerberosSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedKerberosSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedKerberosSourceRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedKerberosSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedKerberosSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedKerberosSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedKerberosSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PatchedKerberosSourceRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedKerberosSourceRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedKerberosSourceRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedKerberosSourceRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedKerberosSourceRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedKerberosSourceRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *PatchedKerberosSourceRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *PatchedKerberosSourceRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *PatchedKerberosSourceRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *PatchedKerberosSourceRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *PatchedKerberosSourceRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *PatchedKerberosSourceRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetUserPropertyMappings

`func (o *PatchedKerberosSourceRequest) GetUserPropertyMappings() []string`

GetUserPropertyMappings returns the UserPropertyMappings field if non-nil, zero value otherwise.

### GetUserPropertyMappingsOk

`func (o *PatchedKerberosSourceRequest) GetUserPropertyMappingsOk() (*[]string, bool)`

GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertyMappings

`func (o *PatchedKerberosSourceRequest) SetUserPropertyMappings(v []string)`

SetUserPropertyMappings sets UserPropertyMappings field to given value.

### HasUserPropertyMappings

`func (o *PatchedKerberosSourceRequest) HasUserPropertyMappings() bool`

HasUserPropertyMappings returns a boolean if a field has been set.

### GetGroupPropertyMappings

`func (o *PatchedKerberosSourceRequest) GetGroupPropertyMappings() []string`

GetGroupPropertyMappings returns the GroupPropertyMappings field if non-nil, zero value otherwise.

### GetGroupPropertyMappingsOk

`func (o *PatchedKerberosSourceRequest) GetGroupPropertyMappingsOk() (*[]string, bool)`

GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPropertyMappings

`func (o *PatchedKerberosSourceRequest) SetGroupPropertyMappings(v []string)`

SetGroupPropertyMappings sets GroupPropertyMappings field to given value.

### HasGroupPropertyMappings

`func (o *PatchedKerberosSourceRequest) HasGroupPropertyMappings() bool`

HasGroupPropertyMappings returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *PatchedKerberosSourceRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PatchedKerberosSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PatchedKerberosSourceRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PatchedKerberosSourceRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *PatchedKerberosSourceRequest) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *PatchedKerberosSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *PatchedKerberosSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *PatchedKerberosSourceRequest) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *PatchedKerberosSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *PatchedKerberosSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *PatchedKerberosSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *PatchedKerberosSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetGroupMatchingMode

`func (o *PatchedKerberosSourceRequest) GetGroupMatchingMode() GroupMatchingModeEnum`

GetGroupMatchingMode returns the GroupMatchingMode field if non-nil, zero value otherwise.

### GetGroupMatchingModeOk

`func (o *PatchedKerberosSourceRequest) GetGroupMatchingModeOk() (*GroupMatchingModeEnum, bool)`

GetGroupMatchingModeOk returns a tuple with the GroupMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMatchingMode

`func (o *PatchedKerberosSourceRequest) SetGroupMatchingMode(v GroupMatchingModeEnum)`

SetGroupMatchingMode sets GroupMatchingMode field to given value.

### HasGroupMatchingMode

`func (o *PatchedKerberosSourceRequest) HasGroupMatchingMode() bool`

HasGroupMatchingMode returns a boolean if a field has been set.

### GetRealm

`func (o *PatchedKerberosSourceRequest) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *PatchedKerberosSourceRequest) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *PatchedKerberosSourceRequest) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *PatchedKerberosSourceRequest) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetKrb5Conf

`func (o *PatchedKerberosSourceRequest) GetKrb5Conf() string`

GetKrb5Conf returns the Krb5Conf field if non-nil, zero value otherwise.

### GetKrb5ConfOk

`func (o *PatchedKerberosSourceRequest) GetKrb5ConfOk() (*string, bool)`

GetKrb5ConfOk returns a tuple with the Krb5Conf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrb5Conf

`func (o *PatchedKerberosSourceRequest) SetKrb5Conf(v string)`

SetKrb5Conf sets Krb5Conf field to given value.

### HasKrb5Conf

`func (o *PatchedKerberosSourceRequest) HasKrb5Conf() bool`

HasKrb5Conf returns a boolean if a field has been set.

### GetSyncUsers

`func (o *PatchedKerberosSourceRequest) GetSyncUsers() bool`

GetSyncUsers returns the SyncUsers field if non-nil, zero value otherwise.

### GetSyncUsersOk

`func (o *PatchedKerberosSourceRequest) GetSyncUsersOk() (*bool, bool)`

GetSyncUsersOk returns a tuple with the SyncUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsers

`func (o *PatchedKerberosSourceRequest) SetSyncUsers(v bool)`

SetSyncUsers sets SyncUsers field to given value.

### HasSyncUsers

`func (o *PatchedKerberosSourceRequest) HasSyncUsers() bool`

HasSyncUsers returns a boolean if a field has been set.

### GetSyncUsersPassword

`func (o *PatchedKerberosSourceRequest) GetSyncUsersPassword() bool`

GetSyncUsersPassword returns the SyncUsersPassword field if non-nil, zero value otherwise.

### GetSyncUsersPasswordOk

`func (o *PatchedKerberosSourceRequest) GetSyncUsersPasswordOk() (*bool, bool)`

GetSyncUsersPasswordOk returns a tuple with the SyncUsersPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsersPassword

`func (o *PatchedKerberosSourceRequest) SetSyncUsersPassword(v bool)`

SetSyncUsersPassword sets SyncUsersPassword field to given value.

### HasSyncUsersPassword

`func (o *PatchedKerberosSourceRequest) HasSyncUsersPassword() bool`

HasSyncUsersPassword returns a boolean if a field has been set.

### GetSyncPrincipal

`func (o *PatchedKerberosSourceRequest) GetSyncPrincipal() string`

GetSyncPrincipal returns the SyncPrincipal field if non-nil, zero value otherwise.

### GetSyncPrincipalOk

`func (o *PatchedKerberosSourceRequest) GetSyncPrincipalOk() (*string, bool)`

GetSyncPrincipalOk returns a tuple with the SyncPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPrincipal

`func (o *PatchedKerberosSourceRequest) SetSyncPrincipal(v string)`

SetSyncPrincipal sets SyncPrincipal field to given value.

### HasSyncPrincipal

`func (o *PatchedKerberosSourceRequest) HasSyncPrincipal() bool`

HasSyncPrincipal returns a boolean if a field has been set.

### GetSyncPassword

`func (o *PatchedKerberosSourceRequest) GetSyncPassword() string`

GetSyncPassword returns the SyncPassword field if non-nil, zero value otherwise.

### GetSyncPasswordOk

`func (o *PatchedKerberosSourceRequest) GetSyncPasswordOk() (*string, bool)`

GetSyncPasswordOk returns a tuple with the SyncPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPassword

`func (o *PatchedKerberosSourceRequest) SetSyncPassword(v string)`

SetSyncPassword sets SyncPassword field to given value.

### HasSyncPassword

`func (o *PatchedKerberosSourceRequest) HasSyncPassword() bool`

HasSyncPassword returns a boolean if a field has been set.

### GetSyncKeytab

`func (o *PatchedKerberosSourceRequest) GetSyncKeytab() string`

GetSyncKeytab returns the SyncKeytab field if non-nil, zero value otherwise.

### GetSyncKeytabOk

`func (o *PatchedKerberosSourceRequest) GetSyncKeytabOk() (*string, bool)`

GetSyncKeytabOk returns a tuple with the SyncKeytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncKeytab

`func (o *PatchedKerberosSourceRequest) SetSyncKeytab(v string)`

SetSyncKeytab sets SyncKeytab field to given value.

### HasSyncKeytab

`func (o *PatchedKerberosSourceRequest) HasSyncKeytab() bool`

HasSyncKeytab returns a boolean if a field has been set.

### GetSyncCcache

`func (o *PatchedKerberosSourceRequest) GetSyncCcache() string`

GetSyncCcache returns the SyncCcache field if non-nil, zero value otherwise.

### GetSyncCcacheOk

`func (o *PatchedKerberosSourceRequest) GetSyncCcacheOk() (*string, bool)`

GetSyncCcacheOk returns a tuple with the SyncCcache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCcache

`func (o *PatchedKerberosSourceRequest) SetSyncCcache(v string)`

SetSyncCcache sets SyncCcache field to given value.

### HasSyncCcache

`func (o *PatchedKerberosSourceRequest) HasSyncCcache() bool`

HasSyncCcache returns a boolean if a field has been set.

### GetSpnegoServerName

`func (o *PatchedKerberosSourceRequest) GetSpnegoServerName() string`

GetSpnegoServerName returns the SpnegoServerName field if non-nil, zero value otherwise.

### GetSpnegoServerNameOk

`func (o *PatchedKerberosSourceRequest) GetSpnegoServerNameOk() (*string, bool)`

GetSpnegoServerNameOk returns a tuple with the SpnegoServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpnegoServerName

`func (o *PatchedKerberosSourceRequest) SetSpnegoServerName(v string)`

SetSpnegoServerName sets SpnegoServerName field to given value.

### HasSpnegoServerName

`func (o *PatchedKerberosSourceRequest) HasSpnegoServerName() bool`

HasSpnegoServerName returns a boolean if a field has been set.

### GetSpnegoKeytab

`func (o *PatchedKerberosSourceRequest) GetSpnegoKeytab() string`

GetSpnegoKeytab returns the SpnegoKeytab field if non-nil, zero value otherwise.

### GetSpnegoKeytabOk

`func (o *PatchedKerberosSourceRequest) GetSpnegoKeytabOk() (*string, bool)`

GetSpnegoKeytabOk returns a tuple with the SpnegoKeytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpnegoKeytab

`func (o *PatchedKerberosSourceRequest) SetSpnegoKeytab(v string)`

SetSpnegoKeytab sets SpnegoKeytab field to given value.

### HasSpnegoKeytab

`func (o *PatchedKerberosSourceRequest) HasSpnegoKeytab() bool`

HasSpnegoKeytab returns a boolean if a field has been set.

### GetSpnegoCcache

`func (o *PatchedKerberosSourceRequest) GetSpnegoCcache() string`

GetSpnegoCcache returns the SpnegoCcache field if non-nil, zero value otherwise.

### GetSpnegoCcacheOk

`func (o *PatchedKerberosSourceRequest) GetSpnegoCcacheOk() (*string, bool)`

GetSpnegoCcacheOk returns a tuple with the SpnegoCcache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpnegoCcache

`func (o *PatchedKerberosSourceRequest) SetSpnegoCcache(v string)`

SetSpnegoCcache sets SpnegoCcache field to given value.

### HasSpnegoCcache

`func (o *PatchedKerberosSourceRequest) HasSpnegoCcache() bool`

HasSpnegoCcache returns a boolean if a field has been set.

### GetPasswordLoginUpdateInternalPassword

`func (o *PatchedKerberosSourceRequest) GetPasswordLoginUpdateInternalPassword() bool`

GetPasswordLoginUpdateInternalPassword returns the PasswordLoginUpdateInternalPassword field if non-nil, zero value otherwise.

### GetPasswordLoginUpdateInternalPasswordOk

`func (o *PatchedKerberosSourceRequest) GetPasswordLoginUpdateInternalPasswordOk() (*bool, bool)`

GetPasswordLoginUpdateInternalPasswordOk returns a tuple with the PasswordLoginUpdateInternalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLoginUpdateInternalPassword

`func (o *PatchedKerberosSourceRequest) SetPasswordLoginUpdateInternalPassword(v bool)`

SetPasswordLoginUpdateInternalPassword sets PasswordLoginUpdateInternalPassword field to given value.

### HasPasswordLoginUpdateInternalPassword

`func (o *PatchedKerberosSourceRequest) HasPasswordLoginUpdateInternalPassword() bool`

HasPasswordLoginUpdateInternalPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


