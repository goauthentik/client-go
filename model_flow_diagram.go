/*
authentik

Making authentication simple.

API version: 2023.1.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FlowDiagram response of the flow's diagram action
type FlowDiagram struct {
	Diagram string `json:"diagram"`
}

// NewFlowDiagram instantiates a new FlowDiagram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowDiagram(diagram string) *FlowDiagram {
	this := FlowDiagram{}
	this.Diagram = diagram
	return &this
}

// NewFlowDiagramWithDefaults instantiates a new FlowDiagram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowDiagramWithDefaults() *FlowDiagram {
	this := FlowDiagram{}
	return &this
}

// GetDiagram returns the Diagram field value
func (o *FlowDiagram) GetDiagram() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Diagram
}

// GetDiagramOk returns a tuple with the Diagram field value
// and a boolean to check if the value has been set.
func (o *FlowDiagram) GetDiagramOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Diagram, true
}

// SetDiagram sets field value
func (o *FlowDiagram) SetDiagram(v string) {
	o.Diagram = v
}

func (o FlowDiagram) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["diagram"] = o.Diagram
	}
	return json.Marshal(toSerialize)
}

type NullableFlowDiagram struct {
	value *FlowDiagram
	isSet bool
}

func (v NullableFlowDiagram) Get() *FlowDiagram {
	return v.value
}

func (v *NullableFlowDiagram) Set(val *FlowDiagram) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowDiagram) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowDiagram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowDiagram(val *FlowDiagram) *NullableFlowDiagram {
	return &NullableFlowDiagram{value: val, isSet: true}
}

func (v NullableFlowDiagram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowDiagram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
