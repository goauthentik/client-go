/*
authentik

Making authentication simple.

API version: 2023.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PasswordChallenge Password challenge UI fields
type PasswordChallenge struct {
	Type              ChallengeChoices          `json:"type"`
	FlowInfo          *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component         *string                   `json:"component,omitempty"`
	ResponseErrors    *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	PendingUser       string                    `json:"pending_user"`
	PendingUserAvatar string                    `json:"pending_user_avatar"`
	RecoveryUrl       *string                   `json:"recovery_url,omitempty"`
}

// NewPasswordChallenge instantiates a new PasswordChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string) *PasswordChallenge {
	this := PasswordChallenge{}
	this.Type = type_
	var component string = "ak-stage-password"
	this.Component = &component
	this.PendingUser = pendingUser
	this.PendingUserAvatar = pendingUserAvatar
	return &this
}

// NewPasswordChallengeWithDefaults instantiates a new PasswordChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordChallengeWithDefaults() *PasswordChallenge {
	this := PasswordChallenge{}
	var component string = "ak-stage-password"
	this.Component = &component
	return &this
}

// GetType returns the Type field value
func (o *PasswordChallenge) GetType() ChallengeChoices {
	if o == nil {
		var ret ChallengeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PasswordChallenge) GetTypeOk() (*ChallengeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PasswordChallenge) SetType(v ChallengeChoices) {
	o.Type = v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *PasswordChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *PasswordChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *PasswordChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *PasswordChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *PasswordChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *PasswordChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *PasswordChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *PasswordChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *PasswordChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetPendingUser returns the PendingUser field value
func (o *PasswordChallenge) GetPendingUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUser
}

// GetPendingUserOk returns a tuple with the PendingUser field value
// and a boolean to check if the value has been set.
func (o *PasswordChallenge) GetPendingUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUser, true
}

// SetPendingUser sets field value
func (o *PasswordChallenge) SetPendingUser(v string) {
	o.PendingUser = v
}

// GetPendingUserAvatar returns the PendingUserAvatar field value
func (o *PasswordChallenge) GetPendingUserAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUserAvatar
}

// GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field value
// and a boolean to check if the value has been set.
func (o *PasswordChallenge) GetPendingUserAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUserAvatar, true
}

// SetPendingUserAvatar sets field value
func (o *PasswordChallenge) SetPendingUserAvatar(v string) {
	o.PendingUserAvatar = v
}

// GetRecoveryUrl returns the RecoveryUrl field value if set, zero value otherwise.
func (o *PasswordChallenge) GetRecoveryUrl() string {
	if o == nil || o.RecoveryUrl == nil {
		var ret string
		return ret
	}
	return *o.RecoveryUrl
}

// GetRecoveryUrlOk returns a tuple with the RecoveryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChallenge) GetRecoveryUrlOk() (*string, bool) {
	if o == nil || o.RecoveryUrl == nil {
		return nil, false
	}
	return o.RecoveryUrl, true
}

// HasRecoveryUrl returns a boolean if a field has been set.
func (o *PasswordChallenge) HasRecoveryUrl() bool {
	if o != nil && o.RecoveryUrl != nil {
		return true
	}

	return false
}

// SetRecoveryUrl gets a reference to the given string and assigns it to the RecoveryUrl field.
func (o *PasswordChallenge) SetRecoveryUrl(v string) {
	o.RecoveryUrl = &v
}

func (o PasswordChallenge) MarshalJSON() ([]byte, error) {
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
	if o.RecoveryUrl != nil {
		toSerialize["recovery_url"] = o.RecoveryUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordChallenge struct {
	value *PasswordChallenge
	isSet bool
}

func (v NullablePasswordChallenge) Get() *PasswordChallenge {
	return v.value
}

func (v *NullablePasswordChallenge) Set(val *PasswordChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordChallenge(val *PasswordChallenge) *NullablePasswordChallenge {
	return &NullablePasswordChallenge{value: val, isSet: true}
}

func (v NullablePasswordChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
