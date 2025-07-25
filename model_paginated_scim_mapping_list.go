/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSCIMMappingList struct for PaginatedSCIMMappingList
type PaginatedSCIMMappingList struct {
	Pagination   Pagination             `json:"pagination"`
	Results      []SCIMMapping          `json:"results"`
	Autocomplete map[string]interface{} `json:"autocomplete"`
}

// NewPaginatedSCIMMappingList instantiates a new PaginatedSCIMMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSCIMMappingList(pagination Pagination, results []SCIMMapping, autocomplete map[string]interface{}) *PaginatedSCIMMappingList {
	this := PaginatedSCIMMappingList{}
	this.Pagination = pagination
	this.Results = results
	this.Autocomplete = autocomplete
	return &this
}

// NewPaginatedSCIMMappingListWithDefaults instantiates a new PaginatedSCIMMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSCIMMappingListWithDefaults() *PaginatedSCIMMappingList {
	this := PaginatedSCIMMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSCIMMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSCIMMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSCIMMappingList) GetResults() []SCIMMapping {
	if o == nil {
		var ret []SCIMMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMMappingList) GetResultsOk() ([]SCIMMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSCIMMappingList) SetResults(v []SCIMMapping) {
	o.Results = v
}

// GetAutocomplete returns the Autocomplete field value
func (o *PaginatedSCIMMappingList) GetAutocomplete() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Autocomplete
}

// GetAutocompleteOk returns a tuple with the Autocomplete field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMMappingList) GetAutocompleteOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Autocomplete, true
}

// SetAutocomplete sets field value
func (o *PaginatedSCIMMappingList) SetAutocomplete(v map[string]interface{}) {
	o.Autocomplete = v
}

func (o PaginatedSCIMMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if true {
		toSerialize["autocomplete"] = o.Autocomplete
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSCIMMappingList struct {
	value *PaginatedSCIMMappingList
	isSet bool
}

func (v NullablePaginatedSCIMMappingList) Get() *PaginatedSCIMMappingList {
	return v.value
}

func (v *NullablePaginatedSCIMMappingList) Set(val *PaginatedSCIMMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSCIMMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSCIMMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSCIMMappingList(val *PaginatedSCIMMappingList) *NullablePaginatedSCIMMappingList {
	return &NullablePaginatedSCIMMappingList{value: val, isSet: true}
}

func (v NullablePaginatedSCIMMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSCIMMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
