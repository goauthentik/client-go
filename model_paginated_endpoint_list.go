/*
authentik

Making authentication simple.

API version: 2025.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedEndpointList struct for PaginatedEndpointList
type PaginatedEndpointList struct {
	Pagination Pagination `json:"pagination"`
	Results    []Endpoint `json:"results"`
}

// NewPaginatedEndpointList instantiates a new PaginatedEndpointList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEndpointList(pagination Pagination, results []Endpoint) *PaginatedEndpointList {
	this := PaginatedEndpointList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedEndpointListWithDefaults instantiates a new PaginatedEndpointList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEndpointListWithDefaults() *PaginatedEndpointList {
	this := PaginatedEndpointList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedEndpointList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedEndpointList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedEndpointList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedEndpointList) GetResults() []Endpoint {
	if o == nil {
		var ret []Endpoint
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedEndpointList) GetResultsOk() ([]Endpoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedEndpointList) SetResults(v []Endpoint) {
	o.Results = v
}

func (o PaginatedEndpointList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedEndpointList struct {
	value *PaginatedEndpointList
	isSet bool
}

func (v NullablePaginatedEndpointList) Get() *PaginatedEndpointList {
	return v.value
}

func (v *NullablePaginatedEndpointList) Set(val *PaginatedEndpointList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEndpointList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEndpointList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEndpointList(val *PaginatedEndpointList) *NullablePaginatedEndpointList {
	return &NullablePaginatedEndpointList{value: val, isSet: true}
}

func (v NullablePaginatedEndpointList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEndpointList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
