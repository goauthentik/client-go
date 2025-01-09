/*
authentik

Making authentication simple.

API version: 2024.12.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedExpiringBaseGrantModelList struct for PaginatedExpiringBaseGrantModelList
type PaginatedExpiringBaseGrantModelList struct {
	Pagination Pagination               `json:"pagination"`
	Results    []ExpiringBaseGrantModel `json:"results"`
}

// NewPaginatedExpiringBaseGrantModelList instantiates a new PaginatedExpiringBaseGrantModelList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedExpiringBaseGrantModelList(pagination Pagination, results []ExpiringBaseGrantModel) *PaginatedExpiringBaseGrantModelList {
	this := PaginatedExpiringBaseGrantModelList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedExpiringBaseGrantModelListWithDefaults instantiates a new PaginatedExpiringBaseGrantModelList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedExpiringBaseGrantModelListWithDefaults() *PaginatedExpiringBaseGrantModelList {
	this := PaginatedExpiringBaseGrantModelList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedExpiringBaseGrantModelList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedExpiringBaseGrantModelList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedExpiringBaseGrantModelList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedExpiringBaseGrantModelList) GetResults() []ExpiringBaseGrantModel {
	if o == nil {
		var ret []ExpiringBaseGrantModel
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedExpiringBaseGrantModelList) GetResultsOk() ([]ExpiringBaseGrantModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedExpiringBaseGrantModelList) SetResults(v []ExpiringBaseGrantModel) {
	o.Results = v
}

func (o PaginatedExpiringBaseGrantModelList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedExpiringBaseGrantModelList struct {
	value *PaginatedExpiringBaseGrantModelList
	isSet bool
}

func (v NullablePaginatedExpiringBaseGrantModelList) Get() *PaginatedExpiringBaseGrantModelList {
	return v.value
}

func (v *NullablePaginatedExpiringBaseGrantModelList) Set(val *PaginatedExpiringBaseGrantModelList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedExpiringBaseGrantModelList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedExpiringBaseGrantModelList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedExpiringBaseGrantModelList(val *PaginatedExpiringBaseGrantModelList) *NullablePaginatedExpiringBaseGrantModelList {
	return &NullablePaginatedExpiringBaseGrantModelList{value: val, isSet: true}
}

func (v NullablePaginatedExpiringBaseGrantModelList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedExpiringBaseGrantModelList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
