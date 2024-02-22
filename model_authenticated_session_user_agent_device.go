/*
authentik

Making authentication simple.

API version: 2024.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatedSessionUserAgentDevice User agent device
type AuthenticatedSessionUserAgentDevice struct {
	Brand  string `json:"brand"`
	Family string `json:"family"`
	Model  string `json:"model"`
}

// NewAuthenticatedSessionUserAgentDevice instantiates a new AuthenticatedSessionUserAgentDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSessionUserAgentDevice(brand string, family string, model string) *AuthenticatedSessionUserAgentDevice {
	this := AuthenticatedSessionUserAgentDevice{}
	this.Brand = brand
	this.Family = family
	this.Model = model
	return &this
}

// NewAuthenticatedSessionUserAgentDeviceWithDefaults instantiates a new AuthenticatedSessionUserAgentDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionUserAgentDeviceWithDefaults() *AuthenticatedSessionUserAgentDevice {
	this := AuthenticatedSessionUserAgentDevice{}
	return &this
}

// GetBrand returns the Brand field value
func (o *AuthenticatedSessionUserAgentDevice) GetBrand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Brand
}

// GetBrandOk returns a tuple with the Brand field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentDevice) GetBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Brand, true
}

// SetBrand sets field value
func (o *AuthenticatedSessionUserAgentDevice) SetBrand(v string) {
	o.Brand = v
}

// GetFamily returns the Family field value
func (o *AuthenticatedSessionUserAgentDevice) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentDevice) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *AuthenticatedSessionUserAgentDevice) SetFamily(v string) {
	o.Family = v
}

// GetModel returns the Model field value
func (o *AuthenticatedSessionUserAgentDevice) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentDevice) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *AuthenticatedSessionUserAgentDevice) SetModel(v string) {
	o.Model = v
}

func (o AuthenticatedSessionUserAgentDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["brand"] = o.Brand
	}
	if true {
		toSerialize["family"] = o.Family
	}
	if true {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatedSessionUserAgentDevice struct {
	value *AuthenticatedSessionUserAgentDevice
	isSet bool
}

func (v NullableAuthenticatedSessionUserAgentDevice) Get() *AuthenticatedSessionUserAgentDevice {
	return v.value
}

func (v *NullableAuthenticatedSessionUserAgentDevice) Set(val *AuthenticatedSessionUserAgentDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatedSessionUserAgentDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatedSessionUserAgentDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatedSessionUserAgentDevice(val *AuthenticatedSessionUserAgentDevice) *NullableAuthenticatedSessionUserAgentDevice {
	return &NullableAuthenticatedSessionUserAgentDevice{value: val, isSet: true}
}

func (v NullableAuthenticatedSessionUserAgentDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatedSessionUserAgentDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
