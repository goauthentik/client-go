/*
authentik

Making authentication simple.

API version: 2023.4.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// BlueprintFile struct for BlueprintFile
type BlueprintFile struct {
	Path  string    `json:"path"`
	LastM time.Time `json:"last_m"`
	Hash  string    `json:"hash"`
	Meta  Metadata  `json:"meta"`
}

// NewBlueprintFile instantiates a new BlueprintFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintFile(path string, lastM time.Time, hash string, meta Metadata) *BlueprintFile {
	this := BlueprintFile{}
	this.Path = path
	this.LastM = lastM
	this.Hash = hash
	this.Meta = meta
	return &this
}

// NewBlueprintFileWithDefaults instantiates a new BlueprintFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintFileWithDefaults() *BlueprintFile {
	this := BlueprintFile{}
	return &this
}

// GetPath returns the Path field value
func (o *BlueprintFile) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *BlueprintFile) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *BlueprintFile) SetPath(v string) {
	o.Path = v
}

// GetLastM returns the LastM field value
func (o *BlueprintFile) GetLastM() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastM
}

// GetLastMOk returns a tuple with the LastM field value
// and a boolean to check if the value has been set.
func (o *BlueprintFile) GetLastMOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastM, true
}

// SetLastM sets field value
func (o *BlueprintFile) SetLastM(v time.Time) {
	o.LastM = v
}

// GetHash returns the Hash field value
func (o *BlueprintFile) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *BlueprintFile) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *BlueprintFile) SetHash(v string) {
	o.Hash = v
}

// GetMeta returns the Meta field value
func (o *BlueprintFile) GetMeta() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *BlueprintFile) GetMetaOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *BlueprintFile) SetMeta(v Metadata) {
	o.Meta = v
}

func (o BlueprintFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["last_m"] = o.LastM
	}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableBlueprintFile struct {
	value *BlueprintFile
	isSet bool
}

func (v NullableBlueprintFile) Get() *BlueprintFile {
	return v.value
}

func (v *NullableBlueprintFile) Set(val *BlueprintFile) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintFile) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintFile(val *BlueprintFile) *NullableBlueprintFile {
	return &NullableBlueprintFile{value: val, isSet: true}
}

func (v NullableBlueprintFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
