/*
authentik

Making authentication simple.

API version: 2024.6.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// FlowDesignationEnum the model 'FlowDesignationEnum'
type FlowDesignationEnum string

// List of FlowDesignationEnum
const (
	FLOWDESIGNATIONENUM_AUTHENTICATION      FlowDesignationEnum = "authentication"
	FLOWDESIGNATIONENUM_AUTHORIZATION       FlowDesignationEnum = "authorization"
	FLOWDESIGNATIONENUM_INVALIDATION        FlowDesignationEnum = "invalidation"
	FLOWDESIGNATIONENUM_ENROLLMENT          FlowDesignationEnum = "enrollment"
	FLOWDESIGNATIONENUM_UNENROLLMENT        FlowDesignationEnum = "unenrollment"
	FLOWDESIGNATIONENUM_RECOVERY            FlowDesignationEnum = "recovery"
	FLOWDESIGNATIONENUM_STAGE_CONFIGURATION FlowDesignationEnum = "stage_configuration"
)

// All allowed values of FlowDesignationEnum enum
var AllowedFlowDesignationEnumEnumValues = []FlowDesignationEnum{
	"authentication",
	"authorization",
	"invalidation",
	"enrollment",
	"unenrollment",
	"recovery",
	"stage_configuration",
}

func (v *FlowDesignationEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlowDesignationEnum(value)
	for _, existing := range AllowedFlowDesignationEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FlowDesignationEnum", value)
}

// NewFlowDesignationEnumFromValue returns a pointer to a valid FlowDesignationEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlowDesignationEnumFromValue(v string) (*FlowDesignationEnum, error) {
	ev := FlowDesignationEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlowDesignationEnum: valid values are %v", v, AllowedFlowDesignationEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlowDesignationEnum) IsValid() bool {
	for _, existing := range AllowedFlowDesignationEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlowDesignationEnum value
func (v FlowDesignationEnum) Ptr() *FlowDesignationEnum {
	return &v
}

type NullableFlowDesignationEnum struct {
	value *FlowDesignationEnum
	isSet bool
}

func (v NullableFlowDesignationEnum) Get() *FlowDesignationEnum {
	return v.value
}

func (v *NullableFlowDesignationEnum) Set(val *FlowDesignationEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowDesignationEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowDesignationEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowDesignationEnum(val *FlowDesignationEnum) *NullableFlowDesignationEnum {
	return &NullableFlowDesignationEnum{value: val, isSet: true}
}

func (v NullableFlowDesignationEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowDesignationEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
