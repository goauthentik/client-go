/*
authentik

Making authentication simple.

API version: 2024.4.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedEventList struct for PaginatedEventList
type PaginatedEventList struct {
	Pagination Pagination `json:"pagination"`
	Results    []Event    `json:"results"`
}

// NewPaginatedEventList instantiates a new PaginatedEventList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEventList(pagination Pagination, results []Event) *PaginatedEventList {
	this := PaginatedEventList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedEventListWithDefaults instantiates a new PaginatedEventList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEventListWithDefaults() *PaginatedEventList {
	this := PaginatedEventList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedEventList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedEventList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedEventList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedEventList) GetResults() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedEventList) GetResultsOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedEventList) SetResults(v []Event) {
	o.Results = v
}

func (o PaginatedEventList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedEventList struct {
	value *PaginatedEventList
	isSet bool
}

func (v NullablePaginatedEventList) Get() *PaginatedEventList {
	return v.value
}

func (v *NullablePaginatedEventList) Set(val *PaginatedEventList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEventList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEventList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEventList(val *PaginatedEventList) *NullablePaginatedEventList {
	return &NullablePaginatedEventList{value: val, isSet: true}
}

func (v NullablePaginatedEventList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEventList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
