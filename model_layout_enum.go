/*
authentik

Making authentication simple.

API version: 2022.9.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// LayoutEnum the model 'LayoutEnum'
type LayoutEnum string

// List of LayoutEnum
const (
	LAYOUTENUM_STACKED       LayoutEnum = "stacked"
	LAYOUTENUM_CONTENT_LEFT  LayoutEnum = "content_left"
	LAYOUTENUM_CONTENT_RIGHT LayoutEnum = "content_right"
	LAYOUTENUM_SIDEBAR_LEFT  LayoutEnum = "sidebar_left"
	LAYOUTENUM_SIDEBAR_RIGHT LayoutEnum = "sidebar_right"
)

// All allowed values of LayoutEnum enum
var AllowedLayoutEnumEnumValues = []LayoutEnum{
	"stacked",
	"content_left",
	"content_right",
	"sidebar_left",
	"sidebar_right",
}

func (v *LayoutEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LayoutEnum(value)
	for _, existing := range AllowedLayoutEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LayoutEnum", value)
}

// NewLayoutEnumFromValue returns a pointer to a valid LayoutEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLayoutEnumFromValue(v string) (*LayoutEnum, error) {
	ev := LayoutEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LayoutEnum: valid values are %v", v, AllowedLayoutEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LayoutEnum) IsValid() bool {
	for _, existing := range AllowedLayoutEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LayoutEnum value
func (v LayoutEnum) Ptr() *LayoutEnum {
	return &v
}

type NullableLayoutEnum struct {
	value *LayoutEnum
	isSet bool
}

func (v NullableLayoutEnum) Get() *LayoutEnum {
	return v.value
}

func (v *NullableLayoutEnum) Set(val *LayoutEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableLayoutEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableLayoutEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayoutEnum(val *LayoutEnum) *NullableLayoutEnum {
	return &NullableLayoutEnum{value: val, isSet: true}
}

func (v NullableLayoutEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayoutEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
