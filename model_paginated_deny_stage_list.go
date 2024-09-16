/*
authentik

Making authentication simple.

API version: 2024.8.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedDenyStageList struct for PaginatedDenyStageList
type PaginatedDenyStageList struct {
	Pagination Pagination  `json:"pagination"`
	Results    []DenyStage `json:"results"`
}

// NewPaginatedDenyStageList instantiates a new PaginatedDenyStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDenyStageList(pagination Pagination, results []DenyStage) *PaginatedDenyStageList {
	this := PaginatedDenyStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedDenyStageListWithDefaults instantiates a new PaginatedDenyStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDenyStageListWithDefaults() *PaginatedDenyStageList {
	this := PaginatedDenyStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedDenyStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedDenyStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedDenyStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedDenyStageList) GetResults() []DenyStage {
	if o == nil {
		var ret []DenyStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedDenyStageList) GetResultsOk() ([]DenyStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedDenyStageList) SetResults(v []DenyStage) {
	o.Results = v
}

func (o PaginatedDenyStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedDenyStageList struct {
	value *PaginatedDenyStageList
	isSet bool
}

func (v NullablePaginatedDenyStageList) Get() *PaginatedDenyStageList {
	return v.value
}

func (v *NullablePaginatedDenyStageList) Set(val *PaginatedDenyStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedDenyStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedDenyStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedDenyStageList(val *PaginatedDenyStageList) *NullablePaginatedDenyStageList {
	return &NullablePaginatedDenyStageList{value: val, isSet: true}
}

func (v NullablePaginatedDenyStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedDenyStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
