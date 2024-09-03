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

// PaginatedGoogleWorkspaceProviderList struct for PaginatedGoogleWorkspaceProviderList
type PaginatedGoogleWorkspaceProviderList struct {
	Pagination Pagination                `json:"pagination"`
	Results    []GoogleWorkspaceProvider `json:"results"`
}

// NewPaginatedGoogleWorkspaceProviderList instantiates a new PaginatedGoogleWorkspaceProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGoogleWorkspaceProviderList(pagination Pagination, results []GoogleWorkspaceProvider) *PaginatedGoogleWorkspaceProviderList {
	this := PaginatedGoogleWorkspaceProviderList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedGoogleWorkspaceProviderListWithDefaults instantiates a new PaginatedGoogleWorkspaceProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGoogleWorkspaceProviderListWithDefaults() *PaginatedGoogleWorkspaceProviderList {
	this := PaginatedGoogleWorkspaceProviderList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedGoogleWorkspaceProviderList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedGoogleWorkspaceProviderList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedGoogleWorkspaceProviderList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedGoogleWorkspaceProviderList) GetResults() []GoogleWorkspaceProvider {
	if o == nil {
		var ret []GoogleWorkspaceProvider
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedGoogleWorkspaceProviderList) GetResultsOk() ([]GoogleWorkspaceProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedGoogleWorkspaceProviderList) SetResults(v []GoogleWorkspaceProvider) {
	o.Results = v
}

func (o PaginatedGoogleWorkspaceProviderList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedGoogleWorkspaceProviderList struct {
	value *PaginatedGoogleWorkspaceProviderList
	isSet bool
}

func (v NullablePaginatedGoogleWorkspaceProviderList) Get() *PaginatedGoogleWorkspaceProviderList {
	return v.value
}

func (v *NullablePaginatedGoogleWorkspaceProviderList) Set(val *PaginatedGoogleWorkspaceProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGoogleWorkspaceProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGoogleWorkspaceProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGoogleWorkspaceProviderList(val *PaginatedGoogleWorkspaceProviderList) *NullablePaginatedGoogleWorkspaceProviderList {
	return &NullablePaginatedGoogleWorkspaceProviderList{value: val, isSet: true}
}

func (v NullablePaginatedGoogleWorkspaceProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGoogleWorkspaceProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
