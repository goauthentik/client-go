/*
authentik

Making authentication simple.

API version: 2024.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedRoleList struct for PaginatedRoleList
type PaginatedRoleList struct {
	Pagination Pagination `json:"pagination"`
	Results    []Role     `json:"results"`
}

// NewPaginatedRoleList instantiates a new PaginatedRoleList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedRoleList(pagination Pagination, results []Role) *PaginatedRoleList {
	this := PaginatedRoleList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedRoleListWithDefaults instantiates a new PaginatedRoleList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedRoleListWithDefaults() *PaginatedRoleList {
	this := PaginatedRoleList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedRoleList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedRoleList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedRoleList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedRoleList) GetResults() []Role {
	if o == nil {
		var ret []Role
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedRoleList) GetResultsOk() ([]Role, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedRoleList) SetResults(v []Role) {
	o.Results = v
}

func (o PaginatedRoleList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedRoleList struct {
	value *PaginatedRoleList
	isSet bool
}

func (v NullablePaginatedRoleList) Get() *PaginatedRoleList {
	return v.value
}

func (v *NullablePaginatedRoleList) Set(val *PaginatedRoleList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedRoleList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedRoleList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedRoleList(val *PaginatedRoleList) *NullablePaginatedRoleList {
	return &NullablePaginatedRoleList{value: val, isSet: true}
}

func (v NullablePaginatedRoleList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedRoleList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
