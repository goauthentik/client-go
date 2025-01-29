/*
authentik

Making authentication simple.

API version: 2024.12.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatedSessionUserAgentUserAgent User agent browser
type AuthenticatedSessionUserAgentUserAgent struct {
	Family string `json:"family"`
	Major  string `json:"major"`
	Minor  string `json:"minor"`
	Patch  string `json:"patch"`
}

// NewAuthenticatedSessionUserAgentUserAgent instantiates a new AuthenticatedSessionUserAgentUserAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSessionUserAgentUserAgent(family string, major string, minor string, patch string) *AuthenticatedSessionUserAgentUserAgent {
	this := AuthenticatedSessionUserAgentUserAgent{}
	this.Family = family
	this.Major = major
	this.Minor = minor
	this.Patch = patch
	return &this
}

// NewAuthenticatedSessionUserAgentUserAgentWithDefaults instantiates a new AuthenticatedSessionUserAgentUserAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionUserAgentUserAgentWithDefaults() *AuthenticatedSessionUserAgentUserAgent {
	this := AuthenticatedSessionUserAgentUserAgent{}
	return &this
}

// GetFamily returns the Family field value
func (o *AuthenticatedSessionUserAgentUserAgent) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *AuthenticatedSessionUserAgentUserAgent) SetFamily(v string) {
	o.Family = v
}

// GetMajor returns the Major field value
func (o *AuthenticatedSessionUserAgentUserAgent) GetMajor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Major
}

// GetMajorOk returns a tuple with the Major field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) GetMajorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Major, true
}

// SetMajor sets field value
func (o *AuthenticatedSessionUserAgentUserAgent) SetMajor(v string) {
	o.Major = v
}

// GetMinor returns the Minor field value
func (o *AuthenticatedSessionUserAgentUserAgent) GetMinor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Minor
}

// GetMinorOk returns a tuple with the Minor field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) GetMinorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minor, true
}

// SetMinor sets field value
func (o *AuthenticatedSessionUserAgentUserAgent) SetMinor(v string) {
	o.Minor = v
}

// GetPatch returns the Patch field value
func (o *AuthenticatedSessionUserAgentUserAgent) GetPatch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Patch
}

// GetPatchOk returns a tuple with the Patch field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionUserAgentUserAgent) GetPatchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patch, true
}

// SetPatch sets field value
func (o *AuthenticatedSessionUserAgentUserAgent) SetPatch(v string) {
	o.Patch = v
}

func (o AuthenticatedSessionUserAgentUserAgent) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableAuthenticatedSessionUserAgentUserAgent struct {
	value *AuthenticatedSessionUserAgentUserAgent
	isSet bool
}

func (v NullableAuthenticatedSessionUserAgentUserAgent) Get() *AuthenticatedSessionUserAgentUserAgent {
	return v.value
}

func (v *NullableAuthenticatedSessionUserAgentUserAgent) Set(val *AuthenticatedSessionUserAgentUserAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatedSessionUserAgentUserAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatedSessionUserAgentUserAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatedSessionUserAgentUserAgent(val *AuthenticatedSessionUserAgentUserAgent) *NullableAuthenticatedSessionUserAgentUserAgent {
	return &NullableAuthenticatedSessionUserAgentUserAgent{value: val, isSet: true}
}

func (v NullableAuthenticatedSessionUserAgentUserAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatedSessionUserAgentUserAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
