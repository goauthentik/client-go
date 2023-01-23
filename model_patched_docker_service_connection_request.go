/*
authentik

Making authentication simple.

API version: 2023.1.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedDockerServiceConnectionRequest DockerServiceConnection Serializer
type PatchedDockerServiceConnectionRequest struct {
	Name *string `json:"name,omitempty"`
	// If enabled, use the local connection. Required Docker socket/Kubernetes Integration
	Local *bool `json:"local,omitempty"`
	// Can be in the format of 'unix://<path>' when connecting to a local docker daemon, or 'https://<hostname>:2376' when connecting to a remote system.
	Url *string `json:"url,omitempty"`
	// CA which the endpoint's Certificate is verified against. Can be left empty for no validation.
	TlsVerification NullableString `json:"tls_verification,omitempty"`
	// Certificate/Key used for authentication. Can be left empty for no authentication.
	TlsAuthentication NullableString `json:"tls_authentication,omitempty"`
}

// NewPatchedDockerServiceConnectionRequest instantiates a new PatchedDockerServiceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDockerServiceConnectionRequest() *PatchedDockerServiceConnectionRequest {
	this := PatchedDockerServiceConnectionRequest{}
	return &this
}

// NewPatchedDockerServiceConnectionRequestWithDefaults instantiates a new PatchedDockerServiceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDockerServiceConnectionRequestWithDefaults() *PatchedDockerServiceConnectionRequest {
	this := PatchedDockerServiceConnectionRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedDockerServiceConnectionRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDockerServiceConnectionRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedDockerServiceConnectionRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedDockerServiceConnectionRequest) SetName(v string) {
	o.Name = &v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *PatchedDockerServiceConnectionRequest) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDockerServiceConnectionRequest) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *PatchedDockerServiceConnectionRequest) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *PatchedDockerServiceConnectionRequest) SetLocal(v bool) {
	o.Local = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedDockerServiceConnectionRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDockerServiceConnectionRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedDockerServiceConnectionRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedDockerServiceConnectionRequest) SetUrl(v string) {
	o.Url = &v
}

// GetTlsVerification returns the TlsVerification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDockerServiceConnectionRequest) GetTlsVerification() string {
	if o == nil || o.TlsVerification.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsVerification.Get()
}

// GetTlsVerificationOk returns a tuple with the TlsVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDockerServiceConnectionRequest) GetTlsVerificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsVerification.Get(), o.TlsVerification.IsSet()
}

// HasTlsVerification returns a boolean if a field has been set.
func (o *PatchedDockerServiceConnectionRequest) HasTlsVerification() bool {
	if o != nil && o.TlsVerification.IsSet() {
		return true
	}

	return false
}

// SetTlsVerification gets a reference to the given NullableString and assigns it to the TlsVerification field.
func (o *PatchedDockerServiceConnectionRequest) SetTlsVerification(v string) {
	o.TlsVerification.Set(&v)
}

// SetTlsVerificationNil sets the value for TlsVerification to be an explicit nil
func (o *PatchedDockerServiceConnectionRequest) SetTlsVerificationNil() {
	o.TlsVerification.Set(nil)
}

// UnsetTlsVerification ensures that no value is present for TlsVerification, not even an explicit nil
func (o *PatchedDockerServiceConnectionRequest) UnsetTlsVerification() {
	o.TlsVerification.Unset()
}

// GetTlsAuthentication returns the TlsAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDockerServiceConnectionRequest) GetTlsAuthentication() string {
	if o == nil || o.TlsAuthentication.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsAuthentication.Get()
}

// GetTlsAuthenticationOk returns a tuple with the TlsAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDockerServiceConnectionRequest) GetTlsAuthenticationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsAuthentication.Get(), o.TlsAuthentication.IsSet()
}

// HasTlsAuthentication returns a boolean if a field has been set.
func (o *PatchedDockerServiceConnectionRequest) HasTlsAuthentication() bool {
	if o != nil && o.TlsAuthentication.IsSet() {
		return true
	}

	return false
}

// SetTlsAuthentication gets a reference to the given NullableString and assigns it to the TlsAuthentication field.
func (o *PatchedDockerServiceConnectionRequest) SetTlsAuthentication(v string) {
	o.TlsAuthentication.Set(&v)
}

// SetTlsAuthenticationNil sets the value for TlsAuthentication to be an explicit nil
func (o *PatchedDockerServiceConnectionRequest) SetTlsAuthenticationNil() {
	o.TlsAuthentication.Set(nil)
}

// UnsetTlsAuthentication ensures that no value is present for TlsAuthentication, not even an explicit nil
func (o *PatchedDockerServiceConnectionRequest) UnsetTlsAuthentication() {
	o.TlsAuthentication.Unset()
}

func (o PatchedDockerServiceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.TlsVerification.IsSet() {
		toSerialize["tls_verification"] = o.TlsVerification.Get()
	}
	if o.TlsAuthentication.IsSet() {
		toSerialize["tls_authentication"] = o.TlsAuthentication.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedDockerServiceConnectionRequest struct {
	value *PatchedDockerServiceConnectionRequest
	isSet bool
}

func (v NullablePatchedDockerServiceConnectionRequest) Get() *PatchedDockerServiceConnectionRequest {
	return v.value
}

func (v *NullablePatchedDockerServiceConnectionRequest) Set(val *PatchedDockerServiceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDockerServiceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDockerServiceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDockerServiceConnectionRequest(val *PatchedDockerServiceConnectionRequest) *NullablePatchedDockerServiceConnectionRequest {
	return &NullablePatchedDockerServiceConnectionRequest{value: val, isSet: true}
}

func (v NullablePatchedDockerServiceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDockerServiceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
