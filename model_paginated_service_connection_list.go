/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedServiceConnectionList struct for PaginatedServiceConnectionList
type PaginatedServiceConnectionList struct {
	Pagination Pagination          `json:"pagination"`
	Results    []ServiceConnection `json:"results"`
}

// NewPaginatedServiceConnectionList instantiates a new PaginatedServiceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedServiceConnectionList(pagination Pagination, results []ServiceConnection) *PaginatedServiceConnectionList {
	this := PaginatedServiceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedServiceConnectionListWithDefaults instantiates a new PaginatedServiceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedServiceConnectionListWithDefaults() *PaginatedServiceConnectionList {
	this := PaginatedServiceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedServiceConnectionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedServiceConnectionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedServiceConnectionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedServiceConnectionList) GetResults() []ServiceConnection {
	if o == nil {
		var ret []ServiceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedServiceConnectionList) GetResultsOk() ([]ServiceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedServiceConnectionList) SetResults(v []ServiceConnection) {
	o.Results = v
}

func (o PaginatedServiceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedServiceConnectionList struct {
	value *PaginatedServiceConnectionList
	isSet bool
}

func (v NullablePaginatedServiceConnectionList) Get() *PaginatedServiceConnectionList {
	return v.value
}

func (v *NullablePaginatedServiceConnectionList) Set(val *PaginatedServiceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedServiceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedServiceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedServiceConnectionList(val *PaginatedServiceConnectionList) *NullablePaginatedServiceConnectionList {
	return &NullablePaginatedServiceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedServiceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedServiceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
