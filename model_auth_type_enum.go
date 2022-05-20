/*
authentik

Making authentication simple.

API version: 2022.5.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// AuthTypeEnum the model 'AuthTypeEnum'
type AuthTypeEnum string

// List of AuthTypeEnum
const (
	AUTHTYPEENUM_BASIC  AuthTypeEnum = "basic"
	AUTHTYPEENUM_BEARER AuthTypeEnum = "bearer"
)

var allowedAuthTypeEnumEnumValues = []AuthTypeEnum{
	"basic",
	"bearer",
}

func (v *AuthTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthTypeEnum(value)
	for _, existing := range allowedAuthTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthTypeEnum", value)
}

// NewAuthTypeEnumFromValue returns a pointer to a valid AuthTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthTypeEnumFromValue(v string) (*AuthTypeEnum, error) {
	ev := AuthTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthTypeEnum: valid values are %v", v, allowedAuthTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthTypeEnum) IsValid() bool {
	for _, existing := range allowedAuthTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthTypeEnum value
func (v AuthTypeEnum) Ptr() *AuthTypeEnum {
	return &v
}

type NullableAuthTypeEnum struct {
	value *AuthTypeEnum
	isSet bool
}

func (v NullableAuthTypeEnum) Get() *AuthTypeEnum {
	return v.value
}

func (v *NullableAuthTypeEnum) Set(val *AuthTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTypeEnum(val *AuthTypeEnum) *NullableAuthTypeEnum {
	return &NullableAuthTypeEnum{value: val, isSet: true}
}

func (v NullableAuthTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
