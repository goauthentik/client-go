# PatchedLDAPSourceRequest

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
**ServerUri** | Pointer to **string** |  | [optional] 
**PeerCertificate** | Pointer to **NullableString** | Optionally verify the LDAP Server&#39;s Certificate against the CA Chain in this keypair. | [optional] 
**ClientCertificate** | Pointer to **NullableString** | Client certificate to authenticate against the LDAP Server&#39;s Certificate. | [optional] 
**BindCn** | Pointer to **string** |  | [optional] 
**BindPassword** | Pointer to **string** |  | [optional] 
**StartTls** | Pointer to **bool** |  | [optional] 
**Sni** | Pointer to **bool** |  | [optional] 
**BaseDn** | Pointer to **string** |  | [optional] 
**AdditionalUserDn** | Pointer to **string** | Prepended to Base DN for User-queries. | [optional] 
**AdditionalGroupDn** | Pointer to **string** | Prepended to Base DN for Group-queries. | [optional] 
**UserObjectFilter** | Pointer to **string** | Consider Objects matching this filter to be Users. | [optional] 
**GroupObjectFilter** | Pointer to **string** | Consider Objects matching this filter to be Groups. | [optional] 
**GroupMembershipField** | Pointer to **string** | Field which contains members of a group. | [optional] 
**UserMembershipAttribute** | Pointer to **string** | Attribute which matches the value of &#x60;group_membership_field&#x60;. | [optional] 
**ObjectUniquenessField** | Pointer to **string** | Field which contains a unique Identifier. | [optional] 
**PasswordLoginUpdateInternalPassword** | Pointer to **bool** | Update internal authentik password when login succeeds with LDAP | [optional] 
**SyncUsers** | Pointer to **bool** |  | [optional] 
**SyncUsersPassword** | Pointer to **bool** | When a user changes their password, sync it back to LDAP. This can only be enabled on a single LDAP source. | [optional] 
**SyncGroups** | Pointer to **bool** |  | [optional] 
**SyncParentGroup** | Pointer to **NullableString** |  | [optional] 
**LookupGroupsFromUser** | Pointer to **bool** | Lookup group membership based on a user attribute instead of a group attribute. This allows nested group resolution on systems like FreeIPA and Active Directory | [optional] 
**DeleteNotFoundObjects** | Pointer to **bool** | Delete authentik users and groups which were previously supplied by this source, but are now missing from it. | [optional] 

## Methods

### NewPatchedLDAPSourceRequest

`func NewPatchedLDAPSourceRequest() *PatchedLDAPSourceRequest`

NewPatchedLDAPSourceRequest instantiates a new PatchedLDAPSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedLDAPSourceRequestWithDefaults

`func NewPatchedLDAPSourceRequestWithDefaults() *PatchedLDAPSourceRequest`

NewPatchedLDAPSourceRequestWithDefaults instantiates a new PatchedLDAPSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedLDAPSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedLDAPSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedLDAPSourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedLDAPSourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedLDAPSourceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedLDAPSourceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedLDAPSourceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedLDAPSourceRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedLDAPSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedLDAPSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedLDAPSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedLDAPSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *PatchedLDAPSourceRequest) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *PatchedLDAPSourceRequest) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *PatchedLDAPSourceRequest) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *PatchedLDAPSourceRequest) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *PatchedLDAPSourceRequest) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *PatchedLDAPSourceRequest) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *PatchedLDAPSourceRequest) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *PatchedLDAPSourceRequest) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *PatchedLDAPSourceRequest) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *PatchedLDAPSourceRequest) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *PatchedLDAPSourceRequest) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *PatchedLDAPSourceRequest) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetUserPropertyMappings

`func (o *PatchedLDAPSourceRequest) GetUserPropertyMappings() []string`

GetUserPropertyMappings returns the UserPropertyMappings field if non-nil, zero value otherwise.

### GetUserPropertyMappingsOk

`func (o *PatchedLDAPSourceRequest) GetUserPropertyMappingsOk() (*[]string, bool)`

GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertyMappings

`func (o *PatchedLDAPSourceRequest) SetUserPropertyMappings(v []string)`

SetUserPropertyMappings sets UserPropertyMappings field to given value.

### HasUserPropertyMappings

`func (o *PatchedLDAPSourceRequest) HasUserPropertyMappings() bool`

HasUserPropertyMappings returns a boolean if a field has been set.

### GetGroupPropertyMappings

`func (o *PatchedLDAPSourceRequest) GetGroupPropertyMappings() []string`

GetGroupPropertyMappings returns the GroupPropertyMappings field if non-nil, zero value otherwise.

### GetGroupPropertyMappingsOk

`func (o *PatchedLDAPSourceRequest) GetGroupPropertyMappingsOk() (*[]string, bool)`

GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPropertyMappings

`func (o *PatchedLDAPSourceRequest) SetGroupPropertyMappings(v []string)`

SetGroupPropertyMappings sets GroupPropertyMappings field to given value.

### HasGroupPropertyMappings

`func (o *PatchedLDAPSourceRequest) HasGroupPropertyMappings() bool`

HasGroupPropertyMappings returns a boolean if a field has been set.

### GetPolicyEngineMode

`func (o *PatchedLDAPSourceRequest) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *PatchedLDAPSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *PatchedLDAPSourceRequest) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *PatchedLDAPSourceRequest) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *PatchedLDAPSourceRequest) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *PatchedLDAPSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *PatchedLDAPSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *PatchedLDAPSourceRequest) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetUserPathTemplate

`func (o *PatchedLDAPSourceRequest) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *PatchedLDAPSourceRequest) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *PatchedLDAPSourceRequest) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *PatchedLDAPSourceRequest) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetServerUri

`func (o *PatchedLDAPSourceRequest) GetServerUri() string`

GetServerUri returns the ServerUri field if non-nil, zero value otherwise.

### GetServerUriOk

`func (o *PatchedLDAPSourceRequest) GetServerUriOk() (*string, bool)`

GetServerUriOk returns a tuple with the ServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUri

`func (o *PatchedLDAPSourceRequest) SetServerUri(v string)`

SetServerUri sets ServerUri field to given value.

### HasServerUri

`func (o *PatchedLDAPSourceRequest) HasServerUri() bool`

HasServerUri returns a boolean if a field has been set.

### GetPeerCertificate

`func (o *PatchedLDAPSourceRequest) GetPeerCertificate() string`

GetPeerCertificate returns the PeerCertificate field if non-nil, zero value otherwise.

### GetPeerCertificateOk

`func (o *PatchedLDAPSourceRequest) GetPeerCertificateOk() (*string, bool)`

GetPeerCertificateOk returns a tuple with the PeerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerCertificate

`func (o *PatchedLDAPSourceRequest) SetPeerCertificate(v string)`

SetPeerCertificate sets PeerCertificate field to given value.

### HasPeerCertificate

`func (o *PatchedLDAPSourceRequest) HasPeerCertificate() bool`

HasPeerCertificate returns a boolean if a field has been set.

### SetPeerCertificateNil

`func (o *PatchedLDAPSourceRequest) SetPeerCertificateNil(b bool)`

 SetPeerCertificateNil sets the value for PeerCertificate to be an explicit nil

### UnsetPeerCertificate
`func (o *PatchedLDAPSourceRequest) UnsetPeerCertificate()`

UnsetPeerCertificate ensures that no value is present for PeerCertificate, not even an explicit nil
### GetClientCertificate

`func (o *PatchedLDAPSourceRequest) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *PatchedLDAPSourceRequest) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *PatchedLDAPSourceRequest) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *PatchedLDAPSourceRequest) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### SetClientCertificateNil

`func (o *PatchedLDAPSourceRequest) SetClientCertificateNil(b bool)`

 SetClientCertificateNil sets the value for ClientCertificate to be an explicit nil

### UnsetClientCertificate
`func (o *PatchedLDAPSourceRequest) UnsetClientCertificate()`

UnsetClientCertificate ensures that no value is present for ClientCertificate, not even an explicit nil
### GetBindCn

`func (o *PatchedLDAPSourceRequest) GetBindCn() string`

GetBindCn returns the BindCn field if non-nil, zero value otherwise.

### GetBindCnOk

`func (o *PatchedLDAPSourceRequest) GetBindCnOk() (*string, bool)`

GetBindCnOk returns a tuple with the BindCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindCn

`func (o *PatchedLDAPSourceRequest) SetBindCn(v string)`

SetBindCn sets BindCn field to given value.

### HasBindCn

`func (o *PatchedLDAPSourceRequest) HasBindCn() bool`

HasBindCn returns a boolean if a field has been set.

### GetBindPassword

`func (o *PatchedLDAPSourceRequest) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *PatchedLDAPSourceRequest) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *PatchedLDAPSourceRequest) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.

### HasBindPassword

`func (o *PatchedLDAPSourceRequest) HasBindPassword() bool`

HasBindPassword returns a boolean if a field has been set.

### GetStartTls

`func (o *PatchedLDAPSourceRequest) GetStartTls() bool`

GetStartTls returns the StartTls field if non-nil, zero value otherwise.

### GetStartTlsOk

`func (o *PatchedLDAPSourceRequest) GetStartTlsOk() (*bool, bool)`

GetStartTlsOk returns a tuple with the StartTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTls

`func (o *PatchedLDAPSourceRequest) SetStartTls(v bool)`

SetStartTls sets StartTls field to given value.

### HasStartTls

`func (o *PatchedLDAPSourceRequest) HasStartTls() bool`

HasStartTls returns a boolean if a field has been set.

### GetSni

`func (o *PatchedLDAPSourceRequest) GetSni() bool`

GetSni returns the Sni field if non-nil, zero value otherwise.

### GetSniOk

`func (o *PatchedLDAPSourceRequest) GetSniOk() (*bool, bool)`

GetSniOk returns a tuple with the Sni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSni

`func (o *PatchedLDAPSourceRequest) SetSni(v bool)`

SetSni sets Sni field to given value.

### HasSni

`func (o *PatchedLDAPSourceRequest) HasSni() bool`

HasSni returns a boolean if a field has been set.

### GetBaseDn

`func (o *PatchedLDAPSourceRequest) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *PatchedLDAPSourceRequest) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *PatchedLDAPSourceRequest) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *PatchedLDAPSourceRequest) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetAdditionalUserDn

`func (o *PatchedLDAPSourceRequest) GetAdditionalUserDn() string`

GetAdditionalUserDn returns the AdditionalUserDn field if non-nil, zero value otherwise.

### GetAdditionalUserDnOk

`func (o *PatchedLDAPSourceRequest) GetAdditionalUserDnOk() (*string, bool)`

GetAdditionalUserDnOk returns a tuple with the AdditionalUserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUserDn

`func (o *PatchedLDAPSourceRequest) SetAdditionalUserDn(v string)`

SetAdditionalUserDn sets AdditionalUserDn field to given value.

### HasAdditionalUserDn

`func (o *PatchedLDAPSourceRequest) HasAdditionalUserDn() bool`

HasAdditionalUserDn returns a boolean if a field has been set.

### GetAdditionalGroupDn

`func (o *PatchedLDAPSourceRequest) GetAdditionalGroupDn() string`

GetAdditionalGroupDn returns the AdditionalGroupDn field if non-nil, zero value otherwise.

### GetAdditionalGroupDnOk

`func (o *PatchedLDAPSourceRequest) GetAdditionalGroupDnOk() (*string, bool)`

GetAdditionalGroupDnOk returns a tuple with the AdditionalGroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalGroupDn

`func (o *PatchedLDAPSourceRequest) SetAdditionalGroupDn(v string)`

SetAdditionalGroupDn sets AdditionalGroupDn field to given value.

### HasAdditionalGroupDn

`func (o *PatchedLDAPSourceRequest) HasAdditionalGroupDn() bool`

HasAdditionalGroupDn returns a boolean if a field has been set.

### GetUserObjectFilter

`func (o *PatchedLDAPSourceRequest) GetUserObjectFilter() string`

GetUserObjectFilter returns the UserObjectFilter field if non-nil, zero value otherwise.

### GetUserObjectFilterOk

`func (o *PatchedLDAPSourceRequest) GetUserObjectFilterOk() (*string, bool)`

GetUserObjectFilterOk returns a tuple with the UserObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectFilter

`func (o *PatchedLDAPSourceRequest) SetUserObjectFilter(v string)`

SetUserObjectFilter sets UserObjectFilter field to given value.

### HasUserObjectFilter

`func (o *PatchedLDAPSourceRequest) HasUserObjectFilter() bool`

HasUserObjectFilter returns a boolean if a field has been set.

### GetGroupObjectFilter

`func (o *PatchedLDAPSourceRequest) GetGroupObjectFilter() string`

GetGroupObjectFilter returns the GroupObjectFilter field if non-nil, zero value otherwise.

### GetGroupObjectFilterOk

`func (o *PatchedLDAPSourceRequest) GetGroupObjectFilterOk() (*string, bool)`

GetGroupObjectFilterOk returns a tuple with the GroupObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectFilter

`func (o *PatchedLDAPSourceRequest) SetGroupObjectFilter(v string)`

SetGroupObjectFilter sets GroupObjectFilter field to given value.

### HasGroupObjectFilter

`func (o *PatchedLDAPSourceRequest) HasGroupObjectFilter() bool`

HasGroupObjectFilter returns a boolean if a field has been set.

### GetGroupMembershipField

`func (o *PatchedLDAPSourceRequest) GetGroupMembershipField() string`

GetGroupMembershipField returns the GroupMembershipField field if non-nil, zero value otherwise.

### GetGroupMembershipFieldOk

`func (o *PatchedLDAPSourceRequest) GetGroupMembershipFieldOk() (*string, bool)`

GetGroupMembershipFieldOk returns a tuple with the GroupMembershipField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMembershipField

`func (o *PatchedLDAPSourceRequest) SetGroupMembershipField(v string)`

SetGroupMembershipField sets GroupMembershipField field to given value.

### HasGroupMembershipField

`func (o *PatchedLDAPSourceRequest) HasGroupMembershipField() bool`

HasGroupMembershipField returns a boolean if a field has been set.

### GetUserMembershipAttribute

`func (o *PatchedLDAPSourceRequest) GetUserMembershipAttribute() string`

GetUserMembershipAttribute returns the UserMembershipAttribute field if non-nil, zero value otherwise.

### GetUserMembershipAttributeOk

`func (o *PatchedLDAPSourceRequest) GetUserMembershipAttributeOk() (*string, bool)`

GetUserMembershipAttributeOk returns a tuple with the UserMembershipAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMembershipAttribute

`func (o *PatchedLDAPSourceRequest) SetUserMembershipAttribute(v string)`

SetUserMembershipAttribute sets UserMembershipAttribute field to given value.

### HasUserMembershipAttribute

`func (o *PatchedLDAPSourceRequest) HasUserMembershipAttribute() bool`

HasUserMembershipAttribute returns a boolean if a field has been set.

### GetObjectUniquenessField

`func (o *PatchedLDAPSourceRequest) GetObjectUniquenessField() string`

GetObjectUniquenessField returns the ObjectUniquenessField field if non-nil, zero value otherwise.

### GetObjectUniquenessFieldOk

`func (o *PatchedLDAPSourceRequest) GetObjectUniquenessFieldOk() (*string, bool)`

GetObjectUniquenessFieldOk returns a tuple with the ObjectUniquenessField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUniquenessField

`func (o *PatchedLDAPSourceRequest) SetObjectUniquenessField(v string)`

SetObjectUniquenessField sets ObjectUniquenessField field to given value.

### HasObjectUniquenessField

`func (o *PatchedLDAPSourceRequest) HasObjectUniquenessField() bool`

HasObjectUniquenessField returns a boolean if a field has been set.

### GetPasswordLoginUpdateInternalPassword

`func (o *PatchedLDAPSourceRequest) GetPasswordLoginUpdateInternalPassword() bool`

GetPasswordLoginUpdateInternalPassword returns the PasswordLoginUpdateInternalPassword field if non-nil, zero value otherwise.

### GetPasswordLoginUpdateInternalPasswordOk

`func (o *PatchedLDAPSourceRequest) GetPasswordLoginUpdateInternalPasswordOk() (*bool, bool)`

GetPasswordLoginUpdateInternalPasswordOk returns a tuple with the PasswordLoginUpdateInternalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLoginUpdateInternalPassword

`func (o *PatchedLDAPSourceRequest) SetPasswordLoginUpdateInternalPassword(v bool)`

SetPasswordLoginUpdateInternalPassword sets PasswordLoginUpdateInternalPassword field to given value.

### HasPasswordLoginUpdateInternalPassword

`func (o *PatchedLDAPSourceRequest) HasPasswordLoginUpdateInternalPassword() bool`

HasPasswordLoginUpdateInternalPassword returns a boolean if a field has been set.

### GetSyncUsers

`func (o *PatchedLDAPSourceRequest) GetSyncUsers() bool`

GetSyncUsers returns the SyncUsers field if non-nil, zero value otherwise.

### GetSyncUsersOk

`func (o *PatchedLDAPSourceRequest) GetSyncUsersOk() (*bool, bool)`

GetSyncUsersOk returns a tuple with the SyncUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsers

`func (o *PatchedLDAPSourceRequest) SetSyncUsers(v bool)`

SetSyncUsers sets SyncUsers field to given value.

### HasSyncUsers

`func (o *PatchedLDAPSourceRequest) HasSyncUsers() bool`

HasSyncUsers returns a boolean if a field has been set.

### GetSyncUsersPassword

`func (o *PatchedLDAPSourceRequest) GetSyncUsersPassword() bool`

GetSyncUsersPassword returns the SyncUsersPassword field if non-nil, zero value otherwise.

### GetSyncUsersPasswordOk

`func (o *PatchedLDAPSourceRequest) GetSyncUsersPasswordOk() (*bool, bool)`

GetSyncUsersPasswordOk returns a tuple with the SyncUsersPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsersPassword

`func (o *PatchedLDAPSourceRequest) SetSyncUsersPassword(v bool)`

SetSyncUsersPassword sets SyncUsersPassword field to given value.

### HasSyncUsersPassword

`func (o *PatchedLDAPSourceRequest) HasSyncUsersPassword() bool`

HasSyncUsersPassword returns a boolean if a field has been set.

### GetSyncGroups

`func (o *PatchedLDAPSourceRequest) GetSyncGroups() bool`

GetSyncGroups returns the SyncGroups field if non-nil, zero value otherwise.

### GetSyncGroupsOk

`func (o *PatchedLDAPSourceRequest) GetSyncGroupsOk() (*bool, bool)`

GetSyncGroupsOk returns a tuple with the SyncGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncGroups

`func (o *PatchedLDAPSourceRequest) SetSyncGroups(v bool)`

SetSyncGroups sets SyncGroups field to given value.

### HasSyncGroups

`func (o *PatchedLDAPSourceRequest) HasSyncGroups() bool`

HasSyncGroups returns a boolean if a field has been set.

### GetSyncParentGroup

`func (o *PatchedLDAPSourceRequest) GetSyncParentGroup() string`

GetSyncParentGroup returns the SyncParentGroup field if non-nil, zero value otherwise.

### GetSyncParentGroupOk

`func (o *PatchedLDAPSourceRequest) GetSyncParentGroupOk() (*string, bool)`

GetSyncParentGroupOk returns a tuple with the SyncParentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncParentGroup

`func (o *PatchedLDAPSourceRequest) SetSyncParentGroup(v string)`

SetSyncParentGroup sets SyncParentGroup field to given value.

### HasSyncParentGroup

`func (o *PatchedLDAPSourceRequest) HasSyncParentGroup() bool`

HasSyncParentGroup returns a boolean if a field has been set.

### SetSyncParentGroupNil

`func (o *PatchedLDAPSourceRequest) SetSyncParentGroupNil(b bool)`

 SetSyncParentGroupNil sets the value for SyncParentGroup to be an explicit nil

### UnsetSyncParentGroup
`func (o *PatchedLDAPSourceRequest) UnsetSyncParentGroup()`

UnsetSyncParentGroup ensures that no value is present for SyncParentGroup, not even an explicit nil
### GetLookupGroupsFromUser

`func (o *PatchedLDAPSourceRequest) GetLookupGroupsFromUser() bool`

GetLookupGroupsFromUser returns the LookupGroupsFromUser field if non-nil, zero value otherwise.

### GetLookupGroupsFromUserOk

`func (o *PatchedLDAPSourceRequest) GetLookupGroupsFromUserOk() (*bool, bool)`

GetLookupGroupsFromUserOk returns a tuple with the LookupGroupsFromUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupGroupsFromUser

`func (o *PatchedLDAPSourceRequest) SetLookupGroupsFromUser(v bool)`

SetLookupGroupsFromUser sets LookupGroupsFromUser field to given value.

### HasLookupGroupsFromUser

`func (o *PatchedLDAPSourceRequest) HasLookupGroupsFromUser() bool`

HasLookupGroupsFromUser returns a boolean if a field has been set.

### GetDeleteNotFoundObjects

`func (o *PatchedLDAPSourceRequest) GetDeleteNotFoundObjects() bool`

GetDeleteNotFoundObjects returns the DeleteNotFoundObjects field if non-nil, zero value otherwise.

### GetDeleteNotFoundObjectsOk

`func (o *PatchedLDAPSourceRequest) GetDeleteNotFoundObjectsOk() (*bool, bool)`

GetDeleteNotFoundObjectsOk returns a tuple with the DeleteNotFoundObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteNotFoundObjects

`func (o *PatchedLDAPSourceRequest) SetDeleteNotFoundObjects(v bool)`

SetDeleteNotFoundObjects sets DeleteNotFoundObjects field to given value.

### HasDeleteNotFoundObjects

`func (o *PatchedLDAPSourceRequest) HasDeleteNotFoundObjects() bool`

HasDeleteNotFoundObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


