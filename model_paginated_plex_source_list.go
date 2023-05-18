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

// PaginatedPlexSourceList struct for PaginatedPlexSourceList
type PaginatedPlexSourceList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []PlexSource                       `json:"results"`
}

// NewPaginatedPlexSourceList instantiates a new PaginatedPlexSourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPlexSourceList(pagination PaginatedApplicationListPagination, results []PlexSource) *PaginatedPlexSourceList {
	this := PaginatedPlexSourceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedPlexSourceListWithDefaults instantiates a new PaginatedPlexSourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPlexSourceListWithDefaults() *PaginatedPlexSourceList {
	this := PaginatedPlexSourceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPlexSourceList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPlexSourceList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPlexSourceList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPlexSourceList) GetResults() []PlexSource {
	if o == nil {
		var ret []PlexSource
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPlexSourceList) GetResultsOk() ([]PlexSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPlexSourceList) SetResults(v []PlexSource) {
	o.Results = v
}

func (o PaginatedPlexSourceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedPlexSourceList struct {
	value *PaginatedPlexSourceList
	isSet bool
}

func (v NullablePaginatedPlexSourceList) Get() *PaginatedPlexSourceList {
	return v.value
}

func (v *NullablePaginatedPlexSourceList) Set(val *PaginatedPlexSourceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPlexSourceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPlexSourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPlexSourceList(val *PaginatedPlexSourceList) *NullablePaginatedPlexSourceList {
	return &NullablePaginatedPlexSourceList{value: val, isSet: true}
}

func (v NullablePaginatedPlexSourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPlexSourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
