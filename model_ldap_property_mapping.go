/*
authentik

Making authentication simple.

API version: 2022.6.2
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LDAPPropertyMapping LDAP PropertyMapping Serializer
type LDAPPropertyMapping struct {
	Pk string `json:"pk"`
	// Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed           NullableString `json:"managed,omitempty"`
	Name              string         `json:"name"`
	Expression        string         `json:"expression"`
	Component         string         `json:"component"`
	VerboseName       string         `json:"verbose_name"`
	VerboseNamePlural string         `json:"verbose_name_plural"`
	MetaModelName     string         `json:"meta_model_name"`
	ObjectField       string         `json:"object_field"`
}

// NewLDAPPropertyMapping instantiates a new LDAPPropertyMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPPropertyMapping(pk string, name string, expression string, component string, verboseName string, verboseNamePlural string, metaModelName string, objectField string) *LDAPPropertyMapping {
	this := LDAPPropertyMapping{}
	this.Pk = pk
	this.Name = name
	this.Expression = expression
	this.Component = component
	this.VerboseName = verboseName
	this.VerboseNamePlural = verboseNamePlural
	this.MetaModelName = metaModelName
	this.ObjectField = objectField
	return &this
}

// NewLDAPPropertyMappingWithDefaults instantiates a new LDAPPropertyMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPPropertyMappingWithDefaults() *LDAPPropertyMapping {
	this := LDAPPropertyMapping{}
	return &this
}

// GetPk returns the Pk field value
func (o *LDAPPropertyMapping) GetPk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pk
}

// GetPkOk returns a tuple with the Pk field value
// and a boolean to check if the value has been set.
func (o *LDAPPropertyMapping) GetPkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pk, true
}

// SetPk sets field value
func (o *LDAPPropertyMapping) SetPk(v string) {
	o.Pk = v
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPPropertyMapping) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPPropertyMapping) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *LDAPPropertyMapping) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *LDAPPropertyMapping) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *LDAPPropertyMapping) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *LDAPPropertyMapping) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *LDAPPropertyMapping) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LDAPPropertyMapping) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LDAPPropertyMapping) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *LDAPPropertyMapping) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *LDAPPropertyMapping) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *LDAPPropertyMapping) SetExpression(v string) {
	o.Expression = v
}

// GetComponent returns the Component field value
func (o *LDAPPropertyMapping) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *LDAPPropertyMapping) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *LDAPPropertyMapping) SetComponent(v string) {
	o.Component = v
}

// GetVerboseName returns the VerboseName field value
func (o *LDAPPropertyMapping) GetVerboseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseName
}

// GetVerboseNameOk returns a tuple with the VerboseName field value
// and a boolean to check if the value has been set.
func (o *LDAPPropertyMapping) GetVerboseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseName, true
}

// SetVerboseName sets field value
func (o *LDAPPropertyMapping) SetVerboseName(v string) {
	o.VerboseName = v
}

// GetVerboseNamePlural returns the VerboseNamePlural field value
func (o *LDAPPropertyMapping) GetVerboseNamePlural() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerboseNamePlural
}

// GetVerboseNamePluralOk returns a tuple with the VerboseNamePlural field value
// and a boolean to check if the value has been set.
func (o *LDAPPropertyMapping) GetVerboseNamePluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerboseNamePlural, true
}

// SetVerboseNamePlural sets field value
func (o *LDAPPropertyMapping) SetVerboseNamePlural(v string) {
	o.VerboseNamePlural = v
}

// GetMetaModelName returns the MetaModelName field value
func (o *LDAPPropertyMapping) GetMetaModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetaModelName
}

// GetMetaModelNameOk returns a tuple with the MetaModelName field value
// and a boolean to check if the value has been set.
func (o *LDAPPropertyMapping) GetMetaModelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetaModelName, true
}

// SetMetaModelName sets field value
func (o *LDAPPropertyMapping) SetMetaModelName(v string) {
	o.MetaModelName = v
}

// GetObjectField returns the ObjectField field value
func (o *LDAPPropertyMapping) GetObjectField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectField
}

// GetObjectFieldOk returns a tuple with the ObjectField field value
// and a boolean to check if the value has been set.
func (o *LDAPPropertyMapping) GetObjectFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectField, true
}

// SetObjectField sets field value
func (o *LDAPPropertyMapping) SetObjectField(v string) {
	o.ObjectField = v
}

func (o LDAPPropertyMapping) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["object_field"] = o.ObjectField
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPPropertyMapping struct {
	value *LDAPPropertyMapping
	isSet bool
}

func (v NullableLDAPPropertyMapping) Get() *LDAPPropertyMapping {
	return v.value
}

func (v *NullableLDAPPropertyMapping) Set(val *LDAPPropertyMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPPropertyMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPPropertyMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPPropertyMapping(val *LDAPPropertyMapping) *NullableLDAPPropertyMapping {
	return &NullableLDAPPropertyMapping{value: val, isSet: true}
}

func (v NullableLDAPPropertyMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPPropertyMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
