/*
authentik

Making authentication simple.

API version: 2021.10.1-rc1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// OutpostDefaultConfig Global default outpost config
type OutpostDefaultConfig struct {
	Config map[string]interface{} `json:"config"`
}

// NewOutpostDefaultConfig instantiates a new OutpostDefaultConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutpostDefaultConfig(config map[string]interface{}) *OutpostDefaultConfig {
	this := OutpostDefaultConfig{}
	this.Config = config
	return &this
}

// NewOutpostDefaultConfigWithDefaults instantiates a new OutpostDefaultConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutpostDefaultConfigWithDefaults() *OutpostDefaultConfig {
	this := OutpostDefaultConfig{}
	return &this
}

// GetConfig returns the Config field value
func (o *OutpostDefaultConfig) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *OutpostDefaultConfig) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *OutpostDefaultConfig) SetConfig(v map[string]interface{}) {
	o.Config = v
}

func (o OutpostDefaultConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableOutpostDefaultConfig struct {
	value *OutpostDefaultConfig
	isSet bool
}

func (v NullableOutpostDefaultConfig) Get() *OutpostDefaultConfig {
	return v.value
}

func (v *NullableOutpostDefaultConfig) Set(val *OutpostDefaultConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOutpostDefaultConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOutpostDefaultConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutpostDefaultConfig(val *OutpostDefaultConfig) *NullableOutpostDefaultConfig {
	return &NullableOutpostDefaultConfig{value: val, isSet: true}
}

func (v NullableOutpostDefaultConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutpostDefaultConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
