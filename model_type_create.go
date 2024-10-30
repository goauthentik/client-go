/*
authentik

Making authentication simple.

API version: 2024.10.0
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TypeCreate Types of an object that can be created
type TypeCreate struct {
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	Component          string  `json:"component"`
	ModelName          string  `json:"model_name"`
	IconUrl            *string `json:"icon_url,omitempty"`
	RequiresEnterprise *bool   `json:"requires_enterprise,omitempty"`
}

// NewTypeCreate instantiates a new TypeCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypeCreate(name string, description string, component string, modelName string) *TypeCreate {
	this := TypeCreate{}
	this.Name = name
	this.Description = description
	this.Component = component
	this.ModelName = modelName
	var requiresEnterprise bool = false
	this.RequiresEnterprise = &requiresEnterprise
	return &this
}

// NewTypeCreateWithDefaults instantiates a new TypeCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypeCreateWithDefaults() *TypeCreate {
	this := TypeCreate{}
	var requiresEnterprise bool = false
	this.RequiresEnterprise = &requiresEnterprise
	return &this
}

// GetName returns the Name field value
func (o *TypeCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TypeCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TypeCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *TypeCreate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TypeCreate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TypeCreate) SetDescription(v string) {
	o.Description = v
}

// GetComponent returns the Component field value
func (o *TypeCreate) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *TypeCreate) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *TypeCreate) SetComponent(v string) {
	o.Component = v
}

// GetModelName returns the ModelName field value
func (o *TypeCreate) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *TypeCreate) GetModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *TypeCreate) SetModelName(v string) {
	o.ModelName = v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *TypeCreate) GetIconUrl() string {
	if o == nil || o.IconUrl == nil {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeCreate) GetIconUrlOk() (*string, bool) {
	if o == nil || o.IconUrl == nil {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *TypeCreate) HasIconUrl() bool {
	if o != nil && o.IconUrl != nil {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *TypeCreate) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetRequiresEnterprise returns the RequiresEnterprise field value if set, zero value otherwise.
func (o *TypeCreate) GetRequiresEnterprise() bool {
	if o == nil || o.RequiresEnterprise == nil {
		var ret bool
		return ret
	}
	return *o.RequiresEnterprise
}

// GetRequiresEnterpriseOk returns a tuple with the RequiresEnterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypeCreate) GetRequiresEnterpriseOk() (*bool, bool) {
	if o == nil || o.RequiresEnterprise == nil {
		return nil, false
	}
	return o.RequiresEnterprise, true
}

// HasRequiresEnterprise returns a boolean if a field has been set.
func (o *TypeCreate) HasRequiresEnterprise() bool {
	if o != nil && o.RequiresEnterprise != nil {
		return true
	}

	return false
}

// SetRequiresEnterprise gets a reference to the given bool and assigns it to the RequiresEnterprise field.
func (o *TypeCreate) SetRequiresEnterprise(v bool) {
	o.RequiresEnterprise = &v
}

func (o TypeCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["model_name"] = o.ModelName
	}
	if o.IconUrl != nil {
		toSerialize["icon_url"] = o.IconUrl
	}
	if o.RequiresEnterprise != nil {
		toSerialize["requires_enterprise"] = o.RequiresEnterprise
	}
	return json.Marshal(toSerialize)
}

type NullableTypeCreate struct {
	value *TypeCreate
	isSet bool
}

func (v NullableTypeCreate) Get() *TypeCreate {
	return v.value
}

func (v *NullableTypeCreate) Set(val *TypeCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeCreate(val *TypeCreate) *NullableTypeCreate {
	return &NullableTypeCreate{value: val, isSet: true}
}

func (v NullableTypeCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
