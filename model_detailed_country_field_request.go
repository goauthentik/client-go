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

// checks if the DetailedCountryFieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailedCountryFieldRequest{}

// DetailedCountryFieldRequest struct for DetailedCountryFieldRequest
type DetailedCountryFieldRequest struct {
	Code CountryCodeEnum `json:"code"`
	Name string          `json:"name"`
}

type _DetailedCountryFieldRequest DetailedCountryFieldRequest

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailedCountryFieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *DetailedCountryFieldRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"name",
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

	varDetailedCountryFieldRequest := _DetailedCountryFieldRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDetailedCountryFieldRequest)

	if err != nil {
		return err
	}

	*o = DetailedCountryFieldRequest(varDetailedCountryFieldRequest)

	return err
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
