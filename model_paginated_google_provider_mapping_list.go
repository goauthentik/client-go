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

// PaginatedGoogleProviderMappingList struct for PaginatedGoogleProviderMappingList
type PaginatedGoogleProviderMappingList struct {
	Pagination Pagination              `json:"pagination"`
	Results    []GoogleProviderMapping `json:"results"`
}

// NewPaginatedGoogleProviderMappingList instantiates a new PaginatedGoogleProviderMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGoogleProviderMappingList(pagination Pagination, results []GoogleProviderMapping) *PaginatedGoogleProviderMappingList {
	this := PaginatedGoogleProviderMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedGoogleProviderMappingListWithDefaults instantiates a new PaginatedGoogleProviderMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGoogleProviderMappingListWithDefaults() *PaginatedGoogleProviderMappingList {
	this := PaginatedGoogleProviderMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedGoogleProviderMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedGoogleProviderMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedGoogleProviderMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedGoogleProviderMappingList) GetResults() []GoogleProviderMapping {
	if o == nil {
		var ret []GoogleProviderMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedGoogleProviderMappingList) GetResultsOk() ([]GoogleProviderMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedGoogleProviderMappingList) SetResults(v []GoogleProviderMapping) {
	o.Results = v
}

func (o PaginatedGoogleProviderMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedGoogleProviderMappingList struct {
	value *PaginatedGoogleProviderMappingList
	isSet bool
}

func (v NullablePaginatedGoogleProviderMappingList) Get() *PaginatedGoogleProviderMappingList {
	return v.value
}

func (v *NullablePaginatedGoogleProviderMappingList) Set(val *PaginatedGoogleProviderMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGoogleProviderMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGoogleProviderMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGoogleProviderMappingList(val *PaginatedGoogleProviderMappingList) *NullablePaginatedGoogleProviderMappingList {
	return &NullablePaginatedGoogleProviderMappingList{value: val, isSet: true}
}

func (v NullablePaginatedGoogleProviderMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGoogleProviderMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
