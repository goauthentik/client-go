/*
authentik

Making authentication simple.

API version: 2022.12.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedCertificateKeyPairList struct for PaginatedCertificateKeyPairList
type PaginatedCertificateKeyPairList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []CertificateKeyPair               `json:"results"`
}

// NewPaginatedCertificateKeyPairList instantiates a new PaginatedCertificateKeyPairList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCertificateKeyPairList(pagination PaginatedApplicationListPagination, results []CertificateKeyPair) *PaginatedCertificateKeyPairList {
	this := PaginatedCertificateKeyPairList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedCertificateKeyPairListWithDefaults instantiates a new PaginatedCertificateKeyPairList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCertificateKeyPairListWithDefaults() *PaginatedCertificateKeyPairList {
	this := PaginatedCertificateKeyPairList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedCertificateKeyPairList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedCertificateKeyPairList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedCertificateKeyPairList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedCertificateKeyPairList) GetResults() []CertificateKeyPair {
	if o == nil {
		var ret []CertificateKeyPair
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedCertificateKeyPairList) GetResultsOk() ([]CertificateKeyPair, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedCertificateKeyPairList) SetResults(v []CertificateKeyPair) {
	o.Results = v
}

func (o PaginatedCertificateKeyPairList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedCertificateKeyPairList struct {
	value *PaginatedCertificateKeyPairList
	isSet bool
}

func (v NullablePaginatedCertificateKeyPairList) Get() *PaginatedCertificateKeyPairList {
	return v.value
}

func (v *NullablePaginatedCertificateKeyPairList) Set(val *PaginatedCertificateKeyPairList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCertificateKeyPairList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCertificateKeyPairList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCertificateKeyPairList(val *PaginatedCertificateKeyPairList) *NullablePaginatedCertificateKeyPairList {
	return &NullablePaginatedCertificateKeyPairList{value: val, isSet: true}
}

func (v NullablePaginatedCertificateKeyPairList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCertificateKeyPairList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
