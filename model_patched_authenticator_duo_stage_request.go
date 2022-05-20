/*
authentik

Making authentication simple.

API version: 2022.5.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedAuthenticatorDuoStageRequest AuthenticatorDuoStage Serializer
type PatchedAuthenticatorDuoStageRequest struct {
	Name    *string        `json:"name,omitempty"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	ClientId      *string        `json:"client_id,omitempty"`
	ClientSecret  *string        `json:"client_secret,omitempty"`
	ApiHostname   *string        `json:"api_hostname,omitempty"`
}

// NewPatchedAuthenticatorDuoStageRequest instantiates a new PatchedAuthenticatorDuoStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedAuthenticatorDuoStageRequest() *PatchedAuthenticatorDuoStageRequest {
	this := PatchedAuthenticatorDuoStageRequest{}
	return &this
}

// NewPatchedAuthenticatorDuoStageRequestWithDefaults instantiates a new PatchedAuthenticatorDuoStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedAuthenticatorDuoStageRequestWithDefaults() *PatchedAuthenticatorDuoStageRequest {
	this := PatchedAuthenticatorDuoStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedAuthenticatorDuoStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorDuoStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedAuthenticatorDuoStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedAuthenticatorDuoStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedAuthenticatorDuoStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorDuoStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedAuthenticatorDuoStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *PatchedAuthenticatorDuoStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedAuthenticatorDuoStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedAuthenticatorDuoStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *PatchedAuthenticatorDuoStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *PatchedAuthenticatorDuoStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *PatchedAuthenticatorDuoStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *PatchedAuthenticatorDuoStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PatchedAuthenticatorDuoStageRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorDuoStageRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PatchedAuthenticatorDuoStageRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PatchedAuthenticatorDuoStageRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *PatchedAuthenticatorDuoStageRequest) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorDuoStageRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *PatchedAuthenticatorDuoStageRequest) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *PatchedAuthenticatorDuoStageRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetApiHostname returns the ApiHostname field value if set, zero value otherwise.
func (o *PatchedAuthenticatorDuoStageRequest) GetApiHostname() string {
	if o == nil || o.ApiHostname == nil {
		var ret string
		return ret
	}
	return *o.ApiHostname
}

// GetApiHostnameOk returns a tuple with the ApiHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedAuthenticatorDuoStageRequest) GetApiHostnameOk() (*string, bool) {
	if o == nil || o.ApiHostname == nil {
		return nil, false
	}
	return o.ApiHostname, true
}

// HasApiHostname returns a boolean if a field has been set.
func (o *PatchedAuthenticatorDuoStageRequest) HasApiHostname() bool {
	if o != nil && o.ApiHostname != nil {
		return true
	}

	return false
}

// SetApiHostname gets a reference to the given string and assigns it to the ApiHostname field.
func (o *PatchedAuthenticatorDuoStageRequest) SetApiHostname(v string) {
	o.ApiHostname = &v
}

func (o PatchedAuthenticatorDuoStageRequest) MarshalJSON() ([]byte, error) {
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
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if o.ApiHostname != nil {
		toSerialize["api_hostname"] = o.ApiHostname
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedAuthenticatorDuoStageRequest struct {
	value *PatchedAuthenticatorDuoStageRequest
	isSet bool
}

func (v NullablePatchedAuthenticatorDuoStageRequest) Get() *PatchedAuthenticatorDuoStageRequest {
	return v.value
}

func (v *NullablePatchedAuthenticatorDuoStageRequest) Set(val *PatchedAuthenticatorDuoStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedAuthenticatorDuoStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedAuthenticatorDuoStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedAuthenticatorDuoStageRequest(val *PatchedAuthenticatorDuoStageRequest) *NullablePatchedAuthenticatorDuoStageRequest {
	return &NullablePatchedAuthenticatorDuoStageRequest{value: val, isSet: true}
}

func (v NullablePatchedAuthenticatorDuoStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedAuthenticatorDuoStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
