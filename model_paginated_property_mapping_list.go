/*
authentik

Making authentication simple.

API version: 2023.5.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedPropertyMappingList struct for PaginatedPropertyMappingList
type PaginatedPropertyMappingList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []PropertyMapping                  `json:"results"`
}

// NewPaginatedPropertyMappingList instantiates a new PaginatedPropertyMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPropertyMappingList(pagination PaginatedApplicationListPagination, results []PropertyMapping) *PaginatedPropertyMappingList {
	this := PaginatedPropertyMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedPropertyMappingListWithDefaults instantiates a new PaginatedPropertyMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPropertyMappingListWithDefaults() *PaginatedPropertyMappingList {
	this := PaginatedPropertyMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPropertyMappingList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPropertyMappingList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPropertyMappingList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPropertyMappingList) GetResults() []PropertyMapping {
	if o == nil {
		var ret []PropertyMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPropertyMappingList) GetResultsOk() ([]PropertyMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPropertyMappingList) SetResults(v []PropertyMapping) {
	o.Results = v
}

func (o PaginatedPropertyMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedPropertyMappingList struct {
	value *PaginatedPropertyMappingList
	isSet bool
}

func (v NullablePaginatedPropertyMappingList) Get() *PaginatedPropertyMappingList {
	return v.value
}

func (v *NullablePaginatedPropertyMappingList) Set(val *PaginatedPropertyMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPropertyMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPropertyMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPropertyMappingList(val *PaginatedPropertyMappingList) *NullablePaginatedPropertyMappingList {
	return &NullablePaginatedPropertyMappingList{value: val, isSet: true}
}

func (v NullablePaginatedPropertyMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPropertyMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
