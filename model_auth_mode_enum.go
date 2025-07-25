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

// AuthModeEnum the model 'AuthModeEnum'
type AuthModeEnum string

// List of AuthModeEnum
const (
	AUTHMODEENUM_STATIC AuthModeEnum = "static"
	AUTHMODEENUM_PROMPT AuthModeEnum = "prompt"
)

// All allowed values of AuthModeEnum enum
var AllowedAuthModeEnumEnumValues = []AuthModeEnum{
	"static",
	"prompt",
}

func (v *AuthModeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthModeEnum(value)
	for _, existing := range AllowedAuthModeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthModeEnum", value)
}

// NewAuthModeEnumFromValue returns a pointer to a valid AuthModeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthModeEnumFromValue(v string) (*AuthModeEnum, error) {
	ev := AuthModeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthModeEnum: valid values are %v", v, AllowedAuthModeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthModeEnum) IsValid() bool {
	for _, existing := range AllowedAuthModeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthModeEnum value
func (v AuthModeEnum) Ptr() *AuthModeEnum {
	return &v
}

type NullableAuthModeEnum struct {
	value *AuthModeEnum
	isSet bool
}

func (v NullableAuthModeEnum) Get() *AuthModeEnum {
	return v.value
}

func (v *NullableAuthModeEnum) Set(val *AuthModeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthModeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthModeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthModeEnum(val *AuthModeEnum) *NullableAuthModeEnum {
	return &NullableAuthModeEnum{value: val, isSet: true}
}

func (v NullableAuthModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthModeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
