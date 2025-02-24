/*
authentik

Making authentication simple.

API version: 2025.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSSFProviderList struct for PaginatedSSFProviderList
type PaginatedSSFProviderList struct {
	Pagination Pagination    `json:"pagination"`
	Results    []SSFProvider `json:"results"`
}

// NewPaginatedSSFProviderList instantiates a new PaginatedSSFProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSSFProviderList(pagination Pagination, results []SSFProvider) *PaginatedSSFProviderList {
	this := PaginatedSSFProviderList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSSFProviderListWithDefaults instantiates a new PaginatedSSFProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSSFProviderListWithDefaults() *PaginatedSSFProviderList {
	this := PaginatedSSFProviderList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSSFProviderList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSSFProviderList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSSFProviderList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSSFProviderList) GetResults() []SSFProvider {
	if o == nil {
		var ret []SSFProvider
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSSFProviderList) GetResultsOk() ([]SSFProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSSFProviderList) SetResults(v []SSFProvider) {
	o.Results = v
}

func (o PaginatedSSFProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSSFProviderList struct {
	value *PaginatedSSFProviderList
	isSet bool
}

func (v NullablePaginatedSSFProviderList) Get() *PaginatedSSFProviderList {
	return v.value
}

func (v *NullablePaginatedSSFProviderList) Set(val *PaginatedSSFProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSSFProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSSFProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSSFProviderList(val *PaginatedSSFProviderList) *NullablePaginatedSSFProviderList {
	return &NullablePaginatedSSFProviderList{value: val, isSet: true}
}

func (v NullablePaginatedSSFProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSSFProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
