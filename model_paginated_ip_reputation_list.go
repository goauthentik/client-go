/*
authentik

Making authentication simple.

API version: 2021.8.5
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedIPReputationList struct for PaginatedIPReputationList
type PaginatedIPReputationList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results []IPReputation `json:"results"`
}

// NewPaginatedIPReputationList instantiates a new PaginatedIPReputationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedIPReputationList(pagination PaginatedApplicationListPagination, results []IPReputation) *PaginatedIPReputationList {
	this := PaginatedIPReputationList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedIPReputationListWithDefaults instantiates a new PaginatedIPReputationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedIPReputationListWithDefaults() *PaginatedIPReputationList {
	this := PaginatedIPReputationList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedIPReputationList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedIPReputationList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedIPReputationList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedIPReputationList) GetResults() []IPReputation {
	if o == nil {
		var ret []IPReputation
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedIPReputationList) GetResultsOk() (*[]IPReputation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedIPReputationList) SetResults(v []IPReputation) {
	o.Results = v
}

func (o PaginatedIPReputationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedIPReputationList struct {
	value *PaginatedIPReputationList
	isSet bool
}

func (v NullablePaginatedIPReputationList) Get() *PaginatedIPReputationList {
	return v.value
}

func (v *NullablePaginatedIPReputationList) Set(val *PaginatedIPReputationList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedIPReputationList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedIPReputationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedIPReputationList(val *PaginatedIPReputationList) *NullablePaginatedIPReputationList {
	return &NullablePaginatedIPReputationList{value: val, isSet: true}
}

func (v NullablePaginatedIPReputationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedIPReputationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


