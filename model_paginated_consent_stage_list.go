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

// PaginatedConsentStageList struct for PaginatedConsentStageList
type PaginatedConsentStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []ConsentStage                     `json:"results"`
}

// NewPaginatedConsentStageList instantiates a new PaginatedConsentStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedConsentStageList(pagination PaginatedApplicationListPagination, results []ConsentStage) *PaginatedConsentStageList {
	this := PaginatedConsentStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedConsentStageListWithDefaults instantiates a new PaginatedConsentStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedConsentStageListWithDefaults() *PaginatedConsentStageList {
	this := PaginatedConsentStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedConsentStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedConsentStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedConsentStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedConsentStageList) GetResults() []ConsentStage {
	if o == nil {
		var ret []ConsentStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedConsentStageList) GetResultsOk() (*[]ConsentStage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedConsentStageList) SetResults(v []ConsentStage) {
	o.Results = v
}

func (o PaginatedConsentStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedConsentStageList struct {
	value *PaginatedConsentStageList
	isSet bool
}

func (v NullablePaginatedConsentStageList) Get() *PaginatedConsentStageList {
	return v.value
}

func (v *NullablePaginatedConsentStageList) Set(val *PaginatedConsentStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedConsentStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedConsentStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedConsentStageList(val *PaginatedConsentStageList) *NullablePaginatedConsentStageList {
	return &NullablePaginatedConsentStageList{value: val, isSet: true}
}

func (v NullablePaginatedConsentStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedConsentStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
