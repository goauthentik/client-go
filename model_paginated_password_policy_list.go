/*
authentik

Making authentication simple.

API version: 2023.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedPasswordPolicyList struct for PaginatedPasswordPolicyList
type PaginatedPasswordPolicyList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []PasswordPolicy                   `json:"results"`
}

// NewPaginatedPasswordPolicyList instantiates a new PaginatedPasswordPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPasswordPolicyList(pagination PaginatedApplicationListPagination, results []PasswordPolicy) *PaginatedPasswordPolicyList {
	this := PaginatedPasswordPolicyList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedPasswordPolicyListWithDefaults instantiates a new PaginatedPasswordPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPasswordPolicyListWithDefaults() *PaginatedPasswordPolicyList {
	this := PaginatedPasswordPolicyList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPasswordPolicyList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPasswordPolicyList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPasswordPolicyList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPasswordPolicyList) GetResults() []PasswordPolicy {
	if o == nil {
		var ret []PasswordPolicy
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPasswordPolicyList) GetResultsOk() ([]PasswordPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPasswordPolicyList) SetResults(v []PasswordPolicy) {
	o.Results = v
}

func (o PaginatedPasswordPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedPasswordPolicyList struct {
	value *PaginatedPasswordPolicyList
	isSet bool
}

func (v NullablePaginatedPasswordPolicyList) Get() *PaginatedPasswordPolicyList {
	return v.value
}

func (v *NullablePaginatedPasswordPolicyList) Set(val *PaginatedPasswordPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPasswordPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPasswordPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPasswordPolicyList(val *PaginatedPasswordPolicyList) *NullablePaginatedPasswordPolicyList {
	return &NullablePaginatedPasswordPolicyList{value: val, isSet: true}
}

func (v NullablePaginatedPasswordPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPasswordPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
