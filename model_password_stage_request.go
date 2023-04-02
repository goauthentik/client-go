/*
authentik

Making authentication simple.

API version: 2023.3.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PasswordStageRequest PasswordStage Serializer
type PasswordStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
	// Selection of backends to test the password against.
	Backends []BackendsEnum `json:"backends"`
	// Flow used by an authenticated user to configure this Stage. If empty, user will not be able to configure this stage.
	ConfigureFlow NullableString `json:"configure_flow,omitempty"`
	// How many attempts a user has before the flow is canceled. To lock the user out, use a reputation policy and a user_write stage.
	FailedAttemptsBeforeCancel *int32 `json:"failed_attempts_before_cancel,omitempty"`
}

// NewPasswordStageRequest instantiates a new PasswordStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordStageRequest(name string, backends []BackendsEnum) *PasswordStageRequest {
	this := PasswordStageRequest{}
	this.Name = name
	this.Backends = backends
	return &this
}

// NewPasswordStageRequestWithDefaults instantiates a new PasswordStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordStageRequestWithDefaults() *PasswordStageRequest {
	this := PasswordStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PasswordStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PasswordStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PasswordStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PasswordStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PasswordStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PasswordStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetBackends returns the Backends field value
func (o *PasswordStageRequest) GetBackends() []BackendsEnum {
	if o == nil {
		var ret []BackendsEnum
		return ret
	}

	return o.Backends
}

// GetBackendsOk returns a tuple with the Backends field value
// and a boolean to check if the value has been set.
func (o *PasswordStageRequest) GetBackendsOk() ([]BackendsEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Backends, true
}

// SetBackends sets field value
func (o *PasswordStageRequest) SetBackends(v []BackendsEnum) {
	o.Backends = v
}

// GetConfigureFlow returns the ConfigureFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordStageRequest) GetConfigureFlow() string {
	if o == nil || o.ConfigureFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfigureFlow.Get()
}

// GetConfigureFlowOk returns a tuple with the ConfigureFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordStageRequest) GetConfigureFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigureFlow.Get(), o.ConfigureFlow.IsSet()
}

// HasConfigureFlow returns a boolean if a field has been set.
func (o *PasswordStageRequest) HasConfigureFlow() bool {
	if o != nil && o.ConfigureFlow.IsSet() {
		return true
	}

	return false
}

// SetConfigureFlow gets a reference to the given NullableString and assigns it to the ConfigureFlow field.
func (o *PasswordStageRequest) SetConfigureFlow(v string) {
	o.ConfigureFlow.Set(&v)
}

// SetConfigureFlowNil sets the value for ConfigureFlow to be an explicit nil
func (o *PasswordStageRequest) SetConfigureFlowNil() {
	o.ConfigureFlow.Set(nil)
}

// UnsetConfigureFlow ensures that no value is present for ConfigureFlow, not even an explicit nil
func (o *PasswordStageRequest) UnsetConfigureFlow() {
	o.ConfigureFlow.Unset()
}

// GetFailedAttemptsBeforeCancel returns the FailedAttemptsBeforeCancel field value if set, zero value otherwise.
func (o *PasswordStageRequest) GetFailedAttemptsBeforeCancel() int32 {
	if o == nil || o.FailedAttemptsBeforeCancel == nil {
		var ret int32
		return ret
	}
	return *o.FailedAttemptsBeforeCancel
}

// GetFailedAttemptsBeforeCancelOk returns a tuple with the FailedAttemptsBeforeCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordStageRequest) GetFailedAttemptsBeforeCancelOk() (*int32, bool) {
	if o == nil || o.FailedAttemptsBeforeCancel == nil {
		return nil, false
	}
	return o.FailedAttemptsBeforeCancel, true
}

// HasFailedAttemptsBeforeCancel returns a boolean if a field has been set.
func (o *PasswordStageRequest) HasFailedAttemptsBeforeCancel() bool {
	if o != nil && o.FailedAttemptsBeforeCancel != nil {
		return true
	}

	return false
}

// SetFailedAttemptsBeforeCancel gets a reference to the given int32 and assigns it to the FailedAttemptsBeforeCancel field.
func (o *PasswordStageRequest) SetFailedAttemptsBeforeCancel(v int32) {
	o.FailedAttemptsBeforeCancel = &v
}

func (o PasswordStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if true {
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

type NullablePasswordStageRequest struct {
	value *PasswordStageRequest
	isSet bool
}

func (v NullablePasswordStageRequest) Get() *PasswordStageRequest {
	return v.value
}

func (v *NullablePasswordStageRequest) Set(val *PasswordStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordStageRequest(val *PasswordStageRequest) *NullablePasswordStageRequest {
	return &NullablePasswordStageRequest{value: val, isSet: true}
}

func (v NullablePasswordStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
