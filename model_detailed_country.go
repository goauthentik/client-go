/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DetailedCountry struct for DetailedCountry
type DetailedCountry struct {
	Code CountryCodeEnum `json:"code"`
	Name string          `json:"name"`
}

// NewDetailedCountry instantiates a new DetailedCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedCountry(code CountryCodeEnum, name string) *DetailedCountry {
	this := DetailedCountry{}
	this.Code = code
	this.Name = name
	return &this
}

// NewDetailedCountryWithDefaults instantiates a new DetailedCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedCountryWithDefaults() *DetailedCountry {
	this := DetailedCountry{}
	return &this
}

// GetCode returns the Code field value
func (o *DetailedCountry) GetCode() CountryCodeEnum {
	if o == nil {
		var ret CountryCodeEnum
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DetailedCountry) GetCodeOk() (*CountryCodeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DetailedCountry) SetCode(v CountryCodeEnum) {
	o.Code = v
}

// GetName returns the Name field value
func (o *DetailedCountry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DetailedCountry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DetailedCountry) SetName(v string) {
	o.Name = v
}

func (o DetailedCountry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDetailedCountry struct {
	value *DetailedCountry
	isSet bool
}

func (v NullableDetailedCountry) Get() *DetailedCountry {
	return v.value
}

func (v *NullableDetailedCountry) Set(val *DetailedCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedCountry(val *DetailedCountry) *NullableDetailedCountry {
	return &NullableDetailedCountry{value: val, isSet: true}
}

func (v NullableDetailedCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
