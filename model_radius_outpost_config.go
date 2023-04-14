/*
authentik

Making authentication simple.

API version: 2023.4.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RadiusOutpostConfig RadiusProvider Serializer
type RadiusOutpostConfig struct {
	Pk              int32  `json:"pk"`
	Name            string `json:"name"`
	ApplicationSlug string `json:"application_slug"`
	AuthFlowSlug    string `json:"auth_flow_slug"`
	// List of CIDRs (comma-separated) that clients can connect from. A more specific CIDR will match before a looser one. Clients connecting from a non-specified CIDR will be dropped.
	ClientNetworks *string `json:"client_networks,omitempty"`
	// Shared secret between clients and server to hash packets.
	SharedSecret *string `json:"shared_secret,omitempty"`
}

// NewRadiusOutpostConfig instantiates a new RadiusOutpostConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadiusOutpostConfig(pk int32, name string, applicationSlug string, authFlowSlug string) *RadiusOutpostConfig {
	this := RadiusOutpostConfig{}
	this.Pk = pk
	this.Name = name
	this.ApplicationSlug = applicationSlug
	this.AuthFlowSlug = authFlowSlug
	return &this
}

// NewRadiusOutpostConfigWithDefaults instantiates a new RadiusOutpostConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusOutpostConfigWithDefaults() *RadiusOutpostConfig {
	this := RadiusOutpostConfig{}
	return &this
}

// GetPk returns the Pk field value
func (o *RadiusOutpostConfig) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *RadiusOutpostConfig) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *RadiusOutpostConfig) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *RadiusOutpostConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RadiusOutpostConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RadiusOutpostConfig) SetName(v string) {
	o.Name = v
}

// GetApplicationSlug returns the ApplicationSlug field value
func (o *RadiusOutpostConfig) GetApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationSlug
}

// GetApplicationSlugOk returns a tuple with the ApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *RadiusOutpostConfig) GetApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationSlug, true
}

// SetApplicationSlug sets field value
func (o *RadiusOutpostConfig) SetApplicationSlug(v string) {
	o.ApplicationSlug = v
}

// GetAuthFlowSlug returns the AuthFlowSlug field value
func (o *RadiusOutpostConfig) GetAuthFlowSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthFlowSlug
}

// GetAuthFlowSlugOk returns a tuple with the AuthFlowSlug field value
// and a boolean to check if the value has been set.
func (o *RadiusOutpostConfig) GetAuthFlowSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthFlowSlug, true
}

// SetAuthFlowSlug sets field value
func (o *RadiusOutpostConfig) SetAuthFlowSlug(v string) {
	o.AuthFlowSlug = v
}

// GetClientNetworks returns the ClientNetworks field value if set, zero value otherwise.
func (o *RadiusOutpostConfig) GetClientNetworks() string {
	if o == nil || o.ClientNetworks == nil {
		var ret string
		return ret
	}
	return *o.ClientNetworks
}

// GetClientNetworksOk returns a tuple with the ClientNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusOutpostConfig) GetClientNetworksOk() (*string, bool) {
	if o == nil || o.ClientNetworks == nil {
		return nil, false
	}
	return o.ClientNetworks, true
}

// HasClientNetworks returns a boolean if a field has been set.
func (o *RadiusOutpostConfig) HasClientNetworks() bool {
	if o != nil && o.ClientNetworks != nil {
		return true
	}

	return false
}

// SetClientNetworks gets a reference to the given string and assigns it to the ClientNetworks field.
func (o *RadiusOutpostConfig) SetClientNetworks(v string) {
	o.ClientNetworks = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *RadiusOutpostConfig) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusOutpostConfig) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *RadiusOutpostConfig) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *RadiusOutpostConfig) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

func (o RadiusOutpostConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["application_slug"] = o.ApplicationSlug
	}
	if true {
		toSerialize["auth_flow_slug"] = o.AuthFlowSlug
	}
	if o.ClientNetworks != nil {
		toSerialize["client_networks"] = o.ClientNetworks
	}
	if o.SharedSecret != nil {
		toSerialize["shared_secret"] = o.SharedSecret
	}
	return json.Marshal(toSerialize)
}

type NullableRadiusOutpostConfig struct {
	value *RadiusOutpostConfig
	isSet bool
}

func (v NullableRadiusOutpostConfig) Get() *RadiusOutpostConfig {
	return v.value
}

func (v *NullableRadiusOutpostConfig) Set(val *RadiusOutpostConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRadiusOutpostConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRadiusOutpostConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadiusOutpostConfig(val *RadiusOutpostConfig) *NullableRadiusOutpostConfig {
	return &NullableRadiusOutpostConfig{value: val, isSet: true}
}

func (v NullableRadiusOutpostConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadiusOutpostConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
