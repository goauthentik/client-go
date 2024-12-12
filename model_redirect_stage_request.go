/*
authentik

Making authentication simple.

API version: 2024.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RedirectStageRequest RedirectStage Serializer
type RedirectStageRequest struct {
	Name         string                `json:"name"`
	FlowSet      []FlowSetRequest      `json:"flow_set,omitempty"`
	KeepContext  *bool                 `json:"keep_context,omitempty"`
	Mode         RedirectStageModeEnum `json:"mode"`
	TargetStatic *string               `json:"target_static,omitempty"`
	TargetFlow   NullableString        `json:"target_flow,omitempty"`
}

// NewRedirectStageRequest instantiates a new RedirectStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectStageRequest(name string, mode RedirectStageModeEnum) *RedirectStageRequest {
	this := RedirectStageRequest{}
	this.Name = name
	this.Mode = mode
	return &this
}

// NewRedirectStageRequestWithDefaults instantiates a new RedirectStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectStageRequestWithDefaults() *RedirectStageRequest {
	this := RedirectStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *RedirectStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RedirectStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RedirectStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *RedirectStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *RedirectStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *RedirectStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetKeepContext returns the KeepContext field value if set, zero value otherwise.
func (o *RedirectStageRequest) GetKeepContext() bool {
	if o == nil || o.KeepContext == nil {
		var ret bool
		return ret
	}
	return *o.KeepContext
}

// GetKeepContextOk returns a tuple with the KeepContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectStageRequest) GetKeepContextOk() (*bool, bool) {
	if o == nil || o.KeepContext == nil {
		return nil, false
	}
	return o.KeepContext, true
}

// HasKeepContext returns a boolean if a field has been set.
func (o *RedirectStageRequest) HasKeepContext() bool {
	if o != nil && o.KeepContext != nil {
		return true
	}

	return false
}

// SetKeepContext gets a reference to the given bool and assigns it to the KeepContext field.
func (o *RedirectStageRequest) SetKeepContext(v bool) {
	o.KeepContext = &v
}

// GetMode returns the Mode field value
func (o *RedirectStageRequest) GetMode() RedirectStageModeEnum {
	if o == nil {
		var ret RedirectStageModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *RedirectStageRequest) GetModeOk() (*RedirectStageModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *RedirectStageRequest) SetMode(v RedirectStageModeEnum) {
	o.Mode = v
}

// GetTargetStatic returns the TargetStatic field value if set, zero value otherwise.
func (o *RedirectStageRequest) GetTargetStatic() string {
	if o == nil || o.TargetStatic == nil {
		var ret string
		return ret
	}
	return *o.TargetStatic
}

// GetTargetStaticOk returns a tuple with the TargetStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectStageRequest) GetTargetStaticOk() (*string, bool) {
	if o == nil || o.TargetStatic == nil {
		return nil, false
	}
	return o.TargetStatic, true
}

// HasTargetStatic returns a boolean if a field has been set.
func (o *RedirectStageRequest) HasTargetStatic() bool {
	if o != nil && o.TargetStatic != nil {
		return true
	}

	return false
}

// SetTargetStatic gets a reference to the given string and assigns it to the TargetStatic field.
func (o *RedirectStageRequest) SetTargetStatic(v string) {
	o.TargetStatic = &v
}

// GetTargetFlow returns the TargetFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RedirectStageRequest) GetTargetFlow() string {
	if o == nil || o.TargetFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetFlow.Get()
}

// GetTargetFlowOk returns a tuple with the TargetFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RedirectStageRequest) GetTargetFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetFlow.Get(), o.TargetFlow.IsSet()
}

// HasTargetFlow returns a boolean if a field has been set.
func (o *RedirectStageRequest) HasTargetFlow() bool {
	if o != nil && o.TargetFlow.IsSet() {
		return true
	}

	return false
}

// SetTargetFlow gets a reference to the given NullableString and assigns it to the TargetFlow field.
func (o *RedirectStageRequest) SetTargetFlow(v string) {
	o.TargetFlow.Set(&v)
}

// SetTargetFlowNil sets the value for TargetFlow to be an explicit nil
func (o *RedirectStageRequest) SetTargetFlowNil() {
	o.TargetFlow.Set(nil)
}

// UnsetTargetFlow ensures that no value is present for TargetFlow, not even an explicit nil
func (o *RedirectStageRequest) UnsetTargetFlow() {
	o.TargetFlow.Unset()
}

func (o RedirectStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.KeepContext != nil {
		toSerialize["keep_context"] = o.KeepContext
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.TargetStatic != nil {
		toSerialize["target_static"] = o.TargetStatic
	}
	if o.TargetFlow.IsSet() {
		toSerialize["target_flow"] = o.TargetFlow.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRedirectStageRequest struct {
	value *RedirectStageRequest
	isSet bool
}

func (v NullableRedirectStageRequest) Get() *RedirectStageRequest {
	return v.value
}

func (v *NullableRedirectStageRequest) Set(val *RedirectStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectStageRequest(val *RedirectStageRequest) *NullableRedirectStageRequest {
	return &NullableRedirectStageRequest{value: val, isSet: true}
}

func (v NullableRedirectStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
