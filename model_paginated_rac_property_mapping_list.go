/*
authentik

Making authentication simple.

API version: 2024.8.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedRACPropertyMappingList struct for PaginatedRACPropertyMappingList
type PaginatedRACPropertyMappingList struct {
	Pagination Pagination           `json:"pagination"`
	Results    []RACPropertyMapping `json:"results"`
}

// NewPaginatedRACPropertyMappingList instantiates a new PaginatedRACPropertyMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedRACPropertyMappingList(pagination Pagination, results []RACPropertyMapping) *PaginatedRACPropertyMappingList {
	this := PaginatedRACPropertyMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedRACPropertyMappingListWithDefaults instantiates a new PaginatedRACPropertyMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedRACPropertyMappingListWithDefaults() *PaginatedRACPropertyMappingList {
	this := PaginatedRACPropertyMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedRACPropertyMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedRACPropertyMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedRACPropertyMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedRACPropertyMappingList) GetResults() []RACPropertyMapping {
	if o == nil {
		var ret []RACPropertyMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedRACPropertyMappingList) GetResultsOk() ([]RACPropertyMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedRACPropertyMappingList) SetResults(v []RACPropertyMapping) {
	o.Results = v
}

func (o PaginatedRACPropertyMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedRACPropertyMappingList struct {
	value *PaginatedRACPropertyMappingList
	isSet bool
}

func (v NullablePaginatedRACPropertyMappingList) Get() *PaginatedRACPropertyMappingList {
	return v.value
}

func (v *NullablePaginatedRACPropertyMappingList) Set(val *PaginatedRACPropertyMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedRACPropertyMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedRACPropertyMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedRACPropertyMappingList(val *PaginatedRACPropertyMappingList) *NullablePaginatedRACPropertyMappingList {
	return &NullablePaginatedRACPropertyMappingList{value: val, isSet: true}
}

func (v NullablePaginatedRACPropertyMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedRACPropertyMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
