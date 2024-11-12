/*
authentik

Making authentication simple.

API version: 2024.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedUserKerberosSourceConnectionList struct for PaginatedUserKerberosSourceConnectionList
type PaginatedUserKerberosSourceConnectionList struct {
	Pagination Pagination                     `json:"pagination"`
	Results    []UserKerberosSourceConnection `json:"results"`
}

// NewPaginatedUserKerberosSourceConnectionList instantiates a new PaginatedUserKerberosSourceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserKerberosSourceConnectionList(pagination Pagination, results []UserKerberosSourceConnection) *PaginatedUserKerberosSourceConnectionList {
	this := PaginatedUserKerberosSourceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserKerberosSourceConnectionListWithDefaults instantiates a new PaginatedUserKerberosSourceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserKerberosSourceConnectionListWithDefaults() *PaginatedUserKerberosSourceConnectionList {
	this := PaginatedUserKerberosSourceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserKerberosSourceConnectionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserKerberosSourceConnectionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserKerberosSourceConnectionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserKerberosSourceConnectionList) GetResults() []UserKerberosSourceConnection {
	if o == nil {
		var ret []UserKerberosSourceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserKerberosSourceConnectionList) GetResultsOk() ([]UserKerberosSourceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserKerberosSourceConnectionList) SetResults(v []UserKerberosSourceConnection) {
	o.Results = v
}

func (o PaginatedUserKerberosSourceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedUserKerberosSourceConnectionList struct {
	value *PaginatedUserKerberosSourceConnectionList
	isSet bool
}

func (v NullablePaginatedUserKerberosSourceConnectionList) Get() *PaginatedUserKerberosSourceConnectionList {
	return v.value
}

func (v *NullablePaginatedUserKerberosSourceConnectionList) Set(val *PaginatedUserKerberosSourceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserKerberosSourceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserKerberosSourceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserKerberosSourceConnectionList(val *PaginatedUserKerberosSourceConnectionList) *NullablePaginatedUserKerberosSourceConnectionList {
	return &NullablePaginatedUserKerberosSourceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedUserKerberosSourceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserKerberosSourceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}