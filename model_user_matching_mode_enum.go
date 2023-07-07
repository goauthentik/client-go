/*
authentik

Making authentication simple.

API version: 2023.6.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// UserMatchingModeEnum * `identifier` - Use the source-specific identifier * `email_link` - Link to a user with identical email address. Can have security implications when a source doesn't validate email addresses. * `email_deny` - Use the user's email address, but deny enrollment when the email address already exists. * `username_link` - Link to a user with identical username. Can have security implications when a username is used with another source. * `username_deny` - Use the user's username, but deny enrollment when the username already exists.
type UserMatchingModeEnum string

// List of UserMatchingModeEnum
const (
	USERMATCHINGMODEENUM_IDENTIFIER    UserMatchingModeEnum = "identifier"
	USERMATCHINGMODEENUM_EMAIL_LINK    UserMatchingModeEnum = "email_link"
	USERMATCHINGMODEENUM_EMAIL_DENY    UserMatchingModeEnum = "email_deny"
	USERMATCHINGMODEENUM_USERNAME_LINK UserMatchingModeEnum = "username_link"
	USERMATCHINGMODEENUM_USERNAME_DENY UserMatchingModeEnum = "username_deny"
)

// All allowed values of UserMatchingModeEnum enum
var AllowedUserMatchingModeEnumEnumValues = []UserMatchingModeEnum{
	"identifier",
	"email_link",
	"email_deny",
	"username_link",
	"username_deny",
}

func (v *UserMatchingModeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserMatchingModeEnum(value)
	for _, existing := range AllowedUserMatchingModeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserMatchingModeEnum", value)
}

// NewUserMatchingModeEnumFromValue returns a pointer to a valid UserMatchingModeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserMatchingModeEnumFromValue(v string) (*UserMatchingModeEnum, error) {
	ev := UserMatchingModeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserMatchingModeEnum: valid values are %v", v, AllowedUserMatchingModeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserMatchingModeEnum) IsValid() bool {
	for _, existing := range AllowedUserMatchingModeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserMatchingModeEnum value
func (v UserMatchingModeEnum) Ptr() *UserMatchingModeEnum {
	return &v
}

type NullableUserMatchingModeEnum struct {
	value *UserMatchingModeEnum
	isSet bool
}

func (v NullableUserMatchingModeEnum) Get() *UserMatchingModeEnum {
	return v.value
}

func (v *NullableUserMatchingModeEnum) Set(val *UserMatchingModeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMatchingModeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMatchingModeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMatchingModeEnum(val *UserMatchingModeEnum) *NullableUserMatchingModeEnum {
	return &NullableUserMatchingModeEnum{value: val, isSet: true}
}

func (v NullableUserMatchingModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMatchingModeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
