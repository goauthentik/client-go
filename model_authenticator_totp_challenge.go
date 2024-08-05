/*
authentik

Making authentication simple.

API version: 2024.6.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorTOTPChallenge TOTP Setup challenge
type AuthenticatorTOTPChallenge struct {
	FlowInfo          *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component         *string                   `json:"component,omitempty"`
	ResponseErrors    *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	PendingUser       string                    `json:"pending_user"`
	PendingUserAvatar string                    `json:"pending_user_avatar"`
	ConfigUrl         string                    `json:"config_url"`
}

// NewAuthenticatorTOTPChallenge instantiates a new AuthenticatorTOTPChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorTOTPChallenge(pendingUser string, pendingUserAvatar string, configUrl string) *AuthenticatorTOTPChallenge {
	this := AuthenticatorTOTPChallenge{}
	var component string = "ak-stage-authenticator-totp"
	this.Component = &component
	this.PendingUser = pendingUser
	this.PendingUserAvatar = pendingUserAvatar
	this.ConfigUrl = configUrl
	return &this
}

// NewAuthenticatorTOTPChallengeWithDefaults instantiates a new AuthenticatorTOTPChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorTOTPChallengeWithDefaults() *AuthenticatorTOTPChallenge {
	this := AuthenticatorTOTPChallenge{}
	var component string = "ak-stage-authenticator-totp"
	this.Component = &component
	return &this
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *AuthenticatorTOTPChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorTOTPChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *AuthenticatorTOTPChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *AuthenticatorTOTPChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorTOTPChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorTOTPChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorTOTPChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorTOTPChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *AuthenticatorTOTPChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorTOTPChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *AuthenticatorTOTPChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *AuthenticatorTOTPChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetPendingUser returns the PendingUser field value
func (o *AuthenticatorTOTPChallenge) GetPendingUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUser
}

// GetPendingUserOk returns a tuple with the PendingUser field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorTOTPChallenge) GetPendingUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUser, true
}

// SetPendingUser sets field value
func (o *AuthenticatorTOTPChallenge) SetPendingUser(v string) {
	o.PendingUser = v
}

// GetPendingUserAvatar returns the PendingUserAvatar field value
func (o *AuthenticatorTOTPChallenge) GetPendingUserAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUserAvatar
}

// GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorTOTPChallenge) GetPendingUserAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUserAvatar, true
}

// SetPendingUserAvatar sets field value
func (o *AuthenticatorTOTPChallenge) SetPendingUserAvatar(v string) {
	o.PendingUserAvatar = v
}

// GetConfigUrl returns the ConfigUrl field value
func (o *AuthenticatorTOTPChallenge) GetConfigUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigUrl
}

// GetConfigUrlOk returns a tuple with the ConfigUrl field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorTOTPChallenge) GetConfigUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigUrl, true
}

// SetConfigUrl sets field value
func (o *AuthenticatorTOTPChallenge) SetConfigUrl(v string) {
	o.ConfigUrl = v
}

func (o AuthenticatorTOTPChallenge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FlowInfo != nil {
		toSerialize["flow_info"] = o.FlowInfo
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.ResponseErrors != nil {
		toSerialize["response_errors"] = o.ResponseErrors
	}
	if true {
		toSerialize["pending_user"] = o.PendingUser
	}
	if true {
		toSerialize["pending_user_avatar"] = o.PendingUserAvatar
	}
	if true {
		toSerialize["config_url"] = o.ConfigUrl
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorTOTPChallenge struct {
	value *AuthenticatorTOTPChallenge
	isSet bool
}

func (v NullableAuthenticatorTOTPChallenge) Get() *AuthenticatorTOTPChallenge {
	return v.value
}

func (v *NullableAuthenticatorTOTPChallenge) Set(val *AuthenticatorTOTPChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorTOTPChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorTOTPChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorTOTPChallenge(val *AuthenticatorTOTPChallenge) *NullableAuthenticatorTOTPChallenge {
	return &NullableAuthenticatorTOTPChallenge{value: val, isSet: true}
}

func (v NullableAuthenticatorTOTPChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorTOTPChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
