/*
authentik

Making authentication simple.

API version: 2023.5.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedRadiusOutpostConfigList struct for PaginatedRadiusOutpostConfigList
type PaginatedRadiusOutpostConfigList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []RadiusOutpostConfig              `json:"results"`
}

// NewPaginatedRadiusOutpostConfigList instantiates a new PaginatedRadiusOutpostConfigList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedRadiusOutpostConfigList(pagination PaginatedApplicationListPagination, results []RadiusOutpostConfig) *PaginatedRadiusOutpostConfigList {
	this := PaginatedRadiusOutpostConfigList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedRadiusOutpostConfigListWithDefaults instantiates a new PaginatedRadiusOutpostConfigList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedRadiusOutpostConfigListWithDefaults() *PaginatedRadiusOutpostConfigList {
	this := PaginatedRadiusOutpostConfigList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedRadiusOutpostConfigList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedRadiusOutpostConfigList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedRadiusOutpostConfigList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedRadiusOutpostConfigList) GetResults() []RadiusOutpostConfig {
	if o == nil {
		var ret []RadiusOutpostConfig
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedRadiusOutpostConfigList) GetResultsOk() ([]RadiusOutpostConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedRadiusOutpostConfigList) SetResults(v []RadiusOutpostConfig) {
	o.Results = v
}

func (o PaginatedRadiusOutpostConfigList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedRadiusOutpostConfigList struct {
	value *PaginatedRadiusOutpostConfigList
	isSet bool
}

func (v NullablePaginatedRadiusOutpostConfigList) Get() *PaginatedRadiusOutpostConfigList {
	return v.value
}

func (v *NullablePaginatedRadiusOutpostConfigList) Set(val *PaginatedRadiusOutpostConfigList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedRadiusOutpostConfigList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedRadiusOutpostConfigList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedRadiusOutpostConfigList(val *PaginatedRadiusOutpostConfigList) *NullablePaginatedRadiusOutpostConfigList {
	return &NullablePaginatedRadiusOutpostConfigList{value: val, isSet: true}
}

func (v NullablePaginatedRadiusOutpostConfigList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedRadiusOutpostConfigList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
