/*
authentik

Making authentication simple.

API version: 2024.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// OutpostTypeEnum the model 'OutpostTypeEnum'
type OutpostTypeEnum string

// List of OutpostTypeEnum
const (
	OUTPOSTTYPEENUM_PROXY  OutpostTypeEnum = "proxy"
	OUTPOSTTYPEENUM_LDAP   OutpostTypeEnum = "ldap"
	OUTPOSTTYPEENUM_RADIUS OutpostTypeEnum = "radius"
	OUTPOSTTYPEENUM_RAC    OutpostTypeEnum = "rac"
)

// All allowed values of OutpostTypeEnum enum
var AllowedOutpostTypeEnumEnumValues = []OutpostTypeEnum{
	"proxy",
	"ldap",
	"radius",
	"rac",
}

func (v *OutpostTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OutpostTypeEnum(value)
	for _, existing := range AllowedOutpostTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OutpostTypeEnum", value)
}

// NewOutpostTypeEnumFromValue returns a pointer to a valid OutpostTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOutpostTypeEnumFromValue(v string) (*OutpostTypeEnum, error) {
	ev := OutpostTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OutpostTypeEnum: valid values are %v", v, AllowedOutpostTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OutpostTypeEnum) IsValid() bool {
	for _, existing := range AllowedOutpostTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutpostTypeEnum value
func (v OutpostTypeEnum) Ptr() *OutpostTypeEnum {
	return &v
}

type NullableOutpostTypeEnum struct {
	value *OutpostTypeEnum
	isSet bool
}

func (v NullableOutpostTypeEnum) Get() *OutpostTypeEnum {
	return v.value
}

func (v *NullableOutpostTypeEnum) Set(val *OutpostTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOutpostTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOutpostTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutpostTypeEnum(val *OutpostTypeEnum) *NullableOutpostTypeEnum {
	return &NullableOutpostTypeEnum{value: val, isSet: true}
}

func (v NullableOutpostTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutpostTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
