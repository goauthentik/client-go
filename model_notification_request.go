/*
authentik

Making authentication simple.

API version: 2023.6.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// NotificationRequest Notification Serializer
type NotificationRequest struct {
	Event *EventRequest `json:"event,omitempty"`
	Seen  *bool         `json:"seen,omitempty"`
}

// NewNotificationRequest instantiates a new NotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationRequest() *NotificationRequest {
	this := NotificationRequest{}
	return &this
}

// NewNotificationRequestWithDefaults instantiates a new NotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationRequestWithDefaults() *NotificationRequest {
	this := NotificationRequest{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *NotificationRequest) GetEvent() EventRequest {
	if o == nil || o.Event == nil {
		var ret EventRequest
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetEventOk() (*EventRequest, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *NotificationRequest) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given EventRequest and assigns it to the Event field.
func (o *NotificationRequest) SetEvent(v EventRequest) {
	o.Event = &v
}

// GetSeen returns the Seen field value if set, zero value otherwise.
func (o *NotificationRequest) GetSeen() bool {
	if o == nil || o.Seen == nil {
		var ret bool
		return ret
	}
	return *o.Seen
}

// GetSeenOk returns a tuple with the Seen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRequest) GetSeenOk() (*bool, bool) {
	if o == nil || o.Seen == nil {
		return nil, false
	}
	return o.Seen, true
}

// HasSeen returns a boolean if a field has been set.
func (o *NotificationRequest) HasSeen() bool {
	if o != nil && o.Seen != nil {
		return true
	}

	return false
}

// SetSeen gets a reference to the given bool and assigns it to the Seen field.
func (o *NotificationRequest) SetSeen(v bool) {
	o.Seen = &v
}

func (o NotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Seen != nil {
		toSerialize["seen"] = o.Seen
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationRequest struct {
	value *NotificationRequest
	isSet bool
}

func (v NullableNotificationRequest) Get() *NotificationRequest {
	return v.value
}

func (v *NullableNotificationRequest) Set(val *NotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationRequest(val *NotificationRequest) *NullableNotificationRequest {
	return &NullableNotificationRequest{value: val, isSet: true}
}

func (v NullableNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
