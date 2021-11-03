/*
authentik

Making authentication simple.

API version: 2021.10.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedUserLogoutStageList struct for PaginatedUserLogoutStageList
type PaginatedUserLogoutStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []UserLogoutStage                  `json:"results"`
}

// NewPaginatedUserLogoutStageList instantiates a new PaginatedUserLogoutStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserLogoutStageList(pagination PaginatedApplicationListPagination, results []UserLogoutStage) *PaginatedUserLogoutStageList {
	this := PaginatedUserLogoutStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserLogoutStageListWithDefaults instantiates a new PaginatedUserLogoutStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserLogoutStageListWithDefaults() *PaginatedUserLogoutStageList {
	this := PaginatedUserLogoutStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserLogoutStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserLogoutStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserLogoutStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserLogoutStageList) GetResults() []UserLogoutStage {
	if o == nil {
		var ret []UserLogoutStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserLogoutStageList) GetResultsOk() (*[]UserLogoutStage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserLogoutStageList) SetResults(v []UserLogoutStage) {
	o.Results = v
}

func (o PaginatedUserLogoutStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedUserLogoutStageList struct {
	value *PaginatedUserLogoutStageList
	isSet bool
}

func (v NullablePaginatedUserLogoutStageList) Get() *PaginatedUserLogoutStageList {
	return v.value
}

func (v *NullablePaginatedUserLogoutStageList) Set(val *PaginatedUserLogoutStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserLogoutStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserLogoutStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserLogoutStageList(val *PaginatedUserLogoutStageList) *NullablePaginatedUserLogoutStageList {
	return &NullablePaginatedUserLogoutStageList{value: val, isSet: true}
}

func (v NullablePaginatedUserLogoutStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserLogoutStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
