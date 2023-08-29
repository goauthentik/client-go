/*
authentik

Making authentication simple.

API version: 2023.8.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticateWebAuthnStageRequest AuthenticateWebAuthnStage Serializer
type AuthenticateWebAuthnStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow           NullableString                      `json:"configure_flow,omitempty"`
	FriendlyName            NullableString                      `json:"friendly_name,omitempty"`
	UserVerification        *UserVerificationEnum               `json:"user_verification,omitempty"`
	AuthenticatorAttachment NullableAuthenticatorAttachmentEnum `json:"authenticator_attachment,omitempty"`
	ResidentKeyRequirement  *ResidentKeyRequirementEnum         `json:"resident_key_requirement,omitempty"`
}

// NewAuthenticateWebAuthnStageRequest instantiates a new AuthenticateWebAuthnStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticateWebAuthnStageRequest(name string) *AuthenticateWebAuthnStageRequest {
	this := AuthenticateWebAuthnStageRequest{}
	this.Name = name
	return &this
}

// NewAuthenticateWebAuthnStageRequestWithDefaults instantiates a new AuthenticateWebAuthnStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticateWebAuthnStageRequestWithDefaults() *AuthenticateWebAuthnStageRequest {
	this := AuthenticateWebAuthnStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AuthenticateWebAuthnStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticateWebAuthnStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticateWebAuthnStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *AuthenticateWebAuthnStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticateWebAuthnStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticateWebAuthnStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *AuthenticateWebAuthnStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *AuthenticateWebAuthnStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *AuthenticateWebAuthnStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticateWebAuthnStageRequest) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticateWebAuthnStageRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStageRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *AuthenticateWebAuthnStageRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *AuthenticateWebAuthnStageRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *AuthenticateWebAuthnStageRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *AuthenticateWebAuthnStageRequest) GetUserVerification() UserVerificationEnum {
	if o == nil || o.UserVerification == nil {
		var ret UserVerificationEnum
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStageRequest) GetUserVerificationOk() (*UserVerificationEnum, bool) {
	if o == nil || o.UserVerification == nil {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStageRequest) HasUserVerification() bool {
	if o != nil && o.UserVerification != nil {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given UserVerificationEnum and assigns it to the UserVerification field.
func (o *AuthenticateWebAuthnStageRequest) SetUserVerification(v UserVerificationEnum) {
	o.UserVerification = &v
}

// GetAuthenticatorAttachment returns the AuthenticatorAttachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticateWebAuthnStageRequest) GetAuthenticatorAttachment() AuthenticatorAttachmentEnum {
	if o == nil || o.AuthenticatorAttachment.Get() == nil {
		var ret AuthenticatorAttachmentEnum
		return ret
	}
	return *o.AuthenticatorAttachment.Get()
}

// GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticateWebAuthnStageRequest) GetAuthenticatorAttachmentOk() (*AuthenticatorAttachmentEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticatorAttachment.Get(), o.AuthenticatorAttachment.IsSet()
}

// HasAuthenticatorAttachment returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStageRequest) HasAuthenticatorAttachment() bool {
	if o != nil && o.AuthenticatorAttachment.IsSet() {
		return true
	}

	return false
}

// SetAuthenticatorAttachment gets a reference to the given NullableAuthenticatorAttachmentEnum and assigns it to the AuthenticatorAttachment field.
func (o *AuthenticateWebAuthnStageRequest) SetAuthenticatorAttachment(v AuthenticatorAttachmentEnum) {
	o.AuthenticatorAttachment.Set(&v)
}

// SetAuthenticatorAttachmentNil sets the value for AuthenticatorAttachment to be an explicit nil
func (o *AuthenticateWebAuthnStageRequest) SetAuthenticatorAttachmentNil() {
	o.AuthenticatorAttachment.Set(nil)
}

// UnsetAuthenticatorAttachment ensures that no value is present for AuthenticatorAttachment, not even an explicit nil
func (o *AuthenticateWebAuthnStageRequest) UnsetAuthenticatorAttachment() {
	o.AuthenticatorAttachment.Unset()
}

// GetResidentKeyRequirement returns the ResidentKeyRequirement field value if set, zero value otherwise.
func (o *AuthenticateWebAuthnStageRequest) GetResidentKeyRequirement() ResidentKeyRequirementEnum {
	if o == nil || o.ResidentKeyRequirement == nil {
		var ret ResidentKeyRequirementEnum
		return ret
	}
	return *o.ResidentKeyRequirement
}

// GetResidentKeyRequirementOk returns a tuple with the ResidentKeyRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStageRequest) GetResidentKeyRequirementOk() (*ResidentKeyRequirementEnum, bool) {
	if o == nil || o.ResidentKeyRequirement == nil {
		return nil, false
	}
	return o.ResidentKeyRequirement, true
}

// HasResidentKeyRequirement returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStageRequest) HasResidentKeyRequirement() bool {
	if o != nil && o.ResidentKeyRequirement != nil {
		return true
	}

	return false
}

// SetResidentKeyRequirement gets a reference to the given ResidentKeyRequirementEnum and assigns it to the ResidentKeyRequirement field.
func (o *AuthenticateWebAuthnStageRequest) SetResidentKeyRequirement(v ResidentKeyRequirementEnum) {
	o.ResidentKeyRequirement = &v
}

func (o AuthenticateWebAuthnStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ConfigureFlow.IsSet() {
		toSerialize["configure_flow"] = o.ConfigureFlow.Get()
	}
	if o.FriendlyName.IsSet() {
		toSerialize["friendly_name"] = o.FriendlyName.Get()
	}
	if o.UserVerification != nil {
		toSerialize["user_verification"] = o.UserVerification
	}
	if o.AuthenticatorAttachment.IsSet() {
		toSerialize["authenticator_attachment"] = o.AuthenticatorAttachment.Get()
	}
	if o.ResidentKeyRequirement != nil {
		toSerialize["resident_key_requirement"] = o.ResidentKeyRequirement
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticateWebAuthnStageRequest struct {
	value *AuthenticateWebAuthnStageRequest
	isSet bool
}

func (v NullableAuthenticateWebAuthnStageRequest) Get() *AuthenticateWebAuthnStageRequest {
	return v.value
}

func (v *NullableAuthenticateWebAuthnStageRequest) Set(val *AuthenticateWebAuthnStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticateWebAuthnStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticateWebAuthnStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticateWebAuthnStageRequest(val *AuthenticateWebAuthnStageRequest) *NullableAuthenticateWebAuthnStageRequest {
	return &NullableAuthenticateWebAuthnStageRequest{value: val, isSet: true}
}

func (v NullableAuthenticateWebAuthnStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticateWebAuthnStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
