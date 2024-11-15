/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// AuthenticatorAttachmentEnum the model 'AuthenticatorAttachmentEnum'
type AuthenticatorAttachmentEnum string

// List of AuthenticatorAttachmentEnum
const (
	AUTHENTICATORATTACHMENTENUM_PLATFORM       AuthenticatorAttachmentEnum = "platform"
	AUTHENTICATORATTACHMENTENUM_CROSS_PLATFORM AuthenticatorAttachmentEnum = "cross-platform"
)

// All allowed values of AuthenticatorAttachmentEnum enum
var AllowedAuthenticatorAttachmentEnumEnumValues = []AuthenticatorAttachmentEnum{
	"platform",
	"cross-platform",
}

func (v *AuthenticatorAttachmentEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthenticatorAttachmentEnum(value)
	for _, existing := range AllowedAuthenticatorAttachmentEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthenticatorAttachmentEnum", value)
}

// NewAuthenticatorAttachmentEnumFromValue returns a pointer to a valid AuthenticatorAttachmentEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthenticatorAttachmentEnumFromValue(v string) (*AuthenticatorAttachmentEnum, error) {
	ev := AuthenticatorAttachmentEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthenticatorAttachmentEnum: valid values are %v", v, AllowedAuthenticatorAttachmentEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthenticatorAttachmentEnum) IsValid() bool {
	for _, existing := range AllowedAuthenticatorAttachmentEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthenticatorAttachmentEnum value
func (v AuthenticatorAttachmentEnum) Ptr() *AuthenticatorAttachmentEnum {
	return &v
}

type NullableAuthenticatorAttachmentEnum struct {
	value *AuthenticatorAttachmentEnum
	isSet bool
}

func (v NullableAuthenticatorAttachmentEnum) Get() *AuthenticatorAttachmentEnum {
	return v.value
}

func (v *NullableAuthenticatorAttachmentEnum) Set(val *AuthenticatorAttachmentEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorAttachmentEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorAttachmentEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorAttachmentEnum(val *AuthenticatorAttachmentEnum) *NullableAuthenticatorAttachmentEnum {
	return &NullableAuthenticatorAttachmentEnum{value: val, isSet: true}
}

func (v NullableAuthenticatorAttachmentEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorAttachmentEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
