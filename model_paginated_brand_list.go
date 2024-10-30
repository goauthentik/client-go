/*
authentik

Making authentication simple.

API version: 2024.10.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedBrandList struct for PaginatedBrandList
type PaginatedBrandList struct {
	Pagination Pagination `json:"pagination"`
	Results    []Brand    `json:"results"`
}

// NewPaginatedBrandList instantiates a new PaginatedBrandList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedBrandList(pagination Pagination, results []Brand) *PaginatedBrandList {
	this := PaginatedBrandList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedBrandListWithDefaults instantiates a new PaginatedBrandList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedBrandListWithDefaults() *PaginatedBrandList {
	this := PaginatedBrandList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedBrandList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedBrandList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedBrandList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedBrandList) GetResults() []Brand {
	if o == nil {
		var ret []Brand
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedBrandList) GetResultsOk() ([]Brand, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedBrandList) SetResults(v []Brand) {
	o.Results = v
}

func (o PaginatedBrandList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedBrandList struct {
	value *PaginatedBrandList
	isSet bool
}

func (v NullablePaginatedBrandList) Get() *PaginatedBrandList {
	return v.value
}

func (v *NullablePaginatedBrandList) Set(val *PaginatedBrandList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedBrandList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedBrandList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedBrandList(val *PaginatedBrandList) *NullablePaginatedBrandList {
	return &NullablePaginatedBrandList{value: val, isSet: true}
}

func (v NullablePaginatedBrandList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedBrandList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
