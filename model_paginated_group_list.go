/*
authentik

Making authentication simple.

API version: 2023.1.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedGroupList struct for PaginatedGroupList
type PaginatedGroupList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []Group                            `json:"results"`
}

// NewPaginatedGroupList instantiates a new PaginatedGroupList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGroupList(pagination PaginatedApplicationListPagination, results []Group) *PaginatedGroupList {
	this := PaginatedGroupList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedGroupListWithDefaults instantiates a new PaginatedGroupList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGroupListWithDefaults() *PaginatedGroupList {
	this := PaginatedGroupList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedGroupList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedGroupList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedGroupList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedGroupList) GetResults() []Group {
	if o == nil {
		var ret []Group
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedGroupList) GetResultsOk() ([]Group, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedGroupList) SetResults(v []Group) {
	o.Results = v
}

func (o PaginatedGroupList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedGroupList struct {
	value *PaginatedGroupList
	isSet bool
}

func (v NullablePaginatedGroupList) Get() *PaginatedGroupList {
	return v.value
}

func (v *NullablePaginatedGroupList) Set(val *PaginatedGroupList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGroupList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGroupList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGroupList(val *PaginatedGroupList) *NullablePaginatedGroupList {
	return &NullablePaginatedGroupList{value: val, isSet: true}
}

func (v NullablePaginatedGroupList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGroupList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
