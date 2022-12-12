/*
authentik

Making authentication simple.

API version: 2022.11.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorValidationChallengeResponseRequest Challenge used for Code-based and WebAuthn authenticators
type AuthenticatorValidationChallengeResponseRequest struct {
	Component         *string                 `json:"component,omitempty"`
	SelectedChallenge *DeviceChallengeRequest `json:"selected_challenge,omitempty"`
	SelectedStage     *string                 `json:"selected_stage,omitempty"`
	Code              *string                 `json:"code,omitempty"`
	Webauthn          map[string]interface{}  `json:"webauthn,omitempty"`
	Duo               *int32                  `json:"duo,omitempty"`
}

// NewAuthenticatorValidationChallengeResponseRequest instantiates a new AuthenticatorValidationChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorValidationChallengeResponseRequest() *AuthenticatorValidationChallengeResponseRequest {
	this := AuthenticatorValidationChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-validate"
	this.Component = &component
	return &this
}

// NewAuthenticatorValidationChallengeResponseRequestWithDefaults instantiates a new AuthenticatorValidationChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorValidationChallengeResponseRequestWithDefaults() *AuthenticatorValidationChallengeResponseRequest {
	this := AuthenticatorValidationChallengeResponseRequest{}
	var component string = "ak-stage-authenticator-validate"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AuthenticatorValidationChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AuthenticatorValidationChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

// GetSelectedChallenge returns the SelectedChallenge field value if set, zero value otherwise.
func (o *AuthenticatorValidationChallengeResponseRequest) GetSelectedChallenge() DeviceChallengeRequest {
	if o == nil || o.SelectedChallenge == nil {
		var ret DeviceChallengeRequest
		return ret
	}
	return *o.SelectedChallenge
}

// GetSelectedChallengeOk returns a tuple with the SelectedChallenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) GetSelectedChallengeOk() (*DeviceChallengeRequest, bool) {
	if o == nil || o.SelectedChallenge == nil {
		return nil, false
	}
	return o.SelectedChallenge, true
}

// HasSelectedChallenge returns a boolean if a field has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) HasSelectedChallenge() bool {
	if o != nil && o.SelectedChallenge != nil {
		return true
	}

	return false
}

// SetSelectedChallenge gets a reference to the given DeviceChallengeRequest and assigns it to the SelectedChallenge field.
func (o *AuthenticatorValidationChallengeResponseRequest) SetSelectedChallenge(v DeviceChallengeRequest) {
	o.SelectedChallenge = &v
}

// GetSelectedStage returns the SelectedStage field value if set, zero value otherwise.
func (o *AuthenticatorValidationChallengeResponseRequest) GetSelectedStage() string {
	if o == nil || o.SelectedStage == nil {
		var ret string
		return ret
	}
	return *o.SelectedStage
}

// GetSelectedStageOk returns a tuple with the SelectedStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) GetSelectedStageOk() (*string, bool) {
	if o == nil || o.SelectedStage == nil {
		return nil, false
	}
	return o.SelectedStage, true
}

// HasSelectedStage returns a boolean if a field has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) HasSelectedStage() bool {
	if o != nil && o.SelectedStage != nil {
		return true
	}

	return false
}

// SetSelectedStage gets a reference to the given string and assigns it to the SelectedStage field.
func (o *AuthenticatorValidationChallengeResponseRequest) SetSelectedStage(v string) {
	o.SelectedStage = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AuthenticatorValidationChallengeResponseRequest) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AuthenticatorValidationChallengeResponseRequest) SetCode(v string) {
	o.Code = &v
}

// GetWebauthn returns the Webauthn field value if set, zero value otherwise.
func (o *AuthenticatorValidationChallengeResponseRequest) GetWebauthn() map[string]interface{} {
	if o == nil || o.Webauthn == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Webauthn
}

// GetWebauthnOk returns a tuple with the Webauthn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) GetWebauthnOk() (map[string]interface{}, bool) {
	if o == nil || o.Webauthn == nil {
		return nil, false
	}
	return o.Webauthn, true
}

// HasWebauthn returns a boolean if a field has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) HasWebauthn() bool {
	if o != nil && o.Webauthn != nil {
		return true
	}

	return false
}

// SetWebauthn gets a reference to the given map[string]interface{} and assigns it to the Webauthn field.
func (o *AuthenticatorValidationChallengeResponseRequest) SetWebauthn(v map[string]interface{}) {
	o.Webauthn = v
}

// GetDuo returns the Duo field value if set, zero value otherwise.
func (o *AuthenticatorValidationChallengeResponseRequest) GetDuo() int32 {
	if o == nil || o.Duo == nil {
		var ret int32
		return ret
	}
	return *o.Duo
}

// GetDuoOk returns a tuple with the Duo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) GetDuoOk() (*int32, bool) {
	if o == nil || o.Duo == nil {
		return nil, false
	}
	return o.Duo, true
}

// HasDuo returns a boolean if a field has been set.
func (o *AuthenticatorValidationChallengeResponseRequest) HasDuo() bool {
	if o != nil && o.Duo != nil {
		return true
	}

	return false
}

// SetDuo gets a reference to the given int32 and assigns it to the Duo field.
func (o *AuthenticatorValidationChallengeResponseRequest) SetDuo(v int32) {
	o.Duo = &v
}

func (o AuthenticatorValidationChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.SelectedChallenge != nil {
		toSerialize["selected_challenge"] = o.SelectedChallenge
	}
	if o.SelectedStage != nil {
		toSerialize["selected_stage"] = o.SelectedStage
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Webauthn != nil {
		toSerialize["webauthn"] = o.Webauthn
	}
	if o.Duo != nil {
		toSerialize["duo"] = o.Duo
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorValidationChallengeResponseRequest struct {
	value *AuthenticatorValidationChallengeResponseRequest
	isSet bool
}

func (v NullableAuthenticatorValidationChallengeResponseRequest) Get() *AuthenticatorValidationChallengeResponseRequest {
	return v.value
}

func (v *NullableAuthenticatorValidationChallengeResponseRequest) Set(val *AuthenticatorValidationChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorValidationChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorValidationChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorValidationChallengeResponseRequest(val *AuthenticatorValidationChallengeResponseRequest) *NullableAuthenticatorValidationChallengeResponseRequest {
	return &NullableAuthenticatorValidationChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorValidationChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorValidationChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
