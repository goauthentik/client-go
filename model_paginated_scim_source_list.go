/*
authentik

Making authentication simple.

API version: 2024.4.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSCIMSourceList struct for PaginatedSCIMSourceList
type PaginatedSCIMSourceList struct {
	Pagination Pagination   `json:"pagination"`
	Results    []SCIMSource `json:"results"`
}

// NewPaginatedSCIMSourceList instantiates a new PaginatedSCIMSourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSCIMSourceList(pagination Pagination, results []SCIMSource) *PaginatedSCIMSourceList {
	this := PaginatedSCIMSourceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSCIMSourceListWithDefaults instantiates a new PaginatedSCIMSourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSCIMSourceListWithDefaults() *PaginatedSCIMSourceList {
	this := PaginatedSCIMSourceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSCIMSourceList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMSourceList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSCIMSourceList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSCIMSourceList) GetResults() []SCIMSource {
	if o == nil {
		var ret []SCIMSource
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMSourceList) GetResultsOk() ([]SCIMSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSCIMSourceList) SetResults(v []SCIMSource) {
	o.Results = v
}

func (o PaginatedSCIMSourceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSCIMSourceList struct {
	value *PaginatedSCIMSourceList
	isSet bool
}

func (v NullablePaginatedSCIMSourceList) Get() *PaginatedSCIMSourceList {
	return v.value
}

func (v *NullablePaginatedSCIMSourceList) Set(val *PaginatedSCIMSourceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSCIMSourceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSCIMSourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSCIMSourceList(val *PaginatedSCIMSourceList) *NullablePaginatedSCIMSourceList {
	return &NullablePaginatedSCIMSourceList{value: val, isSet: true}
}

func (v NullablePaginatedSCIMSourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSCIMSourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}