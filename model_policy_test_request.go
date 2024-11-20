/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PolicyTestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyTestRequest{}

// PolicyTestRequest Test policy execution for a user with context
type PolicyTestRequest struct {
	User    int32                  `json:"user"`
	Context map[string]interface{} `json:"context,omitempty"`
}

type _PolicyTestRequest PolicyTestRequest

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
	if o == nil || IsNil(o.Context) {
		var ret map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyTestRequest) GetContextOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Context) {
		return map[string]interface{}{}, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PolicyTestRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *PolicyTestRequest) SetContext(v map[string]interface{}) {
	o.Context = v
}

func (o PolicyTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	return toSerialize, nil
}

func (o *PolicyTestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPolicyTestRequest := _PolicyTestRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolicyTestRequest)

	if err != nil {
		return err
	}

	*o = PolicyTestRequest(varPolicyTestRequest)

	return err
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
