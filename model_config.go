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

// checks if the Config type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Config{}

// Config Serialize authentik Config into DRF Object
type Config struct {
	ErrorReporting         ErrorReportingConfig `json:"error_reporting"`
	Capabilities           []CapabilitiesEnum   `json:"capabilities"`
	CacheTimeout           int32                `json:"cache_timeout"`
	CacheTimeoutFlows      int32                `json:"cache_timeout_flows"`
	CacheTimeoutPolicies   int32                `json:"cache_timeout_policies"`
	CacheTimeoutReputation int32                `json:"cache_timeout_reputation"`
}

type _Config Config

// NewConfig instantiates a new Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfig(errorReporting ErrorReportingConfig, capabilities []CapabilitiesEnum, cacheTimeout int32, cacheTimeoutFlows int32, cacheTimeoutPolicies int32, cacheTimeoutReputation int32) *Config {
	this := Config{}
	this.ErrorReporting = errorReporting
	this.Capabilities = capabilities
	this.CacheTimeout = cacheTimeout
	this.CacheTimeoutFlows = cacheTimeoutFlows
	this.CacheTimeoutPolicies = cacheTimeoutPolicies
	this.CacheTimeoutReputation = cacheTimeoutReputation
	return &this
}

// NewConfigWithDefaults instantiates a new Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigWithDefaults() *Config {
	this := Config{}
	return &this
}

// GetErrorReporting returns the ErrorReporting field value
func (o *Config) GetErrorReporting() ErrorReportingConfig {
	if o == nil {
		var ret ErrorReportingConfig
		return ret
	}

	return o.ErrorReporting
}

// GetErrorReportingOk returns a tuple with the ErrorReporting field value
// and a boolean to check if the value has been set.
func (o *Config) GetErrorReportingOk() (*ErrorReportingConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorReporting, true
}

// SetErrorReporting sets field value
func (o *Config) SetErrorReporting(v ErrorReportingConfig) {
	o.ErrorReporting = v
}

// GetCapabilities returns the Capabilities field value
func (o *Config) GetCapabilities() []CapabilitiesEnum {
	if o == nil {
		var ret []CapabilitiesEnum
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *Config) GetCapabilitiesOk() ([]CapabilitiesEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *Config) SetCapabilities(v []CapabilitiesEnum) {
	o.Capabilities = v
}

// GetCacheTimeout returns the CacheTimeout field value
func (o *Config) GetCacheTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CacheTimeout
}

// GetCacheTimeoutOk returns a tuple with the CacheTimeout field value
// and a boolean to check if the value has been set.
func (o *Config) GetCacheTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheTimeout, true
}

// SetCacheTimeout sets field value
func (o *Config) SetCacheTimeout(v int32) {
	o.CacheTimeout = v
}

// GetCacheTimeoutFlows returns the CacheTimeoutFlows field value
func (o *Config) GetCacheTimeoutFlows() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CacheTimeoutFlows
}

// GetCacheTimeoutFlowsOk returns a tuple with the CacheTimeoutFlows field value
// and a boolean to check if the value has been set.
func (o *Config) GetCacheTimeoutFlowsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheTimeoutFlows, true
}

// SetCacheTimeoutFlows sets field value
func (o *Config) SetCacheTimeoutFlows(v int32) {
	o.CacheTimeoutFlows = v
}

// GetCacheTimeoutPolicies returns the CacheTimeoutPolicies field value
func (o *Config) GetCacheTimeoutPolicies() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CacheTimeoutPolicies
}

// GetCacheTimeoutPoliciesOk returns a tuple with the CacheTimeoutPolicies field value
// and a boolean to check if the value has been set.
func (o *Config) GetCacheTimeoutPoliciesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheTimeoutPolicies, true
}

// SetCacheTimeoutPolicies sets field value
func (o *Config) SetCacheTimeoutPolicies(v int32) {
	o.CacheTimeoutPolicies = v
}

// GetCacheTimeoutReputation returns the CacheTimeoutReputation field value
func (o *Config) GetCacheTimeoutReputation() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CacheTimeoutReputation
}

// GetCacheTimeoutReputationOk returns a tuple with the CacheTimeoutReputation field value
// and a boolean to check if the value has been set.
func (o *Config) GetCacheTimeoutReputationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheTimeoutReputation, true
}

// SetCacheTimeoutReputation sets field value
func (o *Config) SetCacheTimeoutReputation(v int32) {
	o.CacheTimeoutReputation = v
}

func (o Config) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Config) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_reporting"] = o.ErrorReporting
	toSerialize["capabilities"] = o.Capabilities
	toSerialize["cache_timeout"] = o.CacheTimeout
	toSerialize["cache_timeout_flows"] = o.CacheTimeoutFlows
	toSerialize["cache_timeout_policies"] = o.CacheTimeoutPolicies
	toSerialize["cache_timeout_reputation"] = o.CacheTimeoutReputation
	return toSerialize, nil
}

func (o *Config) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error_reporting",
		"capabilities",
		"cache_timeout",
		"cache_timeout_flows",
		"cache_timeout_policies",
		"cache_timeout_reputation",
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

	varConfig := _Config{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfig)

	if err != nil {
		return err
	}

	*o = Config(varConfig)

	return err
}

type NullableConfig struct {
	value *Config
	isSet bool
}

func (v NullableConfig) Get() *Config {
	return v.value
}

func (v *NullableConfig) Set(val *Config) {
	v.value = val
	v.isSet = true
}

func (v NullableConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfig(val *Config) *NullableConfig {
	return &NullableConfig{value: val, isSet: true}
}

func (v NullableConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
