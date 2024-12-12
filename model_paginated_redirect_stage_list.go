/*
authentik

Making authentication simple.

API version: 2024.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedRedirectStageList struct for PaginatedRedirectStageList
type PaginatedRedirectStageList struct {
	Pagination Pagination      `json:"pagination"`
	Results    []RedirectStage `json:"results"`
}

// NewPaginatedRedirectStageList instantiates a new PaginatedRedirectStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedRedirectStageList(pagination Pagination, results []RedirectStage) *PaginatedRedirectStageList {
	this := PaginatedRedirectStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedRedirectStageListWithDefaults instantiates a new PaginatedRedirectStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedRedirectStageListWithDefaults() *PaginatedRedirectStageList {
	this := PaginatedRedirectStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedRedirectStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedRedirectStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedRedirectStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedRedirectStageList) GetResults() []RedirectStage {
	if o == nil {
		var ret []RedirectStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedRedirectStageList) GetResultsOk() ([]RedirectStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedRedirectStageList) SetResults(v []RedirectStage) {
	o.Results = v
}

func (o PaginatedRedirectStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedRedirectStageList struct {
	value *PaginatedRedirectStageList
	isSet bool
}

func (v NullablePaginatedRedirectStageList) Get() *PaginatedRedirectStageList {
	return v.value
}

func (v *NullablePaginatedRedirectStageList) Set(val *PaginatedRedirectStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedRedirectStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedRedirectStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedRedirectStageList(val *PaginatedRedirectStageList) *NullablePaginatedRedirectStageList {
	return &NullablePaginatedRedirectStageList{value: val, isSet: true}
}

func (v NullablePaginatedRedirectStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedRedirectStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
