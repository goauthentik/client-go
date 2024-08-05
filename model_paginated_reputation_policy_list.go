/*
authentik

Making authentication simple.

API version: 2024.6.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedReputationPolicyList struct for PaginatedReputationPolicyList
type PaginatedReputationPolicyList struct {
	Pagination Pagination         `json:"pagination"`
	Results    []ReputationPolicy `json:"results"`
}

// NewPaginatedReputationPolicyList instantiates a new PaginatedReputationPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedReputationPolicyList(pagination Pagination, results []ReputationPolicy) *PaginatedReputationPolicyList {
	this := PaginatedReputationPolicyList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedReputationPolicyListWithDefaults instantiates a new PaginatedReputationPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedReputationPolicyListWithDefaults() *PaginatedReputationPolicyList {
	this := PaginatedReputationPolicyList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedReputationPolicyList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedReputationPolicyList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedReputationPolicyList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedReputationPolicyList) GetResults() []ReputationPolicy {
	if o == nil {
		var ret []ReputationPolicy
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedReputationPolicyList) GetResultsOk() ([]ReputationPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedReputationPolicyList) SetResults(v []ReputationPolicy) {
	o.Results = v
}

func (o PaginatedReputationPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedReputationPolicyList struct {
	value *PaginatedReputationPolicyList
	isSet bool
}

func (v NullablePaginatedReputationPolicyList) Get() *PaginatedReputationPolicyList {
	return v.value
}

func (v *NullablePaginatedReputationPolicyList) Set(val *PaginatedReputationPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedReputationPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedReputationPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedReputationPolicyList(val *PaginatedReputationPolicyList) *NullablePaginatedReputationPolicyList {
	return &NullablePaginatedReputationPolicyList{value: val, isSet: true}
}

func (v NullablePaginatedReputationPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedReputationPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
