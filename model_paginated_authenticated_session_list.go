/*
authentik

Making authentication simple.

API version: 2022.12.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedAuthenticatedSessionList struct for PaginatedAuthenticatedSessionList
type PaginatedAuthenticatedSessionList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []AuthenticatedSession             `json:"results"`
}

// NewPaginatedAuthenticatedSessionList instantiates a new PaginatedAuthenticatedSessionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAuthenticatedSessionList(pagination PaginatedApplicationListPagination, results []AuthenticatedSession) *PaginatedAuthenticatedSessionList {
	this := PaginatedAuthenticatedSessionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedAuthenticatedSessionListWithDefaults instantiates a new PaginatedAuthenticatedSessionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAuthenticatedSessionListWithDefaults() *PaginatedAuthenticatedSessionList {
	this := PaginatedAuthenticatedSessionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedAuthenticatedSessionList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatedSessionList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedAuthenticatedSessionList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedAuthenticatedSessionList) GetResults() []AuthenticatedSession {
	if o == nil {
		var ret []AuthenticatedSession
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatedSessionList) GetResultsOk() ([]AuthenticatedSession, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAuthenticatedSessionList) SetResults(v []AuthenticatedSession) {
	o.Results = v
}

func (o PaginatedAuthenticatedSessionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedAuthenticatedSessionList struct {
	value *PaginatedAuthenticatedSessionList
	isSet bool
}

func (v NullablePaginatedAuthenticatedSessionList) Get() *PaginatedAuthenticatedSessionList {
	return v.value
}

func (v *NullablePaginatedAuthenticatedSessionList) Set(val *PaginatedAuthenticatedSessionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAuthenticatedSessionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAuthenticatedSessionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAuthenticatedSessionList(val *PaginatedAuthenticatedSessionList) *NullablePaginatedAuthenticatedSessionList {
	return &NullablePaginatedAuthenticatedSessionList{value: val, isSet: true}
}

func (v NullablePaginatedAuthenticatedSessionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAuthenticatedSessionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
