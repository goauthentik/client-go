/*
authentik

Making authentication simple.

API version: 2022.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// DeniedActionEnum the model 'DeniedActionEnum'
type DeniedActionEnum string

// List of DeniedActionEnum
const (
	DENIEDACTIONENUM_MESSAGE_CONTINUE DeniedActionEnum = "message_continue"
	DENIEDACTIONENUM_MESSAGE          DeniedActionEnum = "message"
	DENIEDACTIONENUM_CONTINUE         DeniedActionEnum = "continue"
)

// All allowed values of DeniedActionEnum enum
var AllowedDeniedActionEnumEnumValues = []DeniedActionEnum{
	"message_continue",
	"message",
	"continue",
}

func (v *DeniedActionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeniedActionEnum(value)
	for _, existing := range AllowedDeniedActionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeniedActionEnum", value)
}

// NewDeniedActionEnumFromValue returns a pointer to a valid DeniedActionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeniedActionEnumFromValue(v string) (*DeniedActionEnum, error) {
	ev := DeniedActionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeniedActionEnum: valid values are %v", v, AllowedDeniedActionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeniedActionEnum) IsValid() bool {
	for _, existing := range AllowedDeniedActionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeniedActionEnum value
func (v DeniedActionEnum) Ptr() *DeniedActionEnum {
	return &v
}

type NullableDeniedActionEnum struct {
	value *DeniedActionEnum
	isSet bool
}

func (v NullableDeniedActionEnum) Get() *DeniedActionEnum {
	return v.value
}

func (v *NullableDeniedActionEnum) Set(val *DeniedActionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDeniedActionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDeniedActionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeniedActionEnum(val *DeniedActionEnum) *NullableDeniedActionEnum {
	return &NullableDeniedActionEnum{value: val, isSet: true}
}

func (v NullableDeniedActionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeniedActionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
