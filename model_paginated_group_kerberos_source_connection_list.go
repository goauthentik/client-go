/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedGroupKerberosSourceConnectionList struct for PaginatedGroupKerberosSourceConnectionList
type PaginatedGroupKerberosSourceConnectionList struct {
	Pagination Pagination                      `json:"pagination"`
	Results    []GroupKerberosSourceConnection `json:"results"`
}

// NewPaginatedGroupKerberosSourceConnectionList instantiates a new PaginatedGroupKerberosSourceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGroupKerberosSourceConnectionList(pagination Pagination, results []GroupKerberosSourceConnection) *PaginatedGroupKerberosSourceConnectionList {
	this := PaginatedGroupKerberosSourceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedGroupKerberosSourceConnectionListWithDefaults instantiates a new PaginatedGroupKerberosSourceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGroupKerberosSourceConnectionListWithDefaults() *PaginatedGroupKerberosSourceConnectionList {
	this := PaginatedGroupKerberosSourceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedGroupKerberosSourceConnectionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedGroupKerberosSourceConnectionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedGroupKerberosSourceConnectionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedGroupKerberosSourceConnectionList) GetResults() []GroupKerberosSourceConnection {
	if o == nil {
		var ret []GroupKerberosSourceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedGroupKerberosSourceConnectionList) GetResultsOk() ([]GroupKerberosSourceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedGroupKerberosSourceConnectionList) SetResults(v []GroupKerberosSourceConnection) {
	o.Results = v
}

func (o PaginatedGroupKerberosSourceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedGroupKerberosSourceConnectionList struct {
	value *PaginatedGroupKerberosSourceConnectionList
	isSet bool
}

func (v NullablePaginatedGroupKerberosSourceConnectionList) Get() *PaginatedGroupKerberosSourceConnectionList {
	return v.value
}

func (v *NullablePaginatedGroupKerberosSourceConnectionList) Set(val *PaginatedGroupKerberosSourceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGroupKerberosSourceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGroupKerberosSourceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGroupKerberosSourceConnectionList(val *PaginatedGroupKerberosSourceConnectionList) *NullablePaginatedGroupKerberosSourceConnectionList {
	return &NullablePaginatedGroupKerberosSourceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedGroupKerberosSourceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGroupKerberosSourceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
