/*
authentik

Making authentication simple.

API version: 2024.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedGroupPlexSourceConnectionList struct for PaginatedGroupPlexSourceConnectionList
type PaginatedGroupPlexSourceConnectionList struct {
	Pagination Pagination                  `json:"pagination"`
	Results    []GroupPlexSourceConnection `json:"results"`
}

// NewPaginatedGroupPlexSourceConnectionList instantiates a new PaginatedGroupPlexSourceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGroupPlexSourceConnectionList(pagination Pagination, results []GroupPlexSourceConnection) *PaginatedGroupPlexSourceConnectionList {
	this := PaginatedGroupPlexSourceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedGroupPlexSourceConnectionListWithDefaults instantiates a new PaginatedGroupPlexSourceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGroupPlexSourceConnectionListWithDefaults() *PaginatedGroupPlexSourceConnectionList {
	this := PaginatedGroupPlexSourceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedGroupPlexSourceConnectionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedGroupPlexSourceConnectionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedGroupPlexSourceConnectionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedGroupPlexSourceConnectionList) GetResults() []GroupPlexSourceConnection {
	if o == nil {
		var ret []GroupPlexSourceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedGroupPlexSourceConnectionList) GetResultsOk() ([]GroupPlexSourceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedGroupPlexSourceConnectionList) SetResults(v []GroupPlexSourceConnection) {
	o.Results = v
}

func (o PaginatedGroupPlexSourceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedGroupPlexSourceConnectionList struct {
	value *PaginatedGroupPlexSourceConnectionList
	isSet bool
}

func (v NullablePaginatedGroupPlexSourceConnectionList) Get() *PaginatedGroupPlexSourceConnectionList {
	return v.value
}

func (v *NullablePaginatedGroupPlexSourceConnectionList) Set(val *PaginatedGroupPlexSourceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGroupPlexSourceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGroupPlexSourceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGroupPlexSourceConnectionList(val *PaginatedGroupPlexSourceConnectionList) *NullablePaginatedGroupPlexSourceConnectionList {
	return &NullablePaginatedGroupPlexSourceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedGroupPlexSourceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGroupPlexSourceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
