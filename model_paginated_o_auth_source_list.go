/*
authentik

Making authentication simple.

API version: 2023.5.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedOAuthSourceList struct for PaginatedOAuthSourceList
type PaginatedOAuthSourceList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []OAuthSource                      `json:"results"`
}

// NewPaginatedOAuthSourceList instantiates a new PaginatedOAuthSourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedOAuthSourceList(pagination PaginatedApplicationListPagination, results []OAuthSource) *PaginatedOAuthSourceList {
	this := PaginatedOAuthSourceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedOAuthSourceListWithDefaults instantiates a new PaginatedOAuthSourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedOAuthSourceListWithDefaults() *PaginatedOAuthSourceList {
	this := PaginatedOAuthSourceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedOAuthSourceList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedOAuthSourceList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedOAuthSourceList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedOAuthSourceList) GetResults() []OAuthSource {
	if o == nil {
		var ret []OAuthSource
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedOAuthSourceList) GetResultsOk() ([]OAuthSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedOAuthSourceList) SetResults(v []OAuthSource) {
	o.Results = v
}

func (o PaginatedOAuthSourceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedOAuthSourceList struct {
	value *PaginatedOAuthSourceList
	isSet bool
}

func (v NullablePaginatedOAuthSourceList) Get() *PaginatedOAuthSourceList {
	return v.value
}

func (v *NullablePaginatedOAuthSourceList) Set(val *PaginatedOAuthSourceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedOAuthSourceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedOAuthSourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedOAuthSourceList(val *PaginatedOAuthSourceList) *NullablePaginatedOAuthSourceList {
	return &NullablePaginatedOAuthSourceList{value: val, isSet: true}
}

func (v NullablePaginatedOAuthSourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedOAuthSourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
