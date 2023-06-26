/*
authentik

Making authentication simple.

API version: 2023.5.4
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SystemRuntime Get versions
type SystemRuntime struct {
	PythonVersion   string `json:"python_version"`
	GunicornVersion string `json:"gunicorn_version"`
	Environment     string `json:"environment"`
	Architecture    string `json:"architecture"`
	Platform        string `json:"platform"`
	Uname           string `json:"uname"`
}

// NewSystemRuntime instantiates a new SystemRuntime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemRuntime(pythonVersion string, gunicornVersion string, environment string, architecture string, platform string, uname string) *SystemRuntime {
	this := SystemRuntime{}
	this.PythonVersion = pythonVersion
	this.GunicornVersion = gunicornVersion
	this.Environment = environment
	this.Architecture = architecture
	this.Platform = platform
	this.Uname = uname
	return &this
}

// NewSystemRuntimeWithDefaults instantiates a new SystemRuntime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemRuntimeWithDefaults() *SystemRuntime {
	this := SystemRuntime{}
	return &this
}

// GetPythonVersion returns the PythonVersion field value
func (o *SystemRuntime) GetPythonVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PythonVersion
}

// GetPythonVersionOk returns a tuple with the PythonVersion field value
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetPythonVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PythonVersion, true
}

// SetPythonVersion sets field value
func (o *SystemRuntime) SetPythonVersion(v string) {
	o.PythonVersion = v
}

// GetGunicornVersion returns the GunicornVersion field value
func (o *SystemRuntime) GetGunicornVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GunicornVersion
}

// GetGunicornVersionOk returns a tuple with the GunicornVersion field value
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetGunicornVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GunicornVersion, true
}

// SetGunicornVersion sets field value
func (o *SystemRuntime) SetGunicornVersion(v string) {
	o.GunicornVersion = v
}

// GetEnvironment returns the Environment field value
func (o *SystemRuntime) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *SystemRuntime) SetEnvironment(v string) {
	o.Environment = v
}

// GetArchitecture returns the Architecture field value
func (o *SystemRuntime) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *SystemRuntime) SetArchitecture(v string) {
	o.Architecture = v
}

// GetPlatform returns the Platform field value
func (o *SystemRuntime) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *SystemRuntime) SetPlatform(v string) {
	o.Platform = v
}

// GetUname returns the Uname field value
func (o *SystemRuntime) GetUname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uname
}

// GetUnameOk returns a tuple with the Uname field value
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetUnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uname, true
}

// SetUname sets field value
func (o *SystemRuntime) SetUname(v string) {
	o.Uname = v
}

func (o SystemRuntime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["python_version"] = o.PythonVersion
	}
	if true {
		toSerialize["gunicorn_version"] = o.GunicornVersion
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
	return json.Marshal(toSerialize)
}

type NullableSystemRuntime struct {
	value *SystemRuntime
	isSet bool
}

func (v NullableSystemRuntime) Get() *SystemRuntime {
	return v.value
}

func (v *NullableSystemRuntime) Set(val *SystemRuntime) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemRuntime) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemRuntime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemRuntime(val *SystemRuntime) *NullableSystemRuntime {
	return &NullableSystemRuntime{value: val, isSet: true}
}

func (v NullableSystemRuntime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemRuntime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
