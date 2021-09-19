/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedAuthenticatorStaticStageList struct for PaginatedAuthenticatorStaticStageList
type PaginatedAuthenticatorStaticStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results []AuthenticatorStaticStage `json:"results"`
}

// NewPaginatedAuthenticatorStaticStageList instantiates a new PaginatedAuthenticatorStaticStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAuthenticatorStaticStageList(pagination PaginatedApplicationListPagination, results []AuthenticatorStaticStage) *PaginatedAuthenticatorStaticStageList {
	this := PaginatedAuthenticatorStaticStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedAuthenticatorStaticStageListWithDefaults instantiates a new PaginatedAuthenticatorStaticStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAuthenticatorStaticStageListWithDefaults() *PaginatedAuthenticatorStaticStageList {
	this := PaginatedAuthenticatorStaticStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedAuthenticatorStaticStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorStaticStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedAuthenticatorStaticStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedAuthenticatorStaticStageList) GetResults() []AuthenticatorStaticStage {
	if o == nil {
		var ret []AuthenticatorStaticStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorStaticStageList) GetResultsOk() (*[]AuthenticatorStaticStage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedAuthenticatorStaticStageList) SetResults(v []AuthenticatorStaticStage) {
	o.Results = v
}

func (o PaginatedAuthenticatorStaticStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedAuthenticatorStaticStageList struct {
	value *PaginatedAuthenticatorStaticStageList
	isSet bool
}

func (v NullablePaginatedAuthenticatorStaticStageList) Get() *PaginatedAuthenticatorStaticStageList {
	return v.value
}

func (v *NullablePaginatedAuthenticatorStaticStageList) Set(val *PaginatedAuthenticatorStaticStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAuthenticatorStaticStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAuthenticatorStaticStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAuthenticatorStaticStageList(val *PaginatedAuthenticatorStaticStageList) *NullablePaginatedAuthenticatorStaticStageList {
	return &NullablePaginatedAuthenticatorStaticStageList{value: val, isSet: true}
}

func (v NullablePaginatedAuthenticatorStaticStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAuthenticatorStaticStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


