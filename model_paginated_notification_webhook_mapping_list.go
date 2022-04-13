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

// PaginatedNotificationWebhookMappingList struct for PaginatedNotificationWebhookMappingList
type PaginatedNotificationWebhookMappingList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []NotificationWebhookMapping       `json:"results"`
}

// NewPaginatedNotificationWebhookMappingList instantiates a new PaginatedNotificationWebhookMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedNotificationWebhookMappingList(pagination PaginatedApplicationListPagination, results []NotificationWebhookMapping) *PaginatedNotificationWebhookMappingList {
	this := PaginatedNotificationWebhookMappingList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedNotificationWebhookMappingListWithDefaults instantiates a new PaginatedNotificationWebhookMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedNotificationWebhookMappingListWithDefaults() *PaginatedNotificationWebhookMappingList {
	this := PaginatedNotificationWebhookMappingList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedNotificationWebhookMappingList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedNotificationWebhookMappingList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedNotificationWebhookMappingList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedNotificationWebhookMappingList) GetResults() []NotificationWebhookMapping {
	if o == nil {
		var ret []NotificationWebhookMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedNotificationWebhookMappingList) GetResultsOk() (*[]NotificationWebhookMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedNotificationWebhookMappingList) SetResults(v []NotificationWebhookMapping) {
	o.Results = v
}

func (o PaginatedNotificationWebhookMappingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedNotificationWebhookMappingList struct {
	value *PaginatedNotificationWebhookMappingList
	isSet bool
}

func (v NullablePaginatedNotificationWebhookMappingList) Get() *PaginatedNotificationWebhookMappingList {
	return v.value
}

func (v *NullablePaginatedNotificationWebhookMappingList) Set(val *PaginatedNotificationWebhookMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedNotificationWebhookMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedNotificationWebhookMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedNotificationWebhookMappingList(val *PaginatedNotificationWebhookMappingList) *NullablePaginatedNotificationWebhookMappingList {
	return &NullablePaginatedNotificationWebhookMappingList{value: val, isSet: true}
}

func (v NullablePaginatedNotificationWebhookMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedNotificationWebhookMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
