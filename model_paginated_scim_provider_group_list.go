/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSCIMProviderGroupList struct for PaginatedSCIMProviderGroupList
type PaginatedSCIMProviderGroupList struct {
	Pagination Pagination          `json:"pagination"`
	Results    []SCIMProviderGroup `json:"results"`
}

// NewPaginatedSCIMProviderGroupList instantiates a new PaginatedSCIMProviderGroupList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSCIMProviderGroupList(pagination Pagination, results []SCIMProviderGroup) *PaginatedSCIMProviderGroupList {
	this := PaginatedSCIMProviderGroupList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSCIMProviderGroupListWithDefaults instantiates a new PaginatedSCIMProviderGroupList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSCIMProviderGroupListWithDefaults() *PaginatedSCIMProviderGroupList {
	this := PaginatedSCIMProviderGroupList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSCIMProviderGroupList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMProviderGroupList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSCIMProviderGroupList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSCIMProviderGroupList) GetResults() []SCIMProviderGroup {
	if o == nil {
		var ret []SCIMProviderGroup
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMProviderGroupList) GetResultsOk() ([]SCIMProviderGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSCIMProviderGroupList) SetResults(v []SCIMProviderGroup) {
	o.Results = v
}

func (o PaginatedSCIMProviderGroupList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSCIMProviderGroupList struct {
	value *PaginatedSCIMProviderGroupList
	isSet bool
}

func (v NullablePaginatedSCIMProviderGroupList) Get() *PaginatedSCIMProviderGroupList {
	return v.value
}

func (v *NullablePaginatedSCIMProviderGroupList) Set(val *PaginatedSCIMProviderGroupList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSCIMProviderGroupList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSCIMProviderGroupList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSCIMProviderGroupList(val *PaginatedSCIMProviderGroupList) *NullablePaginatedSCIMProviderGroupList {
	return &NullablePaginatedSCIMProviderGroupList{value: val, isSet: true}
}

func (v NullablePaginatedSCIMProviderGroupList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSCIMProviderGroupList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
