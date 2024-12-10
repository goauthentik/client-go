/*
authentik

Making authentication simple.

API version: 2024.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedDomainList struct for PaginatedDomainList
type PaginatedDomainList struct {
	Pagination Pagination `json:"pagination"`
	Results    []Domain   `json:"results"`
}

// NewPaginatedDomainList instantiates a new PaginatedDomainList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDomainList(pagination Pagination, results []Domain) *PaginatedDomainList {
	this := PaginatedDomainList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedDomainListWithDefaults instantiates a new PaginatedDomainList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDomainListWithDefaults() *PaginatedDomainList {
	this := PaginatedDomainList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedDomainList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedDomainList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedDomainList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedDomainList) GetResults() []Domain {
	if o == nil {
		var ret []Domain
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedDomainList) GetResultsOk() ([]Domain, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedDomainList) SetResults(v []Domain) {
	o.Results = v
}

func (o PaginatedDomainList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedDomainList struct {
	value *PaginatedDomainList
	isSet bool
}

func (v NullablePaginatedDomainList) Get() *PaginatedDomainList {
	return v.value
}

func (v *NullablePaginatedDomainList) Set(val *PaginatedDomainList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedDomainList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedDomainList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedDomainList(val *PaginatedDomainList) *NullablePaginatedDomainList {
	return &NullablePaginatedDomainList{value: val, isSet: true}
}

func (v NullablePaginatedDomainList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedDomainList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
