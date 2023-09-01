/*
authentik

Making authentication simple.

API version: 2023.8.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RadiusProvider RadiusProvider Serializer
type RadiusProvider struct {
	Pk   int32  `json:"pk"`
	Name string `json:"name"`
	// Flow used for authentication when the associated application is accessed by an un-authenticated user.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow string   `json:"authorization_flow"`
	PropertyMappings  []string `json:"property_mappings,omitempty"`
	// Get object component so that we know how to edit the object
	Component string `json:"component"`
	// Internal application name, used in URLs.
	AssignedApplicationSlug string `json:"assigned_application_slug"`
	// Application's display Name.
	AssignedApplicationName string `json:"assigned_application_name"`
	// Internal application name, used in URLs.
	AssignedBackchannelApplicationSlug string `json:"assigned_backchannel_application_slug"`
	// Application's display Name.
	AssignedBackchannelApplicationName string `json:"assigned_backchannel_application_name"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string `json:"meta_model_name"`
	// List of CIDRs (comma-separated) that clients can connect from. A more specific CIDR will match before a looser one. Clients connecting from a non-specified CIDR will be dropped.
	ClientNetworks *string `json:"client_networks,omitempty"`
	// Shared secret between clients and server to hash packets.
	SharedSecret *string  `json:"shared_secret,omitempty"`
	OutpostSet   []string `json:"outpost_set"`
}

// NewRadiusProvider instantiates a new RadiusProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadiusProvider(pk int32, name string, authorizationFlow string, component string, assignedApplicationSlug string, assignedApplicationName string, assignedBackchannelApplicationSlug string, assignedBackchannelApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, outpostSet []string) *RadiusProvider {
	this := RadiusProvider{}
	this.Pk = pk
	this.Name = name
	this.AuthorizationFlow = authorizationFlow
	this.Component = component
	this.AssignedApplicationSlug = assignedApplicationSlug
	this.AssignedApplicationName = assignedApplicationName
	this.AssignedBackchannelApplicationSlug = assignedBackchannelApplicationSlug
	this.AssignedBackchannelApplicationName = assignedBackchannelApplicationName
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.OutpostSet = outpostSet
	return &this
}

// NewRadiusProviderWithDefaults instantiates a new RadiusProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusProviderWithDefaults() *RadiusProvider {
	this := RadiusProvider{}
	return &this
}

// GetPk returns the Pk field value
func (o *RadiusProvider) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *RadiusProvider) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *RadiusProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RadiusProvider) SetName(v string) {
	o.Name = v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RadiusProvider) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RadiusProvider) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *RadiusProvider) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *RadiusProvider) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *RadiusProvider) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *RadiusProvider) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetAuthorizationFlow returns the AuthorizationFlow field value
func (o *RadiusProvider) GetAuthorizationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationFlow, true
}

// SetAuthorizationFlow sets field value
func (o *RadiusProvider) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *RadiusProvider) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *RadiusProvider) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *RadiusProvider) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetComponent returns the Component field value
func (o *RadiusProvider) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *RadiusProvider) SetComponent(v string) {
	o.Component = v
}

// GetAssignedApplicationSlug returns the AssignedApplicationSlug field value
func (o *RadiusProvider) GetAssignedApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationSlug
}

// GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetAssignedApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationSlug, true
}

// SetAssignedApplicationSlug sets field value
func (o *RadiusProvider) SetAssignedApplicationSlug(v string) {
	o.AssignedApplicationSlug = v
}

// GetAssignedApplicationName returns the AssignedApplicationName field value
func (o *RadiusProvider) GetAssignedApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationName
}

// GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetAssignedApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationName, true
}

// SetAssignedApplicationName sets field value
func (o *RadiusProvider) SetAssignedApplicationName(v string) {
	o.AssignedApplicationName = v
}

// GetAssignedBackchannelApplicationSlug returns the AssignedBackchannelApplicationSlug field value
func (o *RadiusProvider) GetAssignedBackchannelApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedBackchannelApplicationSlug
}

// GetAssignedBackchannelApplicationSlugOk returns a tuple with the AssignedBackchannelApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetAssignedBackchannelApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedBackchannelApplicationSlug, true
}

// SetAssignedBackchannelApplicationSlug sets field value
func (o *RadiusProvider) SetAssignedBackchannelApplicationSlug(v string) {
	o.AssignedBackchannelApplicationSlug = v
}

// GetAssignedBackchannelApplicationName returns the AssignedBackchannelApplicationName field value
func (o *RadiusProvider) GetAssignedBackchannelApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedBackchannelApplicationName
}

// GetAssignedBackchannelApplicationNameOk returns a tuple with the AssignedBackchannelApplicationName field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetAssignedBackchannelApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedBackchannelApplicationName, true
}

// SetAssignedBackchannelApplicationName sets field value
func (o *RadiusProvider) SetAssignedBackchannelApplicationName(v string) {
	o.AssignedBackchannelApplicationName = v
}

// GetVerboseName returns the VerboseName field value
func (o *RadiusProvider) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *RadiusProvider) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *RadiusProvider) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *RadiusProvider) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *RadiusProvider) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *RadiusProvider) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetClientNetworks returns the ClientNetworks field value if set, zero value otherwise.
func (o *RadiusProvider) GetClientNetworks() string {
	if o == nil || o.ClientNetworks == nil {
		var ret string
		return ret
	}
	return *o.ClientNetworks
}

// GetClientNetworksOk returns a tuple with the ClientNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetClientNetworksOk() (*string, bool) {
	if o == nil || o.ClientNetworks == nil {
		return nil, false
	}
	return o.ClientNetworks, true
}

// HasClientNetworks returns a boolean if a field has been set.
func (o *RadiusProvider) HasClientNetworks() bool {
	if o != nil && o.ClientNetworks != nil {
		return true
	}

	return false
}

// SetClientNetworks gets a reference to the given string and assigns it to the ClientNetworks field.
func (o *RadiusProvider) SetClientNetworks(v string) {
	o.ClientNetworks = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *RadiusProvider) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *RadiusProvider) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *RadiusProvider) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetOutpostSet returns the OutpostSet field value
func (o *RadiusProvider) GetOutpostSet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OutpostSet
}

// GetOutpostSetOk returns a tuple with the OutpostSet field value
// and a boolean to check if the value has been set.
func (o *RadiusProvider) GetOutpostSetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutpostSet, true
}

// SetOutpostSet sets field value
func (o *RadiusProvider) SetOutpostSet(v []string) {
	o.OutpostSet = v
}

func (o RadiusProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
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
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["assigned_application_slug"] = o.AssignedApplicationSlug
	}
	if true {
		toSerialize["assigned_application_name"] = o.AssignedApplicationName
	}
	if true {
		toSerialize["assigned_backchannel_application_slug"] = o.AssignedBackchannelApplicationSlug
	}
	if true {
		toSerialize["assigned_backchannel_application_name"] = o.AssignedBackchannelApplicationName
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if o.ClientNetworks != nil {
		toSerialize["client_networks"] = o.ClientNetworks
	}
	if o.SharedSecret != nil {
		toSerialize["shared_secret"] = o.SharedSecret
	}
	if true {
		toSerialize["outpost_set"] = o.OutpostSet
	}
	return json.Marshal(toSerialize)
}

type NullableRadiusProvider struct {
	value *RadiusProvider
	isSet bool
}

func (v NullableRadiusProvider) Get() *RadiusProvider {
	return v.value
}

func (v *NullableRadiusProvider) Set(val *RadiusProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableRadiusProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableRadiusProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadiusProvider(val *RadiusProvider) *NullableRadiusProvider {
	return &NullableRadiusProvider{value: val, isSet: true}
}

func (v NullableRadiusProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadiusProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
