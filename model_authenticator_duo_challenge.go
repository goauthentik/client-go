/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorDuoChallenge Duo Challenge
type AuthenticatorDuoChallenge struct {
	FlowInfo          *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component         *string                   `json:"component,omitempty"`
	ResponseErrors    *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	PendingUser       string                    `json:"pending_user"`
	PendingUserAvatar string                    `json:"pending_user_avatar"`
	ActivationBarcode string                    `json:"activation_barcode"`
	ActivationCode    string                    `json:"activation_code"`
	StageUuid         string                    `json:"stage_uuid"`
}

// NewAuthenticatorDuoChallenge instantiates a new AuthenticatorDuoChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorDuoChallenge(pendingUser string, pendingUserAvatar string, activationBarcode string, activationCode string, stageUuid string) *AuthenticatorDuoChallenge {
	this := AuthenticatorDuoChallenge{}
	var component string = "ak-stage-authenticator-duo"
	this.Component = &component
	this.PendingUser = pendingUser
	this.PendingUserAvatar = pendingUserAvatar
	this.ActivationBarcode = activationBarcode
	this.ActivationCode = activationCode
	this.StageUuid = stageUuid
	return &this
}

// NewAuthenticatorDuoChallengeWithDefaults instantiates a new AuthenticatorDuoChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorDuoChallengeWithDefaults() *AuthenticatorDuoChallenge {
	this := AuthenticatorDuoChallenge{}
	var component string = "ak-stage-authenticator-duo"
	this.Component = &component
	return &this
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *AuthenticatorDuoChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *AuthenticatorDuoChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *AuthenticatorDuoChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorDuoChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorDuoChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorDuoChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *AuthenticatorDuoChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *AuthenticatorDuoChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *AuthenticatorDuoChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetPendingUser returns the PendingUser field value
func (o *AuthenticatorDuoChallenge) GetPendingUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUser
}

// GetPendingUserOk returns a tuple with the PendingUser field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoChallenge) GetPendingUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUser, true
}

// SetPendingUser sets field value
func (o *AuthenticatorDuoChallenge) SetPendingUser(v string) {
	o.PendingUser = v
}

// GetPendingUserAvatar returns the PendingUserAvatar field value
func (o *AuthenticatorDuoChallenge) GetPendingUserAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUserAvatar
}

// GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoChallenge) GetPendingUserAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUserAvatar, true
}

// SetPendingUserAvatar sets field value
func (o *AuthenticatorDuoChallenge) SetPendingUserAvatar(v string) {
	o.PendingUserAvatar = v
}

// GetActivationBarcode returns the ActivationBarcode field value
func (o *AuthenticatorDuoChallenge) GetActivationBarcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActivationBarcode
}

// GetActivationBarcodeOk returns a tuple with the ActivationBarcode field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoChallenge) GetActivationBarcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivationBarcode, true
}

// SetActivationBarcode sets field value
func (o *AuthenticatorDuoChallenge) SetActivationBarcode(v string) {
	o.ActivationBarcode = v
}

// GetActivationCode returns the ActivationCode field value
func (o *AuthenticatorDuoChallenge) GetActivationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActivationCode
}

// GetActivationCodeOk returns a tuple with the ActivationCode field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoChallenge) GetActivationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActivationCode, true
}

// SetActivationCode sets field value
func (o *AuthenticatorDuoChallenge) SetActivationCode(v string) {
	o.ActivationCode = v
}

// GetStageUuid returns the StageUuid field value
func (o *AuthenticatorDuoChallenge) GetStageUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StageUuid
}

// GetStageUuidOk returns a tuple with the StageUuid field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorDuoChallenge) GetStageUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StageUuid, true
}

// SetStageUuid sets field value
func (o *AuthenticatorDuoChallenge) SetStageUuid(v string) {
	o.StageUuid = v
}

func (o AuthenticatorDuoChallenge) MarshalJSON() ([]byte, error) {
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
		toSerialize["activation_barcode"] = o.ActivationBarcode
	}
	if true {
		toSerialize["activation_code"] = o.ActivationCode
	}
	if true {
		toSerialize["stage_uuid"] = o.StageUuid
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorDuoChallenge struct {
	value *AuthenticatorDuoChallenge
	isSet bool
}

func (v NullableAuthenticatorDuoChallenge) Get() *AuthenticatorDuoChallenge {
	return v.value
}

func (v *NullableAuthenticatorDuoChallenge) Set(val *AuthenticatorDuoChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorDuoChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorDuoChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorDuoChallenge(val *AuthenticatorDuoChallenge) *NullableAuthenticatorDuoChallenge {
	return &NullableAuthenticatorDuoChallenge{value: val, isSet: true}
}

func (v NullableAuthenticatorDuoChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorDuoChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
