/*
authentik

Making authentication simple.

API version: 2024.4.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Endpoint Endpoint Serializer
type Endpoint struct {
	Pk               string       `json:"pk"`
	Name             string       `json:"name"`
	Provider         int32        `json:"provider"`
	ProviderObj      RACProvider  `json:"provider_obj"`
	Protocol         ProtocolEnum `json:"protocol"`
	Host             string       `json:"host"`
	Settings         interface{}  `json:"settings,omitempty"`
	PropertyMappings []string     `json:"property_mappings,omitempty"`
	AuthMode         AuthModeEnum `json:"auth_mode"`
	// Build actual launch URL (the provider itself does not have one, just individual endpoints)
	LaunchUrl          NullableString `json:"launch_url"`
	MaximumConnections *int32         `json:"maximum_connections,omitempty"`
}

// NewEndpoint instantiates a new Endpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpoint(pk string, name string, provider int32, providerObj RACProvider, protocol ProtocolEnum, host string, authMode AuthModeEnum, launchUrl NullableString) *Endpoint {
	this := Endpoint{}
	this.Pk = pk
	this.Name = name
	this.Provider = provider
	this.ProviderObj = providerObj
	this.Protocol = protocol
	this.Host = host
	this.AuthMode = authMode
	this.LaunchUrl = launchUrl
	return &this
}

// NewEndpointWithDefaults instantiates a new Endpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointWithDefaults() *Endpoint {
	this := Endpoint{}
	return &this
}

// GetPk returns the Pk field value
func (o *Endpoint) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *Endpoint) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *Endpoint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Endpoint) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value
func (o *Endpoint) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *Endpoint) SetProvider(v int32) {
	o.Provider = v
}

// GetProviderObj returns the ProviderObj field value
func (o *Endpoint) GetProviderObj() RACProvider {
	if o == nil {
		var ret RACProvider
		return ret
	}

	return o.ProviderObj
}

// GetProviderObjOk returns a tuple with the ProviderObj field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetProviderObjOk() (*RACProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderObj, true
}

// SetProviderObj sets field value
func (o *Endpoint) SetProviderObj(v RACProvider) {
	o.ProviderObj = v
}

// GetProtocol returns the Protocol field value
func (o *Endpoint) GetProtocol() ProtocolEnum {
	if o == nil {
		var ret ProtocolEnum
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetProtocolOk() (*ProtocolEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *Endpoint) SetProtocol(v ProtocolEnum) {
	o.Protocol = v
}

// GetHost returns the Host field value
func (o *Endpoint) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *Endpoint) SetHost(v string) {
	o.Host = v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Endpoint) GetSettings() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Endpoint) GetSettingsOk() (*interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Endpoint) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given interface{} and assigns it to the Settings field.
func (o *Endpoint) SetSettings(v interface{}) {
	o.Settings = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *Endpoint) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Endpoint) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *Endpoint) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *Endpoint) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetAuthMode returns the AuthMode field value
func (o *Endpoint) GetAuthMode() AuthModeEnum {
	if o == nil {
		var ret AuthModeEnum
		return ret
	}

	return o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetAuthModeOk() (*AuthModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthMode, true
}

// SetAuthMode sets field value
func (o *Endpoint) SetAuthMode(v AuthModeEnum) {
	o.AuthMode = v
}

// GetLaunchUrl returns the LaunchUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Endpoint) GetLaunchUrl() string {
	if o == nil || o.LaunchUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.LaunchUrl.Get()
}

// GetLaunchUrlOk returns a tuple with the LaunchUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Endpoint) GetLaunchUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LaunchUrl.Get(), o.LaunchUrl.IsSet()
}

// SetLaunchUrl sets field value
func (o *Endpoint) SetLaunchUrl(v string) {
	o.LaunchUrl.Set(&v)
}

// GetMaximumConnections returns the MaximumConnections field value if set, zero value otherwise.
func (o *Endpoint) GetMaximumConnections() int32 {
	if o == nil || o.MaximumConnections == nil {
		var ret int32
		return ret
	}
	return *o.MaximumConnections
}

// GetMaximumConnectionsOk returns a tuple with the MaximumConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Endpoint) GetMaximumConnectionsOk() (*int32, bool) {
	if o == nil || o.MaximumConnections == nil {
		return nil, false
	}
	return o.MaximumConnections, true
}

// HasMaximumConnections returns a boolean if a field has been set.
func (o *Endpoint) HasMaximumConnections() bool {
	if o != nil && o.MaximumConnections != nil {
		return true
	}

	return false
}

// SetMaximumConnections gets a reference to the given int32 and assigns it to the MaximumConnections field.
func (o *Endpoint) SetMaximumConnections(v int32) {
	o.MaximumConnections = &v
}

func (o Endpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["provider_obj"] = o.ProviderObj
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if true {
		toSerialize["host"] = o.Host
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if true {
		toSerialize["auth_mode"] = o.AuthMode
	}
	if true {
		toSerialize["launch_url"] = o.LaunchUrl.Get()
	}
	if o.MaximumConnections != nil {
		toSerialize["maximum_connections"] = o.MaximumConnections
	}
	return json.Marshal(toSerialize)
}

type NullableEndpoint struct {
	value *Endpoint
	isSet bool
}

func (v NullableEndpoint) Get() *Endpoint {
	return v.value
}

func (v *NullableEndpoint) Set(val *Endpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpoint(val *Endpoint) *NullableEndpoint {
	return &NullableEndpoint{value: val, isSet: true}
}

func (v NullableEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
