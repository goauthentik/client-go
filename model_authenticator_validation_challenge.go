/*
authentik

Making authentication simple.

API version: 2022.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorValidationChallenge Authenticator challenge
type AuthenticatorValidationChallenge struct {
	Type                ChallengeChoices          `json:"type"`
	FlowInfo            *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component           *string                   `json:"component,omitempty"`
	ResponseErrors      *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	PendingUser         string                    `json:"pending_user"`
	PendingUserAvatar   string                    `json:"pending_user_avatar"`
	DeviceChallenges    []DeviceChallenge         `json:"device_challenges"`
	ConfigurationStages []SelectableStage         `json:"configuration_stages"`
}

// NewAuthenticatorValidationChallenge instantiates a new AuthenticatorValidationChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorValidationChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string, deviceChallenges []DeviceChallenge, configurationStages []SelectableStage) *AuthenticatorValidationChallenge {
	this := AuthenticatorValidationChallenge{}
	this.Type = type_
	var component string = "ak-stage-authenticator-validate"
	this.Component = &component
	this.PendingUser = pendingUser
	this.PendingUserAvatar = pendingUserAvatar
	this.DeviceChallenges = deviceChallenges
	this.ConfigurationStages = configurationStages
	return &this
}

// NewAuthenticatorValidationChallengeWithDefaults instantiates a new AuthenticatorValidationChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorValidationChallengeWithDefaults() *AuthenticatorValidationChallenge {
	this := AuthenticatorValidationChallenge{}
	var component string = "ak-stage-authenticator-validate"
	this.Component = &component
	return &this
}

// GetType returns the Type field value
func (o *AuthenticatorValidationChallenge) GetType() ChallengeChoices {
	if o == nil {
		var ret ChallengeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallenge) GetTypeOk() (*ChallengeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthenticatorValidationChallenge) SetType(v ChallengeChoices) {
	o.Type = v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *AuthenticatorValidationChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *AuthenticatorValidationChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *AuthenticatorValidationChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorValidationChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorValidationChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorValidationChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *AuthenticatorValidationChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *AuthenticatorValidationChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *AuthenticatorValidationChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetPendingUser returns the PendingUser field value
func (o *AuthenticatorValidationChallenge) GetPendingUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUser
}

// GetPendingUserOk returns a tuple with the PendingUser field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallenge) GetPendingUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUser, true
}

// SetPendingUser sets field value
func (o *AuthenticatorValidationChallenge) SetPendingUser(v string) {
	o.PendingUser = v
}

// GetPendingUserAvatar returns the PendingUserAvatar field value
func (o *AuthenticatorValidationChallenge) GetPendingUserAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUserAvatar
}

// GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallenge) GetPendingUserAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUserAvatar, true
}

// SetPendingUserAvatar sets field value
func (o *AuthenticatorValidationChallenge) SetPendingUserAvatar(v string) {
	o.PendingUserAvatar = v
}

// GetDeviceChallenges returns the DeviceChallenges field value
func (o *AuthenticatorValidationChallenge) GetDeviceChallenges() []DeviceChallenge {
	if o == nil {
		var ret []DeviceChallenge
		return ret
	}

	return o.DeviceChallenges
}

// GetDeviceChallengesOk returns a tuple with the DeviceChallenges field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallenge) GetDeviceChallengesOk() ([]DeviceChallenge, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceChallenges, true
}

// SetDeviceChallenges sets field value
func (o *AuthenticatorValidationChallenge) SetDeviceChallenges(v []DeviceChallenge) {
	o.DeviceChallenges = v
}

// GetConfigurationStages returns the ConfigurationStages field value
func (o *AuthenticatorValidationChallenge) GetConfigurationStages() []SelectableStage {
	if o == nil {
		var ret []SelectableStage
		return ret
	}

	return o.ConfigurationStages
}

// GetConfigurationStagesOk returns a tuple with the ConfigurationStages field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallenge) GetConfigurationStagesOk() ([]SelectableStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationStages, true
}

// SetConfigurationStages sets field value
func (o *AuthenticatorValidationChallenge) SetConfigurationStages(v []SelectableStage) {
	o.ConfigurationStages = v
}

func (o AuthenticatorValidationChallenge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
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
		toSerialize["device_challenges"] = o.DeviceChallenges
	}
	if true {
		toSerialize["configuration_stages"] = o.ConfigurationStages
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorValidationChallenge struct {
	value *AuthenticatorValidationChallenge
	isSet bool
}

func (v NullableAuthenticatorValidationChallenge) Get() *AuthenticatorValidationChallenge {
	return v.value
}

func (v *NullableAuthenticatorValidationChallenge) Set(val *AuthenticatorValidationChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorValidationChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorValidationChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorValidationChallenge(val *AuthenticatorValidationChallenge) *NullableAuthenticatorValidationChallenge {
	return &NullableAuthenticatorValidationChallenge{value: val, isSet: true}
}

func (v NullableAuthenticatorValidationChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorValidationChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
