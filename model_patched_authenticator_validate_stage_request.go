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

// PatchedAuthenticatorValidateStageRequest AuthenticatorValidateStage Serializer
type PatchedAuthenticatorValidateStageRequest struct {
	Name                *string                  `json:"name,omitempty"`
	FlowSet             []FlowSetRequest         `json:"flow_set,omitempty"`
	NotConfiguredAction *NotConfiguredActionEnum `json:"not_configured_action,omitempty"`
	// Device classes which can be used to authenticate
	DeviceClasses []DeviceClassesEnum `json:"device_classes,omitempty"`
	// Stages used to configure Authenticator when user doesn't have any compatible devices. After this configuration Stage passes, the user is not prompted again.
	ConfigurationStages []string `json:"configuration_stages,omitempty"`
	// If any of the user's device has been used within this threshold, this stage will be skipped
	LastAuthThreshold *string `json:"last_auth_threshold,omitempty"`
	// Enforce user verification for WebAuthn devices.
	WebauthnUserVerification   *UserVerificationEnum `json:"webauthn_user_verification,omitempty"`
	WebauthnAllowedDeviceTypes []string              `json:"webauthn_allowed_device_types,omitempty"`
}

// NewPatchedAuthenticatorValidateStageRequest instantiates a new PatchedAuthenticatorValidateStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAuthenticatorValidateStageRequest() *PatchedAuthenticatorValidateStageRequest {
	this := PatchedAuthenticatorValidateStageRequest{}
	return &this
}

// NewPatchedAuthenticatorValidateStageRequestWithDefaults instantiates a new PatchedAuthenticatorValidateStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAuthenticatorValidateStageRequestWithDefaults() *PatchedAuthenticatorValidateStageRequest {
	this := PatchedAuthenticatorValidateStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAuthenticatorValidateStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorValidateStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAuthenticatorValidateStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAuthenticatorValidateStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedAuthenticatorValidateStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorValidateStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedAuthenticatorValidateStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedAuthenticatorValidateStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetNotConfiguredAction returns the NotConfiguredAction field value if set, zero value otherwise.
func (o *PatchedAuthenticatorValidateStageRequest) GetNotConfiguredAction() NotConfiguredActionEnum {
	if o == nil || o.NotConfiguredAction == nil {
		var ret NotConfiguredActionEnum
		return ret
	}
	return *o.NotConfiguredAction
}

// GetNotConfiguredActionOk returns a tuple with the NotConfiguredAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorValidateStageRequest) GetNotConfiguredActionOk() (*NotConfiguredActionEnum, bool) {
	if o == nil || o.NotConfiguredAction == nil {
		return nil, false
	}
	return o.NotConfiguredAction, true
}

// HasNotConfiguredAction returns a boolean if a field has been set.
func (o *PatchedAuthenticatorValidateStageRequest) HasNotConfiguredAction() bool {
	if o != nil && o.NotConfiguredAction != nil {
		return true
	}

	return false
}

// SetNotConfiguredAction gets a reference to the given NotConfiguredActionEnum and assigns it to the NotConfiguredAction field.
func (o *PatchedAuthenticatorValidateStageRequest) SetNotConfiguredAction(v NotConfiguredActionEnum) {
	o.NotConfiguredAction = &v
}

// GetDeviceClasses returns the DeviceClasses field value if set, zero value otherwise.
func (o *PatchedAuthenticatorValidateStageRequest) GetDeviceClasses() []DeviceClassesEnum {
	if o == nil || o.DeviceClasses == nil {
		var ret []DeviceClassesEnum
		return ret
	}
	return o.DeviceClasses
}

// GetDeviceClassesOk returns a tuple with the DeviceClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorValidateStageRequest) GetDeviceClassesOk() ([]DeviceClassesEnum, bool) {
	if o == nil || o.DeviceClasses == nil {
		return nil, false
	}
	return o.DeviceClasses, true
}

// HasDeviceClasses returns a boolean if a field has been set.
func (o *PatchedAuthenticatorValidateStageRequest) HasDeviceClasses() bool {
	if o != nil && o.DeviceClasses != nil {
		return true
	}

	return false
}

// SetDeviceClasses gets a reference to the given []DeviceClassesEnum and assigns it to the DeviceClasses field.
func (o *PatchedAuthenticatorValidateStageRequest) SetDeviceClasses(v []DeviceClassesEnum) {
	o.DeviceClasses = v
}

// GetConfigurationStages returns the ConfigurationStages field value if set, zero value otherwise.
func (o *PatchedAuthenticatorValidateStageRequest) GetConfigurationStages() []string {
	if o == nil || o.ConfigurationStages == nil {
		var ret []string
		return ret
	}
	return o.ConfigurationStages
}

// GetConfigurationStagesOk returns a tuple with the ConfigurationStages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorValidateStageRequest) GetConfigurationStagesOk() ([]string, bool) {
	if o == nil || o.ConfigurationStages == nil {
		return nil, false
	}
	return o.ConfigurationStages, true
}

// HasConfigurationStages returns a boolean if a field has been set.
func (o *PatchedAuthenticatorValidateStageRequest) HasConfigurationStages() bool {
	if o != nil && o.ConfigurationStages != nil {
		return true
	}

	return false
}

// SetConfigurationStages gets a reference to the given []string and assigns it to the ConfigurationStages field.
func (o *PatchedAuthenticatorValidateStageRequest) SetConfigurationStages(v []string) {
	o.ConfigurationStages = v
}

// GetLastAuthThreshold returns the LastAuthThreshold field value if set, zero value otherwise.
func (o *PatchedAuthenticatorValidateStageRequest) GetLastAuthThreshold() string {
	if o == nil || o.LastAuthThreshold == nil {
		var ret string
		return ret
	}
	return *o.LastAuthThreshold
}

// GetLastAuthThresholdOk returns a tuple with the LastAuthThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorValidateStageRequest) GetLastAuthThresholdOk() (*string, bool) {
	if o == nil || o.LastAuthThreshold == nil {
		return nil, false
	}
	return o.LastAuthThreshold, true
}

// HasLastAuthThreshold returns a boolean if a field has been set.
func (o *PatchedAuthenticatorValidateStageRequest) HasLastAuthThreshold() bool {
	if o != nil && o.LastAuthThreshold != nil {
		return true
	}

	return false
}

// SetLastAuthThreshold gets a reference to the given string and assigns it to the LastAuthThreshold field.
func (o *PatchedAuthenticatorValidateStageRequest) SetLastAuthThreshold(v string) {
	o.LastAuthThreshold = &v
}

// GetWebauthnUserVerification returns the WebauthnUserVerification field value if set, zero value otherwise.
func (o *PatchedAuthenticatorValidateStageRequest) GetWebauthnUserVerification() UserVerificationEnum {
	if o == nil || o.WebauthnUserVerification == nil {
		var ret UserVerificationEnum
		return ret
	}
	return *o.WebauthnUserVerification
}

// GetWebauthnUserVerificationOk returns a tuple with the WebauthnUserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorValidateStageRequest) GetWebauthnUserVerificationOk() (*UserVerificationEnum, bool) {
	if o == nil || o.WebauthnUserVerification == nil {
		return nil, false
	}
	return o.WebauthnUserVerification, true
}

// HasWebauthnUserVerification returns a boolean if a field has been set.
func (o *PatchedAuthenticatorValidateStageRequest) HasWebauthnUserVerification() bool {
	if o != nil && o.WebauthnUserVerification != nil {
		return true
	}

	return false
}

// SetWebauthnUserVerification gets a reference to the given UserVerificationEnum and assigns it to the WebauthnUserVerification field.
func (o *PatchedAuthenticatorValidateStageRequest) SetWebauthnUserVerification(v UserVerificationEnum) {
	o.WebauthnUserVerification = &v
}

// GetWebauthnAllowedDeviceTypes returns the WebauthnAllowedDeviceTypes field value if set, zero value otherwise.
func (o *PatchedAuthenticatorValidateStageRequest) GetWebauthnAllowedDeviceTypes() []string {
	if o == nil || o.WebauthnAllowedDeviceTypes == nil {
		var ret []string
		return ret
	}
	return o.WebauthnAllowedDeviceTypes
}

// GetWebauthnAllowedDeviceTypesOk returns a tuple with the WebauthnAllowedDeviceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorValidateStageRequest) GetWebauthnAllowedDeviceTypesOk() ([]string, bool) {
	if o == nil || o.WebauthnAllowedDeviceTypes == nil {
		return nil, false
	}
	return o.WebauthnAllowedDeviceTypes, true
}

// HasWebauthnAllowedDeviceTypes returns a boolean if a field has been set.
func (o *PatchedAuthenticatorValidateStageRequest) HasWebauthnAllowedDeviceTypes() bool {
	if o != nil && o.WebauthnAllowedDeviceTypes != nil {
		return true
	}

	return false
}

// SetWebauthnAllowedDeviceTypes gets a reference to the given []string and assigns it to the WebauthnAllowedDeviceTypes field.
func (o *PatchedAuthenticatorValidateStageRequest) SetWebauthnAllowedDeviceTypes(v []string) {
	o.WebauthnAllowedDeviceTypes = v
}

func (o PatchedAuthenticatorValidateStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	return json.Marshal(toSerialize)
}

type NullablePatchedAuthenticatorValidateStageRequest struct {
	value *PatchedAuthenticatorValidateStageRequest
	isSet bool
}

func (v NullablePatchedAuthenticatorValidateStageRequest) Get() *PatchedAuthenticatorValidateStageRequest {
	return v.value
}

func (v *NullablePatchedAuthenticatorValidateStageRequest) Set(val *PatchedAuthenticatorValidateStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAuthenticatorValidateStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAuthenticatorValidateStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAuthenticatorValidateStageRequest(val *PatchedAuthenticatorValidateStageRequest) *NullablePatchedAuthenticatorValidateStageRequest {
	return &NullablePatchedAuthenticatorValidateStageRequest{value: val, isSet: true}
}

func (v NullablePatchedAuthenticatorValidateStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAuthenticatorValidateStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
