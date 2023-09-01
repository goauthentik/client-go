/*
authentik

Making authentication simple.

API version: 2023.8.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GenericError Generic API Error
type GenericError struct {
	Detail string  `json:"detail"`
	Code   *string `json:"code,omitempty"`
}

// NewGenericError instantiates a new GenericError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericError(detail string) *GenericError {
	this := GenericError{}
	this.Detail = detail
	return &this
}

// NewGenericErrorWithDefaults instantiates a new GenericError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericErrorWithDefaults() *GenericError {
	this := GenericError{}
	return &this
}

// GetDetail returns the Detail field value
func (o *GenericError) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *GenericError) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *GenericError) SetDetail(v string) {
	o.Detail = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GenericError) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericError) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GenericError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GenericError) SetCode(v string) {
	o.Code = &v
}

func (o GenericError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["detail"] = o.Detail
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableGenericError struct {
	value *GenericError
	isSet bool
}

func (v NullableGenericError) Get() *GenericError {
	return v.value
}

func (v *NullableGenericError) Set(val *GenericError) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericError) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericError(val *GenericError) *NullableGenericError {
	return &NullableGenericError{value: val, isSet: true}
}

func (v NullableGenericError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
