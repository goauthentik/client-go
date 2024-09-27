/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DetailedCountryFieldRequest struct for DetailedCountryFieldRequest
type DetailedCountryFieldRequest struct {
	Code CountryCodeEnum `json:"code"`
	Name string          `json:"name"`
}

// NewDetailedCountryFieldRequest instantiates a new DetailedCountryFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedCountryFieldRequest(code CountryCodeEnum, name string) *DetailedCountryFieldRequest {
	this := DetailedCountryFieldRequest{}
	this.Code = code
	this.Name = name
	return &this
}

// NewDetailedCountryFieldRequestWithDefaults instantiates a new DetailedCountryFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedCountryFieldRequestWithDefaults() *DetailedCountryFieldRequest {
	this := DetailedCountryFieldRequest{}
	return &this
}

// GetCode returns the Code field value
func (o *DetailedCountryFieldRequest) GetCode() CountryCodeEnum {
	if o == nil {
		var ret CountryCodeEnum
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DetailedCountryFieldRequest) GetCodeOk() (*CountryCodeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DetailedCountryFieldRequest) SetCode(v CountryCodeEnum) {
	o.Code = v
}

// GetName returns the Name field value
func (o *DetailedCountryFieldRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DetailedCountryFieldRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DetailedCountryFieldRequest) SetName(v string) {
	o.Name = v
}

func (o DetailedCountryFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDetailedCountryFieldRequest struct {
	value *DetailedCountryFieldRequest
	isSet bool
}

func (v NullableDetailedCountryFieldRequest) Get() *DetailedCountryFieldRequest {
	return v.value
}

func (v *NullableDetailedCountryFieldRequest) Set(val *DetailedCountryFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedCountryFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedCountryFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedCountryFieldRequest(val *DetailedCountryFieldRequest) *NullableDetailedCountryFieldRequest {
	return &NullableDetailedCountryFieldRequest{value: val, isSet: true}
}

func (v NullableDetailedCountryFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedCountryFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
