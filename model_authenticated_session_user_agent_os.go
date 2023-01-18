/*
authentik

Making authentication simple.

API version: 2023.1.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatedSessionUserAgentOs User agent os
type AuthenticatedSessionUserAgentOs struct {
	Family     string `json:"family"`
	Major      string `json:"major"`
	Minor      string `json:"minor"`
	Patch      string `json:"patch"`
	PatchMinor string `json:"patch_minor"`
}

// NewAuthenticatedSessionUserAgentOs instantiates a new AuthenticatedSessionUserAgentOs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSessionUserAgentOs(family string, major string, minor string, patch string, patchMinor string) *AuthenticatedSessionUserAgentOs {
	this := AuthenticatedSessionUserAgentOs{}
	this.Family = family
	this.Major = major
	this.Minor = minor
	this.Patch = patch
	this.PatchMinor = patchMinor
	return &this
}

// NewAuthenticatedSessionUserAgentOsWithDefaults instantiates a new AuthenticatedSessionUserAgentOs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionUserAgentOsWithDefaults() *AuthenticatedSessionUserAgentOs {
	this := AuthenticatedSessionUserAgentOs{}
	return &this
}

// GetFamily returns the Family field value
func (o *AuthenticatedSessionUserAgentOs) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentOs) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *AuthenticatedSessionUserAgentOs) SetFamily(v string) {
	o.Family = v
}

// GetMajor returns the Major field value
func (o *AuthenticatedSessionUserAgentOs) GetMajor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Major
}

// GetMajorOk returns a tuple with the Major field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentOs) GetMajorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Major, true
}

// SetMajor sets field value
func (o *AuthenticatedSessionUserAgentOs) SetMajor(v string) {
	o.Major = v
}

// GetMinor returns the Minor field value
func (o *AuthenticatedSessionUserAgentOs) GetMinor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Minor
}

// GetMinorOk returns a tuple with the Minor field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentOs) GetMinorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minor, true
}

// SetMinor sets field value
func (o *AuthenticatedSessionUserAgentOs) SetMinor(v string) {
	o.Minor = v
}

// GetPatch returns the Patch field value
func (o *AuthenticatedSessionUserAgentOs) GetPatch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Patch
}

// GetPatchOk returns a tuple with the Patch field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentOs) GetPatchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patch, true
}

// SetPatch sets field value
func (o *AuthenticatedSessionUserAgentOs) SetPatch(v string) {
	o.Patch = v
}

// GetPatchMinor returns the PatchMinor field value
func (o *AuthenticatedSessionUserAgentOs) GetPatchMinor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PatchMinor
}

// GetPatchMinorOk returns a tuple with the PatchMinor field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentOs) GetPatchMinorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatchMinor, true
}

// SetPatchMinor sets field value
func (o *AuthenticatedSessionUserAgentOs) SetPatchMinor(v string) {
	o.PatchMinor = v
}

func (o AuthenticatedSessionUserAgentOs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["family"] = o.Family
	}
	if true {
		toSerialize["major"] = o.Major
	}
	if true {
		toSerialize["minor"] = o.Minor
	}
	if true {
		toSerialize["patch"] = o.Patch
	}
	if true {
		toSerialize["patch_minor"] = o.PatchMinor
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatedSessionUserAgentOs struct {
	value *AuthenticatedSessionUserAgentOs
	isSet bool
}

func (v NullableAuthenticatedSessionUserAgentOs) Get() *AuthenticatedSessionUserAgentOs {
	return v.value
}

func (v *NullableAuthenticatedSessionUserAgentOs) Set(val *AuthenticatedSessionUserAgentOs) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatedSessionUserAgentOs) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatedSessionUserAgentOs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatedSessionUserAgentOs(val *AuthenticatedSessionUserAgentOs) *NullableAuthenticatedSessionUserAgentOs {
	return &NullableAuthenticatedSessionUserAgentOs{value: val, isSet: true}
}

func (v NullableAuthenticatedSessionUserAgentOs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatedSessionUserAgentOs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
