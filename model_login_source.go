/*
authentik

Making authentication simple.

API version: 2023.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LoginSource Serializer for Login buttons of sources
type LoginSource struct {
	Name      string              `json:"name"`
	IconUrl   NullableString      `json:"icon_url,omitempty"`
	Challenge LoginChallengeTypes `json:"challenge"`
}

// NewLoginSource instantiates a new LoginSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginSource(name string, challenge LoginChallengeTypes) *LoginSource {
	this := LoginSource{}
	this.Name = name
	this.Challenge = challenge
	return &this
}

// NewLoginSourceWithDefaults instantiates a new LoginSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginSourceWithDefaults() *LoginSource {
	this := LoginSource{}
	return &this
}

// GetName returns the Name field value
func (o *LoginSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoginSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoginSource) SetName(v string) {
	o.Name = v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginSource) GetIconUrl() string {
	if o == nil || o.IconUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginSource) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *LoginSource) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *LoginSource) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}

// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *LoginSource) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *LoginSource) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetChallenge returns the Challenge field value
func (o *LoginSource) GetChallenge() LoginChallengeTypes {
	if o == nil {
		var ret LoginChallengeTypes
		return ret
	}

	return o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value
// and a boolean to check if the value has been set.
func (o *LoginSource) GetChallengeOk() (*LoginChallengeTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Challenge, true
}

// SetChallenge sets field value
func (o *LoginSource) SetChallenge(v LoginChallengeTypes) {
	o.Challenge = v
}

func (o LoginSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	if true {
		toSerialize["challenge"] = o.Challenge
	}
	return json.Marshal(toSerialize)
}

type NullableLoginSource struct {
	value *LoginSource
	isSet bool
}

func (v NullableLoginSource) Get() *LoginSource {
	return v.value
}

func (v *NullableLoginSource) Set(val *LoginSource) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginSource) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginSource(val *LoginSource) *NullableLoginSource {
	return &NullableLoginSource{value: val, isSet: true}
}

func (v NullableLoginSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
