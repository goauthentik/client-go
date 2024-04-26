/*
authentik

Making authentication simple.

API version: 2024.4.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorDuoStageRequest AuthenticatorDuoStage Serializer
type AuthenticatorDuoStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow       NullableString `json:"configure_flow,omitempty"`
	FriendlyName        NullableString `json:"friendly_name,omitempty"`
	ClientId            string         `json:"client_id"`
	ClientSecret        string         `json:"client_secret"`
	ApiHostname         string         `json:"api_hostname"`
	AdminIntegrationKey *string        `json:"admin_integration_key,omitempty"`
	AdminSecretKey      *string        `json:"admin_secret_key,omitempty"`
}

// NewAuthenticatorDuoStageRequest instantiates a new AuthenticatorDuoStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorDuoStageRequest(name string, clientId string, clientSecret string, apiHostname string) *AuthenticatorDuoStageRequest {
	this := AuthenticatorDuoStageRequest{}
	this.Name = name
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.ApiHostname = apiHostname
	return &this
}

// NewAuthenticatorDuoStageRequestWithDefaults instantiates a new AuthenticatorDuoStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorDuoStageRequestWithDefaults() *AuthenticatorDuoStageRequest {
	this := AuthenticatorDuoStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AuthenticatorDuoStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorDuoStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorDuoStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorDuoStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *AuthenticatorDuoStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorDuoStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorDuoStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *AuthenticatorDuoStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *AuthenticatorDuoStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *AuthenticatorDuoStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *AuthenticatorDuoStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorDuoStageRequest) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorDuoStageRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *AuthenticatorDuoStageRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *AuthenticatorDuoStageRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *AuthenticatorDuoStageRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *AuthenticatorDuoStageRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetClientId returns the ClientId field value
func (o *AuthenticatorDuoStageRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoStageRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *AuthenticatorDuoStageRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *AuthenticatorDuoStageRequest) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoStageRequest) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *AuthenticatorDuoStageRequest) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetApiHostname returns the ApiHostname field value
func (o *AuthenticatorDuoStageRequest) GetApiHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiHostname
}

// GetApiHostnameOk returns a tuple with the ApiHostname field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoStageRequest) GetApiHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiHostname, true
}

// SetApiHostname sets field value
func (o *AuthenticatorDuoStageRequest) SetApiHostname(v string) {
	o.ApiHostname = v
}

// GetAdminIntegrationKey returns the AdminIntegrationKey field value if set, zero value otherwise.
func (o *AuthenticatorDuoStageRequest) GetAdminIntegrationKey() string {
	if o == nil || o.AdminIntegrationKey == nil {
		var ret string
		return ret
	}
	return *o.AdminIntegrationKey
}

// GetAdminIntegrationKeyOk returns a tuple with the AdminIntegrationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoStageRequest) GetAdminIntegrationKeyOk() (*string, bool) {
	if o == nil || o.AdminIntegrationKey == nil {
		return nil, false
	}
	return o.AdminIntegrationKey, true
}

// HasAdminIntegrationKey returns a boolean if a field has been set.
func (o *AuthenticatorDuoStageRequest) HasAdminIntegrationKey() bool {
	if o != nil && o.AdminIntegrationKey != nil {
		return true
	}

	return false
}

// SetAdminIntegrationKey gets a reference to the given string and assigns it to the AdminIntegrationKey field.
func (o *AuthenticatorDuoStageRequest) SetAdminIntegrationKey(v string) {
	o.AdminIntegrationKey = &v
}

// GetAdminSecretKey returns the AdminSecretKey field value if set, zero value otherwise.
func (o *AuthenticatorDuoStageRequest) GetAdminSecretKey() string {
	if o == nil || o.AdminSecretKey == nil {
		var ret string
		return ret
	}
	return *o.AdminSecretKey
}

// GetAdminSecretKeyOk returns a tuple with the AdminSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoStageRequest) GetAdminSecretKeyOk() (*string, bool) {
	if o == nil || o.AdminSecretKey == nil {
		return nil, false
	}
	return o.AdminSecretKey, true
}

// HasAdminSecretKey returns a boolean if a field has been set.
func (o *AuthenticatorDuoStageRequest) HasAdminSecretKey() bool {
	if o != nil && o.AdminSecretKey != nil {
		return true
	}

	return false
}

// SetAdminSecretKey gets a reference to the given string and assigns it to the AdminSecretKey field.
func (o *AuthenticatorDuoStageRequest) SetAdminSecretKey(v string) {
	o.AdminSecretKey = &v
}

func (o AuthenticatorDuoStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ConfigureFlow.IsSet() {
		toSerialize["configure_flow"] = o.ConfigureFlow.Get()
	}
	if o.FriendlyName.IsSet() {
		toSerialize["friendly_name"] = o.FriendlyName.Get()
	}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if true {
		toSerialize["api_hostname"] = o.ApiHostname
	}
	if o.AdminIntegrationKey != nil {
		toSerialize["admin_integration_key"] = o.AdminIntegrationKey
	}
	if o.AdminSecretKey != nil {
		toSerialize["admin_secret_key"] = o.AdminSecretKey
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorDuoStageRequest struct {
	value *AuthenticatorDuoStageRequest
	isSet bool
}

func (v NullableAuthenticatorDuoStageRequest) Get() *AuthenticatorDuoStageRequest {
	return v.value
}

func (v *NullableAuthenticatorDuoStageRequest) Set(val *AuthenticatorDuoStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorDuoStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorDuoStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorDuoStageRequest(val *AuthenticatorDuoStageRequest) *NullableAuthenticatorDuoStageRequest {
	return &NullableAuthenticatorDuoStageRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorDuoStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorDuoStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
