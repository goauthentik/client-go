/*
authentik

Making authentication simple.

API version: 2022.9.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedLDAPOutpostConfigList struct for PaginatedLDAPOutpostConfigList
type PaginatedLDAPOutpostConfigList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []LDAPOutpostConfig                `json:"results"`
}

// NewPaginatedLDAPOutpostConfigList instantiates a new PaginatedLDAPOutpostConfigList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedLDAPOutpostConfigList(pagination PaginatedApplicationListPagination, results []LDAPOutpostConfig) *PaginatedLDAPOutpostConfigList {
	this := PaginatedLDAPOutpostConfigList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedLDAPOutpostConfigListWithDefaults instantiates a new PaginatedLDAPOutpostConfigList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedLDAPOutpostConfigListWithDefaults() *PaginatedLDAPOutpostConfigList {
	this := PaginatedLDAPOutpostConfigList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedLDAPOutpostConfigList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedLDAPOutpostConfigList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedLDAPOutpostConfigList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedLDAPOutpostConfigList) GetResults() []LDAPOutpostConfig {
	if o == nil {
		var ret []LDAPOutpostConfig
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedLDAPOutpostConfigList) GetResultsOk() ([]LDAPOutpostConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedLDAPOutpostConfigList) SetResults(v []LDAPOutpostConfig) {
	o.Results = v
}

func (o PaginatedLDAPOutpostConfigList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedLDAPOutpostConfigList struct {
	value *PaginatedLDAPOutpostConfigList
	isSet bool
}

func (v NullablePaginatedLDAPOutpostConfigList) Get() *PaginatedLDAPOutpostConfigList {
	return v.value
}

func (v *NullablePaginatedLDAPOutpostConfigList) Set(val *PaginatedLDAPOutpostConfigList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedLDAPOutpostConfigList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedLDAPOutpostConfigList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedLDAPOutpostConfigList(val *PaginatedLDAPOutpostConfigList) *NullablePaginatedLDAPOutpostConfigList {
	return &NullablePaginatedLDAPOutpostConfigList{value: val, isSet: true}
}

func (v NullablePaginatedLDAPOutpostConfigList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedLDAPOutpostConfigList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
