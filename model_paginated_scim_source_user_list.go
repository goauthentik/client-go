/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSCIMSourceUserList struct for PaginatedSCIMSourceUserList
type PaginatedSCIMSourceUserList struct {
	Pagination Pagination       `json:"pagination"`
	Results    []SCIMSourceUser `json:"results"`
}

// NewPaginatedSCIMSourceUserList instantiates a new PaginatedSCIMSourceUserList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSCIMSourceUserList(pagination Pagination, results []SCIMSourceUser) *PaginatedSCIMSourceUserList {
	this := PaginatedSCIMSourceUserList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSCIMSourceUserListWithDefaults instantiates a new PaginatedSCIMSourceUserList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSCIMSourceUserListWithDefaults() *PaginatedSCIMSourceUserList {
	this := PaginatedSCIMSourceUserList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSCIMSourceUserList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMSourceUserList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSCIMSourceUserList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSCIMSourceUserList) GetResults() []SCIMSourceUser {
	if o == nil {
		var ret []SCIMSourceUser
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMSourceUserList) GetResultsOk() ([]SCIMSourceUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSCIMSourceUserList) SetResults(v []SCIMSourceUser) {
	o.Results = v
}

func (o PaginatedSCIMSourceUserList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSCIMSourceUserList struct {
	value *PaginatedSCIMSourceUserList
	isSet bool
}

func (v NullablePaginatedSCIMSourceUserList) Get() *PaginatedSCIMSourceUserList {
	return v.value
}

func (v *NullablePaginatedSCIMSourceUserList) Set(val *PaginatedSCIMSourceUserList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSCIMSourceUserList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSCIMSourceUserList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSCIMSourceUserList(val *PaginatedSCIMSourceUserList) *NullablePaginatedSCIMSourceUserList {
	return &NullablePaginatedSCIMSourceUserList{value: val, isSet: true}
}

func (v NullablePaginatedSCIMSourceUserList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSCIMSourceUserList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
