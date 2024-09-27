/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSCIMSourcePropertyMappingList struct for PaginatedSCIMSourcePropertyMappingList
type PaginatedSCIMSourcePropertyMappingList struct {
	Pagination Pagination                  `json:"pagination"`
	Results    []SCIMSourcePropertyMapping `json:"results"`
}

// NewPaginatedSCIMSourcePropertyMappingList instantiates a new PaginatedSCIMSourcePropertyMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSCIMSourcePropertyMappingList(pagination Pagination, results []SCIMSourcePropertyMapping) *PaginatedSCIMSourcePropertyMappingList {
	this := PaginatedSCIMSourcePropertyMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSCIMSourcePropertyMappingListWithDefaults instantiates a new PaginatedSCIMSourcePropertyMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSCIMSourcePropertyMappingListWithDefaults() *PaginatedSCIMSourcePropertyMappingList {
	this := PaginatedSCIMSourcePropertyMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSCIMSourcePropertyMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMSourcePropertyMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSCIMSourcePropertyMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSCIMSourcePropertyMappingList) GetResults() []SCIMSourcePropertyMapping {
	if o == nil {
		var ret []SCIMSourcePropertyMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMSourcePropertyMappingList) GetResultsOk() ([]SCIMSourcePropertyMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSCIMSourcePropertyMappingList) SetResults(v []SCIMSourcePropertyMapping) {
	o.Results = v
}

func (o PaginatedSCIMSourcePropertyMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSCIMSourcePropertyMappingList struct {
	value *PaginatedSCIMSourcePropertyMappingList
	isSet bool
}

func (v NullablePaginatedSCIMSourcePropertyMappingList) Get() *PaginatedSCIMSourcePropertyMappingList {
	return v.value
}

func (v *NullablePaginatedSCIMSourcePropertyMappingList) Set(val *PaginatedSCIMSourcePropertyMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSCIMSourcePropertyMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSCIMSourcePropertyMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSCIMSourcePropertyMappingList(val *PaginatedSCIMSourcePropertyMappingList) *NullablePaginatedSCIMSourcePropertyMappingList {
	return &NullablePaginatedSCIMSourcePropertyMappingList{value: val, isSet: true}
}

func (v NullablePaginatedSCIMSourcePropertyMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSCIMSourcePropertyMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
