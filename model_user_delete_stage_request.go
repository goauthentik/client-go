/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserDeleteStageRequest UserDeleteStage Serializer
type UserDeleteStageRequest struct {
	Name string `json:"name"`
	FlowSet *[]FlowRequest `json:"flow_set,omitempty"`
}

// NewUserDeleteStageRequest instantiates a new UserDeleteStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDeleteStageRequest(name string) *UserDeleteStageRequest {
	this := UserDeleteStageRequest{}
	this.Name = name
	return &this
}

// NewUserDeleteStageRequestWithDefaults instantiates a new UserDeleteStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDeleteStageRequestWithDefaults() *UserDeleteStageRequest {
	this := UserDeleteStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UserDeleteStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserDeleteStageRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserDeleteStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *UserDeleteStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return *o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDeleteStageRequest) GetFlowSetOk() (*[]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *UserDeleteStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *UserDeleteStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = &v
}

func (o UserDeleteStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	return json.Marshal(toSerialize)
}

type NullableUserDeleteStageRequest struct {
	value *UserDeleteStageRequest
	isSet bool
}

func (v NullableUserDeleteStageRequest) Get() *UserDeleteStageRequest {
	return v.value
}

func (v *NullableUserDeleteStageRequest) Set(val *UserDeleteStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDeleteStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDeleteStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDeleteStageRequest(val *UserDeleteStageRequest) *NullableUserDeleteStageRequest {
	return &NullableUserDeleteStageRequest{value: val, isSet: true}
}

func (v NullableUserDeleteStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDeleteStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


