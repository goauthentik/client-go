/*
authentik

Making authentication simple.

API version: 2022.11.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedKubernetesServiceConnectionList struct for PaginatedKubernetesServiceConnectionList
type PaginatedKubernetesServiceConnectionList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []KubernetesServiceConnection      `json:"results"`
}

// NewPaginatedKubernetesServiceConnectionList instantiates a new PaginatedKubernetesServiceConnectionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedKubernetesServiceConnectionList(pagination PaginatedApplicationListPagination, results []KubernetesServiceConnection) *PaginatedKubernetesServiceConnectionList {
	this := PaginatedKubernetesServiceConnectionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedKubernetesServiceConnectionListWithDefaults instantiates a new PaginatedKubernetesServiceConnectionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedKubernetesServiceConnectionListWithDefaults() *PaginatedKubernetesServiceConnectionList {
	this := PaginatedKubernetesServiceConnectionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedKubernetesServiceConnectionList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedKubernetesServiceConnectionList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedKubernetesServiceConnectionList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedKubernetesServiceConnectionList) GetResults() []KubernetesServiceConnection {
	if o == nil {
		var ret []KubernetesServiceConnection
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedKubernetesServiceConnectionList) GetResultsOk() ([]KubernetesServiceConnection, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedKubernetesServiceConnectionList) SetResults(v []KubernetesServiceConnection) {
	o.Results = v
}

func (o PaginatedKubernetesServiceConnectionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedKubernetesServiceConnectionList struct {
	value *PaginatedKubernetesServiceConnectionList
	isSet bool
}

func (v NullablePaginatedKubernetesServiceConnectionList) Get() *PaginatedKubernetesServiceConnectionList {
	return v.value
}

func (v *NullablePaginatedKubernetesServiceConnectionList) Set(val *PaginatedKubernetesServiceConnectionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedKubernetesServiceConnectionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedKubernetesServiceConnectionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedKubernetesServiceConnectionList(val *PaginatedKubernetesServiceConnectionList) *NullablePaginatedKubernetesServiceConnectionList {
	return &NullablePaginatedKubernetesServiceConnectionList{value: val, isSet: true}
}

func (v NullablePaginatedKubernetesServiceConnectionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedKubernetesServiceConnectionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
