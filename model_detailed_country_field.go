/*
authentik

Making authentication simple.

API version: 2024.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DetailedCountryField struct for DetailedCountryField
type DetailedCountryField struct {
	Code CountryCodeEnum `json:"code"`
	Name string          `json:"name"`
}

// NewDetailedCountryField instantiates a new DetailedCountryField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedCountryField(code CountryCodeEnum, name string) *DetailedCountryField {
	this := DetailedCountryField{}
	this.Code = code
	this.Name = name
	return &this
}

// NewDetailedCountryFieldWithDefaults instantiates a new DetailedCountryField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedCountryFieldWithDefaults() *DetailedCountryField {
	this := DetailedCountryField{}
	return &this
}

// GetCode returns the Code field value
func (o *DetailedCountryField) GetCode() CountryCodeEnum {
	if o == nil {
		var ret CountryCodeEnum
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DetailedCountryField) GetCodeOk() (*CountryCodeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DetailedCountryField) SetCode(v CountryCodeEnum) {
	o.Code = v
}

// GetName returns the Name field value
func (o *DetailedCountryField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DetailedCountryField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DetailedCountryField) SetName(v string) {
	o.Name = v
}

func (o DetailedCountryField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDetailedCountryField struct {
	value *DetailedCountryField
	isSet bool
}

func (v NullableDetailedCountryField) Get() *DetailedCountryField {
	return v.value
}

func (v *NullableDetailedCountryField) Set(val *DetailedCountryField) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedCountryField) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedCountryField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedCountryField(val *DetailedCountryField) *NullableDetailedCountryField {
	return &NullableDetailedCountryField{value: val, isSet: true}
}

func (v NullableDetailedCountryField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedCountryField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
