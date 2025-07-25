/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedPolicyList struct for PaginatedPolicyList
type PaginatedPolicyList struct {
	Pagination   Pagination             `json:"pagination"`
	Results      []Policy               `json:"results"`
	Autocomplete map[string]interface{} `json:"autocomplete"`
}

// NewPaginatedPolicyList instantiates a new PaginatedPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPolicyList(pagination Pagination, results []Policy, autocomplete map[string]interface{}) *PaginatedPolicyList {
	this := PaginatedPolicyList{}
	this.Pagination = pagination
	this.Results = results
	this.Autocomplete = autocomplete
	return &this
}

// NewPaginatedPolicyListWithDefaults instantiates a new PaginatedPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPolicyListWithDefaults() *PaginatedPolicyList {
	this := PaginatedPolicyList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPolicyList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPolicyList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPolicyList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPolicyList) GetResults() []Policy {
	if o == nil {
		var ret []Policy
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPolicyList) GetResultsOk() ([]Policy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPolicyList) SetResults(v []Policy) {
	o.Results = v
}

// GetAutocomplete returns the Autocomplete field value
func (o *PaginatedPolicyList) GetAutocomplete() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Autocomplete
}

// GetAutocompleteOk returns a tuple with the Autocomplete field value
// and a boolean to check if the value has been set.
func (o *PaginatedPolicyList) GetAutocompleteOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Autocomplete, true
}

// SetAutocomplete sets field value
func (o *PaginatedPolicyList) SetAutocomplete(v map[string]interface{}) {
	o.Autocomplete = v
}

func (o PaginatedPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if true {
		toSerialize["autocomplete"] = o.Autocomplete
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedPolicyList struct {
	value *PaginatedPolicyList
	isSet bool
}

func (v NullablePaginatedPolicyList) Get() *PaginatedPolicyList {
	return v.value
}

func (v *NullablePaginatedPolicyList) Set(val *PaginatedPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPolicyList(val *PaginatedPolicyList) *NullablePaginatedPolicyList {
	return &NullablePaginatedPolicyList{value: val, isSet: true}
}

func (v NullablePaginatedPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
