/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DummyStageRequest DummyStage Serializer
type DummyStageRequest struct {
	Name       string           `json:"name"`
	FlowSet    []FlowSetRequest `json:"flow_set,omitempty"`
	ThrowError *bool            `json:"throw_error,omitempty"`
}

// NewDummyStageRequest instantiates a new DummyStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDummyStageRequest(name string) *DummyStageRequest {
	this := DummyStageRequest{}
	this.Name = name
	return &this
}

// NewDummyStageRequestWithDefaults instantiates a new DummyStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDummyStageRequestWithDefaults() *DummyStageRequest {
	this := DummyStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DummyStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DummyStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DummyStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *DummyStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *DummyStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *DummyStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetThrowError returns the ThrowError field value if set, zero value otherwise.
func (o *DummyStageRequest) GetThrowError() bool {
	if o == nil || o.ThrowError == nil {
		var ret bool
		return ret
	}
	return *o.ThrowError
}

// GetThrowErrorOk returns a tuple with the ThrowError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DummyStageRequest) GetThrowErrorOk() (*bool, bool) {
	if o == nil || o.ThrowError == nil {
		return nil, false
	}
	return o.ThrowError, true
}

// HasThrowError returns a boolean if a field has been set.
func (o *DummyStageRequest) HasThrowError() bool {
	if o != nil && o.ThrowError != nil {
		return true
	}

	return false
}

// SetThrowError gets a reference to the given bool and assigns it to the ThrowError field.
func (o *DummyStageRequest) SetThrowError(v bool) {
	o.ThrowError = &v
}

func (o DummyStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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

type NullableDummyStageRequest struct {
	value *DummyStageRequest
	isSet bool
}

func (v NullableDummyStageRequest) Get() *DummyStageRequest {
	return v.value
}

func (v *NullableDummyStageRequest) Set(val *DummyStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDummyStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDummyStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDummyStageRequest(val *DummyStageRequest) *NullableDummyStageRequest {
	return &NullableDummyStageRequest{value: val, isSet: true}
}

func (v NullableDummyStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDummyStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
