/*
authentik

Making authentication simple.

API version: 2022.7.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TokenView Show token's current key
type TokenView struct {
	Key string `json:"key"`
}

// NewTokenView instantiates a new TokenView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenView(key string) *TokenView {
	this := TokenView{}
	this.Key = key
	return &this
}

// NewTokenViewWithDefaults instantiates a new TokenView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenViewWithDefaults() *TokenView {
	this := TokenView{}
	return &this
}

// GetKey returns the Key field value
func (o *TokenView) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TokenView) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TokenView) SetKey(v string) {
	o.Key = v
}

func (o TokenView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableTokenView struct {
	value *TokenView
	isSet bool
}

func (v NullableTokenView) Get() *TokenView {
	return v.value
}

func (v *NullableTokenView) Set(val *TokenView) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenView) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenView(val *TokenView) *NullableTokenView {
	return &NullableTokenView{value: val, isSet: true}
}

func (v NullableTokenView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
