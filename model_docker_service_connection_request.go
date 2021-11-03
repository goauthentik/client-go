/*
authentik

Making authentication simple.

API version: 2021.10.2
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DockerServiceConnectionRequest DockerServiceConnection Serializer
type DockerServiceConnectionRequest struct {
	Name string `json:"name"`
	// If enabled, use the local connection. Required Docker socket/Kubernetes Integration
	Local *bool `json:"local,omitempty"`
	// Can be in the format of 'unix://<path>' when connecting to a local docker daemon, or 'https://<hostname>:2376' when connecting to a remote system.
	Url string `json:"url"`
	// CA which the endpoint's Certificate is verified against. Can be left empty for no validation.
	TlsVerification NullableString `json:"tls_verification,omitempty"`
	// Certificate/Key used for authentication. Can be left empty for no authentication.
	TlsAuthentication NullableString `json:"tls_authentication,omitempty"`
}

// NewDockerServiceConnectionRequest instantiates a new DockerServiceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerServiceConnectionRequest(name string, url string) *DockerServiceConnectionRequest {
	this := DockerServiceConnectionRequest{}
	this.Name = name
	this.Url = url
	return &this
}

// NewDockerServiceConnectionRequestWithDefaults instantiates a new DockerServiceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerServiceConnectionRequestWithDefaults() *DockerServiceConnectionRequest {
	this := DockerServiceConnectionRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DockerServiceConnectionRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DockerServiceConnectionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DockerServiceConnectionRequest) SetName(v string) {
	o.Name = v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *DockerServiceConnectionRequest) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerServiceConnectionRequest) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *DockerServiceConnectionRequest) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *DockerServiceConnectionRequest) SetLocal(v bool) {
	o.Local = &v
}

// GetUrl returns the Url field value
func (o *DockerServiceConnectionRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DockerServiceConnectionRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DockerServiceConnectionRequest) SetUrl(v string) {
	o.Url = v
}

// GetTlsVerification returns the TlsVerification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DockerServiceConnectionRequest) GetTlsVerification() string {
	if o == nil || o.TlsVerification.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsVerification.Get()
}

// GetTlsVerificationOk returns a tuple with the TlsVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DockerServiceConnectionRequest) GetTlsVerificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsVerification.Get(), o.TlsVerification.IsSet()
}

// HasTlsVerification returns a boolean if a field has been set.
func (o *DockerServiceConnectionRequest) HasTlsVerification() bool {
	if o != nil && o.TlsVerification.IsSet() {
		return true
	}

	return false
}

// SetTlsVerification gets a reference to the given NullableString and assigns it to the TlsVerification field.
func (o *DockerServiceConnectionRequest) SetTlsVerification(v string) {
	o.TlsVerification.Set(&v)
}

// SetTlsVerificationNil sets the value for TlsVerification to be an explicit nil
func (o *DockerServiceConnectionRequest) SetTlsVerificationNil() {
	o.TlsVerification.Set(nil)
}

// UnsetTlsVerification ensures that no value is present for TlsVerification, not even an explicit nil
func (o *DockerServiceConnectionRequest) UnsetTlsVerification() {
	o.TlsVerification.Unset()
}

// GetTlsAuthentication returns the TlsAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DockerServiceConnectionRequest) GetTlsAuthentication() string {
	if o == nil || o.TlsAuthentication.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsAuthentication.Get()
}

// GetTlsAuthenticationOk returns a tuple with the TlsAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DockerServiceConnectionRequest) GetTlsAuthenticationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsAuthentication.Get(), o.TlsAuthentication.IsSet()
}

// HasTlsAuthentication returns a boolean if a field has been set.
func (o *DockerServiceConnectionRequest) HasTlsAuthentication() bool {
	if o != nil && o.TlsAuthentication.IsSet() {
		return true
	}

	return false
}

// SetTlsAuthentication gets a reference to the given NullableString and assigns it to the TlsAuthentication field.
func (o *DockerServiceConnectionRequest) SetTlsAuthentication(v string) {
	o.TlsAuthentication.Set(&v)
}

// SetTlsAuthenticationNil sets the value for TlsAuthentication to be an explicit nil
func (o *DockerServiceConnectionRequest) SetTlsAuthenticationNil() {
	o.TlsAuthentication.Set(nil)
}

// UnsetTlsAuthentication ensures that no value is present for TlsAuthentication, not even an explicit nil
func (o *DockerServiceConnectionRequest) UnsetTlsAuthentication() {
	o.TlsAuthentication.Unset()
}

func (o DockerServiceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
	}
	if true {
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

type NullableDockerServiceConnectionRequest struct {
	value *DockerServiceConnectionRequest
	isSet bool
}

func (v NullableDockerServiceConnectionRequest) Get() *DockerServiceConnectionRequest {
	return v.value
}

func (v *NullableDockerServiceConnectionRequest) Set(val *DockerServiceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerServiceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerServiceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerServiceConnectionRequest(val *DockerServiceConnectionRequest) *NullableDockerServiceConnectionRequest {
	return &NullableDockerServiceConnectionRequest{value: val, isSet: true}
}

func (v NullableDockerServiceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerServiceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
