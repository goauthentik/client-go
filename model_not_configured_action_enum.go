/*
authentik

Making authentication simple.

API version: 2024.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// NotConfiguredActionEnum the model 'NotConfiguredActionEnum'
type NotConfiguredActionEnum string

// List of NotConfiguredActionEnum
const (
	NOTCONFIGUREDACTIONENUM_SKIP      NotConfiguredActionEnum = "skip"
	NOTCONFIGUREDACTIONENUM_DENY      NotConfiguredActionEnum = "deny"
	NOTCONFIGUREDACTIONENUM_CONFIGURE NotConfiguredActionEnum = "configure"
)

// All allowed values of NotConfiguredActionEnum enum
var AllowedNotConfiguredActionEnumEnumValues = []NotConfiguredActionEnum{
	"skip",
	"deny",
	"configure",
}

func (v *NotConfiguredActionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotConfiguredActionEnum(value)
	for _, existing := range AllowedNotConfiguredActionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotConfiguredActionEnum", value)
}

// NewNotConfiguredActionEnumFromValue returns a pointer to a valid NotConfiguredActionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotConfiguredActionEnumFromValue(v string) (*NotConfiguredActionEnum, error) {
	ev := NotConfiguredActionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotConfiguredActionEnum: valid values are %v", v, AllowedNotConfiguredActionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotConfiguredActionEnum) IsValid() bool {
	for _, existing := range AllowedNotConfiguredActionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotConfiguredActionEnum value
func (v NotConfiguredActionEnum) Ptr() *NotConfiguredActionEnum {
	return &v
}

type NullableNotConfiguredActionEnum struct {
	value *NotConfiguredActionEnum
	isSet bool
}

func (v NullableNotConfiguredActionEnum) Get() *NotConfiguredActionEnum {
	return v.value
}

func (v *NullableNotConfiguredActionEnum) Set(val *NotConfiguredActionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableNotConfiguredActionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableNotConfiguredActionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotConfiguredActionEnum(val *NotConfiguredActionEnum) *NullableNotConfiguredActionEnum {
	return &NullableNotConfiguredActionEnum{value: val, isSet: true}
}

func (v NullableNotConfiguredActionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotConfiguredActionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
