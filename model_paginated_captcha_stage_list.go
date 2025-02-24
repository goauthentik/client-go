/*
authentik

Making authentication simple.

API version: 2025.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaginatedCaptchaStageList struct for PaginatedCaptchaStageList
type PaginatedCaptchaStageList struct {
	Pagination Pagination     `json:"pagination"`
	Results    []CaptchaStage `json:"results"`
}

// NewPaginatedCaptchaStageList instantiates a new PaginatedCaptchaStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCaptchaStageList(pagination Pagination, results []CaptchaStage) *PaginatedCaptchaStageList {
	this := PaginatedCaptchaStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedCaptchaStageListWithDefaults instantiates a new PaginatedCaptchaStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCaptchaStageListWithDefaults() *PaginatedCaptchaStageList {
	this := PaginatedCaptchaStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedCaptchaStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedCaptchaStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedCaptchaStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedCaptchaStageList) GetResults() []CaptchaStage {
	if o == nil {
		var ret []CaptchaStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedCaptchaStageList) GetResultsOk() ([]CaptchaStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedCaptchaStageList) SetResults(v []CaptchaStage) {
	o.Results = v
}

func (o PaginatedCaptchaStageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedCaptchaStageList struct {
	value *PaginatedCaptchaStageList
	isSet bool
}

func (v NullablePaginatedCaptchaStageList) Get() *PaginatedCaptchaStageList {
	return v.value
}

func (v *NullablePaginatedCaptchaStageList) Set(val *PaginatedCaptchaStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCaptchaStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCaptchaStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCaptchaStageList(val *PaginatedCaptchaStageList) *NullablePaginatedCaptchaStageList {
	return &NullablePaginatedCaptchaStageList{value: val, isSet: true}
}

func (v NullablePaginatedCaptchaStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCaptchaStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
