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

// ServiceConnectionState Serializer for Service connection state
type ServiceConnectionState struct {
	Healthy bool   `json:"healthy"`
	Version string `json:"version"`
}

// NewServiceConnectionState instantiates a new ServiceConnectionState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceConnectionState(healthy bool, version string) *ServiceConnectionState {
	this := ServiceConnectionState{}
	this.Healthy = healthy
	this.Version = version
	return &this
}

// NewServiceConnectionStateWithDefaults instantiates a new ServiceConnectionState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceConnectionStateWithDefaults() *ServiceConnectionState {
	this := ServiceConnectionState{}
	return &this
}

// GetHealthy returns the Healthy field value
func (o *ServiceConnectionState) GetHealthy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Healthy
}

// GetHealthyOk returns a tuple with the Healthy field value
// and a boolean to check if the value has been set.
func (o *ServiceConnectionState) GetHealthyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Healthy, true
}

// SetHealthy sets field value
func (o *ServiceConnectionState) SetHealthy(v bool) {
	o.Healthy = v
}

// GetVersion returns the Version field value
func (o *ServiceConnectionState) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ServiceConnectionState) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ServiceConnectionState) SetVersion(v string) {
	o.Version = v
}

func (o ServiceConnectionState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["healthy"] = o.Healthy
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableServiceConnectionState struct {
	value *ServiceConnectionState
	isSet bool
}

func (v NullableServiceConnectionState) Get() *ServiceConnectionState {
	return v.value
}

func (v *NullableServiceConnectionState) Set(val *ServiceConnectionState) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceConnectionState) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceConnectionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceConnectionState(val *ServiceConnectionState) *NullableServiceConnectionState {
	return &NullableServiceConnectionState{value: val, isSet: true}
}

func (v NullableServiceConnectionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceConnectionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
