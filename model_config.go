/*
authentik

Making authentication simple.

API version: 2022.5.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Config Serialize authentik Config into DRF Object
type Config struct {
	ErrorReporting         ErrorReportingConfig `json:"error_reporting"`
	Capabilities           []CapabilitiesEnum   `json:"capabilities"`
	CacheTimeout           int32                `json:"cache_timeout"`
	CacheTimeoutFlows      int32                `json:"cache_timeout_flows"`
	CacheTimeoutPolicies   int32                `json:"cache_timeout_policies"`
	CacheTimeoutReputation int32                `json:"cache_timeout_reputation"`
}

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
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error_reporting"] = o.ErrorReporting
	}
	if true {
		toSerialize["capabilities"] = o.Capabilities
	}
	if true {
		toSerialize["cache_timeout"] = o.CacheTimeout
	}
	if true {
		toSerialize["cache_timeout_flows"] = o.CacheTimeoutFlows
	}
	if true {
		toSerialize["cache_timeout_policies"] = o.CacheTimeoutPolicies
	}
	if true {
		toSerialize["cache_timeout_reputation"] = o.CacheTimeoutReputation
	}
	return json.Marshal(toSerialize)
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
