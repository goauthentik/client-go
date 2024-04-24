/*
authentik

Making authentication simple.

API version: 2024.4.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedPromptStageRequest PromptStage Serializer
type PatchedPromptStageRequest struct {
	Name               *string          `json:"name,omitempty"`
	FlowSet            []FlowSetRequest `json:"flow_set,omitempty"`
	Fields             []string         `json:"fields,omitempty"`
	ValidationPolicies []string         `json:"validation_policies,omitempty"`
}

// NewPatchedPromptStageRequest instantiates a new PatchedPromptStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedPromptStageRequest() *PatchedPromptStageRequest {
	this := PatchedPromptStageRequest{}
	return &this
}

// NewPatchedPromptStageRequestWithDefaults instantiates a new PatchedPromptStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedPromptStageRequestWithDefaults() *PatchedPromptStageRequest {
	this := PatchedPromptStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedPromptStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedPromptStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedPromptStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedPromptStageRequest) GetFlowSet() []FlowSetRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowSetRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptStageRequest) GetFlowSetOk() ([]FlowSetRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedPromptStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowSetRequest and assigns it to the FlowSet field.
func (o *PatchedPromptStageRequest) SetFlowSet(v []FlowSetRequest) {
	o.FlowSet = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *PatchedPromptStageRequest) GetFields() []string {
	if o == nil || o.Fields == nil {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptStageRequest) GetFieldsOk() ([]string, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *PatchedPromptStageRequest) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *PatchedPromptStageRequest) SetFields(v []string) {
	o.Fields = v
}

// GetValidationPolicies returns the ValidationPolicies field value if set, zero value otherwise.
func (o *PatchedPromptStageRequest) GetValidationPolicies() []string {
	if o == nil || o.ValidationPolicies == nil {
		var ret []string
		return ret
	}
	return o.ValidationPolicies
}

// GetValidationPoliciesOk returns a tuple with the ValidationPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedPromptStageRequest) GetValidationPoliciesOk() ([]string, bool) {
	if o == nil || o.ValidationPolicies == nil {
		return nil, false
	}
	return o.ValidationPolicies, true
}

// HasValidationPolicies returns a boolean if a field has been set.
func (o *PatchedPromptStageRequest) HasValidationPolicies() bool {
	if o != nil && o.ValidationPolicies != nil {
		return true
	}

	return false
}

// SetValidationPolicies gets a reference to the given []string and assigns it to the ValidationPolicies field.
func (o *PatchedPromptStageRequest) SetValidationPolicies(v []string) {
	o.ValidationPolicies = v
}

func (o PatchedPromptStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.ValidationPolicies != nil {
		toSerialize["validation_policies"] = o.ValidationPolicies
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedPromptStageRequest struct {
	value *PatchedPromptStageRequest
	isSet bool
}

func (v NullablePatchedPromptStageRequest) Get() *PatchedPromptStageRequest {
	return v.value
}

func (v *NullablePatchedPromptStageRequest) Set(val *PatchedPromptStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedPromptStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedPromptStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedPromptStageRequest(val *PatchedPromptStageRequest) *NullablePatchedPromptStageRequest {
	return &NullablePatchedPromptStageRequest{value: val, isSet: true}
}

func (v NullablePatchedPromptStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedPromptStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
