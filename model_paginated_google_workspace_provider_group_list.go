/*
authentik

Making authentication simple.

API version: 2024.8.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedGoogleWorkspaceProviderGroupList struct for PaginatedGoogleWorkspaceProviderGroupList
type PaginatedGoogleWorkspaceProviderGroupList struct {
	Pagination Pagination                     `json:"pagination"`
	Results    []GoogleWorkspaceProviderGroup `json:"results"`
}

// NewPaginatedGoogleWorkspaceProviderGroupList instantiates a new PaginatedGoogleWorkspaceProviderGroupList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGoogleWorkspaceProviderGroupList(pagination Pagination, results []GoogleWorkspaceProviderGroup) *PaginatedGoogleWorkspaceProviderGroupList {
	this := PaginatedGoogleWorkspaceProviderGroupList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedGoogleWorkspaceProviderGroupListWithDefaults instantiates a new PaginatedGoogleWorkspaceProviderGroupList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGoogleWorkspaceProviderGroupListWithDefaults() *PaginatedGoogleWorkspaceProviderGroupList {
	this := PaginatedGoogleWorkspaceProviderGroupList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedGoogleWorkspaceProviderGroupList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedGoogleWorkspaceProviderGroupList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedGoogleWorkspaceProviderGroupList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedGoogleWorkspaceProviderGroupList) GetResults() []GoogleWorkspaceProviderGroup {
	if o == nil {
		var ret []GoogleWorkspaceProviderGroup
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedGoogleWorkspaceProviderGroupList) GetResultsOk() ([]GoogleWorkspaceProviderGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedGoogleWorkspaceProviderGroupList) SetResults(v []GoogleWorkspaceProviderGroup) {
	o.Results = v
}

func (o PaginatedGoogleWorkspaceProviderGroupList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedGoogleWorkspaceProviderGroupList struct {
	value *PaginatedGoogleWorkspaceProviderGroupList
	isSet bool
}

func (v NullablePaginatedGoogleWorkspaceProviderGroupList) Get() *PaginatedGoogleWorkspaceProviderGroupList {
	return v.value
}

func (v *NullablePaginatedGoogleWorkspaceProviderGroupList) Set(val *PaginatedGoogleWorkspaceProviderGroupList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGoogleWorkspaceProviderGroupList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGoogleWorkspaceProviderGroupList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGoogleWorkspaceProviderGroupList(val *PaginatedGoogleWorkspaceProviderGroupList) *NullablePaginatedGoogleWorkspaceProviderGroupList {
	return &NullablePaginatedGoogleWorkspaceProviderGroupList{value: val, isSet: true}
}

func (v NullablePaginatedGoogleWorkspaceProviderGroupList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGoogleWorkspaceProviderGroupList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
