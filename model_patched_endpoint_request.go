/*
authentik

Making authentication simple.

API version: 2025.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedEndpointRequest Endpoint Serializer
type PatchedEndpointRequest struct {
	Name               *string       `json:"name,omitempty"`
	Provider           *int32        `json:"provider,omitempty"`
	Protocol           *ProtocolEnum `json:"protocol,omitempty"`
	Host               *string       `json:"host,omitempty"`
	Settings           interface{}   `json:"settings,omitempty"`
	PropertyMappings   []string      `json:"property_mappings,omitempty"`
	AuthMode           *AuthModeEnum `json:"auth_mode,omitempty"`
	MaximumConnections *int32        `json:"maximum_connections,omitempty"`
}

// NewPatchedEndpointRequest instantiates a new PatchedEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEndpointRequest() *PatchedEndpointRequest {
	this := PatchedEndpointRequest{}
	return &this
}

// NewPatchedEndpointRequestWithDefaults instantiates a new PatchedEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEndpointRequestWithDefaults() *PatchedEndpointRequest {
	this := PatchedEndpointRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedEndpointRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEndpointRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedEndpointRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedEndpointRequest) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PatchedEndpointRequest) GetProvider() int32 {
	if o == nil || o.Provider == nil {
		var ret int32
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEndpointRequest) GetProviderOk() (*int32, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PatchedEndpointRequest) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given int32 and assigns it to the Provider field.
func (o *PatchedEndpointRequest) SetProvider(v int32) {
	o.Provider = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *PatchedEndpointRequest) GetProtocol() ProtocolEnum {
	if o == nil || o.Protocol == nil {
		var ret ProtocolEnum
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEndpointRequest) GetProtocolOk() (*ProtocolEnum, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *PatchedEndpointRequest) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given ProtocolEnum and assigns it to the Protocol field.
func (o *PatchedEndpointRequest) SetProtocol(v ProtocolEnum) {
	o.Protocol = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *PatchedEndpointRequest) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEndpointRequest) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *PatchedEndpointRequest) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *PatchedEndpointRequest) SetHost(v string) {
	o.Host = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEndpointRequest) GetSettings() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEndpointRequest) GetSettingsOk() (*interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *PatchedEndpointRequest) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given interface{} and assigns it to the Settings field.
func (o *PatchedEndpointRequest) SetSettings(v interface{}) {
	o.Settings = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *PatchedEndpointRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEndpointRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedEndpointRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *PatchedEndpointRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *PatchedEndpointRequest) GetAuthMode() AuthModeEnum {
	if o == nil || o.AuthMode == nil {
		var ret AuthModeEnum
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEndpointRequest) GetAuthModeOk() (*AuthModeEnum, bool) {
	if o == nil || o.AuthMode == nil {
		return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *PatchedEndpointRequest) HasAuthMode() bool {
	if o != nil && o.AuthMode != nil {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given AuthModeEnum and assigns it to the AuthMode field.
func (o *PatchedEndpointRequest) SetAuthMode(v AuthModeEnum) {
	o.AuthMode = &v
}

// GetMaximumConnections returns the MaximumConnections field value if set, zero value otherwise.
func (o *PatchedEndpointRequest) GetMaximumConnections() int32 {
	if o == nil || o.MaximumConnections == nil {
		var ret int32
		return ret
	}
	return *o.MaximumConnections
}

// GetMaximumConnectionsOk returns a tuple with the MaximumConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEndpointRequest) GetMaximumConnectionsOk() (*int32, bool) {
	if o == nil || o.MaximumConnections == nil {
		return nil, false
	}
	return o.MaximumConnections, true
}

// HasMaximumConnections returns a boolean if a field has been set.
func (o *PatchedEndpointRequest) HasMaximumConnections() bool {
	if o != nil && o.MaximumConnections != nil {
		return true
	}

	return false
}

// SetMaximumConnections gets a reference to the given int32 and assigns it to the MaximumConnections field.
func (o *PatchedEndpointRequest) SetMaximumConnections(v int32) {
	o.MaximumConnections = &v
}

func (o PatchedEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	if o.AuthMode != nil {
		toSerialize["auth_mode"] = o.AuthMode
	}
	if o.MaximumConnections != nil {
		toSerialize["maximum_connections"] = o.MaximumConnections
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedEndpointRequest struct {
	value *PatchedEndpointRequest
	isSet bool
}

func (v NullablePatchedEndpointRequest) Get() *PatchedEndpointRequest {
	return v.value
}

func (v *NullablePatchedEndpointRequest) Set(val *PatchedEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEndpointRequest(val *PatchedEndpointRequest) *NullablePatchedEndpointRequest {
	return &NullablePatchedEndpointRequest{value: val, isSet: true}
}

func (v NullablePatchedEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
