/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedMicrosoftEntraProviderList struct for PaginatedMicrosoftEntraProviderList
type PaginatedMicrosoftEntraProviderList struct {
	Pagination Pagination               `json:"pagination"`
	Results    []MicrosoftEntraProvider `json:"results"`
}

// NewPaginatedMicrosoftEntraProviderList instantiates a new PaginatedMicrosoftEntraProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedMicrosoftEntraProviderList(pagination Pagination, results []MicrosoftEntraProvider) *PaginatedMicrosoftEntraProviderList {
	this := PaginatedMicrosoftEntraProviderList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedMicrosoftEntraProviderListWithDefaults instantiates a new PaginatedMicrosoftEntraProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedMicrosoftEntraProviderListWithDefaults() *PaginatedMicrosoftEntraProviderList {
	this := PaginatedMicrosoftEntraProviderList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedMicrosoftEntraProviderList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedMicrosoftEntraProviderList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedMicrosoftEntraProviderList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedMicrosoftEntraProviderList) GetResults() []MicrosoftEntraProvider {
	if o == nil {
		var ret []MicrosoftEntraProvider
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedMicrosoftEntraProviderList) GetResultsOk() ([]MicrosoftEntraProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedMicrosoftEntraProviderList) SetResults(v []MicrosoftEntraProvider) {
	o.Results = v
}

func (o PaginatedMicrosoftEntraProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedMicrosoftEntraProviderList struct {
	value *PaginatedMicrosoftEntraProviderList
	isSet bool
}

func (v NullablePaginatedMicrosoftEntraProviderList) Get() *PaginatedMicrosoftEntraProviderList {
	return v.value
}

func (v *NullablePaginatedMicrosoftEntraProviderList) Set(val *PaginatedMicrosoftEntraProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedMicrosoftEntraProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedMicrosoftEntraProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedMicrosoftEntraProviderList(val *PaginatedMicrosoftEntraProviderList) *NullablePaginatedMicrosoftEntraProviderList {
	return &NullablePaginatedMicrosoftEntraProviderList{value: val, isSet: true}
}

func (v NullablePaginatedMicrosoftEntraProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedMicrosoftEntraProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
