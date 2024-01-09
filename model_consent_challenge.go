/*
authentik

Making authentication simple.

API version: 2023.10.6
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ConsentChallenge Challenge info for consent screens
type ConsentChallenge struct {
	Type                  ChallengeChoices          `json:"type"`
	FlowInfo              *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component             *string                   `json:"component,omitempty"`
	ResponseErrors        *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	PendingUser           string                    `json:"pending_user"`
	PendingUserAvatar     string                    `json:"pending_user_avatar"`
	HeaderText            *string                   `json:"header_text,omitempty"`
	Permissions           []ConsentPermission       `json:"permissions"`
	AdditionalPermissions []ConsentPermission       `json:"additional_permissions"`
	Token                 string                    `json:"token"`
}

// NewConsentChallenge instantiates a new ConsentChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string, permissions []ConsentPermission, additionalPermissions []ConsentPermission, token string) *ConsentChallenge {
	this := ConsentChallenge{}
	this.Type = type_
	var component string = "ak-stage-consent"
	this.Component = &component
	this.PendingUser = pendingUser
	this.PendingUserAvatar = pendingUserAvatar
	this.Permissions = permissions
	this.AdditionalPermissions = additionalPermissions
	this.Token = token
	return &this
}

// NewConsentChallengeWithDefaults instantiates a new ConsentChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentChallengeWithDefaults() *ConsentChallenge {
	this := ConsentChallenge{}
	var component string = "ak-stage-consent"
	this.Component = &component
	return &this
}

// GetType returns the Type field value
func (o *ConsentChallenge) GetType() ChallengeChoices {
	if o == nil {
		var ret ChallengeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConsentChallenge) GetTypeOk() (*ChallengeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConsentChallenge) SetType(v ChallengeChoices) {
	o.Type = v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *ConsentChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *ConsentChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *ConsentChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ConsentChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ConsentChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ConsentChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *ConsentChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *ConsentChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *ConsentChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetPendingUser returns the PendingUser field value
func (o *ConsentChallenge) GetPendingUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUser
}

// GetPendingUserOk returns a tuple with the PendingUser field value
// and a boolean to check if the value has been set.
func (o *ConsentChallenge) GetPendingUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUser, true
}

// SetPendingUser sets field value
func (o *ConsentChallenge) SetPendingUser(v string) {
	o.PendingUser = v
}

// GetPendingUserAvatar returns the PendingUserAvatar field value
func (o *ConsentChallenge) GetPendingUserAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUserAvatar
}

// GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field value
// and a boolean to check if the value has been set.
func (o *ConsentChallenge) GetPendingUserAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUserAvatar, true
}

// SetPendingUserAvatar sets field value
func (o *ConsentChallenge) SetPendingUserAvatar(v string) {
	o.PendingUserAvatar = v
}

// GetHeaderText returns the HeaderText field value if set, zero value otherwise.
func (o *ConsentChallenge) GetHeaderText() string {
	if o == nil || o.HeaderText == nil {
		var ret string
		return ret
	}
	return *o.HeaderText
}

// GetHeaderTextOk returns a tuple with the HeaderText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentChallenge) GetHeaderTextOk() (*string, bool) {
	if o == nil || o.HeaderText == nil {
		return nil, false
	}
	return o.HeaderText, true
}

// HasHeaderText returns a boolean if a field has been set.
func (o *ConsentChallenge) HasHeaderText() bool {
	if o != nil && o.HeaderText != nil {
		return true
	}

	return false
}

// SetHeaderText gets a reference to the given string and assigns it to the HeaderText field.
func (o *ConsentChallenge) SetHeaderText(v string) {
	o.HeaderText = &v
}

// GetPermissions returns the Permissions field value
func (o *ConsentChallenge) GetPermissions() []ConsentPermission {
	if o == nil {
		var ret []ConsentPermission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *ConsentChallenge) GetPermissionsOk() ([]ConsentPermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *ConsentChallenge) SetPermissions(v []ConsentPermission) {
	o.Permissions = v
}

// GetAdditionalPermissions returns the AdditionalPermissions field value
func (o *ConsentChallenge) GetAdditionalPermissions() []ConsentPermission {
	if o == nil {
		var ret []ConsentPermission
		return ret
	}

	return o.AdditionalPermissions
}

// GetAdditionalPermissionsOk returns a tuple with the AdditionalPermissions field value
// and a boolean to check if the value has been set.
func (o *ConsentChallenge) GetAdditionalPermissionsOk() ([]ConsentPermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalPermissions, true
}

// SetAdditionalPermissions sets field value
func (o *ConsentChallenge) SetAdditionalPermissions(v []ConsentPermission) {
	o.AdditionalPermissions = v
}

// GetToken returns the Token field value
func (o *ConsentChallenge) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ConsentChallenge) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ConsentChallenge) SetToken(v string) {
	o.Token = v
}

func (o ConsentChallenge) MarshalJSON() ([]byte, error) {
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
	if o.HeaderText != nil {
		toSerialize["header_text"] = o.HeaderText
	}
	if true {
		toSerialize["permissions"] = o.Permissions
	}
	if true {
		toSerialize["additional_permissions"] = o.AdditionalPermissions
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableConsentChallenge struct {
	value *ConsentChallenge
	isSet bool
}

func (v NullableConsentChallenge) Get() *ConsentChallenge {
	return v.value
}

func (v *NullableConsentChallenge) Set(val *ConsentChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentChallenge(val *ConsentChallenge) *NullableConsentChallenge {
	return &NullableConsentChallenge{value: val, isSet: true}
}

func (v NullableConsentChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
