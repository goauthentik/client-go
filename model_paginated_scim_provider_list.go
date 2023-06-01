/*
authentik

Making authentication simple.

API version: 2023.5.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSCIMProviderList struct for PaginatedSCIMProviderList
type PaginatedSCIMProviderList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []SCIMProvider                     `json:"results"`
}

// NewPaginatedSCIMProviderList instantiates a new PaginatedSCIMProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSCIMProviderList(pagination PaginatedApplicationListPagination, results []SCIMProvider) *PaginatedSCIMProviderList {
	this := PaginatedSCIMProviderList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSCIMProviderListWithDefaults instantiates a new PaginatedSCIMProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSCIMProviderListWithDefaults() *PaginatedSCIMProviderList {
	this := PaginatedSCIMProviderList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSCIMProviderList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMProviderList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSCIMProviderList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSCIMProviderList) GetResults() []SCIMProvider {
	if o == nil {
		var ret []SCIMProvider
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMProviderList) GetResultsOk() ([]SCIMProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSCIMProviderList) SetResults(v []SCIMProvider) {
	o.Results = v
}

func (o PaginatedSCIMProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSCIMProviderList struct {
	value *PaginatedSCIMProviderList
	isSet bool
}

func (v NullablePaginatedSCIMProviderList) Get() *PaginatedSCIMProviderList {
	return v.value
}

func (v *NullablePaginatedSCIMProviderList) Set(val *PaginatedSCIMProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSCIMProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSCIMProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSCIMProviderList(val *PaginatedSCIMProviderList) *NullablePaginatedSCIMProviderList {
	return &NullablePaginatedSCIMProviderList{value: val, isSet: true}
}

func (v NullablePaginatedSCIMProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSCIMProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
