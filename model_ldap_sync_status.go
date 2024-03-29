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

// LDAPSyncStatus LDAP Source sync status
type LDAPSyncStatus struct {
	IsRunning bool         `json:"is_running"`
	Tasks     []SystemTask `json:"tasks"`
}

// NewLDAPSyncStatus instantiates a new LDAPSyncStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPSyncStatus(isRunning bool, tasks []SystemTask) *LDAPSyncStatus {
	this := LDAPSyncStatus{}
	this.IsRunning = isRunning
	this.Tasks = tasks
	return &this
}

// NewLDAPSyncStatusWithDefaults instantiates a new LDAPSyncStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPSyncStatusWithDefaults() *LDAPSyncStatus {
	this := LDAPSyncStatus{}
	return &this
}

// GetIsRunning returns the IsRunning field value
func (o *LDAPSyncStatus) GetIsRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value
// and a boolean to check if the value has been set.
func (o *LDAPSyncStatus) GetIsRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRunning, true
}

// SetIsRunning sets field value
func (o *LDAPSyncStatus) SetIsRunning(v bool) {
	o.IsRunning = v
}

// GetTasks returns the Tasks field value
func (o *LDAPSyncStatus) GetTasks() []SystemTask {
	if o == nil {
		var ret []SystemTask
		return ret
	}

	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value
// and a boolean to check if the value has been set.
func (o *LDAPSyncStatus) GetTasksOk() ([]SystemTask, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tasks, true
}

// SetTasks sets field value
func (o *LDAPSyncStatus) SetTasks(v []SystemTask) {
	o.Tasks = v
}

func (o LDAPSyncStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["is_running"] = o.IsRunning
	}
	if true {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPSyncStatus struct {
	value *LDAPSyncStatus
	isSet bool
}

func (v NullableLDAPSyncStatus) Get() *LDAPSyncStatus {
	return v.value
}

func (v *NullableLDAPSyncStatus) Set(val *LDAPSyncStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPSyncStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPSyncStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPSyncStatus(val *LDAPSyncStatus) *NullableLDAPSyncStatus {
	return &NullableLDAPSyncStatus{value: val, isSet: true}
}

func (v NullableLDAPSyncStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPSyncStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
