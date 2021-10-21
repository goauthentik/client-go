/*
authentik

Making authentication simple.

API version: 2021.10.1-rc2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedDummyStageRequest DummyStage Serializer
type PatchedDummyStageRequest struct {
	Name    *string        `json:"name,omitempty"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
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
func (o *PatchedDummyStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDummyStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
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

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *PatchedDummyStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

func (o PatchedDummyStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
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
