/*
authentik

Making authentication simple.

API version: 2024.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DockerServiceConnection DockerServiceConnection Serializer
type DockerServiceConnection struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// If enabled, use the local connection. Required Docker socket/Kubernetes Integration
	Local     *bool  `json:"local,omitempty"`
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string `json:"meta_model_name"`
	// Can be in the format of 'unix://<path>' when connecting to a local docker daemon, or 'https://<hostname>:2376' when connecting to a remote system.
	Url string `json:"url"`
	// CA which the endpoint's Certificate is verified against. Can be left empty for no validation.
	TlsVerification NullableString `json:"tls_verification,omitempty"`
	// Certificate/Key used for authentication. Can be left empty for no authentication.
	TlsAuthentication NullableString `json:"tls_authentication,omitempty"`
}

// NewDockerServiceConnection instantiates a new DockerServiceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDockerServiceConnection(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, url string) *DockerServiceConnection {
	this := DockerServiceConnection{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.Url = url
	return &this
}

// NewDockerServiceConnectionWithDefaults instantiates a new DockerServiceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerServiceConnectionWithDefaults() *DockerServiceConnection {
	this := DockerServiceConnection{}
	return &this
}

// GetPk returns the Pk field value
func (o *DockerServiceConnection) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *DockerServiceConnection) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *DockerServiceConnection) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *DockerServiceConnection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DockerServiceConnection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DockerServiceConnection) SetName(v string) {
	o.Name = v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *DockerServiceConnection) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DockerServiceConnection) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *DockerServiceConnection) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *DockerServiceConnection) SetLocal(v bool) {
	o.Local = &v
}

// GetComponent returns the Component field value
func (o *DockerServiceConnection) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DockerServiceConnection) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *DockerServiceConnection) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *DockerServiceConnection) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *DockerServiceConnection) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *DockerServiceConnection) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *DockerServiceConnection) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *DockerServiceConnection) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *DockerServiceConnection) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *DockerServiceConnection) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *DockerServiceConnection) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *DockerServiceConnection) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetUrl returns the Url field value
func (o *DockerServiceConnection) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DockerServiceConnection) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DockerServiceConnection) SetUrl(v string) {
	o.Url = v
}

// GetTlsVerification returns the TlsVerification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DockerServiceConnection) GetTlsVerification() string {
	if o == nil || o.TlsVerification.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsVerification.Get()
}

// GetTlsVerificationOk returns a tuple with the TlsVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DockerServiceConnection) GetTlsVerificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsVerification.Get(), o.TlsVerification.IsSet()
}

// HasTlsVerification returns a boolean if a field has been set.
func (o *DockerServiceConnection) HasTlsVerification() bool {
	if o != nil && o.TlsVerification.IsSet() {
		return true
	}

	return false
}

// SetTlsVerification gets a reference to the given NullableString and assigns it to the TlsVerification field.
func (o *DockerServiceConnection) SetTlsVerification(v string) {
	o.TlsVerification.Set(&v)
}

// SetTlsVerificationNil sets the value for TlsVerification to be an explicit nil
func (o *DockerServiceConnection) SetTlsVerificationNil() {
	o.TlsVerification.Set(nil)
}

// UnsetTlsVerification ensures that no value is present for TlsVerification, not even an explicit nil
func (o *DockerServiceConnection) UnsetTlsVerification() {
	o.TlsVerification.Unset()
}

// GetTlsAuthentication returns the TlsAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DockerServiceConnection) GetTlsAuthentication() string {
	if o == nil || o.TlsAuthentication.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsAuthentication.Get()
}

// GetTlsAuthenticationOk returns a tuple with the TlsAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DockerServiceConnection) GetTlsAuthenticationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsAuthentication.Get(), o.TlsAuthentication.IsSet()
}

// HasTlsAuthentication returns a boolean if a field has been set.
func (o *DockerServiceConnection) HasTlsAuthentication() bool {
	if o != nil && o.TlsAuthentication.IsSet() {
		return true
	}

	return false
}

// SetTlsAuthentication gets a reference to the given NullableString and assigns it to the TlsAuthentication field.
func (o *DockerServiceConnection) SetTlsAuthentication(v string) {
	o.TlsAuthentication.Set(&v)
}

// SetTlsAuthenticationNil sets the value for TlsAuthentication to be an explicit nil
func (o *DockerServiceConnection) SetTlsAuthenticationNil() {
	o.TlsAuthentication.Set(nil)
}

// UnsetTlsAuthentication ensures that no value is present for TlsAuthentication, not even an explicit nil
func (o *DockerServiceConnection) UnsetTlsAuthentication() {
	o.TlsAuthentication.Unset()
}

func (o DockerServiceConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
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

type NullableDockerServiceConnection struct {
	value *DockerServiceConnection
	isSet bool
}

func (v NullableDockerServiceConnection) Get() *DockerServiceConnection {
	return v.value
}

func (v *NullableDockerServiceConnection) Set(val *DockerServiceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableDockerServiceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableDockerServiceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDockerServiceConnection(val *DockerServiceConnection) *NullableDockerServiceConnection {
	return &NullableDockerServiceConnection{value: val, isSet: true}
}

func (v NullableDockerServiceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDockerServiceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
