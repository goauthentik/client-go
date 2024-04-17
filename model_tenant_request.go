/*
authentik

Making authentication simple.

API version: 2024.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TenantRequest Tenant Serializer
type TenantRequest struct {
	SchemaName string `json:"schema_name"`
	Name       string `json:"name"`
	Ready      *bool  `json:"ready,omitempty"`
}

// NewTenantRequest instantiates a new TenantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantRequest(schemaName string, name string) *TenantRequest {
	this := TenantRequest{}
	this.SchemaName = schemaName
	this.Name = name
	return &this
}

// NewTenantRequestWithDefaults instantiates a new TenantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantRequestWithDefaults() *TenantRequest {
	this := TenantRequest{}
	return &this
}

// GetSchemaName returns the SchemaName field value
func (o *TenantRequest) GetSchemaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value
// and a boolean to check if the value has been set.
func (o *TenantRequest) GetSchemaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaName, true
}

// SetSchemaName sets field value
func (o *TenantRequest) SetSchemaName(v string) {
	o.SchemaName = v
}

// GetName returns the Name field value
func (o *TenantRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TenantRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TenantRequest) SetName(v string) {
	o.Name = v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *TenantRequest) GetReady() bool {
	if o == nil || o.Ready == nil {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantRequest) GetReadyOk() (*bool, bool) {
	if o == nil || o.Ready == nil {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *TenantRequest) HasReady() bool {
	if o != nil && o.Ready != nil {
		return true
	}

	return false
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *TenantRequest) SetReady(v bool) {
	o.Ready = &v
}

func (o TenantRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schema_name"] = o.SchemaName
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Ready != nil {
		toSerialize["ready"] = o.Ready
	}
	return json.Marshal(toSerialize)
}

type NullableTenantRequest struct {
	value *TenantRequest
	isSet bool
}

func (v NullableTenantRequest) Get() *TenantRequest {
	return v.value
}

func (v *NullableTenantRequest) Set(val *TenantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantRequest(val *TenantRequest) *NullableTenantRequest {
	return &NullableTenantRequest{value: val, isSet: true}
}

func (v NullableTenantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
