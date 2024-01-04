/*
authentik

Making authentication simple.

API version: 2023.10.5
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RACProvider RACProvider Serializer
type RACProvider struct {
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
	MetaModelName string      `json:"meta_model_name"`
	Settings      interface{} `json:"settings,omitempty"`
	OutpostSet    []string    `json:"outpost_set"`
	// Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours=-1;minutes=-2;seconds=-3)
	ConnectionExpiry *string `json:"connection_expiry,omitempty"`
}

// NewRACProvider instantiates a new RACProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRACProvider(pk int32, name string, authorizationFlow string, component string, assignedApplicationSlug string, assignedApplicationName string, assignedBackchannelApplicationSlug string, assignedBackchannelApplicationName string, verboseName string, verboseNamePlural string, metaModelName string, outpostSet []string) *RACProvider {
	this := RACProvider{}
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

// NewRACProviderWithDefaults instantiates a new RACProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRACProviderWithDefaults() *RACProvider {
	this := RACProvider{}
	return &this
}

// GetPk returns the Pk field value
func (o *RACProvider) GetPk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetPkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *RACProvider) SetPk(v int32) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *RACProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RACProvider) SetName(v string) {
	o.Name = v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RACProvider) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RACProvider) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *RACProvider) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *RACProvider) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *RACProvider) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *RACProvider) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetAuthorizationFlow returns the AuthorizationFlow field value
func (o *RACProvider) GetAuthorizationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationFlow, true
}

// SetAuthorizationFlow sets field value
func (o *RACProvider) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *RACProvider) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RACProvider) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *RACProvider) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *RACProvider) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetComponent returns the Component field value
func (o *RACProvider) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *RACProvider) SetComponent(v string) {
	o.Component = v
}

// GetAssignedApplicationSlug returns the AssignedApplicationSlug field value
func (o *RACProvider) GetAssignedApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationSlug
}

// GetAssignedApplicationSlugOk returns a tuple with the AssignedApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetAssignedApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationSlug, true
}

// SetAssignedApplicationSlug sets field value
func (o *RACProvider) SetAssignedApplicationSlug(v string) {
	o.AssignedApplicationSlug = v
}

// GetAssignedApplicationName returns the AssignedApplicationName field value
func (o *RACProvider) GetAssignedApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedApplicationName
}

// GetAssignedApplicationNameOk returns a tuple with the AssignedApplicationName field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetAssignedApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedApplicationName, true
}

// SetAssignedApplicationName sets field value
func (o *RACProvider) SetAssignedApplicationName(v string) {
	o.AssignedApplicationName = v
}

// GetAssignedBackchannelApplicationSlug returns the AssignedBackchannelApplicationSlug field value
func (o *RACProvider) GetAssignedBackchannelApplicationSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedBackchannelApplicationSlug
}

// GetAssignedBackchannelApplicationSlugOk returns a tuple with the AssignedBackchannelApplicationSlug field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetAssignedBackchannelApplicationSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedBackchannelApplicationSlug, true
}

// SetAssignedBackchannelApplicationSlug sets field value
func (o *RACProvider) SetAssignedBackchannelApplicationSlug(v string) {
	o.AssignedBackchannelApplicationSlug = v
}

// GetAssignedBackchannelApplicationName returns the AssignedBackchannelApplicationName field value
func (o *RACProvider) GetAssignedBackchannelApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedBackchannelApplicationName
}

// GetAssignedBackchannelApplicationNameOk returns a tuple with the AssignedBackchannelApplicationName field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetAssignedBackchannelApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedBackchannelApplicationName, true
}

// SetAssignedBackchannelApplicationName sets field value
func (o *RACProvider) SetAssignedBackchannelApplicationName(v string) {
	o.AssignedBackchannelApplicationName = v
}

// GetVerboseName returns the VerboseName field value
func (o *RACProvider) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *RACProvider) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *RACProvider) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *RACProvider) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *RACProvider) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *RACProvider) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RACProvider) GetSettings() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RACProvider) GetSettingsOk() (*interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *RACProvider) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given interface{} and assigns it to the Settings field.
func (o *RACProvider) SetSettings(v interface{}) {
	o.Settings = v
}

// GetOutpostSet returns the OutpostSet field value
func (o *RACProvider) GetOutpostSet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OutpostSet
}

// GetOutpostSetOk returns a tuple with the OutpostSet field value
// and a boolean to check if the value has been set.
func (o *RACProvider) GetOutpostSetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutpostSet, true
}

// SetOutpostSet sets field value
func (o *RACProvider) SetOutpostSet(v []string) {
	o.OutpostSet = v
}

// GetConnectionExpiry returns the ConnectionExpiry field value if set, zero value otherwise.
func (o *RACProvider) GetConnectionExpiry() string {
	if o == nil || o.ConnectionExpiry == nil {
		var ret string
		return ret
	}
	return *o.ConnectionExpiry
}

// GetConnectionExpiryOk returns a tuple with the ConnectionExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RACProvider) GetConnectionExpiryOk() (*string, bool) {
	if o == nil || o.ConnectionExpiry == nil {
		return nil, false
	}
	return o.ConnectionExpiry, true
}

// HasConnectionExpiry returns a boolean if a field has been set.
func (o *RACProvider) HasConnectionExpiry() bool {
	if o != nil && o.ConnectionExpiry != nil {
		return true
	}

	return false
}

// SetConnectionExpiry gets a reference to the given string and assigns it to the ConnectionExpiry field.
func (o *RACProvider) SetConnectionExpiry(v string) {
	o.ConnectionExpiry = &v
}

func (o RACProvider) MarshalJSON() ([]byte, error) {
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
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if true {
		toSerialize["outpost_set"] = o.OutpostSet
	}
	if o.ConnectionExpiry != nil {
		toSerialize["connection_expiry"] = o.ConnectionExpiry
	}
	return json.Marshal(toSerialize)
}

type NullableRACProvider struct {
	value *RACProvider
	isSet bool
}

func (v NullableRACProvider) Get() *RACProvider {
	return v.value
}

func (v *NullableRACProvider) Set(val *RACProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableRACProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableRACProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRACProvider(val *RACProvider) *NullableRACProvider {
	return &NullableRACProvider{value: val, isSet: true}
}

func (v NullableRACProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRACProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}