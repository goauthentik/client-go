/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// SubModeEnum the model 'SubModeEnum'
type SubModeEnum string

// List of SubModeEnum
const (
	SUBMODEENUM_HASHED_USER_ID SubModeEnum = "hashed_user_id"
	SUBMODEENUM_USER_ID        SubModeEnum = "user_id"
	SUBMODEENUM_USER_UUID      SubModeEnum = "user_uuid"
	SUBMODEENUM_USER_USERNAME  SubModeEnum = "user_username"
	SUBMODEENUM_USER_EMAIL     SubModeEnum = "user_email"
	SUBMODEENUM_USER_UPN       SubModeEnum = "user_upn"
)

// All allowed values of SubModeEnum enum
var AllowedSubModeEnumEnumValues = []SubModeEnum{
	"hashed_user_id",
	"user_id",
	"user_uuid",
	"user_username",
	"user_email",
	"user_upn",
}

func (v *SubModeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubModeEnum(value)
	for _, existing := range AllowedSubModeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubModeEnum", value)
}

// NewSubModeEnumFromValue returns a pointer to a valid SubModeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubModeEnumFromValue(v string) (*SubModeEnum, error) {
	ev := SubModeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubModeEnum: valid values are %v", v, AllowedSubModeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubModeEnum) IsValid() bool {
	for _, existing := range AllowedSubModeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubModeEnum value
func (v SubModeEnum) Ptr() *SubModeEnum {
	return &v
}

type NullableSubModeEnum struct {
	value *SubModeEnum
	isSet bool
}

func (v NullableSubModeEnum) Get() *SubModeEnum {
	return v.value
}

func (v *NullableSubModeEnum) Set(val *SubModeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSubModeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSubModeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubModeEnum(val *SubModeEnum) *NullableSubModeEnum {
	return &NullableSubModeEnum{value: val, isSet: true}
}

func (v NullableSubModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubModeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
