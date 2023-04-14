/*
authentik

Making authentication simple.

API version: 2023.4.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedSCIMProviderRequest SCIMProvider Serializer
type PatchedSCIMProviderRequest struct {
	Name             *string  `json:"name,omitempty"`
	PropertyMappings []string `json:"property_mappings,omitempty"`
	// Property mappings used for group creation/updating.
	PropertyMappingsGroup []string `json:"property_mappings_group,omitempty"`
	// Base URL to SCIM requests, usually ends in /v2
	Url *string `json:"url,omitempty"`
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
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasName() bool {
	if o != nil && o.Name != nil {
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
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
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
	if o == nil || o.PropertyMappingsGroup == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappingsGroup
}

// GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetPropertyMappingsGroupOk() ([]string, bool) {
	if o == nil || o.PropertyMappingsGroup == nil {
		return nil, false
	}
	return o.PropertyMappingsGroup, true
}

// HasPropertyMappingsGroup returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasPropertyMappingsGroup() bool {
	if o != nil && o.PropertyMappingsGroup != nil {
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
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedSCIMProviderRequest) SetUrl(v string) {
	o.Url = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PatchedSCIMProviderRequest) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasToken() bool {
	if o != nil && o.Token != nil {
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
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeUsersServiceAccount
}

// GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSCIMProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool) {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		return nil, false
	}
	return o.ExcludeUsersServiceAccount, true
}

// HasExcludeUsersServiceAccount returns a boolean if a field has been set.
func (o *PatchedSCIMProviderRequest) HasExcludeUsersServiceAccount() bool {
	if o != nil && o.ExcludeUsersServiceAccount != nil {
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
	if o == nil || o.FilterGroup.Get() == nil {
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
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if o.PropertyMappingsGroup != nil {
		toSerialize["property_mappings_group"] = o.PropertyMappingsGroup
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Token != nil {
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
