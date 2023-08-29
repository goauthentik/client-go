/*
authentik

Making authentication simple.

API version: 2023.8.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RadiusProviderRequest RadiusProvider Serializer
type RadiusProviderRequest struct {
	Name string `json:"name"`
	// Flow used for authentication when the associated application is accessed by an un-authenticated user.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow string   `json:"authorization_flow"`
	PropertyMappings  []string `json:"property_mappings,omitempty"`
	// List of CIDRs (comma-separated) that clients can connect from. A more specific CIDR will match before a looser one. Clients connecting from a non-specified CIDR will be dropped.
	ClientNetworks *string `json:"client_networks,omitempty"`
	// Shared secret between clients and server to hash packets.
	SharedSecret *string `json:"shared_secret,omitempty"`
}

// NewRadiusProviderRequest instantiates a new RadiusProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadiusProviderRequest(name string, authorizationFlow string) *RadiusProviderRequest {
	this := RadiusProviderRequest{}
	this.Name = name
	this.AuthorizationFlow = authorizationFlow
	return &this
}

// NewRadiusProviderRequestWithDefaults instantiates a new RadiusProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusProviderRequestWithDefaults() *RadiusProviderRequest {
	this := RadiusProviderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *RadiusProviderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RadiusProviderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RadiusProviderRequest) SetName(v string) {
	o.Name = v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RadiusProviderRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RadiusProviderRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *RadiusProviderRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *RadiusProviderRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *RadiusProviderRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *RadiusProviderRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetAuthorizationFlow returns the AuthorizationFlow field value
func (o *RadiusProviderRequest) GetAuthorizationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value
// and a boolean to check if the value has been set.
func (o *RadiusProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationFlow, true
}

// SetAuthorizationFlow sets field value
func (o *RadiusProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *RadiusProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *RadiusProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *RadiusProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetClientNetworks returns the ClientNetworks field value if set, zero value otherwise.
func (o *RadiusProviderRequest) GetClientNetworks() string {
	if o == nil || o.ClientNetworks == nil {
		var ret string
		return ret
	}
	return *o.ClientNetworks
}

// GetClientNetworksOk returns a tuple with the ClientNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusProviderRequest) GetClientNetworksOk() (*string, bool) {
	if o == nil || o.ClientNetworks == nil {
		return nil, false
	}
	return o.ClientNetworks, true
}

// HasClientNetworks returns a boolean if a field has been set.
func (o *RadiusProviderRequest) HasClientNetworks() bool {
	if o != nil && o.ClientNetworks != nil {
		return true
	}

	return false
}

// SetClientNetworks gets a reference to the given string and assigns it to the ClientNetworks field.
func (o *RadiusProviderRequest) SetClientNetworks(v string) {
	o.ClientNetworks = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *RadiusProviderRequest) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusProviderRequest) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *RadiusProviderRequest) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *RadiusProviderRequest) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

func (o RadiusProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.AuthenticationFlow.IsSet() {
		toSerialize["authentication_flow"] = o.AuthenticationFlow.Get()
	}
	if true {
		toSerialize["authorization_flow"] = o.AuthorizationFlow
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if o.ClientNetworks != nil {
		toSerialize["client_networks"] = o.ClientNetworks
	}
	if o.SharedSecret != nil {
		toSerialize["shared_secret"] = o.SharedSecret
	}
	return json.Marshal(toSerialize)
}

type NullableRadiusProviderRequest struct {
	value *RadiusProviderRequest
	isSet bool
}

func (v NullableRadiusProviderRequest) Get() *RadiusProviderRequest {
	return v.value
}

func (v *NullableRadiusProviderRequest) Set(val *RadiusProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRadiusProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRadiusProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadiusProviderRequest(val *RadiusProviderRequest) *NullableRadiusProviderRequest {
	return &NullableRadiusProviderRequest{value: val, isSet: true}
}

func (v NullableRadiusProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadiusProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
