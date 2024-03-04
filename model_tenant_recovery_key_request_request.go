/*
authentik

Making authentication simple.

API version: 2024.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TenantRecoveryKeyRequestRequest Tenant recovery key creation request serializer
type TenantRecoveryKeyRequestRequest struct {
	User         string `json:"user"`
	DurationDays int32  `json:"duration_days"`
}

// NewTenantRecoveryKeyRequestRequest instantiates a new TenantRecoveryKeyRequestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantRecoveryKeyRequestRequest(user string, durationDays int32) *TenantRecoveryKeyRequestRequest {
	this := TenantRecoveryKeyRequestRequest{}
	this.User = user
	this.DurationDays = durationDays
	return &this
}

// NewTenantRecoveryKeyRequestRequestWithDefaults instantiates a new TenantRecoveryKeyRequestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantRecoveryKeyRequestRequestWithDefaults() *TenantRecoveryKeyRequestRequest {
	this := TenantRecoveryKeyRequestRequest{}
	return &this
}

// GetUser returns the User field value
func (o *TenantRecoveryKeyRequestRequest) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *TenantRecoveryKeyRequestRequest) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *TenantRecoveryKeyRequestRequest) SetUser(v string) {
	o.User = v
}

// GetDurationDays returns the DurationDays field value
func (o *TenantRecoveryKeyRequestRequest) GetDurationDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationDays
}

// GetDurationDaysOk returns a tuple with the DurationDays field value
// and a boolean to check if the value has been set.
func (o *TenantRecoveryKeyRequestRequest) GetDurationDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationDays, true
}

// SetDurationDays sets field value
func (o *TenantRecoveryKeyRequestRequest) SetDurationDays(v int32) {
	o.DurationDays = v
}

func (o TenantRecoveryKeyRequestRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["duration_days"] = o.DurationDays
	}
	return json.Marshal(toSerialize)
}

type NullableTenantRecoveryKeyRequestRequest struct {
	value *TenantRecoveryKeyRequestRequest
	isSet bool
}

func (v NullableTenantRecoveryKeyRequestRequest) Get() *TenantRecoveryKeyRequestRequest {
	return v.value
}

func (v *NullableTenantRecoveryKeyRequestRequest) Set(val *TenantRecoveryKeyRequestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantRecoveryKeyRequestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantRecoveryKeyRequestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantRecoveryKeyRequestRequest(val *TenantRecoveryKeyRequestRequest) *NullableTenantRecoveryKeyRequestRequest {
	return &NullableTenantRecoveryKeyRequestRequest{value: val, isSet: true}
}

func (v NullableTenantRecoveryKeyRequestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantRecoveryKeyRequestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
