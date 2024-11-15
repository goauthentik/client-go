/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorWebAuthnStage AuthenticatorWebAuthnStage Serializer
type AuthenticatorWebAuthnStage struct {
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
	ConfigureFlow             NullableString                      `json:"configure_flow,omitempty"`
	FriendlyName              NullableString                      `json:"friendly_name,omitempty"`
	UserVerification          *UserVerificationEnum               `json:"user_verification,omitempty"`
	AuthenticatorAttachment   NullableAuthenticatorAttachmentEnum `json:"authenticator_attachment,omitempty"`
	ResidentKeyRequirement    *ResidentKeyRequirementEnum         `json:"resident_key_requirement,omitempty"`
	DeviceTypeRestrictions    []string                            `json:"device_type_restrictions,omitempty"`
	DeviceTypeRestrictionsObj []WebAuthnDeviceType                `json:"device_type_restrictions_obj"`
}

// NewAuthenticatorWebAuthnStage instantiates a new AuthenticatorWebAuthnStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorWebAuthnStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, deviceTypeRestrictionsObj []WebAuthnDeviceType) *AuthenticatorWebAuthnStage {
	this := AuthenticatorWebAuthnStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.DeviceTypeRestrictionsObj = deviceTypeRestrictionsObj
	return &this
}

// NewAuthenticatorWebAuthnStageWithDefaults instantiates a new AuthenticatorWebAuthnStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorWebAuthnStageWithDefaults() *AuthenticatorWebAuthnStage {
	this := AuthenticatorWebAuthnStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *AuthenticatorWebAuthnStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *AuthenticatorWebAuthnStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *AuthenticatorWebAuthnStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorWebAuthnStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *AuthenticatorWebAuthnStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *AuthenticatorWebAuthnStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *AuthenticatorWebAuthnStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *AuthenticatorWebAuthnStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *AuthenticatorWebAuthnStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *AuthenticatorWebAuthnStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *AuthenticatorWebAuthnStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *AuthenticatorWebAuthnStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorWebAuthnStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorWebAuthnStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *AuthenticatorWebAuthnStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorWebAuthnStage) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorWebAuthnStage) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *AuthenticatorWebAuthnStage) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *AuthenticatorWebAuthnStage) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *AuthenticatorWebAuthnStage) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *AuthenticatorWebAuthnStage) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorWebAuthnStage) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorWebAuthnStage) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *AuthenticatorWebAuthnStage) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *AuthenticatorWebAuthnStage) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *AuthenticatorWebAuthnStage) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *AuthenticatorWebAuthnStage) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *AuthenticatorWebAuthnStage) GetUserVerification() UserVerificationEnum {
	if o == nil || o.UserVerification == nil {
		var ret UserVerificationEnum
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnStage) GetUserVerificationOk() (*UserVerificationEnum, bool) {
	if o == nil || o.UserVerification == nil {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *AuthenticatorWebAuthnStage) HasUserVerification() bool {
	if o != nil && o.UserVerification != nil {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given UserVerificationEnum and assigns it to the UserVerification field.
func (o *AuthenticatorWebAuthnStage) SetUserVerification(v UserVerificationEnum) {
	o.UserVerification = &v
}

// GetAuthenticatorAttachment returns the AuthenticatorAttachment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorWebAuthnStage) GetAuthenticatorAttachment() AuthenticatorAttachmentEnum {
	if o == nil || o.AuthenticatorAttachment.Get() == nil {
		var ret AuthenticatorAttachmentEnum
		return ret
	}
	return *o.AuthenticatorAttachment.Get()
}

// GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorWebAuthnStage) GetAuthenticatorAttachmentOk() (*AuthenticatorAttachmentEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticatorAttachment.Get(), o.AuthenticatorAttachment.IsSet()
}

// HasAuthenticatorAttachment returns a boolean if a field has been set.
func (o *AuthenticatorWebAuthnStage) HasAuthenticatorAttachment() bool {
	if o != nil && o.AuthenticatorAttachment.IsSet() {
		return true
	}

	return false
}

// SetAuthenticatorAttachment gets a reference to the given NullableAuthenticatorAttachmentEnum and assigns it to the AuthenticatorAttachment field.
func (o *AuthenticatorWebAuthnStage) SetAuthenticatorAttachment(v AuthenticatorAttachmentEnum) {
	o.AuthenticatorAttachment.Set(&v)
}

// SetAuthenticatorAttachmentNil sets the value for AuthenticatorAttachment to be an explicit nil
func (o *AuthenticatorWebAuthnStage) SetAuthenticatorAttachmentNil() {
	o.AuthenticatorAttachment.Set(nil)
}

// UnsetAuthenticatorAttachment ensures that no value is present for AuthenticatorAttachment, not even an explicit nil
func (o *AuthenticatorWebAuthnStage) UnsetAuthenticatorAttachment() {
	o.AuthenticatorAttachment.Unset()
}

// GetResidentKeyRequirement returns the ResidentKeyRequirement field value if set, zero value otherwise.
func (o *AuthenticatorWebAuthnStage) GetResidentKeyRequirement() ResidentKeyRequirementEnum {
	if o == nil || o.ResidentKeyRequirement == nil {
		var ret ResidentKeyRequirementEnum
		return ret
	}
	return *o.ResidentKeyRequirement
}

// GetResidentKeyRequirementOk returns a tuple with the ResidentKeyRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnStage) GetResidentKeyRequirementOk() (*ResidentKeyRequirementEnum, bool) {
	if o == nil || o.ResidentKeyRequirement == nil {
		return nil, false
	}
	return o.ResidentKeyRequirement, true
}

// HasResidentKeyRequirement returns a boolean if a field has been set.
func (o *AuthenticatorWebAuthnStage) HasResidentKeyRequirement() bool {
	if o != nil && o.ResidentKeyRequirement != nil {
		return true
	}

	return false
}

// SetResidentKeyRequirement gets a reference to the given ResidentKeyRequirementEnum and assigns it to the ResidentKeyRequirement field.
func (o *AuthenticatorWebAuthnStage) SetResidentKeyRequirement(v ResidentKeyRequirementEnum) {
	o.ResidentKeyRequirement = &v
}

// GetDeviceTypeRestrictions returns the DeviceTypeRestrictions field value if set, zero value otherwise.
func (o *AuthenticatorWebAuthnStage) GetDeviceTypeRestrictions() []string {
	if o == nil || o.DeviceTypeRestrictions == nil {
		var ret []string
		return ret
	}
	return o.DeviceTypeRestrictions
}

// GetDeviceTypeRestrictionsOk returns a tuple with the DeviceTypeRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnStage) GetDeviceTypeRestrictionsOk() ([]string, bool) {
	if o == nil || o.DeviceTypeRestrictions == nil {
		return nil, false
	}
	return o.DeviceTypeRestrictions, true
}

// HasDeviceTypeRestrictions returns a boolean if a field has been set.
func (o *AuthenticatorWebAuthnStage) HasDeviceTypeRestrictions() bool {
	if o != nil && o.DeviceTypeRestrictions != nil {
		return true
	}

	return false
}

// SetDeviceTypeRestrictions gets a reference to the given []string and assigns it to the DeviceTypeRestrictions field.
func (o *AuthenticatorWebAuthnStage) SetDeviceTypeRestrictions(v []string) {
	o.DeviceTypeRestrictions = v
}

// GetDeviceTypeRestrictionsObj returns the DeviceTypeRestrictionsObj field value
func (o *AuthenticatorWebAuthnStage) GetDeviceTypeRestrictionsObj() []WebAuthnDeviceType {
	if o == nil {
		var ret []WebAuthnDeviceType
		return ret
	}

	return o.DeviceTypeRestrictionsObj
}

// GetDeviceTypeRestrictionsObjOk returns a tuple with the DeviceTypeRestrictionsObj field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorWebAuthnStage) GetDeviceTypeRestrictionsObjOk() ([]WebAuthnDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceTypeRestrictionsObj, true
}

// SetDeviceTypeRestrictionsObj sets field value
func (o *AuthenticatorWebAuthnStage) SetDeviceTypeRestrictionsObj(v []WebAuthnDeviceType) {
	o.DeviceTypeRestrictionsObj = v
}

func (o AuthenticatorWebAuthnStage) MarshalJSON() ([]byte, error) {
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
	if o.DeviceTypeRestrictions != nil {
		toSerialize["device_type_restrictions"] = o.DeviceTypeRestrictions
	}
	if true {
		toSerialize["device_type_restrictions_obj"] = o.DeviceTypeRestrictionsObj
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorWebAuthnStage struct {
	value *AuthenticatorWebAuthnStage
	isSet bool
}

func (v NullableAuthenticatorWebAuthnStage) Get() *AuthenticatorWebAuthnStage {
	return v.value
}

func (v *NullableAuthenticatorWebAuthnStage) Set(val *AuthenticatorWebAuthnStage) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorWebAuthnStage) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorWebAuthnStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorWebAuthnStage(val *AuthenticatorWebAuthnStage) *NullableAuthenticatorWebAuthnStage {
	return &NullableAuthenticatorWebAuthnStage{value: val, isSet: true}
}

func (v NullableAuthenticatorWebAuthnStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorWebAuthnStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
