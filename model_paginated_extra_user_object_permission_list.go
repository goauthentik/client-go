/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedExtraUserObjectPermissionList struct for PaginatedExtraUserObjectPermissionList
type PaginatedExtraUserObjectPermissionList struct {
	Pagination Pagination                  `json:"pagination"`
	Results    []ExtraUserObjectPermission `json:"results"`
}

// NewPaginatedExtraUserObjectPermissionList instantiates a new PaginatedExtraUserObjectPermissionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedExtraUserObjectPermissionList(pagination Pagination, results []ExtraUserObjectPermission) *PaginatedExtraUserObjectPermissionList {
	this := PaginatedExtraUserObjectPermissionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedExtraUserObjectPermissionListWithDefaults instantiates a new PaginatedExtraUserObjectPermissionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedExtraUserObjectPermissionListWithDefaults() *PaginatedExtraUserObjectPermissionList {
	this := PaginatedExtraUserObjectPermissionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedExtraUserObjectPermissionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedExtraUserObjectPermissionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedExtraUserObjectPermissionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedExtraUserObjectPermissionList) GetResults() []ExtraUserObjectPermission {
	if o == nil {
		var ret []ExtraUserObjectPermission
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedExtraUserObjectPermissionList) GetResultsOk() ([]ExtraUserObjectPermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedExtraUserObjectPermissionList) SetResults(v []ExtraUserObjectPermission) {
	o.Results = v
}

func (o PaginatedExtraUserObjectPermissionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedExtraUserObjectPermissionList struct {
	value *PaginatedExtraUserObjectPermissionList
	isSet bool
}

func (v NullablePaginatedExtraUserObjectPermissionList) Get() *PaginatedExtraUserObjectPermissionList {
	return v.value
}

func (v *NullablePaginatedExtraUserObjectPermissionList) Set(val *PaginatedExtraUserObjectPermissionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedExtraUserObjectPermissionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedExtraUserObjectPermissionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedExtraUserObjectPermissionList(val *PaginatedExtraUserObjectPermissionList) *NullablePaginatedExtraUserObjectPermissionList {
	return &NullablePaginatedExtraUserObjectPermissionList{value: val, isSet: true}
}

func (v NullablePaginatedExtraUserObjectPermissionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedExtraUserObjectPermissionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
