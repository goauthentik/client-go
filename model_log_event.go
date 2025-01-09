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
	"time"
)

// LogEvent Single log message with all context logged.
type LogEvent struct {
	Timestamp  time.Time              `json:"timestamp"`
	LogLevel   LogLevelEnum           `json:"log_level"`
	Logger     string                 `json:"logger"`
	Event      string                 `json:"event"`
	Attributes map[string]interface{} `json:"attributes"`
}

// NewLogEvent instantiates a new LogEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogEvent(timestamp time.Time, logLevel LogLevelEnum, logger string, event string, attributes map[string]interface{}) *LogEvent {
	this := LogEvent{}
	this.Timestamp = timestamp
	this.LogLevel = logLevel
	this.Logger = logger
	this.Event = event
	this.Attributes = attributes
	return &this
}

// NewLogEventWithDefaults instantiates a new LogEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogEventWithDefaults() *LogEvent {
	this := LogEvent{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *LogEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *LogEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *LogEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetLogLevel returns the LogLevel field value
func (o *LogEvent) GetLogLevel() LogLevelEnum {
	if o == nil {
		var ret LogLevelEnum
		return ret
	}

	return o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value
// and a boolean to check if the value has been set.
func (o *LogEvent) GetLogLevelOk() (*LogLevelEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogLevel, true
}

// SetLogLevel sets field value
func (o *LogEvent) SetLogLevel(v LogLevelEnum) {
	o.LogLevel = v
}

// GetLogger returns the Logger field value
func (o *LogEvent) GetLogger() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logger
}

// GetLoggerOk returns a tuple with the Logger field value
// and a boolean to check if the value has been set.
func (o *LogEvent) GetLoggerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logger, true
}

// SetLogger sets field value
func (o *LogEvent) SetLogger(v string) {
	o.Logger = v
}

// GetEvent returns the Event field value
func (o *LogEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *LogEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *LogEvent) SetEvent(v string) {
	o.Event = v
}

// GetAttributes returns the Attributes field value
func (o *LogEvent) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *LogEvent) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *LogEvent) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o LogEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["log_level"] = o.LogLevel
	}
	if true {
		toSerialize["logger"] = o.Logger
	}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableLogEvent struct {
	value *LogEvent
	isSet bool
}

func (v NullableLogEvent) Get() *LogEvent {
	return v.value
}

func (v *NullableLogEvent) Set(val *LogEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableLogEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableLogEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogEvent(val *LogEvent) *NullableLogEvent {
	return &NullableLogEvent{value: val, isSet: true}
}

func (v NullableLogEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
