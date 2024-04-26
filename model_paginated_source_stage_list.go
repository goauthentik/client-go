/*
authentik

Making authentication simple.

API version: 2024.4.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSourceStageList struct for PaginatedSourceStageList
type PaginatedSourceStageList struct {
	Pagination Pagination    `json:"pagination"`
	Results    []SourceStage `json:"results"`
}

// NewPaginatedSourceStageList instantiates a new PaginatedSourceStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSourceStageList(pagination Pagination, results []SourceStage) *PaginatedSourceStageList {
	this := PaginatedSourceStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSourceStageListWithDefaults instantiates a new PaginatedSourceStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSourceStageListWithDefaults() *PaginatedSourceStageList {
	this := PaginatedSourceStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSourceStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSourceStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSourceStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSourceStageList) GetResults() []SourceStage {
	if o == nil {
		var ret []SourceStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSourceStageList) GetResultsOk() ([]SourceStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSourceStageList) SetResults(v []SourceStage) {
	o.Results = v
}

func (o PaginatedSourceStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSourceStageList struct {
	value *PaginatedSourceStageList
	isSet bool
}

func (v NullablePaginatedSourceStageList) Get() *PaginatedSourceStageList {
	return v.value
}

func (v *NullablePaginatedSourceStageList) Set(val *PaginatedSourceStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSourceStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSourceStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSourceStageList(val *PaginatedSourceStageList) *NullablePaginatedSourceStageList {
	return &NullablePaginatedSourceStageList{value: val, isSet: true}
}

func (v NullablePaginatedSourceStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSourceStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
