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

// KerberosSourcePropertyMapping Kerberos PropertyMapping Serializer
type KerberosSourcePropertyMapping struct {
	Pk string `json:"pk"`
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed    NullableString `json:"managed,omitempty"`
	Name       string         `json:"name"`
	Expression string         `json:"expression"`
	// Get object's component so that we know how to edit the object
	Component string `json:"component"`
	// Return object's verbose_name
	VerboseName string `json:"verbose_name"`
	// Return object's plural verbose_name
	VerboseNamePlural string `json:"verbose_name_plural"`
	// Return internal model name
	MetaModelName string `json:"meta_model_name"`
}

// NewKerberosSourcePropertyMapping instantiates a new KerberosSourcePropertyMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosSourcePropertyMapping(pk string, name string, expression string, component string, verboseName string, verboseNamePlural string, metaModelName string) *KerberosSourcePropertyMapping {
	this := KerberosSourcePropertyMapping{}
	this.Pk = pk
	this.Name = name
	this.Expression = expression
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	return &this
}

// NewKerberosSourcePropertyMappingWithDefaults instantiates a new KerberosSourcePropertyMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosSourcePropertyMappingWithDefaults() *KerberosSourcePropertyMapping {
	this := KerberosSourcePropertyMapping{}
	return &this
}

// GetPk returns the Pk field value
func (o *KerberosSourcePropertyMapping) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *KerberosSourcePropertyMapping) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *KerberosSourcePropertyMapping) SetPk(v string) {
	o.Pk = v
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KerberosSourcePropertyMapping) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KerberosSourcePropertyMapping) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *KerberosSourcePropertyMapping) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *KerberosSourcePropertyMapping) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *KerberosSourcePropertyMapping) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *KerberosSourcePropertyMapping) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *KerberosSourcePropertyMapping) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KerberosSourcePropertyMapping) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KerberosSourcePropertyMapping) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *KerberosSourcePropertyMapping) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *KerberosSourcePropertyMapping) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *KerberosSourcePropertyMapping) SetExpression(v string) {
	o.Expression = v
}

// GetComponent returns the Component field value
func (o *KerberosSourcePropertyMapping) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *KerberosSourcePropertyMapping) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *KerberosSourcePropertyMapping) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *KerberosSourcePropertyMapping) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *KerberosSourcePropertyMapping) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *KerberosSourcePropertyMapping) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *KerberosSourcePropertyMapping) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *KerberosSourcePropertyMapping) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *KerberosSourcePropertyMapping) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *KerberosSourcePropertyMapping) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *KerberosSourcePropertyMapping) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *KerberosSourcePropertyMapping) SetMetaModelName(v string) {
	o.MetaModelName = v
}

func (o KerberosSourcePropertyMapping) MarshalJSON() ([]byte, error) {
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
	if true {
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
	return json.Marshal(toSerialize)
}

type NullableKerberosSourcePropertyMapping struct {
	value *KerberosSourcePropertyMapping
	isSet bool
}

func (v NullableKerberosSourcePropertyMapping) Get() *KerberosSourcePropertyMapping {
	return v.value
}

func (v *NullableKerberosSourcePropertyMapping) Set(val *KerberosSourcePropertyMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosSourcePropertyMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosSourcePropertyMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosSourcePropertyMapping(val *KerberosSourcePropertyMapping) *NullableKerberosSourcePropertyMapping {
	return &NullableKerberosSourcePropertyMapping{value: val, isSet: true}
}

func (v NullableKerberosSourcePropertyMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosSourcePropertyMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
