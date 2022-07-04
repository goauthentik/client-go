/*
authentik

Making authentication simple.

API version: 2022.7.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedUserLoginStageList struct for PaginatedUserLoginStageList
type PaginatedUserLoginStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []UserLoginStage                   `json:"results"`
}

// NewPaginatedUserLoginStageList instantiates a new PaginatedUserLoginStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserLoginStageList(pagination PaginatedApplicationListPagination, results []UserLoginStage) *PaginatedUserLoginStageList {
	this := PaginatedUserLoginStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserLoginStageListWithDefaults instantiates a new PaginatedUserLoginStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserLoginStageListWithDefaults() *PaginatedUserLoginStageList {
	this := PaginatedUserLoginStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserLoginStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserLoginStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserLoginStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserLoginStageList) GetResults() []UserLoginStage {
	if o == nil {
		var ret []UserLoginStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserLoginStageList) GetResultsOk() ([]UserLoginStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserLoginStageList) SetResults(v []UserLoginStage) {
	o.Results = v
}

func (o PaginatedUserLoginStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedUserLoginStageList struct {
	value *PaginatedUserLoginStageList
	isSet bool
}

func (v NullablePaginatedUserLoginStageList) Get() *PaginatedUserLoginStageList {
	return v.value
}

func (v *NullablePaginatedUserLoginStageList) Set(val *PaginatedUserLoginStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserLoginStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserLoginStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserLoginStageList(val *PaginatedUserLoginStageList) *NullablePaginatedUserLoginStageList {
	return &NullablePaginatedUserLoginStageList{value: val, isSet: true}
}

func (v NullablePaginatedUserLoginStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserLoginStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
