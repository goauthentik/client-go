/*
authentik

Making authentication simple.

API version: 2024.6.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CaptchaChallenge Site public key
type CaptchaChallenge struct {
	Type              ChallengeChoices          `json:"type"`
	FlowInfo          *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component         *string                   `json:"component,omitempty"`
	ResponseErrors    *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	PendingUser       string                    `json:"pending_user"`
	PendingUserAvatar string                    `json:"pending_user_avatar"`
	SiteKey           string                    `json:"site_key"`
	JsUrl             string                    `json:"js_url"`
}

// NewCaptchaChallenge instantiates a new CaptchaChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptchaChallenge(type_ ChallengeChoices, pendingUser string, pendingUserAvatar string, siteKey string, jsUrl string) *CaptchaChallenge {
	this := CaptchaChallenge{}
	this.Type = type_
	var component string = "ak-stage-captcha"
	this.Component = &component
	this.PendingUser = pendingUser
	this.PendingUserAvatar = pendingUserAvatar
	this.SiteKey = siteKey
	this.JsUrl = jsUrl
	return &this
}

// NewCaptchaChallengeWithDefaults instantiates a new CaptchaChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptchaChallengeWithDefaults() *CaptchaChallenge {
	this := CaptchaChallenge{}
	var component string = "ak-stage-captcha"
	this.Component = &component
	return &this
}

// GetType returns the Type field value
func (o *CaptchaChallenge) GetType() ChallengeChoices {
	if o == nil {
		var ret ChallengeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaptchaChallenge) GetTypeOk() (*ChallengeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CaptchaChallenge) SetType(v ChallengeChoices) {
	o.Type = v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *CaptchaChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *CaptchaChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *CaptchaChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *CaptchaChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *CaptchaChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *CaptchaChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *CaptchaChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *CaptchaChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *CaptchaChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetPendingUser returns the PendingUser field value
func (o *CaptchaChallenge) GetPendingUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUser
}

// GetPendingUserOk returns a tuple with the PendingUser field value
// and a boolean to check if the value has been set.
func (o *CaptchaChallenge) GetPendingUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUser, true
}

// SetPendingUser sets field value
func (o *CaptchaChallenge) SetPendingUser(v string) {
	o.PendingUser = v
}

// GetPendingUserAvatar returns the PendingUserAvatar field value
func (o *CaptchaChallenge) GetPendingUserAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUserAvatar
}

// GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field value
// and a boolean to check if the value has been set.
func (o *CaptchaChallenge) GetPendingUserAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUserAvatar, true
}

// SetPendingUserAvatar sets field value
func (o *CaptchaChallenge) SetPendingUserAvatar(v string) {
	o.PendingUserAvatar = v
}

// GetSiteKey returns the SiteKey field value
func (o *CaptchaChallenge) GetSiteKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SiteKey
}

// GetSiteKeyOk returns a tuple with the SiteKey field value
// and a boolean to check if the value has been set.
func (o *CaptchaChallenge) GetSiteKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteKey, true
}

// SetSiteKey sets field value
func (o *CaptchaChallenge) SetSiteKey(v string) {
	o.SiteKey = v
}

// GetJsUrl returns the JsUrl field value
func (o *CaptchaChallenge) GetJsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsUrl
}

// GetJsUrlOk returns a tuple with the JsUrl field value
// and a boolean to check if the value has been set.
func (o *CaptchaChallenge) GetJsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsUrl, true
}

// SetJsUrl sets field value
func (o *CaptchaChallenge) SetJsUrl(v string) {
	o.JsUrl = v
}

func (o CaptchaChallenge) MarshalJSON() ([]byte, error) {
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
		toSerialize["site_key"] = o.SiteKey
	}
	if true {
		toSerialize["js_url"] = o.JsUrl
	}
	return json.Marshal(toSerialize)
}

type NullableCaptchaChallenge struct {
	value *CaptchaChallenge
	isSet bool
}

func (v NullableCaptchaChallenge) Get() *CaptchaChallenge {
	return v.value
}

func (v *NullableCaptchaChallenge) Set(val *CaptchaChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptchaChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptchaChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptchaChallenge(val *CaptchaChallenge) *NullableCaptchaChallenge {
	return &NullableCaptchaChallenge{value: val, isSet: true}
}

func (v NullableCaptchaChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptchaChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
