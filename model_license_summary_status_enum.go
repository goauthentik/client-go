/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// LicenseSummaryStatusEnum the model 'LicenseSummaryStatusEnum'
type LicenseSummaryStatusEnum string

// List of LicenseSummaryStatusEnum
const (
	LICENSESUMMARYSTATUSENUM_UNLICENSED           LicenseSummaryStatusEnum = "unlicensed"
	LICENSESUMMARYSTATUSENUM_VALID                LicenseSummaryStatusEnum = "valid"
	LICENSESUMMARYSTATUSENUM_EXPIRED              LicenseSummaryStatusEnum = "expired"
	LICENSESUMMARYSTATUSENUM_EXPIRY_SOON          LicenseSummaryStatusEnum = "expiry_soon"
	LICENSESUMMARYSTATUSENUM_LIMIT_EXCEEDED_ADMIN LicenseSummaryStatusEnum = "limit_exceeded_admin"
	LICENSESUMMARYSTATUSENUM_LIMIT_EXCEEDED_USER  LicenseSummaryStatusEnum = "limit_exceeded_user"
	LICENSESUMMARYSTATUSENUM_READ_ONLY            LicenseSummaryStatusEnum = "read_only"
)

// All allowed values of LicenseSummaryStatusEnum enum
var AllowedLicenseSummaryStatusEnumEnumValues = []LicenseSummaryStatusEnum{
	"unlicensed",
	"valid",
	"expired",
	"expiry_soon",
	"limit_exceeded_admin",
	"limit_exceeded_user",
	"read_only",
}

func (v *LicenseSummaryStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LicenseSummaryStatusEnum(value)
	for _, existing := range AllowedLicenseSummaryStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LicenseSummaryStatusEnum", value)
}

// NewLicenseSummaryStatusEnumFromValue returns a pointer to a valid LicenseSummaryStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLicenseSummaryStatusEnumFromValue(v string) (*LicenseSummaryStatusEnum, error) {
	ev := LicenseSummaryStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LicenseSummaryStatusEnum: valid values are %v", v, AllowedLicenseSummaryStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicenseSummaryStatusEnum) IsValid() bool {
	for _, existing := range AllowedLicenseSummaryStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LicenseSummaryStatusEnum value
func (v LicenseSummaryStatusEnum) Ptr() *LicenseSummaryStatusEnum {
	return &v
}

type NullableLicenseSummaryStatusEnum struct {
	value *LicenseSummaryStatusEnum
	isSet bool
}

func (v NullableLicenseSummaryStatusEnum) Get() *LicenseSummaryStatusEnum {
	return v.value
}

func (v *NullableLicenseSummaryStatusEnum) Set(val *LicenseSummaryStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseSummaryStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseSummaryStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseSummaryStatusEnum(val *LicenseSummaryStatusEnum) *NullableLicenseSummaryStatusEnum {
	return &NullableLicenseSummaryStatusEnum{value: val, isSet: true}
}

func (v NullableLicenseSummaryStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseSummaryStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
