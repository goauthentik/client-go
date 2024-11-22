/*
authentik

Making authentication simple.

API version: 2024.10.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UserMetrics User Metrics
type UserMetrics struct {
	Logins         []Coordinate `json:"logins"`
	LoginsFailed   []Coordinate `json:"logins_failed"`
	Authorizations []Coordinate `json:"authorizations"`
}

// NewUserMetrics instantiates a new UserMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetrics(logins []Coordinate, loginsFailed []Coordinate, authorizations []Coordinate) *UserMetrics {
	this := UserMetrics{}
	this.Logins = logins
	this.LoginsFailed = loginsFailed
	this.Authorizations = authorizations
	return &this
}

// NewUserMetricsWithDefaults instantiates a new UserMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetricsWithDefaults() *UserMetrics {
	this := UserMetrics{}
	return &this
}

// GetLogins returns the Logins field value
func (o *UserMetrics) GetLogins() []Coordinate {
	if o == nil {
		var ret []Coordinate
		return ret
	}

	return o.Logins
}

// GetLoginsOk returns a tuple with the Logins field value
// and a boolean to check if the value has been set.
func (o *UserMetrics) GetLoginsOk() ([]Coordinate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logins, true
}

// SetLogins sets field value
func (o *UserMetrics) SetLogins(v []Coordinate) {
	o.Logins = v
}

// GetLoginsFailed returns the LoginsFailed field value
func (o *UserMetrics) GetLoginsFailed() []Coordinate {
	if o == nil {
		var ret []Coordinate
		return ret
	}

	return o.LoginsFailed
}

// GetLoginsFailedOk returns a tuple with the LoginsFailed field value
// and a boolean to check if the value has been set.
func (o *UserMetrics) GetLoginsFailedOk() ([]Coordinate, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoginsFailed, true
}

// SetLoginsFailed sets field value
func (o *UserMetrics) SetLoginsFailed(v []Coordinate) {
	o.LoginsFailed = v
}

// GetAuthorizations returns the Authorizations field value
func (o *UserMetrics) GetAuthorizations() []Coordinate {
	if o == nil {
		var ret []Coordinate
		return ret
	}

	return o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value
// and a boolean to check if the value has been set.
func (o *UserMetrics) GetAuthorizationsOk() ([]Coordinate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authorizations, true
}

// SetAuthorizations sets field value
func (o *UserMetrics) SetAuthorizations(v []Coordinate) {
	o.Authorizations = v
}

func (o UserMetrics) MarshalJSON() ([]byte, error) {
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

type NullableUserMetrics struct {
	value *UserMetrics
	isSet bool
}

func (v NullableUserMetrics) Get() *UserMetrics {
	return v.value
}

func (v *NullableUserMetrics) Set(val *UserMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetrics(val *UserMetrics) *NullableUserMetrics {
	return &NullableUserMetrics{value: val, isSet: true}
}

func (v NullableUserMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
