/*
authentik

Making authentication simple.

API version: 2024.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ProxyMode the model 'ProxyMode'
type ProxyMode string

// List of ProxyMode
const (
	PROXYMODE_PROXY          ProxyMode = "proxy"
	PROXYMODE_FORWARD_SINGLE ProxyMode = "forward_single"
	PROXYMODE_FORWARD_DOMAIN ProxyMode = "forward_domain"
)

// All allowed values of ProxyMode enum
var AllowedProxyModeEnumValues = []ProxyMode{
	"proxy",
	"forward_single",
	"forward_domain",
}

func (v *ProxyMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProxyMode(value)
	for _, existing := range AllowedProxyModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProxyMode", value)
}

// NewProxyModeFromValue returns a pointer to a valid ProxyMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProxyModeFromValue(v string) (*ProxyMode, error) {
	ev := ProxyMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProxyMode: valid values are %v", v, AllowedProxyModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProxyMode) IsValid() bool {
	for _, existing := range AllowedProxyModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProxyMode value
func (v ProxyMode) Ptr() *ProxyMode {
	return &v
}

type NullableProxyMode struct {
	value *ProxyMode
	isSet bool
}

func (v NullableProxyMode) Get() *ProxyMode {
	return v.value
}

func (v *NullableProxyMode) Set(val *ProxyMode) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyMode) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyMode(val *ProxyMode) *NullableProxyMode {
	return &NullableProxyMode{value: val, isSet: true}
}

func (v NullableProxyMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
