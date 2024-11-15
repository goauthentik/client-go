/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedAuthenticatorWebAuthnStageList struct for PaginatedAuthenticatorWebAuthnStageList
type PaginatedAuthenticatorWebAuthnStageList struct {
	Pagination Pagination                   `json:"pagination"`
	Results    []AuthenticatorWebAuthnStage `json:"results"`
}

// NewPaginatedAuthenticatorWebAuthnStageList instantiates a new PaginatedAuthenticatorWebAuthnStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAuthenticatorWebAuthnStageList(pagination Pagination, results []AuthenticatorWebAuthnStage) *PaginatedAuthenticatorWebAuthnStageList {
	this := PaginatedAuthenticatorWebAuthnStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedAuthenticatorWebAuthnStageListWithDefaults instantiates a new PaginatedAuthenticatorWebAuthnStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAuthenticatorWebAuthnStageListWithDefaults() *PaginatedAuthenticatorWebAuthnStageList {
	this := PaginatedAuthenticatorWebAuthnStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedAuthenticatorWebAuthnStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorWebAuthnStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedAuthenticatorWebAuthnStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedAuthenticatorWebAuthnStageList) GetResults() []AuthenticatorWebAuthnStage {
	if o == nil {
		var ret []AuthenticatorWebAuthnStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorWebAuthnStageList) GetResultsOk() ([]AuthenticatorWebAuthnStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAuthenticatorWebAuthnStageList) SetResults(v []AuthenticatorWebAuthnStage) {
	o.Results = v
}

func (o PaginatedAuthenticatorWebAuthnStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedAuthenticatorWebAuthnStageList struct {
	value *PaginatedAuthenticatorWebAuthnStageList
	isSet bool
}

func (v NullablePaginatedAuthenticatorWebAuthnStageList) Get() *PaginatedAuthenticatorWebAuthnStageList {
	return v.value
}

func (v *NullablePaginatedAuthenticatorWebAuthnStageList) Set(val *PaginatedAuthenticatorWebAuthnStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAuthenticatorWebAuthnStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAuthenticatorWebAuthnStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAuthenticatorWebAuthnStageList(val *PaginatedAuthenticatorWebAuthnStageList) *NullablePaginatedAuthenticatorWebAuthnStageList {
	return &NullablePaginatedAuthenticatorWebAuthnStageList{value: val, isSet: true}
}

func (v NullablePaginatedAuthenticatorWebAuthnStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAuthenticatorWebAuthnStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
