/*
authentik

Making authentication simple.

API version: 2021.9.1-rc1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedTenantList struct for PaginatedTenantList
type PaginatedTenantList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results []Tenant `json:"results"`
}

// NewPaginatedTenantList instantiates a new PaginatedTenantList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedTenantList(pagination PaginatedApplicationListPagination, results []Tenant) *PaginatedTenantList {
	this := PaginatedTenantList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedTenantListWithDefaults instantiates a new PaginatedTenantList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedTenantListWithDefaults() *PaginatedTenantList {
	this := PaginatedTenantList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedTenantList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedTenantList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedTenantList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedTenantList) GetResults() []Tenant {
	if o == nil {
		var ret []Tenant
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedTenantList) GetResultsOk() (*[]Tenant, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedTenantList) SetResults(v []Tenant) {
	o.Results = v
}

func (o PaginatedTenantList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedTenantList struct {
	value *PaginatedTenantList
	isSet bool
}

func (v NullablePaginatedTenantList) Get() *PaginatedTenantList {
	return v.value
}

func (v *NullablePaginatedTenantList) Set(val *PaginatedTenantList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedTenantList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedTenantList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedTenantList(val *PaginatedTenantList) *NullablePaginatedTenantList {
	return &NullablePaginatedTenantList{value: val, isSet: true}
}

func (v NullablePaginatedTenantList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedTenantList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


