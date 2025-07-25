/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Worker struct for Worker
type Worker struct {
	WorkerId        string `json:"worker_id"`
	Version         string `json:"version"`
	VersionMatching bool   `json:"version_matching"`
}

// NewWorker instantiates a new Worker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorker(workerId string, version string, versionMatching bool) *Worker {
	this := Worker{}
	this.WorkerId = workerId
	this.Version = version
	this.VersionMatching = versionMatching
	return &this
}

// NewWorkerWithDefaults instantiates a new Worker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkerWithDefaults() *Worker {
	this := Worker{}
	return &this
}

// GetWorkerId returns the WorkerId field value
func (o *Worker) GetWorkerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value
// and a boolean to check if the value has been set.
func (o *Worker) GetWorkerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkerId, true
}

// SetWorkerId sets field value
func (o *Worker) SetWorkerId(v string) {
	o.WorkerId = v
}

// GetVersion returns the Version field value
func (o *Worker) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Worker) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Worker) SetVersion(v string) {
	o.Version = v
}

// GetVersionMatching returns the VersionMatching field value
func (o *Worker) GetVersionMatching() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VersionMatching
}

// GetVersionMatchingOk returns a tuple with the VersionMatching field value
// and a boolean to check if the value has been set.
func (o *Worker) GetVersionMatchingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionMatching, true
}

// SetVersionMatching sets field value
func (o *Worker) SetVersionMatching(v bool) {
	o.VersionMatching = v
}

func (o Worker) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["worker_id"] = o.WorkerId
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["version_matching"] = o.VersionMatching
	}
	return json.Marshal(toSerialize)
}

type NullableWorker struct {
	value *Worker
	isSet bool
}

func (v NullableWorker) Get() *Worker {
	return v.value
}

func (v *NullableWorker) Set(val *Worker) {
	v.value = val
	v.isSet = true
}

func (v NullableWorker) IsSet() bool {
	return v.isSet
}

func (v *NullableWorker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorker(val *Worker) *NullableWorker {
	return &NullableWorker{value: val, isSet: true}
}

func (v NullableWorker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
