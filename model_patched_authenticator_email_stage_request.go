/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedAuthenticatorEmailStageRequest AuthenticatorEmailStage Serializer
type PatchedAuthenticatorEmailStageRequest struct {
	Name    *string          `json:"name,omitempty"`
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

// NewPatchedAuthenticatorEmailStageRequest instantiates a new PatchedAuthenticatorEmailStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAuthenticatorEmailStageRequest() *PatchedAuthenticatorEmailStageRequest {
	this := PatchedAuthenticatorEmailStageRequest{}
	return &this
}

// NewPatchedAuthenticatorEmailStageRequestWithDefaults instantiates a new PatchedAuthenticatorEmailStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAuthenticatorEmailStageRequestWithDefaults() *PatchedAuthenticatorEmailStageRequest {
	this := PatchedAuthenticatorEmailStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAuthenticatorEmailStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedAuthenticatorEmailStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorEmailStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorEmailStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *PatchedAuthenticatorEmailStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *PatchedAuthenticatorEmailStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *PatchedAuthenticatorEmailStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorEmailStageRequest) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorEmailStageRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *PatchedAuthenticatorEmailStageRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *PatchedAuthenticatorEmailStageRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *PatchedAuthenticatorEmailStageRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetUseGlobalSettings returns the UseGlobalSettings field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetUseGlobalSettings() bool {
	if o == nil || o.UseGlobalSettings == nil {
		var ret bool
		return ret
	}
	return *o.UseGlobalSettings
}

// GetUseGlobalSettingsOk returns a tuple with the UseGlobalSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetUseGlobalSettingsOk() (*bool, bool) {
	if o == nil || o.UseGlobalSettings == nil {
		return nil, false
	}
	return o.UseGlobalSettings, true
}

// HasUseGlobalSettings returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasUseGlobalSettings() bool {
	if o != nil && o.UseGlobalSettings != nil {
		return true
	}

	return false
}

// SetUseGlobalSettings gets a reference to the given bool and assigns it to the UseGlobalSettings field.
func (o *PatchedAuthenticatorEmailStageRequest) SetUseGlobalSettings(v bool) {
	o.UseGlobalSettings = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *PatchedAuthenticatorEmailStageRequest) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *PatchedAuthenticatorEmailStageRequest) SetPort(v int32) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PatchedAuthenticatorEmailStageRequest) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PatchedAuthenticatorEmailStageRequest) SetPassword(v string) {
	o.Password = &v
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetUseTls() bool {
	if o == nil || o.UseTls == nil {
		var ret bool
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetUseTlsOk() (*bool, bool) {
	if o == nil || o.UseTls == nil {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasUseTls() bool {
	if o != nil && o.UseTls != nil {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given bool and assigns it to the UseTls field.
func (o *PatchedAuthenticatorEmailStageRequest) SetUseTls(v bool) {
	o.UseTls = &v
}

// GetUseSsl returns the UseSsl field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetUseSsl() bool {
	if o == nil || o.UseSsl == nil {
		var ret bool
		return ret
	}
	return *o.UseSsl
}

// GetUseSslOk returns a tuple with the UseSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetUseSslOk() (*bool, bool) {
	if o == nil || o.UseSsl == nil {
		return nil, false
	}
	return o.UseSsl, true
}

// HasUseSsl returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasUseSsl() bool {
	if o != nil && o.UseSsl != nil {
		return true
	}

	return false
}

// SetUseSsl gets a reference to the given bool and assigns it to the UseSsl field.
func (o *PatchedAuthenticatorEmailStageRequest) SetUseSsl(v bool) {
	o.UseSsl = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *PatchedAuthenticatorEmailStageRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetFromAddress() string {
	if o == nil || o.FromAddress == nil {
		var ret string
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetFromAddressOk() (*string, bool) {
	if o == nil || o.FromAddress == nil {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasFromAddress() bool {
	if o != nil && o.FromAddress != nil {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given string and assigns it to the FromAddress field.
func (o *PatchedAuthenticatorEmailStageRequest) SetFromAddress(v string) {
	o.FromAddress = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *PatchedAuthenticatorEmailStageRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetTokenExpiry returns the TokenExpiry field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetTokenExpiry() string {
	if o == nil || o.TokenExpiry == nil {
		var ret string
		return ret
	}
	return *o.TokenExpiry
}

// GetTokenExpiryOk returns a tuple with the TokenExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetTokenExpiryOk() (*string, bool) {
	if o == nil || o.TokenExpiry == nil {
		return nil, false
	}
	return o.TokenExpiry, true
}

// HasTokenExpiry returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasTokenExpiry() bool {
	if o != nil && o.TokenExpiry != nil {
		return true
	}

	return false
}

// SetTokenExpiry gets a reference to the given string and assigns it to the TokenExpiry field.
func (o *PatchedAuthenticatorEmailStageRequest) SetTokenExpiry(v string) {
	o.TokenExpiry = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEmailStageRequest) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEmailStageRequest) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEmailStageRequest) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *PatchedAuthenticatorEmailStageRequest) SetTemplate(v string) {
	o.Template = &v
}

func (o PatchedAuthenticatorEmailStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
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

type NullablePatchedAuthenticatorEmailStageRequest struct {
	value *PatchedAuthenticatorEmailStageRequest
	isSet bool
}

func (v NullablePatchedAuthenticatorEmailStageRequest) Get() *PatchedAuthenticatorEmailStageRequest {
	return v.value
}

func (v *NullablePatchedAuthenticatorEmailStageRequest) Set(val *PatchedAuthenticatorEmailStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAuthenticatorEmailStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAuthenticatorEmailStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAuthenticatorEmailStageRequest(val *PatchedAuthenticatorEmailStageRequest) *NullablePatchedAuthenticatorEmailStageRequest {
	return &NullablePatchedAuthenticatorEmailStageRequest{value: val, isSet: true}
}

func (v NullablePatchedAuthenticatorEmailStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAuthenticatorEmailStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
