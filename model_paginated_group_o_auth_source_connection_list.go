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

// PaginatedGroupOAuthSourceConnectionList struct for PaginatedGroupOAuthSourceConnectionList
type PaginatedGroupOAuthSourceConnectionList struct {
	Pagination Pagination                   `json:"pagination"`
	Results    []GroupOAuthSourceConnection `json:"results"`
}

// NewPaginatedGroupOAuthSourceConnectionList instantiates a new PaginatedGroupOAuthSourceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGroupOAuthSourceConnectionList(pagination Pagination, results []GroupOAuthSourceConnection) *PaginatedGroupOAuthSourceConnectionList {
	this := PaginatedGroupOAuthSourceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedGroupOAuthSourceConnectionListWithDefaults instantiates a new PaginatedGroupOAuthSourceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGroupOAuthSourceConnectionListWithDefaults() *PaginatedGroupOAuthSourceConnectionList {
	this := PaginatedGroupOAuthSourceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedGroupOAuthSourceConnectionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedGroupOAuthSourceConnectionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedGroupOAuthSourceConnectionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedGroupOAuthSourceConnectionList) GetResults() []GroupOAuthSourceConnection {
	if o == nil {
		var ret []GroupOAuthSourceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedGroupOAuthSourceConnectionList) GetResultsOk() ([]GroupOAuthSourceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedGroupOAuthSourceConnectionList) SetResults(v []GroupOAuthSourceConnection) {
	o.Results = v
}

func (o PaginatedGroupOAuthSourceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedGroupOAuthSourceConnectionList struct {
	value *PaginatedGroupOAuthSourceConnectionList
	isSet bool
}

func (v NullablePaginatedGroupOAuthSourceConnectionList) Get() *PaginatedGroupOAuthSourceConnectionList {
	return v.value
}

func (v *NullablePaginatedGroupOAuthSourceConnectionList) Set(val *PaginatedGroupOAuthSourceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGroupOAuthSourceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGroupOAuthSourceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGroupOAuthSourceConnectionList(val *PaginatedGroupOAuthSourceConnectionList) *NullablePaginatedGroupOAuthSourceConnectionList {
	return &NullablePaginatedGroupOAuthSourceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedGroupOAuthSourceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGroupOAuthSourceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
