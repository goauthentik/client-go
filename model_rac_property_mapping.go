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

// RACPropertyMapping RACPropertyMapping Serializer
type RACPropertyMapping struct {
	Pk string `json:"pk"`
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed    NullableString `json:"managed,omitempty"`
	Name       string         `json:"name"`
	Expression *string        `json:"expression,omitempty"`
	// Get object's component so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName  string                 `json:"meta_model_name"`
	StaticSettings map[string]interface{} `json:"static_settings"`
}

// NewRACPropertyMapping instantiates a new RACPropertyMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRACPropertyMapping(pk string, name string, component string, verboseName string, verboseNamePlural string, metaModelName string, staticSettings map[string]interface{}) *RACPropertyMapping {
	this := RACPropertyMapping{}
	this.Pk = pk
	this.Name = name
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.StaticSettings = staticSettings
	return &this
}

// NewRACPropertyMappingWithDefaults instantiates a new RACPropertyMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRACPropertyMappingWithDefaults() *RACPropertyMapping {
	this := RACPropertyMapping{}
	return &this
}

// GetPk returns the Pk field value
func (o *RACPropertyMapping) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *RACPropertyMapping) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *RACPropertyMapping) SetPk(v string) {
	o.Pk = v
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RACPropertyMapping) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RACPropertyMapping) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *RACPropertyMapping) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *RACPropertyMapping) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *RACPropertyMapping) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *RACPropertyMapping) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *RACPropertyMapping) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RACPropertyMapping) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RACPropertyMapping) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *RACPropertyMapping) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RACPropertyMapping) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *RACPropertyMapping) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *RACPropertyMapping) SetExpression(v string) {
	o.Expression = &v
}

// GetComponent returns the Component field value
func (o *RACPropertyMapping) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *RACPropertyMapping) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *RACPropertyMapping) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *RACPropertyMapping) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *RACPropertyMapping) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *RACPropertyMapping) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *RACPropertyMapping) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *RACPropertyMapping) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *RACPropertyMapping) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *RACPropertyMapping) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *RACPropertyMapping) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *RACPropertyMapping) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetStaticSettings returns the StaticSettings field value
func (o *RACPropertyMapping) GetStaticSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.StaticSettings
}

// GetStaticSettingsOk returns a tuple with the StaticSettings field value
// and a boolean to check if the value has been set.
func (o *RACPropertyMapping) GetStaticSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.StaticSettings, true
}

// SetStaticSettings sets field value
func (o *RACPropertyMapping) SetStaticSettings(v map[string]interface{}) {
	o.StaticSettings = v
}

func (o RACPropertyMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pk"] = o.Pk
	}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if true {
		toSerialize["component"] = o.Component
	}
	if true {
		toSerialize["verbose_name"] = o.VerboseName
	}
	if true {
		toSerialize["verbose_name_plural"] = o.VerboseNamePlural
	}
	if true {
		toSerialize["meta_model_name"] = o.MetaModelName
	}
	if true {
		toSerialize["static_settings"] = o.StaticSettings
	}
	return json.Marshal(toSerialize)
}

type NullableRACPropertyMapping struct {
	value *RACPropertyMapping
	isSet bool
}

func (v NullableRACPropertyMapping) Get() *RACPropertyMapping {
	return v.value
}

func (v *NullableRACPropertyMapping) Set(val *RACPropertyMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableRACPropertyMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableRACPropertyMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRACPropertyMapping(val *RACPropertyMapping) *NullableRACPropertyMapping {
	return &NullableRACPropertyMapping{value: val, isSet: true}
}

func (v NullableRACPropertyMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRACPropertyMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
