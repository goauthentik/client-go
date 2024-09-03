/*
authentik

Making authentication simple.

API version: 2024.8.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedLDAPProviderList struct for PaginatedLDAPProviderList
type PaginatedLDAPProviderList struct {
	Pagination Pagination     `json:"pagination"`
	Results    []LDAPProvider `json:"results"`
}

// NewPaginatedLDAPProviderList instantiates a new PaginatedLDAPProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedLDAPProviderList(pagination Pagination, results []LDAPProvider) *PaginatedLDAPProviderList {
	this := PaginatedLDAPProviderList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedLDAPProviderListWithDefaults instantiates a new PaginatedLDAPProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedLDAPProviderListWithDefaults() *PaginatedLDAPProviderList {
	this := PaginatedLDAPProviderList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedLDAPProviderList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedLDAPProviderList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedLDAPProviderList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedLDAPProviderList) GetResults() []LDAPProvider {
	if o == nil {
		var ret []LDAPProvider
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedLDAPProviderList) GetResultsOk() ([]LDAPProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedLDAPProviderList) SetResults(v []LDAPProvider) {
	o.Results = v
}

func (o PaginatedLDAPProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedLDAPProviderList struct {
	value *PaginatedLDAPProviderList
	isSet bool
}

func (v NullablePaginatedLDAPProviderList) Get() *PaginatedLDAPProviderList {
	return v.value
}

func (v *NullablePaginatedLDAPProviderList) Set(val *PaginatedLDAPProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedLDAPProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedLDAPProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedLDAPProviderList(val *PaginatedLDAPProviderList) *NullablePaginatedLDAPProviderList {
	return &NullablePaginatedLDAPProviderList{value: val, isSet: true}
}

func (v NullablePaginatedLDAPProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedLDAPProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
