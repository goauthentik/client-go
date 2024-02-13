/*
authentik

Making authentication simple.

API version: 2023.10.7
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedLicenseList struct for PaginatedLicenseList
type PaginatedLicenseList struct {
	Pagination Pagination `json:"pagination"`
	Results    []License  `json:"results"`
}

// NewPaginatedLicenseList instantiates a new PaginatedLicenseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedLicenseList(pagination Pagination, results []License) *PaginatedLicenseList {
	this := PaginatedLicenseList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedLicenseListWithDefaults instantiates a new PaginatedLicenseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedLicenseListWithDefaults() *PaginatedLicenseList {
	this := PaginatedLicenseList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedLicenseList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedLicenseList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedLicenseList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedLicenseList) GetResults() []License {
	if o == nil {
		var ret []License
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedLicenseList) GetResultsOk() ([]License, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedLicenseList) SetResults(v []License) {
	o.Results = v
}

func (o PaginatedLicenseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedLicenseList struct {
	value *PaginatedLicenseList
	isSet bool
}

func (v NullablePaginatedLicenseList) Get() *PaginatedLicenseList {
	return v.value
}

func (v *NullablePaginatedLicenseList) Set(val *PaginatedLicenseList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedLicenseList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedLicenseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedLicenseList(val *PaginatedLicenseList) *NullablePaginatedLicenseList {
	return &NullablePaginatedLicenseList{value: val, isSet: true}
}

func (v NullablePaginatedLicenseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedLicenseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
