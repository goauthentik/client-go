/*
authentik

Making authentication simple.

API version: 2024.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// KerberosSourceRequest Kerberos Source Serializer
type KerberosSourceRequest struct {
	// Source's display Name.
	Name string `json:"name"`
	// Internal source name, used in URLs.
	Slug    string `json:"slug"`
	Enabled *bool  `json:"enabled,omitempty"`
	// Flow to use when authenticating existing users.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow to use when enrolling new users.
	EnrollmentFlow        NullableString    `json:"enrollment_flow,omitempty"`
	UserPropertyMappings  []string          `json:"user_property_mappings,omitempty"`
	GroupPropertyMappings []string          `json:"group_property_mappings,omitempty"`
	PolicyEngineMode      *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// How the source determines if an existing user should be authenticated or a new user enrolled.
	UserMatchingMode *UserMatchingModeEnum `json:"user_matching_mode,omitempty"`
	UserPathTemplate *string               `json:"user_path_template,omitempty"`
	// How the source determines if an existing group should be used or a new group created.
	GroupMatchingMode *GroupMatchingModeEnum `json:"group_matching_mode,omitempty"`
	// Kerberos realm
	Realm string `json:"realm"`
	// Custom krb5.conf to use. Uses the system one by default
	Krb5Conf *string `json:"krb5_conf,omitempty"`
	// KAdmin server type
	KadminType *KadminTypeEnum `json:"kadmin_type,omitempty"`
	// Sync users from Kerberos into authentik
	SyncUsers *bool `json:"sync_users,omitempty"`
	// When a user changes their password, sync it back to Kerberos
	SyncUsersPassword *bool `json:"sync_users_password,omitempty"`
	// Principal to authenticate to kadmin for sync.
	SyncPrincipal *string `json:"sync_principal,omitempty"`
	// Password to authenticate to kadmin for sync
	SyncPassword *string `json:"sync_password,omitempty"`
	// Keytab to authenticate to kadmin for sync. Must be base64-encoded or in the form TYPE:residual
	SyncKeytab *string `json:"sync_keytab,omitempty"`
	// Credentials cache to authenticate to kadmin for sync. Must be in the form TYPE:residual
	SyncCcache *string `json:"sync_ccache,omitempty"`
	// Force the use of a specific server name for SPNEGO. Must be in the form HTTP@hostname
	SpnegoServerName *string `json:"spnego_server_name,omitempty"`
	// SPNEGO keytab base64-encoded or path to keytab in the form FILE:path
	SpnegoKeytab *string `json:"spnego_keytab,omitempty"`
	// Credential cache to use for SPNEGO in form type:residual
	SpnegoCcache *string `json:"spnego_ccache,omitempty"`
	// If enabled, the authentik-stored password will be updated upon login with the Kerberos password backend
	PasswordLoginUpdateInternalPassword *bool `json:"password_login_update_internal_password,omitempty"`
}

// NewKerberosSourceRequest instantiates a new KerberosSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosSourceRequest(name string, slug string, realm string) *KerberosSourceRequest {
	this := KerberosSourceRequest{}
	this.Name = name
	this.Slug = slug
	this.Realm = realm
	return &this
}

// NewKerberosSourceRequestWithDefaults instantiates a new KerberosSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosSourceRequestWithDefaults() *KerberosSourceRequest {
	this := KerberosSourceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *KerberosSourceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KerberosSourceRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *KerberosSourceRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *KerberosSourceRequest) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KerberosSourceRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KerberosSourceRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KerberosSourceRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *KerberosSourceRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *KerberosSourceRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *KerberosSourceRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KerberosSourceRequest) GetEnrollmentFlow() string {
	if o == nil || o.EnrollmentFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KerberosSourceRequest) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *KerberosSourceRequest) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *KerberosSourceRequest) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *KerberosSourceRequest) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetUserPropertyMappings returns the UserPropertyMappings field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetUserPropertyMappings() []string {
	if o == nil || o.UserPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.UserPropertyMappings
}

// GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetUserPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.UserPropertyMappings == nil {
		return nil, false
	}
	return o.UserPropertyMappings, true
}

// HasUserPropertyMappings returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasUserPropertyMappings() bool {
	if o != nil && o.UserPropertyMappings != nil {
		return true
	}

	return false
}

// SetUserPropertyMappings gets a reference to the given []string and assigns it to the UserPropertyMappings field.
func (o *KerberosSourceRequest) SetUserPropertyMappings(v []string) {
	o.UserPropertyMappings = v
}

// GetGroupPropertyMappings returns the GroupPropertyMappings field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetGroupPropertyMappings() []string {
	if o == nil || o.GroupPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.GroupPropertyMappings
}

// GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetGroupPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.GroupPropertyMappings == nil {
		return nil, false
	}
	return o.GroupPropertyMappings, true
}

// HasGroupPropertyMappings returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasGroupPropertyMappings() bool {
	if o != nil && o.GroupPropertyMappings != nil {
		return true
	}

	return false
}

// SetGroupPropertyMappings gets a reference to the given []string and assigns it to the GroupPropertyMappings field.
func (o *KerberosSourceRequest) SetGroupPropertyMappings(v []string) {
	o.GroupPropertyMappings = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *KerberosSourceRequest) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || o.UserMatchingMode == nil {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || o.UserMatchingMode == nil {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasUserMatchingMode() bool {
	if o != nil && o.UserMatchingMode != nil {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *KerberosSourceRequest) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *KerberosSourceRequest) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

// GetGroupMatchingMode returns the GroupMatchingMode field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetGroupMatchingMode() GroupMatchingModeEnum {
	if o == nil || o.GroupMatchingMode == nil {
		var ret GroupMatchingModeEnum
		return ret
	}
	return *o.GroupMatchingMode
}

// GetGroupMatchingModeOk returns a tuple with the GroupMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetGroupMatchingModeOk() (*GroupMatchingModeEnum, bool) {
	if o == nil || o.GroupMatchingMode == nil {
		return nil, false
	}
	return o.GroupMatchingMode, true
}

// HasGroupMatchingMode returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasGroupMatchingMode() bool {
	if o != nil && o.GroupMatchingMode != nil {
		return true
	}

	return false
}

// SetGroupMatchingMode gets a reference to the given GroupMatchingModeEnum and assigns it to the GroupMatchingMode field.
func (o *KerberosSourceRequest) SetGroupMatchingMode(v GroupMatchingModeEnum) {
	o.GroupMatchingMode = &v
}

// GetRealm returns the Realm field value
func (o *KerberosSourceRequest) GetRealm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Realm
}

// GetRealmOk returns a tuple with the Realm field value
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetRealmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Realm, true
}

// SetRealm sets field value
func (o *KerberosSourceRequest) SetRealm(v string) {
	o.Realm = v
}

// GetKrb5Conf returns the Krb5Conf field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetKrb5Conf() string {
	if o == nil || o.Krb5Conf == nil {
		var ret string
		return ret
	}
	return *o.Krb5Conf
}

// GetKrb5ConfOk returns a tuple with the Krb5Conf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetKrb5ConfOk() (*string, bool) {
	if o == nil || o.Krb5Conf == nil {
		return nil, false
	}
	return o.Krb5Conf, true
}

// HasKrb5Conf returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasKrb5Conf() bool {
	if o != nil && o.Krb5Conf != nil {
		return true
	}

	return false
}

// SetKrb5Conf gets a reference to the given string and assigns it to the Krb5Conf field.
func (o *KerberosSourceRequest) SetKrb5Conf(v string) {
	o.Krb5Conf = &v
}

// GetKadminType returns the KadminType field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetKadminType() KadminTypeEnum {
	if o == nil || o.KadminType == nil {
		var ret KadminTypeEnum
		return ret
	}
	return *o.KadminType
}

// GetKadminTypeOk returns a tuple with the KadminType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetKadminTypeOk() (*KadminTypeEnum, bool) {
	if o == nil || o.KadminType == nil {
		return nil, false
	}
	return o.KadminType, true
}

// HasKadminType returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasKadminType() bool {
	if o != nil && o.KadminType != nil {
		return true
	}

	return false
}

// SetKadminType gets a reference to the given KadminTypeEnum and assigns it to the KadminType field.
func (o *KerberosSourceRequest) SetKadminType(v KadminTypeEnum) {
	o.KadminType = &v
}

// GetSyncUsers returns the SyncUsers field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetSyncUsers() bool {
	if o == nil || o.SyncUsers == nil {
		var ret bool
		return ret
	}
	return *o.SyncUsers
}

// GetSyncUsersOk returns a tuple with the SyncUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetSyncUsersOk() (*bool, bool) {
	if o == nil || o.SyncUsers == nil {
		return nil, false
	}
	return o.SyncUsers, true
}

// HasSyncUsers returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasSyncUsers() bool {
	if o != nil && o.SyncUsers != nil {
		return true
	}

	return false
}

// SetSyncUsers gets a reference to the given bool and assigns it to the SyncUsers field.
func (o *KerberosSourceRequest) SetSyncUsers(v bool) {
	o.SyncUsers = &v
}

// GetSyncUsersPassword returns the SyncUsersPassword field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetSyncUsersPassword() bool {
	if o == nil || o.SyncUsersPassword == nil {
		var ret bool
		return ret
	}
	return *o.SyncUsersPassword
}

// GetSyncUsersPasswordOk returns a tuple with the SyncUsersPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetSyncUsersPasswordOk() (*bool, bool) {
	if o == nil || o.SyncUsersPassword == nil {
		return nil, false
	}
	return o.SyncUsersPassword, true
}

// HasSyncUsersPassword returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasSyncUsersPassword() bool {
	if o != nil && o.SyncUsersPassword != nil {
		return true
	}

	return false
}

// SetSyncUsersPassword gets a reference to the given bool and assigns it to the SyncUsersPassword field.
func (o *KerberosSourceRequest) SetSyncUsersPassword(v bool) {
	o.SyncUsersPassword = &v
}

// GetSyncPrincipal returns the SyncPrincipal field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetSyncPrincipal() string {
	if o == nil || o.SyncPrincipal == nil {
		var ret string
		return ret
	}
	return *o.SyncPrincipal
}

// GetSyncPrincipalOk returns a tuple with the SyncPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetSyncPrincipalOk() (*string, bool) {
	if o == nil || o.SyncPrincipal == nil {
		return nil, false
	}
	return o.SyncPrincipal, true
}

// HasSyncPrincipal returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasSyncPrincipal() bool {
	if o != nil && o.SyncPrincipal != nil {
		return true
	}

	return false
}

// SetSyncPrincipal gets a reference to the given string and assigns it to the SyncPrincipal field.
func (o *KerberosSourceRequest) SetSyncPrincipal(v string) {
	o.SyncPrincipal = &v
}

// GetSyncPassword returns the SyncPassword field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetSyncPassword() string {
	if o == nil || o.SyncPassword == nil {
		var ret string
		return ret
	}
	return *o.SyncPassword
}

// GetSyncPasswordOk returns a tuple with the SyncPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetSyncPasswordOk() (*string, bool) {
	if o == nil || o.SyncPassword == nil {
		return nil, false
	}
	return o.SyncPassword, true
}

// HasSyncPassword returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasSyncPassword() bool {
	if o != nil && o.SyncPassword != nil {
		return true
	}

	return false
}

// SetSyncPassword gets a reference to the given string and assigns it to the SyncPassword field.
func (o *KerberosSourceRequest) SetSyncPassword(v string) {
	o.SyncPassword = &v
}

// GetSyncKeytab returns the SyncKeytab field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetSyncKeytab() string {
	if o == nil || o.SyncKeytab == nil {
		var ret string
		return ret
	}
	return *o.SyncKeytab
}

// GetSyncKeytabOk returns a tuple with the SyncKeytab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetSyncKeytabOk() (*string, bool) {
	if o == nil || o.SyncKeytab == nil {
		return nil, false
	}
	return o.SyncKeytab, true
}

// HasSyncKeytab returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasSyncKeytab() bool {
	if o != nil && o.SyncKeytab != nil {
		return true
	}

	return false
}

// SetSyncKeytab gets a reference to the given string and assigns it to the SyncKeytab field.
func (o *KerberosSourceRequest) SetSyncKeytab(v string) {
	o.SyncKeytab = &v
}

// GetSyncCcache returns the SyncCcache field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetSyncCcache() string {
	if o == nil || o.SyncCcache == nil {
		var ret string
		return ret
	}
	return *o.SyncCcache
}

// GetSyncCcacheOk returns a tuple with the SyncCcache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetSyncCcacheOk() (*string, bool) {
	if o == nil || o.SyncCcache == nil {
		return nil, false
	}
	return o.SyncCcache, true
}

// HasSyncCcache returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasSyncCcache() bool {
	if o != nil && o.SyncCcache != nil {
		return true
	}

	return false
}

// SetSyncCcache gets a reference to the given string and assigns it to the SyncCcache field.
func (o *KerberosSourceRequest) SetSyncCcache(v string) {
	o.SyncCcache = &v
}

// GetSpnegoServerName returns the SpnegoServerName field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetSpnegoServerName() string {
	if o == nil || o.SpnegoServerName == nil {
		var ret string
		return ret
	}
	return *o.SpnegoServerName
}

// GetSpnegoServerNameOk returns a tuple with the SpnegoServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetSpnegoServerNameOk() (*string, bool) {
	if o == nil || o.SpnegoServerName == nil {
		return nil, false
	}
	return o.SpnegoServerName, true
}

// HasSpnegoServerName returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasSpnegoServerName() bool {
	if o != nil && o.SpnegoServerName != nil {
		return true
	}

	return false
}

// SetSpnegoServerName gets a reference to the given string and assigns it to the SpnegoServerName field.
func (o *KerberosSourceRequest) SetSpnegoServerName(v string) {
	o.SpnegoServerName = &v
}

// GetSpnegoKeytab returns the SpnegoKeytab field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetSpnegoKeytab() string {
	if o == nil || o.SpnegoKeytab == nil {
		var ret string
		return ret
	}
	return *o.SpnegoKeytab
}

// GetSpnegoKeytabOk returns a tuple with the SpnegoKeytab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetSpnegoKeytabOk() (*string, bool) {
	if o == nil || o.SpnegoKeytab == nil {
		return nil, false
	}
	return o.SpnegoKeytab, true
}

// HasSpnegoKeytab returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasSpnegoKeytab() bool {
	if o != nil && o.SpnegoKeytab != nil {
		return true
	}

	return false
}

// SetSpnegoKeytab gets a reference to the given string and assigns it to the SpnegoKeytab field.
func (o *KerberosSourceRequest) SetSpnegoKeytab(v string) {
	o.SpnegoKeytab = &v
}

// GetSpnegoCcache returns the SpnegoCcache field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetSpnegoCcache() string {
	if o == nil || o.SpnegoCcache == nil {
		var ret string
		return ret
	}
	return *o.SpnegoCcache
}

// GetSpnegoCcacheOk returns a tuple with the SpnegoCcache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetSpnegoCcacheOk() (*string, bool) {
	if o == nil || o.SpnegoCcache == nil {
		return nil, false
	}
	return o.SpnegoCcache, true
}

// HasSpnegoCcache returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasSpnegoCcache() bool {
	if o != nil && o.SpnegoCcache != nil {
		return true
	}

	return false
}

// SetSpnegoCcache gets a reference to the given string and assigns it to the SpnegoCcache field.
func (o *KerberosSourceRequest) SetSpnegoCcache(v string) {
	o.SpnegoCcache = &v
}

// GetPasswordLoginUpdateInternalPassword returns the PasswordLoginUpdateInternalPassword field value if set, zero value otherwise.
func (o *KerberosSourceRequest) GetPasswordLoginUpdateInternalPassword() bool {
	if o == nil || o.PasswordLoginUpdateInternalPassword == nil {
		var ret bool
		return ret
	}
	return *o.PasswordLoginUpdateInternalPassword
}

// GetPasswordLoginUpdateInternalPasswordOk returns a tuple with the PasswordLoginUpdateInternalPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSourceRequest) GetPasswordLoginUpdateInternalPasswordOk() (*bool, bool) {
	if o == nil || o.PasswordLoginUpdateInternalPassword == nil {
		return nil, false
	}
	return o.PasswordLoginUpdateInternalPassword, true
}

// HasPasswordLoginUpdateInternalPassword returns a boolean if a field has been set.
func (o *KerberosSourceRequest) HasPasswordLoginUpdateInternalPassword() bool {
	if o != nil && o.PasswordLoginUpdateInternalPassword != nil {
		return true
	}

	return false
}

// SetPasswordLoginUpdateInternalPassword gets a reference to the given bool and assigns it to the PasswordLoginUpdateInternalPassword field.
func (o *KerberosSourceRequest) SetPasswordLoginUpdateInternalPassword(v bool) {
	o.PasswordLoginUpdateInternalPassword = &v
}

func (o KerberosSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.AuthenticationFlow.IsSet() {
		toSerialize["authentication_flow"] = o.AuthenticationFlow.Get()
	}
	if o.EnrollmentFlow.IsSet() {
		toSerialize["enrollment_flow"] = o.EnrollmentFlow.Get()
	}
	if o.UserPropertyMappings != nil {
		toSerialize["user_property_mappings"] = o.UserPropertyMappings
	}
	if o.GroupPropertyMappings != nil {
		toSerialize["group_property_mappings"] = o.GroupPropertyMappings
	}
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.UserMatchingMode != nil {
		toSerialize["user_matching_mode"] = o.UserMatchingMode
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	if o.GroupMatchingMode != nil {
		toSerialize["group_matching_mode"] = o.GroupMatchingMode
	}
	if true {
		toSerialize["realm"] = o.Realm
	}
	if o.Krb5Conf != nil {
		toSerialize["krb5_conf"] = o.Krb5Conf
	}
	if o.KadminType != nil {
		toSerialize["kadmin_type"] = o.KadminType
	}
	if o.SyncUsers != nil {
		toSerialize["sync_users"] = o.SyncUsers
	}
	if o.SyncUsersPassword != nil {
		toSerialize["sync_users_password"] = o.SyncUsersPassword
	}
	if o.SyncPrincipal != nil {
		toSerialize["sync_principal"] = o.SyncPrincipal
	}
	if o.SyncPassword != nil {
		toSerialize["sync_password"] = o.SyncPassword
	}
	if o.SyncKeytab != nil {
		toSerialize["sync_keytab"] = o.SyncKeytab
	}
	if o.SyncCcache != nil {
		toSerialize["sync_ccache"] = o.SyncCcache
	}
	if o.SpnegoServerName != nil {
		toSerialize["spnego_server_name"] = o.SpnegoServerName
	}
	if o.SpnegoKeytab != nil {
		toSerialize["spnego_keytab"] = o.SpnegoKeytab
	}
	if o.SpnegoCcache != nil {
		toSerialize["spnego_ccache"] = o.SpnegoCcache
	}
	if o.PasswordLoginUpdateInternalPassword != nil {
		toSerialize["password_login_update_internal_password"] = o.PasswordLoginUpdateInternalPassword
	}
	return json.Marshal(toSerialize)
}

type NullableKerberosSourceRequest struct {
	value *KerberosSourceRequest
	isSet bool
}

func (v NullableKerberosSourceRequest) Get() *KerberosSourceRequest {
	return v.value
}

func (v *NullableKerberosSourceRequest) Set(val *KerberosSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosSourceRequest(val *KerberosSourceRequest) *NullableKerberosSourceRequest {
	return &NullableKerberosSourceRequest{value: val, isSet: true}
}

func (v NullableKerberosSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
