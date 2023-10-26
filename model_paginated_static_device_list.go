/*
authentik

Making authentication simple.

API version: 2023.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedStaticDeviceList struct for PaginatedStaticDeviceList
type PaginatedStaticDeviceList struct {
	Pagination Pagination     `json:"pagination"`
	Results    []StaticDevice `json:"results"`
}

// NewPaginatedStaticDeviceList instantiates a new PaginatedStaticDeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedStaticDeviceList(pagination Pagination, results []StaticDevice) *PaginatedStaticDeviceList {
	this := PaginatedStaticDeviceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedStaticDeviceListWithDefaults instantiates a new PaginatedStaticDeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedStaticDeviceListWithDefaults() *PaginatedStaticDeviceList {
	this := PaginatedStaticDeviceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedStaticDeviceList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedStaticDeviceList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedStaticDeviceList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedStaticDeviceList) GetResults() []StaticDevice {
	if o == nil {
		var ret []StaticDevice
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedStaticDeviceList) GetResultsOk() ([]StaticDevice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedStaticDeviceList) SetResults(v []StaticDevice) {
	o.Results = v
}

func (o PaginatedStaticDeviceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedStaticDeviceList struct {
	value *PaginatedStaticDeviceList
	isSet bool
}

func (v NullablePaginatedStaticDeviceList) Get() *PaginatedStaticDeviceList {
	return v.value
}

func (v *NullablePaginatedStaticDeviceList) Set(val *PaginatedStaticDeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedStaticDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedStaticDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedStaticDeviceList(val *PaginatedStaticDeviceList) *NullablePaginatedStaticDeviceList {
	return &NullablePaginatedStaticDeviceList{value: val, isSet: true}
}

func (v NullablePaginatedStaticDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedStaticDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
