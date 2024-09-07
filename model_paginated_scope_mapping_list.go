/*
authentik

Making authentication simple.

API version: 2024.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedScopeMappingList struct for PaginatedScopeMappingList
type PaginatedScopeMappingList struct {
	Pagination Pagination     `json:"pagination"`
	Results    []ScopeMapping `json:"results"`
}

// NewPaginatedScopeMappingList instantiates a new PaginatedScopeMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedScopeMappingList(pagination Pagination, results []ScopeMapping) *PaginatedScopeMappingList {
	this := PaginatedScopeMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedScopeMappingListWithDefaults instantiates a new PaginatedScopeMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedScopeMappingListWithDefaults() *PaginatedScopeMappingList {
	this := PaginatedScopeMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedScopeMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedScopeMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedScopeMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedScopeMappingList) GetResults() []ScopeMapping {
	if o == nil {
		var ret []ScopeMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedScopeMappingList) GetResultsOk() ([]ScopeMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedScopeMappingList) SetResults(v []ScopeMapping) {
	o.Results = v
}

func (o PaginatedScopeMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedScopeMappingList struct {
	value *PaginatedScopeMappingList
	isSet bool
}

func (v NullablePaginatedScopeMappingList) Get() *PaginatedScopeMappingList {
	return v.value
}

func (v *NullablePaginatedScopeMappingList) Set(val *PaginatedScopeMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedScopeMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedScopeMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedScopeMappingList(val *PaginatedScopeMappingList) *NullablePaginatedScopeMappingList {
	return &NullablePaginatedScopeMappingList{value: val, isSet: true}
}

func (v NullablePaginatedScopeMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedScopeMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
