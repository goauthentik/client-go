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

// EmailChallengeResponseRequest Email challenge resposen. No fields. This challenge is always declared invalid to give the user a chance to retry
type EmailChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
}

// NewEmailChallengeResponseRequest instantiates a new EmailChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailChallengeResponseRequest() *EmailChallengeResponseRequest {
	this := EmailChallengeResponseRequest{}
	var component string = "ak-stage-email"
	this.Component = &component
	return &this
}

// NewEmailChallengeResponseRequestWithDefaults instantiates a new EmailChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailChallengeResponseRequestWithDefaults() *EmailChallengeResponseRequest {
	this := EmailChallengeResponseRequest{}
	var component string = "ak-stage-email"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *EmailChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *EmailChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *EmailChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

func (o EmailChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	return json.Marshal(toSerialize)
}

type NullableEmailChallengeResponseRequest struct {
	value *EmailChallengeResponseRequest
	isSet bool
}

func (v NullableEmailChallengeResponseRequest) Get() *EmailChallengeResponseRequest {
	return v.value
}

func (v *NullableEmailChallengeResponseRequest) Set(val *EmailChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailChallengeResponseRequest(val *EmailChallengeResponseRequest) *NullableEmailChallengeResponseRequest {
	return &NullableEmailChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableEmailChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


