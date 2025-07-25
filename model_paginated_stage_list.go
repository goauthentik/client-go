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

// PaginatedStageList struct for PaginatedStageList
type PaginatedStageList struct {
	Pagination   Pagination             `json:"pagination"`
	Results      []Stage                `json:"results"`
	Autocomplete map[string]interface{} `json:"autocomplete"`
}

// NewPaginatedStageList instantiates a new PaginatedStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedStageList(pagination Pagination, results []Stage, autocomplete map[string]interface{}) *PaginatedStageList {
	this := PaginatedStageList{}
	this.Pagination = pagination
	this.Results = results
	this.Autocomplete = autocomplete
	return &this
}

// NewPaginatedStageListWithDefaults instantiates a new PaginatedStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedStageListWithDefaults() *PaginatedStageList {
	this := PaginatedStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedStageList) GetResults() []Stage {
	if o == nil {
		var ret []Stage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedStageList) GetResultsOk() ([]Stage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedStageList) SetResults(v []Stage) {
	o.Results = v
}

// GetAutocomplete returns the Autocomplete field value
func (o *PaginatedStageList) GetAutocomplete() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Autocomplete
}

// GetAutocompleteOk returns a tuple with the Autocomplete field value
// and a boolean to check if the value has been set.
func (o *PaginatedStageList) GetAutocompleteOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Autocomplete, true
}

// SetAutocomplete sets field value
func (o *PaginatedStageList) SetAutocomplete(v map[string]interface{}) {
	o.Autocomplete = v
}

func (o PaginatedStageList) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedStageList struct {
	value *PaginatedStageList
	isSet bool
}

func (v NullablePaginatedStageList) Get() *PaginatedStageList {
	return v.value
}

func (v *NullablePaginatedStageList) Set(val *PaginatedStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedStageList(val *PaginatedStageList) *NullablePaginatedStageList {
	return &NullablePaginatedStageList{value: val, isSet: true}
}

func (v NullablePaginatedStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
