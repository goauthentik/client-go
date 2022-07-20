/*
authentik

Making authentication simple.

API version: 2022.7.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedTOTPDeviceList struct for PaginatedTOTPDeviceList
type PaginatedTOTPDeviceList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []TOTPDevice                       `json:"results"`
}

// NewPaginatedTOTPDeviceList instantiates a new PaginatedTOTPDeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedTOTPDeviceList(pagination PaginatedApplicationListPagination, results []TOTPDevice) *PaginatedTOTPDeviceList {
	this := PaginatedTOTPDeviceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedTOTPDeviceListWithDefaults instantiates a new PaginatedTOTPDeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedTOTPDeviceListWithDefaults() *PaginatedTOTPDeviceList {
	this := PaginatedTOTPDeviceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedTOTPDeviceList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedTOTPDeviceList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedTOTPDeviceList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedTOTPDeviceList) GetResults() []TOTPDevice {
	if o == nil {
		var ret []TOTPDevice
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedTOTPDeviceList) GetResultsOk() ([]TOTPDevice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedTOTPDeviceList) SetResults(v []TOTPDevice) {
	o.Results = v
}

func (o PaginatedTOTPDeviceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedTOTPDeviceList struct {
	value *PaginatedTOTPDeviceList
	isSet bool
}

func (v NullablePaginatedTOTPDeviceList) Get() *PaginatedTOTPDeviceList {
	return v.value
}

func (v *NullablePaginatedTOTPDeviceList) Set(val *PaginatedTOTPDeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedTOTPDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedTOTPDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedTOTPDeviceList(val *PaginatedTOTPDeviceList) *NullablePaginatedTOTPDeviceList {
	return &NullablePaginatedTOTPDeviceList{value: val, isSet: true}
}

func (v NullablePaginatedTOTPDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedTOTPDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
