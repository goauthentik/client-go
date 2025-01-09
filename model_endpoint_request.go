/*
authentik

Making authentication simple.

API version: 2024.12.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EndpointRequest Endpoint Serializer
type EndpointRequest struct {
	Name               string       `json:"name"`
	Provider           int32        `json:"provider"`
	Protocol           ProtocolEnum `json:"protocol"`
	Host               string       `json:"host"`
	Settings           interface{}  `json:"settings,omitempty"`
	PropertyMappings   []string     `json:"property_mappings,omitempty"`
	AuthMode           AuthModeEnum `json:"auth_mode"`
	MaximumConnections *int32       `json:"maximum_connections,omitempty"`
}

// NewEndpointRequest instantiates a new EndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointRequest(name string, provider int32, protocol ProtocolEnum, host string, authMode AuthModeEnum) *EndpointRequest {
	this := EndpointRequest{}
	this.Name = name
	this.Provider = provider
	this.Protocol = protocol
	this.Host = host
	this.AuthMode = authMode
	return &this
}

// NewEndpointRequestWithDefaults instantiates a new EndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointRequestWithDefaults() *EndpointRequest {
	this := EndpointRequest{}
	return &this
}

// GetName returns the Name field value
func (o *EndpointRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EndpointRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EndpointRequest) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value
func (o *EndpointRequest) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *EndpointRequest) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *EndpointRequest) SetProvider(v int32) {
	o.Provider = v
}

// GetProtocol returns the Protocol field value
func (o *EndpointRequest) GetProtocol() ProtocolEnum {
	if o == nil {
		var ret ProtocolEnum
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *EndpointRequest) GetProtocolOk() (*ProtocolEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *EndpointRequest) SetProtocol(v ProtocolEnum) {
	o.Protocol = v
}

// GetHost returns the Host field value
func (o *EndpointRequest) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *EndpointRequest) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *EndpointRequest) SetHost(v string) {
	o.Host = v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointRequest) GetSettings() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointRequest) GetSettingsOk() (*interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *EndpointRequest) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given interface{} and assigns it to the Settings field.
func (o *EndpointRequest) SetSettings(v interface{}) {
	o.Settings = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *EndpointRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *EndpointRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *EndpointRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetAuthMode returns the AuthMode field value
func (o *EndpointRequest) GetAuthMode() AuthModeEnum {
	if o == nil {
		var ret AuthModeEnum
		return ret
	}

	return o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value
// and a boolean to check if the value has been set.
func (o *EndpointRequest) GetAuthModeOk() (*AuthModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthMode, true
}

// SetAuthMode sets field value
func (o *EndpointRequest) SetAuthMode(v AuthModeEnum) {
	o.AuthMode = v
}

// GetMaximumConnections returns the MaximumConnections field value if set, zero value otherwise.
func (o *EndpointRequest) GetMaximumConnections() int32 {
	if o == nil || o.MaximumConnections == nil {
		var ret int32
		return ret
	}
	return *o.MaximumConnections
}

// GetMaximumConnectionsOk returns a tuple with the MaximumConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRequest) GetMaximumConnectionsOk() (*int32, bool) {
	if o == nil || o.MaximumConnections == nil {
		return nil, false
	}
	return o.MaximumConnections, true
}

// HasMaximumConnections returns a boolean if a field has been set.
func (o *EndpointRequest) HasMaximumConnections() bool {
	if o != nil && o.MaximumConnections != nil {
		return true
	}

	return false
}

// SetMaximumConnections gets a reference to the given int32 and assigns it to the MaximumConnections field.
func (o *EndpointRequest) SetMaximumConnections(v int32) {
	o.MaximumConnections = &v
}

func (o EndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["provider"] = o.Provider
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
	if o.MaximumConnections != nil {
		toSerialize["maximum_connections"] = o.MaximumConnections
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointRequest struct {
	value *EndpointRequest
	isSet bool
}

func (v NullableEndpointRequest) Get() *EndpointRequest {
	return v.value
}

func (v *NullableEndpointRequest) Set(val *EndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointRequest(val *EndpointRequest) *NullableEndpointRequest {
	return &NullableEndpointRequest{value: val, isSet: true}
}

func (v NullableEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
