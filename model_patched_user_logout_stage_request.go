/*
authentik

Making authentication simple.

API version: 2021.9.1-rc1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedUserLogoutStageRequest UserLogoutStage Serializer
type PatchedUserLogoutStageRequest struct {
	Name *string `json:"name,omitempty"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
}

// NewPatchedUserLogoutStageRequest instantiates a new PatchedUserLogoutStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserLogoutStageRequest() *PatchedUserLogoutStageRequest {
	this := PatchedUserLogoutStageRequest{}
	return &this
}

// NewPatchedUserLogoutStageRequestWithDefaults instantiates a new PatchedUserLogoutStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserLogoutStageRequestWithDefaults() *PatchedUserLogoutStageRequest {
	this := PatchedUserLogoutStageRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedUserLogoutStageRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserLogoutStageRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedUserLogoutStageRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedUserLogoutStageRequest) SetName(v string) {
	o.Name = &v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *PatchedUserLogoutStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserLogoutStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *PatchedUserLogoutStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *PatchedUserLogoutStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

func (o PatchedUserLogoutStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedUserLogoutStageRequest struct {
	value *PatchedUserLogoutStageRequest
	isSet bool
}

func (v NullablePatchedUserLogoutStageRequest) Get() *PatchedUserLogoutStageRequest {
	return v.value
}

func (v *NullablePatchedUserLogoutStageRequest) Set(val *PatchedUserLogoutStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserLogoutStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserLogoutStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserLogoutStageRequest(val *PatchedUserLogoutStageRequest) *NullablePatchedUserLogoutStageRequest {
	return &NullablePatchedUserLogoutStageRequest{value: val, isSet: true}
}

func (v NullablePatchedUserLogoutStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserLogoutStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


