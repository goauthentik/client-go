/*
authentik

Making authentication simple.

API version: 2024.6.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SyncStatus Provider sync status
type SyncStatus struct {
	IsRunning bool         `json:"is_running"`
	Tasks     []SystemTask `json:"tasks"`
}

// NewSyncStatus instantiates a new SyncStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncStatus(isRunning bool, tasks []SystemTask) *SyncStatus {
	this := SyncStatus{}
	this.IsRunning = isRunning
	this.Tasks = tasks
	return &this
}

// NewSyncStatusWithDefaults instantiates a new SyncStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncStatusWithDefaults() *SyncStatus {
	this := SyncStatus{}
	return &this
}

// GetIsRunning returns the IsRunning field value
func (o *SyncStatus) GetIsRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value
// and a boolean to check if the value has been set.
func (o *SyncStatus) GetIsRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRunning, true
}

// SetIsRunning sets field value
func (o *SyncStatus) SetIsRunning(v bool) {
	o.IsRunning = v
}

// GetTasks returns the Tasks field value
func (o *SyncStatus) GetTasks() []SystemTask {
	if o == nil {
		var ret []SystemTask
		return ret
	}

	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value
// and a boolean to check if the value has been set.
func (o *SyncStatus) GetTasksOk() ([]SystemTask, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tasks, true
}

// SetTasks sets field value
func (o *SyncStatus) SetTasks(v []SystemTask) {
	o.Tasks = v
}

func (o SyncStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["is_running"] = o.IsRunning
	}
	if true {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableSyncStatus struct {
	value *SyncStatus
	isSet bool
}

func (v NullableSyncStatus) Get() *SyncStatus {
	return v.value
}

func (v *NullableSyncStatus) Set(val *SyncStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncStatus(val *SyncStatus) *NullableSyncStatus {
	return &NullableSyncStatus{value: val, isSet: true}
}

func (v NullableSyncStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
