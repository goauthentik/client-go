/*
authentik

Making authentication simple.

API version: 2024.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMProviderRequest SCIMProvider Serializer
type SCIMProviderRequest struct {
	Name             string   `json:"name"`
	PropertyMappings []string `json:"property_mappings,omitempty"`
	// Property mappings used for group creation/updating.
	PropertyMappingsGroup []string `json:"property_mappings_group,omitempty"`
	// Base URL to SCIM requests, usually ends in /v2
	Url                string `json:"url"`
	VerifyCertificates *bool  `json:"verify_certificates,omitempty"`
	// Authentication token
	Token                      string         `json:"token"`
	ExcludeUsersServiceAccount *bool          `json:"exclude_users_service_account,omitempty"`
	FilterGroup                NullableString `json:"filter_group,omitempty"`
}

// NewSCIMProviderRequest instantiates a new SCIMProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMProviderRequest(name string, url string, token string) *SCIMProviderRequest {
	this := SCIMProviderRequest{}
	this.Name = name
	this.Url = url
	this.Token = token
	return &this
}

// NewSCIMProviderRequestWithDefaults instantiates a new SCIMProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMProviderRequestWithDefaults() *SCIMProviderRequest {
	this := SCIMProviderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SCIMProviderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SCIMProviderRequest) SetName(v string) {
	o.Name = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *SCIMProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *SCIMProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *SCIMProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetPropertyMappingsGroup returns the PropertyMappingsGroup field value if set, zero value otherwise.
func (o *SCIMProviderRequest) GetPropertyMappingsGroup() []string {
	if o == nil || o.PropertyMappingsGroup == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappingsGroup
}

// GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMProviderRequest) GetPropertyMappingsGroupOk() ([]string, bool) {
	if o == nil || o.PropertyMappingsGroup == nil {
		return nil, false
	}
	return o.PropertyMappingsGroup, true
}

// HasPropertyMappingsGroup returns a boolean if a field has been set.
func (o *SCIMProviderRequest) HasPropertyMappingsGroup() bool {
	if o != nil && o.PropertyMappingsGroup != nil {
		return true
	}

	return false
}

// SetPropertyMappingsGroup gets a reference to the given []string and assigns it to the PropertyMappingsGroup field.
func (o *SCIMProviderRequest) SetPropertyMappingsGroup(v []string) {
	o.PropertyMappingsGroup = v
}

// GetUrl returns the Url field value
func (o *SCIMProviderRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SCIMProviderRequest) SetUrl(v string) {
	o.Url = v
}

// GetVerifyCertificates returns the VerifyCertificates field value if set, zero value otherwise.
func (o *SCIMProviderRequest) GetVerifyCertificates() bool {
	if o == nil || o.VerifyCertificates == nil {
		var ret bool
		return ret
	}
	return *o.VerifyCertificates
}

// GetVerifyCertificatesOk returns a tuple with the VerifyCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMProviderRequest) GetVerifyCertificatesOk() (*bool, bool) {
	if o == nil || o.VerifyCertificates == nil {
		return nil, false
	}
	return o.VerifyCertificates, true
}

// HasVerifyCertificates returns a boolean if a field has been set.
func (o *SCIMProviderRequest) HasVerifyCertificates() bool {
	if o != nil && o.VerifyCertificates != nil {
		return true
	}

	return false
}

// SetVerifyCertificates gets a reference to the given bool and assigns it to the VerifyCertificates field.
func (o *SCIMProviderRequest) SetVerifyCertificates(v bool) {
	o.VerifyCertificates = &v
}

// GetToken returns the Token field value
func (o *SCIMProviderRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *SCIMProviderRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *SCIMProviderRequest) SetToken(v string) {
	o.Token = v
}

// GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field value if set, zero value otherwise.
func (o *SCIMProviderRequest) GetExcludeUsersServiceAccount() bool {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeUsersServiceAccount
}

// GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool) {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		return nil, false
	}
	return o.ExcludeUsersServiceAccount, true
}

// HasExcludeUsersServiceAccount returns a boolean if a field has been set.
func (o *SCIMProviderRequest) HasExcludeUsersServiceAccount() bool {
	if o != nil && o.ExcludeUsersServiceAccount != nil {
		return true
	}

	return false
}

// SetExcludeUsersServiceAccount gets a reference to the given bool and assigns it to the ExcludeUsersServiceAccount field.
func (o *SCIMProviderRequest) SetExcludeUsersServiceAccount(v bool) {
	o.ExcludeUsersServiceAccount = &v
}

// GetFilterGroup returns the FilterGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SCIMProviderRequest) GetFilterGroup() string {
	if o == nil || o.FilterGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.FilterGroup.Get()
}

// GetFilterGroupOk returns a tuple with the FilterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SCIMProviderRequest) GetFilterGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterGroup.Get(), o.FilterGroup.IsSet()
}

// HasFilterGroup returns a boolean if a field has been set.
func (o *SCIMProviderRequest) HasFilterGroup() bool {
	if o != nil && o.FilterGroup.IsSet() {
		return true
	}

	return false
}

// SetFilterGroup gets a reference to the given NullableString and assigns it to the FilterGroup field.
func (o *SCIMProviderRequest) SetFilterGroup(v string) {
	o.FilterGroup.Set(&v)
}

// SetFilterGroupNil sets the value for FilterGroup to be an explicit nil
func (o *SCIMProviderRequest) SetFilterGroupNil() {
	o.FilterGroup.Set(nil)
}

// UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
func (o *SCIMProviderRequest) UnsetFilterGroup() {
	o.FilterGroup.Unset()
}

func (o SCIMProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if o.PropertyMappingsGroup != nil {
		toSerialize["property_mappings_group"] = o.PropertyMappingsGroup
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.VerifyCertificates != nil {
		toSerialize["verify_certificates"] = o.VerifyCertificates
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if o.ExcludeUsersServiceAccount != nil {
		toSerialize["exclude_users_service_account"] = o.ExcludeUsersServiceAccount
	}
	if o.FilterGroup.IsSet() {
		toSerialize["filter_group"] = o.FilterGroup.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMProviderRequest struct {
	value *SCIMProviderRequest
	isSet bool
}

func (v NullableSCIMProviderRequest) Get() *SCIMProviderRequest {
	return v.value
}

func (v *NullableSCIMProviderRequest) Set(val *SCIMProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMProviderRequest(val *SCIMProviderRequest) *NullableSCIMProviderRequest {
	return &NullableSCIMProviderRequest{value: val, isSet: true}
}

func (v NullableSCIMProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
