# LDAPSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Source&#39;s display Name. | 
**Slug** | **string** | Internal source name, used in URLs. | 
**Enabled** | Pointer to **bool** |  | [optional] 
**AuthenticationFlow** | Pointer to **NullableString** | Flow to use when authenticating existing users. | [optional] 
**EnrollmentFlow** | Pointer to **NullableString** | Flow to use when enrolling new users. | [optional] 
**PolicyEngineMode** | Pointer to [**PolicyEngineMode**](PolicyEngineMode.md) |  | [optional] 
**UserMatchingMode** | Pointer to [**UserMatchingModeEnum**](UserMatchingModeEnum.md) | How the source determines if an existing user should be authenticated or a new user enrolled. | [optional] 
**UserPathTemplate** | Pointer to **string** |  | [optional] 
**ServerUri** | **string** |  | 
**PeerCertificate** | Pointer to **NullableString** | Optionally verify the LDAP Server&#39;s Certificate against the CA Chain in this keypair. | [optional] 
**ClientCertificate** | Pointer to **NullableString** | Client certificate to authenticate against the LDAP Server&#39;s Certificate. | [optional] 
**BindCn** | Pointer to **string** |  | [optional] 
**BindPassword** | Pointer to **string** |  | [optional] 
**StartTls** | Pointer to **bool** |  | [optional] 
**Sni** | Pointer to **bool** |  | [optional] 
**BaseDn** | **string** |  | 
**AdditionalUserDn** | Pointer to **string** | Prepended to Base DN for User-queries. | [optional] 
**AdditionalGroupDn** | Pointer to **string** | Prepended to Base DN for Group-queries. | [optional] 
**UserObjectFilter** | Pointer to **string** | Consider Objects matching this filter to be Users. | [optional] 
**GroupObjectFilter** | Pointer to **string** | Consider Objects matching this filter to be Groups. | [optional] 
**GroupMembershipField** | Pointer to **string** | Field which contains members of a group. | [optional] 
**ObjectUniquenessField** | Pointer to **string** | Field which contains a unique Identifier. | [optional] 
**PasswordLoginUpdateInternalPassword** | Pointer to **bool** | Update internal authentik password when login succeeds with LDAP | [optional] 
**SyncUsers** | Pointer to **bool** |  | [optional] 
**SyncUsersPassword** | Pointer to **bool** | When a user changes their password, sync it back to LDAP. This can only be enabled on a single LDAP source. | [optional] 
**SyncGroups** | Pointer to **bool** |  | [optional] 
**SyncParentGroup** | Pointer to **NullableString** |  | [optional] 
**PropertyMappings** | Pointer to **[]string** |  | [optional] 
**PropertyMappingsGroup** | Pointer to **[]string** | Property mappings used for group creation/updating. | [optional] 

## Methods

### NewLDAPSourceRequest

`func NewLDAPSourceRequest(name string, slug string, serverUri string, baseDn string, ) *LDAPSourceRequest`

NewLDAPSourceRequest instantiates a new LDAPSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPSourceRequestWithDefaults

`func NewLDAPSourceRequestWithDefaults() *LDAPSourceRequest`

NewLDAPSourceRequestWithDefaults instantiates a new LDAPSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LDAPSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *LDAPSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *LDAPSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *LDAPSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *LDAPSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LDAPSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LDAPSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LDAPSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *LDAPSourceRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *LDAPSourceRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *LDAPSourceRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *LDAPSourceRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *LDAPSourceRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *LDAPSourceRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *LDAPSourceRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *LDAPSourceRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *LDAPSourceRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *LDAPSourceRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *LDAPSourceRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *LDAPSourceRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetPolicyEngineMode

`func (o *LDAPSourceRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *LDAPSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *LDAPSourceRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *LDAPSourceRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *LDAPSourceRequest) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *LDAPSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *LDAPSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *LDAPSourceRequest) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *LDAPSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *LDAPSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *LDAPSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *LDAPSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetServerUri

`func (o *LDAPSourceRequest) GetServerUri() string`

GetServerUri returns the ServerUri field if non-nil, zero value otherwise.

### GetServerUriOk

`func (o *LDAPSourceRequest) GetServerUriOk() (*string, bool)`

GetServerUriOk returns a tuple with the ServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUri

`func (o *LDAPSourceRequest) SetServerUri(v string)`

SetServerUri sets ServerUri field to given value.


### GetPeerCertificate

`func (o *LDAPSourceRequest) GetPeerCertificate() string`

GetPeerCertificate returns the PeerCertificate field if non-nil, zero value otherwise.

### GetPeerCertificateOk

`func (o *LDAPSourceRequest) GetPeerCertificateOk() (*string, bool)`

GetPeerCertificateOk returns a tuple with the PeerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerCertificate

`func (o *LDAPSourceRequest) SetPeerCertificate(v string)`

SetPeerCertificate sets PeerCertificate field to given value.

### HasPeerCertificate

`func (o *LDAPSourceRequest) HasPeerCertificate() bool`

HasPeerCertificate returns a boolean if a field has been set.

### SetPeerCertificateNil

`func (o *LDAPSourceRequest) SetPeerCertificateNil(b bool)`

 SetPeerCertificateNil sets the value for PeerCertificate to be an explicit nil

### UnsetPeerCertificate
`func (o *LDAPSourceRequest) UnsetPeerCertificate()`

UnsetPeerCertificate ensures that no value is present for PeerCertificate, not even an explicit nil
### GetClientCertificate

`func (o *LDAPSourceRequest) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *LDAPSourceRequest) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *LDAPSourceRequest) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *LDAPSourceRequest) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### SetClientCertificateNil

`func (o *LDAPSourceRequest) SetClientCertificateNil(b bool)`

 SetClientCertificateNil sets the value for ClientCertificate to be an explicit nil

### UnsetClientCertificate
`func (o *LDAPSourceRequest) UnsetClientCertificate()`

UnsetClientCertificate ensures that no value is present for ClientCertificate, not even an explicit nil
### GetBindCn

`func (o *LDAPSourceRequest) GetBindCn() string`

GetBindCn returns the BindCn field if non-nil, zero value otherwise.

### GetBindCnOk

`func (o *LDAPSourceRequest) GetBindCnOk() (*string, bool)`

GetBindCnOk returns a tuple with the BindCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindCn

`func (o *LDAPSourceRequest) SetBindCn(v string)`

SetBindCn sets BindCn field to given value.

### HasBindCn

`func (o *LDAPSourceRequest) HasBindCn() bool`

HasBindCn returns a boolean if a field has been set.

### GetBindPassword

`func (o *LDAPSourceRequest) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *LDAPSourceRequest) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *LDAPSourceRequest) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.

### HasBindPassword

`func (o *LDAPSourceRequest) HasBindPassword() bool`

HasBindPassword returns a boolean if a field has been set.

### GetStartTls

`func (o *LDAPSourceRequest) GetStartTls() bool`

GetStartTls returns the StartTls field if non-nil, zero value otherwise.

### GetStartTlsOk

`func (o *LDAPSourceRequest) GetStartTlsOk() (*bool, bool)`

GetStartTlsOk returns a tuple with the StartTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTls

`func (o *LDAPSourceRequest) SetStartTls(v bool)`

SetStartTls sets StartTls field to given value.

### HasStartTls

`func (o *LDAPSourceRequest) HasStartTls() bool`

HasStartTls returns a boolean if a field has been set.

### GetSni

`func (o *LDAPSourceRequest) GetSni() bool`

GetSni returns the Sni field if non-nil, zero value otherwise.

### GetSniOk

`func (o *LDAPSourceRequest) GetSniOk() (*bool, bool)`

GetSniOk returns a tuple with the Sni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSni

`func (o *LDAPSourceRequest) SetSni(v bool)`

SetSni sets Sni field to given value.

### HasSni

`func (o *LDAPSourceRequest) HasSni() bool`

HasSni returns a boolean if a field has been set.

### GetBaseDn

`func (o *LDAPSourceRequest) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LDAPSourceRequest) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LDAPSourceRequest) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetAdditionalUserDn

`func (o *LDAPSourceRequest) GetAdditionalUserDn() string`

GetAdditionalUserDn returns the AdditionalUserDn field if non-nil, zero value otherwise.

### GetAdditionalUserDnOk

`func (o *LDAPSourceRequest) GetAdditionalUserDnOk() (*string, bool)`

GetAdditionalUserDnOk returns a tuple with the AdditionalUserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUserDn

`func (o *LDAPSourceRequest) SetAdditionalUserDn(v string)`

SetAdditionalUserDn sets AdditionalUserDn field to given value.

### HasAdditionalUserDn

`func (o *LDAPSourceRequest) HasAdditionalUserDn() bool`

HasAdditionalUserDn returns a boolean if a field has been set.

### GetAdditionalGroupDn

`func (o *LDAPSourceRequest) GetAdditionalGroupDn() string`

GetAdditionalGroupDn returns the AdditionalGroupDn field if non-nil, zero value otherwise.

### GetAdditionalGroupDnOk

`func (o *LDAPSourceRequest) GetAdditionalGroupDnOk() (*string, bool)`

GetAdditionalGroupDnOk returns a tuple with the AdditionalGroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalGroupDn

`func (o *LDAPSourceRequest) SetAdditionalGroupDn(v string)`

SetAdditionalGroupDn sets AdditionalGroupDn field to given value.

### HasAdditionalGroupDn

`func (o *LDAPSourceRequest) HasAdditionalGroupDn() bool`

HasAdditionalGroupDn returns a boolean if a field has been set.

### GetUserObjectFilter

`func (o *LDAPSourceRequest) GetUserObjectFilter() string`

GetUserObjectFilter returns the UserObjectFilter field if non-nil, zero value otherwise.

### GetUserObjectFilterOk

`func (o *LDAPSourceRequest) GetUserObjectFilterOk() (*string, bool)`

GetUserObjectFilterOk returns a tuple with the UserObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectFilter

`func (o *LDAPSourceRequest) SetUserObjectFilter(v string)`

SetUserObjectFilter sets UserObjectFilter field to given value.

### HasUserObjectFilter

`func (o *LDAPSourceRequest) HasUserObjectFilter() bool`

HasUserObjectFilter returns a boolean if a field has been set.

### GetGroupObjectFilter

`func (o *LDAPSourceRequest) GetGroupObjectFilter() string`

GetGroupObjectFilter returns the GroupObjectFilter field if non-nil, zero value otherwise.

### GetGroupObjectFilterOk

`func (o *LDAPSourceRequest) GetGroupObjectFilterOk() (*string, bool)`

GetGroupObjectFilterOk returns a tuple with the GroupObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectFilter

`func (o *LDAPSourceRequest) SetGroupObjectFilter(v string)`

SetGroupObjectFilter sets GroupObjectFilter field to given value.

### HasGroupObjectFilter

`func (o *LDAPSourceRequest) HasGroupObjectFilter() bool`

HasGroupObjectFilter returns a boolean if a field has been set.

### GetGroupMembershipField

`func (o *LDAPSourceRequest) GetGroupMembershipField() string`

GetGroupMembershipField returns the GroupMembershipField field if non-nil, zero value otherwise.

### GetGroupMembershipFieldOk

`func (o *LDAPSourceRequest) GetGroupMembershipFieldOk() (*string, bool)`

GetGroupMembershipFieldOk returns a tuple with the GroupMembershipField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMembershipField

`func (o *LDAPSourceRequest) SetGroupMembershipField(v string)`

SetGroupMembershipField sets GroupMembershipField field to given value.

### HasGroupMembershipField

`func (o *LDAPSourceRequest) HasGroupMembershipField() bool`

HasGroupMembershipField returns a boolean if a field has been set.

### GetObjectUniquenessField

`func (o *LDAPSourceRequest) GetObjectUniquenessField() string`

GetObjectUniquenessField returns the ObjectUniquenessField field if non-nil, zero value otherwise.

### GetObjectUniquenessFieldOk

`func (o *LDAPSourceRequest) GetObjectUniquenessFieldOk() (*string, bool)`

GetObjectUniquenessFieldOk returns a tuple with the ObjectUniquenessField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUniquenessField

`func (o *LDAPSourceRequest) SetObjectUniquenessField(v string)`

SetObjectUniquenessField sets ObjectUniquenessField field to given value.

### HasObjectUniquenessField

`func (o *LDAPSourceRequest) HasObjectUniquenessField() bool`

HasObjectUniquenessField returns a boolean if a field has been set.

### GetPasswordLoginUpdateInternalPassword

`func (o *LDAPSourceRequest) GetPasswordLoginUpdateInternalPassword() bool`

GetPasswordLoginUpdateInternalPassword returns the PasswordLoginUpdateInternalPassword field if non-nil, zero value otherwise.

### GetPasswordLoginUpdateInternalPasswordOk

`func (o *LDAPSourceRequest) GetPasswordLoginUpdateInternalPasswordOk() (*bool, bool)`

GetPasswordLoginUpdateInternalPasswordOk returns a tuple with the PasswordLoginUpdateInternalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLoginUpdateInternalPassword

`func (o *LDAPSourceRequest) SetPasswordLoginUpdateInternalPassword(v bool)`

SetPasswordLoginUpdateInternalPassword sets PasswordLoginUpdateInternalPassword field to given value.

### HasPasswordLoginUpdateInternalPassword

`func (o *LDAPSourceRequest) HasPasswordLoginUpdateInternalPassword() bool`

HasPasswordLoginUpdateInternalPassword returns a boolean if a field has been set.

### GetSyncUsers

`func (o *LDAPSourceRequest) GetSyncUsers() bool`

GetSyncUsers returns the SyncUsers field if non-nil, zero value otherwise.

### GetSyncUsersOk

`func (o *LDAPSourceRequest) GetSyncUsersOk() (*bool, bool)`

GetSyncUsersOk returns a tuple with the SyncUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsers

`func (o *LDAPSourceRequest) SetSyncUsers(v bool)`

SetSyncUsers sets SyncUsers field to given value.

### HasSyncUsers

`func (o *LDAPSourceRequest) HasSyncUsers() bool`

HasSyncUsers returns a boolean if a field has been set.

### GetSyncUsersPassword

`func (o *LDAPSourceRequest) GetSyncUsersPassword() bool`

GetSyncUsersPassword returns the SyncUsersPassword field if non-nil, zero value otherwise.

### GetSyncUsersPasswordOk

`func (o *LDAPSourceRequest) GetSyncUsersPasswordOk() (*bool, bool)`

GetSyncUsersPasswordOk returns a tuple with the SyncUsersPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsersPassword

`func (o *LDAPSourceRequest) SetSyncUsersPassword(v bool)`

SetSyncUsersPassword sets SyncUsersPassword field to given value.

### HasSyncUsersPassword

`func (o *LDAPSourceRequest) HasSyncUsersPassword() bool`

HasSyncUsersPassword returns a boolean if a field has been set.

### GetSyncGroups

`func (o *LDAPSourceRequest) GetSyncGroups() bool`

GetSyncGroups returns the SyncGroups field if non-nil, zero value otherwise.

### GetSyncGroupsOk

`func (o *LDAPSourceRequest) GetSyncGroupsOk() (*bool, bool)`

GetSyncGroupsOk returns a tuple with the SyncGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncGroups

`func (o *LDAPSourceRequest) SetSyncGroups(v bool)`

SetSyncGroups sets SyncGroups field to given value.

### HasSyncGroups

`func (o *LDAPSourceRequest) HasSyncGroups() bool`

HasSyncGroups returns a boolean if a field has been set.

### GetSyncParentGroup

`func (o *LDAPSourceRequest) GetSyncParentGroup() string`

GetSyncParentGroup returns the SyncParentGroup field if non-nil, zero value otherwise.

### GetSyncParentGroupOk

`func (o *LDAPSourceRequest) GetSyncParentGroupOk() (*string, bool)`

GetSyncParentGroupOk returns a tuple with the SyncParentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncParentGroup

`func (o *LDAPSourceRequest) SetSyncParentGroup(v string)`

SetSyncParentGroup sets SyncParentGroup field to given value.

### HasSyncParentGroup

`func (o *LDAPSourceRequest) HasSyncParentGroup() bool`

HasSyncParentGroup returns a boolean if a field has been set.

### SetSyncParentGroupNil

`func (o *LDAPSourceRequest) SetSyncParentGroupNil(b bool)`

 SetSyncParentGroupNil sets the value for SyncParentGroup to be an explicit nil

### UnsetSyncParentGroup
`func (o *LDAPSourceRequest) UnsetSyncParentGroup()`

UnsetSyncParentGroup ensures that no value is present for SyncParentGroup, not even an explicit nil
### GetPropertyMappings

`func (o *LDAPSourceRequest) GetPropertyMappings() []string`

GetPropertyMappings returns the PropertyMappings field if non-nil, zero value otherwise.

### GetPropertyMappingsOk

`func (o *LDAPSourceRequest) GetPropertyMappingsOk() (*[]string, bool)`

GetPropertyMappingsOk returns a tuple with the PropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappings

`func (o *LDAPSourceRequest) SetPropertyMappings(v []string)`

SetPropertyMappings sets PropertyMappings field to given value.

### HasPropertyMappings

`func (o *LDAPSourceRequest) HasPropertyMappings() bool`

HasPropertyMappings returns a boolean if a field has been set.

### GetPropertyMappingsGroup

`func (o *LDAPSourceRequest) GetPropertyMappingsGroup() []string`

GetPropertyMappingsGroup returns the PropertyMappingsGroup field if non-nil, zero value otherwise.

### GetPropertyMappingsGroupOk

`func (o *LDAPSourceRequest) GetPropertyMappingsGroupOk() (*[]string, bool)`

GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyMappingsGroup

`func (o *LDAPSourceRequest) SetPropertyMappingsGroup(v []string)`

SetPropertyMappingsGroup sets PropertyMappingsGroup field to given value.

### HasPropertyMappingsGroup

`func (o *LDAPSourceRequest) HasPropertyMappingsGroup() bool`

HasPropertyMappingsGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


