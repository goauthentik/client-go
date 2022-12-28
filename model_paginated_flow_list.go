/*
authentik

Making authentication simple.

API version: 2022.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedFlowList struct for PaginatedFlowList
type PaginatedFlowList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []Flow                             `json:"results"`
}

// NewPaginatedFlowList instantiates a new PaginatedFlowList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedFlowList(pagination PaginatedApplicationListPagination, results []Flow) *PaginatedFlowList {
	this := PaginatedFlowList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedFlowListWithDefaults instantiates a new PaginatedFlowList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedFlowListWithDefaults() *PaginatedFlowList {
	this := PaginatedFlowList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedFlowList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedFlowList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedFlowList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedFlowList) GetResults() []Flow {
	if o == nil {
		var ret []Flow
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedFlowList) GetResultsOk() ([]Flow, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedFlowList) SetResults(v []Flow) {
	o.Results = v
}

func (o PaginatedFlowList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedFlowList struct {
	value *PaginatedFlowList
	isSet bool
}

func (v NullablePaginatedFlowList) Get() *PaginatedFlowList {
	return v.value
}

func (v *NullablePaginatedFlowList) Set(val *PaginatedFlowList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedFlowList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedFlowList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedFlowList(val *PaginatedFlowList) *NullablePaginatedFlowList {
	return &NullablePaginatedFlowList{value: val, isSet: true}
}

func (v NullablePaginatedFlowList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedFlowList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
