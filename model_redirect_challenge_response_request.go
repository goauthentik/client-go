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

// RedirectChallengeResponseRequest Redirect challenge response
type RedirectChallengeResponseRequest struct {
	Component *string `json:"component,omitempty"`
	To        string  `json:"to"`
}

// NewRedirectChallengeResponseRequest instantiates a new RedirectChallengeResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectChallengeResponseRequest(to string) *RedirectChallengeResponseRequest {
	this := RedirectChallengeResponseRequest{}
	var component string = "xak-flow-redirect"
	this.Component = &component
	this.To = to
	return &this
}

// NewRedirectChallengeResponseRequestWithDefaults instantiates a new RedirectChallengeResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectChallengeResponseRequestWithDefaults() *RedirectChallengeResponseRequest {
	this := RedirectChallengeResponseRequest{}
	var component string = "xak-flow-redirect"
	this.Component = &component
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *RedirectChallengeResponseRequest) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectChallengeResponseRequest) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *RedirectChallengeResponseRequest) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *RedirectChallengeResponseRequest) SetComponent(v string) {
	o.Component = &v
}

// GetTo returns the To field value
func (o *RedirectChallengeResponseRequest) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *RedirectChallengeResponseRequest) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *RedirectChallengeResponseRequest) SetTo(v string) {
	o.To = v
}

func (o RedirectChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableRedirectChallengeResponseRequest struct {
	value *RedirectChallengeResponseRequest
	isSet bool
}

func (v NullableRedirectChallengeResponseRequest) Get() *RedirectChallengeResponseRequest {
	return v.value
}

func (v *NullableRedirectChallengeResponseRequest) Set(val *RedirectChallengeResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectChallengeResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectChallengeResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectChallengeResponseRequest(val *RedirectChallengeResponseRequest) *NullableRedirectChallengeResponseRequest {
	return &NullableRedirectChallengeResponseRequest{value: val, isSet: true}
}

func (v NullableRedirectChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectChallengeResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
