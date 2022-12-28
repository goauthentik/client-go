/*
authentik

Making authentication simple.

API version: 2022.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AutoSubmitChallengeResponseRequest Pseudo class for autosubmit response
type AutoSubmitChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
}

// NewAutoSubmitChallengeResponseRequest instantiates a new AutoSubmitChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoSubmitChallengeResponseRequest() *AutoSubmitChallengeResponseRequest {
	this := AutoSubmitChallengeResponseRequest{}
	var component string = "ak-stage-autosubmit"
	this.Component = &component
	return &this
}

// NewAutoSubmitChallengeResponseRequestWithDefaults instantiates a new AutoSubmitChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoSubmitChallengeResponseRequestWithDefaults() *AutoSubmitChallengeResponseRequest {
	this := AutoSubmitChallengeResponseRequest{}
	var component string = "ak-stage-autosubmit"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AutoSubmitChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoSubmitChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AutoSubmitChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AutoSubmitChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

func (o AutoSubmitChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullableAutoSubmitChallengeResponseRequest struct {
	value *AutoSubmitChallengeResponseRequest
	isSet bool
}

func (v NullableAutoSubmitChallengeResponseRequest) Get() *AutoSubmitChallengeResponseRequest {
	return v.value
}

func (v *NullableAutoSubmitChallengeResponseRequest) Set(val *AutoSubmitChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoSubmitChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoSubmitChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoSubmitChallengeResponseRequest(val *AutoSubmitChallengeResponseRequest) *NullableAutoSubmitChallengeResponseRequest {
	return &NullableAutoSubmitChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableAutoSubmitChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoSubmitChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
