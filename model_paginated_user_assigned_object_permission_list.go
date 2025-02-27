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

// PaginatedUserAssignedObjectPermissionList struct for PaginatedUserAssignedObjectPermissionList
type PaginatedUserAssignedObjectPermissionList struct {
	Pagination Pagination                     `json:"pagination"`
	Results    []UserAssignedObjectPermission `json:"results"`
}

// NewPaginatedUserAssignedObjectPermissionList instantiates a new PaginatedUserAssignedObjectPermissionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserAssignedObjectPermissionList(pagination Pagination, results []UserAssignedObjectPermission) *PaginatedUserAssignedObjectPermissionList {
	this := PaginatedUserAssignedObjectPermissionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserAssignedObjectPermissionListWithDefaults instantiates a new PaginatedUserAssignedObjectPermissionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserAssignedObjectPermissionListWithDefaults() *PaginatedUserAssignedObjectPermissionList {
	this := PaginatedUserAssignedObjectPermissionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserAssignedObjectPermissionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserAssignedObjectPermissionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserAssignedObjectPermissionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserAssignedObjectPermissionList) GetResults() []UserAssignedObjectPermission {
	if o == nil {
		var ret []UserAssignedObjectPermission
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserAssignedObjectPermissionList) GetResultsOk() ([]UserAssignedObjectPermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserAssignedObjectPermissionList) SetResults(v []UserAssignedObjectPermission) {
	o.Results = v
}

func (o PaginatedUserAssignedObjectPermissionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedUserAssignedObjectPermissionList struct {
	value *PaginatedUserAssignedObjectPermissionList
	isSet bool
}

func (v NullablePaginatedUserAssignedObjectPermissionList) Get() *PaginatedUserAssignedObjectPermissionList {
	return v.value
}

func (v *NullablePaginatedUserAssignedObjectPermissionList) Set(val *PaginatedUserAssignedObjectPermissionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserAssignedObjectPermissionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserAssignedObjectPermissionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserAssignedObjectPermissionList(val *PaginatedUserAssignedObjectPermissionList) *NullablePaginatedUserAssignedObjectPermissionList {
	return &NullablePaginatedUserAssignedObjectPermissionList{value: val, isSet: true}
}

func (v NullablePaginatedUserAssignedObjectPermissionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserAssignedObjectPermissionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
