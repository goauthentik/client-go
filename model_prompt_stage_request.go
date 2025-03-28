/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PromptStageRequest PromptStage Serializer
type PromptStageRequest struct {
	Name               string           `json:"name"`
	FlowSet            []FlowSetRequest `json:"flow_set,omitempty"`
	Fields             []string         `json:"fields"`
	ValidationPolicies []string         `json:"validation_policies,omitempty"`
}

// NewPromptStageRequest instantiates a new PromptStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptStageRequest(name string, fields []string) *PromptStageRequest {
	this := PromptStageRequest{}
	this.Name = name
	this.Fields = fields
	return &this
}

// NewPromptStageRequestWithDefaults instantiates a new PromptStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptStageRequestWithDefaults() *PromptStageRequest {
	this := PromptStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PromptStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PromptStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PromptStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PromptStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PromptStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PromptStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetFields returns the Fields field value
func (o *PromptStageRequest) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *PromptStageRequest) GetFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *PromptStageRequest) SetFields(v []string) {
	o.Fields = v
}

// GetValidationPolicies returns the ValidationPolicies field value if set, zero value otherwise.
func (o *PromptStageRequest) GetValidationPolicies() []string {
	if o == nil || o.ValidationPolicies == nil {
		var ret []string
		return ret
	}
	return o.ValidationPolicies
}

// GetValidationPoliciesOk returns a tuple with the ValidationPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptStageRequest) GetValidationPoliciesOk() ([]string, bool) {
	if o == nil || o.ValidationPolicies == nil {
		return nil, false
	}
	return o.ValidationPolicies, true
}

// HasValidationPolicies returns a boolean if a field has been set.
func (o *PromptStageRequest) HasValidationPolicies() bool {
	if o != nil && o.ValidationPolicies != nil {
		return true
	}

	return false
}

// SetValidationPolicies gets a reference to the given []string and assigns it to the ValidationPolicies field.
func (o *PromptStageRequest) SetValidationPolicies(v []string) {
	o.ValidationPolicies = v
}

func (o PromptStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if true {
		toSerialize["fields"] = o.Fields
	}
	if o.ValidationPolicies != nil {
		toSerialize["validation_policies"] = o.ValidationPolicies
	}
	return json.Marshal(toSerialize)
}

type NullablePromptStageRequest struct {
	value *PromptStageRequest
	isSet bool
}

func (v NullablePromptStageRequest) Get() *PromptStageRequest {
	return v.value
}

func (v *NullablePromptStageRequest) Set(val *PromptStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptStageRequest(val *PromptStageRequest) *NullablePromptStageRequest {
	return &NullablePromptStageRequest{value: val, isSet: true}
}

func (v NullablePromptStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
