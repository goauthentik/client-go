/*
authentik

Making authentication simple.

API version: 2024.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// SystemTask Serialize TaskInfo and TaskResult
type SystemTask struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	// Get full name with UID
	FullName        string               `json:"full_name"`
	Uid             *string              `json:"uid,omitempty"`
	Description     string               `json:"description"`
	StartTimestamp  time.Time            `json:"start_timestamp"`
	FinishTimestamp time.Time            `json:"finish_timestamp"`
	Duration        float64              `json:"duration"`
	Status          SystemTaskStatusEnum `json:"status"`
	Messages        []LogEvent           `json:"messages"`
	Expires         NullableTime         `json:"expires,omitempty"`
	Expiring        *bool                `json:"expiring,omitempty"`
}

// NewSystemTask instantiates a new SystemTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemTask(uuid string, name string, fullName string, description string, startTimestamp time.Time, finishTimestamp time.Time, duration float64, status SystemTaskStatusEnum, messages []LogEvent) *SystemTask {
	this := SystemTask{}
	this.Uuid = uuid
	this.Name = name
	this.FullName = fullName
	this.Description = description
	this.StartTimestamp = startTimestamp
	this.FinishTimestamp = finishTimestamp
	this.Duration = duration
	this.Status = status
	this.Messages = messages
	return &this
}

// NewSystemTaskWithDefaults instantiates a new SystemTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemTaskWithDefaults() *SystemTask {
	this := SystemTask{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *SystemTask) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SystemTask) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SystemTask) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *SystemTask) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SystemTask) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SystemTask) SetName(v string) {
	o.Name = v
}

// GetFullName returns the FullName field value
func (o *SystemTask) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *SystemTask) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *SystemTask) SetFullName(v string) {
	o.FullName = v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *SystemTask) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemTask) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *SystemTask) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *SystemTask) SetUid(v string) {
	o.Uid = &v
}

// GetDescription returns the Description field value
func (o *SystemTask) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SystemTask) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SystemTask) SetDescription(v string) {
	o.Description = v
}

// GetStartTimestamp returns the StartTimestamp field value
func (o *SystemTask) GetStartTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value
// and a boolean to check if the value has been set.
func (o *SystemTask) GetStartTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTimestamp, true
}

// SetStartTimestamp sets field value
func (o *SystemTask) SetStartTimestamp(v time.Time) {
	o.StartTimestamp = v
}

// GetFinishTimestamp returns the FinishTimestamp field value
func (o *SystemTask) GetFinishTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.FinishTimestamp
}

// GetFinishTimestampOk returns a tuple with the FinishTimestamp field value
// and a boolean to check if the value has been set.
func (o *SystemTask) GetFinishTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishTimestamp, true
}

// SetFinishTimestamp sets field value
func (o *SystemTask) SetFinishTimestamp(v time.Time) {
	o.FinishTimestamp = v
}

// GetDuration returns the Duration field value
func (o *SystemTask) GetDuration() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SystemTask) GetDurationOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *SystemTask) SetDuration(v float64) {
	o.Duration = v
}

// GetStatus returns the Status field value
func (o *SystemTask) GetStatus() SystemTaskStatusEnum {
	if o == nil {
		var ret SystemTaskStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SystemTask) GetStatusOk() (*SystemTaskStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SystemTask) SetStatus(v SystemTaskStatusEnum) {
	o.Status = v
}

// GetMessages returns the Messages field value
func (o *SystemTask) GetMessages() []LogEvent {
	if o == nil {
		var ret []LogEvent
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *SystemTask) GetMessagesOk() ([]LogEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *SystemTask) SetMessages(v []LogEvent) {
	o.Messages = v
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemTask) GetExpires() time.Time {
	if o == nil || o.Expires.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemTask) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *SystemTask) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *SystemTask) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}

// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *SystemTask) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *SystemTask) UnsetExpires() {
	o.Expires.Unset()
}

// GetExpiring returns the Expiring field value if set, zero value otherwise.
func (o *SystemTask) GetExpiring() bool {
	if o == nil || o.Expiring == nil {
		var ret bool
		return ret
	}
	return *o.Expiring
}

// GetExpiringOk returns a tuple with the Expiring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemTask) GetExpiringOk() (*bool, bool) {
	if o == nil || o.Expiring == nil {
		return nil, false
	}
	return o.Expiring, true
}

// HasExpiring returns a boolean if a field has been set.
func (o *SystemTask) HasExpiring() bool {
	if o != nil && o.Expiring != nil {
		return true
	}

	return false
}

// SetExpiring gets a reference to the given bool and assigns it to the Expiring field.
func (o *SystemTask) SetExpiring(v bool) {
	o.Expiring = &v
}

func (o SystemTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uuid"] = o.Uuid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["full_name"] = o.FullName
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["start_timestamp"] = o.StartTimestamp
	}
	if true {
		toSerialize["finish_timestamp"] = o.FinishTimestamp
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["messages"] = o.Messages
	}
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	if o.Expiring != nil {
		toSerialize["expiring"] = o.Expiring
	}
	return json.Marshal(toSerialize)
}

type NullableSystemTask struct {
	value *SystemTask
	isSet bool
}

func (v NullableSystemTask) Get() *SystemTask {
	return v.value
}

func (v *NullableSystemTask) Set(val *SystemTask) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemTask) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemTask(val *SystemTask) *NullableSystemTask {
	return &NullableSystemTask{value: val, isSet: true}
}

func (v NullableSystemTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
