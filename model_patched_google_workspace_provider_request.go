/*
authentik

Making authentication simple.

API version: 2025.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedGoogleWorkspaceProviderRequest GoogleWorkspaceProvider Serializer
type PatchedGoogleWorkspaceProviderRequest struct {
	Name             *string  `json:"name,omitempty"`
	PropertyMappings []string `json:"property_mappings,omitempty"`
	// Property mappings used for group creation/updating.
	PropertyMappingsGroup      []string                  `json:"property_mappings_group,omitempty"`
	DelegatedSubject           *string                   `json:"delegated_subject,omitempty"`
	Credentials                interface{}               `json:"credentials,omitempty"`
	Scopes                     *string                   `json:"scopes,omitempty"`
	ExcludeUsersServiceAccount *bool                     `json:"exclude_users_service_account,omitempty"`
	FilterGroup                NullableString            `json:"filter_group,omitempty"`
	UserDeleteAction           *OutgoingSyncDeleteAction `json:"user_delete_action,omitempty"`
	GroupDeleteAction          *OutgoingSyncDeleteAction `json:"group_delete_action,omitempty"`
	DefaultGroupEmailDomain    *string                   `json:"default_group_email_domain,omitempty"`
}

// NewPatchedGoogleWorkspaceProviderRequest instantiates a new PatchedGoogleWorkspaceProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedGoogleWorkspaceProviderRequest() *PatchedGoogleWorkspaceProviderRequest {
	this := PatchedGoogleWorkspaceProviderRequest{}
	return &this
}

// NewPatchedGoogleWorkspaceProviderRequestWithDefaults instantiates a new PatchedGoogleWorkspaceProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedGoogleWorkspaceProviderRequestWithDefaults() *PatchedGoogleWorkspaceProviderRequest {
	this := PatchedGoogleWorkspaceProviderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedGoogleWorkspaceProviderRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedGoogleWorkspaceProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *PatchedGoogleWorkspaceProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *PatchedGoogleWorkspaceProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetPropertyMappingsGroup returns the PropertyMappingsGroup field value if set, zero value otherwise.
func (o *PatchedGoogleWorkspaceProviderRequest) GetPropertyMappingsGroup() []string {
	if o == nil || o.PropertyMappingsGroup == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappingsGroup
}

// GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) GetPropertyMappingsGroupOk() ([]string, bool) {
	if o == nil || o.PropertyMappingsGroup == nil {
		return nil, false
	}
	return o.PropertyMappingsGroup, true
}

// HasPropertyMappingsGroup returns a boolean if a field has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) HasPropertyMappingsGroup() bool {
	if o != nil && o.PropertyMappingsGroup != nil {
		return true
	}

	return false
}

// SetPropertyMappingsGroup gets a reference to the given []string and assigns it to the PropertyMappingsGroup field.
func (o *PatchedGoogleWorkspaceProviderRequest) SetPropertyMappingsGroup(v []string) {
	o.PropertyMappingsGroup = v
}

// GetDelegatedSubject returns the DelegatedSubject field value if set, zero value otherwise.
func (o *PatchedGoogleWorkspaceProviderRequest) GetDelegatedSubject() string {
	if o == nil || o.DelegatedSubject == nil {
		var ret string
		return ret
	}
	return *o.DelegatedSubject
}

// GetDelegatedSubjectOk returns a tuple with the DelegatedSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) GetDelegatedSubjectOk() (*string, bool) {
	if o == nil || o.DelegatedSubject == nil {
		return nil, false
	}
	return o.DelegatedSubject, true
}

// HasDelegatedSubject returns a boolean if a field has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) HasDelegatedSubject() bool {
	if o != nil && o.DelegatedSubject != nil {
		return true
	}

	return false
}

// SetDelegatedSubject gets a reference to the given string and assigns it to the DelegatedSubject field.
func (o *PatchedGoogleWorkspaceProviderRequest) SetDelegatedSubject(v string) {
	o.DelegatedSubject = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedGoogleWorkspaceProviderRequest) GetCredentials() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedGoogleWorkspaceProviderRequest) GetCredentialsOk() (*interface{}, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given interface{} and assigns it to the Credentials field.
func (o *PatchedGoogleWorkspaceProviderRequest) SetCredentials(v interface{}) {
	o.Credentials = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *PatchedGoogleWorkspaceProviderRequest) GetScopes() string {
	if o == nil || o.Scopes == nil {
		var ret string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) GetScopesOk() (*string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given string and assigns it to the Scopes field.
func (o *PatchedGoogleWorkspaceProviderRequest) SetScopes(v string) {
	o.Scopes = &v
}

// GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field value if set, zero value otherwise.
func (o *PatchedGoogleWorkspaceProviderRequest) GetExcludeUsersServiceAccount() bool {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeUsersServiceAccount
}

// GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool) {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		return nil, false
	}
	return o.ExcludeUsersServiceAccount, true
}

// HasExcludeUsersServiceAccount returns a boolean if a field has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) HasExcludeUsersServiceAccount() bool {
	if o != nil && o.ExcludeUsersServiceAccount != nil {
		return true
	}

	return false
}

// SetExcludeUsersServiceAccount gets a reference to the given bool and assigns it to the ExcludeUsersServiceAccount field.
func (o *PatchedGoogleWorkspaceProviderRequest) SetExcludeUsersServiceAccount(v bool) {
	o.ExcludeUsersServiceAccount = &v
}

// GetFilterGroup returns the FilterGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedGoogleWorkspaceProviderRequest) GetFilterGroup() string {
	if o == nil || o.FilterGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.FilterGroup.Get()
}

// GetFilterGroupOk returns a tuple with the FilterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedGoogleWorkspaceProviderRequest) GetFilterGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterGroup.Get(), o.FilterGroup.IsSet()
}

// HasFilterGroup returns a boolean if a field has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) HasFilterGroup() bool {
	if o != nil && o.FilterGroup.IsSet() {
		return true
	}

	return false
}

// SetFilterGroup gets a reference to the given NullableString and assigns it to the FilterGroup field.
func (o *PatchedGoogleWorkspaceProviderRequest) SetFilterGroup(v string) {
	o.FilterGroup.Set(&v)
}

// SetFilterGroupNil sets the value for FilterGroup to be an explicit nil
func (o *PatchedGoogleWorkspaceProviderRequest) SetFilterGroupNil() {
	o.FilterGroup.Set(nil)
}

// UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
func (o *PatchedGoogleWorkspaceProviderRequest) UnsetFilterGroup() {
	o.FilterGroup.Unset()
}

// GetUserDeleteAction returns the UserDeleteAction field value if set, zero value otherwise.
func (o *PatchedGoogleWorkspaceProviderRequest) GetUserDeleteAction() OutgoingSyncDeleteAction {
	if o == nil || o.UserDeleteAction == nil {
		var ret OutgoingSyncDeleteAction
		return ret
	}
	return *o.UserDeleteAction
}

// GetUserDeleteActionOk returns a tuple with the UserDeleteAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) GetUserDeleteActionOk() (*OutgoingSyncDeleteAction, bool) {
	if o == nil || o.UserDeleteAction == nil {
		return nil, false
	}
	return o.UserDeleteAction, true
}

// HasUserDeleteAction returns a boolean if a field has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) HasUserDeleteAction() bool {
	if o != nil && o.UserDeleteAction != nil {
		return true
	}

	return false
}

// SetUserDeleteAction gets a reference to the given OutgoingSyncDeleteAction and assigns it to the UserDeleteAction field.
func (o *PatchedGoogleWorkspaceProviderRequest) SetUserDeleteAction(v OutgoingSyncDeleteAction) {
	o.UserDeleteAction = &v
}

// GetGroupDeleteAction returns the GroupDeleteAction field value if set, zero value otherwise.
func (o *PatchedGoogleWorkspaceProviderRequest) GetGroupDeleteAction() OutgoingSyncDeleteAction {
	if o == nil || o.GroupDeleteAction == nil {
		var ret OutgoingSyncDeleteAction
		return ret
	}
	return *o.GroupDeleteAction
}

// GetGroupDeleteActionOk returns a tuple with the GroupDeleteAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) GetGroupDeleteActionOk() (*OutgoingSyncDeleteAction, bool) {
	if o == nil || o.GroupDeleteAction == nil {
		return nil, false
	}
	return o.GroupDeleteAction, true
}

// HasGroupDeleteAction returns a boolean if a field has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) HasGroupDeleteAction() bool {
	if o != nil && o.GroupDeleteAction != nil {
		return true
	}

	return false
}

// SetGroupDeleteAction gets a reference to the given OutgoingSyncDeleteAction and assigns it to the GroupDeleteAction field.
func (o *PatchedGoogleWorkspaceProviderRequest) SetGroupDeleteAction(v OutgoingSyncDeleteAction) {
	o.GroupDeleteAction = &v
}

// GetDefaultGroupEmailDomain returns the DefaultGroupEmailDomain field value if set, zero value otherwise.
func (o *PatchedGoogleWorkspaceProviderRequest) GetDefaultGroupEmailDomain() string {
	if o == nil || o.DefaultGroupEmailDomain == nil {
		var ret string
		return ret
	}
	return *o.DefaultGroupEmailDomain
}

// GetDefaultGroupEmailDomainOk returns a tuple with the DefaultGroupEmailDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) GetDefaultGroupEmailDomainOk() (*string, bool) {
	if o == nil || o.DefaultGroupEmailDomain == nil {
		return nil, false
	}
	return o.DefaultGroupEmailDomain, true
}

// HasDefaultGroupEmailDomain returns a boolean if a field has been set.
func (o *PatchedGoogleWorkspaceProviderRequest) HasDefaultGroupEmailDomain() bool {
	if o != nil && o.DefaultGroupEmailDomain != nil {
		return true
	}

	return false
}

// SetDefaultGroupEmailDomain gets a reference to the given string and assigns it to the DefaultGroupEmailDomain field.
func (o *PatchedGoogleWorkspaceProviderRequest) SetDefaultGroupEmailDomain(v string) {
	o.DefaultGroupEmailDomain = &v
}

func (o PatchedGoogleWorkspaceProviderRequest) MarshalJSON() ([]byte, error) {
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
	if o.DelegatedSubject != nil {
		toSerialize["delegated_subject"] = o.DelegatedSubject
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
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
	if o.DefaultGroupEmailDomain != nil {
		toSerialize["default_group_email_domain"] = o.DefaultGroupEmailDomain
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedGoogleWorkspaceProviderRequest struct {
	value *PatchedGoogleWorkspaceProviderRequest
	isSet bool
}

func (v NullablePatchedGoogleWorkspaceProviderRequest) Get() *PatchedGoogleWorkspaceProviderRequest {
	return v.value
}

func (v *NullablePatchedGoogleWorkspaceProviderRequest) Set(val *PatchedGoogleWorkspaceProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedGoogleWorkspaceProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedGoogleWorkspaceProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedGoogleWorkspaceProviderRequest(val *PatchedGoogleWorkspaceProviderRequest) *NullablePatchedGoogleWorkspaceProviderRequest {
	return &NullablePatchedGoogleWorkspaceProviderRequest{value: val, isSet: true}
}

func (v NullablePatchedGoogleWorkspaceProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedGoogleWorkspaceProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
