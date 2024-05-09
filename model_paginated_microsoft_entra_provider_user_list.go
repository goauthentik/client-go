/*
authentik

Making authentication simple.

API version: 2024.4.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedMicrosoftEntraProviderUserList struct for PaginatedMicrosoftEntraProviderUserList
type PaginatedMicrosoftEntraProviderUserList struct {
	Pagination Pagination                   `json:"pagination"`
	Results    []MicrosoftEntraProviderUser `json:"results"`
}

// NewPaginatedMicrosoftEntraProviderUserList instantiates a new PaginatedMicrosoftEntraProviderUserList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedMicrosoftEntraProviderUserList(pagination Pagination, results []MicrosoftEntraProviderUser) *PaginatedMicrosoftEntraProviderUserList {
	this := PaginatedMicrosoftEntraProviderUserList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedMicrosoftEntraProviderUserListWithDefaults instantiates a new PaginatedMicrosoftEntraProviderUserList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedMicrosoftEntraProviderUserListWithDefaults() *PaginatedMicrosoftEntraProviderUserList {
	this := PaginatedMicrosoftEntraProviderUserList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedMicrosoftEntraProviderUserList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedMicrosoftEntraProviderUserList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedMicrosoftEntraProviderUserList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedMicrosoftEntraProviderUserList) GetResults() []MicrosoftEntraProviderUser {
	if o == nil {
		var ret []MicrosoftEntraProviderUser
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedMicrosoftEntraProviderUserList) GetResultsOk() ([]MicrosoftEntraProviderUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedMicrosoftEntraProviderUserList) SetResults(v []MicrosoftEntraProviderUser) {
	o.Results = v
}

func (o PaginatedMicrosoftEntraProviderUserList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedMicrosoftEntraProviderUserList struct {
	value *PaginatedMicrosoftEntraProviderUserList
	isSet bool
}

func (v NullablePaginatedMicrosoftEntraProviderUserList) Get() *PaginatedMicrosoftEntraProviderUserList {
	return v.value
}

func (v *NullablePaginatedMicrosoftEntraProviderUserList) Set(val *PaginatedMicrosoftEntraProviderUserList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedMicrosoftEntraProviderUserList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedMicrosoftEntraProviderUserList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedMicrosoftEntraProviderUserList(val *PaginatedMicrosoftEntraProviderUserList) *NullablePaginatedMicrosoftEntraProviderUserList {
	return &NullablePaginatedMicrosoftEntraProviderUserList{value: val, isSet: true}
}

func (v NullablePaginatedMicrosoftEntraProviderUserList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedMicrosoftEntraProviderUserList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
