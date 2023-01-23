/*
authentik

Making authentication simple.

API version: 2023.1.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// TaskStatusEnum the model 'TaskStatusEnum'
type TaskStatusEnum string

// List of TaskStatusEnum
const (
	TASKSTATUSENUM_SUCCESSFUL TaskStatusEnum = "SUCCESSFUL"
	TASKSTATUSENUM_WARNING    TaskStatusEnum = "WARNING"
	TASKSTATUSENUM_ERROR      TaskStatusEnum = "ERROR"
	TASKSTATUSENUM_UNKNOWN    TaskStatusEnum = "UNKNOWN"
)

// All allowed values of TaskStatusEnum enum
var AllowedTaskStatusEnumEnumValues = []TaskStatusEnum{
	"SUCCESSFUL",
	"WARNING",
	"ERROR",
	"UNKNOWN",
}

func (v *TaskStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskStatusEnum(value)
	for _, existing := range AllowedTaskStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskStatusEnum", value)
}

// NewTaskStatusEnumFromValue returns a pointer to a valid TaskStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskStatusEnumFromValue(v string) (*TaskStatusEnum, error) {
	ev := TaskStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaskStatusEnum: valid values are %v", v, AllowedTaskStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskStatusEnum) IsValid() bool {
	for _, existing := range AllowedTaskStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskStatusEnum value
func (v TaskStatusEnum) Ptr() *TaskStatusEnum {
	return &v
}

type NullableTaskStatusEnum struct {
	value *TaskStatusEnum
	isSet bool
}

func (v NullableTaskStatusEnum) Get() *TaskStatusEnum {
	return v.value
}

func (v *NullableTaskStatusEnum) Set(val *TaskStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskStatusEnum(val *TaskStatusEnum) *NullableTaskStatusEnum {
	return &NullableTaskStatusEnum{value: val, isSet: true}
}

func (v NullableTaskStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
