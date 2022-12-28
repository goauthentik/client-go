/*
authentik

Making authentication simple.

API version: 2022.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedUserConsentList struct for PaginatedUserConsentList
type PaginatedUserConsentList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []UserConsent                      `json:"results"`
}

// NewPaginatedUserConsentList instantiates a new PaginatedUserConsentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserConsentList(pagination PaginatedApplicationListPagination, results []UserConsent) *PaginatedUserConsentList {
	this := PaginatedUserConsentList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserConsentListWithDefaults instantiates a new PaginatedUserConsentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserConsentListWithDefaults() *PaginatedUserConsentList {
	this := PaginatedUserConsentList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserConsentList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserConsentList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserConsentList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserConsentList) GetResults() []UserConsent {
	if o == nil {
		var ret []UserConsent
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserConsentList) GetResultsOk() ([]UserConsent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserConsentList) SetResults(v []UserConsent) {
	o.Results = v
}

func (o PaginatedUserConsentList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedUserConsentList struct {
	value *PaginatedUserConsentList
	isSet bool
}

func (v NullablePaginatedUserConsentList) Get() *PaginatedUserConsentList {
	return v.value
}

func (v *NullablePaginatedUserConsentList) Set(val *PaginatedUserConsentList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserConsentList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserConsentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserConsentList(val *PaginatedUserConsentList) *NullablePaginatedUserConsentList {
	return &NullablePaginatedUserConsentList{value: val, isSet: true}
}

func (v NullablePaginatedUserConsentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserConsentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
