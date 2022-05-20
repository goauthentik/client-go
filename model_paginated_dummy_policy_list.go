/*
authentik

Making authentication simple.

API version: 2022.5.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedDummyPolicyList struct for PaginatedDummyPolicyList
type PaginatedDummyPolicyList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []DummyPolicy                      `json:"results"`
}

// NewPaginatedDummyPolicyList instantiates a new PaginatedDummyPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDummyPolicyList(pagination PaginatedApplicationListPagination, results []DummyPolicy) *PaginatedDummyPolicyList {
	this := PaginatedDummyPolicyList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedDummyPolicyListWithDefaults instantiates a new PaginatedDummyPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDummyPolicyListWithDefaults() *PaginatedDummyPolicyList {
	this := PaginatedDummyPolicyList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedDummyPolicyList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedDummyPolicyList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedDummyPolicyList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedDummyPolicyList) GetResults() []DummyPolicy {
	if o == nil {
		var ret []DummyPolicy
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedDummyPolicyList) GetResultsOk() (*[]DummyPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedDummyPolicyList) SetResults(v []DummyPolicy) {
	o.Results = v
}

func (o PaginatedDummyPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedDummyPolicyList struct {
	value *PaginatedDummyPolicyList
	isSet bool
}

func (v NullablePaginatedDummyPolicyList) Get() *PaginatedDummyPolicyList {
	return v.value
}

func (v *NullablePaginatedDummyPolicyList) Set(val *PaginatedDummyPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedDummyPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedDummyPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedDummyPolicyList(val *PaginatedDummyPolicyList) *NullablePaginatedDummyPolicyList {
	return &NullablePaginatedDummyPolicyList{value: val, isSet: true}
}

func (v NullablePaginatedDummyPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedDummyPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
