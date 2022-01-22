/*
authentik

Making authentication simple.

API version: 2022.1.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedDenyStageRequest DenyStage Serializer
type PatchedDenyStageRequest struct {
	Name    *string        `json:"name,omitempty"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
}

// NewPatchedDenyStageRequest instantiates a new PatchedDenyStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDenyStageRequest() *PatchedDenyStageRequest {
	this := PatchedDenyStageRequest{}
	return &this
}

// NewPatchedDenyStageRequestWithDefaults instantiates a new PatchedDenyStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDenyStageRequestWithDefaults() *PatchedDenyStageRequest {
	this := PatchedDenyStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedDenyStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDenyStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedDenyStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedDenyStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedDenyStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDenyStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedDenyStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *PatchedDenyStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

func (o PatchedDenyStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedDenyStageRequest struct {
	value *PatchedDenyStageRequest
	isSet bool
}

func (v NullablePatchedDenyStageRequest) Get() *PatchedDenyStageRequest {
	return v.value
}

func (v *NullablePatchedDenyStageRequest) Set(val *PatchedDenyStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDenyStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDenyStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDenyStageRequest(val *PatchedDenyStageRequest) *NullablePatchedDenyStageRequest {
	return &NullablePatchedDenyStageRequest{value: val, isSet: true}
}

func (v NullablePatchedDenyStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDenyStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
