/*
authentik

Making authentication simple.

API version: 2022.10.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BlueprintFileMeta struct for BlueprintFileMeta
type BlueprintFileMeta struct {
	Name   string                 `json:"name"`
	Labels map[string]interface{} `json:"labels"`
}

// NewBlueprintFileMeta instantiates a new BlueprintFileMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintFileMeta(name string, labels map[string]interface{}) *BlueprintFileMeta {
	this := BlueprintFileMeta{}
	this.Name = name
	this.Labels = labels
	return &this
}

// NewBlueprintFileMetaWithDefaults instantiates a new BlueprintFileMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintFileMetaWithDefaults() *BlueprintFileMeta {
	this := BlueprintFileMeta{}
	return &this
}

// GetName returns the Name field value
func (o *BlueprintFileMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlueprintFileMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BlueprintFileMeta) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value
func (o *BlueprintFileMeta) GetLabels() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *BlueprintFileMeta) GetLabelsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *BlueprintFileMeta) SetLabels(v map[string]interface{}) {
	o.Labels = v
}

func (o BlueprintFileMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableBlueprintFileMeta struct {
	value *BlueprintFileMeta
	isSet bool
}

func (v NullableBlueprintFileMeta) Get() *BlueprintFileMeta {
	return v.value
}

func (v *NullableBlueprintFileMeta) Set(val *BlueprintFileMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintFileMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintFileMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintFileMeta(val *BlueprintFileMeta) *NullableBlueprintFileMeta {
	return &NullableBlueprintFileMeta{value: val, isSet: true}
}

func (v NullableBlueprintFileMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintFileMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
