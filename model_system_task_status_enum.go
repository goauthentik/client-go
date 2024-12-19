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

// SystemTaskStatusEnum the model 'SystemTaskStatusEnum'
type SystemTaskStatusEnum string

// List of SystemTaskStatusEnum
const (
	SYSTEMTASKSTATUSENUM_UNKNOWN    SystemTaskStatusEnum = "unknown"
	SYSTEMTASKSTATUSENUM_SUCCESSFUL SystemTaskStatusEnum = "successful"
	SYSTEMTASKSTATUSENUM_WARNING    SystemTaskStatusEnum = "warning"
	SYSTEMTASKSTATUSENUM_ERROR      SystemTaskStatusEnum = "error"
)

// All allowed values of SystemTaskStatusEnum enum
var AllowedSystemTaskStatusEnumEnumValues = []SystemTaskStatusEnum{
	"unknown",
	"successful",
	"warning",
	"error",
}

func (v *SystemTaskStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SystemTaskStatusEnum(value)
	for _, existing := range AllowedSystemTaskStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SystemTaskStatusEnum", value)
}

// NewSystemTaskStatusEnumFromValue returns a pointer to a valid SystemTaskStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSystemTaskStatusEnumFromValue(v string) (*SystemTaskStatusEnum, error) {
	ev := SystemTaskStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SystemTaskStatusEnum: valid values are %v", v, AllowedSystemTaskStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SystemTaskStatusEnum) IsValid() bool {
	for _, existing := range AllowedSystemTaskStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SystemTaskStatusEnum value
func (v SystemTaskStatusEnum) Ptr() *SystemTaskStatusEnum {
	return &v
}

type NullableSystemTaskStatusEnum struct {
	value *SystemTaskStatusEnum
	isSet bool
}

func (v NullableSystemTaskStatusEnum) Get() *SystemTaskStatusEnum {
	return v.value
}

func (v *NullableSystemTaskStatusEnum) Set(val *SystemTaskStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemTaskStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemTaskStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemTaskStatusEnum(val *SystemTaskStatusEnum) *NullableSystemTaskStatusEnum {
	return &NullableSystemTaskStatusEnum{value: val, isSet: true}
}

func (v NullableSystemTaskStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemTaskStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
