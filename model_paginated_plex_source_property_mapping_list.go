/*
authentik

Making authentication simple.

API version: 2024.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedPlexSourcePropertyMappingList struct for PaginatedPlexSourcePropertyMappingList
type PaginatedPlexSourcePropertyMappingList struct {
	Pagination Pagination                  `json:"pagination"`
	Results    []PlexSourcePropertyMapping `json:"results"`
}

// NewPaginatedPlexSourcePropertyMappingList instantiates a new PaginatedPlexSourcePropertyMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPlexSourcePropertyMappingList(pagination Pagination, results []PlexSourcePropertyMapping) *PaginatedPlexSourcePropertyMappingList {
	this := PaginatedPlexSourcePropertyMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedPlexSourcePropertyMappingListWithDefaults instantiates a new PaginatedPlexSourcePropertyMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPlexSourcePropertyMappingListWithDefaults() *PaginatedPlexSourcePropertyMappingList {
	this := PaginatedPlexSourcePropertyMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPlexSourcePropertyMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPlexSourcePropertyMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPlexSourcePropertyMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPlexSourcePropertyMappingList) GetResults() []PlexSourcePropertyMapping {
	if o == nil {
		var ret []PlexSourcePropertyMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPlexSourcePropertyMappingList) GetResultsOk() ([]PlexSourcePropertyMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPlexSourcePropertyMappingList) SetResults(v []PlexSourcePropertyMapping) {
	o.Results = v
}

func (o PaginatedPlexSourcePropertyMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedPlexSourcePropertyMappingList struct {
	value *PaginatedPlexSourcePropertyMappingList
	isSet bool
}

func (v NullablePaginatedPlexSourcePropertyMappingList) Get() *PaginatedPlexSourcePropertyMappingList {
	return v.value
}

func (v *NullablePaginatedPlexSourcePropertyMappingList) Set(val *PaginatedPlexSourcePropertyMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPlexSourcePropertyMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPlexSourcePropertyMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPlexSourcePropertyMappingList(val *PaginatedPlexSourcePropertyMappingList) *NullablePaginatedPlexSourcePropertyMappingList {
	return &NullablePaginatedPlexSourcePropertyMappingList{value: val, isSet: true}
}

func (v NullablePaginatedPlexSourcePropertyMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPlexSourcePropertyMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
