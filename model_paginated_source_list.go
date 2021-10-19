/*
authentik

Making authentication simple.

API version: 2021.10.1-rc1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSourceList struct for PaginatedSourceList
type PaginatedSourceList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []Source                           `json:"results"`
}

// NewPaginatedSourceList instantiates a new PaginatedSourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSourceList(pagination PaginatedApplicationListPagination, results []Source) *PaginatedSourceList {
	this := PaginatedSourceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSourceListWithDefaults instantiates a new PaginatedSourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSourceListWithDefaults() *PaginatedSourceList {
	this := PaginatedSourceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSourceList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSourceList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSourceList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSourceList) GetResults() []Source {
	if o == nil {
		var ret []Source
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSourceList) GetResultsOk() (*[]Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedSourceList) SetResults(v []Source) {
	o.Results = v
}

func (o PaginatedSourceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSourceList struct {
	value *PaginatedSourceList
	isSet bool
}

func (v NullablePaginatedSourceList) Get() *PaginatedSourceList {
	return v.value
}

func (v *NullablePaginatedSourceList) Set(val *PaginatedSourceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSourceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSourceList(val *PaginatedSourceList) *NullablePaginatedSourceList {
	return &NullablePaginatedSourceList{value: val, isSet: true}
}

func (v NullablePaginatedSourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
