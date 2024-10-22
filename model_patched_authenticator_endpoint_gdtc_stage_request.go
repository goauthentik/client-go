/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedAuthenticatorEndpointGDTCStageRequest AuthenticatorEndpointGDTCStage Serializer
type PatchedAuthenticatorEndpointGDTCStageRequest struct {
	Name    *string          `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	FriendlyName  NullableString `json:"friendly_name,omitempty"`
	Credentials   interface{}    `json:"credentials,omitempty"`
}

// NewPatchedAuthenticatorEndpointGDTCStageRequest instantiates a new PatchedAuthenticatorEndpointGDTCStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAuthenticatorEndpointGDTCStageRequest() *PatchedAuthenticatorEndpointGDTCStageRequest {
	this := PatchedAuthenticatorEndpointGDTCStageRequest{}
	return &this
}

// NewPatchedAuthenticatorEndpointGDTCStageRequestWithDefaults instantiates a new PatchedAuthenticatorEndpointGDTCStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAuthenticatorEndpointGDTCStageRequestWithDefaults() *PatchedAuthenticatorEndpointGDTCStageRequest {
	this := PatchedAuthenticatorEndpointGDTCStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetCredentials returns the Credentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) GetCredentials() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) GetCredentialsOk() (*interface{}, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given interface{} and assigns it to the Credentials field.
func (o *PatchedAuthenticatorEndpointGDTCStageRequest) SetCredentials(v interface{}) {
	o.Credentials = v
}

func (o PatchedAuthenticatorEndpointGDTCStageRequest) MarshalJSON() ([]byte, error) {
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
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedAuthenticatorEndpointGDTCStageRequest struct {
	value *PatchedAuthenticatorEndpointGDTCStageRequest
	isSet bool
}

func (v NullablePatchedAuthenticatorEndpointGDTCStageRequest) Get() *PatchedAuthenticatorEndpointGDTCStageRequest {
	return v.value
}

func (v *NullablePatchedAuthenticatorEndpointGDTCStageRequest) Set(val *PatchedAuthenticatorEndpointGDTCStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAuthenticatorEndpointGDTCStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAuthenticatorEndpointGDTCStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAuthenticatorEndpointGDTCStageRequest(val *PatchedAuthenticatorEndpointGDTCStageRequest) *NullablePatchedAuthenticatorEndpointGDTCStageRequest {
	return &NullablePatchedAuthenticatorEndpointGDTCStageRequest{value: val, isSet: true}
}

func (v NullablePatchedAuthenticatorEndpointGDTCStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAuthenticatorEndpointGDTCStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
