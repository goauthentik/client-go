/*
authentik

Making authentication simple.

API version: 2024.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// UserCreationModeEnum the model 'UserCreationModeEnum'
type UserCreationModeEnum string

// List of UserCreationModeEnum
const (
	USERCREATIONMODEENUM_NEVER_CREATE         UserCreationModeEnum = "never_create"
	USERCREATIONMODEENUM_CREATE_WHEN_REQUIRED UserCreationModeEnum = "create_when_required"
	USERCREATIONMODEENUM_ALWAYS_CREATE        UserCreationModeEnum = "always_create"
)

// All allowed values of UserCreationModeEnum enum
var AllowedUserCreationModeEnumEnumValues = []UserCreationModeEnum{
	"never_create",
	"create_when_required",
	"always_create",
}

func (v *UserCreationModeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserCreationModeEnum(value)
	for _, existing := range AllowedUserCreationModeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserCreationModeEnum", value)
}

// NewUserCreationModeEnumFromValue returns a pointer to a valid UserCreationModeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserCreationModeEnumFromValue(v string) (*UserCreationModeEnum, error) {
	ev := UserCreationModeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserCreationModeEnum: valid values are %v", v, AllowedUserCreationModeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserCreationModeEnum) IsValid() bool {
	for _, existing := range AllowedUserCreationModeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserCreationModeEnum value
func (v UserCreationModeEnum) Ptr() *UserCreationModeEnum {
	return &v
}

type NullableUserCreationModeEnum struct {
	value *UserCreationModeEnum
	isSet bool
}

func (v NullableUserCreationModeEnum) Get() *UserCreationModeEnum {
	return v.value
}

func (v *NullableUserCreationModeEnum) Set(val *UserCreationModeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreationModeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreationModeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreationModeEnum(val *UserCreationModeEnum) *NullableUserCreationModeEnum {
	return &NullableUserCreationModeEnum{value: val, isSet: true}
}

func (v NullableUserCreationModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreationModeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
