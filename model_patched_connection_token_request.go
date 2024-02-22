/*
authentik

Making authentication simple.

API version: 2024.2.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedConnectionTokenRequest ConnectionToken Serializer
type PatchedConnectionTokenRequest struct {
	Provider *int32 `json:"provider,omitempty"`
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

func (o PatchedConnectionTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
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
