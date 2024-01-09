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

// PatchedRadiusProviderRequest RadiusProvider Serializer
type PatchedRadiusProviderRequest struct {
	Name *string `json:"name,omitempty"`
	// Flow used for authentication when the associated application is accessed by an un-authenticated user.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow *string  `json:"authorization_flow,omitempty"`
	PropertyMappings  []string `json:"property_mappings,omitempty"`
	// List of CIDRs (comma-separated) that clients can connect from. A more specific CIDR will match before a looser one. Clients connecting from a non-specified CIDR will be dropped.
	ClientNetworks *string `json:"client_networks,omitempty"`
	// Shared secret between clients and server to hash packets.
	SharedSecret *string `json:"shared_secret,omitempty"`
	// When enabled, code-based multi-factor authentication can be used by appending a semicolon and the TOTP code to the password. This should only be enabled if all users that will bind to this provider have a TOTP device configured, as otherwise a password may incorrectly be rejected if it contains a semicolon.
	MfaSupport *bool `json:"mfa_support,omitempty"`
}

// NewPatchedRadiusProviderRequest instantiates a new PatchedRadiusProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRadiusProviderRequest() *PatchedRadiusProviderRequest {
	this := PatchedRadiusProviderRequest{}
	return &this
}

// NewPatchedRadiusProviderRequestWithDefaults instantiates a new PatchedRadiusProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRadiusProviderRequestWithDefaults() *PatchedRadiusProviderRequest {
	this := PatchedRadiusProviderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedRadiusProviderRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRadiusProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedRadiusProviderRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedRadiusProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRadiusProviderRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRadiusProviderRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *PatchedRadiusProviderRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *PatchedRadiusProviderRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *PatchedRadiusProviderRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *PatchedRadiusProviderRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetAuthorizationFlow returns the AuthorizationFlow field value if set, zero value otherwise.
func (o *PatchedRadiusProviderRequest) GetAuthorizationFlow() string {
	if o == nil || o.AuthorizationFlow == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRadiusProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil || o.AuthorizationFlow == nil {
		return nil, false
	}
	return o.AuthorizationFlow, true
}

// HasAuthorizationFlow returns a boolean if a field has been set.
func (o *PatchedRadiusProviderRequest) HasAuthorizationFlow() bool {
	if o != nil && o.AuthorizationFlow != nil {
		return true
	}

	return false
}

// SetAuthorizationFlow gets a reference to the given string and assigns it to the AuthorizationFlow field.
func (o *PatchedRadiusProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = &v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *PatchedRadiusProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRadiusProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedRadiusProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *PatchedRadiusProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetClientNetworks returns the ClientNetworks field value if set, zero value otherwise.
func (o *PatchedRadiusProviderRequest) GetClientNetworks() string {
	if o == nil || o.ClientNetworks == nil {
		var ret string
		return ret
	}
	return *o.ClientNetworks
}

// GetClientNetworksOk returns a tuple with the ClientNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRadiusProviderRequest) GetClientNetworksOk() (*string, bool) {
	if o == nil || o.ClientNetworks == nil {
		return nil, false
	}
	return o.ClientNetworks, true
}

// HasClientNetworks returns a boolean if a field has been set.
func (o *PatchedRadiusProviderRequest) HasClientNetworks() bool {
	if o != nil && o.ClientNetworks != nil {
		return true
	}

	return false
}

// SetClientNetworks gets a reference to the given string and assigns it to the ClientNetworks field.
func (o *PatchedRadiusProviderRequest) SetClientNetworks(v string) {
	o.ClientNetworks = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *PatchedRadiusProviderRequest) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRadiusProviderRequest) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *PatchedRadiusProviderRequest) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *PatchedRadiusProviderRequest) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetMfaSupport returns the MfaSupport field value if set, zero value otherwise.
func (o *PatchedRadiusProviderRequest) GetMfaSupport() bool {
	if o == nil || o.MfaSupport == nil {
		var ret bool
		return ret
	}
	return *o.MfaSupport
}

// GetMfaSupportOk returns a tuple with the MfaSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRadiusProviderRequest) GetMfaSupportOk() (*bool, bool) {
	if o == nil || o.MfaSupport == nil {
		return nil, false
	}
	return o.MfaSupport, true
}

// HasMfaSupport returns a boolean if a field has been set.
func (o *PatchedRadiusProviderRequest) HasMfaSupport() bool {
	if o != nil && o.MfaSupport != nil {
		return true
	}

	return false
}

// SetMfaSupport gets a reference to the given bool and assigns it to the MfaSupport field.
func (o *PatchedRadiusProviderRequest) SetMfaSupport(v bool) {
	o.MfaSupport = &v
}

func (o PatchedRadiusProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AuthenticationFlow.IsSet() {
		toSerialize["authentication_flow"] = o.AuthenticationFlow.Get()
	}
	if o.AuthorizationFlow != nil {
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
	if o.MfaSupport != nil {
		toSerialize["mfa_support"] = o.MfaSupport
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedRadiusProviderRequest struct {
	value *PatchedRadiusProviderRequest
	isSet bool
}

func (v NullablePatchedRadiusProviderRequest) Get() *PatchedRadiusProviderRequest {
	return v.value
}

func (v *NullablePatchedRadiusProviderRequest) Set(val *PatchedRadiusProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRadiusProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRadiusProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRadiusProviderRequest(val *PatchedRadiusProviderRequest) *NullablePatchedRadiusProviderRequest {
	return &NullablePatchedRadiusProviderRequest{value: val, isSet: true}
}

func (v NullablePatchedRadiusProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRadiusProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
