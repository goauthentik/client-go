/*
authentik

Making authentication simple.

API version: 2022.11.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ResidentKeyRequirementEnum the model 'ResidentKeyRequirementEnum'
type ResidentKeyRequirementEnum string

// List of ResidentKeyRequirementEnum
const (
	RESIDENTKEYREQUIREMENTENUM_DISCOURAGED ResidentKeyRequirementEnum = "discouraged"
	RESIDENTKEYREQUIREMENTENUM_PREFERRED   ResidentKeyRequirementEnum = "preferred"
	RESIDENTKEYREQUIREMENTENUM_REQUIRED    ResidentKeyRequirementEnum = "required"
)

// All allowed values of ResidentKeyRequirementEnum enum
var AllowedResidentKeyRequirementEnumEnumValues = []ResidentKeyRequirementEnum{
	"discouraged",
	"preferred",
	"required",
}

func (v *ResidentKeyRequirementEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResidentKeyRequirementEnum(value)
	for _, existing := range AllowedResidentKeyRequirementEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResidentKeyRequirementEnum", value)
}

// NewResidentKeyRequirementEnumFromValue returns a pointer to a valid ResidentKeyRequirementEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResidentKeyRequirementEnumFromValue(v string) (*ResidentKeyRequirementEnum, error) {
	ev := ResidentKeyRequirementEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResidentKeyRequirementEnum: valid values are %v", v, AllowedResidentKeyRequirementEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResidentKeyRequirementEnum) IsValid() bool {
	for _, existing := range AllowedResidentKeyRequirementEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResidentKeyRequirementEnum value
func (v ResidentKeyRequirementEnum) Ptr() *ResidentKeyRequirementEnum {
	return &v
}

type NullableResidentKeyRequirementEnum struct {
	value *ResidentKeyRequirementEnum
	isSet bool
}

func (v NullableResidentKeyRequirementEnum) Get() *ResidentKeyRequirementEnum {
	return v.value
}

func (v *NullableResidentKeyRequirementEnum) Set(val *ResidentKeyRequirementEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableResidentKeyRequirementEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableResidentKeyRequirementEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResidentKeyRequirementEnum(val *ResidentKeyRequirementEnum) *NullableResidentKeyRequirementEnum {
	return &NullableResidentKeyRequirementEnum{value: val, isSet: true}
}

func (v NullableResidentKeyRequirementEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResidentKeyRequirementEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
