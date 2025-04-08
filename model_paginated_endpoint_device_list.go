/*
authentik

Making authentication simple.

API version: 2025.2.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedEndpointDeviceList struct for PaginatedEndpointDeviceList
type PaginatedEndpointDeviceList struct {
	Pagination Pagination       `json:"pagination"`
	Results    []EndpointDevice `json:"results"`
}

// NewPaginatedEndpointDeviceList instantiates a new PaginatedEndpointDeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEndpointDeviceList(pagination Pagination, results []EndpointDevice) *PaginatedEndpointDeviceList {
	this := PaginatedEndpointDeviceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedEndpointDeviceListWithDefaults instantiates a new PaginatedEndpointDeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEndpointDeviceListWithDefaults() *PaginatedEndpointDeviceList {
	this := PaginatedEndpointDeviceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedEndpointDeviceList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedEndpointDeviceList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedEndpointDeviceList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedEndpointDeviceList) GetResults() []EndpointDevice {
	if o == nil {
		var ret []EndpointDevice
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedEndpointDeviceList) GetResultsOk() ([]EndpointDevice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedEndpointDeviceList) SetResults(v []EndpointDevice) {
	o.Results = v
}

func (o PaginatedEndpointDeviceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedEndpointDeviceList struct {
	value *PaginatedEndpointDeviceList
	isSet bool
}

func (v NullablePaginatedEndpointDeviceList) Get() *PaginatedEndpointDeviceList {
	return v.value
}

func (v *NullablePaginatedEndpointDeviceList) Set(val *PaginatedEndpointDeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEndpointDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEndpointDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEndpointDeviceList(val *PaginatedEndpointDeviceList) *NullablePaginatedEndpointDeviceList {
	return &NullablePaginatedEndpointDeviceList{value: val, isSet: true}
}

func (v NullablePaginatedEndpointDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEndpointDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
