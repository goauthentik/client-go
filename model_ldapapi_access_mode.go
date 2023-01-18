/*
authentik

Making authentication simple.

API version: 2023.1.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// LDAPAPIAccessMode the model 'LDAPAPIAccessMode'
type LDAPAPIAccessMode string

// List of LDAPAPIAccessMode
const (
	LDAPAPIACCESSMODE_DIRECT LDAPAPIAccessMode = "direct"
	LDAPAPIACCESSMODE_CACHED LDAPAPIAccessMode = "cached"
)

// All allowed values of LDAPAPIAccessMode enum
var AllowedLDAPAPIAccessModeEnumValues = []LDAPAPIAccessMode{
	"direct",
	"cached",
}

func (v *LDAPAPIAccessMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LDAPAPIAccessMode(value)
	for _, existing := range AllowedLDAPAPIAccessModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LDAPAPIAccessMode", value)
}

// NewLDAPAPIAccessModeFromValue returns a pointer to a valid LDAPAPIAccessMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLDAPAPIAccessModeFromValue(v string) (*LDAPAPIAccessMode, error) {
	ev := LDAPAPIAccessMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LDAPAPIAccessMode: valid values are %v", v, AllowedLDAPAPIAccessModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LDAPAPIAccessMode) IsValid() bool {
	for _, existing := range AllowedLDAPAPIAccessModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LDAPAPIAccessMode value
func (v LDAPAPIAccessMode) Ptr() *LDAPAPIAccessMode {
	return &v
}

type NullableLDAPAPIAccessMode struct {
	value *LDAPAPIAccessMode
	isSet bool
}

func (v NullableLDAPAPIAccessMode) Get() *LDAPAPIAccessMode {
	return v.value
}

func (v *NullableLDAPAPIAccessMode) Set(val *LDAPAPIAccessMode) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPAPIAccessMode) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPAPIAccessMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPAPIAccessMode(val *LDAPAPIAccessMode) *NullableLDAPAPIAccessMode {
	return &NullableLDAPAPIAccessMode{value: val, isSet: true}
}

func (v NullableLDAPAPIAccessMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPAPIAccessMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
