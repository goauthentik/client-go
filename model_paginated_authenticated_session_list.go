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

// checks if the PaginatedAuthenticatedSessionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedAuthenticatedSessionList{}

// PaginatedAuthenticatedSessionList struct for PaginatedAuthenticatedSessionList
type PaginatedAuthenticatedSessionList struct {
	Pagination Pagination             `json:"pagination"`
	Results    []AuthenticatedSession `json:"results"`
}

type _PaginatedAuthenticatedSessionList PaginatedAuthenticatedSessionList

// NewPaginatedAuthenticatedSessionList instantiates a new PaginatedAuthenticatedSessionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedAuthenticatedSessionList(pagination Pagination, results []AuthenticatedSession) *PaginatedAuthenticatedSessionList {
	this := PaginatedAuthenticatedSessionList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedAuthenticatedSessionListWithDefaults instantiates a new PaginatedAuthenticatedSessionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedAuthenticatedSessionListWithDefaults() *PaginatedAuthenticatedSessionList {
	this := PaginatedAuthenticatedSessionList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedAuthenticatedSessionList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatedSessionList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedAuthenticatedSessionList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedAuthenticatedSessionList) GetResults() []AuthenticatedSession {
	if o == nil {
		var ret []AuthenticatedSession
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedAuthenticatedSessionList) GetResultsOk() ([]AuthenticatedSession, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedAuthenticatedSessionList) SetResults(v []AuthenticatedSession) {
	o.Results = v
}

func (o PaginatedAuthenticatedSessionList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedAuthenticatedSessionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PaginatedAuthenticatedSessionList) UnmarshalJSON(data []byte) (err error) {
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

	varPaginatedAuthenticatedSessionList := _PaginatedAuthenticatedSessionList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginatedAuthenticatedSessionList)

	if err != nil {
		return err
	}

	*o = PaginatedAuthenticatedSessionList(varPaginatedAuthenticatedSessionList)

	return err
}

type NullablePaginatedAuthenticatedSessionList struct {
	value *PaginatedAuthenticatedSessionList
	isSet bool
}

func (v NullablePaginatedAuthenticatedSessionList) Get() *PaginatedAuthenticatedSessionList {
	return v.value
}

func (v *NullablePaginatedAuthenticatedSessionList) Set(val *PaginatedAuthenticatedSessionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedAuthenticatedSessionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedAuthenticatedSessionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedAuthenticatedSessionList(val *PaginatedAuthenticatedSessionList) *NullablePaginatedAuthenticatedSessionList {
	return &NullablePaginatedAuthenticatedSessionList{value: val, isSet: true}
}

func (v NullablePaginatedAuthenticatedSessionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedAuthenticatedSessionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
