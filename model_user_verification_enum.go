/*
authentik

Making authentication simple.

API version: 2024.4.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// UserVerificationEnum the model 'UserVerificationEnum'
type UserVerificationEnum string

// List of UserVerificationEnum
const (
	USERVERIFICATIONENUM_REQUIRED    UserVerificationEnum = "required"
	USERVERIFICATIONENUM_PREFERRED   UserVerificationEnum = "preferred"
	USERVERIFICATIONENUM_DISCOURAGED UserVerificationEnum = "discouraged"
)

// All allowed values of UserVerificationEnum enum
var AllowedUserVerificationEnumEnumValues = []UserVerificationEnum{
	"required",
	"preferred",
	"discouraged",
}

func (v *UserVerificationEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserVerificationEnum(value)
	for _, existing := range AllowedUserVerificationEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserVerificationEnum", value)
}

// NewUserVerificationEnumFromValue returns a pointer to a valid UserVerificationEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserVerificationEnumFromValue(v string) (*UserVerificationEnum, error) {
	ev := UserVerificationEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserVerificationEnum: valid values are %v", v, AllowedUserVerificationEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserVerificationEnum) IsValid() bool {
	for _, existing := range AllowedUserVerificationEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserVerificationEnum value
func (v UserVerificationEnum) Ptr() *UserVerificationEnum {
	return &v
}

type NullableUserVerificationEnum struct {
	value *UserVerificationEnum
	isSet bool
}

func (v NullableUserVerificationEnum) Get() *UserVerificationEnum {
	return v.value
}

func (v *NullableUserVerificationEnum) Set(val *UserVerificationEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUserVerificationEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUserVerificationEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserVerificationEnum(val *UserVerificationEnum) *NullableUserVerificationEnum {
	return &NullableUserVerificationEnum{value: val, isSet: true}
}

func (v NullableUserVerificationEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserVerificationEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
