/*
authentik

Making authentication simple.

API version: 2024.10.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// EndpointDeviceRequest Serializer for Endpoint authenticator devices
type EndpointDeviceRequest struct {
	Pk *string `json:"pk,omitempty"`
	// The human-readable name of this device.
	Name string `json:"name"`
}

// NewEndpointDeviceRequest instantiates a new EndpointDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointDeviceRequest(name string) *EndpointDeviceRequest {
	this := EndpointDeviceRequest{}
	this.Name = name
	return &this
}

// NewEndpointDeviceRequestWithDefaults instantiates a new EndpointDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointDeviceRequestWithDefaults() *EndpointDeviceRequest {
	this := EndpointDeviceRequest{}
	return &this
}

// GetPk returns the Pk field value if set, zero value otherwise.
func (o *EndpointDeviceRequest) GetPk() string {
	if o == nil || o.Pk == nil {
		var ret string
		return ret
	}
	return *o.Pk
}

// GetPkOk returns a tuple with the Pk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointDeviceRequest) GetPkOk() (*string, bool) {
	if o == nil || o.Pk == nil {
		return nil, false
	}
	return o.Pk, true
}

// HasPk returns a boolean if a field has been set.
func (o *EndpointDeviceRequest) HasPk() bool {
	if o != nil && o.Pk != nil {
		return true
	}

	return false
}

// SetPk gets a reference to the given string and assigns it to the Pk field.
func (o *EndpointDeviceRequest) SetPk(v string) {
	o.Pk = &v
}

// GetName returns the Name field value
func (o *EndpointDeviceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EndpointDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EndpointDeviceRequest) SetName(v string) {
	o.Name = v
}

func (o EndpointDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pk != nil {
		toSerialize["pk"] = o.Pk
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableEndpointDeviceRequest struct {
	value *EndpointDeviceRequest
	isSet bool
}

func (v NullableEndpointDeviceRequest) Get() *EndpointDeviceRequest {
	return v.value
}

func (v *NullableEndpointDeviceRequest) Set(val *EndpointDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointDeviceRequest(val *EndpointDeviceRequest) *NullableEndpointDeviceRequest {
	return &NullableEndpointDeviceRequest{value: val, isSet: true}
}

func (v NullableEndpointDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
