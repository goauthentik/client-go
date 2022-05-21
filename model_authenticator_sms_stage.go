/*
authentik

Making authentication simple.

API version: 2022.5.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorSMSStage AuthenticatorSMSStage Serializer
type AuthenticatorSMSStage struct {
	Pk                string  `json:"pk"`
	Name              string  `json:"name"`
	Component         string  `json:"component"`
	VerboseName       string  `json:"verbose_name"`
	VerboseNamePlural string  `json:"verbose_name_plural"`
	MetaModelName     string  `json:"meta_model_name"`
	FlowSet           *[]Flow `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	Provider      ProviderEnum   `json:"provider"`
	FromNumber    string         `json:"from_number"`
	AccountSid    string         `json:"account_sid"`
	Auth          string         `json:"auth"`
	AuthPassword  *string        `json:"auth_password,omitempty"`
	AuthType      *AuthTypeEnum  `json:"auth_type,omitempty"`
}

// NewAuthenticatorSMSStage instantiates a new AuthenticatorSMSStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorSMSStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, provider ProviderEnum, fromNumber string, accountSid string, auth string) *AuthenticatorSMSStage {
	this := AuthenticatorSMSStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Provider = provider
	this.FromNumber = fromNumber
	this.AccountSid = accountSid
	this.Auth = auth
	return &this
}

// NewAuthenticatorSMSStageWithDefaults instantiates a new AuthenticatorSMSStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorSMSStageWithDefaults() *AuthenticatorSMSStage {
	this := AuthenticatorSMSStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *AuthenticatorSMSStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *AuthenticatorSMSStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *AuthenticatorSMSStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorSMSStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *AuthenticatorSMSStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *AuthenticatorSMSStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *AuthenticatorSMSStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *AuthenticatorSMSStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *AuthenticatorSMSStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *AuthenticatorSMSStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *AuthenticatorSMSStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *AuthenticatorSMSStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorSMSStage) GetFlowSet() []Flow {
	if o == nil || o.FlowSet == nil {
		var ret []Flow
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetFlowSetOk() (*[]Flow, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorSMSStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []Flow and assigns it to the FlowSet field.
func (o *AuthenticatorSMSStage) SetFlowSet(v []Flow) {
	o.FlowSet = &v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorSMSStage) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorSMSStage) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *AuthenticatorSMSStage) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *AuthenticatorSMSStage) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *AuthenticatorSMSStage) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *AuthenticatorSMSStage) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetProvider returns the Provider field value
func (o *AuthenticatorSMSStage) GetProvider() ProviderEnum {
	if o == nil {
		var ret ProviderEnum
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetProviderOk() (*ProviderEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *AuthenticatorSMSStage) SetProvider(v ProviderEnum) {
	o.Provider = v
}

// GetFromNumber returns the FromNumber field value
func (o *AuthenticatorSMSStage) GetFromNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromNumber
}

// GetFromNumberOk returns a tuple with the FromNumber field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetFromNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromNumber, true
}

// SetFromNumber sets field value
func (o *AuthenticatorSMSStage) SetFromNumber(v string) {
	o.FromNumber = v
}

// GetAccountSid returns the AccountSid field value
func (o *AuthenticatorSMSStage) GetAccountSid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountSid
}

// GetAccountSidOk returns a tuple with the AccountSid field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetAccountSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountSid, true
}

// SetAccountSid sets field value
func (o *AuthenticatorSMSStage) SetAccountSid(v string) {
	o.AccountSid = v
}

// GetAuth returns the Auth field value
func (o *AuthenticatorSMSStage) GetAuth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetAuthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *AuthenticatorSMSStage) SetAuth(v string) {
	o.Auth = v
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *AuthenticatorSMSStage) GetAuthPassword() string {
	if o == nil || o.AuthPassword == nil {
		var ret string
		return ret
	}
	return *o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetAuthPasswordOk() (*string, bool) {
	if o == nil || o.AuthPassword == nil {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *AuthenticatorSMSStage) HasAuthPassword() bool {
	if o != nil && o.AuthPassword != nil {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *AuthenticatorSMSStage) SetAuthPassword(v string) {
	o.AuthPassword = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *AuthenticatorSMSStage) GetAuthType() AuthTypeEnum {
	if o == nil || o.AuthType == nil {
		var ret AuthTypeEnum
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorSMSStage) GetAuthTypeOk() (*AuthTypeEnum, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *AuthenticatorSMSStage) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given AuthTypeEnum and assigns it to the AuthType field.
func (o *AuthenticatorSMSStage) SetAuthType(v AuthTypeEnum) {
	o.AuthType = &v
}

func (o AuthenticatorSMSStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["component"] = o.Component
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
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ConfigureFlow.IsSet() {
		toSerialize["configure_flow"] = o.ConfigureFlow.Get()
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["from_number"] = o.FromNumber
	}
	if true {
		toSerialize["account_sid"] = o.AccountSid
	}
	if true {
		toSerialize["auth"] = o.Auth
	}
	if o.AuthPassword != nil {
		toSerialize["auth_password"] = o.AuthPassword
	}
	if o.AuthType != nil {
		toSerialize["auth_type"] = o.AuthType
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorSMSStage struct {
	value *AuthenticatorSMSStage
	isSet bool
}

func (v NullableAuthenticatorSMSStage) Get() *AuthenticatorSMSStage {
	return v.value
}

func (v *NullableAuthenticatorSMSStage) Set(val *AuthenticatorSMSStage) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorSMSStage) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorSMSStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorSMSStage(val *AuthenticatorSMSStage) *NullableAuthenticatorSMSStage {
	return &NullableAuthenticatorSMSStage{value: val, isSet: true}
}

func (v NullableAuthenticatorSMSStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorSMSStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
