/*
authentik

Making authentication simple.

API version: 2023.10.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ServiceConnectionRequest ServiceConnection Serializer
type ServiceConnectionRequest struct {
	Name string `json:"name"`
	// If enabled, use the local connection. Required Docker socket/Kubernetes Integration
	Local *bool `json:"local,omitempty"`
}

// NewServiceConnectionRequest instantiates a new ServiceConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceConnectionRequest(name string) *ServiceConnectionRequest {
	this := ServiceConnectionRequest{}
	this.Name = name
	return &this
}

// NewServiceConnectionRequestWithDefaults instantiates a new ServiceConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceConnectionRequestWithDefaults() *ServiceConnectionRequest {
	this := ServiceConnectionRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ServiceConnectionRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceConnectionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceConnectionRequest) SetName(v string) {
	o.Name = v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *ServiceConnectionRequest) GetLocal() bool {
	if o == nil || o.Local == nil {
		var ret bool
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceConnectionRequest) GetLocalOk() (*bool, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *ServiceConnectionRequest) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given bool and assigns it to the Local field.
func (o *ServiceConnectionRequest) SetLocal(v bool) {
	o.Local = &v
}

func (o ServiceConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
	}
	return json.Marshal(toSerialize)
}

type NullableServiceConnectionRequest struct {
	value *ServiceConnectionRequest
	isSet bool
}

func (v NullableServiceConnectionRequest) Get() *ServiceConnectionRequest {
	return v.value
}

func (v *NullableServiceConnectionRequest) Set(val *ServiceConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceConnectionRequest(val *ServiceConnectionRequest) *NullableServiceConnectionRequest {
	return &NullableServiceConnectionRequest{value: val, isSet: true}
}

func (v NullableServiceConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
