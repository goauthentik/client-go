/*
authentik

Making authentication simple.

API version: 2022.11.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedAuthenticatorTOTPStageList struct for PaginatedAuthenticatorTOTPStageList
type PaginatedAuthenticatorTOTPStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []AuthenticatorTOTPStage           `json:"results"`
}

// NewPaginatedAuthenticatorTOTPStageList instantiates a new PaginatedAuthenticatorTOTPStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAuthenticatorTOTPStageList(pagination PaginatedApplicationListPagination, results []AuthenticatorTOTPStage) *PaginatedAuthenticatorTOTPStageList {
	this := PaginatedAuthenticatorTOTPStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedAuthenticatorTOTPStageListWithDefaults instantiates a new PaginatedAuthenticatorTOTPStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAuthenticatorTOTPStageListWithDefaults() *PaginatedAuthenticatorTOTPStageList {
	this := PaginatedAuthenticatorTOTPStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedAuthenticatorTOTPStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorTOTPStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedAuthenticatorTOTPStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedAuthenticatorTOTPStageList) GetResults() []AuthenticatorTOTPStage {
	if o == nil {
		var ret []AuthenticatorTOTPStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorTOTPStageList) GetResultsOk() ([]AuthenticatorTOTPStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAuthenticatorTOTPStageList) SetResults(v []AuthenticatorTOTPStage) {
	o.Results = v
}

func (o PaginatedAuthenticatorTOTPStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedAuthenticatorTOTPStageList struct {
	value *PaginatedAuthenticatorTOTPStageList
	isSet bool
}

func (v NullablePaginatedAuthenticatorTOTPStageList) Get() *PaginatedAuthenticatorTOTPStageList {
	return v.value
}

func (v *NullablePaginatedAuthenticatorTOTPStageList) Set(val *PaginatedAuthenticatorTOTPStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAuthenticatorTOTPStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAuthenticatorTOTPStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAuthenticatorTOTPStageList(val *PaginatedAuthenticatorTOTPStageList) *NullablePaginatedAuthenticatorTOTPStageList {
	return &NullablePaginatedAuthenticatorTOTPStageList{value: val, isSet: true}
}

func (v NullablePaginatedAuthenticatorTOTPStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAuthenticatorTOTPStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
