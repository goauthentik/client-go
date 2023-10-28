/*
authentik

Making authentication simple.

API version: 2023.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedNotificationList struct for PaginatedNotificationList
type PaginatedNotificationList struct {
	Pagination Pagination     `json:"pagination"`
	Results    []Notification `json:"results"`
}

// NewPaginatedNotificationList instantiates a new PaginatedNotificationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedNotificationList(pagination Pagination, results []Notification) *PaginatedNotificationList {
	this := PaginatedNotificationList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedNotificationListWithDefaults instantiates a new PaginatedNotificationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedNotificationListWithDefaults() *PaginatedNotificationList {
	this := PaginatedNotificationList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedNotificationList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedNotificationList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedNotificationList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedNotificationList) GetResults() []Notification {
	if o == nil {
		var ret []Notification
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedNotificationList) GetResultsOk() ([]Notification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedNotificationList) SetResults(v []Notification) {
	o.Results = v
}

func (o PaginatedNotificationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedNotificationList struct {
	value *PaginatedNotificationList
	isSet bool
}

func (v NullablePaginatedNotificationList) Get() *PaginatedNotificationList {
	return v.value
}

func (v *NullablePaginatedNotificationList) Set(val *PaginatedNotificationList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedNotificationList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedNotificationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedNotificationList(val *PaginatedNotificationList) *NullablePaginatedNotificationList {
	return &NullablePaginatedNotificationList{value: val, isSet: true}
}

func (v NullablePaginatedNotificationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedNotificationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
