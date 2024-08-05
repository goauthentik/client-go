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

// OAuthDeviceCodeChallenge OAuth Device code challenge
type OAuthDeviceCodeChallenge struct {
	FlowInfo       *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component      *string                   `json:"component,omitempty"`
	ResponseErrors *map[string][]ErrorDetail `json:"response_errors,omitempty"`
}

// NewOAuthDeviceCodeChallenge instantiates a new OAuthDeviceCodeChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthDeviceCodeChallenge() *OAuthDeviceCodeChallenge {
	this := OAuthDeviceCodeChallenge{}
	var component string = "ak-provider-oauth2-device-code"
	this.Component = &component
	return &this
}

// NewOAuthDeviceCodeChallengeWithDefaults instantiates a new OAuthDeviceCodeChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthDeviceCodeChallengeWithDefaults() *OAuthDeviceCodeChallenge {
	this := OAuthDeviceCodeChallenge{}
	var component string = "ak-provider-oauth2-device-code"
	this.Component = &component
	return &this
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *OAuthDeviceCodeChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthDeviceCodeChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *OAuthDeviceCodeChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *OAuthDeviceCodeChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *OAuthDeviceCodeChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthDeviceCodeChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *OAuthDeviceCodeChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *OAuthDeviceCodeChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *OAuthDeviceCodeChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthDeviceCodeChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *OAuthDeviceCodeChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *OAuthDeviceCodeChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

func (o OAuthDeviceCodeChallenge) MarshalJSON() ([]byte, error) {
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

type NullableOAuthDeviceCodeChallenge struct {
	value *OAuthDeviceCodeChallenge
	isSet bool
}

func (v NullableOAuthDeviceCodeChallenge) Get() *OAuthDeviceCodeChallenge {
	return v.value
}

func (v *NullableOAuthDeviceCodeChallenge) Set(val *OAuthDeviceCodeChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthDeviceCodeChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthDeviceCodeChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthDeviceCodeChallenge(val *OAuthDeviceCodeChallenge) *NullableOAuthDeviceCodeChallenge {
	return &NullableOAuthDeviceCodeChallenge{value: val, isSet: true}
}

func (v NullableOAuthDeviceCodeChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthDeviceCodeChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
