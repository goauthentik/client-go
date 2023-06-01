/*
authentik

Making authentication simple.

API version: 2023.5.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// NotificationTransportModeEnum * `local` - authentik inbuilt notifications * `webhook` - Generic Webhook * `webhook_slack` - Slack Webhook (Slack/Discord) * `email` - Email
type NotificationTransportModeEnum string

// List of NotificationTransportModeEnum
const (
	NOTIFICATIONTRANSPORTMODEENUM_LOCAL         NotificationTransportModeEnum = "local"
	NOTIFICATIONTRANSPORTMODEENUM_WEBHOOK       NotificationTransportModeEnum = "webhook"
	NOTIFICATIONTRANSPORTMODEENUM_WEBHOOK_SLACK NotificationTransportModeEnum = "webhook_slack"
	NOTIFICATIONTRANSPORTMODEENUM_EMAIL         NotificationTransportModeEnum = "email"
)

// All allowed values of NotificationTransportModeEnum enum
var AllowedNotificationTransportModeEnumEnumValues = []NotificationTransportModeEnum{
	"local",
	"webhook",
	"webhook_slack",
	"email",
}

func (v *NotificationTransportModeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationTransportModeEnum(value)
	for _, existing := range AllowedNotificationTransportModeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationTransportModeEnum", value)
}

// NewNotificationTransportModeEnumFromValue returns a pointer to a valid NotificationTransportModeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationTransportModeEnumFromValue(v string) (*NotificationTransportModeEnum, error) {
	ev := NotificationTransportModeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationTransportModeEnum: valid values are %v", v, AllowedNotificationTransportModeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationTransportModeEnum) IsValid() bool {
	for _, existing := range AllowedNotificationTransportModeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationTransportModeEnum value
func (v NotificationTransportModeEnum) Ptr() *NotificationTransportModeEnum {
	return &v
}

type NullableNotificationTransportModeEnum struct {
	value *NotificationTransportModeEnum
	isSet bool
}

func (v NullableNotificationTransportModeEnum) Get() *NotificationTransportModeEnum {
	return v.value
}

func (v *NullableNotificationTransportModeEnum) Set(val *NotificationTransportModeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTransportModeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTransportModeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTransportModeEnum(val *NotificationTransportModeEnum) *NullableNotificationTransportModeEnum {
	return &NullableNotificationTransportModeEnum{value: val, isSet: true}
}

func (v NullableNotificationTransportModeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTransportModeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
