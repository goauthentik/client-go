/*
authentik

Making authentication simple.

API version: 2025.2.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ProviderEnum the model 'ProviderEnum'
type ProviderEnum string

// List of ProviderEnum
const (
	PROVIDERENUM_TWILIO  ProviderEnum = "twilio"
	PROVIDERENUM_GENERIC ProviderEnum = "generic"
)

// All allowed values of ProviderEnum enum
var AllowedProviderEnumEnumValues = []ProviderEnum{
	"twilio",
	"generic",
}

func (v *ProviderEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProviderEnum(value)
	for _, existing := range AllowedProviderEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProviderEnum", value)
}

// NewProviderEnumFromValue returns a pointer to a valid ProviderEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProviderEnumFromValue(v string) (*ProviderEnum, error) {
	ev := ProviderEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProviderEnum: valid values are %v", v, AllowedProviderEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProviderEnum) IsValid() bool {
	for _, existing := range AllowedProviderEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProviderEnum value
func (v ProviderEnum) Ptr() *ProviderEnum {
	return &v
}

type NullableProviderEnum struct {
	value *ProviderEnum
	isSet bool
}

func (v NullableProviderEnum) Get() *ProviderEnum {
	return v.value
}

func (v *NullableProviderEnum) Set(val *ProviderEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderEnum(val *ProviderEnum) *NullableProviderEnum {
	return &NullableProviderEnum{value: val, isSet: true}
}

func (v NullableProviderEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
