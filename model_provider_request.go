/*
authentik

Making authentication simple.

API version: 2025.6.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ProviderRequest Provider Serializer
type ProviderRequest struct {
	Name string `json:"name"`
	// Flow used for authentication when the associated application is accessed by an un-authenticated user.
	AuthenticationFlow NullableString `json:"authentication_flow,omitempty"`
	// Flow used when authorizing this provider.
	AuthorizationFlow string `json:"authorization_flow"`
	// Flow used ending the session from a provider.
	InvalidationFlow string   `json:"invalidation_flow"`
	PropertyMappings []string `json:"property_mappings,omitempty"`
}

// NewProviderRequest instantiates a new ProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderRequest(name string, authorizationFlow string, invalidationFlow string) *ProviderRequest {
	this := ProviderRequest{}
	this.Name = name
	this.AuthorizationFlow = authorizationFlow
	this.InvalidationFlow = invalidationFlow
	return &this
}

// NewProviderRequestWithDefaults instantiates a new ProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderRequestWithDefaults() *ProviderRequest {
	this := ProviderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ProviderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProviderRequest) SetName(v string) {
	o.Name = v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderRequest) GetAuthenticationFlow() string {
	if o == nil || o.AuthenticationFlow.Get() == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationFlow.Get()
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderRequest) GetAuthenticationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationFlow.Get(), o.AuthenticationFlow.IsSet()
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *ProviderRequest) HasAuthenticationFlow() bool {
	if o != nil && o.AuthenticationFlow.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given NullableString and assigns it to the AuthenticationFlow field.
func (o *ProviderRequest) SetAuthenticationFlow(v string) {
	o.AuthenticationFlow.Set(&v)
}

// SetAuthenticationFlowNil sets the value for AuthenticationFlow to be an explicit nil
func (o *ProviderRequest) SetAuthenticationFlowNil() {
	o.AuthenticationFlow.Set(nil)
}

// UnsetAuthenticationFlow ensures that no value is present for AuthenticationFlow, not even an explicit nil
func (o *ProviderRequest) UnsetAuthenticationFlow() {
	o.AuthenticationFlow.Unset()
}

// GetAuthorizationFlow returns the AuthorizationFlow field value
func (o *ProviderRequest) GetAuthorizationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationFlow
}

// GetAuthorizationFlowOk returns a tuple with the AuthorizationFlow field value
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetAuthorizationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationFlow, true
}

// SetAuthorizationFlow sets field value
func (o *ProviderRequest) SetAuthorizationFlow(v string) {
	o.AuthorizationFlow = v
}

// GetInvalidationFlow returns the InvalidationFlow field value
func (o *ProviderRequest) GetInvalidationFlow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvalidationFlow
}

// GetInvalidationFlowOk returns a tuple with the InvalidationFlow field value
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetInvalidationFlowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvalidationFlow, true
}

// SetInvalidationFlow sets field value
func (o *ProviderRequest) SetInvalidationFlow(v string) {
	o.InvalidationFlow = v
}

// GetPropertyMappings returns the PropertyMappings field value if set, zero value otherwise.
func (o *ProviderRequest) GetPropertyMappings() []string {
	if o == nil || o.PropertyMappings == nil {
		var ret []string
		return ret
	}
	return o.PropertyMappings
}

// GetPropertyMappingsOk returns a tuple with the PropertyMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRequest) GetPropertyMappingsOk() ([]string, bool) {
	if o == nil || o.PropertyMappings == nil {
		return nil, false
	}
	return o.PropertyMappings, true
}

// HasPropertyMappings returns a boolean if a field has been set.
func (o *ProviderRequest) HasPropertyMappings() bool {
	if o != nil && o.PropertyMappings != nil {
		return true
	}

	return false
}

// SetPropertyMappings gets a reference to the given []string and assigns it to the PropertyMappings field.
func (o *ProviderRequest) SetPropertyMappings(v []string) {
	o.PropertyMappings = v
}

func (o ProviderRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["invalidation_flow"] = o.InvalidationFlow
	}
	if o.PropertyMappings != nil {
		toSerialize["property_mappings"] = o.PropertyMappings
	}
	return json.Marshal(toSerialize)
}

type NullableProviderRequest struct {
	value *ProviderRequest
	isSet bool
}

func (v NullableProviderRequest) Get() *ProviderRequest {
	return v.value
}

func (v *NullableProviderRequest) Set(val *ProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderRequest(val *ProviderRequest) *NullableProviderRequest {
	return &NullableProviderRequest{value: val, isSet: true}
}

func (v NullableProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
