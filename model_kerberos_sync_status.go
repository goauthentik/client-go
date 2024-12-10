/*
authentik

Making authentication simple.

API version: 2024.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// KerberosSyncStatus Kerberos Source sync status
type KerberosSyncStatus struct {
	IsRunning bool         `json:"is_running"`
	Tasks     []SystemTask `json:"tasks"`
}

// NewKerberosSyncStatus instantiates a new KerberosSyncStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosSyncStatus(isRunning bool, tasks []SystemTask) *KerberosSyncStatus {
	this := KerberosSyncStatus{}
	this.IsRunning = isRunning
	this.Tasks = tasks
	return &this
}

// NewKerberosSyncStatusWithDefaults instantiates a new KerberosSyncStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosSyncStatusWithDefaults() *KerberosSyncStatus {
	this := KerberosSyncStatus{}
	return &this
}

// GetIsRunning returns the IsRunning field value
func (o *KerberosSyncStatus) GetIsRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value
// and a boolean to check if the value has been set.
func (o *KerberosSyncStatus) GetIsRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRunning, true
}

// SetIsRunning sets field value
func (o *KerberosSyncStatus) SetIsRunning(v bool) {
	o.IsRunning = v
}

// GetTasks returns the Tasks field value
func (o *KerberosSyncStatus) GetTasks() []SystemTask {
	if o == nil {
		var ret []SystemTask
		return ret
	}

	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value
// and a boolean to check if the value has been set.
func (o *KerberosSyncStatus) GetTasksOk() ([]SystemTask, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tasks, true
}

// SetTasks sets field value
func (o *KerberosSyncStatus) SetTasks(v []SystemTask) {
	o.Tasks = v
}

func (o KerberosSyncStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["is_running"] = o.IsRunning
	}
	if true {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableKerberosSyncStatus struct {
	value *KerberosSyncStatus
	isSet bool
}

func (v NullableKerberosSyncStatus) Get() *KerberosSyncStatus {
	return v.value
}

func (v *NullableKerberosSyncStatus) Set(val *KerberosSyncStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosSyncStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosSyncStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosSyncStatus(val *KerberosSyncStatus) *NullableKerberosSyncStatus {
	return &NullableKerberosSyncStatus{value: val, isSet: true}
}

func (v NullableKerberosSyncStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosSyncStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
