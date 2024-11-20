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

// checks if the PaginatedUserLoginStageList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedUserLoginStageList{}

// PaginatedUserLoginStageList struct for PaginatedUserLoginStageList
type PaginatedUserLoginStageList struct {
	Pagination Pagination       `json:"pagination"`
	Results    []UserLoginStage `json:"results"`
}

type _PaginatedUserLoginStageList PaginatedUserLoginStageList

// NewPaginatedUserLoginStageList instantiates a new PaginatedUserLoginStageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserLoginStageList(pagination Pagination, results []UserLoginStage) *PaginatedUserLoginStageList {
	this := PaginatedUserLoginStageList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedUserLoginStageListWithDefaults instantiates a new PaginatedUserLoginStageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserLoginStageListWithDefaults() *PaginatedUserLoginStageList {
	this := PaginatedUserLoginStageList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedUserLoginStageList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserLoginStageList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedUserLoginStageList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedUserLoginStageList) GetResults() []UserLoginStage {
	if o == nil {
		var ret []UserLoginStage
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedUserLoginStageList) GetResultsOk() ([]UserLoginStage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedUserLoginStageList) SetResults(v []UserLoginStage) {
	o.Results = v
}

func (o PaginatedUserLoginStageList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedUserLoginStageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PaginatedUserLoginStageList) UnmarshalJSON(data []byte) (err error) {
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

	varPaginatedUserLoginStageList := _PaginatedUserLoginStageList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginatedUserLoginStageList)

	if err != nil {
		return err
	}

	*o = PaginatedUserLoginStageList(varPaginatedUserLoginStageList)

	return err
}

type NullablePaginatedUserLoginStageList struct {
	value *PaginatedUserLoginStageList
	isSet bool
}

func (v NullablePaginatedUserLoginStageList) Get() *PaginatedUserLoginStageList {
	return v.value
}

func (v *NullablePaginatedUserLoginStageList) Set(val *PaginatedUserLoginStageList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserLoginStageList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserLoginStageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserLoginStageList(val *PaginatedUserLoginStageList) *NullablePaginatedUserLoginStageList {
	return &NullablePaginatedUserLoginStageList{value: val, isSet: true}
}

func (v NullablePaginatedUserLoginStageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserLoginStageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
