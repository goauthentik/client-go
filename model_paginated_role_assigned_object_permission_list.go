/*
authentik

Making authentication simple.

API version: 2025.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedRoleAssignedObjectPermissionList struct for PaginatedRoleAssignedObjectPermissionList
type PaginatedRoleAssignedObjectPermissionList struct {
	Pagination Pagination                     `json:"pagination"`
	Results    []RoleAssignedObjectPermission `json:"results"`
}

// NewPaginatedRoleAssignedObjectPermissionList instantiates a new PaginatedRoleAssignedObjectPermissionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedRoleAssignedObjectPermissionList(pagination Pagination, results []RoleAssignedObjectPermission) *PaginatedRoleAssignedObjectPermissionList {
	this := PaginatedRoleAssignedObjectPermissionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedRoleAssignedObjectPermissionListWithDefaults instantiates a new PaginatedRoleAssignedObjectPermissionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedRoleAssignedObjectPermissionListWithDefaults() *PaginatedRoleAssignedObjectPermissionList {
	this := PaginatedRoleAssignedObjectPermissionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedRoleAssignedObjectPermissionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedRoleAssignedObjectPermissionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedRoleAssignedObjectPermissionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedRoleAssignedObjectPermissionList) GetResults() []RoleAssignedObjectPermission {
	if o == nil {
		var ret []RoleAssignedObjectPermission
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedRoleAssignedObjectPermissionList) GetResultsOk() ([]RoleAssignedObjectPermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedRoleAssignedObjectPermissionList) SetResults(v []RoleAssignedObjectPermission) {
	o.Results = v
}

func (o PaginatedRoleAssignedObjectPermissionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedRoleAssignedObjectPermissionList struct {
	value *PaginatedRoleAssignedObjectPermissionList
	isSet bool
}

func (v NullablePaginatedRoleAssignedObjectPermissionList) Get() *PaginatedRoleAssignedObjectPermissionList {
	return v.value
}

func (v *NullablePaginatedRoleAssignedObjectPermissionList) Set(val *PaginatedRoleAssignedObjectPermissionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedRoleAssignedObjectPermissionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedRoleAssignedObjectPermissionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedRoleAssignedObjectPermissionList(val *PaginatedRoleAssignedObjectPermissionList) *NullablePaginatedRoleAssignedObjectPermissionList {
	return &NullablePaginatedRoleAssignedObjectPermissionList{value: val, isSet: true}
}

func (v NullablePaginatedRoleAssignedObjectPermissionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedRoleAssignedObjectPermissionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
