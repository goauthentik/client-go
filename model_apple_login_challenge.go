/*
authentik

Making authentication simple.

API version: 2024.4.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AppleLoginChallenge Special challenge for apple-native authentication flow, which happens on the client.
type AppleLoginChallenge struct {
	Type           ChallengeChoices          `json:"type"`
	FlowInfo       *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component      *string                   `json:"component,omitempty"`
	ResponseErrors *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	ClientId       string                    `json:"client_id"`
	Scope          string                    `json:"scope"`
	RedirectUri    string                    `json:"redirect_uri"`
	State          string                    `json:"state"`
}

// NewAppleLoginChallenge instantiates a new AppleLoginChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppleLoginChallenge(type_ ChallengeChoices, clientId string, scope string, redirectUri string, state string) *AppleLoginChallenge {
	this := AppleLoginChallenge{}
	this.Type = type_
	var component string = "ak-source-oauth-apple"
	this.Component = &component
	this.ClientId = clientId
	this.Scope = scope
	this.RedirectUri = redirectUri
	this.State = state
	return &this
}

// NewAppleLoginChallengeWithDefaults instantiates a new AppleLoginChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppleLoginChallengeWithDefaults() *AppleLoginChallenge {
	this := AppleLoginChallenge{}
	var component string = "ak-source-oauth-apple"
	this.Component = &component
	return &this
}

// GetType returns the Type field value
func (o *AppleLoginChallenge) GetType() ChallengeChoices {
	if o == nil {
		var ret ChallengeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppleLoginChallenge) GetTypeOk() (*ChallengeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppleLoginChallenge) SetType(v ChallengeChoices) {
	o.Type = v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *AppleLoginChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleLoginChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *AppleLoginChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *AppleLoginChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AppleLoginChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleLoginChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AppleLoginChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AppleLoginChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *AppleLoginChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleLoginChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *AppleLoginChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *AppleLoginChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetClientId returns the ClientId field value
func (o *AppleLoginChallenge) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *AppleLoginChallenge) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *AppleLoginChallenge) SetClientId(v string) {
	o.ClientId = v
}

// GetScope returns the Scope field value
func (o *AppleLoginChallenge) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AppleLoginChallenge) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *AppleLoginChallenge) SetScope(v string) {
	o.Scope = v
}

// GetRedirectUri returns the RedirectUri field value
func (o *AppleLoginChallenge) GetRedirectUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value
// and a boolean to check if the value has been set.
func (o *AppleLoginChallenge) GetRedirectUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUri, true
}

// SetRedirectUri sets field value
func (o *AppleLoginChallenge) SetRedirectUri(v string) {
	o.RedirectUri = v
}

// GetState returns the State field value
func (o *AppleLoginChallenge) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *AppleLoginChallenge) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *AppleLoginChallenge) SetState(v string) {
	o.State = v
}

func (o AppleLoginChallenge) MarshalJSON() ([]byte, error) {
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
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableAppleLoginChallenge struct {
	value *AppleLoginChallenge
	isSet bool
}

func (v NullableAppleLoginChallenge) Get() *AppleLoginChallenge {
	return v.value
}

func (v *NullableAppleLoginChallenge) Set(val *AppleLoginChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableAppleLoginChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableAppleLoginChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppleLoginChallenge(val *AppleLoginChallenge) *NullableAppleLoginChallenge {
	return &NullableAppleLoginChallenge{value: val, isSet: true}
}

func (v NullableAppleLoginChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppleLoginChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
