/*
authentik

Making authentication simple.

API version: 2023.5.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedRadiusProviderList struct for PaginatedRadiusProviderList
type PaginatedRadiusProviderList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []RadiusProvider                   `json:"results"`
}

// NewPaginatedRadiusProviderList instantiates a new PaginatedRadiusProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedRadiusProviderList(pagination PaginatedApplicationListPagination, results []RadiusProvider) *PaginatedRadiusProviderList {
	this := PaginatedRadiusProviderList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedRadiusProviderListWithDefaults instantiates a new PaginatedRadiusProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedRadiusProviderListWithDefaults() *PaginatedRadiusProviderList {
	this := PaginatedRadiusProviderList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedRadiusProviderList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedRadiusProviderList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedRadiusProviderList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedRadiusProviderList) GetResults() []RadiusProvider {
	if o == nil {
		var ret []RadiusProvider
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedRadiusProviderList) GetResultsOk() ([]RadiusProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedRadiusProviderList) SetResults(v []RadiusProvider) {
	o.Results = v
}

func (o PaginatedRadiusProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedRadiusProviderList struct {
	value *PaginatedRadiusProviderList
	isSet bool
}

func (v NullablePaginatedRadiusProviderList) Get() *PaginatedRadiusProviderList {
	return v.value
}

func (v *NullablePaginatedRadiusProviderList) Set(val *PaginatedRadiusProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedRadiusProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedRadiusProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedRadiusProviderList(val *PaginatedRadiusProviderList) *NullablePaginatedRadiusProviderList {
	return &NullablePaginatedRadiusProviderList{value: val, isSet: true}
}

func (v NullablePaginatedRadiusProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedRadiusProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
