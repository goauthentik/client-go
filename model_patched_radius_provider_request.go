/*
authentik

Making authentication simple.

API version: 2023.3.1
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
	// Flow used when authorizing this provider.
	AuthorizationFlow *string  `json:"authorization_flow,omitempty"`
	PropertyMappings  []string `json:"property_mappings,omitempty"`
	// List of CIDRs (comma-seperated) that clients can connect from. A more specific CIDR will match before a looser one. Clients connecting from a non-specified CIDR will be dropped.
	ClientNetworks *string `json:"client_networks,omitempty"`
	// Shared secret between clients and server to hash packets.
	SharedSecret *string `json:"shared_secret,omitempty"`
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

func (o PatchedRadiusProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
