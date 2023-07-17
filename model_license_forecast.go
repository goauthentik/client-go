/*
authentik

Making authentication simple.

API version: 2023.6.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LicenseForecast Serializer for license forecast
type LicenseForecast struct {
	Users         int32 `json:"users"`
	ExternalUsers int32 `json:"external_users"`
}

// NewLicenseForecast instantiates a new LicenseForecast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseForecast(users int32, externalUsers int32) *LicenseForecast {
	this := LicenseForecast{}
	this.Users = users
	this.ExternalUsers = externalUsers
	return &this
}

// NewLicenseForecastWithDefaults instantiates a new LicenseForecast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseForecastWithDefaults() *LicenseForecast {
	this := LicenseForecast{}
	return &this
}

// GetUsers returns the Users field value
func (o *LicenseForecast) GetUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *LicenseForecast) GetUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value
func (o *LicenseForecast) SetUsers(v int32) {
	o.Users = v
}

// GetExternalUsers returns the ExternalUsers field value
func (o *LicenseForecast) GetExternalUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExternalUsers
}

// GetExternalUsersOk returns a tuple with the ExternalUsers field value
// and a boolean to check if the value has been set.
func (o *LicenseForecast) GetExternalUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUsers, true
}

// SetExternalUsers sets field value
func (o *LicenseForecast) SetExternalUsers(v int32) {
	o.ExternalUsers = v
}

func (o LicenseForecast) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["users"] = o.Users
	}
	if true {
		toSerialize["external_users"] = o.ExternalUsers
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseForecast struct {
	value *LicenseForecast
	isSet bool
}

func (v NullableLicenseForecast) Get() *LicenseForecast {
	return v.value
}

func (v *NullableLicenseForecast) Set(val *LicenseForecast) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseForecast) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseForecast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseForecast(val *LicenseForecast) *NullableLicenseForecast {
	return &NullableLicenseForecast{value: val, isSet: true}
}

func (v NullableLicenseForecast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseForecast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
