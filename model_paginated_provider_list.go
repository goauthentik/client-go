/*
authentik

Making authentication simple.

API version: 2023.8.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedProviderList struct for PaginatedProviderList
type PaginatedProviderList struct {
	Pagination Pagination `json:"pagination"`
	Results    []Provider `json:"results"`
}

// NewPaginatedProviderList instantiates a new PaginatedProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedProviderList(pagination Pagination, results []Provider) *PaginatedProviderList {
	this := PaginatedProviderList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedProviderListWithDefaults instantiates a new PaginatedProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedProviderListWithDefaults() *PaginatedProviderList {
	this := PaginatedProviderList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedProviderList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedProviderList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedProviderList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedProviderList) GetResults() []Provider {
	if o == nil {
		var ret []Provider
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedProviderList) GetResultsOk() ([]Provider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedProviderList) SetResults(v []Provider) {
	o.Results = v
}

func (o PaginatedProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedProviderList struct {
	value *PaginatedProviderList
	isSet bool
}

func (v NullablePaginatedProviderList) Get() *PaginatedProviderList {
	return v.value
}

func (v *NullablePaginatedProviderList) Set(val *PaginatedProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedProviderList(val *PaginatedProviderList) *NullablePaginatedProviderList {
	return &NullablePaginatedProviderList{value: val, isSet: true}
}

func (v NullablePaginatedProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
