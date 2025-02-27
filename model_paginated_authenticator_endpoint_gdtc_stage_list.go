/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedAuthenticatorEndpointGDTCStageList struct for PaginatedAuthenticatorEndpointGDTCStageList
type PaginatedAuthenticatorEndpointGDTCStageList struct {
	Pagination Pagination                       `json:"pagination"`
	Results    []AuthenticatorEndpointGDTCStage `json:"results"`
}

// NewPaginatedAuthenticatorEndpointGDTCStageList instantiates a new PaginatedAuthenticatorEndpointGDTCStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAuthenticatorEndpointGDTCStageList(pagination Pagination, results []AuthenticatorEndpointGDTCStage) *PaginatedAuthenticatorEndpointGDTCStageList {
	this := PaginatedAuthenticatorEndpointGDTCStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedAuthenticatorEndpointGDTCStageListWithDefaults instantiates a new PaginatedAuthenticatorEndpointGDTCStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAuthenticatorEndpointGDTCStageListWithDefaults() *PaginatedAuthenticatorEndpointGDTCStageList {
	this := PaginatedAuthenticatorEndpointGDTCStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedAuthenticatorEndpointGDTCStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorEndpointGDTCStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedAuthenticatorEndpointGDTCStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedAuthenticatorEndpointGDTCStageList) GetResults() []AuthenticatorEndpointGDTCStage {
	if o == nil {
		var ret []AuthenticatorEndpointGDTCStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorEndpointGDTCStageList) GetResultsOk() ([]AuthenticatorEndpointGDTCStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAuthenticatorEndpointGDTCStageList) SetResults(v []AuthenticatorEndpointGDTCStage) {
	o.Results = v
}

func (o PaginatedAuthenticatorEndpointGDTCStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedAuthenticatorEndpointGDTCStageList struct {
	value *PaginatedAuthenticatorEndpointGDTCStageList
	isSet bool
}

func (v NullablePaginatedAuthenticatorEndpointGDTCStageList) Get() *PaginatedAuthenticatorEndpointGDTCStageList {
	return v.value
}

func (v *NullablePaginatedAuthenticatorEndpointGDTCStageList) Set(val *PaginatedAuthenticatorEndpointGDTCStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAuthenticatorEndpointGDTCStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAuthenticatorEndpointGDTCStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAuthenticatorEndpointGDTCStageList(val *PaginatedAuthenticatorEndpointGDTCStageList) *NullablePaginatedAuthenticatorEndpointGDTCStageList {
	return &NullablePaginatedAuthenticatorEndpointGDTCStageList{value: val, isSet: true}
}

func (v NullablePaginatedAuthenticatorEndpointGDTCStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAuthenticatorEndpointGDTCStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
