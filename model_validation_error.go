/*
authentik

Making authentication simple.

API version: 2024.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ValidationError Validation Error
type ValidationError struct {
	NonFieldErrors []string `json:"non_field_errors,omitempty"`
	Code           *string  `json:"code,omitempty"`
}

// NewValidationError instantiates a new ValidationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationError() *ValidationError {
	this := ValidationError{}
	return &this
}

// NewValidationErrorWithDefaults instantiates a new ValidationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorWithDefaults() *ValidationError {
	this := ValidationError{}
	return &this
}

// GetNonFieldErrors returns the NonFieldErrors field value if set, zero value otherwise.
func (o *ValidationError) GetNonFieldErrors() []string {
	if o == nil || o.NonFieldErrors == nil {
		var ret []string
		return ret
	}
	return o.NonFieldErrors
}

// GetNonFieldErrorsOk returns a tuple with the NonFieldErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationError) GetNonFieldErrorsOk() ([]string, bool) {
	if o == nil || o.NonFieldErrors == nil {
		return nil, false
	}
	return o.NonFieldErrors, true
}

// HasNonFieldErrors returns a boolean if a field has been set.
func (o *ValidationError) HasNonFieldErrors() bool {
	if o != nil && o.NonFieldErrors != nil {
		return true
	}

	return false
}

// SetNonFieldErrors gets a reference to the given []string and assigns it to the NonFieldErrors field.
func (o *ValidationError) SetNonFieldErrors(v []string) {
	o.NonFieldErrors = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ValidationError) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationError) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ValidationError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ValidationError) SetCode(v string) {
	o.Code = &v
}

func (o ValidationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NonFieldErrors != nil {
		toSerialize["non_field_errors"] = o.NonFieldErrors
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableValidationError struct {
	value *ValidationError
	isSet bool
}

func (v NullableValidationError) Get() *ValidationError {
	return v.value
}

func (v *NullableValidationError) Set(val *ValidationError) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationError) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationError(val *ValidationError) *NullableValidationError {
	return &NullableValidationError{value: val, isSet: true}
}

func (v NullableValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
