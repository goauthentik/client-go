/*
authentik

Making authentication simple.

API version: 2024.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedOAuthSourcePropertyMappingList struct for PaginatedOAuthSourcePropertyMappingList
type PaginatedOAuthSourcePropertyMappingList struct {
	Pagination Pagination                   `json:"pagination"`
	Results    []OAuthSourcePropertyMapping `json:"results"`
}

// NewPaginatedOAuthSourcePropertyMappingList instantiates a new PaginatedOAuthSourcePropertyMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedOAuthSourcePropertyMappingList(pagination Pagination, results []OAuthSourcePropertyMapping) *PaginatedOAuthSourcePropertyMappingList {
	this := PaginatedOAuthSourcePropertyMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedOAuthSourcePropertyMappingListWithDefaults instantiates a new PaginatedOAuthSourcePropertyMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedOAuthSourcePropertyMappingListWithDefaults() *PaginatedOAuthSourcePropertyMappingList {
	this := PaginatedOAuthSourcePropertyMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedOAuthSourcePropertyMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedOAuthSourcePropertyMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedOAuthSourcePropertyMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedOAuthSourcePropertyMappingList) GetResults() []OAuthSourcePropertyMapping {
	if o == nil {
		var ret []OAuthSourcePropertyMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedOAuthSourcePropertyMappingList) GetResultsOk() ([]OAuthSourcePropertyMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedOAuthSourcePropertyMappingList) SetResults(v []OAuthSourcePropertyMapping) {
	o.Results = v
}

func (o PaginatedOAuthSourcePropertyMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedOAuthSourcePropertyMappingList struct {
	value *PaginatedOAuthSourcePropertyMappingList
	isSet bool
}

func (v NullablePaginatedOAuthSourcePropertyMappingList) Get() *PaginatedOAuthSourcePropertyMappingList {
	return v.value
}

func (v *NullablePaginatedOAuthSourcePropertyMappingList) Set(val *PaginatedOAuthSourcePropertyMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedOAuthSourcePropertyMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedOAuthSourcePropertyMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedOAuthSourcePropertyMappingList(val *PaginatedOAuthSourcePropertyMappingList) *NullablePaginatedOAuthSourcePropertyMappingList {
	return &NullablePaginatedOAuthSourcePropertyMappingList{value: val, isSet: true}
}

func (v NullablePaginatedOAuthSourcePropertyMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedOAuthSourcePropertyMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
