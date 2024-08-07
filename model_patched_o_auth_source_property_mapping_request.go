/*
authentik

Making authentication simple.

API version: 2024.6.3
Contact: hello@goauthentik.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PatchedOAuthSourcePropertyMappingRequest OAuthSourcePropertyMapping Serializer
type PatchedOAuthSourcePropertyMappingRequest struct {
	// Objects that are managed by authentik. These objects are created and updated automatically. This flag only indicates that an object can be overwritten by migrations. You can still modify the objects via the API, but expect changes to be overwritten in a later update.
	Managed    NullableString `json:"managed,omitempty"`
	Name       *string        `json:"name,omitempty"`
	Expression *string        `json:"expression,omitempty"`
}

// NewPatchedOAuthSourcePropertyMappingRequest instantiates a new PatchedOAuthSourcePropertyMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedOAuthSourcePropertyMappingRequest() *PatchedOAuthSourcePropertyMappingRequest {
	this := PatchedOAuthSourcePropertyMappingRequest{}
	return &this
}

// NewPatchedOAuthSourcePropertyMappingRequestWithDefaults instantiates a new PatchedOAuthSourcePropertyMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedOAuthSourcePropertyMappingRequestWithDefaults() *PatchedOAuthSourcePropertyMappingRequest {
	this := PatchedOAuthSourcePropertyMappingRequest{}
	return &this
}

// GetManaged returns the Managed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedOAuthSourcePropertyMappingRequest) GetManaged() string {
	if o == nil || o.Managed.Get() == nil {
		var ret string
		return ret
	}
	return *o.Managed.Get()
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedOAuthSourcePropertyMappingRequest) GetManagedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Managed.Get(), o.Managed.IsSet()
}

// HasManaged returns a boolean if a field has been set.
func (o *PatchedOAuthSourcePropertyMappingRequest) HasManaged() bool {
	if o != nil && o.Managed.IsSet() {
		return true
	}

	return false
}

// SetManaged gets a reference to the given NullableString and assigns it to the Managed field.
func (o *PatchedOAuthSourcePropertyMappingRequest) SetManaged(v string) {
	o.Managed.Set(&v)
}

// SetManagedNil sets the value for Managed to be an explicit nil
func (o *PatchedOAuthSourcePropertyMappingRequest) SetManagedNil() {
	o.Managed.Set(nil)
}

// UnsetManaged ensures that no value is present for Managed, not even an explicit nil
func (o *PatchedOAuthSourcePropertyMappingRequest) UnsetManaged() {
	o.Managed.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedOAuthSourcePropertyMappingRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourcePropertyMappingRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedOAuthSourcePropertyMappingRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedOAuthSourcePropertyMappingRequest) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *PatchedOAuthSourcePropertyMappingRequest) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedOAuthSourcePropertyMappingRequest) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *PatchedOAuthSourcePropertyMappingRequest) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *PatchedOAuthSourcePropertyMappingRequest) SetExpression(v string) {
	o.Expression = &v
}

func (o PatchedOAuthSourcePropertyMappingRequest) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullablePatchedOAuthSourcePropertyMappingRequest struct {
	value *PatchedOAuthSourcePropertyMappingRequest
	isSet bool
}

func (v NullablePatchedOAuthSourcePropertyMappingRequest) Get() *PatchedOAuthSourcePropertyMappingRequest {
	return v.value
}

func (v *NullablePatchedOAuthSourcePropertyMappingRequest) Set(val *PatchedOAuthSourcePropertyMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedOAuthSourcePropertyMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedOAuthSourcePropertyMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedOAuthSourcePropertyMappingRequest(val *PatchedOAuthSourcePropertyMappingRequest) *NullablePatchedOAuthSourcePropertyMappingRequest {
	return &NullablePatchedOAuthSourcePropertyMappingRequest{value: val, isSet: true}
}

func (v NullablePatchedOAuthSourcePropertyMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedOAuthSourcePropertyMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
