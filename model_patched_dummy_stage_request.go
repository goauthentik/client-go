/*
authentik

Making authentication simple.

API version: 2024.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedDummyStageRequest DummyStage Serializer
type PatchedDummyStageRequest struct {
	Name       *string          `json:"name,omitempty"`
	FlowSet    []FlowSetRequest `json:"flow_set,omitempty"`
	ThrowError *bool            `json:"throw_error,omitempty"`
}

// NewPatchedDummyStageRequest instantiates a new PatchedDummyStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDummyStageRequest() *PatchedDummyStageRequest {
	this := PatchedDummyStageRequest{}
	return &this
}

// NewPatchedDummyStageRequestWithDefaults instantiates a new PatchedDummyStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDummyStageRequestWithDefaults() *PatchedDummyStageRequest {
	this := PatchedDummyStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedDummyStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDummyStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedDummyStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedDummyStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedDummyStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDummyStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedDummyStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedDummyStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetThrowError returns the ThrowError field value if set, zero value otherwise.
func (o *PatchedDummyStageRequest) GetThrowError() bool {
	if o == nil || o.ThrowError == nil {
		var ret bool
		return ret
	}
	return *o.ThrowError
}

// GetThrowErrorOk returns a tuple with the ThrowError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDummyStageRequest) GetThrowErrorOk() (*bool, bool) {
	if o == nil || o.ThrowError == nil {
		return nil, false
	}
	return o.ThrowError, true
}

// HasThrowError returns a boolean if a field has been set.
func (o *PatchedDummyStageRequest) HasThrowError() bool {
	if o != nil && o.ThrowError != nil {
		return true
	}

	return false
}

// SetThrowError gets a reference to the given bool and assigns it to the ThrowError field.
func (o *PatchedDummyStageRequest) SetThrowError(v bool) {
	o.ThrowError = &v
}

func (o PatchedDummyStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.ThrowError != nil {
		toSerialize["throw_error"] = o.ThrowError
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedDummyStageRequest struct {
	value *PatchedDummyStageRequest
	isSet bool
}

func (v NullablePatchedDummyStageRequest) Get() *PatchedDummyStageRequest {
	return v.value
}

func (v *NullablePatchedDummyStageRequest) Set(val *PatchedDummyStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDummyStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDummyStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDummyStageRequest(val *PatchedDummyStageRequest) *NullablePatchedDummyStageRequest {
	return &NullablePatchedDummyStageRequest{value: val, isSet: true}
}

func (v NullablePatchedDummyStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDummyStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
