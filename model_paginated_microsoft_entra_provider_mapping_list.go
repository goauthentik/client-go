/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedMicrosoftEntraProviderMappingList struct for PaginatedMicrosoftEntraProviderMappingList
type PaginatedMicrosoftEntraProviderMappingList struct {
	Pagination Pagination                      `json:"pagination"`
	Results    []MicrosoftEntraProviderMapping `json:"results"`
}

// NewPaginatedMicrosoftEntraProviderMappingList instantiates a new PaginatedMicrosoftEntraProviderMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedMicrosoftEntraProviderMappingList(pagination Pagination, results []MicrosoftEntraProviderMapping) *PaginatedMicrosoftEntraProviderMappingList {
	this := PaginatedMicrosoftEntraProviderMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedMicrosoftEntraProviderMappingListWithDefaults instantiates a new PaginatedMicrosoftEntraProviderMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedMicrosoftEntraProviderMappingListWithDefaults() *PaginatedMicrosoftEntraProviderMappingList {
	this := PaginatedMicrosoftEntraProviderMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedMicrosoftEntraProviderMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedMicrosoftEntraProviderMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedMicrosoftEntraProviderMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedMicrosoftEntraProviderMappingList) GetResults() []MicrosoftEntraProviderMapping {
	if o == nil {
		var ret []MicrosoftEntraProviderMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedMicrosoftEntraProviderMappingList) GetResultsOk() ([]MicrosoftEntraProviderMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedMicrosoftEntraProviderMappingList) SetResults(v []MicrosoftEntraProviderMapping) {
	o.Results = v
}

func (o PaginatedMicrosoftEntraProviderMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedMicrosoftEntraProviderMappingList struct {
	value *PaginatedMicrosoftEntraProviderMappingList
	isSet bool
}

func (v NullablePaginatedMicrosoftEntraProviderMappingList) Get() *PaginatedMicrosoftEntraProviderMappingList {
	return v.value
}

func (v *NullablePaginatedMicrosoftEntraProviderMappingList) Set(val *PaginatedMicrosoftEntraProviderMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedMicrosoftEntraProviderMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedMicrosoftEntraProviderMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedMicrosoftEntraProviderMappingList(val *PaginatedMicrosoftEntraProviderMappingList) *NullablePaginatedMicrosoftEntraProviderMappingList {
	return &NullablePaginatedMicrosoftEntraProviderMappingList{value: val, isSet: true}
}

func (v NullablePaginatedMicrosoftEntraProviderMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedMicrosoftEntraProviderMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
