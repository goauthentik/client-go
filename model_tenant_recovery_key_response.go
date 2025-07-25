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
	"time"
)

// TenantRecoveryKeyResponse Tenant recovery key creation response serializer
type TenantRecoveryKeyResponse struct {
	Expiry time.Time `json:"expiry"`
	Url    string    `json:"url"`
}

// NewTenantRecoveryKeyResponse instantiates a new TenantRecoveryKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantRecoveryKeyResponse(expiry time.Time, url string) *TenantRecoveryKeyResponse {
	this := TenantRecoveryKeyResponse{}
	this.Expiry = expiry
	this.Url = url
	return &this
}

// NewTenantRecoveryKeyResponseWithDefaults instantiates a new TenantRecoveryKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantRecoveryKeyResponseWithDefaults() *TenantRecoveryKeyResponse {
	this := TenantRecoveryKeyResponse{}
	return &this
}

// GetExpiry returns the Expiry field value
func (o *TenantRecoveryKeyResponse) GetExpiry() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value
// and a boolean to check if the value has been set.
func (o *TenantRecoveryKeyResponse) GetExpiryOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiry, true
}

// SetExpiry sets field value
func (o *TenantRecoveryKeyResponse) SetExpiry(v time.Time) {
	o.Expiry = v
}

// GetUrl returns the Url field value
func (o *TenantRecoveryKeyResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TenantRecoveryKeyResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TenantRecoveryKeyResponse) SetUrl(v string) {
	o.Url = v
}

func (o TenantRecoveryKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["expiry"] = o.Expiry
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableTenantRecoveryKeyResponse struct {
	value *TenantRecoveryKeyResponse
	isSet bool
}

func (v NullableTenantRecoveryKeyResponse) Get() *TenantRecoveryKeyResponse {
	return v.value
}

func (v *NullableTenantRecoveryKeyResponse) Set(val *TenantRecoveryKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantRecoveryKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantRecoveryKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantRecoveryKeyResponse(val *TenantRecoveryKeyResponse) *NullableTenantRecoveryKeyResponse {
	return &NullableTenantRecoveryKeyResponse{value: val, isSet: true}
}

func (v NullableTenantRecoveryKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantRecoveryKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
