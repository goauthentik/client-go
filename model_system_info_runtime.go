/*
authentik

Making authentication simple.

API version: 2025.2.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SystemInfoRuntime Get versions
type SystemInfoRuntime struct {
	PythonVersion      string       `json:"python_version"`
	Environment        string       `json:"environment"`
	Architecture       string       `json:"architecture"`
	Platform           string       `json:"platform"`
	Uname              string       `json:"uname"`
	OpensslVersion     string       `json:"openssl_version"`
	OpensslFipsEnabled NullableBool `json:"openssl_fips_enabled"`
	AuthentikVersion   string       `json:"authentik_version"`
}

// NewSystemInfoRuntime instantiates a new SystemInfoRuntime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInfoRuntime(pythonVersion string, environment string, architecture string, platform string, uname string, opensslVersion string, opensslFipsEnabled NullableBool, authentikVersion string) *SystemInfoRuntime {
	this := SystemInfoRuntime{}
	this.PythonVersion = pythonVersion
	this.Environment = environment
	this.Architecture = architecture
	this.Platform = platform
	this.Uname = uname
	this.OpensslVersion = opensslVersion
	this.OpensslFipsEnabled = opensslFipsEnabled
	this.AuthentikVersion = authentikVersion
	return &this
}

// NewSystemInfoRuntimeWithDefaults instantiates a new SystemInfoRuntime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInfoRuntimeWithDefaults() *SystemInfoRuntime {
	this := SystemInfoRuntime{}
	return &this
}

// GetPythonVersion returns the PythonVersion field value
func (o *SystemInfoRuntime) GetPythonVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PythonVersion
}

// GetPythonVersionOk returns a tuple with the PythonVersion field value
// and a boolean to check if the value has been set.
func (o *SystemInfoRuntime) GetPythonVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PythonVersion, true
}

// SetPythonVersion sets field value
func (o *SystemInfoRuntime) SetPythonVersion(v string) {
	o.PythonVersion = v
}

// GetEnvironment returns the Environment field value
func (o *SystemInfoRuntime) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *SystemInfoRuntime) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *SystemInfoRuntime) SetEnvironment(v string) {
	o.Environment = v
}

// GetArchitecture returns the Architecture field value
func (o *SystemInfoRuntime) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *SystemInfoRuntime) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *SystemInfoRuntime) SetArchitecture(v string) {
	o.Architecture = v
}

// GetPlatform returns the Platform field value
func (o *SystemInfoRuntime) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *SystemInfoRuntime) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *SystemInfoRuntime) SetPlatform(v string) {
	o.Platform = v
}

// GetUname returns the Uname field value
func (o *SystemInfoRuntime) GetUname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uname
}

// GetUnameOk returns a tuple with the Uname field value
// and a boolean to check if the value has been set.
func (o *SystemInfoRuntime) GetUnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uname, true
}

// SetUname sets field value
func (o *SystemInfoRuntime) SetUname(v string) {
	o.Uname = v
}

// GetOpensslVersion returns the OpensslVersion field value
func (o *SystemInfoRuntime) GetOpensslVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpensslVersion
}

// GetOpensslVersionOk returns a tuple with the OpensslVersion field value
// and a boolean to check if the value has been set.
func (o *SystemInfoRuntime) GetOpensslVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpensslVersion, true
}

// SetOpensslVersion sets field value
func (o *SystemInfoRuntime) SetOpensslVersion(v string) {
	o.OpensslVersion = v
}

// GetOpensslFipsEnabled returns the OpensslFipsEnabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *SystemInfoRuntime) GetOpensslFipsEnabled() bool {
	if o == nil || o.OpensslFipsEnabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.OpensslFipsEnabled.Get()
}

// GetOpensslFipsEnabledOk returns a tuple with the OpensslFipsEnabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemInfoRuntime) GetOpensslFipsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpensslFipsEnabled.Get(), o.OpensslFipsEnabled.IsSet()
}

// SetOpensslFipsEnabled sets field value
func (o *SystemInfoRuntime) SetOpensslFipsEnabled(v bool) {
	o.OpensslFipsEnabled.Set(&v)
}

// GetAuthentikVersion returns the AuthentikVersion field value
func (o *SystemInfoRuntime) GetAuthentikVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthentikVersion
}

// GetAuthentikVersionOk returns a tuple with the AuthentikVersion field value
// and a boolean to check if the value has been set.
func (o *SystemInfoRuntime) GetAuthentikVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthentikVersion, true
}

// SetAuthentikVersion sets field value
func (o *SystemInfoRuntime) SetAuthentikVersion(v string) {
	o.AuthentikVersion = v
}

func (o SystemInfoRuntime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["python_version"] = o.PythonVersion
	}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["architecture"] = o.Architecture
	}
	if true {
		toSerialize["platform"] = o.Platform
	}
	if true {
		toSerialize["uname"] = o.Uname
	}
	if true {
		toSerialize["openssl_version"] = o.OpensslVersion
	}
	if true {
		toSerialize["openssl_fips_enabled"] = o.OpensslFipsEnabled.Get()
	}
	if true {
		toSerialize["authentik_version"] = o.AuthentikVersion
	}
	return json.Marshal(toSerialize)
}

type NullableSystemInfoRuntime struct {
	value *SystemInfoRuntime
	isSet bool
}

func (v NullableSystemInfoRuntime) Get() *SystemInfoRuntime {
	return v.value
}

func (v *NullableSystemInfoRuntime) Set(val *SystemInfoRuntime) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInfoRuntime) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInfoRuntime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInfoRuntime(val *SystemInfoRuntime) *NullableSystemInfoRuntime {
	return &NullableSystemInfoRuntime{value: val, isSet: true}
}

func (v NullableSystemInfoRuntime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInfoRuntime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
