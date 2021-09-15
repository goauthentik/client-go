/*
authentik

Making authentication simple.

API version: 2021.9.1-rc1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedUserWriteStageList struct for PaginatedUserWriteStageList
type PaginatedUserWriteStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results []UserWriteStage `json:"results"`
}

// NewPaginatedUserWriteStageList instantiates a new PaginatedUserWriteStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserWriteStageList(pagination PaginatedApplicationListPagination, results []UserWriteStage) *PaginatedUserWriteStageList {
	this := PaginatedUserWriteStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserWriteStageListWithDefaults instantiates a new PaginatedUserWriteStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserWriteStageListWithDefaults() *PaginatedUserWriteStageList {
	this := PaginatedUserWriteStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserWriteStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserWriteStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserWriteStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserWriteStageList) GetResults() []UserWriteStage {
	if o == nil {
		var ret []UserWriteStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserWriteStageList) GetResultsOk() (*[]UserWriteStage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserWriteStageList) SetResults(v []UserWriteStage) {
	o.Results = v
}

func (o PaginatedUserWriteStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedUserWriteStageList struct {
	value *PaginatedUserWriteStageList
	isSet bool
}

func (v NullablePaginatedUserWriteStageList) Get() *PaginatedUserWriteStageList {
	return v.value
}

func (v *NullablePaginatedUserWriteStageList) Set(val *PaginatedUserWriteStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserWriteStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserWriteStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserWriteStageList(val *PaginatedUserWriteStageList) *NullablePaginatedUserWriteStageList {
	return &NullablePaginatedUserWriteStageList{value: val, isSet: true}
}

func (v NullablePaginatedUserWriteStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserWriteStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


