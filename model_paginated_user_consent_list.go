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

// checks if the PaginatedUserConsentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedUserConsentList{}

// PaginatedUserConsentList struct for PaginatedUserConsentList
type PaginatedUserConsentList struct {
	Pagination Pagination    `json:"pagination"`
	Results    []UserConsent `json:"results"`
}

type _PaginatedUserConsentList PaginatedUserConsentList

// NewPaginatedUserConsentList instantiates a new PaginatedUserConsentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserConsentList(pagination Pagination, results []UserConsent) *PaginatedUserConsentList {
	this := PaginatedUserConsentList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserConsentListWithDefaults instantiates a new PaginatedUserConsentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserConsentListWithDefaults() *PaginatedUserConsentList {
	this := PaginatedUserConsentList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserConsentList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserConsentList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserConsentList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserConsentList) GetResults() []UserConsent {
	if o == nil {
		var ret []UserConsent
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserConsentList) GetResultsOk() ([]UserConsent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserConsentList) SetResults(v []UserConsent) {
	o.Results = v
}

func (o PaginatedUserConsentList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedUserConsentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PaginatedUserConsentList) UnmarshalJSON(data []byte) (err error) {
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

	varPaginatedUserConsentList := _PaginatedUserConsentList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginatedUserConsentList)

	if err != nil {
		return err
	}

	*o = PaginatedUserConsentList(varPaginatedUserConsentList)

	return err
}

type NullablePaginatedUserConsentList struct {
	value *PaginatedUserConsentList
	isSet bool
}

func (v NullablePaginatedUserConsentList) Get() *PaginatedUserConsentList {
	return v.value
}

func (v *NullablePaginatedUserConsentList) Set(val *PaginatedUserConsentList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserConsentList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserConsentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserConsentList(val *PaginatedUserConsentList) *NullablePaginatedUserConsentList {
	return &NullablePaginatedUserConsentList{value: val, isSet: true}
}

func (v NullablePaginatedUserConsentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserConsentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
