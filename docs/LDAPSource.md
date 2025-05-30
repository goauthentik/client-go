# LDAPSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pk** | **string** |  | [readonly] 
**Name** | **string** | Source&#39;s display Name. | 
**Slug** | **string** | Internal source name, used in URLs. | 
**Enabled** | Pointer to **bool** |  | [optional] 
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
**Icon** | **string** |  | [readonly] 
**ServerUri** | **string** |  | 
**PeerCertificate** | Pointer to **NullableString** | Optionally verify the LDAP Server&#39;s Certificate against the CA Chain in this keypair. | [optional] 
**ClientCertificate** | Pointer to **NullableString** | Client certificate to authenticate against the LDAP Server&#39;s Certificate. | [optional] 
**BindCn** | Pointer to **string** |  | [optional] 
**StartTls** | Pointer to **bool** |  | [optional] 
**Sni** | Pointer to **bool** |  | [optional] 
**BaseDn** | **string** |  | 
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
**Connectivity** | **map[string]map[string]string** | Get cached source connectivity | [readonly] 
**LookupGroupsFromUser** | Pointer to **bool** | Lookup group membership based on a user attribute instead of a group attribute. This allows nested group resolution on systems like FreeIPA and Active Directory | [optional] 
**DeleteNotFoundObjects** | Pointer to **bool** | Delete authentik users and groups which were previously supplied by this source, but are now missing from it. | [optional] 

## Methods

### NewLDAPSource

`func NewLDAPSource(pk string, name string, slug string, component string, verboseName string, verboseNamePlural string, metaModelName string, managed NullableString, icon string, serverUri string, baseDn string, connectivity map[string]map[string]string, ) *LDAPSource`

NewLDAPSource instantiates a new LDAPSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPSourceWithDefaults

`func NewLDAPSourceWithDefaults() *LDAPSource`

NewLDAPSourceWithDefaults instantiates a new LDAPSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPk

`func (o *LDAPSource) GetPk() string`

GetPk returns the Pk field if non-nil, zero value otherwise.

### GetPkOk

`func (o *LDAPSource) GetPkOk() (*string, bool)`

GetPkOk returns a tuple with the Pk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPk

`func (o *LDAPSource) SetPk(v string)`

SetPk sets Pk field to given value.


### GetName

`func (o *LDAPSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPSource) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *LDAPSource) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *LDAPSource) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *LDAPSource) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetEnabled

`func (o *LDAPSource) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LDAPSource) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LDAPSource) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LDAPSource) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *LDAPSource) GetAuthenticationFlow() string`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *LDAPSource) GetAuthenticationFlowOk() (*string, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *LDAPSource) SetAuthenticationFlow(v string)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *LDAPSource) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### SetAuthenticationFlowNil

`func (o *LDAPSource) SetAuthenticationFlowNil(b bool)`

 SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil

### UnsetAuthenticationFlow
`func (o *LDAPSource) UnsetAuthenticationFlow()`

UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
### GetEnrollmentFlow

`func (o *LDAPSource) GetEnrollmentFlow() string`

GetEnrollmentFlow returns the EnrollmentFlow field if non-nil, zero value otherwise.

### GetEnrollmentFlowOk

`func (o *LDAPSource) GetEnrollmentFlowOk() (*string, bool)`

GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentFlow

`func (o *LDAPSource) SetEnrollmentFlow(v string)`

SetEnrollmentFlow sets EnrollmentFlow field to given value.

### HasEnrollmentFlow

`func (o *LDAPSource) HasEnrollmentFlow() bool`

HasEnrollmentFlow returns a boolean if a field has been set.

### SetEnrollmentFlowNil

`func (o *LDAPSource) SetEnrollmentFlowNil(b bool)`

 SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil

### UnsetEnrollmentFlow
`func (o *LDAPSource) UnsetEnrollmentFlow()`

UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
### GetUserPropertyMappings

`func (o *LDAPSource) GetUserPropertyMappings() []string`

GetUserPropertyMappings returns the UserPropertyMappings field if non-nil, zero value otherwise.

### GetUserPropertyMappingsOk

`func (o *LDAPSource) GetUserPropertyMappingsOk() (*[]string, bool)`

GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPropertyMappings

`func (o *LDAPSource) SetUserPropertyMappings(v []string)`

SetUserPropertyMappings sets UserPropertyMappings field to given value.

### HasUserPropertyMappings

`func (o *LDAPSource) HasUserPropertyMappings() bool`

HasUserPropertyMappings returns a boolean if a field has been set.

### GetGroupPropertyMappings

`func (o *LDAPSource) GetGroupPropertyMappings() []string`

GetGroupPropertyMappings returns the GroupPropertyMappings field if non-nil, zero value otherwise.

### GetGroupPropertyMappingsOk

`func (o *LDAPSource) GetGroupPropertyMappingsOk() (*[]string, bool)`

GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPropertyMappings

`func (o *LDAPSource) SetGroupPropertyMappings(v []string)`

SetGroupPropertyMappings sets GroupPropertyMappings field to given value.

### HasGroupPropertyMappings

`func (o *LDAPSource) HasGroupPropertyMappings() bool`

HasGroupPropertyMappings returns a boolean if a field has been set.

### GetComponent

`func (o *LDAPSource) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *LDAPSource) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *LDAPSource) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetVerboseName

`func (o *LDAPSource) GetVerboseName() string`

GetVerboseName returns the VerboseName field if non-nil, zero value otherwise.

### GetVerboseNameOk

`func (o *LDAPSource) GetVerboseNameOk() (*string, bool)`

GetVerboseNameOk returns a tuple with the VerboseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseName

`func (o *LDAPSource) SetVerboseName(v string)`

SetVerboseName sets VerboseName field to given value.


### GetVerboseNamePlural

`func (o *LDAPSource) GetVerboseNamePlural() string`

GetVerboseNamePlural returns the VerboseNamePlural field if non-nil, zero value otherwise.

### GetVerboseNamePluralOk

`func (o *LDAPSource) GetVerboseNamePluralOk() (*string, bool)`

GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseNamePlural

`func (o *LDAPSource) SetVerboseNamePlural(v string)`

SetVerboseNamePlural sets VerboseNamePlural field to given value.


### GetMetaModelName

`func (o *LDAPSource) GetMetaModelName() string`

GetMetaModelName returns the MetaModelName field if non-nil, zero value otherwise.

### GetMetaModelNameOk

`func (o *LDAPSource) GetMetaModelNameOk() (*string, bool)`

GetMetaModelNameOk returns a tuple with the MetaModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaModelName

`func (o *LDAPSource) SetMetaModelName(v string)`

SetMetaModelName sets MetaModelName field to given value.


### GetPolicyEngineMode

`func (o *LDAPSource) GetPolicyEngineMode() PolicyEngineMode`

GetPolicyEngineMode returns the PolicyEngineMode field if non-nil, zero value otherwise.

### GetPolicyEngineModeOk

`func (o *LDAPSource) GetPolicyEngineModeOk() (*PolicyEngineMode, bool)`

GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEngineMode

`func (o *LDAPSource) SetPolicyEngineMode(v PolicyEngineMode)`

SetPolicyEngineMode sets PolicyEngineMode field to given value.

### HasPolicyEngineMode

`func (o *LDAPSource) HasPolicyEngineMode() bool`

HasPolicyEngineMode returns a boolean if a field has been set.

### GetUserMatchingMode

`func (o *LDAPSource) GetUserMatchingMode() UserMatchingModeEnum`

GetUserMatchingMode returns the UserMatchingMode field if non-nil, zero value otherwise.

### GetUserMatchingModeOk

`func (o *LDAPSource) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool)`

GetUserMatchingModeOk returns a tuple with the UserMatchingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMatchingMode

`func (o *LDAPSource) SetUserMatchingMode(v UserMatchingModeEnum)`

SetUserMatchingMode sets UserMatchingMode field to given value.

### HasUserMatchingMode

`func (o *LDAPSource) HasUserMatchingMode() bool`

HasUserMatchingMode returns a boolean if a field has been set.

### GetManaged

`func (o *LDAPSource) GetManaged() string`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *LDAPSource) GetManagedOk() (*string, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *LDAPSource) SetManaged(v string)`

SetManaged sets Managed field to given value.


### SetManagedNil

`func (o *LDAPSource) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *LDAPSource) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil
### GetUserPathTemplate

`func (o *LDAPSource) GetUserPathTemplate() string`

GetUserPathTemplate returns the UserPathTemplate field if non-nil, zero value otherwise.

### GetUserPathTemplateOk

`func (o *LDAPSource) GetUserPathTemplateOk() (*string, bool)`

GetUserPathTemplateOk returns a tuple with the UserPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPathTemplate

`func (o *LDAPSource) SetUserPathTemplate(v string)`

SetUserPathTemplate sets UserPathTemplate field to given value.

### HasUserPathTemplate

`func (o *LDAPSource) HasUserPathTemplate() bool`

HasUserPathTemplate returns a boolean if a field has been set.

### GetIcon

`func (o *LDAPSource) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *LDAPSource) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *LDAPSource) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetServerUri

`func (o *LDAPSource) GetServerUri() string`

GetServerUri returns the ServerUri field if non-nil, zero value otherwise.

### GetServerUriOk

`func (o *LDAPSource) GetServerUriOk() (*string, bool)`

GetServerUriOk returns a tuple with the ServerUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUri

`func (o *LDAPSource) SetServerUri(v string)`

SetServerUri sets ServerUri field to given value.


### GetPeerCertificate

`func (o *LDAPSource) GetPeerCertificate() string`

GetPeerCertificate returns the PeerCertificate field if non-nil, zero value otherwise.

### GetPeerCertificateOk

`func (o *LDAPSource) GetPeerCertificateOk() (*string, bool)`

GetPeerCertificateOk returns a tuple with the PeerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerCertificate

`func (o *LDAPSource) SetPeerCertificate(v string)`

SetPeerCertificate sets PeerCertificate field to given value.

### HasPeerCertificate

`func (o *LDAPSource) HasPeerCertificate() bool`

HasPeerCertificate returns a boolean if a field has been set.

### SetPeerCertificateNil

`func (o *LDAPSource) SetPeerCertificateNil(b bool)`

 SetPeerCertificateNil sets the value for PeerCertificate to be an explicit nil

### UnsetPeerCertificate
`func (o *LDAPSource) UnsetPeerCertificate()`

UnsetPeerCertificate ensures that no value is present for PeerCertificate, not even an explicit nil
### GetClientCertificate

`func (o *LDAPSource) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *LDAPSource) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *LDAPSource) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *LDAPSource) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### SetClientCertificateNil

`func (o *LDAPSource) SetClientCertificateNil(b bool)`

 SetClientCertificateNil sets the value for ClientCertificate to be an explicit nil

### UnsetClientCertificate
`func (o *LDAPSource) UnsetClientCertificate()`

UnsetClientCertificate ensures that no value is present for ClientCertificate, not even an explicit nil
### GetBindCn

`func (o *LDAPSource) GetBindCn() string`

GetBindCn returns the BindCn field if non-nil, zero value otherwise.

### GetBindCnOk

`func (o *LDAPSource) GetBindCnOk() (*string, bool)`

GetBindCnOk returns a tuple with the BindCn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindCn

`func (o *LDAPSource) SetBindCn(v string)`

SetBindCn sets BindCn field to given value.

### HasBindCn

`func (o *LDAPSource) HasBindCn() bool`

HasBindCn returns a boolean if a field has been set.

### GetStartTls

`func (o *LDAPSource) GetStartTls() bool`

GetStartTls returns the StartTls field if non-nil, zero value otherwise.

### GetStartTlsOk

`func (o *LDAPSource) GetStartTlsOk() (*bool, bool)`

GetStartTlsOk returns a tuple with the StartTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTls

`func (o *LDAPSource) SetStartTls(v bool)`

SetStartTls sets StartTls field to given value.

### HasStartTls

`func (o *LDAPSource) HasStartTls() bool`

HasStartTls returns a boolean if a field has been set.

### GetSni

`func (o *LDAPSource) GetSni() bool`

GetSni returns the Sni field if non-nil, zero value otherwise.

### GetSniOk

`func (o *LDAPSource) GetSniOk() (*bool, bool)`

GetSniOk returns a tuple with the Sni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSni

`func (o *LDAPSource) SetSni(v bool)`

SetSni sets Sni field to given value.

### HasSni

`func (o *LDAPSource) HasSni() bool`

HasSni returns a boolean if a field has been set.

### GetBaseDn

`func (o *LDAPSource) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LDAPSource) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LDAPSource) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetAdditionalUserDn

`func (o *LDAPSource) GetAdditionalUserDn() string`

GetAdditionalUserDn returns the AdditionalUserDn field if non-nil, zero value otherwise.

### GetAdditionalUserDnOk

`func (o *LDAPSource) GetAdditionalUserDnOk() (*string, bool)`

GetAdditionalUserDnOk returns a tuple with the AdditionalUserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUserDn

`func (o *LDAPSource) SetAdditionalUserDn(v string)`

SetAdditionalUserDn sets AdditionalUserDn field to given value.

### HasAdditionalUserDn

`func (o *LDAPSource) HasAdditionalUserDn() bool`

HasAdditionalUserDn returns a boolean if a field has been set.

### GetAdditionalGroupDn

`func (o *LDAPSource) GetAdditionalGroupDn() string`

GetAdditionalGroupDn returns the AdditionalGroupDn field if non-nil, zero value otherwise.

### GetAdditionalGroupDnOk

`func (o *LDAPSource) GetAdditionalGroupDnOk() (*string, bool)`

GetAdditionalGroupDnOk returns a tuple with the AdditionalGroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalGroupDn

`func (o *LDAPSource) SetAdditionalGroupDn(v string)`

SetAdditionalGroupDn sets AdditionalGroupDn field to given value.

### HasAdditionalGroupDn

`func (o *LDAPSource) HasAdditionalGroupDn() bool`

HasAdditionalGroupDn returns a boolean if a field has been set.

### GetUserObjectFilter

`func (o *LDAPSource) GetUserObjectFilter() string`

GetUserObjectFilter returns the UserObjectFilter field if non-nil, zero value otherwise.

### GetUserObjectFilterOk

`func (o *LDAPSource) GetUserObjectFilterOk() (*string, bool)`

GetUserObjectFilterOk returns a tuple with the UserObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectFilter

`func (o *LDAPSource) SetUserObjectFilter(v string)`

SetUserObjectFilter sets UserObjectFilter field to given value.

### HasUserObjectFilter

`func (o *LDAPSource) HasUserObjectFilter() bool`

HasUserObjectFilter returns a boolean if a field has been set.

### GetGroupObjectFilter

`func (o *LDAPSource) GetGroupObjectFilter() string`

GetGroupObjectFilter returns the GroupObjectFilter field if non-nil, zero value otherwise.

### GetGroupObjectFilterOk

`func (o *LDAPSource) GetGroupObjectFilterOk() (*string, bool)`

GetGroupObjectFilterOk returns a tuple with the GroupObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectFilter

`func (o *LDAPSource) SetGroupObjectFilter(v string)`

SetGroupObjectFilter sets GroupObjectFilter field to given value.

### HasGroupObjectFilter

`func (o *LDAPSource) HasGroupObjectFilter() bool`

HasGroupObjectFilter returns a boolean if a field has been set.

### GetGroupMembershipField

`func (o *LDAPSource) GetGroupMembershipField() string`

GetGroupMembershipField returns the GroupMembershipField field if non-nil, zero value otherwise.

### GetGroupMembershipFieldOk

`func (o *LDAPSource) GetGroupMembershipFieldOk() (*string, bool)`

GetGroupMembershipFieldOk returns a tuple with the GroupMembershipField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMembershipField

`func (o *LDAPSource) SetGroupMembershipField(v string)`

SetGroupMembershipField sets GroupMembershipField field to given value.

### HasGroupMembershipField

`func (o *LDAPSource) HasGroupMembershipField() bool`

HasGroupMembershipField returns a boolean if a field has been set.

### GetUserMembershipAttribute

`func (o *LDAPSource) GetUserMembershipAttribute() string`

GetUserMembershipAttribute returns the UserMembershipAttribute field if non-nil, zero value otherwise.

### GetUserMembershipAttributeOk

`func (o *LDAPSource) GetUserMembershipAttributeOk() (*string, bool)`

GetUserMembershipAttributeOk returns a tuple with the UserMembershipAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMembershipAttribute

`func (o *LDAPSource) SetUserMembershipAttribute(v string)`

SetUserMembershipAttribute sets UserMembershipAttribute field to given value.

### HasUserMembershipAttribute

`func (o *LDAPSource) HasUserMembershipAttribute() bool`

HasUserMembershipAttribute returns a boolean if a field has been set.

### GetObjectUniquenessField

`func (o *LDAPSource) GetObjectUniquenessField() string`

GetObjectUniquenessField returns the ObjectUniquenessField field if non-nil, zero value otherwise.

### GetObjectUniquenessFieldOk

`func (o *LDAPSource) GetObjectUniquenessFieldOk() (*string, bool)`

GetObjectUniquenessFieldOk returns a tuple with the ObjectUniquenessField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUniquenessField

`func (o *LDAPSource) SetObjectUniquenessField(v string)`

SetObjectUniquenessField sets ObjectUniquenessField field to given value.

### HasObjectUniquenessField

`func (o *LDAPSource) HasObjectUniquenessField() bool`

HasObjectUniquenessField returns a boolean if a field has been set.

### GetPasswordLoginUpdateInternalPassword

`func (o *LDAPSource) GetPasswordLoginUpdateInternalPassword() bool`

GetPasswordLoginUpdateInternalPassword returns the PasswordLoginUpdateInternalPassword field if non-nil, zero value otherwise.

### GetPasswordLoginUpdateInternalPasswordOk

`func (o *LDAPSource) GetPasswordLoginUpdateInternalPasswordOk() (*bool, bool)`

GetPasswordLoginUpdateInternalPasswordOk returns a tuple with the PasswordLoginUpdateInternalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLoginUpdateInternalPassword

`func (o *LDAPSource) SetPasswordLoginUpdateInternalPassword(v bool)`

SetPasswordLoginUpdateInternalPassword sets PasswordLoginUpdateInternalPassword field to given value.

### HasPasswordLoginUpdateInternalPassword

`func (o *LDAPSource) HasPasswordLoginUpdateInternalPassword() bool`

HasPasswordLoginUpdateInternalPassword returns a boolean if a field has been set.

### GetSyncUsers

`func (o *LDAPSource) GetSyncUsers() bool`

GetSyncUsers returns the SyncUsers field if non-nil, zero value otherwise.

### GetSyncUsersOk

`func (o *LDAPSource) GetSyncUsersOk() (*bool, bool)`

GetSyncUsersOk returns a tuple with the SyncUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsers

`func (o *LDAPSource) SetSyncUsers(v bool)`

SetSyncUsers sets SyncUsers field to given value.

### HasSyncUsers

`func (o *LDAPSource) HasSyncUsers() bool`

HasSyncUsers returns a boolean if a field has been set.

### GetSyncUsersPassword

`func (o *LDAPSource) GetSyncUsersPassword() bool`

GetSyncUsersPassword returns the SyncUsersPassword field if non-nil, zero value otherwise.

### GetSyncUsersPasswordOk

`func (o *LDAPSource) GetSyncUsersPasswordOk() (*bool, bool)`

GetSyncUsersPasswordOk returns a tuple with the SyncUsersPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsersPassword

`func (o *LDAPSource) SetSyncUsersPassword(v bool)`

SetSyncUsersPassword sets SyncUsersPassword field to given value.

### HasSyncUsersPassword

`func (o *LDAPSource) HasSyncUsersPassword() bool`

HasSyncUsersPassword returns a boolean if a field has been set.

### GetSyncGroups

`func (o *LDAPSource) GetSyncGroups() bool`

GetSyncGroups returns the SyncGroups field if non-nil, zero value otherwise.

### GetSyncGroupsOk

`func (o *LDAPSource) GetSyncGroupsOk() (*bool, bool)`

GetSyncGroupsOk returns a tuple with the SyncGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncGroups

`func (o *LDAPSource) SetSyncGroups(v bool)`

SetSyncGroups sets SyncGroups field to given value.

### HasSyncGroups

`func (o *LDAPSource) HasSyncGroups() bool`

HasSyncGroups returns a boolean if a field has been set.

### GetSyncParentGroup

`func (o *LDAPSource) GetSyncParentGroup() string`

GetSyncParentGroup returns the SyncParentGroup field if non-nil, zero value otherwise.

### GetSyncParentGroupOk

`func (o *LDAPSource) GetSyncParentGroupOk() (*string, bool)`

GetSyncParentGroupOk returns a tuple with the SyncParentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncParentGroup

`func (o *LDAPSource) SetSyncParentGroup(v string)`

SetSyncParentGroup sets SyncParentGroup field to given value.

### HasSyncParentGroup

`func (o *LDAPSource) HasSyncParentGroup() bool`

HasSyncParentGroup returns a boolean if a field has been set.

### SetSyncParentGroupNil

`func (o *LDAPSource) SetSyncParentGroupNil(b bool)`

 SetSyncParentGroupNil sets the value for SyncParentGroup to be an explicit nil

### UnsetSyncParentGroup
`func (o *LDAPSource) UnsetSyncParentGroup()`

UnsetSyncParentGroup ensures that no value is present for SyncParentGroup, not even an explicit nil
### GetConnectivity

`func (o *LDAPSource) GetConnectivity() map[string]map[string]string`

GetConnectivity returns the Connectivity field if non-nil, zero value otherwise.

### GetConnectivityOk

`func (o *LDAPSource) GetConnectivityOk() (*map[string]map[string]string, bool)`

GetConnectivityOk returns a tuple with the Connectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivity

`func (o *LDAPSource) SetConnectivity(v map[string]map[string]string)`

SetConnectivity sets Connectivity field to given value.


### SetConnectivityNil

`func (o *LDAPSource) SetConnectivityNil(b bool)`

 SetConnectivityNil sets the value for Connectivity to be an explicit nil

### UnsetConnectivity
`func (o *LDAPSource) UnsetConnectivity()`

UnsetConnectivity ensures that no value is present for Connectivity, not even an explicit nil
### GetLookupGroupsFromUser

`func (o *LDAPSource) GetLookupGroupsFromUser() bool`

GetLookupGroupsFromUser returns the LookupGroupsFromUser field if non-nil, zero value otherwise.

### GetLookupGroupsFromUserOk

`func (o *LDAPSource) GetLookupGroupsFromUserOk() (*bool, bool)`

GetLookupGroupsFromUserOk returns a tuple with the LookupGroupsFromUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupGroupsFromUser

`func (o *LDAPSource) SetLookupGroupsFromUser(v bool)`

SetLookupGroupsFromUser sets LookupGroupsFromUser field to given value.

### HasLookupGroupsFromUser

`func (o *LDAPSource) HasLookupGroupsFromUser() bool`

HasLookupGroupsFromUser returns a boolean if a field has been set.

### GetDeleteNotFoundObjects

`func (o *LDAPSource) GetDeleteNotFoundObjects() bool`

GetDeleteNotFoundObjects returns the DeleteNotFoundObjects field if non-nil, zero value otherwise.

### GetDeleteNotFoundObjectsOk

`func (o *LDAPSource) GetDeleteNotFoundObjectsOk() (*bool, bool)`

GetDeleteNotFoundObjectsOk returns a tuple with the DeleteNotFoundObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteNotFoundObjects

`func (o *LDAPSource) SetDeleteNotFoundObjects(v bool)`

SetDeleteNotFoundObjects sets DeleteNotFoundObjects field to given value.

### HasDeleteNotFoundObjects

`func (o *LDAPSource) HasDeleteNotFoundObjects() bool`

HasDeleteNotFoundObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


