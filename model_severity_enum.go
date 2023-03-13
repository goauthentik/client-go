/*
authentik

Making authentication simple.

API version: 2023.3.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// SeverityEnum * `notice` - Notice * `warning` - Warning * `alert` - Alert
type SeverityEnum string

// List of SeverityEnum
const (
	SEVERITYENUM_NOTICE  SeverityEnum = "notice"
	SEVERITYENUM_WARNING SeverityEnum = "warning"
	SEVERITYENUM_ALERT   SeverityEnum = "alert"
)

// All allowed values of SeverityEnum enum
var AllowedSeverityEnumEnumValues = []SeverityEnum{
	"notice",
	"warning",
	"alert",
}

func (v *SeverityEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SeverityEnum(value)
	for _, existing := range AllowedSeverityEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SeverityEnum", value)
}

// NewSeverityEnumFromValue returns a pointer to a valid SeverityEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSeverityEnumFromValue(v string) (*SeverityEnum, error) {
	ev := SeverityEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SeverityEnum: valid values are %v", v, AllowedSeverityEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SeverityEnum) IsValid() bool {
	for _, existing := range AllowedSeverityEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeverityEnum value
func (v SeverityEnum) Ptr() *SeverityEnum {
	return &v
}

type NullableSeverityEnum struct {
	value *SeverityEnum
	isSet bool
}

func (v NullableSeverityEnum) Get() *SeverityEnum {
	return v.value
}

func (v *NullableSeverityEnum) Set(val *SeverityEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSeverityEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSeverityEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeverityEnum(val *SeverityEnum) *NullableSeverityEnum {
	return &NullableSeverityEnum{value: val, isSet: true}
}

func (v NullableSeverityEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeverityEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
