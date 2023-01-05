/*
authentik

Making authentication simple.

API version: 2022.12.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// StageRequest Stage Serializer
type StageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
}

// NewStageRequest instantiates a new StageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStageRequest(name string) *StageRequest {
	this := StageRequest{}
	this.Name = name
	return &this
}

// NewStageRequestWithDefaults instantiates a new StageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStageRequestWithDefaults() *StageRequest {
	this := StageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *StageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *StageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *StageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *StageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

func (o StageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	return json.Marshal(toSerialize)
}

type NullableStageRequest struct {
	value *StageRequest
	isSet bool
}

func (v NullableStageRequest) Get() *StageRequest {
	return v.value
}

func (v *NullableStageRequest) Set(val *StageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStageRequest(val *StageRequest) *NullableStageRequest {
	return &NullableStageRequest{value: val, isSet: true}
}

func (v NullableStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
