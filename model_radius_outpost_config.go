/*
authentik

Making authentication simple.

API version: 2024.10.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the RadiusOutpostConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RadiusOutpostConfig{}

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
	// When enabled, code-based multi-factor authentication can be used by appending a semicolon and the TOTP code to the password. This should only be enabled if all users that will bind to this provider have a TOTP device configured, as otherwise a password may incorrectly be rejected if it contains a semicolon.
	MfaSupport *bool `json:"mfa_support,omitempty"`
}

type _RadiusOutpostConfig RadiusOutpostConfig

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
	if o == nil || IsNil(o.ClientNetworks) {
		var ret string
		return ret
	}
	return *o.ClientNetworks
}

// GetClientNetworksOk returns a tuple with the ClientNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusOutpostConfig) GetClientNetworksOk() (*string, bool) {
	if o == nil || IsNil(o.ClientNetworks) {
		return nil, false
	}
	return o.ClientNetworks, true
}

// HasClientNetworks returns a boolean if a field has been set.
func (o *RadiusOutpostConfig) HasClientNetworks() bool {
	if o != nil && !IsNil(o.ClientNetworks) {
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
	if o == nil || IsNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusOutpostConfig) GetSharedSecretOk() (*string, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *RadiusOutpostConfig) HasSharedSecret() bool {
	if o != nil && !IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *RadiusOutpostConfig) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetMfaSupport returns the MfaSupport field value if set, zero value otherwise.
func (o *RadiusOutpostConfig) GetMfaSupport() bool {
	if o == nil || IsNil(o.MfaSupport) {
		var ret bool
		return ret
	}
	return *o.MfaSupport
}

// GetMfaSupportOk returns a tuple with the MfaSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusOutpostConfig) GetMfaSupportOk() (*bool, bool) {
	if o == nil || IsNil(o.MfaSupport) {
		return nil, false
	}
	return o.MfaSupport, true
}

// HasMfaSupport returns a boolean if a field has been set.
func (o *RadiusOutpostConfig) HasMfaSupport() bool {
	if o != nil && !IsNil(o.MfaSupport) {
		return true
	}

	return false
}

// SetMfaSupport gets a reference to the given bool and assigns it to the MfaSupport field.
func (o *RadiusOutpostConfig) SetMfaSupport(v bool) {
	o.MfaSupport = &v
}

func (o RadiusOutpostConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RadiusOutpostConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pk"] = o.Pk
	toSerialize["name"] = o.Name
	toSerialize["application_slug"] = o.ApplicationSlug
	toSerialize["auth_flow_slug"] = o.AuthFlowSlug
	if !IsNil(o.ClientNetworks) {
		toSerialize["client_networks"] = o.ClientNetworks
	}
	if !IsNil(o.SharedSecret) {
		toSerialize["shared_secret"] = o.SharedSecret
	}
	if !IsNil(o.MfaSupport) {
		toSerialize["mfa_support"] = o.MfaSupport
	}
	return toSerialize, nil
}

func (o *RadiusOutpostConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pk",
		"name",
		"application_slug",
		"auth_flow_slug",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRadiusOutpostConfig := _RadiusOutpostConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRadiusOutpostConfig)

	if err != nil {
		return err
	}

	*o = RadiusOutpostConfig(varRadiusOutpostConfig)

	return err
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
