/*
authentik

Making authentication simple.

API version: 2023.10.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// UiThemeEnum * `automatic` - Automatic * `light` - Light * `dark` - Dark
type UiThemeEnum string

// List of UiThemeEnum
const (
	UITHEMEENUM_AUTOMATIC UiThemeEnum = "automatic"
	UITHEMEENUM_LIGHT     UiThemeEnum = "light"
	UITHEMEENUM_DARK      UiThemeEnum = "dark"
)

// All allowed values of UiThemeEnum enum
var AllowedUiThemeEnumEnumValues = []UiThemeEnum{
	"automatic",
	"light",
	"dark",
}

func (v *UiThemeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UiThemeEnum(value)
	for _, existing := range AllowedUiThemeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UiThemeEnum", value)
}

// NewUiThemeEnumFromValue returns a pointer to a valid UiThemeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUiThemeEnumFromValue(v string) (*UiThemeEnum, error) {
	ev := UiThemeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UiThemeEnum: valid values are %v", v, AllowedUiThemeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UiThemeEnum) IsValid() bool {
	for _, existing := range AllowedUiThemeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UiThemeEnum value
func (v UiThemeEnum) Ptr() *UiThemeEnum {
	return &v
}

type NullableUiThemeEnum struct {
	value *UiThemeEnum
	isSet bool
}

func (v NullableUiThemeEnum) Get() *UiThemeEnum {
	return v.value
}

func (v *NullableUiThemeEnum) Set(val *UiThemeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUiThemeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUiThemeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiThemeEnum(val *UiThemeEnum) *NullableUiThemeEnum {
	return &NullableUiThemeEnum{value: val, isSet: true}
}

func (v NullableUiThemeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiThemeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
