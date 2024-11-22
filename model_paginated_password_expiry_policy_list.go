/*
authentik

Making authentication simple.

API version: 2024.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedPasswordExpiryPolicyList struct for PaginatedPasswordExpiryPolicyList
type PaginatedPasswordExpiryPolicyList struct {
	Pagination Pagination             `json:"pagination"`
	Results    []PasswordExpiryPolicy `json:"results"`
}

// NewPaginatedPasswordExpiryPolicyList instantiates a new PaginatedPasswordExpiryPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPasswordExpiryPolicyList(pagination Pagination, results []PasswordExpiryPolicy) *PaginatedPasswordExpiryPolicyList {
	this := PaginatedPasswordExpiryPolicyList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedPasswordExpiryPolicyListWithDefaults instantiates a new PaginatedPasswordExpiryPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPasswordExpiryPolicyListWithDefaults() *PaginatedPasswordExpiryPolicyList {
	this := PaginatedPasswordExpiryPolicyList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPasswordExpiryPolicyList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPasswordExpiryPolicyList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPasswordExpiryPolicyList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPasswordExpiryPolicyList) GetResults() []PasswordExpiryPolicy {
	if o == nil {
		var ret []PasswordExpiryPolicy
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPasswordExpiryPolicyList) GetResultsOk() ([]PasswordExpiryPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPasswordExpiryPolicyList) SetResults(v []PasswordExpiryPolicy) {
	o.Results = v
}

func (o PaginatedPasswordExpiryPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedPasswordExpiryPolicyList struct {
	value *PaginatedPasswordExpiryPolicyList
	isSet bool
}

func (v NullablePaginatedPasswordExpiryPolicyList) Get() *PaginatedPasswordExpiryPolicyList {
	return v.value
}

func (v *NullablePaginatedPasswordExpiryPolicyList) Set(val *PaginatedPasswordExpiryPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPasswordExpiryPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPasswordExpiryPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPasswordExpiryPolicyList(val *PaginatedPasswordExpiryPolicyList) *NullablePaginatedPasswordExpiryPolicyList {
	return &NullablePaginatedPasswordExpiryPolicyList{value: val, isSet: true}
}

func (v NullablePaginatedPasswordExpiryPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPasswordExpiryPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
