/*
authentik

Making authentication simple.

API version: 2024.4.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// AlgEnum the model 'AlgEnum'
type AlgEnum string

// List of AlgEnum
const (
	ALGENUM_RSA   AlgEnum = "rsa"
	ALGENUM_ECDSA AlgEnum = "ecdsa"
)

// All allowed values of AlgEnum enum
var AllowedAlgEnumEnumValues = []AlgEnum{
	"rsa",
	"ecdsa",
}

func (v *AlgEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlgEnum(value)
	for _, existing := range AllowedAlgEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlgEnum", value)
}

// NewAlgEnumFromValue returns a pointer to a valid AlgEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlgEnumFromValue(v string) (*AlgEnum, error) {
	ev := AlgEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlgEnum: valid values are %v", v, AllowedAlgEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlgEnum) IsValid() bool {
	for _, existing := range AllowedAlgEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlgEnum value
func (v AlgEnum) Ptr() *AlgEnum {
	return &v
}

type NullableAlgEnum struct {
	value *AlgEnum
	isSet bool
}

func (v NullableAlgEnum) Get() *AlgEnum {
	return v.value
}

func (v *NullableAlgEnum) Set(val *AlgEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgEnum(val *AlgEnum) *NullableAlgEnum {
	return &NullableAlgEnum{value: val, isSet: true}
}

func (v NullableAlgEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
