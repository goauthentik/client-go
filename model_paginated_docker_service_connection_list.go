/*
authentik

Making authentication simple.

API version: 2024.4.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedDockerServiceConnectionList struct for PaginatedDockerServiceConnectionList
type PaginatedDockerServiceConnectionList struct {
	Pagination Pagination                `json:"pagination"`
	Results    []DockerServiceConnection `json:"results"`
}

// NewPaginatedDockerServiceConnectionList instantiates a new PaginatedDockerServiceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDockerServiceConnectionList(pagination Pagination, results []DockerServiceConnection) *PaginatedDockerServiceConnectionList {
	this := PaginatedDockerServiceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedDockerServiceConnectionListWithDefaults instantiates a new PaginatedDockerServiceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDockerServiceConnectionListWithDefaults() *PaginatedDockerServiceConnectionList {
	this := PaginatedDockerServiceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedDockerServiceConnectionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedDockerServiceConnectionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedDockerServiceConnectionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedDockerServiceConnectionList) GetResults() []DockerServiceConnection {
	if o == nil {
		var ret []DockerServiceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedDockerServiceConnectionList) GetResultsOk() ([]DockerServiceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedDockerServiceConnectionList) SetResults(v []DockerServiceConnection) {
	o.Results = v
}

func (o PaginatedDockerServiceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedDockerServiceConnectionList struct {
	value *PaginatedDockerServiceConnectionList
	isSet bool
}

func (v NullablePaginatedDockerServiceConnectionList) Get() *PaginatedDockerServiceConnectionList {
	return v.value
}

func (v *NullablePaginatedDockerServiceConnectionList) Set(val *PaginatedDockerServiceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedDockerServiceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedDockerServiceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedDockerServiceConnectionList(val *PaginatedDockerServiceConnectionList) *NullablePaginatedDockerServiceConnectionList {
	return &NullablePaginatedDockerServiceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedDockerServiceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedDockerServiceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
