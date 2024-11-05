/*
authentik

Making authentication simple.

API version: 2024.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedAuthenticatorTOTPStageRequest AuthenticatorTOTPStage Serializer
type PatchedAuthenticatorTOTPStageRequest struct {
	Name    *string          `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	FriendlyName  NullableString `json:"friendly_name,omitempty"`
	Digits        *DigitsEnum    `json:"digits,omitempty"`
}

// NewPatchedAuthenticatorTOTPStageRequest instantiates a new PatchedAuthenticatorTOTPStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAuthenticatorTOTPStageRequest() *PatchedAuthenticatorTOTPStageRequest {
	this := PatchedAuthenticatorTOTPStageRequest{}
	return &this
}

// NewPatchedAuthenticatorTOTPStageRequestWithDefaults instantiates a new PatchedAuthenticatorTOTPStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAuthenticatorTOTPStageRequestWithDefaults() *PatchedAuthenticatorTOTPStageRequest {
	this := PatchedAuthenticatorTOTPStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAuthenticatorTOTPStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorTOTPStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAuthenticatorTOTPStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAuthenticatorTOTPStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedAuthenticatorTOTPStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorTOTPStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedAuthenticatorTOTPStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedAuthenticatorTOTPStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorTOTPStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorTOTPStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *PatchedAuthenticatorTOTPStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *PatchedAuthenticatorTOTPStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *PatchedAuthenticatorTOTPStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *PatchedAuthenticatorTOTPStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorTOTPStageRequest) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorTOTPStageRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *PatchedAuthenticatorTOTPStageRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *PatchedAuthenticatorTOTPStageRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *PatchedAuthenticatorTOTPStageRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *PatchedAuthenticatorTOTPStageRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetDigits returns the Digits field value if set, zero value otherwise.
func (o *PatchedAuthenticatorTOTPStageRequest) GetDigits() DigitsEnum {
	if o == nil || o.Digits == nil {
		var ret DigitsEnum
		return ret
	}
	return *o.Digits
}

// GetDigitsOk returns a tuple with the Digits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorTOTPStageRequest) GetDigitsOk() (*DigitsEnum, bool) {
	if o == nil || o.Digits == nil {
		return nil, false
	}
	return o.Digits, true
}

// HasDigits returns a boolean if a field has been set.
func (o *PatchedAuthenticatorTOTPStageRequest) HasDigits() bool {
	if o != nil && o.Digits != nil {
		return true
	}

	return false
}

// SetDigits gets a reference to the given DigitsEnum and assigns it to the Digits field.
func (o *PatchedAuthenticatorTOTPStageRequest) SetDigits(v DigitsEnum) {
	o.Digits = &v
}

func (o PatchedAuthenticatorTOTPStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
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
	if o.Digits != nil {
		toSerialize["digits"] = o.Digits
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedAuthenticatorTOTPStageRequest struct {
	value *PatchedAuthenticatorTOTPStageRequest
	isSet bool
}

func (v NullablePatchedAuthenticatorTOTPStageRequest) Get() *PatchedAuthenticatorTOTPStageRequest {
	return v.value
}

func (v *NullablePatchedAuthenticatorTOTPStageRequest) Set(val *PatchedAuthenticatorTOTPStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAuthenticatorTOTPStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAuthenticatorTOTPStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAuthenticatorTOTPStageRequest(val *PatchedAuthenticatorTOTPStageRequest) *NullablePatchedAuthenticatorTOTPStageRequest {
	return &NullablePatchedAuthenticatorTOTPStageRequest{value: val, isSet: true}
}

func (v NullablePatchedAuthenticatorTOTPStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAuthenticatorTOTPStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
