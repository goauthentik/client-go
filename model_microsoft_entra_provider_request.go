/*
authentik

Making authentication simple.

API version: 2024.4.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MicrosoftEntraProviderRequest MicrosoftEntraProvider Serializer
type MicrosoftEntraProviderRequest struct {
	Name             string   `json:"name"`
	PropertyMappings []string `json:"property_mappings,omitempty"`
	// Property mappings used for group creation/updating.
	PropertyMappingsGroup      []string                  `json:"property_mappings_group,omitempty"`
	ClientId                   string                    `json:"client_id"`
	ClientSecret               string                    `json:"client_secret"`
	TenantId                   string                    `json:"tenant_id"`
	ExcludeUsersServiceAccount *bool                     `json:"exclude_users_service_account,omitempty"`
	FilterGroup                NullableString            `json:"filter_group,omitempty"`
	UserDeleteAction           *OutgoingSyncDeleteAction `json:"user_delete_action,omitempty"`
	GroupDeleteAction          *OutgoingSyncDeleteAction `json:"group_delete_action,omitempty"`
}

// NewMicrosoftEntraProviderRequest instantiates a new MicrosoftEntraProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftEntraProviderRequest(name string, clientId string, clientSecret string, tenantId string) *MicrosoftEntraProviderRequest {
	this := MicrosoftEntraProviderRequest{}
	this.Name = name
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.TenantId = tenantId
	return &this
}

// NewMicrosoftEntraProviderRequestWithDefaults instantiates a new MicrosoftEntraProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftEntraProviderRequestWithDefaults() *MicrosoftEntraProviderRequest {
	this := MicrosoftEntraProviderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *MicrosoftEntraProviderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MicrosoftEntraProviderRequest) SetName(v string) {
	o.Name = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *MicrosoftEntraProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *MicrosoftEntraProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *MicrosoftEntraProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetPropertyMappingsGroup returns the PropertyMappingsGroup field value if set, zero value otherwise.
func (o *MicrosoftEntraProviderRequest) GetPropertyMappingsGroup() []string {
	if o == nil || o.PropertyMappingsGroup == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappingsGroup
}

// GetPropertyMappingsGroupOk returns a tuple with the PropertyMappingsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderRequest) GetPropertyMappingsGroupOk() ([]string, bool) {
	if o == nil || o.PropertyMappingsGroup == nil {
		return nil, false
	}
	return o.PropertyMappingsGroup, true
}

// HasPropertyMappingsGroup returns a boolean if a field has been set.
func (o *MicrosoftEntraProviderRequest) HasPropertyMappingsGroup() bool {
	if o != nil && o.PropertyMappingsGroup != nil {
		return true
	}

	return false
}

// SetPropertyMappingsGroup gets a reference to the given []string and assigns it to the PropertyMappingsGroup field.
func (o *MicrosoftEntraProviderRequest) SetPropertyMappingsGroup(v []string) {
	o.PropertyMappingsGroup = v
}

// GetClientId returns the ClientId field value
func (o *MicrosoftEntraProviderRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *MicrosoftEntraProviderRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *MicrosoftEntraProviderRequest) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderRequest) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *MicrosoftEntraProviderRequest) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetTenantId returns the TenantId field value
func (o *MicrosoftEntraProviderRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *MicrosoftEntraProviderRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetExcludeUsersServiceAccount returns the ExcludeUsersServiceAccount field value if set, zero value otherwise.
func (o *MicrosoftEntraProviderRequest) GetExcludeUsersServiceAccount() bool {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeUsersServiceAccount
}

// GetExcludeUsersServiceAccountOk returns a tuple with the ExcludeUsersServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderRequest) GetExcludeUsersServiceAccountOk() (*bool, bool) {
	if o == nil || o.ExcludeUsersServiceAccount == nil {
		return nil, false
	}
	return o.ExcludeUsersServiceAccount, true
}

// HasExcludeUsersServiceAccount returns a boolean if a field has been set.
func (o *MicrosoftEntraProviderRequest) HasExcludeUsersServiceAccount() bool {
	if o != nil && o.ExcludeUsersServiceAccount != nil {
		return true
	}

	return false
}

// SetExcludeUsersServiceAccount gets a reference to the given bool and assigns it to the ExcludeUsersServiceAccount field.
func (o *MicrosoftEntraProviderRequest) SetExcludeUsersServiceAccount(v bool) {
	o.ExcludeUsersServiceAccount = &v
}

// GetFilterGroup returns the FilterGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftEntraProviderRequest) GetFilterGroup() string {
	if o == nil || o.FilterGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.FilterGroup.Get()
}

// GetFilterGroupOk returns a tuple with the FilterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftEntraProviderRequest) GetFilterGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterGroup.Get(), o.FilterGroup.IsSet()
}

// HasFilterGroup returns a boolean if a field has been set.
func (o *MicrosoftEntraProviderRequest) HasFilterGroup() bool {
	if o != nil && o.FilterGroup.IsSet() {
		return true
	}

	return false
}

// SetFilterGroup gets a reference to the given NullableString and assigns it to the FilterGroup field.
func (o *MicrosoftEntraProviderRequest) SetFilterGroup(v string) {
	o.FilterGroup.Set(&v)
}

// SetFilterGroupNil sets the value for FilterGroup to be an explicit nil
func (o *MicrosoftEntraProviderRequest) SetFilterGroupNil() {
	o.FilterGroup.Set(nil)
}

// UnsetFilterGroup ensures that no value is present for FilterGroup, not even an explicit nil
func (o *MicrosoftEntraProviderRequest) UnsetFilterGroup() {
	o.FilterGroup.Unset()
}

// GetUserDeleteAction returns the UserDeleteAction field value if set, zero value otherwise.
func (o *MicrosoftEntraProviderRequest) GetUserDeleteAction() OutgoingSyncDeleteAction {
	if o == nil || o.UserDeleteAction == nil {
		var ret OutgoingSyncDeleteAction
		return ret
	}
	return *o.UserDeleteAction
}

// GetUserDeleteActionOk returns a tuple with the UserDeleteAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderRequest) GetUserDeleteActionOk() (*OutgoingSyncDeleteAction, bool) {
	if o == nil || o.UserDeleteAction == nil {
		return nil, false
	}
	return o.UserDeleteAction, true
}

// HasUserDeleteAction returns a boolean if a field has been set.
func (o *MicrosoftEntraProviderRequest) HasUserDeleteAction() bool {
	if o != nil && o.UserDeleteAction != nil {
		return true
	}

	return false
}

// SetUserDeleteAction gets a reference to the given OutgoingSyncDeleteAction and assigns it to the UserDeleteAction field.
func (o *MicrosoftEntraProviderRequest) SetUserDeleteAction(v OutgoingSyncDeleteAction) {
	o.UserDeleteAction = &v
}

// GetGroupDeleteAction returns the GroupDeleteAction field value if set, zero value otherwise.
func (o *MicrosoftEntraProviderRequest) GetGroupDeleteAction() OutgoingSyncDeleteAction {
	if o == nil || o.GroupDeleteAction == nil {
		var ret OutgoingSyncDeleteAction
		return ret
	}
	return *o.GroupDeleteAction
}

// GetGroupDeleteActionOk returns a tuple with the GroupDeleteAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftEntraProviderRequest) GetGroupDeleteActionOk() (*OutgoingSyncDeleteAction, bool) {
	if o == nil || o.GroupDeleteAction == nil {
		return nil, false
	}
	return o.GroupDeleteAction, true
}

// HasGroupDeleteAction returns a boolean if a field has been set.
func (o *MicrosoftEntraProviderRequest) HasGroupDeleteAction() bool {
	if o != nil && o.GroupDeleteAction != nil {
		return true
	}

	return false
}

// SetGroupDeleteAction gets a reference to the given OutgoingSyncDeleteAction and assigns it to the GroupDeleteAction field.
func (o *MicrosoftEntraProviderRequest) SetGroupDeleteAction(v OutgoingSyncDeleteAction) {
	o.GroupDeleteAction = &v
}

func (o MicrosoftEntraProviderRequest) MarshalJSON() ([]byte, error) {
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

type NullableMicrosoftEntraProviderRequest struct {
	value *MicrosoftEntraProviderRequest
	isSet bool
}

func (v NullableMicrosoftEntraProviderRequest) Get() *MicrosoftEntraProviderRequest {
	return v.value
}

func (v *NullableMicrosoftEntraProviderRequest) Set(val *MicrosoftEntraProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftEntraProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftEntraProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftEntraProviderRequest(val *MicrosoftEntraProviderRequest) *NullableMicrosoftEntraProviderRequest {
	return &NullableMicrosoftEntraProviderRequest{value: val, isSet: true}
}

func (v NullableMicrosoftEntraProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftEntraProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}