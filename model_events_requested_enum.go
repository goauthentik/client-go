/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// EventsRequestedEnum the model 'EventsRequestedEnum'
type EventsRequestedEnum string

// List of EventsRequestedEnum
const (
	EVENTSREQUESTEDENUM_CAEP_EVENT_TYPE_SESSION_REVOKED   EventsRequestedEnum = "https://schemas.openid.net/secevent/caep/event-type/session-revoked"
	EVENTSREQUESTEDENUM_CAEP_EVENT_TYPE_CREDENTIAL_CHANGE EventsRequestedEnum = "https://schemas.openid.net/secevent/caep/event-type/credential-change"
	EVENTSREQUESTEDENUM_SSF_EVENT_TYPE_VERIFICATION       EventsRequestedEnum = "https://schemas.openid.net/secevent/ssf/event-type/verification"
)

// All allowed values of EventsRequestedEnum enum
var AllowedEventsRequestedEnumEnumValues = []EventsRequestedEnum{
	"https://schemas.openid.net/secevent/caep/event-type/session-revoked",
	"https://schemas.openid.net/secevent/caep/event-type/credential-change",
	"https://schemas.openid.net/secevent/ssf/event-type/verification",
}

func (v *EventsRequestedEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventsRequestedEnum(value)
	for _, existing := range AllowedEventsRequestedEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventsRequestedEnum", value)
}

// NewEventsRequestedEnumFromValue returns a pointer to a valid EventsRequestedEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventsRequestedEnumFromValue(v string) (*EventsRequestedEnum, error) {
	ev := EventsRequestedEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventsRequestedEnum: valid values are %v", v, AllowedEventsRequestedEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventsRequestedEnum) IsValid() bool {
	for _, existing := range AllowedEventsRequestedEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventsRequestedEnum value
func (v EventsRequestedEnum) Ptr() *EventsRequestedEnum {
	return &v
}

type NullableEventsRequestedEnum struct {
	value *EventsRequestedEnum
	isSet bool
}

func (v NullableEventsRequestedEnum) Get() *EventsRequestedEnum {
	return v.value
}

func (v *NullableEventsRequestedEnum) Set(val *EventsRequestedEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsRequestedEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsRequestedEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsRequestedEnum(val *EventsRequestedEnum) *NullableEventsRequestedEnum {
	return &NullableEventsRequestedEnum{value: val, isSet: true}
}

func (v NullableEventsRequestedEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsRequestedEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
