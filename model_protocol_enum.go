/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ProtocolEnum the model 'ProtocolEnum'
type ProtocolEnum string

// List of ProtocolEnum
const (
	PROTOCOLENUM_RDP ProtocolEnum = "rdp"
	PROTOCOLENUM_VNC ProtocolEnum = "vnc"
	PROTOCOLENUM_SSH ProtocolEnum = "ssh"
)

// All allowed values of ProtocolEnum enum
var AllowedProtocolEnumEnumValues = []ProtocolEnum{
	"rdp",
	"vnc",
	"ssh",
}

func (v *ProtocolEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProtocolEnum(value)
	for _, existing := range AllowedProtocolEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProtocolEnum", value)
}

// NewProtocolEnumFromValue returns a pointer to a valid ProtocolEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProtocolEnumFromValue(v string) (*ProtocolEnum, error) {
	ev := ProtocolEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProtocolEnum: valid values are %v", v, AllowedProtocolEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProtocolEnum) IsValid() bool {
	for _, existing := range AllowedProtocolEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProtocolEnum value
func (v ProtocolEnum) Ptr() *ProtocolEnum {
	return &v
}

type NullableProtocolEnum struct {
	value *ProtocolEnum
	isSet bool
}

func (v NullableProtocolEnum) Get() *ProtocolEnum {
	return v.value
}

func (v *NullableProtocolEnum) Set(val *ProtocolEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolEnum(val *ProtocolEnum) *NullableProtocolEnum {
	return &NullableProtocolEnum{value: val, isSet: true}
}

func (v NullableProtocolEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
