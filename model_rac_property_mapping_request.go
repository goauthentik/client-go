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

// RACPropertyMappingRequest RACPropertyMapping Serializer
type RACPropertyMappingRequest struct {
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed        NullableString         `json:"managed,omitempty"`
	Name           string                 `json:"name"`
	Expression     *string                `json:"expression,omitempty"`
	StaticSettings map[string]interface{} `json:"static_settings"`
}

// NewRACPropertyMappingRequest instantiates a new RACPropertyMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRACPropertyMappingRequest(name string, staticSettings map[string]interface{}) *RACPropertyMappingRequest {
	this := RACPropertyMappingRequest{}
	this.Name = name
	this.StaticSettings = staticSettings
	return &this
}

// NewRACPropertyMappingRequestWithDefaults instantiates a new RACPropertyMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRACPropertyMappingRequestWithDefaults() *RACPropertyMappingRequest {
	this := RACPropertyMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RACPropertyMappingRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RACPropertyMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *RACPropertyMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *RACPropertyMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *RACPropertyMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *RACPropertyMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *RACPropertyMappingRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RACPropertyMappingRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RACPropertyMappingRequest) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *RACPropertyMappingRequest) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RACPropertyMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *RACPropertyMappingRequest) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *RACPropertyMappingRequest) SetExpression(v string) {
	o.Expression = &v
}

// GetStaticSettings returns the StaticSettings field value
func (o *RACPropertyMappingRequest) GetStaticSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.StaticSettings
}

// GetStaticSettingsOk returns a tuple with the StaticSettings field value
// and a boolean to check if the value has been set.
func (o *RACPropertyMappingRequest) GetStaticSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.StaticSettings, true
}

// SetStaticSettings sets field value
func (o *RACPropertyMappingRequest) SetStaticSettings(v map[string]interface{}) {
	o.StaticSettings = v
}

func (o RACPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["static_settings"] = o.StaticSettings
	}
	return json.Marshal(toSerialize)
}

type NullableRACPropertyMappingRequest struct {
	value *RACPropertyMappingRequest
	isSet bool
}

func (v NullableRACPropertyMappingRequest) Get() *RACPropertyMappingRequest {
	return v.value
}

func (v *NullableRACPropertyMappingRequest) Set(val *RACPropertyMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRACPropertyMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRACPropertyMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRACPropertyMappingRequest(val *RACPropertyMappingRequest) *NullableRACPropertyMappingRequest {
	return &NullableRACPropertyMappingRequest{value: val, isSet: true}
}

func (v NullableRACPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRACPropertyMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
