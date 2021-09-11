/*
authentik

Making authentication simple.

API version: 2021.8.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedUserDeleteStageRequest UserDeleteStage Serializer
type PatchedUserDeleteStageRequest struct {
	Name *string `json:"name,omitempty"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
}

// NewPatchedUserDeleteStageRequest instantiates a new PatchedUserDeleteStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserDeleteStageRequest() *PatchedUserDeleteStageRequest {
	this := PatchedUserDeleteStageRequest{}
	return &this
}

// NewPatchedUserDeleteStageRequestWithDefaults instantiates a new PatchedUserDeleteStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserDeleteStageRequestWithDefaults() *PatchedUserDeleteStageRequest {
	this := PatchedUserDeleteStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedUserDeleteStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserDeleteStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedUserDeleteStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedUserDeleteStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedUserDeleteStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserDeleteStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedUserDeleteStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *PatchedUserDeleteStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

func (o PatchedUserDeleteStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedUserDeleteStageRequest struct {
	value *PatchedUserDeleteStageRequest
	isSet bool
}

func (v NullablePatchedUserDeleteStageRequest) Get() *PatchedUserDeleteStageRequest {
	return v.value
}

func (v *NullablePatchedUserDeleteStageRequest) Set(val *PatchedUserDeleteStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserDeleteStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserDeleteStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserDeleteStageRequest(val *PatchedUserDeleteStageRequest) *NullablePatchedUserDeleteStageRequest {
	return &NullablePatchedUserDeleteStageRequest{value: val, isSet: true}
}

func (v NullablePatchedUserDeleteStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserDeleteStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


