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

// PaginatedGoogleWorkspaceProviderMappingList struct for PaginatedGoogleWorkspaceProviderMappingList
type PaginatedGoogleWorkspaceProviderMappingList struct {
	Pagination   Pagination                       `json:"pagination"`
	Results      []GoogleWorkspaceProviderMapping `json:"results"`
	Autocomplete map[string]interface{}           `json:"autocomplete"`
}

// NewPaginatedGoogleWorkspaceProviderMappingList instantiates a new PaginatedGoogleWorkspaceProviderMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGoogleWorkspaceProviderMappingList(pagination Pagination, results []GoogleWorkspaceProviderMapping, autocomplete map[string]interface{}) *PaginatedGoogleWorkspaceProviderMappingList {
	this := PaginatedGoogleWorkspaceProviderMappingList{}
	this.Pagination = pagination
	this.Results = results
	this.Autocomplete = autocomplete
	return &this
}

// NewPaginatedGoogleWorkspaceProviderMappingListWithDefaults instantiates a new PaginatedGoogleWorkspaceProviderMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGoogleWorkspaceProviderMappingListWithDefaults() *PaginatedGoogleWorkspaceProviderMappingList {
	this := PaginatedGoogleWorkspaceProviderMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedGoogleWorkspaceProviderMappingList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedGoogleWorkspaceProviderMappingList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedGoogleWorkspaceProviderMappingList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedGoogleWorkspaceProviderMappingList) GetResults() []GoogleWorkspaceProviderMapping {
	if o == nil {
		var ret []GoogleWorkspaceProviderMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedGoogleWorkspaceProviderMappingList) GetResultsOk() ([]GoogleWorkspaceProviderMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedGoogleWorkspaceProviderMappingList) SetResults(v []GoogleWorkspaceProviderMapping) {
	o.Results = v
}

// GetAutocomplete returns the Autocomplete field value
func (o *PaginatedGoogleWorkspaceProviderMappingList) GetAutocomplete() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Autocomplete
}

// GetAutocompleteOk returns a tuple with the Autocomplete field value
// and a boolean to check if the value has been set.
func (o *PaginatedGoogleWorkspaceProviderMappingList) GetAutocompleteOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Autocomplete, true
}

// SetAutocomplete sets field value
func (o *PaginatedGoogleWorkspaceProviderMappingList) SetAutocomplete(v map[string]interface{}) {
	o.Autocomplete = v
}

func (o PaginatedGoogleWorkspaceProviderMappingList) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedGoogleWorkspaceProviderMappingList struct {
	value *PaginatedGoogleWorkspaceProviderMappingList
	isSet bool
}

func (v NullablePaginatedGoogleWorkspaceProviderMappingList) Get() *PaginatedGoogleWorkspaceProviderMappingList {
	return v.value
}

func (v *NullablePaginatedGoogleWorkspaceProviderMappingList) Set(val *PaginatedGoogleWorkspaceProviderMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGoogleWorkspaceProviderMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGoogleWorkspaceProviderMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGoogleWorkspaceProviderMappingList(val *PaginatedGoogleWorkspaceProviderMappingList) *NullablePaginatedGoogleWorkspaceProviderMappingList {
	return &NullablePaginatedGoogleWorkspaceProviderMappingList{value: val, isSet: true}
}

func (v NullablePaginatedGoogleWorkspaceProviderMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGoogleWorkspaceProviderMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
