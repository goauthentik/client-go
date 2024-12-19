/*
authentik

Making authentication simple.

API version: 2024.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SessionEndChallenge Challenge for ending a session
type SessionEndChallenge struct {
	FlowInfo             *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component            *string                   `json:"component,omitempty"`
	ResponseErrors       *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	PendingUser          string                    `json:"pending_user"`
	PendingUserAvatar    string                    `json:"pending_user_avatar"`
	ApplicationName      *string                   `json:"application_name,omitempty"`
	ApplicationLaunchUrl *string                   `json:"application_launch_url,omitempty"`
	InvalidationFlowUrl  *string                   `json:"invalidation_flow_url,omitempty"`
	BrandName            string                    `json:"brand_name"`
}

// NewSessionEndChallenge instantiates a new SessionEndChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionEndChallenge(pendingUser string, pendingUserAvatar string, brandName string) *SessionEndChallenge {
	this := SessionEndChallenge{}
	var component string = "ak-stage-session-end"
	this.Component = &component
	this.PendingUser = pendingUser
	this.PendingUserAvatar = pendingUserAvatar
	this.BrandName = brandName
	return &this
}

// NewSessionEndChallengeWithDefaults instantiates a new SessionEndChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionEndChallengeWithDefaults() *SessionEndChallenge {
	this := SessionEndChallenge{}
	var component string = "ak-stage-session-end"
	this.Component = &component
	return &this
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *SessionEndChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionEndChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *SessionEndChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *SessionEndChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *SessionEndChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionEndChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *SessionEndChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *SessionEndChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *SessionEndChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionEndChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *SessionEndChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *SessionEndChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetPendingUser returns the PendingUser field value
func (o *SessionEndChallenge) GetPendingUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUser
}

// GetPendingUserOk returns a tuple with the PendingUser field value
// and a boolean to check if the value has been set.
func (o *SessionEndChallenge) GetPendingUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUser, true
}

// SetPendingUser sets field value
func (o *SessionEndChallenge) SetPendingUser(v string) {
	o.PendingUser = v
}

// GetPendingUserAvatar returns the PendingUserAvatar field value
func (o *SessionEndChallenge) GetPendingUserAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUserAvatar
}

// GetPendingUserAvatarOk returns a tuple with the PendingUserAvatar field value
// and a boolean to check if the value has been set.
func (o *SessionEndChallenge) GetPendingUserAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingUserAvatar, true
}

// SetPendingUserAvatar sets field value
func (o *SessionEndChallenge) SetPendingUserAvatar(v string) {
	o.PendingUserAvatar = v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *SessionEndChallenge) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionEndChallenge) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *SessionEndChallenge) HasApplicationName() bool {
	if o != nil && o.ApplicationName != nil {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *SessionEndChallenge) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetApplicationLaunchUrl returns the ApplicationLaunchUrl field value if set, zero value otherwise.
func (o *SessionEndChallenge) GetApplicationLaunchUrl() string {
	if o == nil || o.ApplicationLaunchUrl == nil {
		var ret string
		return ret
	}
	return *o.ApplicationLaunchUrl
}

// GetApplicationLaunchUrlOk returns a tuple with the ApplicationLaunchUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionEndChallenge) GetApplicationLaunchUrlOk() (*string, bool) {
	if o == nil || o.ApplicationLaunchUrl == nil {
		return nil, false
	}
	return o.ApplicationLaunchUrl, true
}

// HasApplicationLaunchUrl returns a boolean if a field has been set.
func (o *SessionEndChallenge) HasApplicationLaunchUrl() bool {
	if o != nil && o.ApplicationLaunchUrl != nil {
		return true
	}

	return false
}

// SetApplicationLaunchUrl gets a reference to the given string and assigns it to the ApplicationLaunchUrl field.
func (o *SessionEndChallenge) SetApplicationLaunchUrl(v string) {
	o.ApplicationLaunchUrl = &v
}

// GetInvalidationFlowUrl returns the InvalidationFlowUrl field value if set, zero value otherwise.
func (o *SessionEndChallenge) GetInvalidationFlowUrl() string {
	if o == nil || o.InvalidationFlowUrl == nil {
		var ret string
		return ret
	}
	return *o.InvalidationFlowUrl
}

// GetInvalidationFlowUrlOk returns a tuple with the InvalidationFlowUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionEndChallenge) GetInvalidationFlowUrlOk() (*string, bool) {
	if o == nil || o.InvalidationFlowUrl == nil {
		return nil, false
	}
	return o.InvalidationFlowUrl, true
}

// HasInvalidationFlowUrl returns a boolean if a field has been set.
func (o *SessionEndChallenge) HasInvalidationFlowUrl() bool {
	if o != nil && o.InvalidationFlowUrl != nil {
		return true
	}

	return false
}

// SetInvalidationFlowUrl gets a reference to the given string and assigns it to the InvalidationFlowUrl field.
func (o *SessionEndChallenge) SetInvalidationFlowUrl(v string) {
	o.InvalidationFlowUrl = &v
}

// GetBrandName returns the BrandName field value
func (o *SessionEndChallenge) GetBrandName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value
// and a boolean to check if the value has been set.
func (o *SessionEndChallenge) GetBrandNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandName, true
}

// SetBrandName sets field value
func (o *SessionEndChallenge) SetBrandName(v string) {
	o.BrandName = v
}

func (o SessionEndChallenge) MarshalJSON() ([]byte, error) {
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
	if o.ApplicationName != nil {
		toSerialize["application_name"] = o.ApplicationName
	}
	if o.ApplicationLaunchUrl != nil {
		toSerialize["application_launch_url"] = o.ApplicationLaunchUrl
	}
	if o.InvalidationFlowUrl != nil {
		toSerialize["invalidation_flow_url"] = o.InvalidationFlowUrl
	}
	if true {
		toSerialize["brand_name"] = o.BrandName
	}
	return json.Marshal(toSerialize)
}

type NullableSessionEndChallenge struct {
	value *SessionEndChallenge
	isSet bool
}

func (v NullableSessionEndChallenge) Get() *SessionEndChallenge {
	return v.value
}

func (v *NullableSessionEndChallenge) Set(val *SessionEndChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionEndChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionEndChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionEndChallenge(val *SessionEndChallenge) *NullableSessionEndChallenge {
	return &NullableSessionEndChallenge{value: val, isSet: true}
}

func (v NullableSessionEndChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionEndChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
