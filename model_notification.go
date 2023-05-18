/*
authentik

Making authentication simple.

API version: 2023.5.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// Notification Notification Serializer
type Notification struct {
	Pk       string       `json:"pk"`
	Severity SeverityEnum `json:"severity"`
	Body     string       `json:"body"`
	Created  time.Time    `json:"created"`
	Event    *Event       `json:"event,omitempty"`
	Seen     *bool        `json:"seen,omitempty"`
}

// NewNotification instantiates a new Notification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotification(pk string, severity SeverityEnum, body string, created time.Time) *Notification {
	this := Notification{}
	this.Pk = pk
	this.Severity = severity
	this.Body = body
	this.Created = created
	return &this
}

// NewNotificationWithDefaults instantiates a new Notification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWithDefaults() *Notification {
	this := Notification{}
	return &this
}

// GetPk returns the Pk field value
func (o *Notification) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Notification) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Notification) SetPk(v string) {
	o.Pk = v
}

// GetSeverity returns the Severity field value
func (o *Notification) GetSeverity() SeverityEnum {
	if o == nil {
		var ret SeverityEnum
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *Notification) GetSeverityOk() (*SeverityEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *Notification) SetSeverity(v SeverityEnum) {
	o.Severity = v
}

// GetBody returns the Body field value
func (o *Notification) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *Notification) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *Notification) SetBody(v string) {
	o.Body = v
}

// GetCreated returns the Created field value
func (o *Notification) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Notification) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Notification) SetCreated(v time.Time) {
	o.Created = v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *Notification) GetEvent() Event {
	if o == nil || o.Event == nil {
		var ret Event
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetEventOk() (*Event, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *Notification) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given Event and assigns it to the Event field.
func (o *Notification) SetEvent(v Event) {
	o.Event = &v
}

// GetSeen returns the Seen field value if set, zero value otherwise.
func (o *Notification) GetSeen() bool {
	if o == nil || o.Seen == nil {
		var ret bool
		return ret
	}
	return *o.Seen
}

// GetSeenOk returns a tuple with the Seen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetSeenOk() (*bool, bool) {
	if o == nil || o.Seen == nil {
		return nil, false
	}
	return o.Seen, true
}

// HasSeen returns a boolean if a field has been set.
func (o *Notification) HasSeen() bool {
	if o != nil && o.Seen != nil {
		return true
	}

	return false
}

// SetSeen gets a reference to the given bool and assigns it to the Seen field.
func (o *Notification) SetSeen(v bool) {
	o.Seen = &v
}

func (o Notification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["body"] = o.Body
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Seen != nil {
		toSerialize["seen"] = o.Seen
	}
	return json.Marshal(toSerialize)
}

type NullableNotification struct {
	value *Notification
	isSet bool
}

func (v NullableNotification) Get() *Notification {
	return v.value
}

func (v *NullableNotification) Set(val *Notification) {
	v.value = val
	v.isSet = true
}

func (v NullableNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotification(val *Notification) *NullableNotification {
	return &NullableNotification{value: val, isSet: true}
}

func (v NullableNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
