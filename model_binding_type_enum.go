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

// BindingTypeEnum the model 'BindingTypeEnum'
type BindingTypeEnum string

// List of BindingTypeEnum
const (
	BINDINGTYPEENUM_REDIRECT  BindingTypeEnum = "REDIRECT"
	BINDINGTYPEENUM_POST      BindingTypeEnum = "POST"
	BINDINGTYPEENUM_POST_AUTO BindingTypeEnum = "POST_AUTO"
)

// All allowed values of BindingTypeEnum enum
var AllowedBindingTypeEnumEnumValues = []BindingTypeEnum{
	"REDIRECT",
	"POST",
	"POST_AUTO",
}

func (v *BindingTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BindingTypeEnum(value)
	for _, existing := range AllowedBindingTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BindingTypeEnum", value)
}

// NewBindingTypeEnumFromValue returns a pointer to a valid BindingTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBindingTypeEnumFromValue(v string) (*BindingTypeEnum, error) {
	ev := BindingTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BindingTypeEnum: valid values are %v", v, AllowedBindingTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BindingTypeEnum) IsValid() bool {
	for _, existing := range AllowedBindingTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BindingTypeEnum value
func (v BindingTypeEnum) Ptr() *BindingTypeEnum {
	return &v
}

type NullableBindingTypeEnum struct {
	value *BindingTypeEnum
	isSet bool
}

func (v NullableBindingTypeEnum) Get() *BindingTypeEnum {
	return v.value
}

func (v *NullableBindingTypeEnum) Set(val *BindingTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBindingTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBindingTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindingTypeEnum(val *BindingTypeEnum) *NullableBindingTypeEnum {
	return &NullableBindingTypeEnum{value: val, isSet: true}
}

func (v NullableBindingTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindingTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
