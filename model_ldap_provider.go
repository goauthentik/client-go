/*
authentik

Making authentication simple.

API version: 2022.11.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LDAPProvider LDAPProvider Serializer
type LDAPProvider struct {
	Pk   int32  `json:"pk"`
	Name string `json:"name"`
	// Flow used when authorizing this provider.
	AuthorizationFlow string   `json:"authorization_flow"`
	PropertyMappings  []string `json:"property_mappings,omitempty"`
	Component         string   `json:"component"`
	// Internal application name, used in URLs.
	AssignedApplicationSlug string `json:"assigned_application_slug"`
	// Application's display Name.
	AssignedApplicationName string `json:"assigned_application_name"`
	VerboseName             string `json:"verbose_name"`
	VerboseNamePlural       string `json:"verbose_name_plural"`
	MetaModelName           string `json:"meta_model_name"`
	// DN under which objects are accessible.
	BaseDn *string `json:"base_dn,omitempty"`
	// Users in this group can do search queries. If not set, every user can execute search queries.
	SearchGroup   NullableString `json:"search_group,omitempty"`
	Certificate   NullableString `json:"certificate,omitempty"`
	TlsServerName *string        `json:"tls_server_name,omitempty"`
	// The start for uidNumbers, this number is added to the user.Pk to make sure that the numbers aren't too low for POSIX users. Default is 2000 to ensure that we don't collide with local users uidNumber
	UidStartNumber *int32 `json:"uid_start_number,omitempty"`
	// The start for gidNumbers, this number is added to a number generated from the group.Pk to make sure that the numbers aren't too low for POSIX groups. Default is 4000 to ensure that we don't collide with local groups or users primary groups gidNumber
	GidStartNumber *int32             `json:"gid_start_number,omitempty"`
	OutpostSet     []string           `json:"outpost_set"`
	SearchMode     *LDAPAPIAccessMode `json:"search_mode,omitempty"`
	BindMode       *LDAPAPIAccessMode `json:"bind_mode,omitempty"`
}

// NewLDAPProvider instantiates a new LDAPProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPProvider(pk int32, name string, authorizationFlow string, component string, assignedApplicationSlug string, assignedApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, outpostSet []string) *LDAPProvider {
	this := LDAPProvider{}
	this.Pk = pk
	this.Name = name
	this.AuthorizationFlow = authorizationFlow
	this.Component = component
	this.AssignedApplicationSlug = assignedApplicationSlug
	this.AssignedApplicationName = assignedApplicationName
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.OutpostSet = outpostSet
	return &this
}

// NewLDAPProviderWithDefaults instantiates a new LDAPProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPProviderWithDefaults() *LDAPProvider {
	this := LDAPProvider{}
	return &this
}

// GetPk returns the Pk field value
func (o *LDAPProvider) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *LDAPProvider) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *LDAPProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LDAPProvider) SetName(v string) {
	o.Name = v
}

// GetAuthorizationFlow returns the AuthorizationFlow field value
func (o *LDAPProvider) GetAuthorizationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationFlow, true
}

// SetAuthorizationFlow sets field value
func (o *LDAPProvider) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *LDAPProvider) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *LDAPProvider) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *LDAPProvider) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetComponent returns the Component field value
func (o *LDAPProvider) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *LDAPProvider) SetComponent(v string) {
	o.Component = v
}

// GetAssignedApplicationSlug returns the AssignedApplicationSlug field value
func (o *LDAPProvider) GetAssignedApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationSlug
}

// GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetAssignedApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationSlug, true
}

// SetAssignedApplicationSlug sets field value
func (o *LDAPProvider) SetAssignedApplicationSlug(v string) {
	o.AssignedApplicationSlug = v
}

// GetAssignedApplicationName returns the AssignedApplicationName field value
func (o *LDAPProvider) GetAssignedApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationName
}

// GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field value
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetAssignedApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationName, true
}

// SetAssignedApplicationName sets field value
func (o *LDAPProvider) SetAssignedApplicationName(v string) {
	o.AssignedApplicationName = v
}

// GetVerboseName returns the VerboseName field value
func (o *LDAPProvider) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *LDAPProvider) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *LDAPProvider) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *LDAPProvider) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *LDAPProvider) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *LDAPProvider) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetBaseDn returns the BaseDn field value if set, zero value otherwise.
func (o *LDAPProvider) GetBaseDn() string {
	if o == nil || o.BaseDn == nil {
		var ret string
		return ret
	}
	return *o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetBaseDnOk() (*string, bool) {
	if o == nil || o.BaseDn == nil {
		return nil, false
	}
	return o.BaseDn, true
}

// HasBaseDn returns a boolean if a field has been set.
func (o *LDAPProvider) HasBaseDn() bool {
	if o != nil && o.BaseDn != nil {
		return true
	}

	return false
}

// SetBaseDn gets a reference to the given string and assigns it to the BaseDn field.
func (o *LDAPProvider) SetBaseDn(v string) {
	o.BaseDn = &v
}

// GetSearchGroup returns the SearchGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPProvider) GetSearchGroup() string {
	if o == nil || o.SearchGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.SearchGroup.Get()
}

// GetSearchGroupOk returns a tuple with the SearchGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPProvider) GetSearchGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchGroup.Get(), o.SearchGroup.IsSet()
}

// HasSearchGroup returns a boolean if a field has been set.
func (o *LDAPProvider) HasSearchGroup() bool {
	if o != nil && o.SearchGroup.IsSet() {
		return true
	}

	return false
}

// SetSearchGroup gets a reference to the given NullableString and assigns it to the SearchGroup field.
func (o *LDAPProvider) SetSearchGroup(v string) {
	o.SearchGroup.Set(&v)
}

// SetSearchGroupNil sets the value for SearchGroup to be an explicit nil
func (o *LDAPProvider) SetSearchGroupNil() {
	o.SearchGroup.Set(nil)
}

// UnsetSearchGroup ensures that no value is present for SearchGroup, not even an explicit nil
func (o *LDAPProvider) UnsetSearchGroup() {
	o.SearchGroup.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPProvider) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPProvider) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *LDAPProvider) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *LDAPProvider) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *LDAPProvider) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *LDAPProvider) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetTlsServerName returns the TlsServerName field value if set, zero value otherwise.
func (o *LDAPProvider) GetTlsServerName() string {
	if o == nil || o.TlsServerName == nil {
		var ret string
		return ret
	}
	return *o.TlsServerName
}

// GetTlsServerNameOk returns a tuple with the TlsServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetTlsServerNameOk() (*string, bool) {
	if o == nil || o.TlsServerName == nil {
		return nil, false
	}
	return o.TlsServerName, true
}

// HasTlsServerName returns a boolean if a field has been set.
func (o *LDAPProvider) HasTlsServerName() bool {
	if o != nil && o.TlsServerName != nil {
		return true
	}

	return false
}

// SetTlsServerName gets a reference to the given string and assigns it to the TlsServerName field.
func (o *LDAPProvider) SetTlsServerName(v string) {
	o.TlsServerName = &v
}

// GetUidStartNumber returns the UidStartNumber field value if set, zero value otherwise.
func (o *LDAPProvider) GetUidStartNumber() int32 {
	if o == nil || o.UidStartNumber == nil {
		var ret int32
		return ret
	}
	return *o.UidStartNumber
}

// GetUidStartNumberOk returns a tuple with the UidStartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetUidStartNumberOk() (*int32, bool) {
	if o == nil || o.UidStartNumber == nil {
		return nil, false
	}
	return o.UidStartNumber, true
}

// HasUidStartNumber returns a boolean if a field has been set.
func (o *LDAPProvider) HasUidStartNumber() bool {
	if o != nil && o.UidStartNumber != nil {
		return true
	}

	return false
}

// SetUidStartNumber gets a reference to the given int32 and assigns it to the UidStartNumber field.
func (o *LDAPProvider) SetUidStartNumber(v int32) {
	o.UidStartNumber = &v
}

// GetGidStartNumber returns the GidStartNumber field value if set, zero value otherwise.
func (o *LDAPProvider) GetGidStartNumber() int32 {
	if o == nil || o.GidStartNumber == nil {
		var ret int32
		return ret
	}
	return *o.GidStartNumber
}

// GetGidStartNumberOk returns a tuple with the GidStartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetGidStartNumberOk() (*int32, bool) {
	if o == nil || o.GidStartNumber == nil {
		return nil, false
	}
	return o.GidStartNumber, true
}

// HasGidStartNumber returns a boolean if a field has been set.
func (o *LDAPProvider) HasGidStartNumber() bool {
	if o != nil && o.GidStartNumber != nil {
		return true
	}

	return false
}

// SetGidStartNumber gets a reference to the given int32 and assigns it to the GidStartNumber field.
func (o *LDAPProvider) SetGidStartNumber(v int32) {
	o.GidStartNumber = &v
}

// GetOutpostSet returns the OutpostSet field value
func (o *LDAPProvider) GetOutpostSet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OutpostSet
}

// GetOutpostSetOk returns a tuple with the OutpostSet field value
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetOutpostSetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutpostSet, true
}

// SetOutpostSet sets field value
func (o *LDAPProvider) SetOutpostSet(v []string) {
	o.OutpostSet = v
}

// GetSearchMode returns the SearchMode field value if set, zero value otherwise.
func (o *LDAPProvider) GetSearchMode() LDAPAPIAccessMode {
	if o == nil || o.SearchMode == nil {
		var ret LDAPAPIAccessMode
		return ret
	}
	return *o.SearchMode
}

// GetSearchModeOk returns a tuple with the SearchMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetSearchModeOk() (*LDAPAPIAccessMode, bool) {
	if o == nil || o.SearchMode == nil {
		return nil, false
	}
	return o.SearchMode, true
}

// HasSearchMode returns a boolean if a field has been set.
func (o *LDAPProvider) HasSearchMode() bool {
	if o != nil && o.SearchMode != nil {
		return true
	}

	return false
}

// SetSearchMode gets a reference to the given LDAPAPIAccessMode and assigns it to the SearchMode field.
func (o *LDAPProvider) SetSearchMode(v LDAPAPIAccessMode) {
	o.SearchMode = &v
}

// GetBindMode returns the BindMode field value if set, zero value otherwise.
func (o *LDAPProvider) GetBindMode() LDAPAPIAccessMode {
	if o == nil || o.BindMode == nil {
		var ret LDAPAPIAccessMode
		return ret
	}
	return *o.BindMode
}

// GetBindModeOk returns a tuple with the BindMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProvider) GetBindModeOk() (*LDAPAPIAccessMode, bool) {
	if o == nil || o.BindMode == nil {
		return nil, false
	}
	return o.BindMode, true
}

// HasBindMode returns a boolean if a field has been set.
func (o *LDAPProvider) HasBindMode() bool {
	if o != nil && o.BindMode != nil {
		return true
	}

	return false
}

// SetBindMode gets a reference to the given LDAPAPIAccessMode and assigns it to the BindMode field.
func (o *LDAPProvider) SetBindMode(v LDAPAPIAccessMode) {
	o.BindMode = &v
}

func (o LDAPProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["authorization_flow"] = o.AuthorizationFlow
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["assigned_application_slug"] = o.AssignedApplicationSlug
	}
	if true {
		toSerialize["assigned_application_name"] = o.AssignedApplicationName
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
	if o.BaseDn != nil {
		toSerialize["base_dn"] = o.BaseDn
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
	if true {
		toSerialize["outpost_set"] = o.OutpostSet
	}
	if o.SearchMode != nil {
		toSerialize["search_mode"] = o.SearchMode
	}
	if o.BindMode != nil {
		toSerialize["bind_mode"] = o.BindMode
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPProvider struct {
	value *LDAPProvider
	isSet bool
}

func (v NullableLDAPProvider) Get() *LDAPProvider {
	return v.value
}

func (v *NullableLDAPProvider) Set(val *LDAPProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPProvider(val *LDAPProvider) *NullableLDAPProvider {
	return &NullableLDAPProvider{value: val, isSet: true}
}

func (v NullableLDAPProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
