/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LoginMetrics Login Metrics per 1h
type LoginMetrics struct {
	Logins         []Coordinate `json:"logins"`
	LoginsFailed   []Coordinate `json:"logins_failed"`
	Authorizations []Coordinate `json:"authorizations"`
}

// NewLoginMetrics instantiates a new LoginMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginMetrics(logins []Coordinate, loginsFailed []Coordinate, authorizations []Coordinate) *LoginMetrics {
	this := LoginMetrics{}
	this.Logins = logins
	this.LoginsFailed = loginsFailed
	this.Authorizations = authorizations
	return &this
}

// NewLoginMetricsWithDefaults instantiates a new LoginMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginMetricsWithDefaults() *LoginMetrics {
	this := LoginMetrics{}
	return &this
}

// GetLogins returns the Logins field value
func (o *LoginMetrics) GetLogins() []Coordinate {
	if o == nil {
		var ret []Coordinate
		return ret
	}

	return o.Logins
}

// GetLoginsOk returns a tuple with the Logins field value
// and a boolean to check if the value has been set.
func (o *LoginMetrics) GetLoginsOk() ([]Coordinate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logins, true
}

// SetLogins sets field value
func (o *LoginMetrics) SetLogins(v []Coordinate) {
	o.Logins = v
}

// GetLoginsFailed returns the LoginsFailed field value
func (o *LoginMetrics) GetLoginsFailed() []Coordinate {
	if o == nil {
		var ret []Coordinate
		return ret
	}

	return o.LoginsFailed
}

// GetLoginsFailedOk returns a tuple with the LoginsFailed field value
// and a boolean to check if the value has been set.
func (o *LoginMetrics) GetLoginsFailedOk() ([]Coordinate, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoginsFailed, true
}

// SetLoginsFailed sets field value
func (o *LoginMetrics) SetLoginsFailed(v []Coordinate) {
	o.LoginsFailed = v
}

// GetAuthorizations returns the Authorizations field value
func (o *LoginMetrics) GetAuthorizations() []Coordinate {
	if o == nil {
		var ret []Coordinate
		return ret
	}

	return o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value
// and a boolean to check if the value has been set.
func (o *LoginMetrics) GetAuthorizationsOk() ([]Coordinate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authorizations, true
}

// SetAuthorizations sets field value
func (o *LoginMetrics) SetAuthorizations(v []Coordinate) {
	o.Authorizations = v
}

func (o LoginMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["logins"] = o.Logins
	}
	if true {
		toSerialize["logins_failed"] = o.LoginsFailed
	}
	if true {
		toSerialize["authorizations"] = o.Authorizations
	}
	return json.Marshal(toSerialize)
}

type NullableLoginMetrics struct {
	value *LoginMetrics
	isSet bool
}

func (v NullableLoginMetrics) Get() *LoginMetrics {
	return v.value
}

func (v *NullableLoginMetrics) Set(val *LoginMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginMetrics(val *LoginMetrics) *NullableLoginMetrics {
	return &NullableLoginMetrics{value: val, isSet: true}
}

func (v NullableLoginMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
