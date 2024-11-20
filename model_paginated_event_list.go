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

// checks if the PaginatedEventList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedEventList{}

// PaginatedEventList struct for PaginatedEventList
type PaginatedEventList struct {
	Pagination Pagination `json:"pagination"`
	Results    []Event    `json:"results"`
}

type _PaginatedEventList PaginatedEventList

// NewPaginatedEventList instantiates a new PaginatedEventList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEventList(pagination Pagination, results []Event) *PaginatedEventList {
	this := PaginatedEventList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedEventListWithDefaults instantiates a new PaginatedEventList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEventListWithDefaults() *PaginatedEventList {
	this := PaginatedEventList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedEventList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedEventList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedEventList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedEventList) GetResults() []Event {
	if o == nil {
		var ret []Event
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedEventList) GetResultsOk() ([]Event, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedEventList) SetResults(v []Event) {
	o.Results = v
}

func (o PaginatedEventList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedEventList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PaginatedEventList) UnmarshalJSON(data []byte) (err error) {
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

	varPaginatedEventList := _PaginatedEventList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginatedEventList)

	if err != nil {
		return err
	}

	*o = PaginatedEventList(varPaginatedEventList)

	return err
}

type NullablePaginatedEventList struct {
	value *PaginatedEventList
	isSet bool
}

func (v NullablePaginatedEventList) Get() *PaginatedEventList {
	return v.value
}

func (v *NullablePaginatedEventList) Set(val *PaginatedEventList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEventList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEventList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEventList(val *PaginatedEventList) *NullablePaginatedEventList {
	return &NullablePaginatedEventList{value: val, isSet: true}
}

func (v NullablePaginatedEventList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEventList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
