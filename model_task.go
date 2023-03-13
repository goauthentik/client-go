/*
authentik

Making authentication simple.

API version: 2023.3.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// Task Serialize TaskInfo and TaskResult
type Task struct {
	TaskName            string    `json:"task_name"`
	TaskDescription     string    `json:"task_description"`
	TaskFinishTimestamp time.Time `json:"task_finish_timestamp"`
	// Get the duration a task took to run
	TaskDuration int32          `json:"task_duration"`
	Status       TaskStatusEnum `json:"status"`
	Messages     []interface{}  `json:"messages"`
}

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask(taskName string, taskDescription string, taskFinishTimestamp time.Time, taskDuration int32, status TaskStatusEnum, messages []interface{}) *Task {
	this := Task{}
	this.TaskName = taskName
	this.TaskDescription = taskDescription
	this.TaskFinishTimestamp = taskFinishTimestamp
	this.TaskDuration = taskDuration
	this.Status = status
	this.Messages = messages
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetTaskName returns the TaskName field value
func (o *Task) GetTaskName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value
// and a boolean to check if the value has been set.
func (o *Task) GetTaskNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskName, true
}

// SetTaskName sets field value
func (o *Task) SetTaskName(v string) {
	o.TaskName = v
}

// GetTaskDescription returns the TaskDescription field value
func (o *Task) GetTaskDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskDescription
}

// GetTaskDescriptionOk returns a tuple with the TaskDescription field value
// and a boolean to check if the value has been set.
func (o *Task) GetTaskDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskDescription, true
}

// SetTaskDescription sets field value
func (o *Task) SetTaskDescription(v string) {
	o.TaskDescription = v
}

// GetTaskFinishTimestamp returns the TaskFinishTimestamp field value
func (o *Task) GetTaskFinishTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TaskFinishTimestamp
}

// GetTaskFinishTimestampOk returns a tuple with the TaskFinishTimestamp field value
// and a boolean to check if the value has been set.
func (o *Task) GetTaskFinishTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskFinishTimestamp, true
}

// SetTaskFinishTimestamp sets field value
func (o *Task) SetTaskFinishTimestamp(v time.Time) {
	o.TaskFinishTimestamp = v
}

// GetTaskDuration returns the TaskDuration field value
func (o *Task) GetTaskDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TaskDuration
}

// GetTaskDurationOk returns a tuple with the TaskDuration field value
// and a boolean to check if the value has been set.
func (o *Task) GetTaskDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskDuration, true
}

// SetTaskDuration sets field value
func (o *Task) SetTaskDuration(v int32) {
	o.TaskDuration = v
}

// GetStatus returns the Status field value
func (o *Task) GetStatus() TaskStatusEnum {
	if o == nil {
		var ret TaskStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Task) GetStatusOk() (*TaskStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Task) SetStatus(v TaskStatusEnum) {
	o.Status = v
}

// GetMessages returns the Messages field value
func (o *Task) GetMessages() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *Task) GetMessagesOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *Task) SetMessages(v []interface{}) {
	o.Messages = v
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["task_name"] = o.TaskName
	}
	if true {
		toSerialize["task_description"] = o.TaskDescription
	}
	if true {
		toSerialize["task_finish_timestamp"] = o.TaskFinishTimestamp
	}
	if true {
		toSerialize["task_duration"] = o.TaskDuration
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
