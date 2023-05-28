/*
authentik

Making authentication simple.

API version: 2023.5.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticateWebAuthnStage AuthenticateWebAuthnStage Serializer
type AuthenticateWebAuthnStage struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// Get object type so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string    `json:"meta_model_name"`
	FlowSet       []FlowSet `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow           NullableString                      `json:"configure_flow,omitempty"`
	FriendlyName            NullableString                      `json:"friendly_name,omitempty"`
	UserVerification        *UserVerificationEnum               `json:"user_verification,omitempty"`
	AuthenticatorAttachment NullableAuthenticatorAttachmentEnum `json:"authenticator_attachment,omitempty"`
	ResidentKeyRequirement  *ResidentKeyRequirementEnum         `json:"resident_key_requirement,omitempty"`
}

// NewAuthenticateWebAuthnStage instantiates a new AuthenticateWebAuthnStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticateWebAuthnStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string) *AuthenticateWebAuthnStage {
	this := AuthenticateWebAuthnStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewAuthenticateWebAuthnStageWithDefaults instantiates a new AuthenticateWebAuthnStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticateWebAuthnStageWithDefaults() *AuthenticateWebAuthnStage {
	this := AuthenticateWebAuthnStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *AuthenticateWebAuthnStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *AuthenticateWebAuthnStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *AuthenticateWebAuthnStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticateWebAuthnStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *AuthenticateWebAuthnStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *AuthenticateWebAuthnStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *AuthenticateWebAuthnStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *AuthenticateWebAuthnStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *AuthenticateWebAuthnStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *AuthenticateWebAuthnStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *AuthenticateWebAuthnStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *AuthenticateWebAuthnStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticateWebAuthnStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *AuthenticateWebAuthnStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticateWebAuthnStage) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticateWebAuthnStage) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStage) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *AuthenticateWebAuthnStage) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *AuthenticateWebAuthnStage) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *AuthenticateWebAuthnStage) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticateWebAuthnStage) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticateWebAuthnStage) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStage) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *AuthenticateWebAuthnStage) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *AuthenticateWebAuthnStage) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *AuthenticateWebAuthnStage) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *AuthenticateWebAuthnStage) GetUserVerification() UserVerificationEnum {
	if o == nil || o.UserVerification == nil {
		var ret UserVerificationEnum
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStage) GetUserVerificationOk() (*UserVerificationEnum, bool) {
	if o == nil || o.UserVerification == nil {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStage) HasUserVerification() bool {
	if o != nil && o.UserVerification != nil {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given UserVerificationEnum and assigns it to the UserVerification field.
func (o *AuthenticateWebAuthnStage) SetUserVerification(v UserVerificationEnum) {
	o.UserVerification = &v
}

// GetAuthenticatorAttachment returns the AuthenticatorAttachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticateWebAuthnStage) GetAuthenticatorAttachment() AuthenticatorAttachmentEnum {
	if o == nil || o.AuthenticatorAttachment.Get() == nil {
		var ret AuthenticatorAttachmentEnum
		return ret
	}
	return *o.AuthenticatorAttachment.Get()
}

// GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticateWebAuthnStage) GetAuthenticatorAttachmentOk() (*AuthenticatorAttachmentEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticatorAttachment.Get(), o.AuthenticatorAttachment.IsSet()
}

// HasAuthenticatorAttachment returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStage) HasAuthenticatorAttachment() bool {
	if o != nil && o.AuthenticatorAttachment.IsSet() {
		return true
	}

	return false
}

// SetAuthenticatorAttachment gets a reference to the given NullableAuthenticatorAttachmentEnum and assigns it to the AuthenticatorAttachment field.
func (o *AuthenticateWebAuthnStage) SetAuthenticatorAttachment(v AuthenticatorAttachmentEnum) {
	o.AuthenticatorAttachment.Set(&v)
}

// SetAuthenticatorAttachmentNil sets the value for AuthenticatorAttachment to be an explicit nil
func (o *AuthenticateWebAuthnStage) SetAuthenticatorAttachmentNil() {
	o.AuthenticatorAttachment.Set(nil)
}

// UnsetAuthenticatorAttachment ensures that no value is present for AuthenticatorAttachment, not even an explicit nil
func (o *AuthenticateWebAuthnStage) UnsetAuthenticatorAttachment() {
	o.AuthenticatorAttachment.Unset()
}

// GetResidentKeyRequirement returns the ResidentKeyRequirement field value if set, zero value otherwise.
func (o *AuthenticateWebAuthnStage) GetResidentKeyRequirement() ResidentKeyRequirementEnum {
	if o == nil || o.ResidentKeyRequirement == nil {
		var ret ResidentKeyRequirementEnum
		return ret
	}
	return *o.ResidentKeyRequirement
}

// GetResidentKeyRequirementOk returns a tuple with the ResidentKeyRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateWebAuthnStage) GetResidentKeyRequirementOk() (*ResidentKeyRequirementEnum, bool) {
	if o == nil || o.ResidentKeyRequirement == nil {
		return nil, false
	}
	return o.ResidentKeyRequirement, true
}

// HasResidentKeyRequirement returns a boolean if a field has been set.
func (o *AuthenticateWebAuthnStage) HasResidentKeyRequirement() bool {
	if o != nil && o.ResidentKeyRequirement != nil {
		return true
	}

	return false
}

// SetResidentKeyRequirement gets a reference to the given ResidentKeyRequirementEnum and assigns it to the ResidentKeyRequirement field.
func (o *AuthenticateWebAuthnStage) SetResidentKeyRequirement(v ResidentKeyRequirementEnum) {
	o.ResidentKeyRequirement = &v
}

func (o AuthenticateWebAuthnStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
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

type NullableAuthenticateWebAuthnStage struct {
	value *AuthenticateWebAuthnStage
	isSet bool
}

func (v NullableAuthenticateWebAuthnStage) Get() *AuthenticateWebAuthnStage {
	return v.value
}

func (v *NullableAuthenticateWebAuthnStage) Set(val *AuthenticateWebAuthnStage) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticateWebAuthnStage) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticateWebAuthnStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticateWebAuthnStage(val *AuthenticateWebAuthnStage) *NullableAuthenticateWebAuthnStage {
	return &NullableAuthenticateWebAuthnStage{value: val, isSet: true}
}

func (v NullableAuthenticateWebAuthnStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticateWebAuthnStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
