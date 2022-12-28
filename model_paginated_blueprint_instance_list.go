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

// PaginatedBlueprintInstanceList struct for PaginatedBlueprintInstanceList
type PaginatedBlueprintInstanceList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []BlueprintInstance                `json:"results"`
}

// NewPaginatedBlueprintInstanceList instantiates a new PaginatedBlueprintInstanceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedBlueprintInstanceList(pagination PaginatedApplicationListPagination, results []BlueprintInstance) *PaginatedBlueprintInstanceList {
	this := PaginatedBlueprintInstanceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedBlueprintInstanceListWithDefaults instantiates a new PaginatedBlueprintInstanceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedBlueprintInstanceListWithDefaults() *PaginatedBlueprintInstanceList {
	this := PaginatedBlueprintInstanceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedBlueprintInstanceList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedBlueprintInstanceList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedBlueprintInstanceList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedBlueprintInstanceList) GetResults() []BlueprintInstance {
	if o == nil {
		var ret []BlueprintInstance
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedBlueprintInstanceList) GetResultsOk() ([]BlueprintInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedBlueprintInstanceList) SetResults(v []BlueprintInstance) {
	o.Results = v
}

func (o PaginatedBlueprintInstanceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedBlueprintInstanceList struct {
	value *PaginatedBlueprintInstanceList
	isSet bool
}

func (v NullablePaginatedBlueprintInstanceList) Get() *PaginatedBlueprintInstanceList {
	return v.value
}

func (v *NullablePaginatedBlueprintInstanceList) Set(val *PaginatedBlueprintInstanceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedBlueprintInstanceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedBlueprintInstanceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedBlueprintInstanceList(val *PaginatedBlueprintInstanceList) *NullablePaginatedBlueprintInstanceList {
	return &NullablePaginatedBlueprintInstanceList{value: val, isSet: true}
}

func (v NullablePaginatedBlueprintInstanceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedBlueprintInstanceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
