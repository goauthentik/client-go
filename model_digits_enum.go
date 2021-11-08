/*
authentik

Making authentication simple.

API version: 2021.10.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// DigitsEnum the model 'DigitsEnum'
type DigitsEnum int32

// List of DigitsEnum
const (
	DIGITSENUM__6 DigitsEnum = 6
	DIGITSENUM__8 DigitsEnum = 8
)

var allowedDigitsEnumEnumValues = []DigitsEnum{
	6,
	8,
}

func (v *DigitsEnum) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DigitsEnum(value)
	for _, existing := range allowedDigitsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DigitsEnum", value)
}

// NewDigitsEnumFromValue returns a pointer to a valid DigitsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDigitsEnumFromValue(v int32) (*DigitsEnum, error) {
	ev := DigitsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DigitsEnum: valid values are %v", v, allowedDigitsEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DigitsEnum) IsValid() bool {
	for _, existing := range allowedDigitsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DigitsEnum value
func (v DigitsEnum) Ptr() *DigitsEnum {
	return &v
}

type NullableDigitsEnum struct {
	value *DigitsEnum
	isSet bool
}

func (v NullableDigitsEnum) Get() *DigitsEnum {
	return v.value
}

func (v *NullableDigitsEnum) Set(val *DigitsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitsEnum(val *DigitsEnum) *NullableDigitsEnum {
	return &NullableDigitsEnum{value: val, isSet: true}
}

func (v NullableDigitsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
