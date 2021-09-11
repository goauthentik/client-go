/*
authentik

Making authentication simple.

API version: 2021.8.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedEmailStageList struct for PaginatedEmailStageList
type PaginatedEmailStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results []EmailStage `json:"results"`
}

// NewPaginatedEmailStageList instantiates a new PaginatedEmailStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEmailStageList(pagination PaginatedApplicationListPagination, results []EmailStage) *PaginatedEmailStageList {
	this := PaginatedEmailStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedEmailStageListWithDefaults instantiates a new PaginatedEmailStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEmailStageListWithDefaults() *PaginatedEmailStageList {
	this := PaginatedEmailStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedEmailStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedEmailStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedEmailStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedEmailStageList) GetResults() []EmailStage {
	if o == nil {
		var ret []EmailStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedEmailStageList) GetResultsOk() (*[]EmailStage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedEmailStageList) SetResults(v []EmailStage) {
	o.Results = v
}

func (o PaginatedEmailStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedEmailStageList struct {
	value *PaginatedEmailStageList
	isSet bool
}

func (v NullablePaginatedEmailStageList) Get() *PaginatedEmailStageList {
	return v.value
}

func (v *NullablePaginatedEmailStageList) Set(val *PaginatedEmailStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEmailStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEmailStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEmailStageList(val *PaginatedEmailStageList) *NullablePaginatedEmailStageList {
	return &NullablePaginatedEmailStageList{value: val, isSet: true}
}

func (v NullablePaginatedEmailStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEmailStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


