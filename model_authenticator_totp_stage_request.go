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

// AuthenticatorTOTPStageRequest AuthenticatorTOTPStage Serializer
type AuthenticatorTOTPStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	FriendlyName  NullableString `json:"friendly_name,omitempty"`
	Digits        DigitsEnum     `json:"digits"`
}

// NewAuthenticatorTOTPStageRequest instantiates a new AuthenticatorTOTPStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorTOTPStageRequest(name string, digits DigitsEnum) *AuthenticatorTOTPStageRequest {
	this := AuthenticatorTOTPStageRequest{}
	this.Name = name
	this.Digits = digits
	return &this
}

// NewAuthenticatorTOTPStageRequestWithDefaults instantiates a new AuthenticatorTOTPStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorTOTPStageRequestWithDefaults() *AuthenticatorTOTPStageRequest {
	this := AuthenticatorTOTPStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AuthenticatorTOTPStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorTOTPStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorTOTPStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorTOTPStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorTOTPStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorTOTPStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *AuthenticatorTOTPStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorTOTPStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorTOTPStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *AuthenticatorTOTPStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *AuthenticatorTOTPStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *AuthenticatorTOTPStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *AuthenticatorTOTPStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorTOTPStageRequest) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorTOTPStageRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *AuthenticatorTOTPStageRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *AuthenticatorTOTPStageRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *AuthenticatorTOTPStageRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *AuthenticatorTOTPStageRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetDigits returns the Digits field value
func (o *AuthenticatorTOTPStageRequest) GetDigits() DigitsEnum {
	if o == nil {
		var ret DigitsEnum
		return ret
	}

	return o.Digits
}

// GetDigitsOk returns a tuple with the Digits field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorTOTPStageRequest) GetDigitsOk() (*DigitsEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digits, true
}

// SetDigits sets field value
func (o *AuthenticatorTOTPStageRequest) SetDigits(v DigitsEnum) {
	o.Digits = v
}

func (o AuthenticatorTOTPStageRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["digits"] = o.Digits
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorTOTPStageRequest struct {
	value *AuthenticatorTOTPStageRequest
	isSet bool
}

func (v NullableAuthenticatorTOTPStageRequest) Get() *AuthenticatorTOTPStageRequest {
	return v.value
}

func (v *NullableAuthenticatorTOTPStageRequest) Set(val *AuthenticatorTOTPStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorTOTPStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorTOTPStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorTOTPStageRequest(val *AuthenticatorTOTPStageRequest) *NullableAuthenticatorTOTPStageRequest {
	return &NullableAuthenticatorTOTPStageRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorTOTPStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorTOTPStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
