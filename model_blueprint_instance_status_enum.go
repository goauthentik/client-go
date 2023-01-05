/*
authentik

Making authentication simple.

API version: 2022.12.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// BlueprintInstanceStatusEnum the model 'BlueprintInstanceStatusEnum'
type BlueprintInstanceStatusEnum string

// List of BlueprintInstanceStatusEnum
const (
	BLUEPRINTINSTANCESTATUSENUM_SUCCESSFUL BlueprintInstanceStatusEnum = "successful"
	BLUEPRINTINSTANCESTATUSENUM_WARNING    BlueprintInstanceStatusEnum = "warning"
	BLUEPRINTINSTANCESTATUSENUM_ERROR      BlueprintInstanceStatusEnum = "error"
	BLUEPRINTINSTANCESTATUSENUM_ORPHANED   BlueprintInstanceStatusEnum = "orphaned"
	BLUEPRINTINSTANCESTATUSENUM_UNKNOWN    BlueprintInstanceStatusEnum = "unknown"
)

// All allowed values of BlueprintInstanceStatusEnum enum
var AllowedBlueprintInstanceStatusEnumEnumValues = []BlueprintInstanceStatusEnum{
	"successful",
	"warning",
	"error",
	"orphaned",
	"unknown",
}

func (v *BlueprintInstanceStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BlueprintInstanceStatusEnum(value)
	for _, existing := range AllowedBlueprintInstanceStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BlueprintInstanceStatusEnum", value)
}

// NewBlueprintInstanceStatusEnumFromValue returns a pointer to a valid BlueprintInstanceStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBlueprintInstanceStatusEnumFromValue(v string) (*BlueprintInstanceStatusEnum, error) {
	ev := BlueprintInstanceStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BlueprintInstanceStatusEnum: valid values are %v", v, AllowedBlueprintInstanceStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BlueprintInstanceStatusEnum) IsValid() bool {
	for _, existing := range AllowedBlueprintInstanceStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BlueprintInstanceStatusEnum value
func (v BlueprintInstanceStatusEnum) Ptr() *BlueprintInstanceStatusEnum {
	return &v
}

type NullableBlueprintInstanceStatusEnum struct {
	value *BlueprintInstanceStatusEnum
	isSet bool
}

func (v NullableBlueprintInstanceStatusEnum) Get() *BlueprintInstanceStatusEnum {
	return v.value
}

func (v *NullableBlueprintInstanceStatusEnum) Set(val *BlueprintInstanceStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintInstanceStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintInstanceStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintInstanceStatusEnum(val *BlueprintInstanceStatusEnum) *NullableBlueprintInstanceStatusEnum {
	return &NullableBlueprintInstanceStatusEnum{value: val, isSet: true}
}

func (v NullableBlueprintInstanceStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintInstanceStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
