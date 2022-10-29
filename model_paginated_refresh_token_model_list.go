/*
authentik

Making authentication simple.

API version: 2022.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedRefreshTokenModelList struct for PaginatedRefreshTokenModelList
type PaginatedRefreshTokenModelList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []RefreshTokenModel                `json:"results"`
}

// NewPaginatedRefreshTokenModelList instantiates a new PaginatedRefreshTokenModelList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedRefreshTokenModelList(pagination PaginatedApplicationListPagination, results []RefreshTokenModel) *PaginatedRefreshTokenModelList {
	this := PaginatedRefreshTokenModelList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedRefreshTokenModelListWithDefaults instantiates a new PaginatedRefreshTokenModelList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedRefreshTokenModelListWithDefaults() *PaginatedRefreshTokenModelList {
	this := PaginatedRefreshTokenModelList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedRefreshTokenModelList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedRefreshTokenModelList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedRefreshTokenModelList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedRefreshTokenModelList) GetResults() []RefreshTokenModel {
	if o == nil {
		var ret []RefreshTokenModel
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedRefreshTokenModelList) GetResultsOk() ([]RefreshTokenModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedRefreshTokenModelList) SetResults(v []RefreshTokenModel) {
	o.Results = v
}

func (o PaginatedRefreshTokenModelList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedRefreshTokenModelList struct {
	value *PaginatedRefreshTokenModelList
	isSet bool
}

func (v NullablePaginatedRefreshTokenModelList) Get() *PaginatedRefreshTokenModelList {
	return v.value
}

func (v *NullablePaginatedRefreshTokenModelList) Set(val *PaginatedRefreshTokenModelList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedRefreshTokenModelList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedRefreshTokenModelList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedRefreshTokenModelList(val *PaginatedRefreshTokenModelList) *NullablePaginatedRefreshTokenModelList {
	return &NullablePaginatedRefreshTokenModelList{value: val, isSet: true}
}

func (v NullablePaginatedRefreshTokenModelList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedRefreshTokenModelList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
