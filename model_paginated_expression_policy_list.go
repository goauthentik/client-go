/*
authentik

Making authentication simple.

API version: 2024.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedExpressionPolicyList struct for PaginatedExpressionPolicyList
type PaginatedExpressionPolicyList struct {
	Pagination Pagination         `json:"pagination"`
	Results    []ExpressionPolicy `json:"results"`
}

// NewPaginatedExpressionPolicyList instantiates a new PaginatedExpressionPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedExpressionPolicyList(pagination Pagination, results []ExpressionPolicy) *PaginatedExpressionPolicyList {
	this := PaginatedExpressionPolicyList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedExpressionPolicyListWithDefaults instantiates a new PaginatedExpressionPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedExpressionPolicyListWithDefaults() *PaginatedExpressionPolicyList {
	this := PaginatedExpressionPolicyList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedExpressionPolicyList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedExpressionPolicyList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedExpressionPolicyList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedExpressionPolicyList) GetResults() []ExpressionPolicy {
	if o == nil {
		var ret []ExpressionPolicy
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedExpressionPolicyList) GetResultsOk() ([]ExpressionPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedExpressionPolicyList) SetResults(v []ExpressionPolicy) {
	o.Results = v
}

func (o PaginatedExpressionPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedExpressionPolicyList struct {
	value *PaginatedExpressionPolicyList
	isSet bool
}

func (v NullablePaginatedExpressionPolicyList) Get() *PaginatedExpressionPolicyList {
	return v.value
}

func (v *NullablePaginatedExpressionPolicyList) Set(val *PaginatedExpressionPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedExpressionPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedExpressionPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedExpressionPolicyList(val *PaginatedExpressionPolicyList) *NullablePaginatedExpressionPolicyList {
	return &NullablePaginatedExpressionPolicyList{value: val, isSet: true}
}

func (v NullablePaginatedExpressionPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedExpressionPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
