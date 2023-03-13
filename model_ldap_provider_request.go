/*
authentik

Making authentication simple.

API version: 2023.3.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LDAPProviderRequest LDAPProvider Serializer
type LDAPProviderRequest struct {
	Name string `json:"name"`
	// Flow used when authorizing this provider.
	AuthorizationFlow NullableString `json:"authorization_flow,omitempty"`
	PropertyMappings  []string       `json:"property_mappings,omitempty"`
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
	SearchMode     *LDAPAPIAccessMode `json:"search_mode,omitempty"`
	BindMode       *LDAPAPIAccessMode `json:"bind_mode,omitempty"`
}

// NewLDAPProviderRequest instantiates a new LDAPProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPProviderRequest(name string) *LDAPProviderRequest {
	this := LDAPProviderRequest{}
	this.Name = name
	return &this
}

// NewLDAPProviderRequestWithDefaults instantiates a new LDAPProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPProviderRequestWithDefaults() *LDAPProviderRequest {
	this := LDAPProviderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *LDAPProviderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LDAPProviderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LDAPProviderRequest) SetName(v string) {
	o.Name = v
}

// GetAuthorizationFlow returns the AuthorizationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPProviderRequest) GetAuthorizationFlow() string {
	if o == nil || o.AuthorizationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationFlow.Get()
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationFlow.Get(), o.AuthorizationFlow.IsSet()
}

// HasAuthorizationFlow returns a boolean if a field has been set.
func (o *LDAPProviderRequest) HasAuthorizationFlow() bool {
	if o != nil && o.AuthorizationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationFlow gets a reference to the given NullableString and assigns it to the AuthorizationFlow field.
func (o *LDAPProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow.Set(&v)
}

// SetAuthorizationFlowNil sets the value for AuthorizationFlow to be an explicit nil
func (o *LDAPProviderRequest) SetAuthorizationFlowNil() {
	o.AuthorizationFlow.Set(nil)
}

// UnsetAuthorizationFlow ensures that no value is present for AuthorizationFlow, not even an explicit nil
func (o *LDAPProviderRequest) UnsetAuthorizationFlow() {
	o.AuthorizationFlow.Unset()
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *LDAPProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *LDAPProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *LDAPProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetBaseDn returns the BaseDn field value if set, zero value otherwise.
func (o *LDAPProviderRequest) GetBaseDn() string {
	if o == nil || o.BaseDn == nil {
		var ret string
		return ret
	}
	return *o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProviderRequest) GetBaseDnOk() (*string, bool) {
	if o == nil || o.BaseDn == nil {
		return nil, false
	}
	return o.BaseDn, true
}

// HasBaseDn returns a boolean if a field has been set.
func (o *LDAPProviderRequest) HasBaseDn() bool {
	if o != nil && o.BaseDn != nil {
		return true
	}

	return false
}

// SetBaseDn gets a reference to the given string and assigns it to the BaseDn field.
func (o *LDAPProviderRequest) SetBaseDn(v string) {
	o.BaseDn = &v
}

// GetSearchGroup returns the SearchGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPProviderRequest) GetSearchGroup() string {
	if o == nil || o.SearchGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.SearchGroup.Get()
}

// GetSearchGroupOk returns a tuple with the SearchGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPProviderRequest) GetSearchGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchGroup.Get(), o.SearchGroup.IsSet()
}

// HasSearchGroup returns a boolean if a field has been set.
func (o *LDAPProviderRequest) HasSearchGroup() bool {
	if o != nil && o.SearchGroup.IsSet() {
		return true
	}

	return false
}

// SetSearchGroup gets a reference to the given NullableString and assigns it to the SearchGroup field.
func (o *LDAPProviderRequest) SetSearchGroup(v string) {
	o.SearchGroup.Set(&v)
}

// SetSearchGroupNil sets the value for SearchGroup to be an explicit nil
func (o *LDAPProviderRequest) SetSearchGroupNil() {
	o.SearchGroup.Set(nil)
}

// UnsetSearchGroup ensures that no value is present for SearchGroup, not even an explicit nil
func (o *LDAPProviderRequest) UnsetSearchGroup() {
	o.SearchGroup.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPProviderRequest) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPProviderRequest) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *LDAPProviderRequest) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *LDAPProviderRequest) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *LDAPProviderRequest) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *LDAPProviderRequest) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetTlsServerName returns the TlsServerName field value if set, zero value otherwise.
func (o *LDAPProviderRequest) GetTlsServerName() string {
	if o == nil || o.TlsServerName == nil {
		var ret string
		return ret
	}
	return *o.TlsServerName
}

// GetTlsServerNameOk returns a tuple with the TlsServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProviderRequest) GetTlsServerNameOk() (*string, bool) {
	if o == nil || o.TlsServerName == nil {
		return nil, false
	}
	return o.TlsServerName, true
}

// HasTlsServerName returns a boolean if a field has been set.
func (o *LDAPProviderRequest) HasTlsServerName() bool {
	if o != nil && o.TlsServerName != nil {
		return true
	}

	return false
}

// SetTlsServerName gets a reference to the given string and assigns it to the TlsServerName field.
func (o *LDAPProviderRequest) SetTlsServerName(v string) {
	o.TlsServerName = &v
}

// GetUidStartNumber returns the UidStartNumber field value if set, zero value otherwise.
func (o *LDAPProviderRequest) GetUidStartNumber() int32 {
	if o == nil || o.UidStartNumber == nil {
		var ret int32
		return ret
	}
	return *o.UidStartNumber
}

// GetUidStartNumberOk returns a tuple with the UidStartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProviderRequest) GetUidStartNumberOk() (*int32, bool) {
	if o == nil || o.UidStartNumber == nil {
		return nil, false
	}
	return o.UidStartNumber, true
}

// HasUidStartNumber returns a boolean if a field has been set.
func (o *LDAPProviderRequest) HasUidStartNumber() bool {
	if o != nil && o.UidStartNumber != nil {
		return true
	}

	return false
}

// SetUidStartNumber gets a reference to the given int32 and assigns it to the UidStartNumber field.
func (o *LDAPProviderRequest) SetUidStartNumber(v int32) {
	o.UidStartNumber = &v
}

// GetGidStartNumber returns the GidStartNumber field value if set, zero value otherwise.
func (o *LDAPProviderRequest) GetGidStartNumber() int32 {
	if o == nil || o.GidStartNumber == nil {
		var ret int32
		return ret
	}
	return *o.GidStartNumber
}

// GetGidStartNumberOk returns a tuple with the GidStartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProviderRequest) GetGidStartNumberOk() (*int32, bool) {
	if o == nil || o.GidStartNumber == nil {
		return nil, false
	}
	return o.GidStartNumber, true
}

// HasGidStartNumber returns a boolean if a field has been set.
func (o *LDAPProviderRequest) HasGidStartNumber() bool {
	if o != nil && o.GidStartNumber != nil {
		return true
	}

	return false
}

// SetGidStartNumber gets a reference to the given int32 and assigns it to the GidStartNumber field.
func (o *LDAPProviderRequest) SetGidStartNumber(v int32) {
	o.GidStartNumber = &v
}

// GetSearchMode returns the SearchMode field value if set, zero value otherwise.
func (o *LDAPProviderRequest) GetSearchMode() LDAPAPIAccessMode {
	if o == nil || o.SearchMode == nil {
		var ret LDAPAPIAccessMode
		return ret
	}
	return *o.SearchMode
}

// GetSearchModeOk returns a tuple with the SearchMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProviderRequest) GetSearchModeOk() (*LDAPAPIAccessMode, bool) {
	if o == nil || o.SearchMode == nil {
		return nil, false
	}
	return o.SearchMode, true
}

// HasSearchMode returns a boolean if a field has been set.
func (o *LDAPProviderRequest) HasSearchMode() bool {
	if o != nil && o.SearchMode != nil {
		return true
	}

	return false
}

// SetSearchMode gets a reference to the given LDAPAPIAccessMode and assigns it to the SearchMode field.
func (o *LDAPProviderRequest) SetSearchMode(v LDAPAPIAccessMode) {
	o.SearchMode = &v
}

// GetBindMode returns the BindMode field value if set, zero value otherwise.
func (o *LDAPProviderRequest) GetBindMode() LDAPAPIAccessMode {
	if o == nil || o.BindMode == nil {
		var ret LDAPAPIAccessMode
		return ret
	}
	return *o.BindMode
}

// GetBindModeOk returns a tuple with the BindMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPProviderRequest) GetBindModeOk() (*LDAPAPIAccessMode, bool) {
	if o == nil || o.BindMode == nil {
		return nil, false
	}
	return o.BindMode, true
}

// HasBindMode returns a boolean if a field has been set.
func (o *LDAPProviderRequest) HasBindMode() bool {
	if o != nil && o.BindMode != nil {
		return true
	}

	return false
}

// SetBindMode gets a reference to the given LDAPAPIAccessMode and assigns it to the BindMode field.
func (o *LDAPProviderRequest) SetBindMode(v LDAPAPIAccessMode) {
	o.BindMode = &v
}

func (o LDAPProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.AuthorizationFlow.IsSet() {
		toSerialize["authorization_flow"] = o.AuthorizationFlow.Get()
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
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
	if o.SearchMode != nil {
		toSerialize["search_mode"] = o.SearchMode
	}
	if o.BindMode != nil {
		toSerialize["bind_mode"] = o.BindMode
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPProviderRequest struct {
	value *LDAPProviderRequest
	isSet bool
}

func (v NullableLDAPProviderRequest) Get() *LDAPProviderRequest {
	return v.value
}

func (v *NullableLDAPProviderRequest) Set(val *LDAPProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPProviderRequest(val *LDAPProviderRequest) *NullableLDAPProviderRequest {
	return &NullableLDAPProviderRequest{value: val, isSet: true}
}

func (v NullableLDAPProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
