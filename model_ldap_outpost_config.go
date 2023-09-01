/*
authentik

Making authentication simple.

API version: 2023.8.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LDAPOutpostConfig LDAPProvider Serializer
type LDAPOutpostConfig struct {
	Pk   int32  `json:"pk"`
	Name string `json:"name"`
	// DN under which objects are accessible.
	BaseDn       *string `json:"base_dn,omitempty"`
	BindFlowSlug string  `json:"bind_flow_slug"`
	// Prioritise backchannel slug over direct application slug
	ApplicationSlug string `json:"application_slug"`
	// Users in this group can do search queries. If not set, every user can execute search queries.
	SearchGroup   NullableString `json:"search_group,omitempty"`
	Certificate   NullableString `json:"certificate,omitempty"`
	TlsServerName *string        `json:"tls_server_name,omitempty"`
	// The start for uidNumbers, this number is added to the user.pk to make sure that the numbers aren't too low for POSIX users. Default is 2000 to ensure that we don't collide with local users uidNumber
	UidStartNumber *int32 `json:"uid_start_number,omitempty"`
	// The start for gidNumbers, this number is added to a number generated from the group.pk to make sure that the numbers aren't too low for POSIX groups. Default is 4000 to ensure that we don't collide with local groups or users primary groups gidNumber
	GidStartNumber *int32             `json:"gid_start_number,omitempty"`
	SearchMode     *LDAPAPIAccessMode `json:"search_mode,omitempty"`
	BindMode       *LDAPAPIAccessMode `json:"bind_mode,omitempty"`
	// When enabled, code-based multi-factor authentication can be used by appending a semicolon and the TOTP code to the password. This should only be enabled if all users that will bind to this provider have a TOTP device configured, as otherwise a password may incorrectly be rejected if it contains a semicolon.
	MfaSupport *bool `json:"mfa_support,omitempty"`
}

// NewLDAPOutpostConfig instantiates a new LDAPOutpostConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPOutpostConfig(pk int32, name string, bindFlowSlug string, applicationSlug string) *LDAPOutpostConfig {
	this := LDAPOutpostConfig{}
	this.Pk = pk
	this.Name = name
	this.BindFlowSlug = bindFlowSlug
	this.ApplicationSlug = applicationSlug
	return &this
}

// NewLDAPOutpostConfigWithDefaults instantiates a new LDAPOutpostConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPOutpostConfigWithDefaults() *LDAPOutpostConfig {
	this := LDAPOutpostConfig{}
	return &this
}

// GetPk returns the Pk field value
func (o *LDAPOutpostConfig) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *LDAPOutpostConfig) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *LDAPOutpostConfig) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *LDAPOutpostConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LDAPOutpostConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LDAPOutpostConfig) SetName(v string) {
	o.Name = v
}

// GetBaseDn returns the BaseDn field value if set, zero value otherwise.
func (o *LDAPOutpostConfig) GetBaseDn() string {
	if o == nil || o.BaseDn == nil {
		var ret string
		return ret
	}
	return *o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPOutpostConfig) GetBaseDnOk() (*string, bool) {
	if o == nil || o.BaseDn == nil {
		return nil, false
	}
	return o.BaseDn, true
}

// HasBaseDn returns a boolean if a field has been set.
func (o *LDAPOutpostConfig) HasBaseDn() bool {
	if o != nil && o.BaseDn != nil {
		return true
	}

	return false
}

// SetBaseDn gets a reference to the given string and assigns it to the BaseDn field.
func (o *LDAPOutpostConfig) SetBaseDn(v string) {
	o.BaseDn = &v
}

// GetBindFlowSlug returns the BindFlowSlug field value
func (o *LDAPOutpostConfig) GetBindFlowSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BindFlowSlug
}

// GetBindFlowSlugOk returns a tuple with the BindFlowSlug field value
// and a boolean to check if the value has been set.
func (o *LDAPOutpostConfig) GetBindFlowSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BindFlowSlug, true
}

// SetBindFlowSlug sets field value
func (o *LDAPOutpostConfig) SetBindFlowSlug(v string) {
	o.BindFlowSlug = v
}

// GetApplicationSlug returns the ApplicationSlug field value
func (o *LDAPOutpostConfig) GetApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationSlug
}

// GetApplicationSlugOk returns a tuple with the ApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *LDAPOutpostConfig) GetApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationSlug, true
}

// SetApplicationSlug sets field value
func (o *LDAPOutpostConfig) SetApplicationSlug(v string) {
	o.ApplicationSlug = v
}

// GetSearchGroup returns the SearchGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPOutpostConfig) GetSearchGroup() string {
	if o == nil || o.SearchGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.SearchGroup.Get()
}

// GetSearchGroupOk returns a tuple with the SearchGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPOutpostConfig) GetSearchGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchGroup.Get(), o.SearchGroup.IsSet()
}

// HasSearchGroup returns a boolean if a field has been set.
func (o *LDAPOutpostConfig) HasSearchGroup() bool {
	if o != nil && o.SearchGroup.IsSet() {
		return true
	}

	return false
}

// SetSearchGroup gets a reference to the given NullableString and assigns it to the SearchGroup field.
func (o *LDAPOutpostConfig) SetSearchGroup(v string) {
	o.SearchGroup.Set(&v)
}

// SetSearchGroupNil sets the value for SearchGroup to be an explicit nil
func (o *LDAPOutpostConfig) SetSearchGroupNil() {
	o.SearchGroup.Set(nil)
}

// UnsetSearchGroup ensures that no value is present for SearchGroup, not even an explicit nil
func (o *LDAPOutpostConfig) UnsetSearchGroup() {
	o.SearchGroup.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPOutpostConfig) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPOutpostConfig) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *LDAPOutpostConfig) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *LDAPOutpostConfig) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *LDAPOutpostConfig) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *LDAPOutpostConfig) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetTlsServerName returns the TlsServerName field value if set, zero value otherwise.
func (o *LDAPOutpostConfig) GetTlsServerName() string {
	if o == nil || o.TlsServerName == nil {
		var ret string
		return ret
	}
	return *o.TlsServerName
}

// GetTlsServerNameOk returns a tuple with the TlsServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPOutpostConfig) GetTlsServerNameOk() (*string, bool) {
	if o == nil || o.TlsServerName == nil {
		return nil, false
	}
	return o.TlsServerName, true
}

// HasTlsServerName returns a boolean if a field has been set.
func (o *LDAPOutpostConfig) HasTlsServerName() bool {
	if o != nil && o.TlsServerName != nil {
		return true
	}

	return false
}

// SetTlsServerName gets a reference to the given string and assigns it to the TlsServerName field.
func (o *LDAPOutpostConfig) SetTlsServerName(v string) {
	o.TlsServerName = &v
}

// GetUidStartNumber returns the UidStartNumber field value if set, zero value otherwise.
func (o *LDAPOutpostConfig) GetUidStartNumber() int32 {
	if o == nil || o.UidStartNumber == nil {
		var ret int32
		return ret
	}
	return *o.UidStartNumber
}

// GetUidStartNumberOk returns a tuple with the UidStartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPOutpostConfig) GetUidStartNumberOk() (*int32, bool) {
	if o == nil || o.UidStartNumber == nil {
		return nil, false
	}
	return o.UidStartNumber, true
}

// HasUidStartNumber returns a boolean if a field has been set.
func (o *LDAPOutpostConfig) HasUidStartNumber() bool {
	if o != nil && o.UidStartNumber != nil {
		return true
	}

	return false
}

// SetUidStartNumber gets a reference to the given int32 and assigns it to the UidStartNumber field.
func (o *LDAPOutpostConfig) SetUidStartNumber(v int32) {
	o.UidStartNumber = &v
}

// GetGidStartNumber returns the GidStartNumber field value if set, zero value otherwise.
func (o *LDAPOutpostConfig) GetGidStartNumber() int32 {
	if o == nil || o.GidStartNumber == nil {
		var ret int32
		return ret
	}
	return *o.GidStartNumber
}

// GetGidStartNumberOk returns a tuple with the GidStartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPOutpostConfig) GetGidStartNumberOk() (*int32, bool) {
	if o == nil || o.GidStartNumber == nil {
		return nil, false
	}
	return o.GidStartNumber, true
}

// HasGidStartNumber returns a boolean if a field has been set.
func (o *LDAPOutpostConfig) HasGidStartNumber() bool {
	if o != nil && o.GidStartNumber != nil {
		return true
	}

	return false
}

// SetGidStartNumber gets a reference to the given int32 and assigns it to the GidStartNumber field.
func (o *LDAPOutpostConfig) SetGidStartNumber(v int32) {
	o.GidStartNumber = &v
}

// GetSearchMode returns the SearchMode field value if set, zero value otherwise.
func (o *LDAPOutpostConfig) GetSearchMode() LDAPAPIAccessMode {
	if o == nil || o.SearchMode == nil {
		var ret LDAPAPIAccessMode
		return ret
	}
	return *o.SearchMode
}

// GetSearchModeOk returns a tuple with the SearchMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPOutpostConfig) GetSearchModeOk() (*LDAPAPIAccessMode, bool) {
	if o == nil || o.SearchMode == nil {
		return nil, false
	}
	return o.SearchMode, true
}

// HasSearchMode returns a boolean if a field has been set.
func (o *LDAPOutpostConfig) HasSearchMode() bool {
	if o != nil && o.SearchMode != nil {
		return true
	}

	return false
}

// SetSearchMode gets a reference to the given LDAPAPIAccessMode and assigns it to the SearchMode field.
func (o *LDAPOutpostConfig) SetSearchMode(v LDAPAPIAccessMode) {
	o.SearchMode = &v
}

// GetBindMode returns the BindMode field value if set, zero value otherwise.
func (o *LDAPOutpostConfig) GetBindMode() LDAPAPIAccessMode {
	if o == nil || o.BindMode == nil {
		var ret LDAPAPIAccessMode
		return ret
	}
	return *o.BindMode
}

// GetBindModeOk returns a tuple with the BindMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPOutpostConfig) GetBindModeOk() (*LDAPAPIAccessMode, bool) {
	if o == nil || o.BindMode == nil {
		return nil, false
	}
	return o.BindMode, true
}

// HasBindMode returns a boolean if a field has been set.
func (o *LDAPOutpostConfig) HasBindMode() bool {
	if o != nil && o.BindMode != nil {
		return true
	}

	return false
}

// SetBindMode gets a reference to the given LDAPAPIAccessMode and assigns it to the BindMode field.
func (o *LDAPOutpostConfig) SetBindMode(v LDAPAPIAccessMode) {
	o.BindMode = &v
}

// GetMfaSupport returns the MfaSupport field value if set, zero value otherwise.
func (o *LDAPOutpostConfig) GetMfaSupport() bool {
	if o == nil || o.MfaSupport == nil {
		var ret bool
		return ret
	}
	return *o.MfaSupport
}

// GetMfaSupportOk returns a tuple with the MfaSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPOutpostConfig) GetMfaSupportOk() (*bool, bool) {
	if o == nil || o.MfaSupport == nil {
		return nil, false
	}
	return o.MfaSupport, true
}

// HasMfaSupport returns a boolean if a field has been set.
func (o *LDAPOutpostConfig) HasMfaSupport() bool {
	if o != nil && o.MfaSupport != nil {
		return true
	}

	return false
}

// SetMfaSupport gets a reference to the given bool and assigns it to the MfaSupport field.
func (o *LDAPOutpostConfig) SetMfaSupport(v bool) {
	o.MfaSupport = &v
}

func (o LDAPOutpostConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.BaseDn != nil {
		toSerialize["base_dn"] = o.BaseDn
	}
	if true {
		toSerialize["bind_flow_slug"] = o.BindFlowSlug
	}
	if true {
		toSerialize["application_slug"] = o.ApplicationSlug
	}
	if o.SearchGroup.IsSet() {
		toSerialize["search_group"] = o.SearchGroup.Get()
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.TlsServerName != nil {
		toSerialize["tls_server_name"] = o.TlsServerName
	}
	if o.UidStartNumber != nil {
		toSerialize["uid_start_number"] = o.UidStartNumber
	}
	if o.GidStartNumber != nil {
		toSerialize["gid_start_number"] = o.GidStartNumber
	}
	if o.SearchMode != nil {
		toSerialize["search_mode"] = o.SearchMode
	}
	if o.BindMode != nil {
		toSerialize["bind_mode"] = o.BindMode
	}
	if o.MfaSupport != nil {
		toSerialize["mfa_support"] = o.MfaSupport
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPOutpostConfig struct {
	value *LDAPOutpostConfig
	isSet bool
}

func (v NullableLDAPOutpostConfig) Get() *LDAPOutpostConfig {
	return v.value
}

func (v *NullableLDAPOutpostConfig) Set(val *LDAPOutpostConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPOutpostConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPOutpostConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPOutpostConfig(val *LDAPOutpostConfig) *NullableLDAPOutpostConfig {
	return &NullableLDAPOutpostConfig{value: val, isSet: true}
}

func (v NullableLDAPOutpostConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPOutpostConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
