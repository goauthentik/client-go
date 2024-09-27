/*
authentik

Making authentication simple.

API version: 2024.8.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Tenant Tenant Serializer
type Tenant struct {
	TenantUuid string `json:"tenant_uuid"`
	SchemaName string `json:"schema_name"`
	Name       string `json:"name"`
	Ready      *bool  `json:"ready,omitempty"`
}

// NewTenant instantiates a new Tenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenant(tenantUuid string, schemaName string, name string) *Tenant {
	this := Tenant{}
	this.TenantUuid = tenantUuid
	this.SchemaName = schemaName
	this.Name = name
	return &this
}

// NewTenantWithDefaults instantiates a new Tenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWithDefaults() *Tenant {
	this := Tenant{}
	return &this
}

// GetTenantUuid returns the TenantUuid field value
func (o *Tenant) GetTenantUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetTenantUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantUuid, true
}

// SetTenantUuid sets field value
func (o *Tenant) SetTenantUuid(v string) {
	o.TenantUuid = v
}

// GetSchemaName returns the SchemaName field value
func (o *Tenant) GetSchemaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetSchemaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaName, true
}

// SetSchemaName sets field value
func (o *Tenant) SetSchemaName(v string) {
	o.SchemaName = v
}

// GetName returns the Name field value
func (o *Tenant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tenant) SetName(v string) {
	o.Name = v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *Tenant) GetReady() bool {
	if o == nil || o.Ready == nil {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetReadyOk() (*bool, bool) {
	if o == nil || o.Ready == nil {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *Tenant) HasReady() bool {
	if o != nil && o.Ready != nil {
		return true
	}

	return false
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *Tenant) SetReady(v bool) {
	o.Ready = &v
}

func (o Tenant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
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

type NullableTenant struct {
	value *Tenant
	isSet bool
}

func (v NullableTenant) Get() *Tenant {
	return v.value
}

func (v *NullableTenant) Set(val *Tenant) {
	v.value = val
	v.isSet = true
}

func (v NullableTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenant(val *Tenant) *NullableTenant {
	return &NullableTenant{value: val, isSet: true}
}

func (v NullableTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
