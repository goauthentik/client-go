/*
authentik

Making authentication simple.

API version: 2024.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FrameChallengeResponseRequest Base class for all challenge responses
type FrameChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
}

// NewFrameChallengeResponseRequest instantiates a new FrameChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameChallengeResponseRequest() *FrameChallengeResponseRequest {
	this := FrameChallengeResponseRequest{}
	var component string = "xak-flow-frame"
	this.Component = &component
	return &this
}

// NewFrameChallengeResponseRequestWithDefaults instantiates a new FrameChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameChallengeResponseRequestWithDefaults() *FrameChallengeResponseRequest {
	this := FrameChallengeResponseRequest{}
	var component string = "xak-flow-frame"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *FrameChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *FrameChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *FrameChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

func (o FrameChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullableFrameChallengeResponseRequest struct {
	value *FrameChallengeResponseRequest
	isSet bool
}

func (v NullableFrameChallengeResponseRequest) Get() *FrameChallengeResponseRequest {
	return v.value
}

func (v *NullableFrameChallengeResponseRequest) Set(val *FrameChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameChallengeResponseRequest(val *FrameChallengeResponseRequest) *NullableFrameChallengeResponseRequest {
	return &NullableFrameChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableFrameChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
