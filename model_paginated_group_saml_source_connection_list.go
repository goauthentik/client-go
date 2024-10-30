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

// PaginatedGroupSAMLSourceConnectionList struct for PaginatedGroupSAMLSourceConnectionList
type PaginatedGroupSAMLSourceConnectionList struct {
	Pagination Pagination                  `json:"pagination"`
	Results    []GroupSAMLSourceConnection `json:"results"`
}

// NewPaginatedGroupSAMLSourceConnectionList instantiates a new PaginatedGroupSAMLSourceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGroupSAMLSourceConnectionList(pagination Pagination, results []GroupSAMLSourceConnection) *PaginatedGroupSAMLSourceConnectionList {
	this := PaginatedGroupSAMLSourceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedGroupSAMLSourceConnectionListWithDefaults instantiates a new PaginatedGroupSAMLSourceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGroupSAMLSourceConnectionListWithDefaults() *PaginatedGroupSAMLSourceConnectionList {
	this := PaginatedGroupSAMLSourceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedGroupSAMLSourceConnectionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedGroupSAMLSourceConnectionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedGroupSAMLSourceConnectionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedGroupSAMLSourceConnectionList) GetResults() []GroupSAMLSourceConnection {
	if o == nil {
		var ret []GroupSAMLSourceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedGroupSAMLSourceConnectionList) GetResultsOk() ([]GroupSAMLSourceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedGroupSAMLSourceConnectionList) SetResults(v []GroupSAMLSourceConnection) {
	o.Results = v
}

func (o PaginatedGroupSAMLSourceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedGroupSAMLSourceConnectionList struct {
	value *PaginatedGroupSAMLSourceConnectionList
	isSet bool
}

func (v NullablePaginatedGroupSAMLSourceConnectionList) Get() *PaginatedGroupSAMLSourceConnectionList {
	return v.value
}

func (v *NullablePaginatedGroupSAMLSourceConnectionList) Set(val *PaginatedGroupSAMLSourceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGroupSAMLSourceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGroupSAMLSourceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGroupSAMLSourceConnectionList(val *PaginatedGroupSAMLSourceConnectionList) *NullablePaginatedGroupSAMLSourceConnectionList {
	return &NullablePaginatedGroupSAMLSourceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedGroupSAMLSourceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGroupSAMLSourceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
