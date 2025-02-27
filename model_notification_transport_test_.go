/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// NotificationTransportTest Notification test serializer
type NotificationTransportTest struct {
	Messages []string `json:"messages"`
}

// NewNotificationTransportTest instantiates a new NotificationTransportTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationTransportTest(messages []string) *NotificationTransportTest {
	this := NotificationTransportTest{}
	this.Messages = messages
	return &this
}

// NewNotificationTransportTestWithDefaults instantiates a new NotificationTransportTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationTransportTestWithDefaults() *NotificationTransportTest {
	this := NotificationTransportTest{}
	return &this
}

// GetMessages returns the Messages field value
func (o *NotificationTransportTest) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *NotificationTransportTest) GetMessagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *NotificationTransportTest) SetMessages(v []string) {
	o.Messages = v
}

func (o NotificationTransportTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationTransportTest struct {
	value *NotificationTransportTest
	isSet bool
}

func (v NullableNotificationTransportTest) Get() *NotificationTransportTest {
	return v.value
}

func (v *NullableNotificationTransportTest) Set(val *NotificationTransportTest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationTransportTest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationTransportTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationTransportTest(val *NotificationTransportTest) *NullableNotificationTransportTest {
	return &NullableNotificationTransportTest{value: val, isSet: true}
}

func (v NullableNotificationTransportTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationTransportTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
