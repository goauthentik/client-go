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

// DeviceClassesEnum the model 'DeviceClassesEnum'
type DeviceClassesEnum string

// List of DeviceClassesEnum
const (
	DEVICECLASSESENUM_STATIC   DeviceClassesEnum = "static"
	DEVICECLASSESENUM_TOTP     DeviceClassesEnum = "totp"
	DEVICECLASSESENUM_WEBAUTHN DeviceClassesEnum = "webauthn"
	DEVICECLASSESENUM_DUO      DeviceClassesEnum = "duo"
	DEVICECLASSESENUM_SMS      DeviceClassesEnum = "sms"
)

// All allowed values of DeviceClassesEnum enum
var AllowedDeviceClassesEnumEnumValues = []DeviceClassesEnum{
	"static",
	"totp",
	"webauthn",
	"duo",
	"sms",
}

func (v *DeviceClassesEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceClassesEnum(value)
	for _, existing := range AllowedDeviceClassesEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceClassesEnum", value)
}

// NewDeviceClassesEnumFromValue returns a pointer to a valid DeviceClassesEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceClassesEnumFromValue(v string) (*DeviceClassesEnum, error) {
	ev := DeviceClassesEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceClassesEnum: valid values are %v", v, AllowedDeviceClassesEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceClassesEnum) IsValid() bool {
	for _, existing := range AllowedDeviceClassesEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeviceClassesEnum value
func (v DeviceClassesEnum) Ptr() *DeviceClassesEnum {
	return &v
}

type NullableDeviceClassesEnum struct {
	value *DeviceClassesEnum
	isSet bool
}

func (v NullableDeviceClassesEnum) Get() *DeviceClassesEnum {
	return v.value
}

func (v *NullableDeviceClassesEnum) Set(val *DeviceClassesEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceClassesEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceClassesEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceClassesEnum(val *DeviceClassesEnum) *NullableDeviceClassesEnum {
	return &NullableDeviceClassesEnum{value: val, isSet: true}
}

func (v NullableDeviceClassesEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceClassesEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
