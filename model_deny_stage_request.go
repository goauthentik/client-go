/*
authentik

Making authentication simple.

API version: 2023.1.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DenyStageRequest DenyStage Serializer
type DenyStageRequest struct {
	Name    string           `json:"name"`
	FlowSet []FlowSetRequest `json:"flow_set,omitempty"`
}

// NewDenyStageRequest instantiates a new DenyStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDenyStageRequest(name string) *DenyStageRequest {
	this := DenyStageRequest{}
	this.Name = name
	return &this
}

// NewDenyStageRequestWithDefaults instantiates a new DenyStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDenyStageRequestWithDefaults() *DenyStageRequest {
	this := DenyStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DenyStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DenyStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DenyStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *DenyStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DenyStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *DenyStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *DenyStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

func (o DenyStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	return json.Marshal(toSerialize)
}

type NullableDenyStageRequest struct {
	value *DenyStageRequest
	isSet bool
}

func (v NullableDenyStageRequest) Get() *DenyStageRequest {
	return v.value
}

func (v *NullableDenyStageRequest) Set(val *DenyStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDenyStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDenyStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDenyStageRequest(val *DenyStageRequest) *NullableDenyStageRequest {
	return &NullableDenyStageRequest{value: val, isSet: true}
}

func (v NullableDenyStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDenyStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
