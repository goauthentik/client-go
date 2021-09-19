/*
authentik

Making authentication simple.

API version: 2021.9.1-rc3
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SystemRuntime struct for SystemRuntime
type SystemRuntime struct {
	PythonVersion *string `json:"python_version,omitempty"`
	GunicornVersion *string `json:"gunicorn_version,omitempty"`
	Environment *string `json:"environment,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Uname *string `json:"uname,omitempty"`
}

// NewSystemRuntime instantiates a new SystemRuntime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemRuntime() *SystemRuntime {
	this := SystemRuntime{}
	return &this
}

// NewSystemRuntimeWithDefaults instantiates a new SystemRuntime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemRuntimeWithDefaults() *SystemRuntime {
	this := SystemRuntime{}
	return &this
}

// GetPythonVersion returns the PythonVersion field value if set, zero value otherwise.
func (o *SystemRuntime) GetPythonVersion() string {
	if o == nil || o.PythonVersion == nil {
		var ret string
		return ret
	}
	return *o.PythonVersion
}

// GetPythonVersionOk returns a tuple with the PythonVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetPythonVersionOk() (*string, bool) {
	if o == nil || o.PythonVersion == nil {
		return nil, false
	}
	return o.PythonVersion, true
}

// HasPythonVersion returns a boolean if a field has been set.
func (o *SystemRuntime) HasPythonVersion() bool {
	if o != nil && o.PythonVersion != nil {
		return true
	}

	return false
}

// SetPythonVersion gets a reference to the given string and assigns it to the PythonVersion field.
func (o *SystemRuntime) SetPythonVersion(v string) {
	o.PythonVersion = &v
}

// GetGunicornVersion returns the GunicornVersion field value if set, zero value otherwise.
func (o *SystemRuntime) GetGunicornVersion() string {
	if o == nil || o.GunicornVersion == nil {
		var ret string
		return ret
	}
	return *o.GunicornVersion
}

// GetGunicornVersionOk returns a tuple with the GunicornVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetGunicornVersionOk() (*string, bool) {
	if o == nil || o.GunicornVersion == nil {
		return nil, false
	}
	return o.GunicornVersion, true
}

// HasGunicornVersion returns a boolean if a field has been set.
func (o *SystemRuntime) HasGunicornVersion() bool {
	if o != nil && o.GunicornVersion != nil {
		return true
	}

	return false
}

// SetGunicornVersion gets a reference to the given string and assigns it to the GunicornVersion field.
func (o *SystemRuntime) SetGunicornVersion(v string) {
	o.GunicornVersion = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SystemRuntime) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SystemRuntime) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *SystemRuntime) SetEnvironment(v string) {
	o.Environment = &v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *SystemRuntime) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetArchitectureOk() (*string, bool) {
	if o == nil || o.Architecture == nil {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *SystemRuntime) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *SystemRuntime) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *SystemRuntime) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *SystemRuntime) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *SystemRuntime) SetPlatform(v string) {
	o.Platform = &v
}

// GetUname returns the Uname field value if set, zero value otherwise.
func (o *SystemRuntime) GetUname() string {
	if o == nil || o.Uname == nil {
		var ret string
		return ret
	}
	return *o.Uname
}

// GetUnameOk returns a tuple with the Uname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRuntime) GetUnameOk() (*string, bool) {
	if o == nil || o.Uname == nil {
		return nil, false
	}
	return o.Uname, true
}

// HasUname returns a boolean if a field has been set.
func (o *SystemRuntime) HasUname() bool {
	if o != nil && o.Uname != nil {
		return true
	}

	return false
}

// SetUname gets a reference to the given string and assigns it to the Uname field.
func (o *SystemRuntime) SetUname(v string) {
	o.Uname = &v
}

func (o SystemRuntime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PythonVersion != nil {
		toSerialize["python_version"] = o.PythonVersion
	}
	if o.GunicornVersion != nil {
		toSerialize["gunicorn_version"] = o.GunicornVersion
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Architecture != nil {
		toSerialize["architecture"] = o.Architecture
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Uname != nil {
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


