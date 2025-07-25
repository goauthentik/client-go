/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ContextualFlowInfoLayoutEnum the model 'ContextualFlowInfoLayoutEnum'
type ContextualFlowInfoLayoutEnum string

// List of ContextualFlowInfoLayoutEnum
const (
	CONTEXTUALFLOWINFOLAYOUTENUM_STACKED       ContextualFlowInfoLayoutEnum = "stacked"
	CONTEXTUALFLOWINFOLAYOUTENUM_CONTENT_LEFT  ContextualFlowInfoLayoutEnum = "content_left"
	CONTEXTUALFLOWINFOLAYOUTENUM_CONTENT_RIGHT ContextualFlowInfoLayoutEnum = "content_right"
	CONTEXTUALFLOWINFOLAYOUTENUM_SIDEBAR_LEFT  ContextualFlowInfoLayoutEnum = "sidebar_left"
	CONTEXTUALFLOWINFOLAYOUTENUM_SIDEBAR_RIGHT ContextualFlowInfoLayoutEnum = "sidebar_right"
)

// All allowed values of ContextualFlowInfoLayoutEnum enum
var AllowedContextualFlowInfoLayoutEnumEnumValues = []ContextualFlowInfoLayoutEnum{
	"stacked",
	"content_left",
	"content_right",
	"sidebar_left",
	"sidebar_right",
}

func (v *ContextualFlowInfoLayoutEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContextualFlowInfoLayoutEnum(value)
	for _, existing := range AllowedContextualFlowInfoLayoutEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContextualFlowInfoLayoutEnum", value)
}

// NewContextualFlowInfoLayoutEnumFromValue returns a pointer to a valid ContextualFlowInfoLayoutEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContextualFlowInfoLayoutEnumFromValue(v string) (*ContextualFlowInfoLayoutEnum, error) {
	ev := ContextualFlowInfoLayoutEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContextualFlowInfoLayoutEnum: valid values are %v", v, AllowedContextualFlowInfoLayoutEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContextualFlowInfoLayoutEnum) IsValid() bool {
	for _, existing := range AllowedContextualFlowInfoLayoutEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContextualFlowInfoLayoutEnum value
func (v ContextualFlowInfoLayoutEnum) Ptr() *ContextualFlowInfoLayoutEnum {
	return &v
}

type NullableContextualFlowInfoLayoutEnum struct {
	value *ContextualFlowInfoLayoutEnum
	isSet bool
}

func (v NullableContextualFlowInfoLayoutEnum) Get() *ContextualFlowInfoLayoutEnum {
	return v.value
}

func (v *NullableContextualFlowInfoLayoutEnum) Set(val *ContextualFlowInfoLayoutEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableContextualFlowInfoLayoutEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableContextualFlowInfoLayoutEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextualFlowInfoLayoutEnum(val *ContextualFlowInfoLayoutEnum) *NullableContextualFlowInfoLayoutEnum {
	return &NullableContextualFlowInfoLayoutEnum{value: val, isSet: true}
}

func (v NullableContextualFlowInfoLayoutEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextualFlowInfoLayoutEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
