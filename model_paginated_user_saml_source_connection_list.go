/*
authentik

Making authentication simple.

API version: 2022.11.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedUserSAMLSourceConnectionList struct for PaginatedUserSAMLSourceConnectionList
type PaginatedUserSAMLSourceConnectionList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []UserSAMLSourceConnection         `json:"results"`
}

// NewPaginatedUserSAMLSourceConnectionList instantiates a new PaginatedUserSAMLSourceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserSAMLSourceConnectionList(pagination PaginatedApplicationListPagination, results []UserSAMLSourceConnection) *PaginatedUserSAMLSourceConnectionList {
	this := PaginatedUserSAMLSourceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserSAMLSourceConnectionListWithDefaults instantiates a new PaginatedUserSAMLSourceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserSAMLSourceConnectionListWithDefaults() *PaginatedUserSAMLSourceConnectionList {
	this := PaginatedUserSAMLSourceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserSAMLSourceConnectionList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserSAMLSourceConnectionList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserSAMLSourceConnectionList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserSAMLSourceConnectionList) GetResults() []UserSAMLSourceConnection {
	if o == nil {
		var ret []UserSAMLSourceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserSAMLSourceConnectionList) GetResultsOk() ([]UserSAMLSourceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserSAMLSourceConnectionList) SetResults(v []UserSAMLSourceConnection) {
	o.Results = v
}

func (o PaginatedUserSAMLSourceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedUserSAMLSourceConnectionList struct {
	value *PaginatedUserSAMLSourceConnectionList
	isSet bool
}

func (v NullablePaginatedUserSAMLSourceConnectionList) Get() *PaginatedUserSAMLSourceConnectionList {
	return v.value
}

func (v *NullablePaginatedUserSAMLSourceConnectionList) Set(val *PaginatedUserSAMLSourceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserSAMLSourceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserSAMLSourceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserSAMLSourceConnectionList(val *PaginatedUserSAMLSourceConnectionList) *NullablePaginatedUserSAMLSourceConnectionList {
	return &NullablePaginatedUserSAMLSourceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedUserSAMLSourceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserSAMLSourceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
