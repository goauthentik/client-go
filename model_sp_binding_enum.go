/*
authentik

Making authentication simple.

API version: 2022.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// SpBindingEnum the model 'SpBindingEnum'
type SpBindingEnum string

// List of SpBindingEnum
const (
	SPBINDINGENUM_REDIRECT SpBindingEnum = "redirect"
	SPBINDINGENUM_POST     SpBindingEnum = "post"
)

// All allowed values of SpBindingEnum enum
var AllowedSpBindingEnumEnumValues = []SpBindingEnum{
	"redirect",
	"post",
}

func (v *SpBindingEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SpBindingEnum(value)
	for _, existing := range AllowedSpBindingEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SpBindingEnum", value)
}

// NewSpBindingEnumFromValue returns a pointer to a valid SpBindingEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSpBindingEnumFromValue(v string) (*SpBindingEnum, error) {
	ev := SpBindingEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SpBindingEnum: valid values are %v", v, AllowedSpBindingEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SpBindingEnum) IsValid() bool {
	for _, existing := range AllowedSpBindingEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpBindingEnum value
func (v SpBindingEnum) Ptr() *SpBindingEnum {
	return &v
}

type NullableSpBindingEnum struct {
	value *SpBindingEnum
	isSet bool
}

func (v NullableSpBindingEnum) Get() *SpBindingEnum {
	return v.value
}

func (v *NullableSpBindingEnum) Set(val *SpBindingEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSpBindingEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSpBindingEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpBindingEnum(val *SpBindingEnum) *NullableSpBindingEnum {
	return &NullableSpBindingEnum{value: val, isSet: true}
}

func (v NullableSpBindingEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpBindingEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
