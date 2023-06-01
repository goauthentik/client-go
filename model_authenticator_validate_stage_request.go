/*
authentik

Making authentication simple.

API version: 2023.5.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorValidateStageRequest AuthenticatorValidateStage Serializer
type AuthenticatorValidateStageRequest struct {
	Name                string                   `json:"name"`
	FlowSet             []FlowSetRequest         `json:"flow_set,omitempty"`
	NotConfiguredAction *NotConfiguredActionEnum `json:"not_configured_action,omitempty"`
	// Device classes which can be used to authenticate
	DeviceClasses []DeviceClassesEnum `json:"device_classes,omitempty"`
	// Stages used to configure Authenticator when user doesn't have any compatible devices. After this configuration Stage passes, the user is not prompted again.
	ConfigurationStages []string `json:"configuration_stages,omitempty"`
	// If any of the user's device has been used within this threshold, this stage will be skipped
	LastAuthThreshold *string `json:"last_auth_threshold,omitempty"`
	// Enforce user verification for WebAuthn devices.  * `required` - Required * `preferred` - Preferred * `discouraged` - Discouraged
	WebauthnUserVerification *UserVerificationEnum `json:"webauthn_user_verification,omitempty"`
}

// NewAuthenticatorValidateStageRequest instantiates a new AuthenticatorValidateStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorValidateStageRequest(name string) *AuthenticatorValidateStageRequest {
	this := AuthenticatorValidateStageRequest{}
	this.Name = name
	return &this
}

// NewAuthenticatorValidateStageRequestWithDefaults instantiates a new AuthenticatorValidateStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorValidateStageRequestWithDefaults() *AuthenticatorValidateStageRequest {
	this := AuthenticatorValidateStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AuthenticatorValidateStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorValidateStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorValidateStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorValidateStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *AuthenticatorValidateStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetNotConfiguredAction returns the NotConfiguredAction field value if set, zero value otherwise.
func (o *AuthenticatorValidateStageRequest) GetNotConfiguredAction() NotConfiguredActionEnum {
	if o == nil || o.NotConfiguredAction == nil {
		var ret NotConfiguredActionEnum
		return ret
	}
	return *o.NotConfiguredAction
}

// GetNotConfiguredActionOk returns a tuple with the NotConfiguredAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStageRequest) GetNotConfiguredActionOk() (*NotConfiguredActionEnum, bool) {
	if o == nil || o.NotConfiguredAction == nil {
		return nil, false
	}
	return o.NotConfiguredAction, true
}

// HasNotConfiguredAction returns a boolean if a field has been set.
func (o *AuthenticatorValidateStageRequest) HasNotConfiguredAction() bool {
	if o != nil && o.NotConfiguredAction != nil {
		return true
	}

	return false
}

// SetNotConfiguredAction gets a reference to the given NotConfiguredActionEnum and assigns it to the NotConfiguredAction field.
func (o *AuthenticatorValidateStageRequest) SetNotConfiguredAction(v NotConfiguredActionEnum) {
	o.NotConfiguredAction = &v
}

// GetDeviceClasses returns the DeviceClasses field value if set, zero value otherwise.
func (o *AuthenticatorValidateStageRequest) GetDeviceClasses() []DeviceClassesEnum {
	if o == nil || o.DeviceClasses == nil {
		var ret []DeviceClassesEnum
		return ret
	}
	return o.DeviceClasses
}

// GetDeviceClassesOk returns a tuple with the DeviceClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStageRequest) GetDeviceClassesOk() ([]DeviceClassesEnum, bool) {
	if o == nil || o.DeviceClasses == nil {
		return nil, false
	}
	return o.DeviceClasses, true
}

// HasDeviceClasses returns a boolean if a field has been set.
func (o *AuthenticatorValidateStageRequest) HasDeviceClasses() bool {
	if o != nil && o.DeviceClasses != nil {
		return true
	}

	return false
}

// SetDeviceClasses gets a reference to the given []DeviceClassesEnum and assigns it to the DeviceClasses field.
func (o *AuthenticatorValidateStageRequest) SetDeviceClasses(v []DeviceClassesEnum) {
	o.DeviceClasses = v
}

// GetConfigurationStages returns the ConfigurationStages field value if set, zero value otherwise.
func (o *AuthenticatorValidateStageRequest) GetConfigurationStages() []string {
	if o == nil || o.ConfigurationStages == nil {
		var ret []string
		return ret
	}
	return o.ConfigurationStages
}

// GetConfigurationStagesOk returns a tuple with the ConfigurationStages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStageRequest) GetConfigurationStagesOk() ([]string, bool) {
	if o == nil || o.ConfigurationStages == nil {
		return nil, false
	}
	return o.ConfigurationStages, true
}

// HasConfigurationStages returns a boolean if a field has been set.
func (o *AuthenticatorValidateStageRequest) HasConfigurationStages() bool {
	if o != nil && o.ConfigurationStages != nil {
		return true
	}

	return false
}

// SetConfigurationStages gets a reference to the given []string and assigns it to the ConfigurationStages field.
func (o *AuthenticatorValidateStageRequest) SetConfigurationStages(v []string) {
	o.ConfigurationStages = v
}

// GetLastAuthThreshold returns the LastAuthThreshold field value if set, zero value otherwise.
func (o *AuthenticatorValidateStageRequest) GetLastAuthThreshold() string {
	if o == nil || o.LastAuthThreshold == nil {
		var ret string
		return ret
	}
	return *o.LastAuthThreshold
}

// GetLastAuthThresholdOk returns a tuple with the LastAuthThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStageRequest) GetLastAuthThresholdOk() (*string, bool) {
	if o == nil || o.LastAuthThreshold == nil {
		return nil, false
	}
	return o.LastAuthThreshold, true
}

// HasLastAuthThreshold returns a boolean if a field has been set.
func (o *AuthenticatorValidateStageRequest) HasLastAuthThreshold() bool {
	if o != nil && o.LastAuthThreshold != nil {
		return true
	}

	return false
}

// SetLastAuthThreshold gets a reference to the given string and assigns it to the LastAuthThreshold field.
func (o *AuthenticatorValidateStageRequest) SetLastAuthThreshold(v string) {
	o.LastAuthThreshold = &v
}

// GetWebauthnUserVerification returns the WebauthnUserVerification field value if set, zero value otherwise.
func (o *AuthenticatorValidateStageRequest) GetWebauthnUserVerification() UserVerificationEnum {
	if o == nil || o.WebauthnUserVerification == nil {
		var ret UserVerificationEnum
		return ret
	}
	return *o.WebauthnUserVerification
}

// GetWebauthnUserVerificationOk returns a tuple with the WebauthnUserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorValidateStageRequest) GetWebauthnUserVerificationOk() (*UserVerificationEnum, bool) {
	if o == nil || o.WebauthnUserVerification == nil {
		return nil, false
	}
	return o.WebauthnUserVerification, true
}

// HasWebauthnUserVerification returns a boolean if a field has been set.
func (o *AuthenticatorValidateStageRequest) HasWebauthnUserVerification() bool {
	if o != nil && o.WebauthnUserVerification != nil {
		return true
	}

	return false
}

// SetWebauthnUserVerification gets a reference to the given UserVerificationEnum and assigns it to the WebauthnUserVerification field.
func (o *AuthenticatorValidateStageRequest) SetWebauthnUserVerification(v UserVerificationEnum) {
	o.WebauthnUserVerification = &v
}

func (o AuthenticatorValidateStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorValidateStageRequest struct {
	value *AuthenticatorValidateStageRequest
	isSet bool
}

func (v NullableAuthenticatorValidateStageRequest) Get() *AuthenticatorValidateStageRequest {
	return v.value
}

func (v *NullableAuthenticatorValidateStageRequest) Set(val *AuthenticatorValidateStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorValidateStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorValidateStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorValidateStageRequest(val *AuthenticatorValidateStageRequest) *NullableAuthenticatorValidateStageRequest {
	return &NullableAuthenticatorValidateStageRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorValidateStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorValidateStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
