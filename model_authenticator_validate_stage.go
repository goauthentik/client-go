/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorValidateStage AuthenticatorValidateStage Serializer
type AuthenticatorValidateStage struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// Get object type so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName       string                   `json:"meta_model_name"`
	FlowSet             []FlowSet                `json:"flow_set,omitempty"`
	NotConfiguredAction *NotConfiguredActionEnum `json:"not_configured_action,omitempty"`
	// Device classes which can be used to authenticate
	DeviceClasses []DeviceClassesEnum `json:"device_classes,omitempty"`
	// Stages used to configure Authenticator when user doesn't have any compatible devices. After this configuration Stage passes, the user is not prompted again.
	ConfigurationStages []string `json:"configuration_stages,omitempty"`
	// If any of the user's device has been used within this threshold, this stage will be skipped
	LastAuthThreshold *string `json:"last_auth_threshold,omitempty"`
	// Enforce user verification for WebAuthn devices.
	WebauthnUserVerification      *UserVerificationEnum `json:"webauthn_user_verification,omitempty"`
	WebauthnAllowedDeviceTypes    []string              `json:"webauthn_allowed_device_types,omitempty"`
	WebauthnAllowedDeviceTypesObj []WebAuthnDeviceType  `json:"webauthn_allowed_device_types_obj"`
}

// NewAuthenticatorValidateStage instantiates a new AuthenticatorValidateStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorValidateStage(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, webauthnAllowedDeviceTypesObj []WebAuthnDeviceType) *AuthenticatorValidateStage {
	this := AuthenticatorValidateStage{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.WebauthnAllowedDeviceTypesObj = webauthnAllowedDeviceTypesObj
	return &this
}

// NewAuthenticatorValidateStageWithDefaults instantiates a new AuthenticatorValidateStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorValidateStageWithDefaults() *AuthenticatorValidateStage {
	this := AuthenticatorValidateStage{}
	return &this
}

// GetPk returns the Pk field value
func (o *AuthenticatorValidateStage) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *AuthenticatorValidateStage) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *AuthenticatorValidateStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorValidateStage) SetName(v string) {
	o.Name = v
}

// GetComponent returns the Component field value
func (o *AuthenticatorValidateStage) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *AuthenticatorValidateStage) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *AuthenticatorValidateStage) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *AuthenticatorValidateStage) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *AuthenticatorValidateStage) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *AuthenticatorValidateStage) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *AuthenticatorValidateStage) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *AuthenticatorValidateStage) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorValidateStage) GetFlowSet() []FlowSet {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSet
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetFlowSetOk() ([]FlowSet, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorValidateStage) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSet and assigns it to the FlowSet field.
func (o *AuthenticatorValidateStage) SetFlowSet(v []FlowSet) {
	o.FlowSet = v
}

// GetNotConfiguredAction returns the NotConfiguredAction field value if set, zero value otherwise.
func (o *AuthenticatorValidateStage) GetNotConfiguredAction() NotConfiguredActionEnum {
	if o == nil || o.NotConfiguredAction == nil {
		var ret NotConfiguredActionEnum
		return ret
	}
	return *o.NotConfiguredAction
}

// GetNotConfiguredActionOk returns a tuple with the NotConfiguredAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetNotConfiguredActionOk() (*NotConfiguredActionEnum, bool) {
	if o == nil || o.NotConfiguredAction == nil {
		return nil, false
	}
	return o.NotConfiguredAction, true
}

// HasNotConfiguredAction returns a boolean if a field has been set.
func (o *AuthenticatorValidateStage) HasNotConfiguredAction() bool {
	if o != nil && o.NotConfiguredAction != nil {
		return true
	}

	return false
}

// SetNotConfiguredAction gets a reference to the given NotConfiguredActionEnum and assigns it to the NotConfiguredAction field.
func (o *AuthenticatorValidateStage) SetNotConfiguredAction(v NotConfiguredActionEnum) {
	o.NotConfiguredAction = &v
}

// GetDeviceClasses returns the DeviceClasses field value if set, zero value otherwise.
func (o *AuthenticatorValidateStage) GetDeviceClasses() []DeviceClassesEnum {
	if o == nil || o.DeviceClasses == nil {
		var ret []DeviceClassesEnum
		return ret
	}
	return o.DeviceClasses
}

// GetDeviceClassesOk returns a tuple with the DeviceClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetDeviceClassesOk() ([]DeviceClassesEnum, bool) {
	if o == nil || o.DeviceClasses == nil {
		return nil, false
	}
	return o.DeviceClasses, true
}

// HasDeviceClasses returns a boolean if a field has been set.
func (o *AuthenticatorValidateStage) HasDeviceClasses() bool {
	if o != nil && o.DeviceClasses != nil {
		return true
	}

	return false
}

// SetDeviceClasses gets a reference to the given []DeviceClassesEnum and assigns it to the DeviceClasses field.
func (o *AuthenticatorValidateStage) SetDeviceClasses(v []DeviceClassesEnum) {
	o.DeviceClasses = v
}

// GetConfigurationStages returns the ConfigurationStages field value if set, zero value otherwise.
func (o *AuthenticatorValidateStage) GetConfigurationStages() []string {
	if o == nil || o.ConfigurationStages == nil {
		var ret []string
		return ret
	}
	return o.ConfigurationStages
}

// GetConfigurationStagesOk returns a tuple with the ConfigurationStages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetConfigurationStagesOk() ([]string, bool) {
	if o == nil || o.ConfigurationStages == nil {
		return nil, false
	}
	return o.ConfigurationStages, true
}

// HasConfigurationStages returns a boolean if a field has been set.
func (o *AuthenticatorValidateStage) HasConfigurationStages() bool {
	if o != nil && o.ConfigurationStages != nil {
		return true
	}

	return false
}

// SetConfigurationStages gets a reference to the given []string and assigns it to the ConfigurationStages field.
func (o *AuthenticatorValidateStage) SetConfigurationStages(v []string) {
	o.ConfigurationStages = v
}

// GetLastAuthThreshold returns the LastAuthThreshold field value if set, zero value otherwise.
func (o *AuthenticatorValidateStage) GetLastAuthThreshold() string {
	if o == nil || o.LastAuthThreshold == nil {
		var ret string
		return ret
	}
	return *o.LastAuthThreshold
}

// GetLastAuthThresholdOk returns a tuple with the LastAuthThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetLastAuthThresholdOk() (*string, bool) {
	if o == nil || o.LastAuthThreshold == nil {
		return nil, false
	}
	return o.LastAuthThreshold, true
}

// HasLastAuthThreshold returns a boolean if a field has been set.
func (o *AuthenticatorValidateStage) HasLastAuthThreshold() bool {
	if o != nil && o.LastAuthThreshold != nil {
		return true
	}

	return false
}

// SetLastAuthThreshold gets a reference to the given string and assigns it to the LastAuthThreshold field.
func (o *AuthenticatorValidateStage) SetLastAuthThreshold(v string) {
	o.LastAuthThreshold = &v
}

// GetWebauthnUserVerification returns the WebauthnUserVerification field value if set, zero value otherwise.
func (o *AuthenticatorValidateStage) GetWebauthnUserVerification() UserVerificationEnum {
	if o == nil || o.WebauthnUserVerification == nil {
		var ret UserVerificationEnum
		return ret
	}
	return *o.WebauthnUserVerification
}

// GetWebauthnUserVerificationOk returns a tuple with the WebauthnUserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetWebauthnUserVerificationOk() (*UserVerificationEnum, bool) {
	if o == nil || o.WebauthnUserVerification == nil {
		return nil, false
	}
	return o.WebauthnUserVerification, true
}

// HasWebauthnUserVerification returns a boolean if a field has been set.
func (o *AuthenticatorValidateStage) HasWebauthnUserVerification() bool {
	if o != nil && o.WebauthnUserVerification != nil {
		return true
	}

	return false
}

// SetWebauthnUserVerification gets a reference to the given UserVerificationEnum and assigns it to the WebauthnUserVerification field.
func (o *AuthenticatorValidateStage) SetWebauthnUserVerification(v UserVerificationEnum) {
	o.WebauthnUserVerification = &v
}

// GetWebauthnAllowedDeviceTypes returns the WebauthnAllowedDeviceTypes field value if set, zero value otherwise.
func (o *AuthenticatorValidateStage) GetWebauthnAllowedDeviceTypes() []string {
	if o == nil || o.WebauthnAllowedDeviceTypes == nil {
		var ret []string
		return ret
	}
	return o.WebauthnAllowedDeviceTypes
}

// GetWebauthnAllowedDeviceTypesOk returns a tuple with the WebauthnAllowedDeviceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetWebauthnAllowedDeviceTypesOk() ([]string, bool) {
	if o == nil || o.WebauthnAllowedDeviceTypes == nil {
		return nil, false
	}
	return o.WebauthnAllowedDeviceTypes, true
}

// HasWebauthnAllowedDeviceTypes returns a boolean if a field has been set.
func (o *AuthenticatorValidateStage) HasWebauthnAllowedDeviceTypes() bool {
	if o != nil && o.WebauthnAllowedDeviceTypes != nil {
		return true
	}

	return false
}

// SetWebauthnAllowedDeviceTypes gets a reference to the given []string and assigns it to the WebauthnAllowedDeviceTypes field.
func (o *AuthenticatorValidateStage) SetWebauthnAllowedDeviceTypes(v []string) {
	o.WebauthnAllowedDeviceTypes = v
}

// GetWebauthnAllowedDeviceTypesObj returns the WebauthnAllowedDeviceTypesObj field value
func (o *AuthenticatorValidateStage) GetWebauthnAllowedDeviceTypesObj() []WebAuthnDeviceType {
	if o == nil {
		var ret []WebAuthnDeviceType
		return ret
	}

	return o.WebauthnAllowedDeviceTypesObj
}

// GetWebauthnAllowedDeviceTypesObjOk returns a tuple with the WebauthnAllowedDeviceTypesObj field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStage) GetWebauthnAllowedDeviceTypesObjOk() ([]WebAuthnDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebauthnAllowedDeviceTypesObj, true
}

// SetWebauthnAllowedDeviceTypesObj sets field value
func (o *AuthenticatorValidateStage) SetWebauthnAllowedDeviceTypesObj(v []WebAuthnDeviceType) {
	o.WebauthnAllowedDeviceTypesObj = v
}

func (o AuthenticatorValidateStage) MarshalJSON() ([]byte, error) {
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
	if o.NotConfiguredAction != nil {
		toSerialize["not_configured_action"] = o.NotConfiguredAction
	}
	if o.DeviceClasses != nil {
		toSerialize["device_classes"] = o.DeviceClasses
	}
	if o.ConfigurationStages != nil {
		toSerialize["configuration_stages"] = o.ConfigurationStages
	}
	if o.LastAuthThreshold != nil {
		toSerialize["last_auth_threshold"] = o.LastAuthThreshold
	}
	if o.WebauthnUserVerification != nil {
		toSerialize["webauthn_user_verification"] = o.WebauthnUserVerification
	}
	if o.WebauthnAllowedDeviceTypes != nil {
		toSerialize["webauthn_allowed_device_types"] = o.WebauthnAllowedDeviceTypes
	}
	if true {
		toSerialize["webauthn_allowed_device_types_obj"] = o.WebauthnAllowedDeviceTypesObj
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorValidateStage struct {
	value *AuthenticatorValidateStage
	isSet bool
}

func (v NullableAuthenticatorValidateStage) Get() *AuthenticatorValidateStage {
	return v.value
}

func (v *NullableAuthenticatorValidateStage) Set(val *AuthenticatorValidateStage) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorValidateStage) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorValidateStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorValidateStage(val *AuthenticatorValidateStage) *NullableAuthenticatorValidateStage {
	return &NullableAuthenticatorValidateStage{value: val, isSet: true}
}

func (v NullableAuthenticatorValidateStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorValidateStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
