/*
authentik

Making authentication simple.

API version: 2024.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// LicenseFlagsEnum the model 'LicenseFlagsEnum'
type LicenseFlagsEnum string

// List of LicenseFlagsEnum
const (
	LICENSEFLAGSENUM_TRIAL          LicenseFlagsEnum = "trial"
	LICENSEFLAGSENUM_NON_PRODUCTION LicenseFlagsEnum = "non_production"
)

// All allowed values of LicenseFlagsEnum enum
var AllowedLicenseFlagsEnumEnumValues = []LicenseFlagsEnum{
	"trial",
	"non_production",
}

func (v *LicenseFlagsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LicenseFlagsEnum(value)
	for _, existing := range AllowedLicenseFlagsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LicenseFlagsEnum", value)
}

// NewLicenseFlagsEnumFromValue returns a pointer to a valid LicenseFlagsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLicenseFlagsEnumFromValue(v string) (*LicenseFlagsEnum, error) {
	ev := LicenseFlagsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LicenseFlagsEnum: valid values are %v", v, AllowedLicenseFlagsEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicenseFlagsEnum) IsValid() bool {
	for _, existing := range AllowedLicenseFlagsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LicenseFlagsEnum value
func (v LicenseFlagsEnum) Ptr() *LicenseFlagsEnum {
	return &v
}

type NullableLicenseFlagsEnum struct {
	value *LicenseFlagsEnum
	isSet bool
}

func (v NullableLicenseFlagsEnum) Get() *LicenseFlagsEnum {
	return v.value
}

func (v *NullableLicenseFlagsEnum) Set(val *LicenseFlagsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseFlagsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseFlagsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseFlagsEnum(val *LicenseFlagsEnum) *NullableLicenseFlagsEnum {
	return &NullableLicenseFlagsEnum{value: val, isSet: true}
}

func (v NullableLicenseFlagsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseFlagsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
