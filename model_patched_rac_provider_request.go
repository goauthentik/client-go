/*
authentik

Making authentication simple.

API version: 2024.8.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedRACProviderRequest RACProvider Serializer
type PatchedRACProviderRequest struct {
	Name *string `json:"name,omitempty"`
	// Flow used for authentication when the associated application is accessed by an un-authenticated user.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow *string     `json:"authorization_flow,omitempty"`
	PropertyMappings  []string    `json:"property_mappings,omitempty"`
	Settings          interface{} `json:"settings,omitempty"`
	// Determines how long a session lasts. Default of 0 means that the sessions lasts until the browser is closed. (Format: hours=-1;minutes=-2;seconds=-3)
	ConnectionExpiry *string `json:"connection_expiry,omitempty"`
	// When set to true, connection tokens will be deleted upon disconnect.
	DeleteTokenOnDisconnect *bool `json:"delete_token_on_disconnect,omitempty"`
}

// NewPatchedRACProviderRequest instantiates a new PatchedRACProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRACProviderRequest() *PatchedRACProviderRequest {
	this := PatchedRACProviderRequest{}
	return &this
}

// NewPatchedRACProviderRequestWithDefaults instantiates a new PatchedRACProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRACProviderRequestWithDefaults() *PatchedRACProviderRequest {
	this := PatchedRACProviderRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedRACProviderRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRACProviderRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedRACProviderRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedRACProviderRequest) SetName(v string) {
	o.Name = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRACProviderRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRACProviderRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *PatchedRACProviderRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *PatchedRACProviderRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *PatchedRACProviderRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *PatchedRACProviderRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetAuthorizationFlow returns the AuthorizationFlow field value if set, zero value otherwise.
func (o *PatchedRACProviderRequest) GetAuthorizationFlow() string {
	if o == nil || o.AuthorizationFlow == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRACProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil || o.AuthorizationFlow == nil {
		return nil, false
	}
	return o.AuthorizationFlow, true
}

// HasAuthorizationFlow returns a boolean if a field has been set.
func (o *PatchedRACProviderRequest) HasAuthorizationFlow() bool {
	if o != nil && o.AuthorizationFlow != nil {
		return true
	}

	return false
}

// SetAuthorizationFlow gets a reference to the given string and assigns it to the AuthorizationFlow field.
func (o *PatchedRACProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = &v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *PatchedRACProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRACProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *PatchedRACProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *PatchedRACProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRACProviderRequest) GetSettings() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRACProviderRequest) GetSettingsOk() (*interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *PatchedRACProviderRequest) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given interface{} and assigns it to the Settings field.
func (o *PatchedRACProviderRequest) SetSettings(v interface{}) {
	o.Settings = v
}

// GetConnectionExpiry returns the ConnectionExpiry field value if set, zero value otherwise.
func (o *PatchedRACProviderRequest) GetConnectionExpiry() string {
	if o == nil || o.ConnectionExpiry == nil {
		var ret string
		return ret
	}
	return *o.ConnectionExpiry
}

// GetConnectionExpiryOk returns a tuple with the ConnectionExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRACProviderRequest) GetConnectionExpiryOk() (*string, bool) {
	if o == nil || o.ConnectionExpiry == nil {
		return nil, false
	}
	return o.ConnectionExpiry, true
}

// HasConnectionExpiry returns a boolean if a field has been set.
func (o *PatchedRACProviderRequest) HasConnectionExpiry() bool {
	if o != nil && o.ConnectionExpiry != nil {
		return true
	}

	return false
}

// SetConnectionExpiry gets a reference to the given string and assigns it to the ConnectionExpiry field.
func (o *PatchedRACProviderRequest) SetConnectionExpiry(v string) {
	o.ConnectionExpiry = &v
}

// GetDeleteTokenOnDisconnect returns the DeleteTokenOnDisconnect field value if set, zero value otherwise.
func (o *PatchedRACProviderRequest) GetDeleteTokenOnDisconnect() bool {
	if o == nil || o.DeleteTokenOnDisconnect == nil {
		var ret bool
		return ret
	}
	return *o.DeleteTokenOnDisconnect
}

// GetDeleteTokenOnDisconnectOk returns a tuple with the DeleteTokenOnDisconnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRACProviderRequest) GetDeleteTokenOnDisconnectOk() (*bool, bool) {
	if o == nil || o.DeleteTokenOnDisconnect == nil {
		return nil, false
	}
	return o.DeleteTokenOnDisconnect, true
}

// HasDeleteTokenOnDisconnect returns a boolean if a field has been set.
func (o *PatchedRACProviderRequest) HasDeleteTokenOnDisconnect() bool {
	if o != nil && o.DeleteTokenOnDisconnect != nil {
		return true
	}

	return false
}

// SetDeleteTokenOnDisconnect gets a reference to the given bool and assigns it to the DeleteTokenOnDisconnect field.
func (o *PatchedRACProviderRequest) SetDeleteTokenOnDisconnect(v bool) {
	o.DeleteTokenOnDisconnect = &v
}

func (o PatchedRACProviderRequest) MarshalJSON() ([]byte, error) {
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
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.ConnectionExpiry != nil {
		toSerialize["connection_expiry"] = o.ConnectionExpiry
	}
	if o.DeleteTokenOnDisconnect != nil {
		toSerialize["delete_token_on_disconnect"] = o.DeleteTokenOnDisconnect
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedRACProviderRequest struct {
	value *PatchedRACProviderRequest
	isSet bool
}

func (v NullablePatchedRACProviderRequest) Get() *PatchedRACProviderRequest {
	return v.value
}

func (v *NullablePatchedRACProviderRequest) Set(val *PatchedRACProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRACProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRACProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRACProviderRequest(val *PatchedRACProviderRequest) *NullablePatchedRACProviderRequest {
	return &NullablePatchedRACProviderRequest{value: val, isSet: true}
}

func (v NullablePatchedRACProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRACProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
