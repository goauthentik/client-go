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

// PatchedScopeMappingRequest ScopeMapping Serializer
type PatchedScopeMappingRequest struct {
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed    NullableString `json:"managed,omitempty"`
	Name       *string        `json:"name,omitempty"`
	Expression *string        `json:"expression,omitempty"`
	// Scope name requested by the client
	ScopeName *string `json:"scope_name,omitempty"`
	// Description shown to the user when consenting. If left empty, the user won't be informed.
	Description *string `json:"description,omitempty"`
}

// NewPatchedScopeMappingRequest instantiates a new PatchedScopeMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedScopeMappingRequest() *PatchedScopeMappingRequest {
	this := PatchedScopeMappingRequest{}
	return &this
}

// NewPatchedScopeMappingRequestWithDefaults instantiates a new PatchedScopeMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedScopeMappingRequestWithDefaults() *PatchedScopeMappingRequest {
	this := PatchedScopeMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedScopeMappingRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedScopeMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *PatchedScopeMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *PatchedScopeMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *PatchedScopeMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *PatchedScopeMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedScopeMappingRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedScopeMappingRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedScopeMappingRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedScopeMappingRequest) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *PatchedScopeMappingRequest) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedScopeMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *PatchedScopeMappingRequest) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *PatchedScopeMappingRequest) SetExpression(v string) {
	o.Expression = &v
}

// GetScopeName returns the ScopeName field value if set, zero value otherwise.
func (o *PatchedScopeMappingRequest) GetScopeName() string {
	if o == nil || o.ScopeName == nil {
		var ret string
		return ret
	}
	return *o.ScopeName
}

// GetScopeNameOk returns a tuple with the ScopeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedScopeMappingRequest) GetScopeNameOk() (*string, bool) {
	if o == nil || o.ScopeName == nil {
		return nil, false
	}
	return o.ScopeName, true
}

// HasScopeName returns a boolean if a field has been set.
func (o *PatchedScopeMappingRequest) HasScopeName() bool {
	if o != nil && o.ScopeName != nil {
		return true
	}

	return false
}

// SetScopeName gets a reference to the given string and assigns it to the ScopeName field.
func (o *PatchedScopeMappingRequest) SetScopeName(v string) {
	o.ScopeName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedScopeMappingRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedScopeMappingRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedScopeMappingRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedScopeMappingRequest) SetDescription(v string) {
	o.Description = &v
}

func (o PatchedScopeMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Managed.IsSet() {
		toSerialize["managed"] = o.Managed.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.ScopeName != nil {
		toSerialize["scope_name"] = o.ScopeName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedScopeMappingRequest struct {
	value *PatchedScopeMappingRequest
	isSet bool
}

func (v NullablePatchedScopeMappingRequest) Get() *PatchedScopeMappingRequest {
	return v.value
}

func (v *NullablePatchedScopeMappingRequest) Set(val *PatchedScopeMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedScopeMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedScopeMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedScopeMappingRequest(val *PatchedScopeMappingRequest) *NullablePatchedScopeMappingRequest {
	return &NullablePatchedScopeMappingRequest{value: val, isSet: true}
}

func (v NullablePatchedScopeMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedScopeMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
