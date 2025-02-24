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

// PatchedMicrosoftEntraProviderRequest MicrosoftEntraProvider Serializer
type PatchedMicrosoftEntraProviderRequest struct {
	Name             *string  `json:"name,omitempty"`
	PropertyMappings []string `json:"property_mappings,omitempty"`
	// Property mappings used for group creation/updating.
	PropertyMappingsGroup      []string                  `json:"property_mappings_group,omitempty"`
	ClientId                   *string                   `json:"client_id,omitempty"`
	ClientSecret               *string                   `json:"client_secret,omitempty"`
	TenantId                   *string                   `json:"tenant_id,omitempty"`
	ExcludeUsersServiceAccount *bool                     `json:"exclude_users_service_account,omitempty"`
	FilterGroup                NullableString            `json:"filter_group,omitempty"`
	UserDeleteAction           *OutgoingSyncDeleteAction `json:"user_delete_action,omitempty"`
	GroupDeleteAction          *OutgoingSyncDeleteAction `json:"group_delete_action,omitempty"`
}

// NewPatchedMicrosoftEntraProviderRequest instantiates a new PatchedMicrosoftEntraProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedMicrosoftEntraProviderRequest() *PatchedMicrosoftEntraProviderRequest {
	this := PatchedMicrosoftEntraProviderRequest{}
	return &this
}

// NewPatchedMicrosoftEntraProviderRequestWithDefaults instantiates a new PatchedMicrosoftEntraProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedMicrosoftEntraProviderRequestWithDefaults() *PatchedMicrosoftEntraProviderRequest {
	this := PatchedMicrosoftEntraProviderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedMicrosoftEntraProviderRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMicrosoftEntraProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedMicrosoftEntraProviderRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedMicrosoftEntraProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *PatchedMicrosoftEntraProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMicrosoftEntraProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedMicrosoftEntraProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *PatchedMicrosoftEntraProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetPropertyMappingsGroup returns the PropertyMappingsGroup field value if set, zero value otherwise.
func (o *PatchedMicrosoftEntraProviderRequest) GetPropertyMappingsGroup() []string {
	if o == nil || o.PropertyMappingsGroup == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappingsGroup
}

// GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMicrosoftEntraProviderRequest) GetPropertyMappingsGroupOk() ([]string, bool) {
	if o == nil || o.PropertyMappingsGroup == nil {
		return nil, false
	}
	return o.PropertyMappingsGroup, true
}

// HasPropertyMappingsGroup returns a boolean if a field has been set.
func (o *PatchedMicrosoftEntraProviderRequest) HasPropertyMappingsGroup() bool {
	if o != nil && o.PropertyMappingsGroup != nil {
		return true
	}

	return false
}

// SetPropertyMappingsGroup gets a reference to the given []string and assigns it to the PropertyMappingsGroup field.
func (o *PatchedMicrosoftEntraProviderRequest) SetPropertyMappingsGroup(v []string) {
	o.PropertyMappingsGroup = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PatchedMicrosoftEntraProviderRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMicrosoftEntraProviderRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PatchedMicrosoftEntraProviderRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PatchedMicrosoftEntraProviderRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *PatchedMicrosoftEntraProviderRequest) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMicrosoftEntraProviderRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *PatchedMicrosoftEntraProviderRequest) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *PatchedMicrosoftEntraProviderRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *PatchedMicrosoftEntraProviderRequest) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMicrosoftEntraProviderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *PatchedMicrosoftEntraProviderRequest) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *PatchedMicrosoftEntraProviderRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field value if set, zero value otherwise.
func (o *PatchedMicrosoftEntraProviderRequest) GetExcludeUsersServiceAccount() bool {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeUsersServiceAccount
}

// GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMicrosoftEntraProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool) {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		return nil, false
	}
	return o.ExcludeUsersServiceAccount, true
}

// HasExcludeUsersServiceAccount returns a boolean if a field has been set.
func (o *PatchedMicrosoftEntraProviderRequest) HasExcludeUsersServiceAccount() bool {
	if o != nil && o.ExcludeUsersServiceAccount != nil {
		return true
	}

	return false
}

// SetExcludeUsersServiceAccount gets a reference to the given bool and assigns it to the ExcludeUsersServiceAccount field.
func (o *PatchedMicrosoftEntraProviderRequest) SetExcludeUsersServiceAccount(v bool) {
	o.ExcludeUsersServiceAccount = &v
}

// GetFilterGroup returns the FilterGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedMicrosoftEntraProviderRequest) GetFilterGroup() string {
	if o == nil || o.FilterGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.FilterGroup.Get()
}

// GetFilterGroupOk returns a tuple with the FilterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedMicrosoftEntraProviderRequest) GetFilterGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterGroup.Get(), o.FilterGroup.IsSet()
}

// HasFilterGroup returns a boolean if a field has been set.
func (o *PatchedMicrosoftEntraProviderRequest) HasFilterGroup() bool {
	if o != nil && o.FilterGroup.IsSet() {
		return true
	}

	return false
}

// SetFilterGroup gets a reference to the given NullableString and assigns it to the FilterGroup field.
func (o *PatchedMicrosoftEntraProviderRequest) SetFilterGroup(v string) {
	o.FilterGroup.Set(&v)
}

// SetFilterGroupNil sets the value for FilterGroup to be an explicit nil
func (o *PatchedMicrosoftEntraProviderRequest) SetFilterGroupNil() {
	o.FilterGroup.Set(nil)
}

// UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
func (o *PatchedMicrosoftEntraProviderRequest) UnsetFilterGroup() {
	o.FilterGroup.Unset()
}

// GetUserDeleteAction returns the UserDeleteAction field value if set, zero value otherwise.
func (o *PatchedMicrosoftEntraProviderRequest) GetUserDeleteAction() OutgoingSyncDeleteAction {
	if o == nil || o.UserDeleteAction == nil {
		var ret OutgoingSyncDeleteAction
		return ret
	}
	return *o.UserDeleteAction
}

// GetUserDeleteActionOk returns a tuple with the UserDeleteAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMicrosoftEntraProviderRequest) GetUserDeleteActionOk() (*OutgoingSyncDeleteAction, bool) {
	if o == nil || o.UserDeleteAction == nil {
		return nil, false
	}
	return o.UserDeleteAction, true
}

// HasUserDeleteAction returns a boolean if a field has been set.
func (o *PatchedMicrosoftEntraProviderRequest) HasUserDeleteAction() bool {
	if o != nil && o.UserDeleteAction != nil {
		return true
	}

	return false
}

// SetUserDeleteAction gets a reference to the given OutgoingSyncDeleteAction and assigns it to the UserDeleteAction field.
func (o *PatchedMicrosoftEntraProviderRequest) SetUserDeleteAction(v OutgoingSyncDeleteAction) {
	o.UserDeleteAction = &v
}

// GetGroupDeleteAction returns the GroupDeleteAction field value if set, zero value otherwise.
func (o *PatchedMicrosoftEntraProviderRequest) GetGroupDeleteAction() OutgoingSyncDeleteAction {
	if o == nil || o.GroupDeleteAction == nil {
		var ret OutgoingSyncDeleteAction
		return ret
	}
	return *o.GroupDeleteAction
}

// GetGroupDeleteActionOk returns a tuple with the GroupDeleteAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedMicrosoftEntraProviderRequest) GetGroupDeleteActionOk() (*OutgoingSyncDeleteAction, bool) {
	if o == nil || o.GroupDeleteAction == nil {
		return nil, false
	}
	return o.GroupDeleteAction, true
}

// HasGroupDeleteAction returns a boolean if a field has been set.
func (o *PatchedMicrosoftEntraProviderRequest) HasGroupDeleteAction() bool {
	if o != nil && o.GroupDeleteAction != nil {
		return true
	}

	return false
}

// SetGroupDeleteAction gets a reference to the given OutgoingSyncDeleteAction and assigns it to the GroupDeleteAction field.
func (o *PatchedMicrosoftEntraProviderRequest) SetGroupDeleteAction(v OutgoingSyncDeleteAction) {
	o.GroupDeleteAction = &v
}

func (o PatchedMicrosoftEntraProviderRequest) MarshalJSON() ([]byte, error) {
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
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.TenantId != nil {
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

type NullablePatchedMicrosoftEntraProviderRequest struct {
	value *PatchedMicrosoftEntraProviderRequest
	isSet bool
}

func (v NullablePatchedMicrosoftEntraProviderRequest) Get() *PatchedMicrosoftEntraProviderRequest {
	return v.value
}

func (v *NullablePatchedMicrosoftEntraProviderRequest) Set(val *PatchedMicrosoftEntraProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedMicrosoftEntraProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedMicrosoftEntraProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedMicrosoftEntraProviderRequest(val *PatchedMicrosoftEntraProviderRequest) *NullablePatchedMicrosoftEntraProviderRequest {
	return &NullablePatchedMicrosoftEntraProviderRequest{value: val, isSet: true}
}

func (v NullablePatchedMicrosoftEntraProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedMicrosoftEntraProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
