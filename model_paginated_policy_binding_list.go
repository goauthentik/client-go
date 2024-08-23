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

// PaginatedPolicyBindingList struct for PaginatedPolicyBindingList
type PaginatedPolicyBindingList struct {
	Pagination Pagination      `json:"pagination"`
	Results    []PolicyBinding `json:"results"`
}

// NewPaginatedPolicyBindingList instantiates a new PaginatedPolicyBindingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPolicyBindingList(pagination Pagination, results []PolicyBinding) *PaginatedPolicyBindingList {
	this := PaginatedPolicyBindingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedPolicyBindingListWithDefaults instantiates a new PaginatedPolicyBindingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPolicyBindingListWithDefaults() *PaginatedPolicyBindingList {
	this := PaginatedPolicyBindingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPolicyBindingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPolicyBindingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPolicyBindingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPolicyBindingList) GetResults() []PolicyBinding {
	if o == nil {
		var ret []PolicyBinding
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPolicyBindingList) GetResultsOk() ([]PolicyBinding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPolicyBindingList) SetResults(v []PolicyBinding) {
	o.Results = v
}

func (o PaginatedPolicyBindingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedPolicyBindingList struct {
	value *PaginatedPolicyBindingList
	isSet bool
}

func (v NullablePaginatedPolicyBindingList) Get() *PaginatedPolicyBindingList {
	return v.value
}

func (v *NullablePaginatedPolicyBindingList) Set(val *PaginatedPolicyBindingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPolicyBindingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPolicyBindingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPolicyBindingList(val *PaginatedPolicyBindingList) *NullablePaginatedPolicyBindingList {
	return &NullablePaginatedPolicyBindingList{value: val, isSet: true}
}

func (v NullablePaginatedPolicyBindingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPolicyBindingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
