/*
authentik

Making authentication simple.

API version: 2021.9.7
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FilePathRequest Serializer to upload file
type FilePathRequest struct {
	Url string `json:"url"`
}

// NewFilePathRequest instantiates a new FilePathRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilePathRequest(url string) *FilePathRequest {
	this := FilePathRequest{}
	this.Url = url
	return &this
}

// NewFilePathRequestWithDefaults instantiates a new FilePathRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilePathRequestWithDefaults() *FilePathRequest {
	this := FilePathRequest{}
	return &this
}

// GetUrl returns the Url field value
func (o *FilePathRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FilePathRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FilePathRequest) SetUrl(v string) {
	o.Url = v
}

func (o FilePathRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableFilePathRequest struct {
	value *FilePathRequest
	isSet bool
}

func (v NullableFilePathRequest) Get() *FilePathRequest {
	return v.value
}

func (v *NullableFilePathRequest) Set(val *FilePathRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFilePathRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFilePathRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilePathRequest(val *FilePathRequest) *NullableFilePathRequest {
	return &NullableFilePathRequest{value: val, isSet: true}
}

func (v NullableFilePathRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilePathRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
