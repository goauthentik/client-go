/*
authentik

Making authentication simple.

API version: 2025.2.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorEmailStageRequest AuthenticatorEmailStage Serializer
type AuthenticatorEmailStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	FriendlyName  NullableString `json:"friendly_name,omitempty"`
	// When enabled, global Email connection settings will be used and connection settings below will be ignored.
	UseGlobalSettings *bool   `json:"use_global_settings,omitempty"`
	Host              *string `json:"host,omitempty"`
	Port              *int32  `json:"port,omitempty"`
	Username          *string `json:"username,omitempty"`
	Password          *string `json:"password,omitempty"`
	UseTls            *bool   `json:"use_tls,omitempty"`
	UseSsl            *bool   `json:"use_ssl,omitempty"`
	Timeout           *int32  `json:"timeout,omitempty"`
	FromAddress       *string `json:"from_address,omitempty"`
	Subject           *string `json:"subject,omitempty"`
	// Time the token sent is valid (Format: hours=3,minutes=17,seconds=300).
	TokenExpiry *string `json:"token_expiry,omitempty"`
	Template    *string `json:"template,omitempty"`
}

// NewAuthenticatorEmailStageRequest instantiates a new AuthenticatorEmailStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEmailStageRequest(name string) *AuthenticatorEmailStageRequest {
	this := AuthenticatorEmailStageRequest{}
	this.Name = name
	return &this
}

// NewAuthenticatorEmailStageRequestWithDefaults instantiates a new AuthenticatorEmailStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEmailStageRequestWithDefaults() *AuthenticatorEmailStageRequest {
	this := AuthenticatorEmailStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AuthenticatorEmailStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorEmailStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *AuthenticatorEmailStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorEmailStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorEmailStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *AuthenticatorEmailStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *AuthenticatorEmailStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *AuthenticatorEmailStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorEmailStageRequest) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorEmailStageRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *AuthenticatorEmailStageRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *AuthenticatorEmailStageRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *AuthenticatorEmailStageRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetUseGlobalSettings returns the UseGlobalSettings field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetUseGlobalSettings() bool {
	if o == nil || o.UseGlobalSettings == nil {
		var ret bool
		return ret
	}
	return *o.UseGlobalSettings
}

// GetUseGlobalSettingsOk returns a tuple with the UseGlobalSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetUseGlobalSettingsOk() (*bool, bool) {
	if o == nil || o.UseGlobalSettings == nil {
		return nil, false
	}
	return o.UseGlobalSettings, true
}

// HasUseGlobalSettings returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasUseGlobalSettings() bool {
	if o != nil && o.UseGlobalSettings != nil {
		return true
	}

	return false
}

// SetUseGlobalSettings gets a reference to the given bool and assigns it to the UseGlobalSettings field.
func (o *AuthenticatorEmailStageRequest) SetUseGlobalSettings(v bool) {
	o.UseGlobalSettings = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *AuthenticatorEmailStageRequest) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *AuthenticatorEmailStageRequest) SetPort(v int32) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *AuthenticatorEmailStageRequest) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AuthenticatorEmailStageRequest) SetPassword(v string) {
	o.Password = &v
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetUseTls() bool {
	if o == nil || o.UseTls == nil {
		var ret bool
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetUseTlsOk() (*bool, bool) {
	if o == nil || o.UseTls == nil {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasUseTls() bool {
	if o != nil && o.UseTls != nil {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given bool and assigns it to the UseTls field.
func (o *AuthenticatorEmailStageRequest) SetUseTls(v bool) {
	o.UseTls = &v
}

// GetUseSsl returns the UseSsl field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetUseSsl() bool {
	if o == nil || o.UseSsl == nil {
		var ret bool
		return ret
	}
	return *o.UseSsl
}

// GetUseSslOk returns a tuple with the UseSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetUseSslOk() (*bool, bool) {
	if o == nil || o.UseSsl == nil {
		return nil, false
	}
	return o.UseSsl, true
}

// HasUseSsl returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasUseSsl() bool {
	if o != nil && o.UseSsl != nil {
		return true
	}

	return false
}

// SetUseSsl gets a reference to the given bool and assigns it to the UseSsl field.
func (o *AuthenticatorEmailStageRequest) SetUseSsl(v bool) {
	o.UseSsl = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *AuthenticatorEmailStageRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetFromAddress() string {
	if o == nil || o.FromAddress == nil {
		var ret string
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetFromAddressOk() (*string, bool) {
	if o == nil || o.FromAddress == nil {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasFromAddress() bool {
	if o != nil && o.FromAddress != nil {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given string and assigns it to the FromAddress field.
func (o *AuthenticatorEmailStageRequest) SetFromAddress(v string) {
	o.FromAddress = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *AuthenticatorEmailStageRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetTokenExpiry returns the TokenExpiry field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetTokenExpiry() string {
	if o == nil || o.TokenExpiry == nil {
		var ret string
		return ret
	}
	return *o.TokenExpiry
}

// GetTokenExpiryOk returns a tuple with the TokenExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetTokenExpiryOk() (*string, bool) {
	if o == nil || o.TokenExpiry == nil {
		return nil, false
	}
	return o.TokenExpiry, true
}

// HasTokenExpiry returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasTokenExpiry() bool {
	if o != nil && o.TokenExpiry != nil {
		return true
	}

	return false
}

// SetTokenExpiry gets a reference to the given string and assigns it to the TokenExpiry field.
func (o *AuthenticatorEmailStageRequest) SetTokenExpiry(v string) {
	o.TokenExpiry = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *AuthenticatorEmailStageRequest) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEmailStageRequest) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *AuthenticatorEmailStageRequest) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *AuthenticatorEmailStageRequest) SetTemplate(v string) {
	o.Template = &v
}

func (o AuthenticatorEmailStageRequest) MarshalJSON() ([]byte, error) {
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
	if o.UseGlobalSettings != nil {
		toSerialize["use_global_settings"] = o.UseGlobalSettings
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.UseTls != nil {
		toSerialize["use_tls"] = o.UseTls
	}
	if o.UseSsl != nil {
		toSerialize["use_ssl"] = o.UseSsl
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.FromAddress != nil {
		toSerialize["from_address"] = o.FromAddress
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.TokenExpiry != nil {
		toSerialize["token_expiry"] = o.TokenExpiry
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorEmailStageRequest struct {
	value *AuthenticatorEmailStageRequest
	isSet bool
}

func (v NullableAuthenticatorEmailStageRequest) Get() *AuthenticatorEmailStageRequest {
	return v.value
}

func (v *NullableAuthenticatorEmailStageRequest) Set(val *AuthenticatorEmailStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEmailStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEmailStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEmailStageRequest(val *AuthenticatorEmailStageRequest) *NullableAuthenticatorEmailStageRequest {
	return &NullableAuthenticatorEmailStageRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorEmailStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEmailStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
