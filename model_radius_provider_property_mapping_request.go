/*
authentik

Making authentication simple.

API version: 2024.8.1
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// RadiusProviderPropertyMappingRequest RadiusProviderPropertyMapping Serializer
type RadiusProviderPropertyMappingRequest struct {
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed    NullableString `json:"managed,omitempty"`
	Name       string         `json:"name"`
	Expression string         `json:"expression"`
}

// NewRadiusProviderPropertyMappingRequest instantiates a new RadiusProviderPropertyMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadiusProviderPropertyMappingRequest(name string, expression string) *RadiusProviderPropertyMappingRequest {
	this := RadiusProviderPropertyMappingRequest{}
	this.Name = name
	this.Expression = expression
	return &this
}

// NewRadiusProviderPropertyMappingRequestWithDefaults instantiates a new RadiusProviderPropertyMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadiusProviderPropertyMappingRequestWithDefaults() *RadiusProviderPropertyMappingRequest {
	this := RadiusProviderPropertyMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RadiusProviderPropertyMappingRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RadiusProviderPropertyMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *RadiusProviderPropertyMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *RadiusProviderPropertyMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *RadiusProviderPropertyMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *RadiusProviderPropertyMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *RadiusProviderPropertyMappingRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RadiusProviderPropertyMappingRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RadiusProviderPropertyMappingRequest) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *RadiusProviderPropertyMappingRequest) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *RadiusProviderPropertyMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *RadiusProviderPropertyMappingRequest) SetExpression(v string) {
	o.Expression = v
}

func (o RadiusProviderPropertyMappingRequest) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableRadiusProviderPropertyMappingRequest struct {
	value *RadiusProviderPropertyMappingRequest
	isSet bool
}

func (v NullableRadiusProviderPropertyMappingRequest) Get() *RadiusProviderPropertyMappingRequest {
	return v.value
}

func (v *NullableRadiusProviderPropertyMappingRequest) Set(val *RadiusProviderPropertyMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRadiusProviderPropertyMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRadiusProviderPropertyMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadiusProviderPropertyMappingRequest(val *RadiusProviderPropertyMappingRequest) *NullableRadiusProviderPropertyMappingRequest {
	return &NullableRadiusProviderPropertyMappingRequest{value: val, isSet: true}
}

func (v NullableRadiusProviderPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadiusProviderPropertyMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
