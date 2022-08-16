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

// UserFieldsEnum the model 'UserFieldsEnum'
type UserFieldsEnum string

// List of UserFieldsEnum
const (
	USERFIELDSENUM_EMAIL    UserFieldsEnum = "email"
	USERFIELDSENUM_USERNAME UserFieldsEnum = "username"
	USERFIELDSENUM_UPN      UserFieldsEnum = "upn"
)

// All allowed values of UserFieldsEnum enum
var AllowedUserFieldsEnumEnumValues = []UserFieldsEnum{
	"email",
	"username",
	"upn",
}

func (v *UserFieldsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserFieldsEnum(value)
	for _, existing := range AllowedUserFieldsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserFieldsEnum", value)
}

// NewUserFieldsEnumFromValue returns a pointer to a valid UserFieldsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserFieldsEnumFromValue(v string) (*UserFieldsEnum, error) {
	ev := UserFieldsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserFieldsEnum: valid values are %v", v, AllowedUserFieldsEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserFieldsEnum) IsValid() bool {
	for _, existing := range AllowedUserFieldsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserFieldsEnum value
func (v UserFieldsEnum) Ptr() *UserFieldsEnum {
	return &v
}

type NullableUserFieldsEnum struct {
	value *UserFieldsEnum
	isSet bool
}

func (v NullableUserFieldsEnum) Get() *UserFieldsEnum {
	return v.value
}

func (v *NullableUserFieldsEnum) Set(val *UserFieldsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFieldsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFieldsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFieldsEnum(val *UserFieldsEnum) *NullableUserFieldsEnum {
	return &NullableUserFieldsEnum{value: val, isSet: true}
}

func (v NullableUserFieldsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFieldsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
