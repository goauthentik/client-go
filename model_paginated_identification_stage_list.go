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

// PaginatedIdentificationStageList struct for PaginatedIdentificationStageList
type PaginatedIdentificationStageList struct {
	Pagination   Pagination             `json:"pagination"`
	Results      []IdentificationStage  `json:"results"`
	Autocomplete map[string]interface{} `json:"autocomplete"`
}

// NewPaginatedIdentificationStageList instantiates a new PaginatedIdentificationStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedIdentificationStageList(pagination Pagination, results []IdentificationStage, autocomplete map[string]interface{}) *PaginatedIdentificationStageList {
	this := PaginatedIdentificationStageList{}
	this.Pagination = pagination
	this.Results = results
	this.Autocomplete = autocomplete
	return &this
}

// NewPaginatedIdentificationStageListWithDefaults instantiates a new PaginatedIdentificationStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedIdentificationStageListWithDefaults() *PaginatedIdentificationStageList {
	this := PaginatedIdentificationStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedIdentificationStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedIdentificationStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedIdentificationStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedIdentificationStageList) GetResults() []IdentificationStage {
	if o == nil {
		var ret []IdentificationStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedIdentificationStageList) GetResultsOk() ([]IdentificationStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedIdentificationStageList) SetResults(v []IdentificationStage) {
	o.Results = v
}

// GetAutocomplete returns the Autocomplete field value
func (o *PaginatedIdentificationStageList) GetAutocomplete() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Autocomplete
}

// GetAutocompleteOk returns a tuple with the Autocomplete field value
// and a boolean to check if the value has been set.
func (o *PaginatedIdentificationStageList) GetAutocompleteOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Autocomplete, true
}

// SetAutocomplete sets field value
func (o *PaginatedIdentificationStageList) SetAutocomplete(v map[string]interface{}) {
	o.Autocomplete = v
}

func (o PaginatedIdentificationStageList) MarshalJSON() ([]byte, error) {
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

type NullablePaginatedIdentificationStageList struct {
	value *PaginatedIdentificationStageList
	isSet bool
}

func (v NullablePaginatedIdentificationStageList) Get() *PaginatedIdentificationStageList {
	return v.value
}

func (v *NullablePaginatedIdentificationStageList) Set(val *PaginatedIdentificationStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedIdentificationStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedIdentificationStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedIdentificationStageList(val *PaginatedIdentificationStageList) *NullablePaginatedIdentificationStageList {
	return &NullablePaginatedIdentificationStageList{value: val, isSet: true}
}

func (v NullablePaginatedIdentificationStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedIdentificationStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
