/*
authentik

Making authentication simple.

API version: 2023.10.6
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// BackendsEnum * `authentik.core.auth.InbuiltBackend` - User database + standard password * `authentik.core.auth.TokenBackend` - User database + app passwords * `authentik.sources.ldap.auth.LDAPBackend` - User database + LDAP password
type BackendsEnum string

// List of BackendsEnum
const (
	BACKENDSENUM_CORE_AUTH_INBUILT_BACKEND      BackendsEnum = "authentik.core.auth.InbuiltBackend"
	BACKENDSENUM_CORE_AUTH_TOKEN_BACKEND        BackendsEnum = "authentik.core.auth.TokenBackend"
	BACKENDSENUM_SOURCES_LDAP_AUTH_LDAP_BACKEND BackendsEnum = "authentik.sources.ldap.auth.LDAPBackend"
)

// All allowed values of BackendsEnum enum
var AllowedBackendsEnumEnumValues = []BackendsEnum{
	"authentik.core.auth.InbuiltBackend",
	"authentik.core.auth.TokenBackend",
	"authentik.sources.ldap.auth.LDAPBackend",
}

func (v *BackendsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackendsEnum(value)
	for _, existing := range AllowedBackendsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackendsEnum", value)
}

// NewBackendsEnumFromValue returns a pointer to a valid BackendsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackendsEnumFromValue(v string) (*BackendsEnum, error) {
	ev := BackendsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackendsEnum: valid values are %v", v, AllowedBackendsEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackendsEnum) IsValid() bool {
	for _, existing := range AllowedBackendsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackendsEnum value
func (v BackendsEnum) Ptr() *BackendsEnum {
	return &v
}

type NullableBackendsEnum struct {
	value *BackendsEnum
	isSet bool
}

func (v NullableBackendsEnum) Get() *BackendsEnum {
	return v.value
}

func (v *NullableBackendsEnum) Set(val *BackendsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBackendsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBackendsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackendsEnum(val *BackendsEnum) *NullableBackendsEnum {
	return &NullableBackendsEnum{value: val, isSet: true}
}

func (v NullableBackendsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackendsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
