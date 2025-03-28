/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedSCIMProviderUserList struct for PaginatedSCIMProviderUserList
type PaginatedSCIMProviderUserList struct {
	Pagination Pagination         `json:"pagination"`
	Results    []SCIMProviderUser `json:"results"`
}

// NewPaginatedSCIMProviderUserList instantiates a new PaginatedSCIMProviderUserList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSCIMProviderUserList(pagination Pagination, results []SCIMProviderUser) *PaginatedSCIMProviderUserList {
	this := PaginatedSCIMProviderUserList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSCIMProviderUserListWithDefaults instantiates a new PaginatedSCIMProviderUserList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSCIMProviderUserListWithDefaults() *PaginatedSCIMProviderUserList {
	this := PaginatedSCIMProviderUserList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSCIMProviderUserList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMProviderUserList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSCIMProviderUserList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSCIMProviderUserList) GetResults() []SCIMProviderUser {
	if o == nil {
		var ret []SCIMProviderUser
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSCIMProviderUserList) GetResultsOk() ([]SCIMProviderUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSCIMProviderUserList) SetResults(v []SCIMProviderUser) {
	o.Results = v
}

func (o PaginatedSCIMProviderUserList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedSCIMProviderUserList struct {
	value *PaginatedSCIMProviderUserList
	isSet bool
}

func (v NullablePaginatedSCIMProviderUserList) Get() *PaginatedSCIMProviderUserList {
	return v.value
}

func (v *NullablePaginatedSCIMProviderUserList) Set(val *PaginatedSCIMProviderUserList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSCIMProviderUserList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSCIMProviderUserList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSCIMProviderUserList(val *PaginatedSCIMProviderUserList) *NullablePaginatedSCIMProviderUserList {
	return &NullablePaginatedSCIMProviderUserList{value: val, isSet: true}
}

func (v NullablePaginatedSCIMProviderUserList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSCIMProviderUserList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
