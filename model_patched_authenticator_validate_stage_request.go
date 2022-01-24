/*
authentik

Making authentication simple.

API version: 2022.1.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedAuthenticatorValidateStageRequest AuthenticatorValidateStage Serializer
type PatchedAuthenticatorValidateStageRequest struct {
	Name                *string                  `json:"name,omitempty"`
	FlowSet             *[]FlowRequest           `json:"flow_set,omitempty"`
	NotConfiguredAction *NotConfiguredActionEnum `json:"not_configured_action,omitempty"`
	// Device classes which can be used to authenticate
	DeviceClasses *[]DeviceClassesEnum `json:"device_classes,omitempty"`
	// Stage used to configure Authenticator when user doesn't have any compatible devices. After this configuration Stage passes, the user is not prompted again.
	ConfigurationStage NullableString `json:"configuration_stage,omitempty"`
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
func (o *PatchedAuthenticatorValidateStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorValidateStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
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

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *PatchedAuthenticatorValidateStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
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
	return *o.DeviceClasses
}

// GetDeviceClassesOk returns a tuple with the DeviceClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorValidateStageRequest) GetDeviceClassesOk() (*[]DeviceClassesEnum, bool) {
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
	o.DeviceClasses = &v
}

// GetConfigurationStage returns the ConfigurationStage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorValidateStageRequest) GetConfigurationStage() string {
	if o == nil || o.ConfigurationStage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationStage.Get()
}

// GetConfigurationStageOk returns a tuple with the ConfigurationStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorValidateStageRequest) GetConfigurationStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationStage.Get(), o.ConfigurationStage.IsSet()
}

// HasConfigurationStage returns a boolean if a field has been set.
func (o *PatchedAuthenticatorValidateStageRequest) HasConfigurationStage() bool {
	if o != nil && o.ConfigurationStage.IsSet() {
		return true
	}

	return false
}

// SetConfigurationStage gets a reference to the given NullableString and assigns it to the ConfigurationStage field.
func (o *PatchedAuthenticatorValidateStageRequest) SetConfigurationStage(v string) {
	o.ConfigurationStage.Set(&v)
}

// SetConfigurationStageNil sets the value for ConfigurationStage to be an explicit nil
func (o *PatchedAuthenticatorValidateStageRequest) SetConfigurationStageNil() {
	o.ConfigurationStage.Set(nil)
}

// UnsetConfigurationStage ensures that no value is present for ConfigurationStage, not even an explicit nil
func (o *PatchedAuthenticatorValidateStageRequest) UnsetConfigurationStage() {
	o.ConfigurationStage.Unset()
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
	if o.ConfigurationStage.IsSet() {
		toSerialize["configuration_stage"] = o.ConfigurationStage.Get()
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
