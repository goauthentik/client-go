/*
authentik

Making authentication simple.

API version: 2023.3.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SAMLMetadata SAML Provider Metadata serializer
type SAMLMetadata struct {
	Metadata    string `json:"metadata"`
	DownloadUrl string `json:"download_url"`
}

// NewSAMLMetadata instantiates a new SAMLMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLMetadata(metadata string, downloadUrl string) *SAMLMetadata {
	this := SAMLMetadata{}
	this.Metadata = metadata
	this.DownloadUrl = downloadUrl
	return &this
}

// NewSAMLMetadataWithDefaults instantiates a new SAMLMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLMetadataWithDefaults() *SAMLMetadata {
	this := SAMLMetadata{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *SAMLMetadata) GetMetadata() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SAMLMetadata) GetMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *SAMLMetadata) SetMetadata(v string) {
	o.Metadata = v
}

// GetDownloadUrl returns the DownloadUrl field value
func (o *SAMLMetadata) GetDownloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value
// and a boolean to check if the value has been set.
func (o *SAMLMetadata) GetDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadUrl, true
}

// SetDownloadUrl sets field value
func (o *SAMLMetadata) SetDownloadUrl(v string) {
	o.DownloadUrl = v
}

func (o SAMLMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["download_url"] = o.DownloadUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSAMLMetadata struct {
	value *SAMLMetadata
	isSet bool
}

func (v NullableSAMLMetadata) Get() *SAMLMetadata {
	return v.value
}

func (v *NullableSAMLMetadata) Set(val *SAMLMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLMetadata(val *SAMLMetadata) *NullableSAMLMetadata {
	return &NullableSAMLMetadata{value: val, isSet: true}
}

func (v NullableSAMLMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
