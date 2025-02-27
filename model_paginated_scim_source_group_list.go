/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSCIMSourceGroupList struct for PaginatedSCIMSourceGroupList
type PaginatedSCIMSourceGroupList struct {
	Pagination Pagination        `json:"pagination"`
	Results    []SCIMSourceGroup `json:"results"`
}

// NewPaginatedSCIMSourceGroupList instantiates a new PaginatedSCIMSourceGroupList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSCIMSourceGroupList(pagination Pagination, results []SCIMSourceGroup) *PaginatedSCIMSourceGroupList {
	this := PaginatedSCIMSourceGroupList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSCIMSourceGroupListWithDefaults instantiates a new PaginatedSCIMSourceGroupList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSCIMSourceGroupListWithDefaults() *PaginatedSCIMSourceGroupList {
	this := PaginatedSCIMSourceGroupList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSCIMSourceGroupList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMSourceGroupList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSCIMSourceGroupList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSCIMSourceGroupList) GetResults() []SCIMSourceGroup {
	if o == nil {
		var ret []SCIMSourceGroup
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMSourceGroupList) GetResultsOk() ([]SCIMSourceGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSCIMSourceGroupList) SetResults(v []SCIMSourceGroup) {
	o.Results = v
}

func (o PaginatedSCIMSourceGroupList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSCIMSourceGroupList struct {
	value *PaginatedSCIMSourceGroupList
	isSet bool
}

func (v NullablePaginatedSCIMSourceGroupList) Get() *PaginatedSCIMSourceGroupList {
	return v.value
}

func (v *NullablePaginatedSCIMSourceGroupList) Set(val *PaginatedSCIMSourceGroupList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSCIMSourceGroupList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSCIMSourceGroupList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSCIMSourceGroupList(val *PaginatedSCIMSourceGroupList) *NullablePaginatedSCIMSourceGroupList {
	return &NullablePaginatedSCIMSourceGroupList{value: val, isSet: true}
}

func (v NullablePaginatedSCIMSourceGroupList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSCIMSourceGroupList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
