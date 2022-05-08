/*
authentik

Making authentication simple.

API version: 2022.4.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// BindModeEnum the model 'BindModeEnum'
type BindModeEnum string

// List of BindModeEnum
const (
	BINDMODEENUM_DIRECT BindModeEnum = "direct"
	BINDMODEENUM_CACHED BindModeEnum = "cached"
)

var allowedBindModeEnumEnumValues = []BindModeEnum{
	"direct",
	"cached",
}

func (v *BindModeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BindModeEnum(value)
	for _, existing := range allowedBindModeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BindModeEnum", value)
}

// NewBindModeEnumFromValue returns a pointer to a valid BindModeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBindModeEnumFromValue(v string) (*BindModeEnum, error) {
	ev := BindModeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BindModeEnum: valid values are %v", v, allowedBindModeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BindModeEnum) IsValid() bool {
	for _, existing := range allowedBindModeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BindModeEnum value
func (v BindModeEnum) Ptr() *BindModeEnum {
	return &v
}

type NullableBindModeEnum struct {
	value *BindModeEnum
	isSet bool
}

func (v NullableBindModeEnum) Get() *BindModeEnum {
	return v.value
}

func (v *NullableBindModeEnum) Set(val *BindModeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBindModeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBindModeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindModeEnum(val *BindModeEnum) *NullableBindModeEnum {
	return &NullableBindModeEnum{value: val, isSet: true}
}

func (v NullableBindModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindModeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
