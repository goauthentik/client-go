/*
authentik

Making authentication simple.

API version: 2023.3.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMProvider SCIMProvider Serializer
type SCIMProvider struct {
	Pk               int32    `json:"pk"`
	Name             string   `json:"name"`
	PropertyMappings []string `json:"property_mappings,omitempty"`
	// Property mappings used for group creation/updating.
	PropertyMappingsGroup []string `json:"property_mappings_group,omitempty"`
	// Get object component so that we know how to edit the object
	Component string `json:"component"`
	// Internal application name, used in URLs.
	AssignedApplicationSlug string `json:"assigned_application_slug"`
	// Application's display Name.
	AssignedApplicationName string `json:"assigned_application_name"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string `json:"meta_model_name"`
	// Base URL to SCIM requests, usually ends in /v2
	Url string `json:"url"`
	// Authentication token
	Token                      string         `json:"token"`
	ExcludeUsersServiceAccount *bool          `json:"exclude_users_service_account,omitempty"`
	FilterGroup                NullableString `json:"filter_group,omitempty"`
}

// NewSCIMProvider instantiates a new SCIMProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMProvider(pk int32, name string, component string, assignedApplicationSlug string, assignedApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, url string, token string) *SCIMProvider {
	this := SCIMProvider{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.AssignedApplicationSlug = assignedApplicationSlug
	this.AssignedApplicationName = assignedApplicationName
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Url = url
	this.Token = token
	return &this
}

// NewSCIMProviderWithDefaults instantiates a new SCIMProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMProviderWithDefaults() *SCIMProvider {
	this := SCIMProvider{}
	return &this
}

// GetPk returns the Pk field value
func (o *SCIMProvider) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *SCIMProvider) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *SCIMProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SCIMProvider) SetName(v string) {
	o.Name = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *SCIMProvider) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *SCIMProvider) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *SCIMProvider) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetPropertyMappingsGroup returns the PropertyMappingsGroup field value if set, zero value otherwise.
func (o *SCIMProvider) GetPropertyMappingsGroup() []string {
	if o == nil || o.PropertyMappingsGroup == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappingsGroup
}

// GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetPropertyMappingsGroupOk() ([]string, bool) {
	if o == nil || o.PropertyMappingsGroup == nil {
		return nil, false
	}
	return o.PropertyMappingsGroup, true
}

// HasPropertyMappingsGroup returns a boolean if a field has been set.
func (o *SCIMProvider) HasPropertyMappingsGroup() bool {
	if o != nil && o.PropertyMappingsGroup != nil {
		return true
	}

	return false
}

// SetPropertyMappingsGroup gets a reference to the given []string and assigns it to the PropertyMappingsGroup field.
func (o *SCIMProvider) SetPropertyMappingsGroup(v []string) {
	o.PropertyMappingsGroup = v
}

// GetComponent returns the Component field value
func (o *SCIMProvider) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *SCIMProvider) SetComponent(v string) {
	o.Component = v
}

// GetAssignedApplicationSlug returns the AssignedApplicationSlug field value
func (o *SCIMProvider) GetAssignedApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationSlug
}

// GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetAssignedApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationSlug, true
}

// SetAssignedApplicationSlug sets field value
func (o *SCIMProvider) SetAssignedApplicationSlug(v string) {
	o.AssignedApplicationSlug = v
}

// GetAssignedApplicationName returns the AssignedApplicationName field value
func (o *SCIMProvider) GetAssignedApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationName
}

// GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field value
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetAssignedApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationName, true
}

// SetAssignedApplicationName sets field value
func (o *SCIMProvider) SetAssignedApplicationName(v string) {
	o.AssignedApplicationName = v
}

// GetVerboseName returns the VerboseName field value
func (o *SCIMProvider) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *SCIMProvider) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *SCIMProvider) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *SCIMProvider) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *SCIMProvider) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *SCIMProvider) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetUrl returns the Url field value
func (o *SCIMProvider) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SCIMProvider) SetUrl(v string) {
	o.Url = v
}

// GetToken returns the Token field value
func (o *SCIMProvider) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *SCIMProvider) SetToken(v string) {
	o.Token = v
}

// GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field value if set, zero value otherwise.
func (o *SCIMProvider) GetExcludeUsersServiceAccount() bool {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeUsersServiceAccount
}

// GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCIMProvider) GetExcludeUsersServiceAccountOk() (*bool, bool) {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		return nil, false
	}
	return o.ExcludeUsersServiceAccount, true
}

// HasExcludeUsersServiceAccount returns a boolean if a field has been set.
func (o *SCIMProvider) HasExcludeUsersServiceAccount() bool {
	if o != nil && o.ExcludeUsersServiceAccount != nil {
		return true
	}

	return false
}

// SetExcludeUsersServiceAccount gets a reference to the given bool and assigns it to the ExcludeUsersServiceAccount field.
func (o *SCIMProvider) SetExcludeUsersServiceAccount(v bool) {
	o.ExcludeUsersServiceAccount = &v
}

// GetFilterGroup returns the FilterGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SCIMProvider) GetFilterGroup() string {
	if o == nil || o.FilterGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.FilterGroup.Get()
}

// GetFilterGroupOk returns a tuple with the FilterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SCIMProvider) GetFilterGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterGroup.Get(), o.FilterGroup.IsSet()
}

// HasFilterGroup returns a boolean if a field has been set.
func (o *SCIMProvider) HasFilterGroup() bool {
	if o != nil && o.FilterGroup.IsSet() {
		return true
	}

	return false
}

// SetFilterGroup gets a reference to the given NullableString and assigns it to the FilterGroup field.
func (o *SCIMProvider) SetFilterGroup(v string) {
	o.FilterGroup.Set(&v)
}

// SetFilterGroupNil sets the value for FilterGroup to be an explicit nil
func (o *SCIMProvider) SetFilterGroupNil() {
	o.FilterGroup.Set(nil)
}

// UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
func (o *SCIMProvider) UnsetFilterGroup() {
	o.FilterGroup.Unset()
}

func (o SCIMProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
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
	if true {
		toSerialize["url"] = o.Url
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

type NullableSCIMProvider struct {
	value *SCIMProvider
	isSet bool
}

func (v NullableSCIMProvider) Get() *SCIMProvider {
	return v.value
}

func (v *NullableSCIMProvider) Set(val *SCIMProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMProvider(val *SCIMProvider) *NullableSCIMProvider {
	return &NullableSCIMProvider{value: val, isSet: true}
}

func (v NullableSCIMProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}