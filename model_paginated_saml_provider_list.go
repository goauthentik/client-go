/*
authentik

Making authentication simple.

API version: 2024.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSAMLProviderList struct for PaginatedSAMLProviderList
type PaginatedSAMLProviderList struct {
	Pagination Pagination     `json:"pagination"`
	Results    []SAMLProvider `json:"results"`
}

// NewPaginatedSAMLProviderList instantiates a new PaginatedSAMLProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSAMLProviderList(pagination Pagination, results []SAMLProvider) *PaginatedSAMLProviderList {
	this := PaginatedSAMLProviderList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSAMLProviderListWithDefaults instantiates a new PaginatedSAMLProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSAMLProviderListWithDefaults() *PaginatedSAMLProviderList {
	this := PaginatedSAMLProviderList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSAMLProviderList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSAMLProviderList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSAMLProviderList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSAMLProviderList) GetResults() []SAMLProvider {
	if o == nil {
		var ret []SAMLProvider
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSAMLProviderList) GetResultsOk() ([]SAMLProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSAMLProviderList) SetResults(v []SAMLProvider) {
	o.Results = v
}

func (o PaginatedSAMLProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSAMLProviderList struct {
	value *PaginatedSAMLProviderList
	isSet bool
}

func (v NullablePaginatedSAMLProviderList) Get() *PaginatedSAMLProviderList {
	return v.value
}

func (v *NullablePaginatedSAMLProviderList) Set(val *PaginatedSAMLProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSAMLProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSAMLProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSAMLProviderList(val *PaginatedSAMLProviderList) *NullablePaginatedSAMLProviderList {
	return &NullablePaginatedSAMLProviderList{value: val, isSet: true}
}

func (v NullablePaginatedSAMLProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSAMLProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
