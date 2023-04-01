/*
authentik

Making authentication simple.

API version: 2023.3.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PaginatedSMSDeviceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedSMSDeviceList{}

// PaginatedSMSDeviceList struct for PaginatedSMSDeviceList
type PaginatedSMSDeviceList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []SMSDevice                        `json:"results"`
}

// NewPaginatedSMSDeviceList instantiates a new PaginatedSMSDeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedSMSDeviceList(pagination PaginatedApplicationListPagination, results []SMSDevice) *PaginatedSMSDeviceList {
	this := PaginatedSMSDeviceList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedSMSDeviceListWithDefaults instantiates a new PaginatedSMSDeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedSMSDeviceListWithDefaults() *PaginatedSMSDeviceList {
	this := PaginatedSMSDeviceList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedSMSDeviceList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedSMSDeviceList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedSMSDeviceList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedSMSDeviceList) GetResults() []SMSDevice {
	if o == nil {
		var ret []SMSDevice
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedSMSDeviceList) GetResultsOk() ([]SMSDevice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedSMSDeviceList) SetResults(v []SMSDevice) {
	o.Results = v
}

func (o PaginatedSMSDeviceList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedSMSDeviceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullablePaginatedSMSDeviceList struct {
	value *PaginatedSMSDeviceList
	isSet bool
}

func (v NullablePaginatedSMSDeviceList) Get() *PaginatedSMSDeviceList {
	return v.value
}

func (v *NullablePaginatedSMSDeviceList) Set(val *PaginatedSMSDeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedSMSDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedSMSDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedSMSDeviceList(val *PaginatedSMSDeviceList) *NullablePaginatedSMSDeviceList {
	return &NullablePaginatedSMSDeviceList{value: val, isSet: true}
}

func (v NullablePaginatedSMSDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedSMSDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
