/*
authentik

Making authentication simple.

API version: 2022.9.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Metadata Serializer for blueprint metadata
type Metadata struct {
	Name   string                 `json:"name"`
	Labels map[string]interface{} `json:"labels"`
}

// NewMetadata instantiates a new Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadata(name string, labels map[string]interface{}) *Metadata {
	this := Metadata{}
	this.Name = name
	this.Labels = labels
	return &this
}

// NewMetadataWithDefaults instantiates a new Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithDefaults() *Metadata {
	this := Metadata{}
	return &this
}

// GetName returns the Name field value
func (o *Metadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Metadata) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value
func (o *Metadata) GetLabels() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *Metadata) GetLabelsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *Metadata) SetLabels(v map[string]interface{}) {
	o.Labels = v
}

func (o Metadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableMetadata struct {
	value *Metadata
	isSet bool
}

func (v NullableMetadata) Get() *Metadata {
	return v.value
}

func (v *NullableMetadata) Set(val *Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata(val *Metadata) *NullableMetadata {
	return &NullableMetadata{value: val, isSet: true}
}

func (v NullableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
