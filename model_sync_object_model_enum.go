/*
authentik

Making authentication simple.

API version: 2025.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// SyncObjectModelEnum the model 'SyncObjectModelEnum'
type SyncObjectModelEnum string

// List of SyncObjectModelEnum
const (
	SYNCOBJECTMODELENUM_USER  SyncObjectModelEnum = "authentik.core.models.User"
	SYNCOBJECTMODELENUM_GROUP SyncObjectModelEnum = "authentik.core.models.Group"
)

// All allowed values of SyncObjectModelEnum enum
var AllowedSyncObjectModelEnumEnumValues = []SyncObjectModelEnum{
	"authentik.core.models.User",
	"authentik.core.models.Group",
}

func (v *SyncObjectModelEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyncObjectModelEnum(value)
	for _, existing := range AllowedSyncObjectModelEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyncObjectModelEnum", value)
}

// NewSyncObjectModelEnumFromValue returns a pointer to a valid SyncObjectModelEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyncObjectModelEnumFromValue(v string) (*SyncObjectModelEnum, error) {
	ev := SyncObjectModelEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyncObjectModelEnum: valid values are %v", v, AllowedSyncObjectModelEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyncObjectModelEnum) IsValid() bool {
	for _, existing := range AllowedSyncObjectModelEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyncObjectModelEnum value
func (v SyncObjectModelEnum) Ptr() *SyncObjectModelEnum {
	return &v
}

type NullableSyncObjectModelEnum struct {
	value *SyncObjectModelEnum
	isSet bool
}

func (v NullableSyncObjectModelEnum) Get() *SyncObjectModelEnum {
	return v.value
}

func (v *NullableSyncObjectModelEnum) Set(val *SyncObjectModelEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncObjectModelEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncObjectModelEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncObjectModelEnum(val *SyncObjectModelEnum) *NullableSyncObjectModelEnum {
	return &NullableSyncObjectModelEnum{value: val, isSet: true}
}

func (v NullableSyncObjectModelEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncObjectModelEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
