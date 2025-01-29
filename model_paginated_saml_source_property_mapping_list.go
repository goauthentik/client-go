/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSAMLSourcePropertyMappingList struct for PaginatedSAMLSourcePropertyMappingList
type PaginatedSAMLSourcePropertyMappingList struct {
	Pagination Pagination                  `json:"pagination"`
	Results    []SAMLSourcePropertyMapping `json:"results"`
}

// NewPaginatedSAMLSourcePropertyMappingList instantiates a new PaginatedSAMLSourcePropertyMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSAMLSourcePropertyMappingList(pagination Pagination, results []SAMLSourcePropertyMapping) *PaginatedSAMLSourcePropertyMappingList {
	this := PaginatedSAMLSourcePropertyMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSAMLSourcePropertyMappingListWithDefaults instantiates a new PaginatedSAMLSourcePropertyMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSAMLSourcePropertyMappingListWithDefaults() *PaginatedSAMLSourcePropertyMappingList {
	this := PaginatedSAMLSourcePropertyMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSAMLSourcePropertyMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSAMLSourcePropertyMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSAMLSourcePropertyMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSAMLSourcePropertyMappingList) GetResults() []SAMLSourcePropertyMapping {
	if o == nil {
		var ret []SAMLSourcePropertyMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSAMLSourcePropertyMappingList) GetResultsOk() ([]SAMLSourcePropertyMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSAMLSourcePropertyMappingList) SetResults(v []SAMLSourcePropertyMapping) {
	o.Results = v
}

func (o PaginatedSAMLSourcePropertyMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSAMLSourcePropertyMappingList struct {
	value *PaginatedSAMLSourcePropertyMappingList
	isSet bool
}

func (v NullablePaginatedSAMLSourcePropertyMappingList) Get() *PaginatedSAMLSourcePropertyMappingList {
	return v.value
}

func (v *NullablePaginatedSAMLSourcePropertyMappingList) Set(val *PaginatedSAMLSourcePropertyMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSAMLSourcePropertyMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSAMLSourcePropertyMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSAMLSourcePropertyMappingList(val *PaginatedSAMLSourcePropertyMappingList) *NullablePaginatedSAMLSourcePropertyMappingList {
	return &NullablePaginatedSAMLSourcePropertyMappingList{value: val, isSet: true}
}

func (v NullablePaginatedSAMLSourcePropertyMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSAMLSourcePropertyMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
