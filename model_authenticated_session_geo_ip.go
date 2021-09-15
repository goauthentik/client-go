/*
authentik

Making authentication simple.

API version: 2021.9.1-rc1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatedSessionGeoIp struct for AuthenticatedSessionGeoIp
type AuthenticatedSessionGeoIp struct {
	Continent *string `json:"continent,omitempty"`
	Country *string `json:"country,omitempty"`
	Lat *float32 `json:"lat,omitempty"`
	Long *float32 `json:"long,omitempty"`
	City *string `json:"city,omitempty"`
}

// NewAuthenticatedSessionGeoIp instantiates a new AuthenticatedSessionGeoIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSessionGeoIp() *AuthenticatedSessionGeoIp {
	this := AuthenticatedSessionGeoIp{}
	return &this
}

// NewAuthenticatedSessionGeoIpWithDefaults instantiates a new AuthenticatedSessionGeoIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionGeoIpWithDefaults() *AuthenticatedSessionGeoIp {
	this := AuthenticatedSessionGeoIp{}
	return &this
}

// GetContinent returns the Continent field value if set, zero value otherwise.
func (o *AuthenticatedSessionGeoIp) GetContinent() string {
	if o == nil || o.Continent == nil {
		var ret string
		return ret
	}
	return *o.Continent
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionGeoIp) GetContinentOk() (*string, bool) {
	if o == nil || o.Continent == nil {
		return nil, false
	}
	return o.Continent, true
}

// HasContinent returns a boolean if a field has been set.
func (o *AuthenticatedSessionGeoIp) HasContinent() bool {
	if o != nil && o.Continent != nil {
		return true
	}

	return false
}

// SetContinent gets a reference to the given string and assigns it to the Continent field.
func (o *AuthenticatedSessionGeoIp) SetContinent(v string) {
	o.Continent = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AuthenticatedSessionGeoIp) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionGeoIp) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AuthenticatedSessionGeoIp) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AuthenticatedSessionGeoIp) SetCountry(v string) {
	o.Country = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *AuthenticatedSessionGeoIp) GetLat() float32 {
	if o == nil || o.Lat == nil {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionGeoIp) GetLatOk() (*float32, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *AuthenticatedSessionGeoIp) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *AuthenticatedSessionGeoIp) SetLat(v float32) {
	o.Lat = &v
}

// GetLong returns the Long field value if set, zero value otherwise.
func (o *AuthenticatedSessionGeoIp) GetLong() float32 {
	if o == nil || o.Long == nil {
		var ret float32
		return ret
	}
	return *o.Long
}

// GetLongOk returns a tuple with the Long field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionGeoIp) GetLongOk() (*float32, bool) {
	if o == nil || o.Long == nil {
		return nil, false
	}
	return o.Long, true
}

// HasLong returns a boolean if a field has been set.
func (o *AuthenticatedSessionGeoIp) HasLong() bool {
	if o != nil && o.Long != nil {
		return true
	}

	return false
}

// SetLong gets a reference to the given float32 and assigns it to the Long field.
func (o *AuthenticatedSessionGeoIp) SetLong(v float32) {
	o.Long = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AuthenticatedSessionGeoIp) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionGeoIp) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AuthenticatedSessionGeoIp) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AuthenticatedSessionGeoIp) SetCity(v string) {
	o.City = &v
}

func (o AuthenticatedSessionGeoIp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Continent != nil {
		toSerialize["continent"] = o.Continent
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Long != nil {
		toSerialize["long"] = o.Long
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatedSessionGeoIp struct {
	value *AuthenticatedSessionGeoIp
	isSet bool
}

func (v NullableAuthenticatedSessionGeoIp) Get() *AuthenticatedSessionGeoIp {
	return v.value
}

func (v *NullableAuthenticatedSessionGeoIp) Set(val *AuthenticatedSessionGeoIp) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatedSessionGeoIp) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatedSessionGeoIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatedSessionGeoIp(val *AuthenticatedSessionGeoIp) *NullableAuthenticatedSessionGeoIp {
	return &NullableAuthenticatedSessionGeoIp{value: val, isSet: true}
}

func (v NullableAuthenticatedSessionGeoIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatedSessionGeoIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


