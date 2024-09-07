/*
authentik

Making authentication simple.

API version: 2024.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TokenSetKeyRequest struct for TokenSetKeyRequest
type TokenSetKeyRequest struct {
	Key string `json:"key"`
}

// NewTokenSetKeyRequest instantiates a new TokenSetKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenSetKeyRequest(key string) *TokenSetKeyRequest {
	this := TokenSetKeyRequest{}
	this.Key = key
	return &this
}

// NewTokenSetKeyRequestWithDefaults instantiates a new TokenSetKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenSetKeyRequestWithDefaults() *TokenSetKeyRequest {
	this := TokenSetKeyRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *TokenSetKeyRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TokenSetKeyRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TokenSetKeyRequest) SetKey(v string) {
	o.Key = v
}

func (o TokenSetKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableTokenSetKeyRequest struct {
	value *TokenSetKeyRequest
	isSet bool
}

func (v NullableTokenSetKeyRequest) Get() *TokenSetKeyRequest {
	return v.value
}

func (v *NullableTokenSetKeyRequest) Set(val *TokenSetKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenSetKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenSetKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenSetKeyRequest(val *TokenSetKeyRequest) *NullableTokenSetKeyRequest {
	return &NullableTokenSetKeyRequest{value: val, isSet: true}
}

func (v NullableTokenSetKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenSetKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
