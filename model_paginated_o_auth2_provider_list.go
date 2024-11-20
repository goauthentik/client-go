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

// checks if the PaginatedOAuth2ProviderList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedOAuth2ProviderList{}

// PaginatedOAuth2ProviderList struct for PaginatedOAuth2ProviderList
type PaginatedOAuth2ProviderList struct {
	Pagination Pagination       `json:"pagination"`
	Results    []OAuth2Provider `json:"results"`
}

type _PaginatedOAuth2ProviderList PaginatedOAuth2ProviderList

// NewPaginatedOAuth2ProviderList instantiates a new PaginatedOAuth2ProviderList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedOAuth2ProviderList(pagination Pagination, results []OAuth2Provider) *PaginatedOAuth2ProviderList {
	this := PaginatedOAuth2ProviderList{}
	this.Pagination = pagination
	this.Results = results
	return &this
}

// NewPaginatedOAuth2ProviderListWithDefaults instantiates a new PaginatedOAuth2ProviderList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedOAuth2ProviderListWithDefaults() *PaginatedOAuth2ProviderList {
	this := PaginatedOAuth2ProviderList{}
	return &this
}

// GetPagination returns the Pagination field value
func (o *PaginatedOAuth2ProviderList) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *PaginatedOAuth2ProviderList) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *PaginatedOAuth2ProviderList) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetResults returns the Results field value
func (o *PaginatedOAuth2ProviderList) GetResults() []OAuth2Provider {
	if o == nil {
		var ret []OAuth2Provider
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedOAuth2ProviderList) GetResultsOk() ([]OAuth2Provider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedOAuth2ProviderList) SetResults(v []OAuth2Provider) {
	o.Results = v
}

func (o PaginatedOAuth2ProviderList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedOAuth2ProviderList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pagination"] = o.Pagination
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *PaginatedOAuth2ProviderList) UnmarshalJSON(data []byte) (err error) {
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

	varPaginatedOAuth2ProviderList := _PaginatedOAuth2ProviderList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginatedOAuth2ProviderList)

	if err != nil {
		return err
	}

	*o = PaginatedOAuth2ProviderList(varPaginatedOAuth2ProviderList)

	return err
}

type NullablePaginatedOAuth2ProviderList struct {
	value *PaginatedOAuth2ProviderList
	isSet bool
}

func (v NullablePaginatedOAuth2ProviderList) Get() *PaginatedOAuth2ProviderList {
	return v.value
}

func (v *NullablePaginatedOAuth2ProviderList) Set(val *PaginatedOAuth2ProviderList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedOAuth2ProviderList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedOAuth2ProviderList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedOAuth2ProviderList(val *PaginatedOAuth2ProviderList) *NullablePaginatedOAuth2ProviderList {
	return &NullablePaginatedOAuth2ProviderList{value: val, isSet: true}
}

func (v NullablePaginatedOAuth2ProviderList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedOAuth2ProviderList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
