/*
authentik

Making authentication simple.

API version: 2023.6.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedOutpostList struct for PaginatedOutpostList
type PaginatedOutpostList struct {
	Pagination Pagination `json:"pagination"`
	Results    []Outpost  `json:"results"`
}

// NewPaginatedOutpostList instantiates a new PaginatedOutpostList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedOutpostList(pagination Pagination, results []Outpost) *PaginatedOutpostList {
	this := PaginatedOutpostList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedOutpostListWithDefaults instantiates a new PaginatedOutpostList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedOutpostListWithDefaults() *PaginatedOutpostList {
	this := PaginatedOutpostList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedOutpostList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedOutpostList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedOutpostList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedOutpostList) GetResults() []Outpost {
	if o == nil {
		var ret []Outpost
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedOutpostList) GetResultsOk() ([]Outpost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedOutpostList) SetResults(v []Outpost) {
	o.Results = v
}

func (o PaginatedOutpostList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedOutpostList struct {
	value *PaginatedOutpostList
	isSet bool
}

func (v NullablePaginatedOutpostList) Get() *PaginatedOutpostList {
	return v.value
}

func (v *NullablePaginatedOutpostList) Set(val *PaginatedOutpostList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedOutpostList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedOutpostList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedOutpostList(val *PaginatedOutpostList) *NullablePaginatedOutpostList {
	return &NullablePaginatedOutpostList{value: val, isSet: true}
}

func (v NullablePaginatedOutpostList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedOutpostList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
