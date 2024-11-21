/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedPromptStageList struct for PaginatedPromptStageList
type PaginatedPromptStageList struct {
	Pagination Pagination    `json:"pagination"`
	Results    []PromptStage `json:"results"`
}

// NewPaginatedPromptStageList instantiates a new PaginatedPromptStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPromptStageList(pagination Pagination, results []PromptStage) *PaginatedPromptStageList {
	this := PaginatedPromptStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedPromptStageListWithDefaults instantiates a new PaginatedPromptStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPromptStageListWithDefaults() *PaginatedPromptStageList {
	this := PaginatedPromptStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPromptStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPromptStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPromptStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPromptStageList) GetResults() []PromptStage {
	if o == nil {
		var ret []PromptStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPromptStageList) GetResultsOk() ([]PromptStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPromptStageList) SetResults(v []PromptStage) {
	o.Results = v
}

func (o PaginatedPromptStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedPromptStageList struct {
	value *PaginatedPromptStageList
	isSet bool
}

func (v NullablePaginatedPromptStageList) Get() *PaginatedPromptStageList {
	return v.value
}

func (v *NullablePaginatedPromptStageList) Set(val *PaginatedPromptStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPromptStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPromptStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPromptStageList(val *PaginatedPromptStageList) *NullablePaginatedPromptStageList {
	return &NullablePaginatedPromptStageList{value: val, isSet: true}
}

func (v NullablePaginatedPromptStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPromptStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
