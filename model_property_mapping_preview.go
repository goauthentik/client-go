/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PropertyMappingPreview Preview how the current user is mapped via the property mappings selected in a provider
type PropertyMappingPreview struct {
	Preview map[string]interface{} `json:"preview"`
}

// NewPropertyMappingPreview instantiates a new PropertyMappingPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyMappingPreview(preview map[string]interface{}) *PropertyMappingPreview {
	this := PropertyMappingPreview{}
	this.Preview = preview
	return &this
}

// NewPropertyMappingPreviewWithDefaults instantiates a new PropertyMappingPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyMappingPreviewWithDefaults() *PropertyMappingPreview {
	this := PropertyMappingPreview{}
	return &this
}

// GetPreview returns the Preview field value
func (o *PropertyMappingPreview) GetPreview() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value
// and a boolean to check if the value has been set.
func (o *PropertyMappingPreview) GetPreviewOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preview, true
}

// SetPreview sets field value
func (o *PropertyMappingPreview) SetPreview(v map[string]interface{}) {
	o.Preview = v
}

func (o PropertyMappingPreview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["preview"] = o.Preview
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyMappingPreview struct {
	value *PropertyMappingPreview
	isSet bool
}

func (v NullablePropertyMappingPreview) Get() *PropertyMappingPreview {
	return v.value
}

func (v *NullablePropertyMappingPreview) Set(val *PropertyMappingPreview) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyMappingPreview) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyMappingPreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyMappingPreview(val *PropertyMappingPreview) *NullablePropertyMappingPreview {
	return &NullablePropertyMappingPreview{value: val, isSet: true}
}

func (v NullablePropertyMappingPreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyMappingPreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
