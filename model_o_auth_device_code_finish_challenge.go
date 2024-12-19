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

// OAuthDeviceCodeFinishChallenge Final challenge after user enters their code
type OAuthDeviceCodeFinishChallenge struct {
	FlowInfo       *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component      *string                   `json:"component,omitempty"`
	ResponseErrors *map[string][]ErrorDetail `json:"response_errors,omitempty"`
}

// NewOAuthDeviceCodeFinishChallenge instantiates a new OAuthDeviceCodeFinishChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthDeviceCodeFinishChallenge() *OAuthDeviceCodeFinishChallenge {
	this := OAuthDeviceCodeFinishChallenge{}
	var component string = "ak-provider-oauth2-device-code-finish"
	this.Component = &component
	return &this
}

// NewOAuthDeviceCodeFinishChallengeWithDefaults instantiates a new OAuthDeviceCodeFinishChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthDeviceCodeFinishChallengeWithDefaults() *OAuthDeviceCodeFinishChallenge {
	this := OAuthDeviceCodeFinishChallenge{}
	var component string = "ak-provider-oauth2-device-code-finish"
	this.Component = &component
	return &this
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *OAuthDeviceCodeFinishChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthDeviceCodeFinishChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *OAuthDeviceCodeFinishChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *OAuthDeviceCodeFinishChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *OAuthDeviceCodeFinishChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthDeviceCodeFinishChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *OAuthDeviceCodeFinishChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *OAuthDeviceCodeFinishChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *OAuthDeviceCodeFinishChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthDeviceCodeFinishChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *OAuthDeviceCodeFinishChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *OAuthDeviceCodeFinishChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

func (o OAuthDeviceCodeFinishChallenge) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableOAuthDeviceCodeFinishChallenge struct {
	value *OAuthDeviceCodeFinishChallenge
	isSet bool
}

func (v NullableOAuthDeviceCodeFinishChallenge) Get() *OAuthDeviceCodeFinishChallenge {
	return v.value
}

func (v *NullableOAuthDeviceCodeFinishChallenge) Set(val *OAuthDeviceCodeFinishChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthDeviceCodeFinishChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthDeviceCodeFinishChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthDeviceCodeFinishChallenge(val *OAuthDeviceCodeFinishChallenge) *NullableOAuthDeviceCodeFinishChallenge {
	return &NullableOAuthDeviceCodeFinishChallenge{value: val, isSet: true}
}

func (v NullableOAuthDeviceCodeFinishChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthDeviceCodeFinishChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
