/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PaginatedReputationList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedReputationList{}

// PaginatedReputationList struct for PaginatedReputationList
type PaginatedReputationList struct {
	Pagination Pagination   `json:"pagination"`
	Results    []Reputation `json:"results"`
}

type _PaginatedReputationList PaginatedReputationList

// NewPaginatedReputationList instantiates a new PaginatedReputationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedReputationList(pagination Pagination, results []Reputation) *PaginatedReputationList {
	this := PaginatedReputationList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedReputationListWithDefaults instantiates a new PaginatedReputationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedReputationListWithDefaults() *PaginatedReputationList {
	this := PaginatedReputationList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedReputationList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedReputationList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedReputationList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedReputationList) GetResults() []Reputation {
	if o == nil {
		var ret []Reputation
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedReputationList) GetResultsOk() ([]Reputation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedReputationList) SetResults(v []Reputation) {
	o.Results = v
}

func (o PaginatedReputationList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedReputationList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PaginatedReputationList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pagination",
		"results",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPaginatedReputationList := _PaginatedReputationList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginatedReputationList)

	if err != nil {
		return err
	}

	*o = PaginatedReputationList(varPaginatedReputationList)

	return err
}

type NullablePaginatedReputationList struct {
	value *PaginatedReputationList
	isSet bool
}

func (v NullablePaginatedReputationList) Get() *PaginatedReputationList {
	return v.value
}

func (v *NullablePaginatedReputationList) Set(val *PaginatedReputationList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedReputationList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedReputationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedReputationList(val *PaginatedReputationList) *NullablePaginatedReputationList {
	return &NullablePaginatedReputationList{value: val, isSet: true}
}

func (v NullablePaginatedReputationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedReputationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
