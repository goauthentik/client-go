/*
authentik

Making authentication simple.

API version: 2024.4.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSystemTaskList struct for PaginatedSystemTaskList
type PaginatedSystemTaskList struct {
	Pagination Pagination   `json:"pagination"`
	Results    []SystemTask `json:"results"`
}

// NewPaginatedSystemTaskList instantiates a new PaginatedSystemTaskList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSystemTaskList(pagination Pagination, results []SystemTask) *PaginatedSystemTaskList {
	this := PaginatedSystemTaskList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSystemTaskListWithDefaults instantiates a new PaginatedSystemTaskList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSystemTaskListWithDefaults() *PaginatedSystemTaskList {
	this := PaginatedSystemTaskList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSystemTaskList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSystemTaskList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSystemTaskList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSystemTaskList) GetResults() []SystemTask {
	if o == nil {
		var ret []SystemTask
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSystemTaskList) GetResultsOk() ([]SystemTask, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSystemTaskList) SetResults(v []SystemTask) {
	o.Results = v
}

func (o PaginatedSystemTaskList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSystemTaskList struct {
	value *PaginatedSystemTaskList
	isSet bool
}

func (v NullablePaginatedSystemTaskList) Get() *PaginatedSystemTaskList {
	return v.value
}

func (v *NullablePaginatedSystemTaskList) Set(val *PaginatedSystemTaskList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSystemTaskList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSystemTaskList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSystemTaskList(val *PaginatedSystemTaskList) *NullablePaginatedSystemTaskList {
	return &NullablePaginatedSystemTaskList{value: val, isSet: true}
}

func (v NullablePaginatedSystemTaskList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSystemTaskList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
