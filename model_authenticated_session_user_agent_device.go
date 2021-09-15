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

// AuthenticatedSessionUserAgentDevice struct for AuthenticatedSessionUserAgentDevice
type AuthenticatedSessionUserAgentDevice struct {
	Brand *string `json:"brand,omitempty"`
	Family *string `json:"family,omitempty"`
	Model *string `json:"model,omitempty"`
}

// NewAuthenticatedSessionUserAgentDevice instantiates a new AuthenticatedSessionUserAgentDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSessionUserAgentDevice() *AuthenticatedSessionUserAgentDevice {
	this := AuthenticatedSessionUserAgentDevice{}
	return &this
}

// NewAuthenticatedSessionUserAgentDeviceWithDefaults instantiates a new AuthenticatedSessionUserAgentDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionUserAgentDeviceWithDefaults() *AuthenticatedSessionUserAgentDevice {
	this := AuthenticatedSessionUserAgentDevice{}
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentDevice) GetBrand() string {
	if o == nil || o.Brand == nil {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentDevice) GetBrandOk() (*string, bool) {
	if o == nil || o.Brand == nil {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentDevice) HasBrand() bool {
	if o != nil && o.Brand != nil {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *AuthenticatedSessionUserAgentDevice) SetBrand(v string) {
	o.Brand = &v
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentDevice) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentDevice) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentDevice) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *AuthenticatedSessionUserAgentDevice) SetFamily(v string) {
	o.Family = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AuthenticatedSessionUserAgentDevice) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentDevice) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AuthenticatedSessionUserAgentDevice) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AuthenticatedSessionUserAgentDevice) SetModel(v string) {
	o.Model = &v
}

func (o AuthenticatedSessionUserAgentDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Brand != nil {
		toSerialize["brand"] = o.Brand
	}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Model != nil {
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


