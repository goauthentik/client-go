/*
authentik

Making authentication simple.

API version: 2024.6.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedRadiusProviderPropertyMappingList struct for PaginatedRadiusProviderPropertyMappingList
type PaginatedRadiusProviderPropertyMappingList struct {
	Pagination Pagination                      `json:"pagination"`
	Results    []RadiusProviderPropertyMapping `json:"results"`
}

// NewPaginatedRadiusProviderPropertyMappingList instantiates a new PaginatedRadiusProviderPropertyMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedRadiusProviderPropertyMappingList(pagination Pagination, results []RadiusProviderPropertyMapping) *PaginatedRadiusProviderPropertyMappingList {
	this := PaginatedRadiusProviderPropertyMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedRadiusProviderPropertyMappingListWithDefaults instantiates a new PaginatedRadiusProviderPropertyMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedRadiusProviderPropertyMappingListWithDefaults() *PaginatedRadiusProviderPropertyMappingList {
	this := PaginatedRadiusProviderPropertyMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedRadiusProviderPropertyMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedRadiusProviderPropertyMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedRadiusProviderPropertyMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedRadiusProviderPropertyMappingList) GetResults() []RadiusProviderPropertyMapping {
	if o == nil {
		var ret []RadiusProviderPropertyMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedRadiusProviderPropertyMappingList) GetResultsOk() ([]RadiusProviderPropertyMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedRadiusProviderPropertyMappingList) SetResults(v []RadiusProviderPropertyMapping) {
	o.Results = v
}

func (o PaginatedRadiusProviderPropertyMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedRadiusProviderPropertyMappingList struct {
	value *PaginatedRadiusProviderPropertyMappingList
	isSet bool
}

func (v NullablePaginatedRadiusProviderPropertyMappingList) Get() *PaginatedRadiusProviderPropertyMappingList {
	return v.value
}

func (v *NullablePaginatedRadiusProviderPropertyMappingList) Set(val *PaginatedRadiusProviderPropertyMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedRadiusProviderPropertyMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedRadiusProviderPropertyMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedRadiusProviderPropertyMappingList(val *PaginatedRadiusProviderPropertyMappingList) *NullablePaginatedRadiusProviderPropertyMappingList {
	return &NullablePaginatedRadiusProviderPropertyMappingList{value: val, isSet: true}
}

func (v NullablePaginatedRadiusProviderPropertyMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedRadiusProviderPropertyMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}