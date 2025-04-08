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

// IdentificationChallengeCaptchaStage struct for IdentificationChallengeCaptchaStage
type IdentificationChallengeCaptchaStage struct {
	FlowInfo          *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component         *string                   `json:"component,omitempty"`
	ResponseErrors    *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	PendingUser       string                    `json:"pending_user"`
	PendingUserAvatar string                    `json:"pending_user_avatar"`
	SiteKey           string                    `json:"site_key"`
	JsUrl             string                    `json:"js_url"`
	Interactive       bool                      `json:"interactive"`
}

// NewIdentificationChallengeCaptchaStage instantiates a new IdentificationChallengeCaptchaStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentificationChallengeCaptchaStage(pendingUser string, pendingUserAvatar string, siteKey string, jsUrl string, interactive bool) *IdentificationChallengeCaptchaStage {
	this := IdentificationChallengeCaptchaStage{}
	var component string = "ak-stage-captcha"
	this.Component = &component
	this.PendingUser = pendingUser
	this.PendingUserAvatar = pendingUserAvatar
	this.SiteKey = siteKey
	this.JsUrl = jsUrl
	this.Interactive = interactive
	return &this
}

// NewIdentificationChallengeCaptchaStageWithDefaults instantiates a new IdentificationChallengeCaptchaStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentificationChallengeCaptchaStageWithDefaults() *IdentificationChallengeCaptchaStage {
	this := IdentificationChallengeCaptchaStage{}
	var component string = "ak-stage-captcha"
	this.Component = &component
	return &this
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *IdentificationChallengeCaptchaStage) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallengeCaptchaStage) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *IdentificationChallengeCaptchaStage) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *IdentificationChallengeCaptchaStage) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *IdentificationChallengeCaptchaStage) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallengeCaptchaStage) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *IdentificationChallengeCaptchaStage) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *IdentificationChallengeCaptchaStage) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *IdentificationChallengeCaptchaStage) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallengeCaptchaStage) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *IdentificationChallengeCaptchaStage) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *IdentificationChallengeCaptchaStage) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetPendingUser returns the PendingUser field value
func (o *IdentificationChallengeCaptchaStage) GetPendingUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUser
}

// GetPendingUserOk returns a tuple with the PendingUser field value
// and a boolean to check if the value has been set.
func (o *IdentificationChallengeCaptchaStage) GetPendingUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUser, true
}

// SetPendingUser sets field value
func (o *IdentificationChallengeCaptchaStage) SetPendingUser(v string) {
	o.PendingUser = v
}

// GetPendingUserAvatar returns the PendingUserAvatar field value
func (o *IdentificationChallengeCaptchaStage) GetPendingUserAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUserAvatar
}

// GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field value
// and a boolean to check if the value has been set.
func (o *IdentificationChallengeCaptchaStage) GetPendingUserAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUserAvatar, true
}

// SetPendingUserAvatar sets field value
func (o *IdentificationChallengeCaptchaStage) SetPendingUserAvatar(v string) {
	o.PendingUserAvatar = v
}

// GetSiteKey returns the SiteKey field value
func (o *IdentificationChallengeCaptchaStage) GetSiteKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteKey
}

// GetSiteKeyOk returns a tuple with the SiteKey field value
// and a boolean to check if the value has been set.
func (o *IdentificationChallengeCaptchaStage) GetSiteKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteKey, true
}

// SetSiteKey sets field value
func (o *IdentificationChallengeCaptchaStage) SetSiteKey(v string) {
	o.SiteKey = v
}

// GetJsUrl returns the JsUrl field value
func (o *IdentificationChallengeCaptchaStage) GetJsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsUrl
}

// GetJsUrlOk returns a tuple with the JsUrl field value
// and a boolean to check if the value has been set.
func (o *IdentificationChallengeCaptchaStage) GetJsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsUrl, true
}

// SetJsUrl sets field value
func (o *IdentificationChallengeCaptchaStage) SetJsUrl(v string) {
	o.JsUrl = v
}

// GetInteractive returns the Interactive field value
func (o *IdentificationChallengeCaptchaStage) GetInteractive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Interactive
}

// GetInteractiveOk returns a tuple with the Interactive field value
// and a boolean to check if the value has been set.
func (o *IdentificationChallengeCaptchaStage) GetInteractiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactive, true
}

// SetInteractive sets field value
func (o *IdentificationChallengeCaptchaStage) SetInteractive(v bool) {
	o.Interactive = v
}

func (o IdentificationChallengeCaptchaStage) MarshalJSON() ([]byte, error) {
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
		toSerialize["site_key"] = o.SiteKey
	}
	if true {
		toSerialize["js_url"] = o.JsUrl
	}
	if true {
		toSerialize["interactive"] = o.Interactive
	}
	return json.Marshal(toSerialize)
}

type NullableIdentificationChallengeCaptchaStage struct {
	value *IdentificationChallengeCaptchaStage
	isSet bool
}

func (v NullableIdentificationChallengeCaptchaStage) Get() *IdentificationChallengeCaptchaStage {
	return v.value
}

func (v *NullableIdentificationChallengeCaptchaStage) Set(val *IdentificationChallengeCaptchaStage) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentificationChallengeCaptchaStage) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentificationChallengeCaptchaStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentificationChallengeCaptchaStage(val *IdentificationChallengeCaptchaStage) *NullableIdentificationChallengeCaptchaStage {
	return &NullableIdentificationChallengeCaptchaStage{value: val, isSet: true}
}

func (v NullableIdentificationChallengeCaptchaStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentificationChallengeCaptchaStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
