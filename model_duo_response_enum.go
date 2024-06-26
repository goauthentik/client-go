/*
authentik

Making authentication simple.

API version: 2024.6.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// DuoResponseEnum the model 'DuoResponseEnum'
type DuoResponseEnum string

// List of DuoResponseEnum
const (
	DUORESPONSEENUM_SUCCESS DuoResponseEnum = "success"
	DUORESPONSEENUM_WAITING DuoResponseEnum = "waiting"
	DUORESPONSEENUM_INVALID DuoResponseEnum = "invalid"
)

// All allowed values of DuoResponseEnum enum
var AllowedDuoResponseEnumEnumValues = []DuoResponseEnum{
	"success",
	"waiting",
	"invalid",
}

func (v *DuoResponseEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DuoResponseEnum(value)
	for _, existing := range AllowedDuoResponseEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DuoResponseEnum", value)
}

// NewDuoResponseEnumFromValue returns a pointer to a valid DuoResponseEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDuoResponseEnumFromValue(v string) (*DuoResponseEnum, error) {
	ev := DuoResponseEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DuoResponseEnum: valid values are %v", v, AllowedDuoResponseEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DuoResponseEnum) IsValid() bool {
	for _, existing := range AllowedDuoResponseEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DuoResponseEnum value
func (v DuoResponseEnum) Ptr() *DuoResponseEnum {
	return &v
}

type NullableDuoResponseEnum struct {
	value *DuoResponseEnum
	isSet bool
}

func (v NullableDuoResponseEnum) Get() *DuoResponseEnum {
	return v.value
}

func (v *NullableDuoResponseEnum) Set(val *DuoResponseEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDuoResponseEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDuoResponseEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDuoResponseEnum(val *DuoResponseEnum) *NullableDuoResponseEnum {
	return &NullableDuoResponseEnum{value: val, isSet: true}
}

func (v NullableDuoResponseEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDuoResponseEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
