/*
authentik

Making authentication simple.

API version: 2023.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedExtraRoleObjectPermissionList struct for PaginatedExtraRoleObjectPermissionList
type PaginatedExtraRoleObjectPermissionList struct {
	Pagination Pagination                  `json:"pagination"`
	Results    []ExtraRoleObjectPermission `json:"results"`
}

// NewPaginatedExtraRoleObjectPermissionList instantiates a new PaginatedExtraRoleObjectPermissionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedExtraRoleObjectPermissionList(pagination Pagination, results []ExtraRoleObjectPermission) *PaginatedExtraRoleObjectPermissionList {
	this := PaginatedExtraRoleObjectPermissionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedExtraRoleObjectPermissionListWithDefaults instantiates a new PaginatedExtraRoleObjectPermissionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedExtraRoleObjectPermissionListWithDefaults() *PaginatedExtraRoleObjectPermissionList {
	this := PaginatedExtraRoleObjectPermissionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedExtraRoleObjectPermissionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedExtraRoleObjectPermissionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedExtraRoleObjectPermissionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedExtraRoleObjectPermissionList) GetResults() []ExtraRoleObjectPermission {
	if o == nil {
		var ret []ExtraRoleObjectPermission
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedExtraRoleObjectPermissionList) GetResultsOk() ([]ExtraRoleObjectPermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedExtraRoleObjectPermissionList) SetResults(v []ExtraRoleObjectPermission) {
	o.Results = v
}

func (o PaginatedExtraRoleObjectPermissionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedExtraRoleObjectPermissionList struct {
	value *PaginatedExtraRoleObjectPermissionList
	isSet bool
}

func (v NullablePaginatedExtraRoleObjectPermissionList) Get() *PaginatedExtraRoleObjectPermissionList {
	return v.value
}

func (v *NullablePaginatedExtraRoleObjectPermissionList) Set(val *PaginatedExtraRoleObjectPermissionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedExtraRoleObjectPermissionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedExtraRoleObjectPermissionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedExtraRoleObjectPermissionList(val *PaginatedExtraRoleObjectPermissionList) *NullablePaginatedExtraRoleObjectPermissionList {
	return &NullablePaginatedExtraRoleObjectPermissionList{value: val, isSet: true}
}

func (v NullablePaginatedExtraRoleObjectPermissionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedExtraRoleObjectPermissionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
