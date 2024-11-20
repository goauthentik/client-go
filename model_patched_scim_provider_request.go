/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PatchedSCIMProviderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedSCIMProviderRequest{}

// PatchedSCIMProviderRequest SCIMProvider Serializer
type PatchedSCIMProviderRequest struct {
	Name             *string  `json:"name,omitempty"`
	PropertyMappings []string `json:"property_mappings,omitempty"`
	// Property mappings used for group creation/updating.
	PropertyMappingsGroup []string `json:"property_mappings_group,omitempty"`
	// Base URL to SCIM requests, usually ends in /v2
	Url                *string `json:"url,omitempty"`
	VerifyCertificates *bool   `json:"verify_certificates,omitempty"`
	// Authentication token
	Token                      *string        `json:"token,omitempty"`
	ExcludeUsersServiceAccount *bool          `json:"exclude_users_service_account,omitempty"`
	FilterGroup                NullableString `json:"filter_group,omitempty"`
}

// NewPatchedSCIMProviderRequest instantiates a new PatchedSCIMProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSCIMProviderRequest() *PatchedSCIMProviderRequest {
	this := PatchedSCIMProviderRequest{}
	return &this
}

// NewPatchedSCIMProviderRequestWithDefaults instantiates a new PatchedSCIMProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSCIMProviderRequestWithDefaults() *PatchedSCIMProviderRequest {
	this := PatchedSCIMProviderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedSCIMProviderRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedSCIMProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *PatchedSCIMProviderRequest) GetPropertyMappings() []string {
	if o == nil || IsNil(o.PropertyMappings) {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || IsNil(o.PropertyMappings) {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasPropertyMappings() bool {
	if o != nil && !IsNil(o.PropertyMappings) {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *PatchedSCIMProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetPropertyMappingsGroup returns the PropertyMappingsGroup field value if set, zero value otherwise.
func (o *PatchedSCIMProviderRequest) GetPropertyMappingsGroup() []string {
	if o == nil || IsNil(o.PropertyMappingsGroup) {
		var ret []string
		return ret
	}
	return o.PropertyMappingsGroup
}

// GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetPropertyMappingsGroupOk() ([]string, bool) {
	if o == nil || IsNil(o.PropertyMappingsGroup) {
		return nil, false
	}
	return o.PropertyMappingsGroup, true
}

// HasPropertyMappingsGroup returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasPropertyMappingsGroup() bool {
	if o != nil && !IsNil(o.PropertyMappingsGroup) {
		return true
	}

	return false
}

// SetPropertyMappingsGroup gets a reference to the given []string and assigns it to the PropertyMappingsGroup field.
func (o *PatchedSCIMProviderRequest) SetPropertyMappingsGroup(v []string) {
	o.PropertyMappingsGroup = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedSCIMProviderRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedSCIMProviderRequest) SetUrl(v string) {
	o.Url = &v
}

// GetVerifyCertificates returns the VerifyCertificates field value if set, zero value otherwise.
func (o *PatchedSCIMProviderRequest) GetVerifyCertificates() bool {
	if o == nil || IsNil(o.VerifyCertificates) {
		var ret bool
		return ret
	}
	return *o.VerifyCertificates
}

// GetVerifyCertificatesOk returns a tuple with the VerifyCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetVerifyCertificatesOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyCertificates) {
		return nil, false
	}
	return o.VerifyCertificates, true
}

// HasVerifyCertificates returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasVerifyCertificates() bool {
	if o != nil && !IsNil(o.VerifyCertificates) {
		return true
	}

	return false
}

// SetVerifyCertificates gets a reference to the given bool and assigns it to the VerifyCertificates field.
func (o *PatchedSCIMProviderRequest) SetVerifyCertificates(v bool) {
	o.VerifyCertificates = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PatchedSCIMProviderRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PatchedSCIMProviderRequest) SetToken(v string) {
	o.Token = &v
}

// GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field value if set, zero value otherwise.
func (o *PatchedSCIMProviderRequest) GetExcludeUsersServiceAccount() bool {
	if o == nil || IsNil(o.ExcludeUsersServiceAccount) {
		var ret bool
		return ret
	}
	return *o.ExcludeUsersServiceAccount
}

// GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeUsersServiceAccount) {
		return nil, false
	}
	return o.ExcludeUsersServiceAccount, true
}

// HasExcludeUsersServiceAccount returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasExcludeUsersServiceAccount() bool {
	if o != nil && !IsNil(o.ExcludeUsersServiceAccount) {
		return true
	}

	return false
}

// SetExcludeUsersServiceAccount gets a reference to the given bool and assigns it to the ExcludeUsersServiceAccount field.
func (o *PatchedSCIMProviderRequest) SetExcludeUsersServiceAccount(v bool) {
	o.ExcludeUsersServiceAccount = &v
}

// GetFilterGroup returns the FilterGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedSCIMProviderRequest) GetFilterGroup() string {
	if o == nil || IsNil(o.FilterGroup.Get()) {
		var ret string
		return ret
	}
	return *o.FilterGroup.Get()
}

// GetFilterGroupOk returns a tuple with the FilterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedSCIMProviderRequest) GetFilterGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterGroup.Get(), o.FilterGroup.IsSet()
}

// HasFilterGroup returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasFilterGroup() bool {
	if o != nil && o.FilterGroup.IsSet() {
		return true
	}

	return false
}

// SetFilterGroup gets a reference to the given NullableString and assigns it to the FilterGroup field.
func (o *PatchedSCIMProviderRequest) SetFilterGroup(v string) {
	o.FilterGroup.Set(&v)
}

// SetFilterGroupNil sets the value for FilterGroup to be an explicit nil
func (o *PatchedSCIMProviderRequest) SetFilterGroupNil() {
	o.FilterGroup.Set(nil)
}

// UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
func (o *PatchedSCIMProviderRequest) UnsetFilterGroup() {
	o.FilterGroup.Unset()
}

func (o PatchedSCIMProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedSCIMProviderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PropertyMappings) {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if !IsNil(o.PropertyMappingsGroup) {
		toSerialize["property_mappings_group"] = o.PropertyMappingsGroup
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.VerifyCertificates) {
		toSerialize["verify_certificates"] = o.VerifyCertificates
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.ExcludeUsersServiceAccount) {
		toSerialize["exclude_users_service_account"] = o.ExcludeUsersServiceAccount
	}
	if o.FilterGroup.IsSet() {
		toSerialize["filter_group"] = o.FilterGroup.Get()
	}
	return toSerialize, nil
}

type NullablePatchedSCIMProviderRequest struct {
	value *PatchedSCIMProviderRequest
	isSet bool
}

func (v NullablePatchedSCIMProviderRequest) Get() *PatchedSCIMProviderRequest {
	return v.value
}

func (v *NullablePatchedSCIMProviderRequest) Set(val *PatchedSCIMProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSCIMProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSCIMProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSCIMProviderRequest(val *PatchedSCIMProviderRequest) *NullablePatchedSCIMProviderRequest {
	return &NullablePatchedSCIMProviderRequest{value: val, isSet: true}
}

func (v NullablePatchedSCIMProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSCIMProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
