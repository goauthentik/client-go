/*
authentik

Making authentication simple.

API version: 2023.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// AuthenticationEnum the model 'AuthenticationEnum'
type AuthenticationEnum string

// List of AuthenticationEnum
const (
	AUTHENTICATIONENUM_NONE                    AuthenticationEnum = "none"
	AUTHENTICATIONENUM_REQUIRE_AUTHENTICATED   AuthenticationEnum = "require_authenticated"
	AUTHENTICATIONENUM_REQUIRE_UNAUTHENTICATED AuthenticationEnum = "require_unauthenticated"
	AUTHENTICATIONENUM_REQUIRE_SUPERUSER       AuthenticationEnum = "require_superuser"
)

// All allowed values of AuthenticationEnum enum
var AllowedAuthenticationEnumEnumValues = []AuthenticationEnum{
	"none",
	"require_authenticated",
	"require_unauthenticated",
	"require_superuser",
}

func (v *AuthenticationEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthenticationEnum(value)
	for _, existing := range AllowedAuthenticationEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthenticationEnum", value)
}

// NewAuthenticationEnumFromValue returns a pointer to a valid AuthenticationEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthenticationEnumFromValue(v string) (*AuthenticationEnum, error) {
	ev := AuthenticationEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthenticationEnum: valid values are %v", v, AllowedAuthenticationEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthenticationEnum) IsValid() bool {
	for _, existing := range AllowedAuthenticationEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthenticationEnum value
func (v AuthenticationEnum) Ptr() *AuthenticationEnum {
	return &v
}

type NullableAuthenticationEnum struct {
	value *AuthenticationEnum
	isSet bool
}

func (v NullableAuthenticationEnum) Get() *AuthenticationEnum {
	return v.value
}

func (v *NullableAuthenticationEnum) Set(val *AuthenticationEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationEnum(val *AuthenticationEnum) *NullableAuthenticationEnum {
	return &NullableAuthenticationEnum{value: val, isSet: true}
}

func (v NullableAuthenticationEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
