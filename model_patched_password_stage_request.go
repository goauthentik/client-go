/*
authentik

Making authentication simple.

API version: 2023.3.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedPasswordStageRequest PasswordStage Serializer
type PatchedPasswordStageRequest struct {
	Name    *string          `json:"name,omitempty"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Selection of backends to test the password against.
	Backends []BackendsEnum `json:"backends,omitempty"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	// How many attempts a user has before the flow is canceled. To lock the user out, use a reputation policy and a user_write stage.
	FailedAttemptsBeforeCancel *int32 `json:"failed_attempts_before_cancel,omitempty"`
}

// NewPatchedPasswordStageRequest instantiates a new PatchedPasswordStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPasswordStageRequest() *PatchedPasswordStageRequest {
	this := PatchedPasswordStageRequest{}
	return &this
}

// NewPatchedPasswordStageRequestWithDefaults instantiates a new PatchedPasswordStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPasswordStageRequestWithDefaults() *PatchedPasswordStageRequest {
	this := PatchedPasswordStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedPasswordStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPasswordStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedPasswordStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedPasswordStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedPasswordStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedPasswordStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetBackends returns the Backends field value if set, zero value otherwise.
func (o *PatchedPasswordStageRequest) GetBackends() []BackendsEnum {
	if o == nil || o.Backends == nil {
		var ret []BackendsEnum
		return ret
	}
	return o.Backends
}

// GetBackendsOk returns a tuple with the Backends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordStageRequest) GetBackendsOk() ([]BackendsEnum, bool) {
	if o == nil || o.Backends == nil {
		return nil, false
	}
	return o.Backends, true
}

// HasBackends returns a boolean if a field has been set.
func (o *PatchedPasswordStageRequest) HasBackends() bool {
	if o != nil && o.Backends != nil {
		return true
	}

	return false
}

// SetBackends gets a reference to the given []BackendsEnum and assigns it to the Backends field.
func (o *PatchedPasswordStageRequest) SetBackends(v []BackendsEnum) {
	o.Backends = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedPasswordStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedPasswordStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *PatchedPasswordStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *PatchedPasswordStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *PatchedPasswordStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *PatchedPasswordStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFailedAttemptsBeforeCancel returns the FailedAttemptsBeforeCancel field value if set, zero value otherwise.
func (o *PatchedPasswordStageRequest) GetFailedAttemptsBeforeCancel() int32 {
	if o == nil || o.FailedAttemptsBeforeCancel == nil {
		var ret int32
		return ret
	}
	return *o.FailedAttemptsBeforeCancel
}

// GetFailedAttemptsBeforeCancelOk returns a tuple with the FailedAttemptsBeforeCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPasswordStageRequest) GetFailedAttemptsBeforeCancelOk() (*int32, bool) {
	if o == nil || o.FailedAttemptsBeforeCancel == nil {
		return nil, false
	}
	return o.FailedAttemptsBeforeCancel, true
}

// HasFailedAttemptsBeforeCancel returns a boolean if a field has been set.
func (o *PatchedPasswordStageRequest) HasFailedAttemptsBeforeCancel() bool {
	if o != nil && o.FailedAttemptsBeforeCancel != nil {
		return true
	}

	return false
}

// SetFailedAttemptsBeforeCancel gets a reference to the given int32 and assigns it to the FailedAttemptsBeforeCancel field.
func (o *PatchedPasswordStageRequest) SetFailedAttemptsBeforeCancel(v int32) {
	o.FailedAttemptsBeforeCancel = &v
}

func (o PatchedPasswordStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.Backends != nil {
		toSerialize["backends"] = o.Backends
	}
	if o.ConfigureFlow.IsSet() {
		toSerialize["configure_flow"] = o.ConfigureFlow.Get()
	}
	if o.FailedAttemptsBeforeCancel != nil {
		toSerialize["failed_attempts_before_cancel"] = o.FailedAttemptsBeforeCancel
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedPasswordStageRequest struct {
	value *PatchedPasswordStageRequest
	isSet bool
}

func (v NullablePatchedPasswordStageRequest) Get() *PatchedPasswordStageRequest {
	return v.value
}

func (v *NullablePatchedPasswordStageRequest) Set(val *PatchedPasswordStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPasswordStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPasswordStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPasswordStageRequest(val *PatchedPasswordStageRequest) *NullablePatchedPasswordStageRequest {
	return &NullablePatchedPasswordStageRequest{value: val, isSet: true}
}

func (v NullablePatchedPasswordStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPasswordStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
