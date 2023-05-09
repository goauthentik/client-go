/*
authentik

Making authentication simple.

API version: 2023.4.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LDAPPropertyMappingRequest LDAP PropertyMapping Serializer
type LDAPPropertyMappingRequest struct {
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed     NullableString `json:"managed,omitempty"`
	Name        string         `json:"name"`
	Expression  string         `json:"expression"`
	ObjectField string         `json:"object_field"`
}

// NewLDAPPropertyMappingRequest instantiates a new LDAPPropertyMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPPropertyMappingRequest(name string, expression string, objectField string) *LDAPPropertyMappingRequest {
	this := LDAPPropertyMappingRequest{}
	this.Name = name
	this.Expression = expression
	this.ObjectField = objectField
	return &this
}

// NewLDAPPropertyMappingRequestWithDefaults instantiates a new LDAPPropertyMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPPropertyMappingRequestWithDefaults() *LDAPPropertyMappingRequest {
	this := LDAPPropertyMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPPropertyMappingRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPPropertyMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *LDAPPropertyMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *LDAPPropertyMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *LDAPPropertyMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *LDAPPropertyMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *LDAPPropertyMappingRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LDAPPropertyMappingRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LDAPPropertyMappingRequest) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *LDAPPropertyMappingRequest) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *LDAPPropertyMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *LDAPPropertyMappingRequest) SetExpression(v string) {
	o.Expression = v
}

// GetObjectField returns the ObjectField field value
func (o *LDAPPropertyMappingRequest) GetObjectField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectField
}

// GetObjectFieldOk returns a tuple with the ObjectField field value
// and a boolean to check if the value has been set.
func (o *LDAPPropertyMappingRequest) GetObjectFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectField, true
}

// SetObjectField sets field value
func (o *LDAPPropertyMappingRequest) SetObjectField(v string) {
	o.ObjectField = v
}

func (o LDAPPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["object_field"] = o.ObjectField
	}
	return json.Marshal(toSerialize)
}

type NullableLDAPPropertyMappingRequest struct {
	value *LDAPPropertyMappingRequest
	isSet bool
}

func (v NullableLDAPPropertyMappingRequest) Get() *LDAPPropertyMappingRequest {
	return v.value
}

func (v *NullableLDAPPropertyMappingRequest) Set(val *LDAPPropertyMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPPropertyMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPPropertyMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPPropertyMappingRequest(val *LDAPPropertyMappingRequest) *NullableLDAPPropertyMappingRequest {
	return &NullableLDAPPropertyMappingRequest{value: val, isSet: true}
}

func (v NullableLDAPPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPPropertyMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
