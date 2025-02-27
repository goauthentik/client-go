/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedGeoIPPolicyList struct for PaginatedGeoIPPolicyList
type PaginatedGeoIPPolicyList struct {
	Pagination Pagination    `json:"pagination"`
	Results    []GeoIPPolicy `json:"results"`
}

// NewPaginatedGeoIPPolicyList instantiates a new PaginatedGeoIPPolicyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGeoIPPolicyList(pagination Pagination, results []GeoIPPolicy) *PaginatedGeoIPPolicyList {
	this := PaginatedGeoIPPolicyList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedGeoIPPolicyListWithDefaults instantiates a new PaginatedGeoIPPolicyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGeoIPPolicyListWithDefaults() *PaginatedGeoIPPolicyList {
	this := PaginatedGeoIPPolicyList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedGeoIPPolicyList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedGeoIPPolicyList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedGeoIPPolicyList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedGeoIPPolicyList) GetResults() []GeoIPPolicy {
	if o == nil {
		var ret []GeoIPPolicy
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedGeoIPPolicyList) GetResultsOk() ([]GeoIPPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedGeoIPPolicyList) SetResults(v []GeoIPPolicy) {
	o.Results = v
}

func (o PaginatedGeoIPPolicyList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedGeoIPPolicyList struct {
	value *PaginatedGeoIPPolicyList
	isSet bool
}

func (v NullablePaginatedGeoIPPolicyList) Get() *PaginatedGeoIPPolicyList {
	return v.value
}

func (v *NullablePaginatedGeoIPPolicyList) Set(val *PaginatedGeoIPPolicyList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGeoIPPolicyList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGeoIPPolicyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGeoIPPolicyList(val *PaginatedGeoIPPolicyList) *NullablePaginatedGeoIPPolicyList {
	return &NullablePaginatedGeoIPPolicyList{value: val, isSet: true}
}

func (v NullablePaginatedGeoIPPolicyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGeoIPPolicyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
