/*
authentik

Making authentication simple.

API version: 2023.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PolicyTestRequest Test policy execution for a user with context
type PolicyTestRequest struct {
	User    int32                  `json:"user"`
	Context map[string]interface{} `json:"context,omitempty"`
}

// NewPolicyTestRequest instantiates a new PolicyTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyTestRequest(user int32) *PolicyTestRequest {
	this := PolicyTestRequest{}
	this.User = user
	return &this
}

// NewPolicyTestRequestWithDefaults instantiates a new PolicyTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyTestRequestWithDefaults() *PolicyTestRequest {
	this := PolicyTestRequest{}
	return &this
}

// GetUser returns the User field value
func (o *PolicyTestRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *PolicyTestRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *PolicyTestRequest) SetUser(v int32) {
	o.User = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *PolicyTestRequest) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyTestRequest) GetContextOk() (map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PolicyTestRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *PolicyTestRequest) SetContext(v map[string]interface{}) {
	o.Context = v
}

func (o PolicyTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyTestRequest struct {
	value *PolicyTestRequest
	isSet bool
}

func (v NullablePolicyTestRequest) Get() *PolicyTestRequest {
	return v.value
}

func (v *NullablePolicyTestRequest) Set(val *PolicyTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyTestRequest(val *PolicyTestRequest) *NullablePolicyTestRequest {
	return &NullablePolicyTestRequest{value: val, isSet: true}
}

func (v NullablePolicyTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
