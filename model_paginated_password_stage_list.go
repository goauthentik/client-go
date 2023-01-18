/*
authentik

Making authentication simple.

API version: 2023.1.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedPasswordStageList struct for PaginatedPasswordStageList
type PaginatedPasswordStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []PasswordStage                    `json:"results"`
}

// NewPaginatedPasswordStageList instantiates a new PaginatedPasswordStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPasswordStageList(pagination PaginatedApplicationListPagination, results []PasswordStage) *PaginatedPasswordStageList {
	this := PaginatedPasswordStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedPasswordStageListWithDefaults instantiates a new PaginatedPasswordStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPasswordStageListWithDefaults() *PaginatedPasswordStageList {
	this := PaginatedPasswordStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPasswordStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPasswordStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPasswordStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPasswordStageList) GetResults() []PasswordStage {
	if o == nil {
		var ret []PasswordStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPasswordStageList) GetResultsOk() ([]PasswordStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPasswordStageList) SetResults(v []PasswordStage) {
	o.Results = v
}

func (o PaginatedPasswordStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedPasswordStageList struct {
	value *PaginatedPasswordStageList
	isSet bool
}

func (v NullablePaginatedPasswordStageList) Get() *PaginatedPasswordStageList {
	return v.value
}

func (v *NullablePaginatedPasswordStageList) Set(val *PaginatedPasswordStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPasswordStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPasswordStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPasswordStageList(val *PaginatedPasswordStageList) *NullablePaginatedPasswordStageList {
	return &NullablePaginatedPasswordStageList{value: val, isSet: true}
}

func (v NullablePaginatedPasswordStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPasswordStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
