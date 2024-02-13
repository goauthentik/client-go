/*
authentik

Making authentication simple.

API version: 2023.10.7
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// IdentificationChallenge Identification challenges with all UI elements
type IdentificationChallenge struct {
	Type             ChallengeChoices          `json:"type"`
	FlowInfo         *ContextualFlowInfo       `json:"flow_info,omitempty"`
	Component        *string                   `json:"component,omitempty"`
	ResponseErrors   *map[string][]ErrorDetail `json:"response_errors,omitempty"`
	UserFields       []string                  `json:"user_fields"`
	PasswordFields   bool                      `json:"password_fields"`
	ApplicationPre   *string                   `json:"application_pre,omitempty"`
	EnrollUrl        *string                   `json:"enroll_url,omitempty"`
	RecoveryUrl      *string                   `json:"recovery_url,omitempty"`
	PasswordlessUrl  *string                   `json:"passwordless_url,omitempty"`
	PrimaryAction    string                    `json:"primary_action"`
	Sources          []LoginSource             `json:"sources,omitempty"`
	ShowSourceLabels bool                      `json:"show_source_labels"`
}

// NewIdentificationChallenge instantiates a new IdentificationChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentificationChallenge(type_ ChallengeChoices, userFields []string, passwordFields bool, primaryAction string, showSourceLabels bool) *IdentificationChallenge {
	this := IdentificationChallenge{}
	this.Type = type_
	var component string = "ak-stage-identification"
	this.Component = &component
	this.UserFields = userFields
	this.PasswordFields = passwordFields
	this.PrimaryAction = primaryAction
	this.ShowSourceLabels = showSourceLabels
	return &this
}

// NewIdentificationChallengeWithDefaults instantiates a new IdentificationChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentificationChallengeWithDefaults() *IdentificationChallenge {
	this := IdentificationChallenge{}
	var component string = "ak-stage-identification"
	this.Component = &component
	return &this
}

// GetType returns the Type field value
func (o *IdentificationChallenge) GetType() ChallengeChoices {
	if o == nil {
		var ret ChallengeChoices
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetTypeOk() (*ChallengeChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IdentificationChallenge) SetType(v ChallengeChoices) {
	o.Type = v
}

// GetFlowInfo returns the FlowInfo field value if set, zero value otherwise.
func (o *IdentificationChallenge) GetFlowInfo() ContextualFlowInfo {
	if o == nil || o.FlowInfo == nil {
		var ret ContextualFlowInfo
		return ret
	}
	return *o.FlowInfo
}

// GetFlowInfoOk returns a tuple with the FlowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetFlowInfoOk() (*ContextualFlowInfo, bool) {
	if o == nil || o.FlowInfo == nil {
		return nil, false
	}
	return o.FlowInfo, true
}

// HasFlowInfo returns a boolean if a field has been set.
func (o *IdentificationChallenge) HasFlowInfo() bool {
	if o != nil && o.FlowInfo != nil {
		return true
	}

	return false
}

// SetFlowInfo gets a reference to the given ContextualFlowInfo and assigns it to the FlowInfo field.
func (o *IdentificationChallenge) SetFlowInfo(v ContextualFlowInfo) {
	o.FlowInfo = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *IdentificationChallenge) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *IdentificationChallenge) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *IdentificationChallenge) SetComponent(v string) {
	o.Component = &v
}

// GetResponseErrors returns the ResponseErrors field value if set, zero value otherwise.
func (o *IdentificationChallenge) GetResponseErrors() map[string][]ErrorDetail {
	if o == nil || o.ResponseErrors == nil {
		var ret map[string][]ErrorDetail
		return ret
	}
	return *o.ResponseErrors
}

// GetResponseErrorsOk returns a tuple with the ResponseErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetResponseErrorsOk() (*map[string][]ErrorDetail, bool) {
	if o == nil || o.ResponseErrors == nil {
		return nil, false
	}
	return o.ResponseErrors, true
}

// HasResponseErrors returns a boolean if a field has been set.
func (o *IdentificationChallenge) HasResponseErrors() bool {
	if o != nil && o.ResponseErrors != nil {
		return true
	}

	return false
}

// SetResponseErrors gets a reference to the given map[string][]ErrorDetail and assigns it to the ResponseErrors field.
func (o *IdentificationChallenge) SetResponseErrors(v map[string][]ErrorDetail) {
	o.ResponseErrors = &v
}

// GetUserFields returns the UserFields field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *IdentificationChallenge) GetUserFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UserFields
}

// GetUserFieldsOk returns a tuple with the UserFields field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentificationChallenge) GetUserFieldsOk() ([]string, bool) {
	if o == nil || o.UserFields == nil {
		return nil, false
	}
	return o.UserFields, true
}

// SetUserFields sets field value
func (o *IdentificationChallenge) SetUserFields(v []string) {
	o.UserFields = v
}

// GetPasswordFields returns the PasswordFields field value
func (o *IdentificationChallenge) GetPasswordFields() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PasswordFields
}

// GetPasswordFieldsOk returns a tuple with the PasswordFields field value
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetPasswordFieldsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordFields, true
}

// SetPasswordFields sets field value
func (o *IdentificationChallenge) SetPasswordFields(v bool) {
	o.PasswordFields = v
}

// GetApplicationPre returns the ApplicationPre field value if set, zero value otherwise.
func (o *IdentificationChallenge) GetApplicationPre() string {
	if o == nil || o.ApplicationPre == nil {
		var ret string
		return ret
	}
	return *o.ApplicationPre
}

// GetApplicationPreOk returns a tuple with the ApplicationPre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetApplicationPreOk() (*string, bool) {
	if o == nil || o.ApplicationPre == nil {
		return nil, false
	}
	return o.ApplicationPre, true
}

// HasApplicationPre returns a boolean if a field has been set.
func (o *IdentificationChallenge) HasApplicationPre() bool {
	if o != nil && o.ApplicationPre != nil {
		return true
	}

	return false
}

// SetApplicationPre gets a reference to the given string and assigns it to the ApplicationPre field.
func (o *IdentificationChallenge) SetApplicationPre(v string) {
	o.ApplicationPre = &v
}

// GetEnrollUrl returns the EnrollUrl field value if set, zero value otherwise.
func (o *IdentificationChallenge) GetEnrollUrl() string {
	if o == nil || o.EnrollUrl == nil {
		var ret string
		return ret
	}
	return *o.EnrollUrl
}

// GetEnrollUrlOk returns a tuple with the EnrollUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetEnrollUrlOk() (*string, bool) {
	if o == nil || o.EnrollUrl == nil {
		return nil, false
	}
	return o.EnrollUrl, true
}

// HasEnrollUrl returns a boolean if a field has been set.
func (o *IdentificationChallenge) HasEnrollUrl() bool {
	if o != nil && o.EnrollUrl != nil {
		return true
	}

	return false
}

// SetEnrollUrl gets a reference to the given string and assigns it to the EnrollUrl field.
func (o *IdentificationChallenge) SetEnrollUrl(v string) {
	o.EnrollUrl = &v
}

// GetRecoveryUrl returns the RecoveryUrl field value if set, zero value otherwise.
func (o *IdentificationChallenge) GetRecoveryUrl() string {
	if o == nil || o.RecoveryUrl == nil {
		var ret string
		return ret
	}
	return *o.RecoveryUrl
}

// GetRecoveryUrlOk returns a tuple with the RecoveryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetRecoveryUrlOk() (*string, bool) {
	if o == nil || o.RecoveryUrl == nil {
		return nil, false
	}
	return o.RecoveryUrl, true
}

// HasRecoveryUrl returns a boolean if a field has been set.
func (o *IdentificationChallenge) HasRecoveryUrl() bool {
	if o != nil && o.RecoveryUrl != nil {
		return true
	}

	return false
}

// SetRecoveryUrl gets a reference to the given string and assigns it to the RecoveryUrl field.
func (o *IdentificationChallenge) SetRecoveryUrl(v string) {
	o.RecoveryUrl = &v
}

// GetPasswordlessUrl returns the PasswordlessUrl field value if set, zero value otherwise.
func (o *IdentificationChallenge) GetPasswordlessUrl() string {
	if o == nil || o.PasswordlessUrl == nil {
		var ret string
		return ret
	}
	return *o.PasswordlessUrl
}

// GetPasswordlessUrlOk returns a tuple with the PasswordlessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetPasswordlessUrlOk() (*string, bool) {
	if o == nil || o.PasswordlessUrl == nil {
		return nil, false
	}
	return o.PasswordlessUrl, true
}

// HasPasswordlessUrl returns a boolean if a field has been set.
func (o *IdentificationChallenge) HasPasswordlessUrl() bool {
	if o != nil && o.PasswordlessUrl != nil {
		return true
	}

	return false
}

// SetPasswordlessUrl gets a reference to the given string and assigns it to the PasswordlessUrl field.
func (o *IdentificationChallenge) SetPasswordlessUrl(v string) {
	o.PasswordlessUrl = &v
}

// GetPrimaryAction returns the PrimaryAction field value
func (o *IdentificationChallenge) GetPrimaryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryAction
}

// GetPrimaryActionOk returns a tuple with the PrimaryAction field value
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetPrimaryActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryAction, true
}

// SetPrimaryAction sets field value
func (o *IdentificationChallenge) SetPrimaryAction(v string) {
	o.PrimaryAction = v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *IdentificationChallenge) GetSources() []LoginSource {
	if o == nil || o.Sources == nil {
		var ret []LoginSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetSourcesOk() ([]LoginSource, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *IdentificationChallenge) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []LoginSource and assigns it to the Sources field.
func (o *IdentificationChallenge) SetSources(v []LoginSource) {
	o.Sources = v
}

// GetShowSourceLabels returns the ShowSourceLabels field value
func (o *IdentificationChallenge) GetShowSourceLabels() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowSourceLabels
}

// GetShowSourceLabelsOk returns a tuple with the ShowSourceLabels field value
// and a boolean to check if the value has been set.
func (o *IdentificationChallenge) GetShowSourceLabelsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowSourceLabels, true
}

// SetShowSourceLabels sets field value
func (o *IdentificationChallenge) SetShowSourceLabels(v bool) {
	o.ShowSourceLabels = v
}

func (o IdentificationChallenge) MarshalJSON() ([]byte, error) {
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
	if o.UserFields != nil {
		toSerialize["user_fields"] = o.UserFields
	}
	if true {
		toSerialize["password_fields"] = o.PasswordFields
	}
	if o.ApplicationPre != nil {
		toSerialize["application_pre"] = o.ApplicationPre
	}
	if o.EnrollUrl != nil {
		toSerialize["enroll_url"] = o.EnrollUrl
	}
	if o.RecoveryUrl != nil {
		toSerialize["recovery_url"] = o.RecoveryUrl
	}
	if o.PasswordlessUrl != nil {
		toSerialize["passwordless_url"] = o.PasswordlessUrl
	}
	if true {
		toSerialize["primary_action"] = o.PrimaryAction
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if true {
		toSerialize["show_source_labels"] = o.ShowSourceLabels
	}
	return json.Marshal(toSerialize)
}

type NullableIdentificationChallenge struct {
	value *IdentificationChallenge
	isSet bool
}

func (v NullableIdentificationChallenge) Get() *IdentificationChallenge {
	return v.value
}

func (v *NullableIdentificationChallenge) Set(val *IdentificationChallenge) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentificationChallenge) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentificationChallenge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentificationChallenge(val *IdentificationChallenge) *NullableIdentificationChallenge {
	return &NullableIdentificationChallenge{value: val, isSet: true}
}

func (v NullableIdentificationChallenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentificationChallenge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
