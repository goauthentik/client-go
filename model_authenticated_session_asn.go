/*
authentik

Making authentication simple.

API version: 2023.10.6
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthenticatedSessionAsn Get ASN Data
type AuthenticatedSessionAsn struct {
	Asn     int32          `json:"asn"`
	AsOrg   NullableString `json:"as_org"`
	Network NullableString `json:"network"`
}

// NewAuthenticatedSessionAsn instantiates a new AuthenticatedSessionAsn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatedSessionAsn(asn int32, asOrg NullableString, network NullableString) *AuthenticatedSessionAsn {
	this := AuthenticatedSessionAsn{}
	this.Asn = asn
	this.AsOrg = asOrg
	this.Network = network
	return &this
}

// NewAuthenticatedSessionAsnWithDefaults instantiates a new AuthenticatedSessionAsn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatedSessionAsnWithDefaults() *AuthenticatedSessionAsn {
	this := AuthenticatedSessionAsn{}
	return &this
}

// GetAsn returns the Asn field value
func (o *AuthenticatedSessionAsn) GetAsn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Asn
}

// GetAsnOk returns a tuple with the Asn field value
// and a boolean to check if the value has been set.
func (o *AuthenticatedSessionAsn) GetAsnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asn, true
}

// SetAsn sets field value
func (o *AuthenticatedSessionAsn) SetAsn(v int32) {
	o.Asn = v
}

// GetAsOrg returns the AsOrg field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuthenticatedSessionAsn) GetAsOrg() string {
	if o == nil || o.AsOrg.Get() == nil {
		var ret string
		return ret
	}

	return *o.AsOrg.Get()
}

// GetAsOrgOk returns a tuple with the AsOrg field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatedSessionAsn) GetAsOrgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AsOrg.Get(), o.AsOrg.IsSet()
}

// SetAsOrg sets field value
func (o *AuthenticatedSessionAsn) SetAsOrg(v string) {
	o.AsOrg.Set(&v)
}

// GetNetwork returns the Network field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuthenticatedSessionAsn) GetNetwork() string {
	if o == nil || o.Network.Get() == nil {
		var ret string
		return ret
	}

	return *o.Network.Get()
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticatedSessionAsn) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Network.Get(), o.Network.IsSet()
}

// SetNetwork sets field value
func (o *AuthenticatedSessionAsn) SetNetwork(v string) {
	o.Network.Set(&v)
}

func (o AuthenticatedSessionAsn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asn"] = o.Asn
	}
	if true {
		toSerialize["as_org"] = o.AsOrg.Get()
	}
	if true {
		toSerialize["network"] = o.Network.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticatedSessionAsn struct {
	value *AuthenticatedSessionAsn
	isSet bool
}

func (v NullableAuthenticatedSessionAsn) Get() *AuthenticatedSessionAsn {
	return v.value
}

func (v *NullableAuthenticatedSessionAsn) Set(val *AuthenticatedSessionAsn) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatedSessionAsn) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatedSessionAsn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatedSessionAsn(val *AuthenticatedSessionAsn) *NullableAuthenticatedSessionAsn {
	return &NullableAuthenticatedSessionAsn{value: val, isSet: true}
}

func (v NullableAuthenticatedSessionAsn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatedSessionAsn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
