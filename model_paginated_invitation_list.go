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

// PaginatedInvitationList struct for PaginatedInvitationList
type PaginatedInvitationList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []Invitation                       `json:"results"`
}

// NewPaginatedInvitationList instantiates a new PaginatedInvitationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedInvitationList(pagination PaginatedApplicationListPagination, results []Invitation) *PaginatedInvitationList {
	this := PaginatedInvitationList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedInvitationListWithDefaults instantiates a new PaginatedInvitationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedInvitationListWithDefaults() *PaginatedInvitationList {
	this := PaginatedInvitationList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedInvitationList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedInvitationList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedInvitationList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedInvitationList) GetResults() []Invitation {
	if o == nil {
		var ret []Invitation
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedInvitationList) GetResultsOk() (*[]Invitation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *PaginatedInvitationList) SetResults(v []Invitation) {
	o.Results = v
}

func (o PaginatedInvitationList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedInvitationList struct {
	value *PaginatedInvitationList
	isSet bool
}

func (v NullablePaginatedInvitationList) Get() *PaginatedInvitationList {
	return v.value
}

func (v *NullablePaginatedInvitationList) Set(val *PaginatedInvitationList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedInvitationList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedInvitationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedInvitationList(val *PaginatedInvitationList) *NullablePaginatedInvitationList {
	return &NullablePaginatedInvitationList{value: val, isSet: true}
}

func (v NullablePaginatedInvitationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedInvitationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
