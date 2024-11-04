# KerberosSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Source&#39;s display Name. | 
**Slug** | **string** | Internal source name, used in URLs. | 
**Enabled** | Pointer to **bool** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow to use when authenticating existing users. | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Flow to use when enrolling new users. | [optional] 
**UserPropertyMappings** | Pointer to **[]string** |  | [optional] 
**GroupPropertyMappings** | Pointer to **[]string** |  | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**UserMatchingMode** | Pointer to [**UserMatchingModeEnum**](UserMatchingModeEnum.md) | How the source determines if an existing user should be authenticated or a new user enrolled. | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 
**GroupMatchingMode** | Pointer to [**GroupMatchingModeEnum**](GroupMatchingModeEnum.md) | How the source determines if an existing group should be used or a new group created. | [optional] 
**Realm** | **string** | Kerberos realm | 
**Krb5Conf** | Pointer to **string** | Custom krb5.conf to use. Uses the system one by default | [optional] 
**SyncUsers** | Pointer to **bool** | Sync users from Kerberos into authentik | [optional] 
**SyncUsersPassword** | Pointer to **bool** | When a user changes their password, sync it back to Kerberos | [optional] 
**SyncPrincipal** | Pointer to **string** | Principal to authenticate to kadmin for sync. | [optional] 
**SyncPassword** | Pointer to **string** | Password to authenticate to kadmin for sync | [optional] 
**SyncKeytab** | Pointer to **string** | Keytab to authenticate to kadmin for sync. Must be base64-encoded or in the form TYPE:residual | [optional] 
**SyncCcache** | Pointer to **string** | Credentials cache to authenticate to kadmin for sync. Must be in the form TYPE:residual | [optional] 
**SpnegoServerName** | Pointer to **string** | Force the use of a specific server name for SPNEGO. Must be in the form HTTP@hostname | [optional] 
**SpnegoKeytab** | Pointer to **string** | SPNEGO keytab base64-encoded or path to keytab in the form FILE:path | [optional] 
**SpnegoCcache** | Pointer to **string** | Credential cache to use for SPNEGO in form type:residual | [optional] 
**PasswordLoginUpdateInternalPassword** | Pointer to **bool** | If enabled, the authentik-stored password will be updated upon login with the Kerberos password backend | [optional] 

## Methods

### NewKerberosSourceRequest

`func NewKerberosSourceRequest(name string, slug string, realm string, ) *KerberosSourceRequest`

NewKerberosSourceRequest instantiates a new KerberosSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosSourceRequestWithDefaults

`func NewKerberosSourceRequestWithDefaults() *KerberosSourceRequest`

NewKerberosSourceRequestWithDefaults instantiates a new KerberosSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KerberosSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KerberosSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KerberosSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *KerberosSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *KerberosSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *KerberosSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *KerberosSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KerberosSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KerberosSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KerberosSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *KerberosSourceRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *KerberosSourceRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *KerberosSourceRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *KerberosSourceRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *KerberosSourceRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *KerberosSourceRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *KerberosSourceRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *KerberosSourceRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *KerberosSourceRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *KerberosSourceRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *KerberosSourceRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *KerberosSourceRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetUserPropertyMappings

`func (o *KerberosSourceRequest) GetUserPropertyMappings() []string`

GetUserPropertyMappings returns the UserPropertyMappings field if non-nil, zero value otherwise.

### GetUserPropertyMappingsOk

`func (o *KerberosSourceRequest) GetUserPropertyMappingsOk() (*[]string, bool)`

GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertyMappings

`func (o *KerberosSourceRequest) SetUserPropertyMappings(v []string)`

SetUserPropertyMappings sets UserPropertyMappings field to given value.

### HasUserPropertyMappings

`func (o *KerberosSourceRequest) HasUserPropertyMappings() bool`

HasUserPropertyMappings returns a boolean if a field has been set.

### GetGroupPropertyMappings

`func (o *KerberosSourceRequest) GetGroupPropertyMappings() []string`

GetGroupPropertyMappings returns the GroupPropertyMappings field if non-nil, zero value otherwise.

### GetGroupPropertyMappingsOk

`func (o *KerberosSourceRequest) GetGroupPropertyMappingsOk() (*[]string, bool)`

GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPropertyMappings

`func (o *KerberosSourceRequest) SetGroupPropertyMappings(v []string)`

SetGroupPropertyMappings sets GroupPropertyMappings field to given value.

### HasGroupPropertyMappings

`func (o *KerberosSourceRequest) HasGroupPropertyMappings() bool`

HasGroupPropertyMappings returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *KerberosSourceRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *KerberosSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *KerberosSourceRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *KerberosSourceRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *KerberosSourceRequest) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *KerberosSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *KerberosSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *KerberosSourceRequest) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *KerberosSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *KerberosSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *KerberosSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *KerberosSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetGroupMatchingMode

`func (o *KerberosSourceRequest) GetGroupMatchingMode() GroupMatchingModeEnum`

GetGroupMatchingMode returns the GroupMatchingMode field if non-nil, zero value otherwise.

### GetGroupMatchingModeOk

`func (o *KerberosSourceRequest) GetGroupMatchingModeOk() (*GroupMatchingModeEnum, bool)`

GetGroupMatchingModeOk returns a tuple with the GroupMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMatchingMode

`func (o *KerberosSourceRequest) SetGroupMatchingMode(v GroupMatchingModeEnum)`

SetGroupMatchingMode sets GroupMatchingMode field to given value.

### HasGroupMatchingMode

`func (o *KerberosSourceRequest) HasGroupMatchingMode() bool`

HasGroupMatchingMode returns a boolean if a field has been set.

### GetRealm

`func (o *KerberosSourceRequest) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *KerberosSourceRequest) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *KerberosSourceRequest) SetRealm(v string)`

SetRealm sets Realm field to given value.


### GetKrb5Conf

`func (o *KerberosSourceRequest) GetKrb5Conf() string`

GetKrb5Conf returns the Krb5Conf field if non-nil, zero value otherwise.

### GetKrb5ConfOk

`func (o *KerberosSourceRequest) GetKrb5ConfOk() (*string, bool)`

GetKrb5ConfOk returns a tuple with the Krb5Conf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrb5Conf

`func (o *KerberosSourceRequest) SetKrb5Conf(v string)`

SetKrb5Conf sets Krb5Conf field to given value.

### HasKrb5Conf

`func (o *KerberosSourceRequest) HasKrb5Conf() bool`

HasKrb5Conf returns a boolean if a field has been set.

### GetSyncUsers

`func (o *KerberosSourceRequest) GetSyncUsers() bool`

GetSyncUsers returns the SyncUsers field if non-nil, zero value otherwise.

### GetSyncUsersOk

`func (o *KerberosSourceRequest) GetSyncUsersOk() (*bool, bool)`

GetSyncUsersOk returns a tuple with the SyncUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsers

`func (o *KerberosSourceRequest) SetSyncUsers(v bool)`

SetSyncUsers sets SyncUsers field to given value.

### HasSyncUsers

`func (o *KerberosSourceRequest) HasSyncUsers() bool`

HasSyncUsers returns a boolean if a field has been set.

### GetSyncUsersPassword

`func (o *KerberosSourceRequest) GetSyncUsersPassword() bool`

GetSyncUsersPassword returns the SyncUsersPassword field if non-nil, zero value otherwise.

### GetSyncUsersPasswordOk

`func (o *KerberosSourceRequest) GetSyncUsersPasswordOk() (*bool, bool)`

GetSyncUsersPasswordOk returns a tuple with the SyncUsersPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsersPassword

`func (o *KerberosSourceRequest) SetSyncUsersPassword(v bool)`

SetSyncUsersPassword sets SyncUsersPassword field to given value.

### HasSyncUsersPassword

`func (o *KerberosSourceRequest) HasSyncUsersPassword() bool`

HasSyncUsersPassword returns a boolean if a field has been set.

### GetSyncPrincipal

`func (o *KerberosSourceRequest) GetSyncPrincipal() string`

GetSyncPrincipal returns the SyncPrincipal field if non-nil, zero value otherwise.

### GetSyncPrincipalOk

`func (o *KerberosSourceRequest) GetSyncPrincipalOk() (*string, bool)`

GetSyncPrincipalOk returns a tuple with the SyncPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPrincipal

`func (o *KerberosSourceRequest) SetSyncPrincipal(v string)`

SetSyncPrincipal sets SyncPrincipal field to given value.

### HasSyncPrincipal

`func (o *KerberosSourceRequest) HasSyncPrincipal() bool`

HasSyncPrincipal returns a boolean if a field has been set.

### GetSyncPassword

`func (o *KerberosSourceRequest) GetSyncPassword() string`

GetSyncPassword returns the SyncPassword field if non-nil, zero value otherwise.

### GetSyncPasswordOk

`func (o *KerberosSourceRequest) GetSyncPasswordOk() (*string, bool)`

GetSyncPasswordOk returns a tuple with the SyncPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPassword

`func (o *KerberosSourceRequest) SetSyncPassword(v string)`

SetSyncPassword sets SyncPassword field to given value.

### HasSyncPassword

`func (o *KerberosSourceRequest) HasSyncPassword() bool`

HasSyncPassword returns a boolean if a field has been set.

### GetSyncKeytab

`func (o *KerberosSourceRequest) GetSyncKeytab() string`

GetSyncKeytab returns the SyncKeytab field if non-nil, zero value otherwise.

### GetSyncKeytabOk

`func (o *KerberosSourceRequest) GetSyncKeytabOk() (*string, bool)`

GetSyncKeytabOk returns a tuple with the SyncKeytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncKeytab

`func (o *KerberosSourceRequest) SetSyncKeytab(v string)`

SetSyncKeytab sets SyncKeytab field to given value.

### HasSyncKeytab

`func (o *KerberosSourceRequest) HasSyncKeytab() bool`

HasSyncKeytab returns a boolean if a field has been set.

### GetSyncCcache

`func (o *KerberosSourceRequest) GetSyncCcache() string`

GetSyncCcache returns the SyncCcache field if non-nil, zero value otherwise.

### GetSyncCcacheOk

`func (o *KerberosSourceRequest) GetSyncCcacheOk() (*string, bool)`

GetSyncCcacheOk returns a tuple with the SyncCcache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCcache

`func (o *KerberosSourceRequest) SetSyncCcache(v string)`

SetSyncCcache sets SyncCcache field to given value.

### HasSyncCcache

`func (o *KerberosSourceRequest) HasSyncCcache() bool`

HasSyncCcache returns a boolean if a field has been set.

### GetSpnegoServerName

`func (o *KerberosSourceRequest) GetSpnegoServerName() string`

GetSpnegoServerName returns the SpnegoServerName field if non-nil, zero value otherwise.

### GetSpnegoServerNameOk

`func (o *KerberosSourceRequest) GetSpnegoServerNameOk() (*string, bool)`

GetSpnegoServerNameOk returns a tuple with the SpnegoServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpnegoServerName

`func (o *KerberosSourceRequest) SetSpnegoServerName(v string)`

SetSpnegoServerName sets SpnegoServerName field to given value.

### HasSpnegoServerName

`func (o *KerberosSourceRequest) HasSpnegoServerName() bool`

HasSpnegoServerName returns a boolean if a field has been set.

### GetSpnegoKeytab

`func (o *KerberosSourceRequest) GetSpnegoKeytab() string`

GetSpnegoKeytab returns the SpnegoKeytab field if non-nil, zero value otherwise.

### GetSpnegoKeytabOk

`func (o *KerberosSourceRequest) GetSpnegoKeytabOk() (*string, bool)`

GetSpnegoKeytabOk returns a tuple with the SpnegoKeytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpnegoKeytab

`func (o *KerberosSourceRequest) SetSpnegoKeytab(v string)`

SetSpnegoKeytab sets SpnegoKeytab field to given value.

### HasSpnegoKeytab

`func (o *KerberosSourceRequest) HasSpnegoKeytab() bool`

HasSpnegoKeytab returns a boolean if a field has been set.

### GetSpnegoCcache

`func (o *KerberosSourceRequest) GetSpnegoCcache() string`

GetSpnegoCcache returns the SpnegoCcache field if non-nil, zero value otherwise.

### GetSpnegoCcacheOk

`func (o *KerberosSourceRequest) GetSpnegoCcacheOk() (*string, bool)`

GetSpnegoCcacheOk returns a tuple with the SpnegoCcache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpnegoCcache

`func (o *KerberosSourceRequest) SetSpnegoCcache(v string)`

SetSpnegoCcache sets SpnegoCcache field to given value.

### HasSpnegoCcache

`func (o *KerberosSourceRequest) HasSpnegoCcache() bool`

HasSpnegoCcache returns a boolean if a field has been set.

### GetPasswordLoginUpdateInternalPassword

`func (o *KerberosSourceRequest) GetPasswordLoginUpdateInternalPassword() bool`

GetPasswordLoginUpdateInternalPassword returns the PasswordLoginUpdateInternalPassword field if non-nil, zero value otherwise.

### GetPasswordLoginUpdateInternalPasswordOk

`func (o *KerberosSourceRequest) GetPasswordLoginUpdateInternalPasswordOk() (*bool, bool)`

GetPasswordLoginUpdateInternalPasswordOk returns a tuple with the PasswordLoginUpdateInternalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLoginUpdateInternalPassword

`func (o *KerberosSourceRequest) SetPasswordLoginUpdateInternalPassword(v bool)`

SetPasswordLoginUpdateInternalPassword sets PasswordLoginUpdateInternalPassword field to given value.

### HasPasswordLoginUpdateInternalPassword

`func (o *KerberosSourceRequest) HasPasswordLoginUpdateInternalPassword() bool`

HasPasswordLoginUpdateInternalPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


