/*
authentik

Making authentication simple.

API version: 2022.12.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// KubernetesServiceConnection KubernetesServiceConnection Serializer
type KubernetesServiceConnection struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
	// If enabled, use the local connection. Required Docker socket/Kubernetes Integration
	Local             *bool  `json:"local,omitempty"`
	Component         string `json:"component"`
	VerboseName       string `json:"verbose_name"`
	VerboseNamePlural string `json:"verbose_name_plural"`
	MetaModelName     string `json:"meta_model_name"`
	// Paste your kubeconfig here. authentik will automatically use the currently selected context.
	Kubeconfig map[string]interface{} `json:"kubeconfig,omitempty"`
	// Verify SSL Certificates of the Kubernetes API endpoint
	VerifySsl *bool `json:"verify_ssl,omitempty"`
}

// NewKubernetesServiceConnection instantiates a new KubernetesServiceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesServiceConnection(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string) *KubernetesServiceConnection {
	this := KubernetesServiceConnection{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewKubernetesServiceConnectionWithDefaults instantiates a new KubernetesServiceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesServiceConnectionWithDefaults() *KubernetesServiceConnection {
	this := KubernetesServiceConnection{}
	return &this
}

// GetPk returns the Pk field value
func (o *KubernetesServiceConnection) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnection) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *KubernetesServiceConnection) SetPk(v string) {
	o.Pk = v
}

// GetName returns the Name field value
func (o *KubernetesServiceConnection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesServiceConnection) SetName(v string) {
	o.Name = v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *KubernetesServiceConnection) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnection) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *KubernetesServiceConnection) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *KubernetesServiceConnection) SetLocal(v bool) {
	o.Local = &v
}

// GetComponent returns the Component field value
func (o *KubernetesServiceConnection) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnection) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *KubernetesServiceConnection) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *KubernetesServiceConnection) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnection) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *KubernetesServiceConnection) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *KubernetesServiceConnection) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnection) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *KubernetesServiceConnection) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *KubernetesServiceConnection) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnection) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *KubernetesServiceConnection) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetKubeconfig returns the Kubeconfig field value if set, zero value otherwise.
func (o *KubernetesServiceConnection) GetKubeconfig() map[string]interface{} {
	if o == nil || o.Kubeconfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Kubeconfig
}

// GetKubeconfigOk returns a tuple with the Kubeconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnection) GetKubeconfigOk() (map[string]interface{}, bool) {
	if o == nil || o.Kubeconfig == nil {
		return nil, false
	}
	return o.Kubeconfig, true
}

// HasKubeconfig returns a boolean if a field has been set.
func (o *KubernetesServiceConnection) HasKubeconfig() bool {
	if o != nil && o.Kubeconfig != nil {
		return true
	}

	return false
}

// SetKubeconfig gets a reference to the given map[string]interface{} and assigns it to the Kubeconfig field.
func (o *KubernetesServiceConnection) SetKubeconfig(v map[string]interface{}) {
	o.Kubeconfig = v
}

// GetVerifySsl returns the VerifySsl field value if set, zero value otherwise.
func (o *KubernetesServiceConnection) GetVerifySsl() bool {
	if o == nil || o.VerifySsl == nil {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesServiceConnection) GetVerifySslOk() (*bool, bool) {
	if o == nil || o.VerifySsl == nil {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *KubernetesServiceConnection) HasVerifySsl() bool {
	if o != nil && o.VerifySsl != nil {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *KubernetesServiceConnection) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

func (o KubernetesServiceConnection) MarshalJSON() ([]byte, error) {
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
	if o.Kubeconfig != nil {
		toSerialize["kubeconfig"] = o.Kubeconfig
	}
	if o.VerifySsl != nil {
		toSerialize["verify_ssl"] = o.VerifySsl
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesServiceConnection struct {
	value *KubernetesServiceConnection
	isSet bool
}

func (v NullableKubernetesServiceConnection) Get() *KubernetesServiceConnection {
	return v.value
}

func (v *NullableKubernetesServiceConnection) Set(val *KubernetesServiceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesServiceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesServiceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesServiceConnection(val *KubernetesServiceConnection) *NullableKubernetesServiceConnection {
	return &NullableKubernetesServiceConnection{value: val, isSet: true}
}

func (v NullableKubernetesServiceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesServiceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
