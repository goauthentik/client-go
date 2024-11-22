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

// LicenseForecast Serializer for license forecast
type LicenseForecast struct {
	InternalUsers           int32 `json:"internal_users"`
	ExternalUsers           int32 `json:"external_users"`
	ForecastedInternalUsers int32 `json:"forecasted_internal_users"`
	ForecastedExternalUsers int32 `json:"forecasted_external_users"`
}

// NewLicenseForecast instantiates a new LicenseForecast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseForecast(internalUsers int32, externalUsers int32, forecastedInternalUsers int32, forecastedExternalUsers int32) *LicenseForecast {
	this := LicenseForecast{}
	this.InternalUsers = internalUsers
	this.ExternalUsers = externalUsers
	this.ForecastedInternalUsers = forecastedInternalUsers
	this.ForecastedExternalUsers = forecastedExternalUsers
	return &this
}

// NewLicenseForecastWithDefaults instantiates a new LicenseForecast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseForecastWithDefaults() *LicenseForecast {
	this := LicenseForecast{}
	return &this
}

// GetInternalUsers returns the InternalUsers field value
func (o *LicenseForecast) GetInternalUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InternalUsers
}

// GetInternalUsersOk returns a tuple with the InternalUsers field value
// and a boolean to check if the value has been set.
func (o *LicenseForecast) GetInternalUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalUsers, true
}

// SetInternalUsers sets field value
func (o *LicenseForecast) SetInternalUsers(v int32) {
	o.InternalUsers = v
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

// GetForecastedInternalUsers returns the ForecastedInternalUsers field value
func (o *LicenseForecast) GetForecastedInternalUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ForecastedInternalUsers
}

// GetForecastedInternalUsersOk returns a tuple with the ForecastedInternalUsers field value
// and a boolean to check if the value has been set.
func (o *LicenseForecast) GetForecastedInternalUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForecastedInternalUsers, true
}

// SetForecastedInternalUsers sets field value
func (o *LicenseForecast) SetForecastedInternalUsers(v int32) {
	o.ForecastedInternalUsers = v
}

// GetForecastedExternalUsers returns the ForecastedExternalUsers field value
func (o *LicenseForecast) GetForecastedExternalUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ForecastedExternalUsers
}

// GetForecastedExternalUsersOk returns a tuple with the ForecastedExternalUsers field value
// and a boolean to check if the value has been set.
func (o *LicenseForecast) GetForecastedExternalUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForecastedExternalUsers, true
}

// SetForecastedExternalUsers sets field value
func (o *LicenseForecast) SetForecastedExternalUsers(v int32) {
	o.ForecastedExternalUsers = v
}

func (o LicenseForecast) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["internal_users"] = o.InternalUsers
	}
	if true {
		toSerialize["external_users"] = o.ExternalUsers
	}
	if true {
		toSerialize["forecasted_internal_users"] = o.ForecastedInternalUsers
	}
	if true {
		toSerialize["forecasted_external_users"] = o.ForecastedExternalUsers
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
