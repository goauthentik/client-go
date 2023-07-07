/*
authentik

Making authentication simple.

API version: 2023.6.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedUserSourceConnectionList struct for PaginatedUserSourceConnectionList
type PaginatedUserSourceConnectionList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []UserSourceConnection             `json:"results"`
}

// NewPaginatedUserSourceConnectionList instantiates a new PaginatedUserSourceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserSourceConnectionList(pagination PaginatedApplicationListPagination, results []UserSourceConnection) *PaginatedUserSourceConnectionList {
	this := PaginatedUserSourceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserSourceConnectionListWithDefaults instantiates a new PaginatedUserSourceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserSourceConnectionListWithDefaults() *PaginatedUserSourceConnectionList {
	this := PaginatedUserSourceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserSourceConnectionList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserSourceConnectionList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserSourceConnectionList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserSourceConnectionList) GetResults() []UserSourceConnection {
	if o == nil {
		var ret []UserSourceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserSourceConnectionList) GetResultsOk() ([]UserSourceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserSourceConnectionList) SetResults(v []UserSourceConnection) {
	o.Results = v
}

func (o PaginatedUserSourceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedUserSourceConnectionList struct {
	value *PaginatedUserSourceConnectionList
	isSet bool
}

func (v NullablePaginatedUserSourceConnectionList) Get() *PaginatedUserSourceConnectionList {
	return v.value
}

func (v *NullablePaginatedUserSourceConnectionList) Set(val *PaginatedUserSourceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserSourceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserSourceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserSourceConnectionList(val *PaginatedUserSourceConnectionList) *NullablePaginatedUserSourceConnectionList {
	return &NullablePaginatedUserSourceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedUserSourceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserSourceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
