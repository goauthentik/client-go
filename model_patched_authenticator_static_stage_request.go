/*
authentik

Making authentication simple.

API version: 2023.5.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedAuthenticatorStaticStageRequest AuthenticatorStaticStage Serializer
type PatchedAuthenticatorStaticStageRequest struct {
	Name    *string          `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	FriendlyName  NullableString `json:"friendly_name,omitempty"`
	TokenCount    *int32         `json:"token_count,omitempty"`
}

// NewPatchedAuthenticatorStaticStageRequest instantiates a new PatchedAuthenticatorStaticStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAuthenticatorStaticStageRequest() *PatchedAuthenticatorStaticStageRequest {
	this := PatchedAuthenticatorStaticStageRequest{}
	return &this
}

// NewPatchedAuthenticatorStaticStageRequestWithDefaults instantiates a new PatchedAuthenticatorStaticStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAuthenticatorStaticStageRequestWithDefaults() *PatchedAuthenticatorStaticStageRequest {
	this := PatchedAuthenticatorStaticStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAuthenticatorStaticStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorStaticStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAuthenticatorStaticStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAuthenticatorStaticStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedAuthenticatorStaticStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorStaticStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedAuthenticatorStaticStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedAuthenticatorStaticStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorStaticStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorStaticStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *PatchedAuthenticatorStaticStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *PatchedAuthenticatorStaticStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *PatchedAuthenticatorStaticStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *PatchedAuthenticatorStaticStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorStaticStageRequest) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorStaticStageRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *PatchedAuthenticatorStaticStageRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *PatchedAuthenticatorStaticStageRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *PatchedAuthenticatorStaticStageRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *PatchedAuthenticatorStaticStageRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetTokenCount returns the TokenCount field value if set, zero value otherwise.
func (o *PatchedAuthenticatorStaticStageRequest) GetTokenCount() int32 {
	if o == nil || o.TokenCount == nil {
		var ret int32
		return ret
	}
	return *o.TokenCount
}

// GetTokenCountOk returns a tuple with the TokenCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorStaticStageRequest) GetTokenCountOk() (*int32, bool) {
	if o == nil || o.TokenCount == nil {
		return nil, false
	}
	return o.TokenCount, true
}

// HasTokenCount returns a boolean if a field has been set.
func (o *PatchedAuthenticatorStaticStageRequest) HasTokenCount() bool {
	if o != nil && o.TokenCount != nil {
		return true
	}

	return false
}

// SetTokenCount gets a reference to the given int32 and assigns it to the TokenCount field.
func (o *PatchedAuthenticatorStaticStageRequest) SetTokenCount(v int32) {
	o.TokenCount = &v
}

func (o PatchedAuthenticatorStaticStageRequest) MarshalJSON() ([]byte, error) {
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
	if o.TokenCount != nil {
		toSerialize["token_count"] = o.TokenCount
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedAuthenticatorStaticStageRequest struct {
	value *PatchedAuthenticatorStaticStageRequest
	isSet bool
}

func (v NullablePatchedAuthenticatorStaticStageRequest) Get() *PatchedAuthenticatorStaticStageRequest {
	return v.value
}

func (v *NullablePatchedAuthenticatorStaticStageRequest) Set(val *PatchedAuthenticatorStaticStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAuthenticatorStaticStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAuthenticatorStaticStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAuthenticatorStaticStageRequest(val *PatchedAuthenticatorStaticStageRequest) *NullablePatchedAuthenticatorStaticStageRequest {
	return &NullablePatchedAuthenticatorStaticStageRequest{value: val, isSet: true}
}

func (v NullablePatchedAuthenticatorStaticStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAuthenticatorStaticStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
