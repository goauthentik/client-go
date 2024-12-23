/*
authentik

Making authentication simple.

API version: 2024.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedRedirectStageRequest RedirectStage Serializer
type PatchedRedirectStageRequest struct {
	Name         *string                `json:"name,omitempty"`
	FlowSet      []FlowSetRequest       `json:"flow_set,omitempty"`
	KeepContext  *bool                  `json:"keep_context,omitempty"`
	Mode         *RedirectStageModeEnum `json:"mode,omitempty"`
	TargetStatic *string                `json:"target_static,omitempty"`
	TargetFlow   NullableString         `json:"target_flow,omitempty"`
}

// NewPatchedRedirectStageRequest instantiates a new PatchedRedirectStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRedirectStageRequest() *PatchedRedirectStageRequest {
	this := PatchedRedirectStageRequest{}
	return &this
}

// NewPatchedRedirectStageRequestWithDefaults instantiates a new PatchedRedirectStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRedirectStageRequestWithDefaults() *PatchedRedirectStageRequest {
	this := PatchedRedirectStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedRedirectStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRedirectStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedRedirectStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedRedirectStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedRedirectStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRedirectStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedRedirectStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedRedirectStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetKeepContext returns the KeepContext field value if set, zero value otherwise.
func (o *PatchedRedirectStageRequest) GetKeepContext() bool {
	if o == nil || o.KeepContext == nil {
		var ret bool
		return ret
	}
	return *o.KeepContext
}

// GetKeepContextOk returns a tuple with the KeepContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRedirectStageRequest) GetKeepContextOk() (*bool, bool) {
	if o == nil || o.KeepContext == nil {
		return nil, false
	}
	return o.KeepContext, true
}

// HasKeepContext returns a boolean if a field has been set.
func (o *PatchedRedirectStageRequest) HasKeepContext() bool {
	if o != nil && o.KeepContext != nil {
		return true
	}

	return false
}

// SetKeepContext gets a reference to the given bool and assigns it to the KeepContext field.
func (o *PatchedRedirectStageRequest) SetKeepContext(v bool) {
	o.KeepContext = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PatchedRedirectStageRequest) GetMode() RedirectStageModeEnum {
	if o == nil || o.Mode == nil {
		var ret RedirectStageModeEnum
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRedirectStageRequest) GetModeOk() (*RedirectStageModeEnum, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PatchedRedirectStageRequest) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given RedirectStageModeEnum and assigns it to the Mode field.
func (o *PatchedRedirectStageRequest) SetMode(v RedirectStageModeEnum) {
	o.Mode = &v
}

// GetTargetStatic returns the TargetStatic field value if set, zero value otherwise.
func (o *PatchedRedirectStageRequest) GetTargetStatic() string {
	if o == nil || o.TargetStatic == nil {
		var ret string
		return ret
	}
	return *o.TargetStatic
}

// GetTargetStaticOk returns a tuple with the TargetStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRedirectStageRequest) GetTargetStaticOk() (*string, bool) {
	if o == nil || o.TargetStatic == nil {
		return nil, false
	}
	return o.TargetStatic, true
}

// HasTargetStatic returns a boolean if a field has been set.
func (o *PatchedRedirectStageRequest) HasTargetStatic() bool {
	if o != nil && o.TargetStatic != nil {
		return true
	}

	return false
}

// SetTargetStatic gets a reference to the given string and assigns it to the TargetStatic field.
func (o *PatchedRedirectStageRequest) SetTargetStatic(v string) {
	o.TargetStatic = &v
}

// GetTargetFlow returns the TargetFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRedirectStageRequest) GetTargetFlow() string {
	if o == nil || o.TargetFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.TargetFlow.Get()
}

// GetTargetFlowOk returns a tuple with the TargetFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRedirectStageRequest) GetTargetFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetFlow.Get(), o.TargetFlow.IsSet()
}

// HasTargetFlow returns a boolean if a field has been set.
func (o *PatchedRedirectStageRequest) HasTargetFlow() bool {
	if o != nil && o.TargetFlow.IsSet() {
		return true
	}

	return false
}

// SetTargetFlow gets a reference to the given NullableString and assigns it to the TargetFlow field.
func (o *PatchedRedirectStageRequest) SetTargetFlow(v string) {
	o.TargetFlow.Set(&v)
}

// SetTargetFlowNil sets the value for TargetFlow to be an explicit nil
func (o *PatchedRedirectStageRequest) SetTargetFlowNil() {
	o.TargetFlow.Set(nil)
}

// UnsetTargetFlow ensures that no value is present for TargetFlow, not even an explicit nil
func (o *PatchedRedirectStageRequest) UnsetTargetFlow() {
	o.TargetFlow.Unset()
}

func (o PatchedRedirectStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.KeepContext != nil {
		toSerialize["keep_context"] = o.KeepContext
	}
	if o.Mode != nil {
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

type NullablePatchedRedirectStageRequest struct {
	value *PatchedRedirectStageRequest
	isSet bool
}

func (v NullablePatchedRedirectStageRequest) Get() *PatchedRedirectStageRequest {
	return v.value
}

func (v *NullablePatchedRedirectStageRequest) Set(val *PatchedRedirectStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRedirectStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRedirectStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRedirectStageRequest(val *PatchedRedirectStageRequest) *NullablePatchedRedirectStageRequest {
	return &NullablePatchedRedirectStageRequest{value: val, isSet: true}
}

func (v NullablePatchedRedirectStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRedirectStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
