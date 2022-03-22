/*
authentik

Making authentication simple.

API version: 2022.3.3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedReputationList struct for PaginatedReputationList
type PaginatedReputationList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []Reputation                       `json:"results"`
}

// NewPaginatedReputationList instantiates a new PaginatedReputationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedReputationList(pagination PaginatedApplicationListPagination, results []Reputation) *PaginatedReputationList {
	this := PaginatedReputationList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedReputationListWithDefaults instantiates a new PaginatedReputationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedReputationListWithDefaults() *PaginatedReputationList {
	this := PaginatedReputationList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedReputationList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedReputationList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedReputationList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedReputationList) GetResults() []Reputation {
	if o == nil {
		var ret []Reputation
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedReputationList) GetResultsOk() (*[]Reputation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedReputationList) SetResults(v []Reputation) {
	o.Results = v
}

func (o PaginatedReputationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedReputationList struct {
	value *PaginatedReputationList
	isSet bool
}

func (v NullablePaginatedReputationList) Get() *PaginatedReputationList {
	return v.value
}

func (v *NullablePaginatedReputationList) Set(val *PaginatedReputationList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedReputationList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedReputationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedReputationList(val *PaginatedReputationList) *NullablePaginatedReputationList {
	return &NullablePaginatedReputationList{value: val, isSet: true}
}

func (v NullablePaginatedReputationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedReputationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}