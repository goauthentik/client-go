/*
authentik

Making authentication simple.

API version: 2023.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// FlowLayoutEnum * `stacked` - Stacked * `content_left` - Content Left * `content_right` - Content Right * `sidebar_left` - Sidebar Left * `sidebar_right` - Sidebar Right
type FlowLayoutEnum string

// List of FlowLayoutEnum
const (
	FLOWLAYOUTENUM_STACKED       FlowLayoutEnum = "stacked"
	FLOWLAYOUTENUM_CONTENT_LEFT  FlowLayoutEnum = "content_left"
	FLOWLAYOUTENUM_CONTENT_RIGHT FlowLayoutEnum = "content_right"
	FLOWLAYOUTENUM_SIDEBAR_LEFT  FlowLayoutEnum = "sidebar_left"
	FLOWLAYOUTENUM_SIDEBAR_RIGHT FlowLayoutEnum = "sidebar_right"
)

// All allowed values of FlowLayoutEnum enum
var AllowedFlowLayoutEnumEnumValues = []FlowLayoutEnum{
	"stacked",
	"content_left",
	"content_right",
	"sidebar_left",
	"sidebar_right",
}

func (v *FlowLayoutEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlowLayoutEnum(value)
	for _, existing := range AllowedFlowLayoutEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FlowLayoutEnum", value)
}

// NewFlowLayoutEnumFromValue returns a pointer to a valid FlowLayoutEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlowLayoutEnumFromValue(v string) (*FlowLayoutEnum, error) {
	ev := FlowLayoutEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlowLayoutEnum: valid values are %v", v, AllowedFlowLayoutEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlowLayoutEnum) IsValid() bool {
	for _, existing := range AllowedFlowLayoutEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlowLayoutEnum value
func (v FlowLayoutEnum) Ptr() *FlowLayoutEnum {
	return &v
}

type NullableFlowLayoutEnum struct {
	value *FlowLayoutEnum
	isSet bool
}

func (v NullableFlowLayoutEnum) Get() *FlowLayoutEnum {
	return v.value
}

func (v *NullableFlowLayoutEnum) Set(val *FlowLayoutEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowLayoutEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowLayoutEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowLayoutEnum(val *FlowLayoutEnum) *NullableFlowLayoutEnum {
	return &NullableFlowLayoutEnum{value: val, isSet: true}
}

func (v NullableFlowLayoutEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowLayoutEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}