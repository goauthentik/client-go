/*
authentik

Making authentication simple.

API version: 2024.4.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ConnectionTokenRequest ConnectionToken Serializer
type ConnectionTokenRequest struct {
	Pk       *string `json:"pk,omitempty"`
	Provider int32   `json:"provider"`
	Endpoint string  `json:"endpoint"`
}

// NewConnectionTokenRequest instantiates a new ConnectionTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionTokenRequest(provider int32, endpoint string) *ConnectionTokenRequest {
	this := ConnectionTokenRequest{}
	this.Provider = provider
	this.Endpoint = endpoint
	return &this
}

// NewConnectionTokenRequestWithDefaults instantiates a new ConnectionTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionTokenRequestWithDefaults() *ConnectionTokenRequest {
	this := ConnectionTokenRequest{}
	return &this
}

// GetPk returns the Pk field value if set, zero value otherwise.
func (o *ConnectionTokenRequest) GetPk() string {
	if o == nil || o.Pk == nil {
		var ret string
		return ret
	}
	return *o.Pk
}

// GetPkOk returns a tuple with the Pk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionTokenRequest) GetPkOk() (*string, bool) {
	if o == nil || o.Pk == nil {
		return nil, false
	}
	return o.Pk, true
}

// HasPk returns a boolean if a field has been set.
func (o *ConnectionTokenRequest) HasPk() bool {
	if o != nil && o.Pk != nil {
		return true
	}

	return false
}

// SetPk gets a reference to the given string and assigns it to the Pk field.
func (o *ConnectionTokenRequest) SetPk(v string) {
	o.Pk = &v
}

// GetProvider returns the Provider field value
func (o *ConnectionTokenRequest) GetProvider() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ConnectionTokenRequest) GetProviderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *ConnectionTokenRequest) SetProvider(v int32) {
	o.Provider = v
}

// GetEndpoint returns the Endpoint field value
func (o *ConnectionTokenRequest) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *ConnectionTokenRequest) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *ConnectionTokenRequest) SetEndpoint(v string) {
	o.Endpoint = v
}

func (o ConnectionTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pk != nil {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["endpoint"] = o.Endpoint
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionTokenRequest struct {
	value *ConnectionTokenRequest
	isSet bool
}

func (v NullableConnectionTokenRequest) Get() *ConnectionTokenRequest {
	return v.value
}

func (v *NullableConnectionTokenRequest) Set(val *ConnectionTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionTokenRequest(val *ConnectionTokenRequest) *NullableConnectionTokenRequest {
	return &NullableConnectionTokenRequest{value: val, isSet: true}
}

func (v NullableConnectionTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
