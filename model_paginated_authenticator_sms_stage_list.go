/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedAuthenticatorSMSStageList struct for PaginatedAuthenticatorSMSStageList
type PaginatedAuthenticatorSMSStageList struct {
	Pagination Pagination              `json:"pagination"`
	Results    []AuthenticatorSMSStage `json:"results"`
}

// NewPaginatedAuthenticatorSMSStageList instantiates a new PaginatedAuthenticatorSMSStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAuthenticatorSMSStageList(pagination Pagination, results []AuthenticatorSMSStage) *PaginatedAuthenticatorSMSStageList {
	this := PaginatedAuthenticatorSMSStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedAuthenticatorSMSStageListWithDefaults instantiates a new PaginatedAuthenticatorSMSStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAuthenticatorSMSStageListWithDefaults() *PaginatedAuthenticatorSMSStageList {
	this := PaginatedAuthenticatorSMSStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedAuthenticatorSMSStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorSMSStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedAuthenticatorSMSStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedAuthenticatorSMSStageList) GetResults() []AuthenticatorSMSStage {
	if o == nil {
		var ret []AuthenticatorSMSStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorSMSStageList) GetResultsOk() ([]AuthenticatorSMSStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAuthenticatorSMSStageList) SetResults(v []AuthenticatorSMSStage) {
	o.Results = v
}

func (o PaginatedAuthenticatorSMSStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedAuthenticatorSMSStageList struct {
	value *PaginatedAuthenticatorSMSStageList
	isSet bool
}

func (v NullablePaginatedAuthenticatorSMSStageList) Get() *PaginatedAuthenticatorSMSStageList {
	return v.value
}

func (v *NullablePaginatedAuthenticatorSMSStageList) Set(val *PaginatedAuthenticatorSMSStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAuthenticatorSMSStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAuthenticatorSMSStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAuthenticatorSMSStageList(val *PaginatedAuthenticatorSMSStageList) *NullablePaginatedAuthenticatorSMSStageList {
	return &NullablePaginatedAuthenticatorSMSStageList{value: val, isSet: true}
}

func (v NullablePaginatedAuthenticatorSMSStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAuthenticatorSMSStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
