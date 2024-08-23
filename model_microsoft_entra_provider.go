/*
authentik

Making authentication simple.

API version: 2024.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MicrosoftEntraProvider MicrosoftEntraProvider Serializer
type MicrosoftEntraProvider struct {
	Pk               int32    `json:"pk"`
	Name             string   `json:"name"`
	PropertyMappings []string `json:"property_mappings,omitempty"`
	// Property mappings used for group creation/updating.
	PropertyMappingsGroup []string `json:"property_mappings_group,omitempty"`
	// Get object component so that we know how to edit the object
	Component string `json:"component"`
	// Internal application name, used in URLs.
	AssignedBackchannelApplicationSlug string `json:"assigned_backchannel_application_slug"`
	// Application's display Name.
	AssignedBackchannelApplicationName string `json:"assigned_backchannel_application_name"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName              string                    `json:"meta_model_name"`
	ClientId                   string                    `json:"client_id"`
	ClientSecret               string                    `json:"client_secret"`
	TenantId                   string                    `json:"tenant_id"`
	ExcludeUsersServiceAccount *bool                     `json:"exclude_users_service_account,omitempty"`
	FilterGroup                NullableString            `json:"filter_group,omitempty"`
	UserDeleteAction           *OutgoingSyncDeleteAction `json:"user_delete_action,omitempty"`
	GroupDeleteAction          *OutgoingSyncDeleteAction `json:"group_delete_action,omitempty"`
}

// NewMicrosoftEntraProvider instantiates a new MicrosoftEntraProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftEntraProvider(pk int32, name string, component string, assignedBackchannelApplicationSlug string, assignedBackchannelApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, clientId string, clientSecret string, tenantId string) *MicrosoftEntraProvider {
	this := MicrosoftEntraProvider{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.AssignedBackchannelApplicationSlug = assignedBackchannelApplicationSlug
	this.AssignedBackchannelApplicationName = assignedBackchannelApplicationName
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.TenantId = tenantId
	return &this
}

// NewMicrosoftEntraProviderWithDefaults instantiates a new MicrosoftEntraProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftEntraProviderWithDefaults() *MicrosoftEntraProvider {
	this := MicrosoftEntraProvider{}
	return &this
}

// GetPk returns the Pk field value
func (o *MicrosoftEntraProvider) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *MicrosoftEntraProvider) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *MicrosoftEntraProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MicrosoftEntraProvider) SetName(v string) {
	o.Name = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *MicrosoftEntraProvider) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *MicrosoftEntraProvider) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *MicrosoftEntraProvider) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetPropertyMappingsGroup returns the PropertyMappingsGroup field value if set, zero value otherwise.
func (o *MicrosoftEntraProvider) GetPropertyMappingsGroup() []string {
	if o == nil || o.PropertyMappingsGroup == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappingsGroup
}

// GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetPropertyMappingsGroupOk() ([]string, bool) {
	if o == nil || o.PropertyMappingsGroup == nil {
		return nil, false
	}
	return o.PropertyMappingsGroup, true
}

// HasPropertyMappingsGroup returns a boolean if a field has been set.
func (o *MicrosoftEntraProvider) HasPropertyMappingsGroup() bool {
	if o != nil && o.PropertyMappingsGroup != nil {
		return true
	}

	return false
}

// SetPropertyMappingsGroup gets a reference to the given []string and assigns it to the PropertyMappingsGroup field.
func (o *MicrosoftEntraProvider) SetPropertyMappingsGroup(v []string) {
	o.PropertyMappingsGroup = v
}

// GetComponent returns the Component field value
func (o *MicrosoftEntraProvider) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *MicrosoftEntraProvider) SetComponent(v string) {
	o.Component = v
}

// GetAssignedBackchannelApplicationSlug returns the AssignedBackchannelApplicationSlug field value
func (o *MicrosoftEntraProvider) GetAssignedBackchannelApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedBackchannelApplicationSlug
}

// GetAssignedBackchannelApplicationSlugOk returns a tuple with the AssignedBackchannelApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetAssignedBackchannelApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedBackchannelApplicationSlug, true
}

// SetAssignedBackchannelApplicationSlug sets field value
func (o *MicrosoftEntraProvider) SetAssignedBackchannelApplicationSlug(v string) {
	o.AssignedBackchannelApplicationSlug = v
}

// GetAssignedBackchannelApplicationName returns the AssignedBackchannelApplicationName field value
func (o *MicrosoftEntraProvider) GetAssignedBackchannelApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedBackchannelApplicationName
}

// GetAssignedBackchannelApplicationNameOk returns a tuple with the AssignedBackchannelApplicationName field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetAssignedBackchannelApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedBackchannelApplicationName, true
}

// SetAssignedBackchannelApplicationName sets field value
func (o *MicrosoftEntraProvider) SetAssignedBackchannelApplicationName(v string) {
	o.AssignedBackchannelApplicationName = v
}

// GetVerboseName returns the VerboseName field value
func (o *MicrosoftEntraProvider) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *MicrosoftEntraProvider) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *MicrosoftEntraProvider) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *MicrosoftEntraProvider) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *MicrosoftEntraProvider) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *MicrosoftEntraProvider) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetClientId returns the ClientId field value
func (o *MicrosoftEntraProvider) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *MicrosoftEntraProvider) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *MicrosoftEntraProvider) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *MicrosoftEntraProvider) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetTenantId returns the TenantId field value
func (o *MicrosoftEntraProvider) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *MicrosoftEntraProvider) SetTenantId(v string) {
	o.TenantId = v
}

// GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field value if set, zero value otherwise.
func (o *MicrosoftEntraProvider) GetExcludeUsersServiceAccount() bool {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeUsersServiceAccount
}

// GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetExcludeUsersServiceAccountOk() (*bool, bool) {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		return nil, false
	}
	return o.ExcludeUsersServiceAccount, true
}

// HasExcludeUsersServiceAccount returns a boolean if a field has been set.
func (o *MicrosoftEntraProvider) HasExcludeUsersServiceAccount() bool {
	if o != nil && o.ExcludeUsersServiceAccount != nil {
		return true
	}

	return false
}

// SetExcludeUsersServiceAccount gets a reference to the given bool and assigns it to the ExcludeUsersServiceAccount field.
func (o *MicrosoftEntraProvider) SetExcludeUsersServiceAccount(v bool) {
	o.ExcludeUsersServiceAccount = &v
}

// GetFilterGroup returns the FilterGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftEntraProvider) GetFilterGroup() string {
	if o == nil || o.FilterGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.FilterGroup.Get()
}

// GetFilterGroupOk returns a tuple with the FilterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftEntraProvider) GetFilterGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterGroup.Get(), o.FilterGroup.IsSet()
}

// HasFilterGroup returns a boolean if a field has been set.
func (o *MicrosoftEntraProvider) HasFilterGroup() bool {
	if o != nil && o.FilterGroup.IsSet() {
		return true
	}

	return false
}

// SetFilterGroup gets a reference to the given NullableString and assigns it to the FilterGroup field.
func (o *MicrosoftEntraProvider) SetFilterGroup(v string) {
	o.FilterGroup.Set(&v)
}

// SetFilterGroupNil sets the value for FilterGroup to be an explicit nil
func (o *MicrosoftEntraProvider) SetFilterGroupNil() {
	o.FilterGroup.Set(nil)
}

// UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
func (o *MicrosoftEntraProvider) UnsetFilterGroup() {
	o.FilterGroup.Unset()
}

// GetUserDeleteAction returns the UserDeleteAction field value if set, zero value otherwise.
func (o *MicrosoftEntraProvider) GetUserDeleteAction() OutgoingSyncDeleteAction {
	if o == nil || o.UserDeleteAction == nil {
		var ret OutgoingSyncDeleteAction
		return ret
	}
	return *o.UserDeleteAction
}

// GetUserDeleteActionOk returns a tuple with the UserDeleteAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetUserDeleteActionOk() (*OutgoingSyncDeleteAction, bool) {
	if o == nil || o.UserDeleteAction == nil {
		return nil, false
	}
	return o.UserDeleteAction, true
}

// HasUserDeleteAction returns a boolean if a field has been set.
func (o *MicrosoftEntraProvider) HasUserDeleteAction() bool {
	if o != nil && o.UserDeleteAction != nil {
		return true
	}

	return false
}

// SetUserDeleteAction gets a reference to the given OutgoingSyncDeleteAction and assigns it to the UserDeleteAction field.
func (o *MicrosoftEntraProvider) SetUserDeleteAction(v OutgoingSyncDeleteAction) {
	o.UserDeleteAction = &v
}

// GetGroupDeleteAction returns the GroupDeleteAction field value if set, zero value otherwise.
func (o *MicrosoftEntraProvider) GetGroupDeleteAction() OutgoingSyncDeleteAction {
	if o == nil || o.GroupDeleteAction == nil {
		var ret OutgoingSyncDeleteAction
		return ret
	}
	return *o.GroupDeleteAction
}

// GetGroupDeleteActionOk returns a tuple with the GroupDeleteAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProvider) GetGroupDeleteActionOk() (*OutgoingSyncDeleteAction, bool) {
	if o == nil || o.GroupDeleteAction == nil {
		return nil, false
	}
	return o.GroupDeleteAction, true
}

// HasGroupDeleteAction returns a boolean if a field has been set.
func (o *MicrosoftEntraProvider) HasGroupDeleteAction() bool {
	if o != nil && o.GroupDeleteAction != nil {
		return true
	}

	return false
}

// SetGroupDeleteAction gets a reference to the given OutgoingSyncDeleteAction and assigns it to the GroupDeleteAction field.
func (o *MicrosoftEntraProvider) SetGroupDeleteAction(v OutgoingSyncDeleteAction) {
	o.GroupDeleteAction = &v
}

func (o MicrosoftEntraProvider) MarshalJSON() ([]byte, error) {
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
		toSerialize["assigned_backchannel_application_slug"] = o.AssignedBackchannelApplicationSlug
	}
	if true {
		toSerialize["assigned_backchannel_application_name"] = o.AssignedBackchannelApplicationName
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
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if true {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.ExcludeUsersServiceAccount != nil {
		toSerialize["exclude_users_service_account"] = o.ExcludeUsersServiceAccount
	}
	if o.FilterGroup.IsSet() {
		toSerialize["filter_group"] = o.FilterGroup.Get()
	}
	if o.UserDeleteAction != nil {
		toSerialize["user_delete_action"] = o.UserDeleteAction
	}
	if o.GroupDeleteAction != nil {
		toSerialize["group_delete_action"] = o.GroupDeleteAction
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftEntraProvider struct {
	value *MicrosoftEntraProvider
	isSet bool
}

func (v NullableMicrosoftEntraProvider) Get() *MicrosoftEntraProvider {
	return v.value
}

func (v *NullableMicrosoftEntraProvider) Set(val *MicrosoftEntraProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftEntraProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftEntraProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftEntraProvider(val *MicrosoftEntraProvider) *NullableMicrosoftEntraProvider {
	return &NullableMicrosoftEntraProvider{value: val, isSet: true}
}

func (v NullableMicrosoftEntraProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftEntraProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
