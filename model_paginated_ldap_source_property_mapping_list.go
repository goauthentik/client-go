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

// PaginatedLDAPSourcePropertyMappingList struct for PaginatedLDAPSourcePropertyMappingList
type PaginatedLDAPSourcePropertyMappingList struct {
	Pagination Pagination                  `json:"pagination"`
	Results    []LDAPSourcePropertyMapping `json:"results"`
}

// NewPaginatedLDAPSourcePropertyMappingList instantiates a new PaginatedLDAPSourcePropertyMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedLDAPSourcePropertyMappingList(pagination Pagination, results []LDAPSourcePropertyMapping) *PaginatedLDAPSourcePropertyMappingList {
	this := PaginatedLDAPSourcePropertyMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedLDAPSourcePropertyMappingListWithDefaults instantiates a new PaginatedLDAPSourcePropertyMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedLDAPSourcePropertyMappingListWithDefaults() *PaginatedLDAPSourcePropertyMappingList {
	this := PaginatedLDAPSourcePropertyMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedLDAPSourcePropertyMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedLDAPSourcePropertyMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedLDAPSourcePropertyMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedLDAPSourcePropertyMappingList) GetResults() []LDAPSourcePropertyMapping {
	if o == nil {
		var ret []LDAPSourcePropertyMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedLDAPSourcePropertyMappingList) GetResultsOk() ([]LDAPSourcePropertyMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedLDAPSourcePropertyMappingList) SetResults(v []LDAPSourcePropertyMapping) {
	o.Results = v
}

func (o PaginatedLDAPSourcePropertyMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedLDAPSourcePropertyMappingList struct {
	value *PaginatedLDAPSourcePropertyMappingList
	isSet bool
}

func (v NullablePaginatedLDAPSourcePropertyMappingList) Get() *PaginatedLDAPSourcePropertyMappingList {
	return v.value
}

func (v *NullablePaginatedLDAPSourcePropertyMappingList) Set(val *PaginatedLDAPSourcePropertyMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedLDAPSourcePropertyMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedLDAPSourcePropertyMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedLDAPSourcePropertyMappingList(val *PaginatedLDAPSourcePropertyMappingList) *NullablePaginatedLDAPSourcePropertyMappingList {
	return &NullablePaginatedLDAPSourcePropertyMappingList{value: val, isSet: true}
}

func (v NullablePaginatedLDAPSourcePropertyMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedLDAPSourcePropertyMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
