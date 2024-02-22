/*
authentik

Making authentication simple.

API version: 2024.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedPermissionList struct for PaginatedPermissionList
type PaginatedPermissionList struct {
	Pagination Pagination   `json:"pagination"`
	Results    []Permission `json:"results"`
}

// NewPaginatedPermissionList instantiates a new PaginatedPermissionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedPermissionList(pagination Pagination, results []Permission) *PaginatedPermissionList {
	this := PaginatedPermissionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedPermissionListWithDefaults instantiates a new PaginatedPermissionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedPermissionListWithDefaults() *PaginatedPermissionList {
	this := PaginatedPermissionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedPermissionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedPermissionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedPermissionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedPermissionList) GetResults() []Permission {
	if o == nil {
		var ret []Permission
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedPermissionList) GetResultsOk() ([]Permission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedPermissionList) SetResults(v []Permission) {
	o.Results = v
}

func (o PaginatedPermissionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedPermissionList struct {
	value *PaginatedPermissionList
	isSet bool
}

func (v NullablePaginatedPermissionList) Get() *PaginatedPermissionList {
	return v.value
}

func (v *NullablePaginatedPermissionList) Set(val *PaginatedPermissionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedPermissionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedPermissionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedPermissionList(val *PaginatedPermissionList) *NullablePaginatedPermissionList {
	return &NullablePaginatedPermissionList{value: val, isSet: true}
}

func (v NullablePaginatedPermissionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedPermissionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
