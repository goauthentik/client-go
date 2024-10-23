/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// KerberosSource Kerberos Source Serializer
type KerberosSource struct {
	Pk string `json:"pk"`
	// Source's display Name.
	Name string `json:"name"`
	// Internal source name, used in URLs.
	Slug    string `json:"slug"`
	Enabled *bool  `json:"enabled,omitempty"`
	// Flow to use when authenticating existing users.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow to use when enrolling new users.
	EnrollmentFlow        NullableString `json:"enrollment_flow,omitempty"`
	UserPropertyMappings  []string       `json:"user_property_mappings,omitempty"`
	GroupPropertyMappings []string       `json:"group_property_mappings,omitempty"`
	// Get object component so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName    string            `json:"meta_model_name"`
	PolicyEngineMode *PolicyEngineMode `json:"policy_engine_mode,omitempty"`
	// How the source determines if an existing user should be authenticated or a new user enrolled.
	UserMatchingMode *UserMatchingModeEnum `json:"user_matching_mode,omitempty"`
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed          NullableString `json:"managed"`
	UserPathTemplate *string        `json:"user_path_template,omitempty"`
	Icon             string         `json:"icon"`
	// How the source determines if an existing group should be used or a new group created.
	GroupMatchingMode *GroupMatchingModeEnum `json:"group_matching_mode,omitempty"`
	// Kerberos realm
	Realm string `json:"realm"`
	// Custom krb5.conf to use. Uses the system one by default
	Krb5Conf *string `json:"krb5_conf,omitempty"`
	// Sync users from Kerberos into authentik
	SyncUsers *bool `json:"sync_users,omitempty"`
	// When a user changes their password, sync it back to Kerberos
	SyncUsersPassword *bool `json:"sync_users_password,omitempty"`
	// Principal to authenticate to kadmin for sync.
	SyncPrincipal *string `json:"sync_principal,omitempty"`
	// Credentials cache to authenticate to kadmin for sync. Must be in the form TYPE:residual
	SyncCcache *string `json:"sync_ccache,omitempty"`
	// Get cached source connectivity
	Connectivity map[string]string `json:"connectivity"`
	// Force the use of a specific server name for SPNEGO
	SpnegoServerName *string `json:"spnego_server_name,omitempty"`
	// Credential cache to use for SPNEGO in form type:residual
	SpnegoCcache *string `json:"spnego_ccache,omitempty"`
	// If enabled, the authentik-stored password will be updated upon login with the Kerberos password backend
	PasswordLoginUpdateInternalPassword *bool `json:"password_login_update_internal_password,omitempty"`
}

// NewKerberosSource instantiates a new KerberosSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosSource(pk string, name string, slug string, component string, verboseName string, verboseNamePlural string, metaModelName string, managed NullableString, icon string, realm string, connectivity map[string]string) *KerberosSource {
	this := KerberosSource{}
	this.Pk = pk
	this.Name = name
	this.Slug = slug
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Managed = managed
	this.Icon = icon
	this.Realm = realm
	this.Connectivity = connectivity
	return &this
}

// NewKerberosSourceWithDefaults instantiates a new KerberosSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosSourceWithDefaults() *KerberosSource {
	this := KerberosSource{}
	return &this
}

// GetPk returns the Pk field value
func (o *KerberosSource) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *KerberosSource) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *KerberosSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KerberosSource) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *KerberosSource) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *KerberosSource) SetSlug(v string) {
	o.Slug = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KerberosSource) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KerberosSource) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KerberosSource) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KerberosSource) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KerberosSource) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *KerberosSource) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *KerberosSource) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *KerberosSource) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *KerberosSource) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetEnrollmentFlow returns the EnrollmentFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KerberosSource) GetEnrollmentFlow() string {
	if o == nil || o.EnrollmentFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentFlow.Get()
}

// GetEnrollmentFlowOk returns a tuple with the EnrollmentFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KerberosSource) GetEnrollmentFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollmentFlow.Get(), o.EnrollmentFlow.IsSet()
}

// HasEnrollmentFlow returns a boolean if a field has been set.
func (o *KerberosSource) HasEnrollmentFlow() bool {
	if o != nil && o.EnrollmentFlow.IsSet() {
		return true
	}

	return false
}

// SetEnrollmentFlow gets a reference to the given NullableString and assigns it to the EnrollmentFlow field.
func (o *KerberosSource) SetEnrollmentFlow(v string) {
	o.EnrollmentFlow.Set(&v)
}

// SetEnrollmentFlowNil sets the value for EnrollmentFlow to be an explicit nil
func (o *KerberosSource) SetEnrollmentFlowNil() {
	o.EnrollmentFlow.Set(nil)
}

// UnsetEnrollmentFlow ensures that no value is present for EnrollmentFlow, not even an explicit nil
func (o *KerberosSource) UnsetEnrollmentFlow() {
	o.EnrollmentFlow.Unset()
}

// GetUserPropertyMappings returns the UserPropertyMappings field value if set, zero value otherwise.
func (o *KerberosSource) GetUserPropertyMappings() []string {
	if o == nil || o.UserPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.UserPropertyMappings
}

// GetUserPropertyMappingsOk returns a tuple with the UserPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetUserPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.UserPropertyMappings == nil {
		return nil, false
	}
	return o.UserPropertyMappings, true
}

// HasUserPropertyMappings returns a boolean if a field has been set.
func (o *KerberosSource) HasUserPropertyMappings() bool {
	if o != nil && o.UserPropertyMappings != nil {
		return true
	}

	return false
}

// SetUserPropertyMappings gets a reference to the given []string and assigns it to the UserPropertyMappings field.
func (o *KerberosSource) SetUserPropertyMappings(v []string) {
	o.UserPropertyMappings = v
}

// GetGroupPropertyMappings returns the GroupPropertyMappings field value if set, zero value otherwise.
func (o *KerberosSource) GetGroupPropertyMappings() []string {
	if o == nil || o.GroupPropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.GroupPropertyMappings
}

// GetGroupPropertyMappingsOk returns a tuple with the GroupPropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetGroupPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.GroupPropertyMappings == nil {
		return nil, false
	}
	return o.GroupPropertyMappings, true
}

// HasGroupPropertyMappings returns a boolean if a field has been set.
func (o *KerberosSource) HasGroupPropertyMappings() bool {
	if o != nil && o.GroupPropertyMappings != nil {
		return true
	}

	return false
}

// SetGroupPropertyMappings gets a reference to the given []string and assigns it to the GroupPropertyMappings field.
func (o *KerberosSource) SetGroupPropertyMappings(v []string) {
	o.GroupPropertyMappings = v
}

// GetComponent returns the Component field value
func (o *KerberosSource) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *KerberosSource) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *KerberosSource) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *KerberosSource) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *KerberosSource) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *KerberosSource) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *KerberosSource) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *KerberosSource) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetPolicyEngineMode returns the PolicyEngineMode field value if set, zero value otherwise.
func (o *KerberosSource) GetPolicyEngineMode() PolicyEngineMode {
	if o == nil || o.PolicyEngineMode == nil {
		var ret PolicyEngineMode
		return ret
	}
	return *o.PolicyEngineMode
}

// GetPolicyEngineModeOk returns a tuple with the PolicyEngineMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetPolicyEngineModeOk() (*PolicyEngineMode, bool) {
	if o == nil || o.PolicyEngineMode == nil {
		return nil, false
	}
	return o.PolicyEngineMode, true
}

// HasPolicyEngineMode returns a boolean if a field has been set.
func (o *KerberosSource) HasPolicyEngineMode() bool {
	if o != nil && o.PolicyEngineMode != nil {
		return true
	}

	return false
}

// SetPolicyEngineMode gets a reference to the given PolicyEngineMode and assigns it to the PolicyEngineMode field.
func (o *KerberosSource) SetPolicyEngineMode(v PolicyEngineMode) {
	o.PolicyEngineMode = &v
}

// GetUserMatchingMode returns the UserMatchingMode field value if set, zero value otherwise.
func (o *KerberosSource) GetUserMatchingMode() UserMatchingModeEnum {
	if o == nil || o.UserMatchingMode == nil {
		var ret UserMatchingModeEnum
		return ret
	}
	return *o.UserMatchingMode
}

// GetUserMatchingModeOk returns a tuple with the UserMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetUserMatchingModeOk() (*UserMatchingModeEnum, bool) {
	if o == nil || o.UserMatchingMode == nil {
		return nil, false
	}
	return o.UserMatchingMode, true
}

// HasUserMatchingMode returns a boolean if a field has been set.
func (o *KerberosSource) HasUserMatchingMode() bool {
	if o != nil && o.UserMatchingMode != nil {
		return true
	}

	return false
}

// SetUserMatchingMode gets a reference to the given UserMatchingModeEnum and assigns it to the UserMatchingMode field.
func (o *KerberosSource) SetUserMatchingMode(v UserMatchingModeEnum) {
	o.UserMatchingMode = &v
}

// GetManaged returns the Managed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KerberosSource) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}

	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KerberosSource) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// SetManaged sets field value
func (o *KerberosSource) SetManaged(v string) {
	o.Managed.Set(&v)
}

// GetUserPathTemplate returns the UserPathTemplate field value if set, zero value otherwise.
func (o *KerberosSource) GetUserPathTemplate() string {
	if o == nil || o.UserPathTemplate == nil {
		var ret string
		return ret
	}
	return *o.UserPathTemplate
}

// GetUserPathTemplateOk returns a tuple with the UserPathTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetUserPathTemplateOk() (*string, bool) {
	if o == nil || o.UserPathTemplate == nil {
		return nil, false
	}
	return o.UserPathTemplate, true
}

// HasUserPathTemplate returns a boolean if a field has been set.
func (o *KerberosSource) HasUserPathTemplate() bool {
	if o != nil && o.UserPathTemplate != nil {
		return true
	}

	return false
}

// SetUserPathTemplate gets a reference to the given string and assigns it to the UserPathTemplate field.
func (o *KerberosSource) SetUserPathTemplate(v string) {
	o.UserPathTemplate = &v
}

// GetIcon returns the Icon field value
func (o *KerberosSource) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *KerberosSource) SetIcon(v string) {
	o.Icon = v
}

// GetGroupMatchingMode returns the GroupMatchingMode field value if set, zero value otherwise.
func (o *KerberosSource) GetGroupMatchingMode() GroupMatchingModeEnum {
	if o == nil || o.GroupMatchingMode == nil {
		var ret GroupMatchingModeEnum
		return ret
	}
	return *o.GroupMatchingMode
}

// GetGroupMatchingModeOk returns a tuple with the GroupMatchingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetGroupMatchingModeOk() (*GroupMatchingModeEnum, bool) {
	if o == nil || o.GroupMatchingMode == nil {
		return nil, false
	}
	return o.GroupMatchingMode, true
}

// HasGroupMatchingMode returns a boolean if a field has been set.
func (o *KerberosSource) HasGroupMatchingMode() bool {
	if o != nil && o.GroupMatchingMode != nil {
		return true
	}

	return false
}

// SetGroupMatchingMode gets a reference to the given GroupMatchingModeEnum and assigns it to the GroupMatchingMode field.
func (o *KerberosSource) SetGroupMatchingMode(v GroupMatchingModeEnum) {
	o.GroupMatchingMode = &v
}

// GetRealm returns the Realm field value
func (o *KerberosSource) GetRealm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Realm
}

// GetRealmOk returns a tuple with the Realm field value
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetRealmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Realm, true
}

// SetRealm sets field value
func (o *KerberosSource) SetRealm(v string) {
	o.Realm = v
}

// GetKrb5Conf returns the Krb5Conf field value if set, zero value otherwise.
func (o *KerberosSource) GetKrb5Conf() string {
	if o == nil || o.Krb5Conf == nil {
		var ret string
		return ret
	}
	return *o.Krb5Conf
}

// GetKrb5ConfOk returns a tuple with the Krb5Conf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetKrb5ConfOk() (*string, bool) {
	if o == nil || o.Krb5Conf == nil {
		return nil, false
	}
	return o.Krb5Conf, true
}

// HasKrb5Conf returns a boolean if a field has been set.
func (o *KerberosSource) HasKrb5Conf() bool {
	if o != nil && o.Krb5Conf != nil {
		return true
	}

	return false
}

// SetKrb5Conf gets a reference to the given string and assigns it to the Krb5Conf field.
func (o *KerberosSource) SetKrb5Conf(v string) {
	o.Krb5Conf = &v
}

// GetSyncUsers returns the SyncUsers field value if set, zero value otherwise.
func (o *KerberosSource) GetSyncUsers() bool {
	if o == nil || o.SyncUsers == nil {
		var ret bool
		return ret
	}
	return *o.SyncUsers
}

// GetSyncUsersOk returns a tuple with the SyncUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetSyncUsersOk() (*bool, bool) {
	if o == nil || o.SyncUsers == nil {
		return nil, false
	}
	return o.SyncUsers, true
}

// HasSyncUsers returns a boolean if a field has been set.
func (o *KerberosSource) HasSyncUsers() bool {
	if o != nil && o.SyncUsers != nil {
		return true
	}

	return false
}

// SetSyncUsers gets a reference to the given bool and assigns it to the SyncUsers field.
func (o *KerberosSource) SetSyncUsers(v bool) {
	o.SyncUsers = &v
}

// GetSyncUsersPassword returns the SyncUsersPassword field value if set, zero value otherwise.
func (o *KerberosSource) GetSyncUsersPassword() bool {
	if o == nil || o.SyncUsersPassword == nil {
		var ret bool
		return ret
	}
	return *o.SyncUsersPassword
}

// GetSyncUsersPasswordOk returns a tuple with the SyncUsersPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetSyncUsersPasswordOk() (*bool, bool) {
	if o == nil || o.SyncUsersPassword == nil {
		return nil, false
	}
	return o.SyncUsersPassword, true
}

// HasSyncUsersPassword returns a boolean if a field has been set.
func (o *KerberosSource) HasSyncUsersPassword() bool {
	if o != nil && o.SyncUsersPassword != nil {
		return true
	}

	return false
}

// SetSyncUsersPassword gets a reference to the given bool and assigns it to the SyncUsersPassword field.
func (o *KerberosSource) SetSyncUsersPassword(v bool) {
	o.SyncUsersPassword = &v
}

// GetSyncPrincipal returns the SyncPrincipal field value if set, zero value otherwise.
func (o *KerberosSource) GetSyncPrincipal() string {
	if o == nil || o.SyncPrincipal == nil {
		var ret string
		return ret
	}
	return *o.SyncPrincipal
}

// GetSyncPrincipalOk returns a tuple with the SyncPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetSyncPrincipalOk() (*string, bool) {
	if o == nil || o.SyncPrincipal == nil {
		return nil, false
	}
	return o.SyncPrincipal, true
}

// HasSyncPrincipal returns a boolean if a field has been set.
func (o *KerberosSource) HasSyncPrincipal() bool {
	if o != nil && o.SyncPrincipal != nil {
		return true
	}

	return false
}

// SetSyncPrincipal gets a reference to the given string and assigns it to the SyncPrincipal field.
func (o *KerberosSource) SetSyncPrincipal(v string) {
	o.SyncPrincipal = &v
}

// GetSyncCcache returns the SyncCcache field value if set, zero value otherwise.
func (o *KerberosSource) GetSyncCcache() string {
	if o == nil || o.SyncCcache == nil {
		var ret string
		return ret
	}
	return *o.SyncCcache
}

// GetSyncCcacheOk returns a tuple with the SyncCcache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetSyncCcacheOk() (*string, bool) {
	if o == nil || o.SyncCcache == nil {
		return nil, false
	}
	return o.SyncCcache, true
}

// HasSyncCcache returns a boolean if a field has been set.
func (o *KerberosSource) HasSyncCcache() bool {
	if o != nil && o.SyncCcache != nil {
		return true
	}

	return false
}

// SetSyncCcache gets a reference to the given string and assigns it to the SyncCcache field.
func (o *KerberosSource) SetSyncCcache(v string) {
	o.SyncCcache = &v
}

// GetConnectivity returns the Connectivity field value
// If the value is explicit nil, the zero value for map[string]string will be returned
func (o *KerberosSource) GetConnectivity() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Connectivity
}

// GetConnectivityOk returns a tuple with the Connectivity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KerberosSource) GetConnectivityOk() (*map[string]string, bool) {
	if o == nil || o.Connectivity == nil {
		return nil, false
	}
	return &o.Connectivity, true
}

// SetConnectivity sets field value
func (o *KerberosSource) SetConnectivity(v map[string]string) {
	o.Connectivity = v
}

// GetSpnegoServerName returns the SpnegoServerName field value if set, zero value otherwise.
func (o *KerberosSource) GetSpnegoServerName() string {
	if o == nil || o.SpnegoServerName == nil {
		var ret string
		return ret
	}
	return *o.SpnegoServerName
}

// GetSpnegoServerNameOk returns a tuple with the SpnegoServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetSpnegoServerNameOk() (*string, bool) {
	if o == nil || o.SpnegoServerName == nil {
		return nil, false
	}
	return o.SpnegoServerName, true
}

// HasSpnegoServerName returns a boolean if a field has been set.
func (o *KerberosSource) HasSpnegoServerName() bool {
	if o != nil && o.SpnegoServerName != nil {
		return true
	}

	return false
}

// SetSpnegoServerName gets a reference to the given string and assigns it to the SpnegoServerName field.
func (o *KerberosSource) SetSpnegoServerName(v string) {
	o.SpnegoServerName = &v
}

// GetSpnegoCcache returns the SpnegoCcache field value if set, zero value otherwise.
func (o *KerberosSource) GetSpnegoCcache() string {
	if o == nil || o.SpnegoCcache == nil {
		var ret string
		return ret
	}
	return *o.SpnegoCcache
}

// GetSpnegoCcacheOk returns a tuple with the SpnegoCcache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetSpnegoCcacheOk() (*string, bool) {
	if o == nil || o.SpnegoCcache == nil {
		return nil, false
	}
	return o.SpnegoCcache, true
}

// HasSpnegoCcache returns a boolean if a field has been set.
func (o *KerberosSource) HasSpnegoCcache() bool {
	if o != nil && o.SpnegoCcache != nil {
		return true
	}

	return false
}

// SetSpnegoCcache gets a reference to the given string and assigns it to the SpnegoCcache field.
func (o *KerberosSource) SetSpnegoCcache(v string) {
	o.SpnegoCcache = &v
}

// GetPasswordLoginUpdateInternalPassword returns the PasswordLoginUpdateInternalPassword field value if set, zero value otherwise.
func (o *KerberosSource) GetPasswordLoginUpdateInternalPassword() bool {
	if o == nil || o.PasswordLoginUpdateInternalPassword == nil {
		var ret bool
		return ret
	}
	return *o.PasswordLoginUpdateInternalPassword
}

// GetPasswordLoginUpdateInternalPasswordOk returns a tuple with the PasswordLoginUpdateInternalPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosSource) GetPasswordLoginUpdateInternalPasswordOk() (*bool, bool) {
	if o == nil || o.PasswordLoginUpdateInternalPassword == nil {
		return nil, false
	}
	return o.PasswordLoginUpdateInternalPassword, true
}

// HasPasswordLoginUpdateInternalPassword returns a boolean if a field has been set.
func (o *KerberosSource) HasPasswordLoginUpdateInternalPassword() bool {
	if o != nil && o.PasswordLoginUpdateInternalPassword != nil {
		return true
	}

	return false
}

// SetPasswordLoginUpdateInternalPassword gets a reference to the given bool and assigns it to the PasswordLoginUpdateInternalPassword field.
func (o *KerberosSource) SetPasswordLoginUpdateInternalPassword(v bool) {
	o.PasswordLoginUpdateInternalPassword = &v
}

func (o KerberosSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
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
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if o.PolicyEngineMode != nil {
		toSerialize["policy_engine_mode"] = o.PolicyEngineMode
	}
	if o.UserMatchingMode != nil {
		toSerialize["user_matching_mode"] = o.UserMatchingMode
	}
	if true {
		toSerialize["managed"] = o.Managed.Get()
	}
	if o.UserPathTemplate != nil {
		toSerialize["user_path_template"] = o.UserPathTemplate
	}
	if true {
		toSerialize["icon"] = o.Icon
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
	if o.SyncUsers != nil {
		toSerialize["sync_users"] = o.SyncUsers
	}
	if o.SyncUsersPassword != nil {
		toSerialize["sync_users_password"] = o.SyncUsersPassword
	}
	if o.SyncPrincipal != nil {
		toSerialize["sync_principal"] = o.SyncPrincipal
	}
	if o.SyncCcache != nil {
		toSerialize["sync_ccache"] = o.SyncCcache
	}
	if o.Connectivity != nil {
		toSerialize["connectivity"] = o.Connectivity
	}
	if o.SpnegoServerName != nil {
		toSerialize["spnego_server_name"] = o.SpnegoServerName
	}
	if o.SpnegoCcache != nil {
		toSerialize["spnego_ccache"] = o.SpnegoCcache
	}
	if o.PasswordLoginUpdateInternalPassword != nil {
		toSerialize["password_login_update_internal_password"] = o.PasswordLoginUpdateInternalPassword
	}
	return json.Marshal(toSerialize)
}

type NullableKerberosSource struct {
	value *KerberosSource
	isSet bool
}

func (v NullableKerberosSource) Get() *KerberosSource {
	return v.value
}

func (v *NullableKerberosSource) Set(val *KerberosSource) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosSource) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosSource(val *KerberosSource) *NullableKerberosSource {
	return &NullableKerberosSource{value: val, isSet: true}
}

func (v NullableKerberosSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
