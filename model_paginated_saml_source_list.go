/*
authentik

Making authentication simple.

API version: 2023.5.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSAMLSourceList struct for PaginatedSAMLSourceList
type PaginatedSAMLSourceList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []SAMLSource                       `json:"results"`
}

// NewPaginatedSAMLSourceList instantiates a new PaginatedSAMLSourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSAMLSourceList(pagination PaginatedApplicationListPagination, results []SAMLSource) *PaginatedSAMLSourceList {
	this := PaginatedSAMLSourceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSAMLSourceListWithDefaults instantiates a new PaginatedSAMLSourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSAMLSourceListWithDefaults() *PaginatedSAMLSourceList {
	this := PaginatedSAMLSourceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSAMLSourceList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSAMLSourceList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSAMLSourceList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSAMLSourceList) GetResults() []SAMLSource {
	if o == nil {
		var ret []SAMLSource
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSAMLSourceList) GetResultsOk() ([]SAMLSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSAMLSourceList) SetResults(v []SAMLSource) {
	o.Results = v
}

func (o PaginatedSAMLSourceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSAMLSourceList struct {
	value *PaginatedSAMLSourceList
	isSet bool
}

func (v NullablePaginatedSAMLSourceList) Get() *PaginatedSAMLSourceList {
	return v.value
}

func (v *NullablePaginatedSAMLSourceList) Set(val *PaginatedSAMLSourceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSAMLSourceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSAMLSourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSAMLSourceList(val *PaginatedSAMLSourceList) *NullablePaginatedSAMLSourceList {
	return &NullablePaginatedSAMLSourceList{value: val, isSet: true}
}

func (v NullablePaginatedSAMLSourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSAMLSourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
