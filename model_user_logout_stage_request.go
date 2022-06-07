/*
authentik

Making authentication simple.

API version: 2022.6.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserLogoutStageRequest UserLogoutStage Serializer
type UserLogoutStageRequest struct {
	Name    string        `json:"name"`
	FlowSet []FlowRequest `json:"flow_set,omitempty"`
}

// NewUserLogoutStageRequest instantiates a new UserLogoutStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLogoutStageRequest(name string) *UserLogoutStageRequest {
	this := UserLogoutStageRequest{}
	this.Name = name
	return &this
}

// NewUserLogoutStageRequestWithDefaults instantiates a new UserLogoutStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLogoutStageRequestWithDefaults() *UserLogoutStageRequest {
	this := UserLogoutStageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UserLogoutStageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserLogoutStageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserLogoutStageRequest) SetName(v string) {
	o.Name = v
}

// GetFlowSet returns the FlowSet field value if set, zero value otherwise.
func (o *UserLogoutStageRequest) GetFlowSet() []FlowRequest {
	if o == nil || o.FlowSet == nil {
		var ret []FlowRequest
		return ret
	}
	return o.FlowSet
}

// GetFlowSetOk returns a tuple with the FlowSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLogoutStageRequest) GetFlowSetOk() ([]FlowRequest, bool) {
	if o == nil || o.FlowSet == nil {
		return nil, false
	}
	return o.FlowSet, true
}

// HasFlowSet returns a boolean if a field has been set.
func (o *UserLogoutStageRequest) HasFlowSet() bool {
	if o != nil && o.FlowSet != nil {
		return true
	}

	return false
}

// SetFlowSet gets a reference to the given []FlowRequest and assigns it to the FlowSet field.
func (o *UserLogoutStageRequest) SetFlowSet(v []FlowRequest) {
	o.FlowSet = v
}

func (o UserLogoutStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.FlowSet != nil {
		toSerialize["flow_set"] = o.FlowSet
	}
	return json.Marshal(toSerialize)
}

type NullableUserLogoutStageRequest struct {
	value *UserLogoutStageRequest
	isSet bool
}

func (v NullableUserLogoutStageRequest) Get() *UserLogoutStageRequest {
	return v.value
}

func (v *NullableUserLogoutStageRequest) Set(val *UserLogoutStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLogoutStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLogoutStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLogoutStageRequest(val *UserLogoutStageRequest) *NullableUserLogoutStageRequest {
	return &NullableUserLogoutStageRequest{value: val, isSet: true}
}

func (v NullableUserLogoutStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLogoutStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
