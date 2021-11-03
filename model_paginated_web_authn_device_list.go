/*
authentik

Making authentication simple.

API version: 2021.10.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedWebAuthnDeviceList struct for PaginatedWebAuthnDeviceList
type PaginatedWebAuthnDeviceList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []WebAuthnDevice                   `json:"results"`
}

// NewPaginatedWebAuthnDeviceList instantiates a new PaginatedWebAuthnDeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedWebAuthnDeviceList(pagination PaginatedApplicationListPagination, results []WebAuthnDevice) *PaginatedWebAuthnDeviceList {
	this := PaginatedWebAuthnDeviceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedWebAuthnDeviceListWithDefaults instantiates a new PaginatedWebAuthnDeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedWebAuthnDeviceListWithDefaults() *PaginatedWebAuthnDeviceList {
	this := PaginatedWebAuthnDeviceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedWebAuthnDeviceList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedWebAuthnDeviceList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedWebAuthnDeviceList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedWebAuthnDeviceList) GetResults() []WebAuthnDevice {
	if o == nil {
		var ret []WebAuthnDevice
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedWebAuthnDeviceList) GetResultsOk() (*[]WebAuthnDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedWebAuthnDeviceList) SetResults(v []WebAuthnDevice) {
	o.Results = v
}

func (o PaginatedWebAuthnDeviceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedWebAuthnDeviceList struct {
	value *PaginatedWebAuthnDeviceList
	isSet bool
}

func (v NullablePaginatedWebAuthnDeviceList) Get() *PaginatedWebAuthnDeviceList {
	return v.value
}

func (v *NullablePaginatedWebAuthnDeviceList) Set(val *PaginatedWebAuthnDeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedWebAuthnDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedWebAuthnDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedWebAuthnDeviceList(val *PaginatedWebAuthnDeviceList) *NullablePaginatedWebAuthnDeviceList {
	return &NullablePaginatedWebAuthnDeviceList{value: val, isSet: true}
}

func (v NullablePaginatedWebAuthnDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedWebAuthnDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
