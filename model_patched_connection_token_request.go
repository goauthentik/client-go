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

// PatchedConnectionTokenRequest ConnectionToken Serializer
type PatchedConnectionTokenRequest struct {
	Pk       *string `json:"pk,omitempty"`
	Provider *int32  `json:"provider,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
}

// NewPatchedConnectionTokenRequest instantiates a new PatchedConnectionTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedConnectionTokenRequest() *PatchedConnectionTokenRequest {
	this := PatchedConnectionTokenRequest{}
	return &this
}

// NewPatchedConnectionTokenRequestWithDefaults instantiates a new PatchedConnectionTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedConnectionTokenRequestWithDefaults() *PatchedConnectionTokenRequest {
	this := PatchedConnectionTokenRequest{}
	return &this
}

// GetPk returns the Pk field value if set, zero value otherwise.
func (o *PatchedConnectionTokenRequest) GetPk() string {
	if o == nil || o.Pk == nil {
		var ret string
		return ret
	}
	return *o.Pk
}

// GetPkOk returns a tuple with the Pk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConnectionTokenRequest) GetPkOk() (*string, bool) {
	if o == nil || o.Pk == nil {
		return nil, false
	}
	return o.Pk, true
}

// HasPk returns a boolean if a field has been set.
func (o *PatchedConnectionTokenRequest) HasPk() bool {
	if o != nil && o.Pk != nil {
		return true
	}

	return false
}

// SetPk gets a reference to the given string and assigns it to the Pk field.
func (o *PatchedConnectionTokenRequest) SetPk(v string) {
	o.Pk = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PatchedConnectionTokenRequest) GetProvider() int32 {
	if o == nil || o.Provider == nil {
		var ret int32
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConnectionTokenRequest) GetProviderOk() (*int32, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PatchedConnectionTokenRequest) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given int32 and assigns it to the Provider field.
func (o *PatchedConnectionTokenRequest) SetProvider(v int32) {
	o.Provider = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *PatchedConnectionTokenRequest) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConnectionTokenRequest) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *PatchedConnectionTokenRequest) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *PatchedConnectionTokenRequest) SetEndpoint(v string) {
	o.Endpoint = &v
}

func (o PatchedConnectionTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pk != nil {
		toSerialize["pk"] = o.Pk
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedConnectionTokenRequest struct {
	value *PatchedConnectionTokenRequest
	isSet bool
}

func (v NullablePatchedConnectionTokenRequest) Get() *PatchedConnectionTokenRequest {
	return v.value
}

func (v *NullablePatchedConnectionTokenRequest) Set(val *PatchedConnectionTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedConnectionTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedConnectionTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedConnectionTokenRequest(val *PatchedConnectionTokenRequest) *NullablePatchedConnectionTokenRequest {
	return &NullablePatchedConnectionTokenRequest{value: val, isSet: true}
}

func (v NullablePatchedConnectionTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedConnectionTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
