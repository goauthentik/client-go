/*
authentik

Making authentication simple.

API version: 2025.2.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ScopeMappingRequest ScopeMapping Serializer
type ScopeMappingRequest struct {
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed    NullableString `json:"managed,omitempty"`
	Name       string         `json:"name"`
	Expression string         `json:"expression"`
	// Scope name requested by the client
	ScopeName string `json:"scope_name"`
	// Description shown to the user when consenting. If left empty, the user won't be informed.
	Description *string `json:"description,omitempty"`
}

// NewScopeMappingRequest instantiates a new ScopeMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeMappingRequest(name string, expression string, scopeName string) *ScopeMappingRequest {
	this := ScopeMappingRequest{}
	this.Name = name
	this.Expression = expression
	this.ScopeName = scopeName
	return &this
}

// NewScopeMappingRequestWithDefaults instantiates a new ScopeMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeMappingRequestWithDefaults() *ScopeMappingRequest {
	this := ScopeMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScopeMappingRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScopeMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *ScopeMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *ScopeMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *ScopeMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *ScopeMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value
func (o *ScopeMappingRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScopeMappingRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScopeMappingRequest) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value
func (o *ScopeMappingRequest) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *ScopeMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *ScopeMappingRequest) SetExpression(v string) {
	o.Expression = v
}

// GetScopeName returns the ScopeName field value
func (o *ScopeMappingRequest) GetScopeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScopeName
}

// GetScopeNameOk returns a tuple with the ScopeName field value
// and a boolean to check if the value has been set.
func (o *ScopeMappingRequest) GetScopeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeName, true
}

// SetScopeName sets field value
func (o *ScopeMappingRequest) SetScopeName(v string) {
	o.ScopeName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScopeMappingRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeMappingRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScopeMappingRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScopeMappingRequest) SetDescription(v string) {
	o.Description = &v
}

func (o ScopeMappingRequest) MarshalJSON() ([]byte, error) {
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
		toSerialize["scope_name"] = o.ScopeName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableScopeMappingRequest struct {
	value *ScopeMappingRequest
	isSet bool
}

func (v NullableScopeMappingRequest) Get() *ScopeMappingRequest {
	return v.value
}

func (v *NullableScopeMappingRequest) Set(val *ScopeMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeMappingRequest(val *ScopeMappingRequest) *NullableScopeMappingRequest {
	return &NullableScopeMappingRequest{value: val, isSet: true}
}

func (v NullableScopeMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
