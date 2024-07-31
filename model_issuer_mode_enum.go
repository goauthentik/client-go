/*
authentik

Making authentication simple.

API version: 2024.6.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// IssuerModeEnum the model 'IssuerModeEnum'
type IssuerModeEnum string

// List of IssuerModeEnum
const (
	ISSUERMODEENUM_GLOBAL       IssuerModeEnum = "global"
	ISSUERMODEENUM_PER_PROVIDER IssuerModeEnum = "per_provider"
)

// All allowed values of IssuerModeEnum enum
var AllowedIssuerModeEnumEnumValues = []IssuerModeEnum{
	"global",
	"per_provider",
}

func (v *IssuerModeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IssuerModeEnum(value)
	for _, existing := range AllowedIssuerModeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IssuerModeEnum", value)
}

// NewIssuerModeEnumFromValue returns a pointer to a valid IssuerModeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIssuerModeEnumFromValue(v string) (*IssuerModeEnum, error) {
	ev := IssuerModeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IssuerModeEnum: valid values are %v", v, AllowedIssuerModeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IssuerModeEnum) IsValid() bool {
	for _, existing := range AllowedIssuerModeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IssuerModeEnum value
func (v IssuerModeEnum) Ptr() *IssuerModeEnum {
	return &v
}

type NullableIssuerModeEnum struct {
	value *IssuerModeEnum
	isSet bool
}

func (v NullableIssuerModeEnum) Get() *IssuerModeEnum {
	return v.value
}

func (v *NullableIssuerModeEnum) Set(val *IssuerModeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuerModeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuerModeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuerModeEnum(val *IssuerModeEnum) *NullableIssuerModeEnum {
	return &NullableIssuerModeEnum{value: val, isSet: true}
}

func (v NullableIssuerModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuerModeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
