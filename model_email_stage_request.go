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

// EmailStageRequest EmailStage Serializer
type EmailStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
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
	// Time in minutes the token sent is valid.
	TokenExpiry *int32  `json:"token_expiry,omitempty"`
	Subject     *string `json:"subject,omitempty"`
	Template    *string `json:"template,omitempty"`
	// Activate users upon completion of stage.
	ActivateUserOnSuccess *bool `json:"activate_user_on_success,omitempty"`
}

// NewEmailStageRequest instantiates a new EmailStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailStageRequest(name string) *EmailStageRequest {
	this := EmailStageRequest{}
	this.Name = name
	return &this
}

// NewEmailStageRequestWithDefaults instantiates a new EmailStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailStageRequestWithDefaults() *EmailStageRequest {
	this := EmailStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *EmailStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EmailStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *EmailStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *EmailStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *EmailStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetUseGlobalSettings returns the UseGlobalSettings field value if set, zero value otherwise.
func (o *EmailStageRequest) GetUseGlobalSettings() bool {
	if o == nil || o.UseGlobalSettings == nil {
		var ret bool
		return ret
	}
	return *o.UseGlobalSettings
}

// GetUseGlobalSettingsOk returns a tuple with the UseGlobalSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetUseGlobalSettingsOk() (*bool, bool) {
	if o == nil || o.UseGlobalSettings == nil {
		return nil, false
	}
	return o.UseGlobalSettings, true
}

// HasUseGlobalSettings returns a boolean if a field has been set.
func (o *EmailStageRequest) HasUseGlobalSettings() bool {
	if o != nil && o.UseGlobalSettings != nil {
		return true
	}

	return false
}

// SetUseGlobalSettings gets a reference to the given bool and assigns it to the UseGlobalSettings field.
func (o *EmailStageRequest) SetUseGlobalSettings(v bool) {
	o.UseGlobalSettings = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *EmailStageRequest) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *EmailStageRequest) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *EmailStageRequest) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *EmailStageRequest) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *EmailStageRequest) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *EmailStageRequest) SetPort(v int32) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *EmailStageRequest) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *EmailStageRequest) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *EmailStageRequest) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *EmailStageRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *EmailStageRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *EmailStageRequest) SetPassword(v string) {
	o.Password = &v
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *EmailStageRequest) GetUseTls() bool {
	if o == nil || o.UseTls == nil {
		var ret bool
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetUseTlsOk() (*bool, bool) {
	if o == nil || o.UseTls == nil {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *EmailStageRequest) HasUseTls() bool {
	if o != nil && o.UseTls != nil {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given bool and assigns it to the UseTls field.
func (o *EmailStageRequest) SetUseTls(v bool) {
	o.UseTls = &v
}

// GetUseSsl returns the UseSsl field value if set, zero value otherwise.
func (o *EmailStageRequest) GetUseSsl() bool {
	if o == nil || o.UseSsl == nil {
		var ret bool
		return ret
	}
	return *o.UseSsl
}

// GetUseSslOk returns a tuple with the UseSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetUseSslOk() (*bool, bool) {
	if o == nil || o.UseSsl == nil {
		return nil, false
	}
	return o.UseSsl, true
}

// HasUseSsl returns a boolean if a field has been set.
func (o *EmailStageRequest) HasUseSsl() bool {
	if o != nil && o.UseSsl != nil {
		return true
	}

	return false
}

// SetUseSsl gets a reference to the given bool and assigns it to the UseSsl field.
func (o *EmailStageRequest) SetUseSsl(v bool) {
	o.UseSsl = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *EmailStageRequest) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *EmailStageRequest) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *EmailStageRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *EmailStageRequest) GetFromAddress() string {
	if o == nil || o.FromAddress == nil {
		var ret string
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetFromAddressOk() (*string, bool) {
	if o == nil || o.FromAddress == nil {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *EmailStageRequest) HasFromAddress() bool {
	if o != nil && o.FromAddress != nil {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given string and assigns it to the FromAddress field.
func (o *EmailStageRequest) SetFromAddress(v string) {
	o.FromAddress = &v
}

// GetTokenExpiry returns the TokenExpiry field value if set, zero value otherwise.
func (o *EmailStageRequest) GetTokenExpiry() int32 {
	if o == nil || o.TokenExpiry == nil {
		var ret int32
		return ret
	}
	return *o.TokenExpiry
}

// GetTokenExpiryOk returns a tuple with the TokenExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetTokenExpiryOk() (*int32, bool) {
	if o == nil || o.TokenExpiry == nil {
		return nil, false
	}
	return o.TokenExpiry, true
}

// HasTokenExpiry returns a boolean if a field has been set.
func (o *EmailStageRequest) HasTokenExpiry() bool {
	if o != nil && o.TokenExpiry != nil {
		return true
	}

	return false
}

// SetTokenExpiry gets a reference to the given int32 and assigns it to the TokenExpiry field.
func (o *EmailStageRequest) SetTokenExpiry(v int32) {
	o.TokenExpiry = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *EmailStageRequest) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailStageRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *EmailStageRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *EmailStageRequest) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *EmailStageRequest) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *EmailStageRequest) SetTemplate(v string) {
	o.Template = &v
}

// GetActivateUserOnSuccess returns the ActivateUserOnSuccess field value if set, zero value otherwise.
func (o *EmailStageRequest) GetActivateUserOnSuccess() bool {
	if o == nil || o.ActivateUserOnSuccess == nil {
		var ret bool
		return ret
	}
	return *o.ActivateUserOnSuccess
}

// GetActivateUserOnSuccessOk returns a tuple with the ActivateUserOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStageRequest) GetActivateUserOnSuccessOk() (*bool, bool) {
	if o == nil || o.ActivateUserOnSuccess == nil {
		return nil, false
	}
	return o.ActivateUserOnSuccess, true
}

// HasActivateUserOnSuccess returns a boolean if a field has been set.
func (o *EmailStageRequest) HasActivateUserOnSuccess() bool {
	if o != nil && o.ActivateUserOnSuccess != nil {
		return true
	}

	return false
}

// SetActivateUserOnSuccess gets a reference to the given bool and assigns it to the ActivateUserOnSuccess field.
func (o *EmailStageRequest) SetActivateUserOnSuccess(v bool) {
	o.ActivateUserOnSuccess = &v
}

func (o EmailStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
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
	if o.TokenExpiry != nil {
		toSerialize["token_expiry"] = o.TokenExpiry
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.ActivateUserOnSuccess != nil {
		toSerialize["activate_user_on_success"] = o.ActivateUserOnSuccess
	}
	return json.Marshal(toSerialize)
}

type NullableEmailStageRequest struct {
	value *EmailStageRequest
	isSet bool
}

func (v NullableEmailStageRequest) Get() *EmailStageRequest {
	return v.value
}

func (v *NullableEmailStageRequest) Set(val *EmailStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailStageRequest(val *EmailStageRequest) *NullableEmailStageRequest {
	return &NullableEmailStageRequest{value: val, isSet: true}
}

func (v NullableEmailStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
