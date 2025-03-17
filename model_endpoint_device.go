/*
authentik

Making authentication simple.

API version: 2025.2.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EndpointDevice Serializer for Endpoint authenticator devices
type EndpointDevice struct {
	Pk *string `json:"pk,omitempty"`
	// The human-readable name of this device.
	Name string `json:"name"`
}

// NewEndpointDevice instantiates a new EndpointDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointDevice(name string) *EndpointDevice {
	this := EndpointDevice{}
	this.Name = name
	return &this
}

// NewEndpointDeviceWithDefaults instantiates a new EndpointDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointDeviceWithDefaults() *EndpointDevice {
	this := EndpointDevice{}
	return &this
}

// GetPk returns the Pk field value if set, zero value otherwise.
func (o *EndpointDevice) GetPk() string {
	if o == nil || o.Pk == nil {
		var ret string
		return ret
	}
	return *o.Pk
}

// GetPkOk returns a tuple with the Pk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointDevice) GetPkOk() (*string, bool) {
	if o == nil || o.Pk == nil {
		return nil, false
	}
	return o.Pk, true
}

// HasPk returns a boolean if a field has been set.
func (o *EndpointDevice) HasPk() bool {
	if o != nil && o.Pk != nil {
		return true
	}

	return false
}

// SetPk gets a reference to the given string and assigns it to the Pk field.
func (o *EndpointDevice) SetPk(v string) {
	o.Pk = &v
}

// GetName returns the Name field value
func (o *EndpointDevice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EndpointDevice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EndpointDevice) SetName(v string) {
	o.Name = v
}

func (o EndpointDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pk != nil {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointDevice struct {
	value *EndpointDevice
	isSet bool
}

func (v NullableEndpointDevice) Get() *EndpointDevice {
	return v.value
}

func (v *NullableEndpointDevice) Set(val *EndpointDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointDevice(val *EndpointDevice) *NullableEndpointDevice {
	return &NullableEndpointDevice{value: val, isSet: true}
}

func (v NullableEndpointDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
