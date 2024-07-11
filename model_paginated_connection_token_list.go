/*
authentik

Making authentication simple.

API version: 2024.6.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedConnectionTokenList struct for PaginatedConnectionTokenList
type PaginatedConnectionTokenList struct {
	Pagination Pagination        `json:"pagination"`
	Results    []ConnectionToken `json:"results"`
}

// NewPaginatedConnectionTokenList instantiates a new PaginatedConnectionTokenList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedConnectionTokenList(pagination Pagination, results []ConnectionToken) *PaginatedConnectionTokenList {
	this := PaginatedConnectionTokenList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedConnectionTokenListWithDefaults instantiates a new PaginatedConnectionTokenList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedConnectionTokenListWithDefaults() *PaginatedConnectionTokenList {
	this := PaginatedConnectionTokenList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedConnectionTokenList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedConnectionTokenList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedConnectionTokenList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedConnectionTokenList) GetResults() []ConnectionToken {
	if o == nil {
		var ret []ConnectionToken
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedConnectionTokenList) GetResultsOk() ([]ConnectionToken, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedConnectionTokenList) SetResults(v []ConnectionToken) {
	o.Results = v
}

func (o PaginatedConnectionTokenList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedConnectionTokenList struct {
	value *PaginatedConnectionTokenList
	isSet bool
}

func (v NullablePaginatedConnectionTokenList) Get() *PaginatedConnectionTokenList {
	return v.value
}

func (v *NullablePaginatedConnectionTokenList) Set(val *PaginatedConnectionTokenList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedConnectionTokenList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedConnectionTokenList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedConnectionTokenList(val *PaginatedConnectionTokenList) *NullablePaginatedConnectionTokenList {
	return &NullablePaginatedConnectionTokenList{value: val, isSet: true}
}

func (v NullablePaginatedConnectionTokenList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedConnectionTokenList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
