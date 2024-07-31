/*
authentik

Making authentication simple.

API version: 2024.6.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedWebAuthnDeviceTypeList struct for PaginatedWebAuthnDeviceTypeList
type PaginatedWebAuthnDeviceTypeList struct {
	Pagination Pagination           `json:"pagination"`
	Results    []WebAuthnDeviceType `json:"results"`
}

// NewPaginatedWebAuthnDeviceTypeList instantiates a new PaginatedWebAuthnDeviceTypeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedWebAuthnDeviceTypeList(pagination Pagination, results []WebAuthnDeviceType) *PaginatedWebAuthnDeviceTypeList {
	this := PaginatedWebAuthnDeviceTypeList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedWebAuthnDeviceTypeListWithDefaults instantiates a new PaginatedWebAuthnDeviceTypeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedWebAuthnDeviceTypeListWithDefaults() *PaginatedWebAuthnDeviceTypeList {
	this := PaginatedWebAuthnDeviceTypeList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedWebAuthnDeviceTypeList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedWebAuthnDeviceTypeList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedWebAuthnDeviceTypeList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedWebAuthnDeviceTypeList) GetResults() []WebAuthnDeviceType {
	if o == nil {
		var ret []WebAuthnDeviceType
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedWebAuthnDeviceTypeList) GetResultsOk() ([]WebAuthnDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedWebAuthnDeviceTypeList) SetResults(v []WebAuthnDeviceType) {
	o.Results = v
}

func (o PaginatedWebAuthnDeviceTypeList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedWebAuthnDeviceTypeList struct {
	value *PaginatedWebAuthnDeviceTypeList
	isSet bool
}

func (v NullablePaginatedWebAuthnDeviceTypeList) Get() *PaginatedWebAuthnDeviceTypeList {
	return v.value
}

func (v *NullablePaginatedWebAuthnDeviceTypeList) Set(val *PaginatedWebAuthnDeviceTypeList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedWebAuthnDeviceTypeList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedWebAuthnDeviceTypeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedWebAuthnDeviceTypeList(val *PaginatedWebAuthnDeviceTypeList) *NullablePaginatedWebAuthnDeviceTypeList {
	return &NullablePaginatedWebAuthnDeviceTypeList{value: val, isSet: true}
}

func (v NullablePaginatedWebAuthnDeviceTypeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedWebAuthnDeviceTypeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
