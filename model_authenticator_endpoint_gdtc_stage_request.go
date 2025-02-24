/*
authentik

Making authentication simple.

API version: 2025.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatorEndpointGDTCStageRequest AuthenticatorEndpointGDTCStage Serializer
type AuthenticatorEndpointGDTCStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	FriendlyName  NullableString `json:"friendly_name,omitempty"`
	Credentials   interface{}    `json:"credentials"`
}

// NewAuthenticatorEndpointGDTCStageRequest instantiates a new AuthenticatorEndpointGDTCStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEndpointGDTCStageRequest(name string, credentials interface{}) *AuthenticatorEndpointGDTCStageRequest {
	this := AuthenticatorEndpointGDTCStageRequest{}
	this.Name = name
	this.Credentials = credentials
	return &this
}

// NewAuthenticatorEndpointGDTCStageRequestWithDefaults instantiates a new AuthenticatorEndpointGDTCStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEndpointGDTCStageRequestWithDefaults() *AuthenticatorEndpointGDTCStageRequest {
	this := AuthenticatorEndpointGDTCStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AuthenticatorEndpointGDTCStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorEndpointGDTCStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticatorEndpointGDTCStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *AuthenticatorEndpointGDTCStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEndpointGDTCStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *AuthenticatorEndpointGDTCStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *AuthenticatorEndpointGDTCStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorEndpointGDTCStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorEndpointGDTCStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *AuthenticatorEndpointGDTCStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *AuthenticatorEndpointGDTCStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *AuthenticatorEndpointGDTCStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *AuthenticatorEndpointGDTCStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticatorEndpointGDTCStageRequest) GetFriendlyName() string {
	if o == nil || o.FriendlyName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorEndpointGDTCStageRequest) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *AuthenticatorEndpointGDTCStageRequest) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *AuthenticatorEndpointGDTCStageRequest) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}

// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *AuthenticatorEndpointGDTCStageRequest) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *AuthenticatorEndpointGDTCStageRequest) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetCredentials returns the Credentials field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AuthenticatorEndpointGDTCStageRequest) GetCredentials() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatorEndpointGDTCStageRequest) GetCredentialsOk() (*interface{}, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *AuthenticatorEndpointGDTCStageRequest) SetCredentials(v interface{}) {
	o.Credentials = v
}

func (o AuthenticatorEndpointGDTCStageRequest) MarshalJSON() ([]byte, error) {
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
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatorEndpointGDTCStageRequest struct {
	value *AuthenticatorEndpointGDTCStageRequest
	isSet bool
}

func (v NullableAuthenticatorEndpointGDTCStageRequest) Get() *AuthenticatorEndpointGDTCStageRequest {
	return v.value
}

func (v *NullableAuthenticatorEndpointGDTCStageRequest) Set(val *AuthenticatorEndpointGDTCStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEndpointGDTCStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEndpointGDTCStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEndpointGDTCStageRequest(val *AuthenticatorEndpointGDTCStageRequest) *NullableAuthenticatorEndpointGDTCStageRequest {
	return &NullableAuthenticatorEndpointGDTCStageRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorEndpointGDTCStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEndpointGDTCStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
