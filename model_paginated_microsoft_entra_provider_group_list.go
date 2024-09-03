/*
authentik

Making authentication simple.

API version: 2024.8.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedMicrosoftEntraProviderGroupList struct for PaginatedMicrosoftEntraProviderGroupList
type PaginatedMicrosoftEntraProviderGroupList struct {
	Pagination Pagination                    `json:"pagination"`
	Results    []MicrosoftEntraProviderGroup `json:"results"`
}

// NewPaginatedMicrosoftEntraProviderGroupList instantiates a new PaginatedMicrosoftEntraProviderGroupList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedMicrosoftEntraProviderGroupList(pagination Pagination, results []MicrosoftEntraProviderGroup) *PaginatedMicrosoftEntraProviderGroupList {
	this := PaginatedMicrosoftEntraProviderGroupList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedMicrosoftEntraProviderGroupListWithDefaults instantiates a new PaginatedMicrosoftEntraProviderGroupList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedMicrosoftEntraProviderGroupListWithDefaults() *PaginatedMicrosoftEntraProviderGroupList {
	this := PaginatedMicrosoftEntraProviderGroupList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedMicrosoftEntraProviderGroupList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedMicrosoftEntraProviderGroupList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedMicrosoftEntraProviderGroupList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedMicrosoftEntraProviderGroupList) GetResults() []MicrosoftEntraProviderGroup {
	if o == nil {
		var ret []MicrosoftEntraProviderGroup
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedMicrosoftEntraProviderGroupList) GetResultsOk() ([]MicrosoftEntraProviderGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedMicrosoftEntraProviderGroupList) SetResults(v []MicrosoftEntraProviderGroup) {
	o.Results = v
}

func (o PaginatedMicrosoftEntraProviderGroupList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedMicrosoftEntraProviderGroupList struct {
	value *PaginatedMicrosoftEntraProviderGroupList
	isSet bool
}

func (v NullablePaginatedMicrosoftEntraProviderGroupList) Get() *PaginatedMicrosoftEntraProviderGroupList {
	return v.value
}

func (v *NullablePaginatedMicrosoftEntraProviderGroupList) Set(val *PaginatedMicrosoftEntraProviderGroupList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedMicrosoftEntraProviderGroupList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedMicrosoftEntraProviderGroupList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedMicrosoftEntraProviderGroupList(val *PaginatedMicrosoftEntraProviderGroupList) *NullablePaginatedMicrosoftEntraProviderGroupList {
	return &NullablePaginatedMicrosoftEntraProviderGroupList{value: val, isSet: true}
}

func (v NullablePaginatedMicrosoftEntraProviderGroupList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedMicrosoftEntraProviderGroupList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
