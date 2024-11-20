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

// checks if the PropertyMappingTestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyMappingTestResult{}

// PropertyMappingTestResult Result of a Property-mapping test
type PropertyMappingTestResult struct {
	Result     string `json:"result"`
	Successful bool   `json:"successful"`
}

type _PropertyMappingTestResult PropertyMappingTestResult

// NewPropertyMappingTestResult instantiates a new PropertyMappingTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyMappingTestResult(result string, successful bool) *PropertyMappingTestResult {
	this := PropertyMappingTestResult{}
	this.Result = result
	this.Successful = successful
	return &this
}

// NewPropertyMappingTestResultWithDefaults instantiates a new PropertyMappingTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyMappingTestResultWithDefaults() *PropertyMappingTestResult {
	this := PropertyMappingTestResult{}
	return &this
}

// GetResult returns the Result field value
func (o *PropertyMappingTestResult) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *PropertyMappingTestResult) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *PropertyMappingTestResult) SetResult(v string) {
	o.Result = v
}

// GetSuccessful returns the Successful field value
func (o *PropertyMappingTestResult) GetSuccessful() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Successful
}

// GetSuccessfulOk returns a tuple with the Successful field value
// and a boolean to check if the value has been set.
func (o *PropertyMappingTestResult) GetSuccessfulOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Successful, true
}

// SetSuccessful sets field value
func (o *PropertyMappingTestResult) SetSuccessful(v bool) {
	o.Successful = v
}

func (o PropertyMappingTestResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyMappingTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["successful"] = o.Successful
	return toSerialize, nil
}

func (o *PropertyMappingTestResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
		"successful",
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

	varPropertyMappingTestResult := _PropertyMappingTestResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPropertyMappingTestResult)

	if err != nil {
		return err
	}

	*o = PropertyMappingTestResult(varPropertyMappingTestResult)

	return err
}

type NullablePropertyMappingTestResult struct {
	value *PropertyMappingTestResult
	isSet bool
}

func (v NullablePropertyMappingTestResult) Get() *PropertyMappingTestResult {
	return v.value
}

func (v *NullablePropertyMappingTestResult) Set(val *PropertyMappingTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyMappingTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyMappingTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyMappingTestResult(val *PropertyMappingTestResult) *NullablePropertyMappingTestResult {
	return &NullablePropertyMappingTestResult{value: val, isSet: true}
}

func (v NullablePropertyMappingTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyMappingTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
