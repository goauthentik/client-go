/*
authentik

Making authentication simple.

API version: 2024.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SCIMSyncStatus SCIM Provider sync status
type SCIMSyncStatus struct {
	IsRunning bool         `json:"is_running"`
	Tasks     []SystemTask `json:"tasks"`
}

// NewSCIMSyncStatus instantiates a new SCIMSyncStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCIMSyncStatus(isRunning bool, tasks []SystemTask) *SCIMSyncStatus {
	this := SCIMSyncStatus{}
	this.IsRunning = isRunning
	this.Tasks = tasks
	return &this
}

// NewSCIMSyncStatusWithDefaults instantiates a new SCIMSyncStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCIMSyncStatusWithDefaults() *SCIMSyncStatus {
	this := SCIMSyncStatus{}
	return &this
}

// GetIsRunning returns the IsRunning field value
func (o *SCIMSyncStatus) GetIsRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value
// and a boolean to check if the value has been set.
func (o *SCIMSyncStatus) GetIsRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRunning, true
}

// SetIsRunning sets field value
func (o *SCIMSyncStatus) SetIsRunning(v bool) {
	o.IsRunning = v
}

// GetTasks returns the Tasks field value
func (o *SCIMSyncStatus) GetTasks() []SystemTask {
	if o == nil {
		var ret []SystemTask
		return ret
	}

	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value
// and a boolean to check if the value has been set.
func (o *SCIMSyncStatus) GetTasksOk() ([]SystemTask, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tasks, true
}

// SetTasks sets field value
func (o *SCIMSyncStatus) SetTasks(v []SystemTask) {
	o.Tasks = v
}

func (o SCIMSyncStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["is_running"] = o.IsRunning
	}
	if true {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableSCIMSyncStatus struct {
	value *SCIMSyncStatus
	isSet bool
}

func (v NullableSCIMSyncStatus) Get() *SCIMSyncStatus {
	return v.value
}

func (v *NullableSCIMSyncStatus) Set(val *SCIMSyncStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSCIMSyncStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSCIMSyncStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCIMSyncStatus(val *SCIMSyncStatus) *NullableSCIMSyncStatus {
	return &NullableSCIMSyncStatus{value: val, isSet: true}
}

func (v NullableSCIMSyncStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCIMSyncStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
