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

// PaginatedProxyOutpostConfigList struct for PaginatedProxyOutpostConfigList
type PaginatedProxyOutpostConfigList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []ProxyOutpostConfig               `json:"results"`
}

// NewPaginatedProxyOutpostConfigList instantiates a new PaginatedProxyOutpostConfigList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedProxyOutpostConfigList(pagination PaginatedApplicationListPagination, results []ProxyOutpostConfig) *PaginatedProxyOutpostConfigList {
	this := PaginatedProxyOutpostConfigList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedProxyOutpostConfigListWithDefaults instantiates a new PaginatedProxyOutpostConfigList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedProxyOutpostConfigListWithDefaults() *PaginatedProxyOutpostConfigList {
	this := PaginatedProxyOutpostConfigList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedProxyOutpostConfigList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedProxyOutpostConfigList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedProxyOutpostConfigList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedProxyOutpostConfigList) GetResults() []ProxyOutpostConfig {
	if o == nil {
		var ret []ProxyOutpostConfig
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedProxyOutpostConfigList) GetResultsOk() ([]ProxyOutpostConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedProxyOutpostConfigList) SetResults(v []ProxyOutpostConfig) {
	o.Results = v
}

func (o PaginatedProxyOutpostConfigList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedProxyOutpostConfigList struct {
	value *PaginatedProxyOutpostConfigList
	isSet bool
}

func (v NullablePaginatedProxyOutpostConfigList) Get() *PaginatedProxyOutpostConfigList {
	return v.value
}

func (v *NullablePaginatedProxyOutpostConfigList) Set(val *PaginatedProxyOutpostConfigList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedProxyOutpostConfigList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedProxyOutpostConfigList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedProxyOutpostConfigList(val *PaginatedProxyOutpostConfigList) *NullablePaginatedProxyOutpostConfigList {
	return &NullablePaginatedProxyOutpostConfigList{value: val, isSet: true}
}

func (v NullablePaginatedProxyOutpostConfigList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedProxyOutpostConfigList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
