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

// checks if the PaginatedAuthenticatorDuoStageList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedAuthenticatorDuoStageList{}

// PaginatedAuthenticatorDuoStageList struct for PaginatedAuthenticatorDuoStageList
type PaginatedAuthenticatorDuoStageList struct {
	Pagination PaginatedApplicationListPagination `json:"pagination"`
	Results    []AuthenticatorDuoStage            `json:"results"`
}

// NewPaginatedAuthenticatorDuoStageList instantiates a new PaginatedAuthenticatorDuoStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAuthenticatorDuoStageList(pagination PaginatedApplicationListPagination, results []AuthenticatorDuoStage) *PaginatedAuthenticatorDuoStageList {
	this := PaginatedAuthenticatorDuoStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedAuthenticatorDuoStageListWithDefaults instantiates a new PaginatedAuthenticatorDuoStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAuthenticatorDuoStageListWithDefaults() *PaginatedAuthenticatorDuoStageList {
	this := PaginatedAuthenticatorDuoStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedAuthenticatorDuoStageList) GetPagination() PaginatedApplicationListPagination {
	if o == nil {
		var ret PaginatedApplicationListPagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorDuoStageList) GetPaginationOk() (*PaginatedApplicationListPagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedAuthenticatorDuoStageList) SetPagination(v PaginatedApplicationListPagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedAuthenticatorDuoStageList) GetResults() []AuthenticatorDuoStage {
	if o == nil {
		var ret []AuthenticatorDuoStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatorDuoStageList) GetResultsOk() ([]AuthenticatorDuoStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAuthenticatorDuoStageList) SetResults(v []AuthenticatorDuoStage) {
	o.Results = v
}

func (o PaginatedAuthenticatorDuoStageList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedAuthenticatorDuoStageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullablePaginatedAuthenticatorDuoStageList struct {
	value *PaginatedAuthenticatorDuoStageList
	isSet bool
}

func (v NullablePaginatedAuthenticatorDuoStageList) Get() *PaginatedAuthenticatorDuoStageList {
	return v.value
}

func (v *NullablePaginatedAuthenticatorDuoStageList) Set(val *PaginatedAuthenticatorDuoStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAuthenticatorDuoStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAuthenticatorDuoStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAuthenticatorDuoStageList(val *PaginatedAuthenticatorDuoStageList) *NullablePaginatedAuthenticatorDuoStageList {
	return &NullablePaginatedAuthenticatorDuoStageList{value: val, isSet: true}
}

func (v NullablePaginatedAuthenticatorDuoStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAuthenticatorDuoStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
