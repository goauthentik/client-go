/*
authentik

Making authentication simple.

API version: 2022.4.1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedDuoDeviceList struct for PaginatedDuoDeviceList
type PaginatedDuoDeviceList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []DuoDevice                        `json:"results"`
}

// NewPaginatedDuoDeviceList instantiates a new PaginatedDuoDeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDuoDeviceList(pagination PaginatedApplicationListPagination, results []DuoDevice) *PaginatedDuoDeviceList {
	this := PaginatedDuoDeviceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedDuoDeviceListWithDefaults instantiates a new PaginatedDuoDeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDuoDeviceListWithDefaults() *PaginatedDuoDeviceList {
	this := PaginatedDuoDeviceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedDuoDeviceList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedDuoDeviceList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedDuoDeviceList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedDuoDeviceList) GetResults() []DuoDevice {
	if o == nil {
		var ret []DuoDevice
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedDuoDeviceList) GetResultsOk() (*[]DuoDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedDuoDeviceList) SetResults(v []DuoDevice) {
	o.Results = v
}

func (o PaginatedDuoDeviceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedDuoDeviceList struct {
	value *PaginatedDuoDeviceList
	isSet bool
}

func (v NullablePaginatedDuoDeviceList) Get() *PaginatedDuoDeviceList {
	return v.value
}

func (v *NullablePaginatedDuoDeviceList) Set(val *PaginatedDuoDeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedDuoDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedDuoDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedDuoDeviceList(val *PaginatedDuoDeviceList) *NullablePaginatedDuoDeviceList {
	return &NullablePaginatedDuoDeviceList{value: val, isSet: true}
}

func (v NullablePaginatedDuoDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedDuoDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
