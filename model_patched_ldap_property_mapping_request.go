/*
authentik

Making authentication simple.

API version: 2021.10.1-rc1
Contact: hello@beryju.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedLDAPPropertyMappingRequest LDAP PropertyMapping Serializer
type PatchedLDAPPropertyMappingRequest struct {
	// Objects which are managed by authentik. These objects are created and updated automatically. This is flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed     NullableString `json:"managed,omitempty"`
	Name        *string        `json:"name,omitempty"`
	Expression  *string        `json:"expression,omitempty"`
	ObjectField *string        `json:"object_field,omitempty"`
}

// NewPatchedLDAPPropertyMappingRequest instantiates a new PatchedLDAPPropertyMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedLDAPPropertyMappingRequest() *PatchedLDAPPropertyMappingRequest {
	this := PatchedLDAPPropertyMappingRequest{}
	return &this
}

// NewPatchedLDAPPropertyMappingRequestWithDefaults instantiates a new PatchedLDAPPropertyMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedLDAPPropertyMappingRequestWithDefaults() *PatchedLDAPPropertyMappingRequest {
	this := PatchedLDAPPropertyMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedLDAPPropertyMappingRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedLDAPPropertyMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *PatchedLDAPPropertyMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *PatchedLDAPPropertyMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *PatchedLDAPPropertyMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *PatchedLDAPPropertyMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedLDAPPropertyMappingRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPPropertyMappingRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedLDAPPropertyMappingRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedLDAPPropertyMappingRequest) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *PatchedLDAPPropertyMappingRequest) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPPropertyMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *PatchedLDAPPropertyMappingRequest) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *PatchedLDAPPropertyMappingRequest) SetExpression(v string) {
	o.Expression = &v
}

// GetObjectField returns the ObjectField field value if set, zero value otherwise.
func (o *PatchedLDAPPropertyMappingRequest) GetObjectField() string {
	if o == nil || o.ObjectField == nil {
		var ret string
		return ret
	}
	return *o.ObjectField
}

// GetObjectFieldOk returns a tuple with the ObjectField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLDAPPropertyMappingRequest) GetObjectFieldOk() (*string, bool) {
	if o == nil || o.ObjectField == nil {
		return nil, false
	}
	return o.ObjectField, true
}

// HasObjectField returns a boolean if a field has been set.
func (o *PatchedLDAPPropertyMappingRequest) HasObjectField() bool {
	if o != nil && o.ObjectField != nil {
		return true
	}

	return false
}

// SetObjectField gets a reference to the given string and assigns it to the ObjectField field.
func (o *PatchedLDAPPropertyMappingRequest) SetObjectField(v string) {
	o.ObjectField = &v
}

func (o PatchedLDAPPropertyMappingRequest) MarshalJSON() ([]byte, error) {
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
	if o.ObjectField != nil {
		toSerialize["object_field"] = o.ObjectField
	}
	return json.Marshal(toSerialize)
}

type NullablePatchedLDAPPropertyMappingRequest struct {
	value *PatchedLDAPPropertyMappingRequest
	isSet bool
}

func (v NullablePatchedLDAPPropertyMappingRequest) Get() *PatchedLDAPPropertyMappingRequest {
	return v.value
}

func (v *NullablePatchedLDAPPropertyMappingRequest) Set(val *PatchedLDAPPropertyMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedLDAPPropertyMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedLDAPPropertyMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedLDAPPropertyMappingRequest(val *PatchedLDAPPropertyMappingRequest) *NullablePatchedLDAPPropertyMappingRequest {
	return &NullablePatchedLDAPPropertyMappingRequest{value: val, isSet: true}
}

func (v NullablePatchedLDAPPropertyMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedLDAPPropertyMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
