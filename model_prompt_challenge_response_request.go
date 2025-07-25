/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PromptChallengeResponseRequest Validate response, fields are dynamically created based on the stage
type PromptChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
}

// NewPromptChallengeResponseRequest instantiates a new PromptChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromptChallengeResponseRequest() *PromptChallengeResponseRequest {
	this := PromptChallengeResponseRequest{}
	var component string = "ak-stage-prompt"
	this.Component = &component
	return &this
}

// NewPromptChallengeResponseRequestWithDefaults instantiates a new PromptChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromptChallengeResponseRequestWithDefaults() *PromptChallengeResponseRequest {
	this := PromptChallengeResponseRequest{}
	var component string = "ak-stage-prompt"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *PromptChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromptChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *PromptChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *PromptChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

func (o PromptChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullablePromptChallengeResponseRequest struct {
	value *PromptChallengeResponseRequest
	isSet bool
}

func (v NullablePromptChallengeResponseRequest) Get() *PromptChallengeResponseRequest {
	return v.value
}

func (v *NullablePromptChallengeResponseRequest) Set(val *PromptChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePromptChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePromptChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromptChallengeResponseRequest(val *PromptChallengeResponseRequest) *NullablePromptChallengeResponseRequest {
	return &NullablePromptChallengeResponseRequest{value: val, isSet: true}
}

func (v NullablePromptChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromptChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
