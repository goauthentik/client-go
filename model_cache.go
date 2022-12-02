/*
authentik

Making authentication simple.

API version: 2022.11.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Cache Generic cache stats for an object
type Cache struct {
	Count int32 `json:"count"`
}

// NewCache instantiates a new Cache object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCache(count int32) *Cache {
	this := Cache{}
	this.Count = count
	return &this
}

// NewCacheWithDefaults instantiates a new Cache object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCacheWithDefaults() *Cache {
	this := Cache{}
	return &this
}

// GetCount returns the Count field value
func (o *Cache) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *Cache) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *Cache) SetCount(v int32) {
	o.Count = v
}

func (o Cache) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableCache struct {
	value *Cache
	isSet bool
}

func (v NullableCache) Get() *Cache {
	return v.value
}

func (v *NullableCache) Set(val *Cache) {
	v.value = val
	v.isSet = true
}

func (v NullableCache) IsSet() bool {
	return v.isSet
}

func (v *NullableCache) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCache(val *Cache) *NullableCache {
	return &NullableCache{value: val, isSet: true}
}

func (v NullableCache) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCache) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
