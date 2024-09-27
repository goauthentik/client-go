/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedRACProviderList struct for PaginatedRACProviderList
type PaginatedRACProviderList struct {
	Pagination Pagination    `json:"pagination"`
	Results    []RACProvider `json:"results"`
}

// NewPaginatedRACProviderList instantiates a new PaginatedRACProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedRACProviderList(pagination Pagination, results []RACProvider) *PaginatedRACProviderList {
	this := PaginatedRACProviderList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedRACProviderListWithDefaults instantiates a new PaginatedRACProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedRACProviderListWithDefaults() *PaginatedRACProviderList {
	this := PaginatedRACProviderList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedRACProviderList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedRACProviderList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedRACProviderList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedRACProviderList) GetResults() []RACProvider {
	if o == nil {
		var ret []RACProvider
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedRACProviderList) GetResultsOk() ([]RACProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedRACProviderList) SetResults(v []RACProvider) {
	o.Results = v
}

func (o PaginatedRACProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedRACProviderList struct {
	value *PaginatedRACProviderList
	isSet bool
}

func (v NullablePaginatedRACProviderList) Get() *PaginatedRACProviderList {
	return v.value
}

func (v *NullablePaginatedRACProviderList) Set(val *PaginatedRACProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedRACProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedRACProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedRACProviderList(val *PaginatedRACProviderList) *NullablePaginatedRACProviderList {
	return &NullablePaginatedRACProviderList{value: val, isSet: true}
}

func (v NullablePaginatedRACProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedRACProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
