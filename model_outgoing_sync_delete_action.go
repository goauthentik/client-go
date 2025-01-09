/*
authentik

Making authentication simple.

API version: 2024.12.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// OutgoingSyncDeleteAction the model 'OutgoingSyncDeleteAction'
type OutgoingSyncDeleteAction string

// List of OutgoingSyncDeleteAction
const (
	OUTGOINGSYNCDELETEACTION_DO_NOTHING OutgoingSyncDeleteAction = "do_nothing"
	OUTGOINGSYNCDELETEACTION_DELETE     OutgoingSyncDeleteAction = "delete"
	OUTGOINGSYNCDELETEACTION_SUSPEND    OutgoingSyncDeleteAction = "suspend"
)

// All allowed values of OutgoingSyncDeleteAction enum
var AllowedOutgoingSyncDeleteActionEnumValues = []OutgoingSyncDeleteAction{
	"do_nothing",
	"delete",
	"suspend",
}

func (v *OutgoingSyncDeleteAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OutgoingSyncDeleteAction(value)
	for _, existing := range AllowedOutgoingSyncDeleteActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OutgoingSyncDeleteAction", value)
}

// NewOutgoingSyncDeleteActionFromValue returns a pointer to a valid OutgoingSyncDeleteAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOutgoingSyncDeleteActionFromValue(v string) (*OutgoingSyncDeleteAction, error) {
	ev := OutgoingSyncDeleteAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OutgoingSyncDeleteAction: valid values are %v", v, AllowedOutgoingSyncDeleteActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OutgoingSyncDeleteAction) IsValid() bool {
	for _, existing := range AllowedOutgoingSyncDeleteActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutgoingSyncDeleteAction value
func (v OutgoingSyncDeleteAction) Ptr() *OutgoingSyncDeleteAction {
	return &v
}

type NullableOutgoingSyncDeleteAction struct {
	value *OutgoingSyncDeleteAction
	isSet bool
}

func (v NullableOutgoingSyncDeleteAction) Get() *OutgoingSyncDeleteAction {
	return v.value
}

func (v *NullableOutgoingSyncDeleteAction) Set(val *OutgoingSyncDeleteAction) {
	v.value = val
	v.isSet = true
}

func (v NullableOutgoingSyncDeleteAction) IsSet() bool {
	return v.isSet
}

func (v *NullableOutgoingSyncDeleteAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutgoingSyncDeleteAction(val *OutgoingSyncDeleteAction) *NullableOutgoingSyncDeleteAction {
	return &NullableOutgoingSyncDeleteAction{value: val, isSet: true}
}

func (v NullableOutgoingSyncDeleteAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutgoingSyncDeleteAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
