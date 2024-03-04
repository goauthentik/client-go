/*
authentik

Making authentication simple.

API version: 2024.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// NetworkBindingEnum * `no_binding` - No Binding * `bind_asn` - Bind Asn * `bind_asn_network` - Bind Asn Network * `bind_asn_network_ip` - Bind Asn Network Ip
type NetworkBindingEnum string

// List of NetworkBindingEnum
const (
	NETWORKBINDINGENUM_NO_BINDING          NetworkBindingEnum = "no_binding"
	NETWORKBINDINGENUM_BIND_ASN            NetworkBindingEnum = "bind_asn"
	NETWORKBINDINGENUM_BIND_ASN_NETWORK    NetworkBindingEnum = "bind_asn_network"
	NETWORKBINDINGENUM_BIND_ASN_NETWORK_IP NetworkBindingEnum = "bind_asn_network_ip"
)

// All allowed values of NetworkBindingEnum enum
var AllowedNetworkBindingEnumEnumValues = []NetworkBindingEnum{
	"no_binding",
	"bind_asn",
	"bind_asn_network",
	"bind_asn_network_ip",
}

func (v *NetworkBindingEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkBindingEnum(value)
	for _, existing := range AllowedNetworkBindingEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkBindingEnum", value)
}

// NewNetworkBindingEnumFromValue returns a pointer to a valid NetworkBindingEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkBindingEnumFromValue(v string) (*NetworkBindingEnum, error) {
	ev := NetworkBindingEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkBindingEnum: valid values are %v", v, AllowedNetworkBindingEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkBindingEnum) IsValid() bool {
	for _, existing := range AllowedNetworkBindingEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkBindingEnum value
func (v NetworkBindingEnum) Ptr() *NetworkBindingEnum {
	return &v
}

type NullableNetworkBindingEnum struct {
	value *NetworkBindingEnum
	isSet bool
}

func (v NullableNetworkBindingEnum) Get() *NetworkBindingEnum {
	return v.value
}

func (v *NullableNetworkBindingEnum) Set(val *NetworkBindingEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkBindingEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkBindingEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkBindingEnum(val *NetworkBindingEnum) *NullableNetworkBindingEnum {
	return &NullableNetworkBindingEnum{value: val, isSet: true}
}

func (v NullableNetworkBindingEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkBindingEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
